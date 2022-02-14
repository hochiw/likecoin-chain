// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likenft/class_data.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ClassData struct {
	Metadata     JsonInput `protobuf:"bytes,1,opt,name=metadata,proto3,customtype=JsonInput" json:"metadata"`
	IscnIdPrefix string    `protobuf:"bytes,2,opt,name=iscnIdPrefix,proto3" json:"iscnIdPrefix,omitempty"`
}

func (m *ClassData) Reset()         { *m = ClassData{} }
func (m *ClassData) String() string { return proto.CompactTextString(m) }
func (*ClassData) ProtoMessage()    {}
func (*ClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_084dbaa322495271, []int{0}
}
func (m *ClassData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassData.Merge(m, src)
}
func (m *ClassData) XXX_Size() int {
	return m.Size()
}
func (m *ClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassData.DiscardUnknown(m)
}

var xxx_messageInfo_ClassData proto.InternalMessageInfo

func (m *ClassData) GetIscnIdPrefix() string {
	if m != nil {
		return m.IscnIdPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*ClassData)(nil), "likecoin.likechain.likenft.ClassData")
}

func init() { proto.RegisterFile("likenft/class_data.proto", fileDescriptor_084dbaa322495271) }

var fileDescriptor_084dbaa322495271 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xc9, 0xcc, 0x4e,
	0xcd, 0x4b, 0x2b, 0xd1, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0x8e, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x02, 0xc9, 0x24, 0xe7, 0x67, 0xe6, 0xe9, 0x81, 0x19, 0x19,
	0x89, 0x50, 0x56, 0x5e, 0x5a, 0x89, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x99, 0x3e, 0x88,
	0x05, 0xd1, 0xa1, 0x14, 0xc7, 0xc5, 0xe9, 0x0c, 0x32, 0xc5, 0x25, 0xb1, 0x24, 0x51, 0x48, 0x97,
	0x8b, 0x23, 0x37, 0xb5, 0x24, 0x11, 0x64, 0xa0, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93, 0xe0,
	0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x73, 0x7a, 0x15, 0xe7, 0xe7, 0x79, 0xe6, 0x15, 0x94,
	0x96, 0x04, 0xc1, 0x95, 0x08, 0x29, 0x71, 0xf1, 0x64, 0x16, 0x27, 0xe7, 0x79, 0xa6, 0x04, 0x14,
	0xa5, 0xa6, 0x65, 0x56, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xa1, 0x88, 0x39, 0xb9, 0x9f,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xcc, 0xd9, 0xfa, 0x70, 0x67, 0xeb, 0x57, 0xe8, 0xc3, 0x7c,
	0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x76, 0xaf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0x39, 0x4e, 0xc7, 0xfd, 0x00, 0x00, 0x00,
}

func (m *ClassData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IscnIdPrefix) > 0 {
		i -= len(m.IscnIdPrefix)
		copy(dAtA[i:], m.IscnIdPrefix)
		i = encodeVarintClassData(dAtA, i, uint64(len(m.IscnIdPrefix)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Metadata.Size()
		i -= size
		if _, err := m.Metadata.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClassData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClassData(dAtA []byte, offset int, v uint64) int {
	offset -= sovClassData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Metadata.Size()
	n += 1 + l + sovClassData(uint64(l))
	l = len(m.IscnIdPrefix)
	if l > 0 {
		n += 1 + l + sovClassData(uint64(l))
	}
	return n
}

func sovClassData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClassData(x uint64) (n int) {
	return sovClassData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassData
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
			return fmt.Errorf("proto: ClassData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
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
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IscnIdPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
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
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IscnIdPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassData
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
func skipClassData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClassData
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
					return 0, ErrIntOverflowClassData
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
					return 0, ErrIntOverflowClassData
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
				return 0, ErrInvalidLengthClassData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClassData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClassData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClassData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClassData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClassData = fmt.Errorf("proto: unexpected end of group")
)
