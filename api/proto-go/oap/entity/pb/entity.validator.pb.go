// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Entity) Validate() error {
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	if this.Key == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Key", fmt.Errorf(`value '%v' must not be an empty string`, this.Key))
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SetEntityRequest) Validate() error {
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SetEntityResponse) Validate() error {
	return nil
}
func (this *RemoveEntityRequest) Validate() error {
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	if this.Key == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Key", fmt.Errorf(`value '%v' must not be an empty string`, this.Key))
	}
	return nil
}
func (this *RemoveEntityResponse) Validate() error {
	return nil
}
func (this *GetEntityRequest) Validate() error {
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	if this.Key == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Key", fmt.Errorf(`value '%v' must not be an empty string`, this.Key))
	}
	return nil
}
func (this *GetEntityResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListEntitiesRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ListEntitiesResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *EntityList) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}