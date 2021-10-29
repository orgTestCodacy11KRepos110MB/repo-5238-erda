// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: event_query.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetEventsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetEventsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetEventsResult)(nil)

// GetEventsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetEventsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "eventId":
				m.EventId = vals[0]
			case "traceId":
				m.TraceId = vals[0]
			case "relationId":
				m.RelationId = vals[0]
			case "relationType":
				m.RelationType = vals[0]
			case "start":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Start = val
			case "end":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.End = val
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "debug":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Debug = val
			}
		}
	}
	return nil
}

// GetEventsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetEventsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &GetEventsResult{}
				}
			}
		}
	}
	return nil
}

// GetEventsResult implement urlenc.URLValuesUnmarshaler.
func (m *GetEventsResult) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
