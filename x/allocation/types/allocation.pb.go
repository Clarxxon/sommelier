// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: allocation/v1/allocation.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_tendermint_tendermint_libs_bytes "github.com/tendermint/tendermint/libs/bytes"
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

// AllocationPrecommit defines an array of hashed data to be used for the precommit phase
// of allocation
type AllocationPrecommit struct {
	Hash     github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=hash,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"hash,omitempty"`
	CellarId string                                               `protobuf:"bytes,2,opt,name=cellar_id,json=cellarId,proto3" json:"cellar_id,omitempty"`
}

func (m *AllocationPrecommit) Reset()         { *m = AllocationPrecommit{} }
func (m *AllocationPrecommit) String() string { return proto.CompactTextString(m) }
func (*AllocationPrecommit) ProtoMessage()    {}
func (*AllocationPrecommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_23d2c35dae4a6cad, []int{0}
}
func (m *AllocationPrecommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllocationPrecommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllocationPrecommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllocationPrecommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationPrecommit.Merge(m, src)
}
func (m *AllocationPrecommit) XXX_Size() int {
	return m.Size()
}
func (m *AllocationPrecommit) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationPrecommit.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationPrecommit proto.InternalMessageInfo

func (m *AllocationPrecommit) GetHash() github_com_tendermint_tendermint_libs_bytes.HexBytes {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *AllocationPrecommit) GetCellarId() string {
	if m != nil {
		return m.CellarId
	}
	return ""
}

// Allocation is the commit for all allocations for a cellar by a validator
type Allocation struct {
	Cellar *Cellar `protobuf:"bytes,1,opt,name=cellar,proto3" json:"cellar,omitempty"`
	Salt   string  `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
}

func (m *Allocation) Reset()         { *m = Allocation{} }
func (m *Allocation) String() string { return proto.CompactTextString(m) }
func (*Allocation) ProtoMessage()    {}
func (*Allocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_23d2c35dae4a6cad, []int{1}
}
func (m *Allocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Allocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Allocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Allocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Allocation.Merge(m, src)
}
func (m *Allocation) XXX_Size() int {
	return m.Size()
}
func (m *Allocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Allocation.DiscardUnknown(m)
}

var xxx_messageInfo_Allocation proto.InternalMessageInfo

func (m *Allocation) GetCellar() *Cellar {
	if m != nil {
		return m.Cellar
	}
	return nil
}

func (m *Allocation) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

// Cellar is a collection of pools for a token pair
type Cellar struct {
	Id         string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TickRanges []*TickRange `protobuf:"bytes,2,rep,name=tick_ranges,json=tickRanges,proto3" json:"tick_ranges,omitempty"`
}

func (m *Cellar) Reset()         { *m = Cellar{} }
func (m *Cellar) String() string { return proto.CompactTextString(m) }
func (*Cellar) ProtoMessage()    {}
func (*Cellar) Descriptor() ([]byte, []int) {
	return fileDescriptor_23d2c35dae4a6cad, []int{2}
}
func (m *Cellar) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cellar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cellar.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cellar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cellar.Merge(m, src)
}
func (m *Cellar) XXX_Size() int {
	return m.Size()
}
func (m *Cellar) XXX_DiscardUnknown() {
	xxx_messageInfo_Cellar.DiscardUnknown(m)
}

var xxx_messageInfo_Cellar proto.InternalMessageInfo

func (m *Cellar) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cellar) GetTickRanges() []*TickRange {
	if m != nil {
		return m.TickRanges
	}
	return nil
}

type TickRange struct {
	Upper  uint64 `protobuf:"varint,1,opt,name=upper,proto3" json:"upper,omitempty"`
	Lower  uint64 `protobuf:"varint,2,opt,name=lower,proto3" json:"lower,omitempty"`
	Weight uint64 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (m *TickRange) Reset()         { *m = TickRange{} }
func (m *TickRange) String() string { return proto.CompactTextString(m) }
func (*TickRange) ProtoMessage()    {}
func (*TickRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_23d2c35dae4a6cad, []int{3}
}
func (m *TickRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickRange.Merge(m, src)
}
func (m *TickRange) XXX_Size() int {
	return m.Size()
}
func (m *TickRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TickRange.DiscardUnknown(m)
}

var xxx_messageInfo_TickRange proto.InternalMessageInfo

func (m *TickRange) GetUpper() uint64 {
	if m != nil {
		return m.Upper
	}
	return 0
}

func (m *TickRange) GetLower() uint64 {
	if m != nil {
		return m.Lower
	}
	return 0
}

func (m *TickRange) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type ManagedCellarsUpdateProposal struct {
	Title       string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Cellars     []*Cellar `protobuf:"bytes,3,rep,name=cellars,proto3" json:"cellars,omitempty"`
}

func (m *ManagedCellarsUpdateProposal) Reset()         { *m = ManagedCellarsUpdateProposal{} }
func (m *ManagedCellarsUpdateProposal) String() string { return proto.CompactTextString(m) }
func (*ManagedCellarsUpdateProposal) ProtoMessage()    {}
func (*ManagedCellarsUpdateProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_23d2c35dae4a6cad, []int{4}
}
func (m *ManagedCellarsUpdateProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManagedCellarsUpdateProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManagedCellarsUpdateProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManagedCellarsUpdateProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedCellarsUpdateProposal.Merge(m, src)
}
func (m *ManagedCellarsUpdateProposal) XXX_Size() int {
	return m.Size()
}
func (m *ManagedCellarsUpdateProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedCellarsUpdateProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedCellarsUpdateProposal proto.InternalMessageInfo

func (m *ManagedCellarsUpdateProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ManagedCellarsUpdateProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ManagedCellarsUpdateProposal) GetCellars() []*Cellar {
	if m != nil {
		return m.Cellars
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationPrecommit)(nil), "allocation.v1.AllocationPrecommit")
	proto.RegisterType((*Allocation)(nil), "allocation.v1.Allocation")
	proto.RegisterType((*Cellar)(nil), "allocation.v1.Cellar")
	proto.RegisterType((*TickRange)(nil), "allocation.v1.TickRange")
	proto.RegisterType((*ManagedCellarsUpdateProposal)(nil), "allocation.v1.ManagedCellarsUpdateProposal")
}

func init() { proto.RegisterFile("allocation/v1/allocation.proto", fileDescriptor_23d2c35dae4a6cad) }

var fileDescriptor_23d2c35dae4a6cad = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x31, 0x6f, 0x13, 0x31,
	0x18, 0xcd, 0x25, 0x21, 0x90, 0x2f, 0xc0, 0x60, 0x0a, 0x3a, 0x01, 0x3a, 0xa2, 0x9b, 0xb2, 0x70,
	0x56, 0x0b, 0x03, 0x8c, 0x84, 0x85, 0x4a, 0xa0, 0x56, 0x07, 0x2c, 0x2c, 0x95, 0x73, 0xf7, 0xc9,
	0x31, 0xf5, 0x9d, 0x2d, 0xdb, 0x4d, 0x9b, 0x8d, 0x8d, 0x95, 0x9f, 0xc5, 0xd8, 0x91, 0x09, 0xa1,
	0xe4, 0x5f, 0x30, 0xa1, 0xb3, 0xaf, 0xcd, 0x81, 0xc4, 0xf6, 0xde, 0xfb, 0x9e, 0x9e, 0xdf, 0x67,
	0x1b, 0x12, 0x26, 0xa5, 0x2a, 0x98, 0x13, 0xaa, 0xa6, 0xab, 0x7d, 0xba, 0x63, 0x99, 0x36, 0xca,
	0x29, 0x72, 0xa7, 0xa3, 0xac, 0xf6, 0x1f, 0xee, 0x71, 0xc5, 0x95, 0x9f, 0xd0, 0x06, 0x05, 0x53,
	0xfa, 0x25, 0x82, 0x7b, 0xaf, 0xae, 0x7d, 0xc7, 0x06, 0x0b, 0x55, 0x55, 0xc2, 0x91, 0xb7, 0x30,
	0x5c, 0x32, 0xbb, 0x8c, 0xa3, 0x69, 0x34, 0xbb, 0x3d, 0x7f, 0xf1, 0xfb, 0xe7, 0x93, 0xe7, 0x5c,
	0xb8, 0xe5, 0xd9, 0x22, 0x2b, 0x54, 0x45, 0x1d, 0xd6, 0x25, 0x9a, 0x4a, 0xd4, 0xae, 0x0b, 0xa5,
	0x58, 0x58, 0xba, 0x58, 0x3b, 0xb4, 0xd9, 0x1b, 0xbc, 0x98, 0x37, 0x20, 0xf7, 0x29, 0xe4, 0x11,
	0x8c, 0x0b, 0x94, 0x92, 0x99, 0x13, 0x51, 0xc6, 0xfd, 0x69, 0x34, 0x1b, 0xe7, 0xb7, 0x82, 0x70,
	0x58, 0xa6, 0x47, 0x00, 0xbb, 0x06, 0xe4, 0x29, 0x8c, 0xc2, 0xc4, 0x1f, 0x3d, 0x39, 0xb8, 0x9f,
	0xfd, 0xb5, 0x46, 0xf6, 0xda, 0x0f, 0xf3, 0xd6, 0x44, 0x08, 0x0c, 0x2d, 0x93, 0xae, 0x0d, 0xf5,
	0x38, 0x7d, 0x0f, 0xa3, 0xe0, 0x22, 0x77, 0xa1, 0x2f, 0x4a, 0x1f, 0x34, 0xce, 0xfb, 0xa2, 0x24,
	0x2f, 0x61, 0xe2, 0x44, 0x71, 0x7a, 0x62, 0x58, 0xcd, 0xd1, 0xc6, 0xfd, 0xe9, 0x60, 0x36, 0x39,
	0x88, 0xff, 0x39, 0xe1, 0x83, 0x28, 0x4e, 0xf3, 0xc6, 0x90, 0x83, 0xbb, 0x82, 0x36, 0x3d, 0x82,
	0xf1, 0xf5, 0x80, 0xec, 0xc1, 0x8d, 0x33, 0xad, 0x31, 0x74, 0x1c, 0xe6, 0x81, 0x34, 0xaa, 0x54,
	0xe7, 0x68, 0x7c, 0x99, 0x61, 0x1e, 0x08, 0x79, 0x00, 0xa3, 0x73, 0x14, 0x7c, 0xe9, 0xe2, 0x81,
	0x97, 0x5b, 0x96, 0x7e, 0x8d, 0xe0, 0xf1, 0x3b, 0x56, 0x33, 0x8e, 0x65, 0x68, 0x6b, 0x3f, 0xea,
	0x92, 0x39, 0x3c, 0x36, 0x4a, 0x2b, 0xcb, 0x64, 0x13, 0xe7, 0x84, 0x93, 0xd8, 0xf6, 0x0f, 0x84,
	0x4c, 0x61, 0x52, 0xa2, 0x2d, 0x8c, 0xd0, 0x4d, 0xdf, 0x76, 0xef, 0xae, 0x44, 0x28, 0xdc, 0x0c,
	0x97, 0x63, 0xe3, 0x81, 0x5f, 0xf0, 0x3f, 0x57, 0x78, 0xe5, 0x9a, 0x1f, 0x7e, 0xdf, 0x24, 0xd1,
	0xe5, 0x26, 0x89, 0x7e, 0x6d, 0x92, 0xe8, 0xdb, 0x36, 0xe9, 0x5d, 0x6e, 0x93, 0xde, 0x8f, 0x6d,
	0xd2, 0xfb, 0x44, 0x3b, 0x6f, 0xae, 0x91, 0xf3, 0xf5, 0xe7, 0x15, 0xb5, 0xaa, 0xaa, 0x50, 0x0a,
	0x34, 0xf4, 0xa2, 0xf3, 0xe7, 0xa8, 0x5b, 0x6b, 0xb4, 0x8b, 0x91, 0xff, 0x55, 0xcf, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0xb3, 0x5d, 0x99, 0x9c, 0x02, 0x00, 0x00,
}

func (m *AllocationPrecommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllocationPrecommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllocationPrecommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CellarId) > 0 {
		i -= len(m.CellarId)
		copy(dAtA[i:], m.CellarId)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.CellarId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Allocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Allocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Allocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0x12
	}
	if m.Cellar != nil {
		{
			size, err := m.Cellar.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAllocation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Cellar) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cellar) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cellar) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TickRanges) > 0 {
		for iNdEx := len(m.TickRanges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickRanges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAllocation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TickRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Weight != 0 {
		i = encodeVarintAllocation(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x18
	}
	if m.Lower != 0 {
		i = encodeVarintAllocation(dAtA, i, uint64(m.Lower))
		i--
		dAtA[i] = 0x10
	}
	if m.Upper != 0 {
		i = encodeVarintAllocation(dAtA, i, uint64(m.Upper))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ManagedCellarsUpdateProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManagedCellarsUpdateProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManagedCellarsUpdateProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cellars) > 0 {
		for iNdEx := len(m.Cellars) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Cellars[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAllocation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAllocation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAllocation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllocationPrecommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	l = len(m.CellarId)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	return n
}

func (m *Allocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cellar != nil {
		l = m.Cellar.Size()
		n += 1 + l + sovAllocation(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	return n
}

func (m *Cellar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	if len(m.TickRanges) > 0 {
		for _, e := range m.TickRanges {
			l = e.Size()
			n += 1 + l + sovAllocation(uint64(l))
		}
	}
	return n
}

func (m *TickRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Upper != 0 {
		n += 1 + sovAllocation(uint64(m.Upper))
	}
	if m.Lower != 0 {
		n += 1 + sovAllocation(uint64(m.Lower))
	}
	if m.Weight != 0 {
		n += 1 + sovAllocation(uint64(m.Weight))
	}
	return n
}

func (m *ManagedCellarsUpdateProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	if len(m.Cellars) > 0 {
		for _, e := range m.Cellars {
			l = e.Size()
			n += 1 + l + sovAllocation(uint64(l))
		}
	}
	return n
}

func sovAllocation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAllocation(x uint64) (n int) {
	return sovAllocation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllocationPrecommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocation
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
			return fmt.Errorf("proto: AllocationPrecommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllocationPrecommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellarId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CellarId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllocation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Allocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocation
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
			return fmt.Errorf("proto: Allocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Allocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cellar", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cellar == nil {
				m.Cellar = &Cellar{}
			}
			if err := m.Cellar.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllocation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Cellar) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocation
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
			return fmt.Errorf("proto: Cellar: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cellar: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickRanges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickRanges = append(m.TickRanges, &TickRange{})
			if err := m.TickRanges[len(m.TickRanges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllocation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TickRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocation
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
			return fmt.Errorf("proto: TickRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upper", wireType)
			}
			m.Upper = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Upper |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lower", wireType)
			}
			m.Lower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAllocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllocation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManagedCellarsUpdateProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocation
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
			return fmt.Errorf("proto: ManagedCellarsUpdateProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManagedCellarsUpdateProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cellars", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cellars = append(m.Cellars, &Cellar{})
			if err := m.Cellars[len(m.Cellars)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllocation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAllocation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllocation
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
					return 0, ErrIntOverflowAllocation
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
					return 0, ErrIntOverflowAllocation
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
				return 0, ErrInvalidLengthAllocation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAllocation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAllocation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAllocation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllocation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAllocation = fmt.Errorf("proto: unexpected end of group")
)
