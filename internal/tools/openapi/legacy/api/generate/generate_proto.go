// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/internal/tools/openapi/legacy/api/apis"
	"github.com/erda-project/erda/pkg/strutil"
)

const protoBaseDir = "openapiv1"

type protoFile struct {
	io.Writer
}

func (pf *protoFile) writeline(s string) {
	if _, err := io.WriteString(pf, s+"\n"); err != nil {
		panic(fmt.Errorf("failed to write string: %s, err: %v", s, err))
	}
}

func (pf *protoFile) prefixWriteline(prefix string, s string) {
	if _, err := io.WriteString(pf, prefix); err != nil {
		panic(fmt.Errorf("failed to write string: %s, err: %v", s, err))
	}
	pf.writeline(s)
}

func (pf *protoFile) newline(n int) {
	for i := 0; i < n; i++ {
		pf.writeline("")
	}
}

type wrappedApiSpec struct {
	FullAPIPath    string
	APIVarName     string
	API            apis.ApiSpec
	Belongs        string // testplatform_autotest
	CompName       string // testplatform
	SubDirs        []string
	ReqMissingVars []string // vars missing in request, but needed in http path
}

func generateProto() {
	// group by APIs
	groupedAPIs := make(map[string][]*wrappedApiSpec) // key1: belongs, key2: apiName
	for i, apiName := range APINames {
		if !strings.Contains(apiName, ".") {
			continue
		}
		belongs := strings.SplitN(apiName, ".", 2)[0]
		compName, subDirs := getCompAndSubDirsByAPIName(apiName)
		groupedAPIs[belongs] = append(groupedAPIs[belongs], &wrappedApiSpec{
			APIVarName:  getAPIVarNameByAPIName(apiName),
			FullAPIPath: apiName,
			API:         APIs[i],
			Belongs:     belongs,
			CompName:    compName,
			SubDirs:     subDirs,
		})
	}

	for belongs, wAPIs := range groupedAPIs {
		groupMessages := make(map[string]*messageV2)
		for _, wAPI := range wAPIs {
			logrus.Infof("API: %s, Request: %v", wAPI.FullAPIPath, reflect.TypeOf(wAPI.API.RequestType))
			//polishRequestTypeBeforeParse(groupMessages, wAPI)
			createEmbedMessage(groupMessages, reflect.TypeOf(wAPI.API.RequestType))

			logrus.Infof("API: %s, Response: %v", wAPI.FullAPIPath, reflect.TypeOf(wAPI.API.ResponseType))
			//polishResponseTypeBeforeParse(wAPI)
			createEmbedMessage(groupMessages, reflect.TypeOf(wAPI.API.ResponseType))

			postHandleMessages(groupMessages, wAPI)
		}
		for k := range groupMessages {
			fmt.Println(k)
		}
		targetProtoPath := filepath.Join("../../../../../../api/proto", protoBaseDir, belongs, getCompNameByAPIName(belongs)+".proto")
		if err := os.MkdirAll(filepath.Dir(targetProtoPath), 0755); err != nil {
			panic(fmt.Errorf("failed to create dir: %s, err: %v", targetProtoPath, err))
		}
		w, err := os.OpenFile(targetProtoPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		f := &protoFile{Writer: w}
		f.writeProtoHeader(belongs)
		writeMessages(f, groupMessages)
		f.writeProtoServices(groupMessages, belongs, wAPIs)
	}
}

//func polishResponseTypeBeforeParse(wAPI *wrappedApiSpec) {
//	respType := reflect.TypeOf(wAPI.API.ResponseType)
//	if respType != nil && respType.Kind() == reflect.Struct {
//		for i := 0; i < respType.NumField(); i++ {
//			f := respType.Field(i)
//			if getJsonTagValue(f.Tag) == "data" {
//				wAPI.API.ResponseType = reflect.Indirect(reflect.New(f.Type)).Interface()
//				break
//			}
//		}
//	}
//}

func postHandleMessages(messages map[string]*messageV2, wAPI *wrappedApiSpec) {
	// add missing vars for message
	pathVarReg := regexp.MustCompile(`<([^<>*]+)>`)
	wAPI.API.Path = strutil.ReplaceAllStringSubmatchFunc(pathVarReg, wAPI.API.Path, func(ss []string) string {
		return ensureVarName(ss[0])
	})
	wAPI.API.BackendPath = strutil.ReplaceAllStringSubmatchFunc(pathVarReg, wAPI.API.Path, func(ss []string) string {
		return ensureVarName(ss[0])
	})
	allVars := pathVarReg.FindAllStringSubmatch(wAPI.API.Path, -1)
	t := reflect.TypeOf(wAPI.API.RequestType)
	var reqMsg *messageV2
	if t != nil {
		reqMsg = messages[t.Name()]
	}
	if t == nil || reqMsg == nil {
		// create messages and add missing vars
		var fields []messageFieldV2
		for _, vars := range allVars {
			varName := ensureVarName(vars[1])
			fields = append(fields, messageFieldV2{
				Name:       varName,
				ProtoType:  "string",
				IsRepeated: false,
				JsonTag:    "",
				Comment:    fmt.Sprintf("generated from path variable: %s. You should change the proto type if necessary.", varName),
			})
		}
		msg := messageV2{Name: makeDefaultMsgRequestName(wAPI), Fields: fields}
		messages[msg.Name] = &msg
		return
	}
	// add missing vars for message
	reqMsgFields := make(map[string]struct{})
	for _, f := range reqMsg.Fields {
		reqMsgFields[strings.ToLower(f.Name)] = struct{}{}
	}
	for _, vars := range allVars {
		varName := ensureVarName(vars[1])
		if _, ok := reqMsgFields[strings.ToLower(varName)]; !ok {
			reqMsg.Fields = append(reqMsg.Fields, messageFieldV2{
				Name:       varName,
				ProtoType:  "string", // if use "google.protobuf.Value", will cause "--go-http_out: not support type "message" for query string"
				IsRepeated: false,
				JsonTag:    "",
				Comment:    fmt.Sprintf("generated from path variable: %s. You should change the proto type if necessary.", varName),
			})
		}
	}
}

func makeDefaultMsgRequestName(wAPI *wrappedApiSpec) string {
	return wAPI.APIVarName + "_Request"
}
func makeDefaultMsgResponseName(wAPI *wrappedApiSpec) string {
	return wAPI.APIVarName + "_Response"
}
func ensureVarName(varName string) string {
	return strings.NewReplacer(" ", "_", "-", "_", "_", "").Replace(varName)
}

func writeMessages(w io.Writer, messages map[string]*messageV2) {
	for _, msg := range messages {
		io.WriteString(w, fmt.Sprintf(`message %s {
`, msg.Name))
		for i, f := range msg.Fields {
			if strings.HasPrefix(f.Name, "ID") {
				// do nothing
			} else {
				f.Name = lowerFirstChar(f.Name)
			}
			var repeatedStr string
			if f.IsRepeated {
				repeatedStr = "repeated "
			}
			var jsonTagString string
			if f.JsonTag != "" && f.JsonTag != f.Name {
				jsonTagString = fmt.Sprintf(` [json_name = "%s"]`, f.JsonTag)
			}
			protoType := f.ProtoType
			var commentStr string
			if f.Comment != "" {
				commentStr = " // " + f.Comment
			}
			line := fmt.Sprintf(`    %s%s %s = %d%s;%s
`, repeatedStr, protoType, f.Name, i+1, jsonTagString, commentStr)
			// cannot use repeated & optional together
			line = strings.NewReplacer("repeated optional", "repeated").Replace(line)
			// cannot use optional inside map
			line = strings.NewReplacer("map<optional ", "map<", ", optional", ", ", ", map>", ", google.protobuf.Value>").Replace(line)
			line = strings.NewReplacer("repeated map", "repeated google.protobuf.Value").Replace(line)

			io.WriteString(w, line)
		}
		io.WriteString(w, `}
`)
	}
}

const TAB = `    `

func (pf *protoFile) writeProtoServices(messages map[string]*messageV2, belongs string, wAPIs []*wrappedApiSpec) {
	if len(wAPIs) == 0 {
		return
	}
	oneWAPI := wAPIs[0]
	pf.writeline(fmt.Sprintf(`// generate service from openapi spec: %s`, belongs))
	pf.writeline(fmt.Sprintf(`service %s {`, oneWAPI.Belongs))

	// openapi option
	pf.prefixWriteline(TAB, fmt.Sprintf(`option (erda.common.openapi_service) = {`))

	pf.prefixWriteline(TAB+TAB, fmt.Sprintf(`service: "%s",`, oneWAPI.CompName))

	pf.prefixWriteline(TAB, fmt.Sprintf(`};`))
	pf.newline(1)

	// rpc s
	for _, wAPI := range wAPIs {
		pf.prefixWriteline(TAB, fmt.Sprintf(`rpc %s (%s) returns (%s) {`, wAPI.APIVarName, getRpcRequestName(messages, wAPI), getRpcResponseName(messages, wAPI)))

		pf.prefixWriteline(TAB+TAB, fmt.Sprintf(`option (google.api.http) = {`))
		if wAPI.API.Method == http.MethodHead {
			pf.prefixWriteline(TAB+TAB+TAB, fmt.Sprintf(`custom: {`))
			pf.prefixWriteline(TAB+TAB+TAB+TAB, fmt.Sprintf(`kind: "head",`))
			pf.prefixWriteline(TAB+TAB+TAB+TAB, fmt.Sprintf(`path: "%s",`, replaceOpenapiV1Path(wAPI.API.BackendPath)))
			pf.prefixWriteline(TAB+TAB+TAB, fmt.Sprintf(`},`))
		} else {
			pf.prefixWriteline(TAB+TAB+TAB, fmt.Sprintf(`%s: "%s",`, strings.ToLower(wAPI.API.Method), replaceOpenapiV1Path(wAPI.API.BackendPath)))
		}
		pf.prefixWriteline(TAB+TAB, fmt.Sprintf(`};`))

		pf.prefixWriteline(TAB+TAB, fmt.Sprintf(`option (erda.common.openapi) = {`))
		pf.prefixWriteline(TAB+TAB+TAB, fmt.Sprintf(`path: "%s",`, replaceOpenapiV1Path(wAPI.API.Path)))
		pf.prefixWriteline(TAB+TAB+TAB, fmt.Sprintf(`auth: {`))
		pf.prefixWriteline(TAB+TAB+TAB+TAB, fmt.Sprintf(`no_check: %t,`, !wAPI.API.CheckBasicAuth && !wAPI.API.CheckLogin && !wAPI.API.CheckToken && !wAPI.API.TryCheckLogin))
		pf.prefixWriteline(TAB+TAB+TAB+TAB, fmt.Sprintf(`check_login: %t,`, wAPI.API.CheckLogin))
		pf.prefixWriteline(TAB+TAB+TAB+TAB, fmt.Sprintf(`try_check_login: %t,`, wAPI.API.TryCheckLogin))
		pf.prefixWriteline(TAB+TAB+TAB+TAB, fmt.Sprintf(`check_token: %t,`, wAPI.API.CheckToken))
		pf.prefixWriteline(TAB+TAB+TAB+TAB, fmt.Sprintf(`check_basic_auth: %t,`, wAPI.API.CheckBasicAuth))
		pf.prefixWriteline(TAB+TAB+TAB, fmt.Sprintf(`},`))
		pf.prefixWriteline(TAB+TAB+TAB, fmt.Sprintf(`doc: %q,`, strings.NewReplacer("\n", "").Replace(wAPI.API.Doc)))
		pf.prefixWriteline(TAB+TAB, fmt.Sprintf(`};`))

		pf.prefixWriteline(TAB, fmt.Sprintf(`};`))
	}

	pf.writeline(`}`)
}

func getRpcRequestName(messages map[string]*messageV2, wAPI *wrappedApiSpec) string {
	var reqMessageName string
	if wAPI.API.RequestType == nil {
		// try message type
		reqMessageName = makeDefaultMsgRequestName(wAPI)
	} else {
		reqMessageName = reflect.TypeOf(wAPI.API.RequestType).Name()
	}
	if _, ok := messages[reqMessageName]; ok {
		return reqMessageName
	}
	return "google.protobuf.Empty"
}
func getRpcResponseName(messages map[string]*messageV2, wAPI *wrappedApiSpec) string {
	var respMessageName string
	if wAPI.API.ResponseType == nil {
		// try message type
		respMessageName = makeDefaultMsgResponseName(wAPI)
	} else {
		respMessageName = reflect.TypeOf(wAPI.API.ResponseType).Name()
	}
	if _, ok := messages[respMessageName]; ok {
		return respMessageName
	}
	return "google.protobuf.Empty"
}

func lowerFirstChar(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func getAPIVarNameByAPIName(apiName string) string {
	ss := strings.Split(apiName, ".")
	varName := ss[len(ss)-1]
	return varName
}

func (pf *protoFile) writeProtoHeader(apiName string) {
	pf.writeline("// generated by openapi-gen-protobuf")
	pf.newline(1)

	pf.writeline(`syntax = "proto3";`)
	pf.newline(1)

	pf.writeline("package " + getProtoPackageByAPIName(apiName) + ";")
	pf.newline(1)

	pf.writeline(`option go_package = "` + getProtoGoPackageByAPIName(apiName) + `";`)
	pf.newline(1)

	pf.writeline(`import "google/api/annotations.proto";`)
	pf.writeline(`import "google/protobuf/empty.proto";`)
	pf.writeline(`import "google/protobuf/struct.proto";`)
	pf.writeline(`import "google/protobuf/timestamp.proto";`)
	pf.writeline(`import "google/protobuf/duration.proto";`)
	pf.writeline(`import "github.com/envoyproxy/protoc-gen-validate/validate/validate.proto";`)
	pf.newline(1)
	pf.writeline(`import "common/openapi.proto";`)
	pf.writeline(`import "common/identity.proto";`)
	pf.newline(1)
}

func getCompNameByAPIName(apiName string) string {
	compName, _ := getCompAndSubDirsByAPIName(apiName)
	// TODO mapping
	return compName
}

func getCompAndSubDirsByAPIName(apiName string) (compName string, subDirs []string) {
	compWithSubDirs := strings.SplitN(apiName, ".", 2)[0]
	ss := strings.Split(compWithSubDirs, "_")
	compName = ss[0]
	if len(ss) > 1 {
		subDirs = ss[1:]
	}
	return
}

func getProtoPackageByAPIName(apiName string) string {
	compName := getCompNameByAPIName(apiName)
	return "erda." + compName
}

func getProtoGoPackageByAPIName(apiName string) string {
	compName, subDirs := getCompAndSubDirsByAPIName(apiName)
	return filepath.Join("github.com/erda-project/erda-proto-go", protoBaseDir, compName, strutil.Join(subDirs, "/"), "pb")
}

func replaceOpenapiV1Path(path string) string {
	path = strings.ReplaceAll(path, "<*>", "**")
	newPath := strings.NewReplacer("<", "{", ">", "}", " ", "_").Replace(path)
	return newPath
}

type messageV2 struct {
	Name   string
	Fields []messageFieldV2
}

type messageFieldV2 struct {
	Name       string
	ProtoType  string
	IsRepeated bool
	JsonTag    string
	Comment    string
}

func createEmbedMessage(messages map[string]*messageV2, t reflect.Type) {
	if t == nil {
		return
	}
	// nested apistructs type doesn't have pkg path
	switch t.Kind() {
	case reflect.Slice:
		t = t.Elem()
	case reflect.Map:
		createEmbedMessage(messages, t.Key())
		createEmbedMessage(messages, t.Elem())
	}
	if !strutil.Contains(t.PkgPath(), "github.com/erda-project/erda") {
		return
	}
	switch t.Kind() {
	case reflect.Interface: // A interface{}
		return
	case reflect.Map: // A map[key]value
		createEmbedMessage(messages, t.Key())
		createEmbedMessage(messages, t.Elem())
	case reflect.Ptr: // A *type
		createEmbedMessage(messages, t.Elem())
	case reflect.Slice:
		createEmbedMessage(messages, t.Elem())
	case reflect.Struct:
		if _, ok := messages[t.Name()]; ok {
			return
		}
		var fields []messageFieldV2
		for i := 0; i < t.NumField(); i++ {
			rf := t.Field(i)
			f := makeMessageField(messages, rf)
			if f != nil {
				fields = append(fields, *f)
			}
		}
		messages[t.Name()] = &messageV2{
			Name:   t.Name(),
			Fields: fields,
		}
	default:
		return
	}
}

func makeMessageField(messages map[string]*messageV2, rf reflect.StructField) *messageFieldV2 {
	if rf.PkgPath != "" {
		return nil // ignore unexported field
	}
	// ignore some fields
	switch rf.Type {
	case reflect.TypeOf(apistructs.Header{}),
		reflect.TypeOf(apistructs.UserInfoHeader{}),
		reflect.TypeOf(apistructs.IdentityInfo{}):
		return nil
	}
	createEmbedMessage(messages, rf.Type)
	f := messageFieldV2{
		Name:       rf.Name,
		IsRepeated: rf.Type.Kind() == reflect.Slice,
		JsonTag:    getJsonTagValue(rf.Tag),
	}
	switch rf.Type.Kind() {
	case reflect.Map:
		f.ProtoType = fmt.Sprintf("map<%s, %s>", getProtoTypeByGoType(messages, rf.Type.Key()), getProtoTypeByGoType(messages, rf.Type.Elem()))
	case reflect.Slice:
		f.ProtoType = getProtoTypeByGoType(messages, rf.Type.Elem())
	default:
		f.ProtoType = getProtoTypeByGoType(messages, rf.Type)
	}
	return &f
}

func getJsonTagValue(tag reflect.StructTag) string {
	v := tag.Get("json")
	if v == "" {
		v = tag.Get("path")
	}
	return strings.SplitN(v, ",", 2)[0]
}

func getProtoTypeByGoType(messages map[string]*messageV2, t reflect.Type) string {
	goTypeStr := t.String()
	optional := false
	if strings.HasPrefix(goTypeStr, "*") {
		goTypeStr = goTypeStr[1:]
		optional = true
	}
	protoTypeStr, ok := goType2ProtoTypeMapping[goTypeStr]
	if !ok {
		// try alias
		alias := t.Kind().String()
		aliasProtoTypeStr, ok := goType2ProtoTypeMapping[alias]
		if ok {
			protoTypeStr = aliasProtoTypeStr
		} else {
			// if exist in messages, use message type
			if _, ok := messages[t.String()]; ok {
				protoTypeStr = t.Name()
			}
		}
	}
	// fallback
	if protoTypeStr == "" {
		protoTypeStr = "google.protobuf.Value"
	}
	if optional {
		return "optional " + protoTypeStr
	}
	return protoTypeStr
}

var goType2ProtoTypeMapping = map[string]string{
	"float64":       "double",
	"float32":       "float",
	"int32":         "int32",
	"int64":         "int64",
	"uint32":        "uint32",
	"uint64":        "uint64",
	"bool":          "bool",
	"string":        "string",
	"uint8":         "uint32",
	"int8":          "int32",
	"byte":          "uint32",
	"uint16":        "uint32",
	"int16":         "int32",
	"int":           "int64",
	"uint":          "uint64",
	"uintptr":       "uint64",
	"rune":          "int32",
	"time.Time":     "google.protobuf.Timestamp",
	"time.Duration": "google.protobuf.Duration",
	"map":           "map",
	"interface {}":  "google.protobuf.Value",
}
