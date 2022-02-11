// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: apm_diagnotor.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// DiagnotorServiceHandler is the server API for DiagnotorService service.
type DiagnotorServiceHandler interface {
	// GET /api/msp/diagnosis/{terminusKey}/services
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// POST /api/msp/diagnosis/{terminusKey}
	StartDiagnosis(context.Context, *StartDiagnosisRequest) (*StartDiagnosisResponse, error)
	// GET /api/msp/diagnosis/{terminusKey}/status
	QueryDiagnosisStatus(context.Context, *QueryDiagnosisStatusRequest) (*QueryDiagnosisStatusResponse, error)
	// DELETE /api/msp/diagnosis/{terminusKey}
	StopDiagnosis(context.Context, *StopDiagnosisRequest) (*StopDiagnosisResponse, error)
	// GET /api/msp/diagnosis/{terminusKey}/processes
	ListProcesses(context.Context, *ListProcessesRequest) (*ListProcessesResponse, error)
}

// RegisterDiagnotorServiceHandler register DiagnotorServiceHandler to http.Router.
func RegisterDiagnotorServiceHandler(r http.Router, srv DiagnotorServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_ListServices := func(method, path string, fn func(context.Context, *ListServicesRequest) (*ListServicesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListServicesRequest))
		}
		var ListServices_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListServices_info = transport.NewServiceInfo("erda.msp.apm.diagnotor.DiagnotorService", "ListServices", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListServices_info)
				}
				r = r.WithContext(ctx)
				var in ListServicesRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "terminusKey":
							in.TerminusKey = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_StartDiagnosis := func(method, path string, fn func(context.Context, *StartDiagnosisRequest) (*StartDiagnosisResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*StartDiagnosisRequest))
		}
		var StartDiagnosis_info transport.ServiceInfo
		if h.Interceptor != nil {
			StartDiagnosis_info = transport.NewServiceInfo("erda.msp.apm.diagnotor.DiagnotorService", "StartDiagnosis", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, StartDiagnosis_info)
				}
				r = r.WithContext(ctx)
				var in StartDiagnosisRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "terminusKey":
							in.TerminusKey = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_QueryDiagnosisStatus := func(method, path string, fn func(context.Context, *QueryDiagnosisStatusRequest) (*QueryDiagnosisStatusResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueryDiagnosisStatusRequest))
		}
		var QueryDiagnosisStatus_info transport.ServiceInfo
		if h.Interceptor != nil {
			QueryDiagnosisStatus_info = transport.NewServiceInfo("erda.msp.apm.diagnotor.DiagnotorService", "QueryDiagnosisStatus", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, QueryDiagnosisStatus_info)
				}
				r = r.WithContext(ctx)
				var in QueryDiagnosisStatusRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "terminusKey":
							in.TerminusKey = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_StopDiagnosis := func(method, path string, fn func(context.Context, *StopDiagnosisRequest) (*StopDiagnosisResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*StopDiagnosisRequest))
		}
		var StopDiagnosis_info transport.ServiceInfo
		if h.Interceptor != nil {
			StopDiagnosis_info = transport.NewServiceInfo("erda.msp.apm.diagnotor.DiagnotorService", "StopDiagnosis", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, StopDiagnosis_info)
				}
				r = r.WithContext(ctx)
				var in StopDiagnosisRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "terminusKey":
							in.TerminusKey = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListProcesses := func(method, path string, fn func(context.Context, *ListProcessesRequest) (*ListProcessesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListProcessesRequest))
		}
		var ListProcesses_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListProcesses_info = transport.NewServiceInfo("erda.msp.apm.diagnotor.DiagnotorService", "ListProcesses", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListProcesses_info)
				}
				r = r.WithContext(ctx)
				var in ListProcessesRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "terminusKey":
							in.TerminusKey = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListServices("GET", "/api/msp/diagnosis/{terminusKey}/services", srv.ListServices)
	add_StartDiagnosis("POST", "/api/msp/diagnosis/{terminusKey}", srv.StartDiagnosis)
	add_QueryDiagnosisStatus("GET", "/api/msp/diagnosis/{terminusKey}/status", srv.QueryDiagnosisStatus)
	add_StopDiagnosis("DELETE", "/api/msp/diagnosis/{terminusKey}", srv.StopDiagnosis)
	add_ListProcesses("GET", "/api/msp/diagnosis/{terminusKey}/processes", srv.ListProcesses)
}
