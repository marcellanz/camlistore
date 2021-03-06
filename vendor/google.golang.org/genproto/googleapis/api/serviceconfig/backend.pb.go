// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/backend.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Backend` defines the backend configuration for a service.
type Backend struct {
	// A list of backend rules providing configuration for individual API
	// elements.
	Rules []*BackendRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *Backend) Reset()                    { *m = Backend{} }
func (m *Backend) String() string            { return proto.CompactTextString(m) }
func (*Backend) ProtoMessage()               {}
func (*Backend) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Backend) GetRules() []*BackendRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A backend rule provides configuration for an individual API element.
type BackendRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// The address of the API backend.
	//
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	// The number of seconds to wait for a response from a request.  The
	// default depends on the deployment context.
	Deadline float64 `protobuf:"fixed64,3,opt,name=deadline" json:"deadline,omitempty"`
}

func (m *BackendRule) Reset()                    { *m = BackendRule{} }
func (m *BackendRule) String() string            { return proto.CompactTextString(m) }
func (*BackendRule) ProtoMessage()               {}
func (*BackendRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*Backend)(nil), "google.api.Backend")
	proto.RegisterType((*BackendRule)(nil), "google.api.BackendRule")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/backend.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x3d, 0x4f, 0x04, 0x21,
	0x10, 0x86, 0x83, 0x17, 0x3d, 0x9d, 0x33, 0x16, 0x34, 0x12, 0x2b, 0x72, 0xd5, 0x36, 0x42, 0xa2,
	0x8d, 0xf5, 0x26, 0xf6, 0x17, 0xfe, 0x80, 0xe1, 0x60, 0x24, 0x44, 0x64, 0x08, 0xdc, 0xfa, 0xfb,
	0xcd, 0x7e, 0xb8, 0x6e, 0xf9, 0xf2, 0x3c, 0xcc, 0xcc, 0x0b, 0xef, 0x81, 0x28, 0x24, 0x54, 0x81,
	0x92, 0xcd, 0x41, 0x51, 0x0d, 0x3a, 0x60, 0x2e, 0x95, 0x2e, 0xa4, 0x67, 0x64, 0x4b, 0x6c, 0xda,
	0x96, 0xa8, 0x1b, 0xd6, 0x9f, 0xe8, 0xd0, 0x51, 0xfe, 0x8c, 0x41, 0x9f, 0xad, 0xfb, 0xc2, 0xec,
	0xd5, 0xa4, 0x72, 0x58, 0xc6, 0xd8, 0x12, 0x8f, 0x6f, 0xb0, 0xef, 0x67, 0xc8, 0x9f, 0xe1, 0xba,
	0x0e, 0x09, 0x9b, 0x60, 0x72, 0xd7, 0x1d, 0x5e, 0x1e, 0xd5, 0xbf, 0xa6, 0x16, 0xc7, 0x0c, 0x09,
	0xcd, 0x6c, 0x1d, 0x3f, 0xe0, 0xb0, 0x79, 0xe5, 0x4f, 0x70, 0xdb, 0x30, 0xa1, 0xbb, 0x50, 0x15,
	0x4c, 0xb2, 0xee, 0xce, 0xac, 0x99, 0x0b, 0xd8, 0x5b, 0xef, 0x2b, 0xb6, 0x26, 0xae, 0x26, 0xf4,
	0x17, 0xc7, 0x5f, 0x1e, 0xad, 0x4f, 0x31, 0xa3, 0xd8, 0x49, 0xd6, 0x31, 0xb3, 0xe6, 0x5e, 0xc2,
	0x83, 0xa3, 0xef, 0xcd, 0x15, 0xfd, 0xfd, 0xb2, 0xf0, 0x34, 0xd6, 0x38, 0xb1, 0xf3, 0xcd, 0xd4,
	0xe7, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x99, 0xf6, 0x26, 0x9c, 0x18, 0x01, 0x00, 0x00,
}
