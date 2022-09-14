// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mapping/testdata/schema/with_complex_id.proto

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

func (this *EntityWithComplexId) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.SomeDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SomeDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SomeDate", err)
		}
	}
	return nil
}
func (this *EntityComplexId) Validate() error {
	if this.IdPart3 != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IdPart3); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IdPart3", err)
		}
	}
	return nil
}
