// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/pkg/work/work.proto

package work

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type State int32

const (
	State_RUNNING State = 0
	State_SUCCESS State = 1
	State_FAILURE State = 2
)

var State_name = map[int32]string{
	0: "RUNNING",
	1: "SUCCESS",
	2: "FAILURE",
}

var State_value = map[string]int32{
	"RUNNING": 0,
	"SUCCESS": 1,
	"FAILURE": 2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_58a68e4647f78187, []int{0}
}

type Task struct {
	ID                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *types.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Subtasks             []*Task    `protobuf:"bytes,3,rep,name=subtasks,proto3" json:"subtasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a68e4647f78187, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Task.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return m.Size()
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Task) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Task) GetSubtasks() []*Task {
	if m != nil {
		return m.Subtasks
	}
	return nil
}

type TaskInfo struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	State                State    `protobuf:"varint,2,opt,name=state,proto3,enum=work.State" json:"state,omitempty"`
	Reason               string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskInfo) Reset()         { *m = TaskInfo{} }
func (m *TaskInfo) String() string { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()    {}
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_58a68e4647f78187, []int{1}
}
func (m *TaskInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInfo.Merge(m, src)
}
func (m *TaskInfo) XXX_Size() int {
	return m.Size()
}
func (m *TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInfo proto.InternalMessageInfo

func (m *TaskInfo) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *TaskInfo) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RUNNING
}

func (m *TaskInfo) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterEnum("work.State", State_name, State_value)
	proto.RegisterType((*Task)(nil), "work.Task")
	proto.RegisterType((*TaskInfo)(nil), "work.TaskInfo")
}

func init() { proto.RegisterFile("server/pkg/work/work.proto", fileDescriptor_58a68e4647f78187) }

var fileDescriptor_58a68e4647f78187 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x37, 0x5d, 0x37, 0x67, 0x0a, 0x32, 0xc2, 0x18, 0x75, 0x87, 0x3a, 0x77, 0x90, 0xe2, 0x21,
	0x81, 0xfa, 0x02, 0x6e, 0x73, 0x4a, 0x41, 0x76, 0x48, 0xdd, 0xc5, 0x5b, 0xba, 0x66, 0xdd, 0xa8,
	0x36, 0x25, 0xc9, 0x94, 0xbd, 0xa1, 0x47, 0x9f, 0x40, 0xa4, 0x4f, 0x22, 0x49, 0x55, 0xd4, 0x4b,
	0xf8, 0xfd, 0xe3, 0xf7, 0x7d, 0xe4, 0x83, 0x43, 0xc5, 0xe5, 0x33, 0x97, 0xa4, 0x2a, 0x72, 0xf2,
	0x22, 0x64, 0x61, 0x1f, 0x5c, 0x49, 0xa1, 0x05, 0x72, 0x0d, 0x1e, 0xf6, 0x73, 0x91, 0x0b, 0x2b,
	0x10, 0x83, 0x1a, 0x6f, 0x78, 0x92, 0x0b, 0x91, 0x3f, 0x72, 0x62, 0x59, 0xba, 0x5b, 0x13, 0x56,
	0xee, 0x1b, 0x6b, 0x5c, 0x41, 0xf7, 0x9e, 0xa9, 0x02, 0x0d, 0xa0, 0xb3, 0xcd, 0x7c, 0x30, 0x02,
	0xe1, 0xd1, 0xb4, 0x53, 0xbf, 0x9f, 0x3a, 0xf1, 0x35, 0x75, 0xb6, 0x19, 0x0a, 0xa1, 0x9b, 0x31,
	0xcd, 0x7c, 0x67, 0x04, 0x42, 0x2f, 0xea, 0xe3, 0xa6, 0x09, 0x7f, 0x37, 0xe1, 0x49, 0xb9, 0xa7,
	0x36, 0x81, 0xce, 0x61, 0x57, 0xed, 0x52, 0xcd, 0x54, 0xa1, 0xfc, 0xd6, 0xa8, 0x15, 0x7a, 0x11,
	0xc4, 0x76, 0x3f, 0xd3, 0x4f, 0x7f, 0xbc, 0x31, 0x87, 0x5d, 0xa3, 0xc4, 0xe5, 0x5a, 0xa0, 0x00,
	0xba, 0x46, 0xb4, 0x73, 0xff, 0xe6, 0xad, 0x8e, 0xce, 0x60, 0x5b, 0x69, 0xa6, 0xb9, 0x1d, 0x7f,
	0x1c, 0x79, 0x4d, 0x20, 0x31, 0x12, 0x6d, 0x1c, 0x34, 0x80, 0x1d, 0xc9, 0x99, 0x12, 0xa5, 0xdf,
	0x32, 0xcb, 0xd3, 0x2f, 0x76, 0x81, 0x61, 0xdb, 0xe6, 0x90, 0x07, 0x0f, 0xe9, 0x72, 0xb1, 0x88,
	0x17, 0xb7, 0xbd, 0x03, 0x43, 0x92, 0xe5, 0x6c, 0x36, 0x4f, 0x92, 0x1e, 0x30, 0xe4, 0x66, 0x12,
	0xdf, 0x2d, 0xe9, 0xbc, 0xe7, 0x4c, 0xaf, 0x5e, 0xeb, 0x00, 0xbc, 0xd5, 0x01, 0xf8, 0xa8, 0x03,
	0xf0, 0x10, 0xe5, 0x5b, 0xbd, 0xd9, 0xa5, 0x78, 0x25, 0x9e, 0x48, 0xc5, 0x56, 0x9b, 0x7d, 0xc6,
	0xe5, 0x6f, 0xa4, 0xe4, 0x8a, 0xfc, 0x3b, 0x46, 0xda, 0xb1, 0x9f, 0x72, 0xf9, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x70, 0xd0, 0x37, 0x42, 0xa6, 0x01, 0x00, 0x00,
}

func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Task) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Subtasks) > 0 {
		for iNdEx := len(m.Subtasks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subtasks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWork(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWork(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintWork(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintWork(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if m.State != 0 {
		i = encodeVarintWork(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if m.Task != nil {
		{
			size, err := m.Task.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWork(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWork(dAtA []byte, offset int, v uint64) int {
	offset -= sovWork(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Task) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovWork(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovWork(uint64(l))
	}
	if len(m.Subtasks) > 0 {
		for _, e := range m.Subtasks {
			l = e.Size()
			n += 1 + l + sovWork(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Task != nil {
		l = m.Task.Size()
		n += 1 + l + sovWork(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovWork(uint64(m.State))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovWork(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWork(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWork(x uint64) (n int) {
	return sovWork(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subtasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subtasks = append(m.Subtasks, &Task{})
			if err := m.Subtasks[len(m.Subtasks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWork
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaskInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Task", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Task == nil {
				m.Task = &Task{}
			}
			if err := m.Task.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWork
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWork(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWork
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWork
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWork
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWork
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWork
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWork
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWork        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWork          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWork = fmt.Errorf("proto: unexpected end of group")
)
