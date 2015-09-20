// Code generated by protoc-gen-go.
// source: pfs/route/route.proto
// DO NOT EDIT!

/*
Package route is a generated protocol buffer package.

It is generated from these files:
	pfs/route/route.proto

It has these top-level messages:
	ServerState
	ServerRole
*/
package route

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerState struct {
	Id      string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address string          `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Version int64           `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Shards  map[uint64]bool `protobuf:"bytes,4,rep,name=shards" json:"shards,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ServerState) Reset()         { *m = ServerState{} }
func (m *ServerState) String() string { return proto.CompactTextString(m) }
func (*ServerState) ProtoMessage()    {}

func (m *ServerState) GetShards() map[uint64]bool {
	if m != nil {
		return m.Shards
	}
	return nil
}

type ServerRole struct {
	Version  int64           `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Masters  map[uint64]bool `protobuf:"bytes,2,rep,name=masters" json:"masters,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Replicas map[uint64]bool `protobuf:"bytes,3,rep,name=replicas" json:"replicas,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ServerRole) Reset()         { *m = ServerRole{} }
func (m *ServerRole) String() string { return proto.CompactTextString(m) }
func (*ServerRole) ProtoMessage()    {}

func (m *ServerRole) GetMasters() map[uint64]bool {
	if m != nil {
		return m.Masters
	}
	return nil
}

func (m *ServerRole) GetReplicas() map[uint64]bool {
	if m != nil {
		return m.Replicas
	}
	return nil
}
