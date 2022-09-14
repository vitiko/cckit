// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cpaper_extended/schema/state.proto

package schema

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CommercialPaper) Validate() error {
	if this.IssueDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IssueDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IssueDate", err)
		}
	}
	if this.MaturityDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MaturityDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MaturityDate", err)
		}
	}
	return nil
}
func (this *CommercialPaperId) Validate() error {
	return nil
}
func (this *CommercialPaperList) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
