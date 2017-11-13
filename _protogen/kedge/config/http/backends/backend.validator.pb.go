// Code generated by protoc-gen-gogo.
// source: kedge/config/http/backends/backend.proto
// DO NOT EDIT!

/*
Package kedge_config_http_backends is a generated protocol buffer package.

It is generated from these files:
	kedge/config/http/backends/backend.proto

It has these top-level messages:
	Backend
	Middleware
	Security
*/
package kedge_config_http_backends

import regexp "regexp"
import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"
import  _ "github.com/improbable-eng/kedge/_protogen/kedge/config/common/resolvers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_Backend_Name = regexp.MustCompile("^[a-z_0-9.]{2,64}$")

func (this *Backend) Validate() error {
	if !_regex_Backend_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z_0-9.]{2,64}$"`, this.Name))
	}
	if this.Security != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Security); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Security", err)
		}
	}
	if oneOfNester, ok := this.GetResolver().(*Backend_Srv); ok {
		if oneOfNester.Srv != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Srv); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Srv", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResolver().(*Backend_K8S); ok {
		if oneOfNester.K8S != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.K8S); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("K8S", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResolver().(*Backend_Host); ok {
		if oneOfNester.Host != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Host); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Host", err)
			}
		}
	}
	return nil
}
func (this *Middleware) Validate() error {
	if oneOfNester, ok := this.GetMiddleware().(*Middleware_Retry_); ok {
		if oneOfNester.Retry != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Retry); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Retry", err)
			}
		}
	}
	return nil
}
func (this *Middleware_Retry) Validate() error {
	return nil
}
func (this *Security) Validate() error {
	return nil
}
