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

package dicehub

import (
	"net/http"

	"github.com/erda-project/erda/internal/core/openapi/legacy/api/apis"
)

var RELEASE_RULE_UPDATE = apis.ApiSpec{
	Path:        "/api/release-rules/<id>",
	BackendPath: "/api/release-rules/<id>",
	Host:        "erda-server.marathon.l4lb.thisdcos.directory:9095",
	Scheme:      "http",
	Method:      http.MethodPut,
	CheckLogin:  true,
	CheckToken:  true,
	IsOpenAPI:   true,
	Doc:         `修改制品规则`,
}