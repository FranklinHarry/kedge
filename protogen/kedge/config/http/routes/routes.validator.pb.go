// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kedge/config/http/routes/routes.proto

/*
Package kedge_config_http_routes is a generated protocol buffer package.

It is generated from these files:
	kedge/config/http/routes/routes.proto

It has these top-level messages:
	Route
*/
package kedge_config_http_routes

import regexp "regexp"
import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_Route_BackendName = regexp.MustCompile(`^[a-z_0-9.]{2,64}$`)

func (this *Route) Validate() error {
	if !_regex_Route_BackendName.MatchString(this.BackendName) {
		return go_proto_validators.FieldError("BackendName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z_0-9.]{2,64}$"`, this.BackendName))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
