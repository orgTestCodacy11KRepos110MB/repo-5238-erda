// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notify.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateNotifyHistoryRequest) Validate() error {
	for _, item := range this.NotifyTargets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NotifyTargets", err)
			}
		}
	}
	if this.NotifySource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NotifySource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NotifySource", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *NotifyTarget) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *Target) Validate() error {
	return nil
}
func (this *NotifySource) Validate() error {
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *CreateNotifyHistoryResponse) Validate() error {
	return nil
}
func (this *QueryNotifyHistoriesRequest) Validate() error {
	return nil
}
func (this *QueryNotifyHistoriesResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QueryNotifyHistoryData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *NotifyHistory) Validate() error {
	for _, item := range this.NotifyTargets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NotifyTargets", err)
			}
		}
	}
	if this.NotifySource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NotifySource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NotifySource", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *GetNotifyStatusRequest) Validate() error {
	return nil
}
func (this *GetNotifyStatusResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetNotifyHistogramRequest) Validate() error {
	return nil
}
func (this *GetNotifyHistogramResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *NotifyHistogramData) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *StatisticValue) Validate() error {
	return nil
}
func (this *QueryAlertNotifyHistoriesRequest) Validate() error {
	return nil
}
func (this *QueryAlertNotifyHistoriesResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AlertNotifyHistories) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *AlertNotifyIndex) Validate() error {
	if this.SendTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SendTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SendTime", err)
		}
	}
	return nil
}
func (this *GetAlertNotifyDetailRequest) Validate() error {
	return nil
}
func (this *GetAlertNotifyDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AlertNotifyDetail) Validate() error {
	if this.SendTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SendTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SendTime", err)
		}
	}
	return nil
}
func (this *GetTypeNotifyHistogramRequest) Validate() error {
	return nil
}
func (this *GetTypeNotifyHistogramResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TypeNotifyHistogram) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}