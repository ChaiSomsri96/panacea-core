// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aol/writer.proto

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

type Writer struct {
	Moniker       string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Description   string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	NanoTimestamp int64  `protobuf:"varint,3,opt,name=nanoTimestamp,proto3" json:"nanoTimestamp,omitempty"`
}

func (m *Writer) Reset()         { *m = Writer{} }
func (m *Writer) String() string { return proto.CompactTextString(m) }
func (*Writer) ProtoMessage()    {}
func (*Writer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3eb7018e147f9d3, []int{0}
}
func (m *Writer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Writer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Writer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Writer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Writer.Merge(m, src)
}
func (m *Writer) XXX_Size() int {
	return m.Size()
}
func (m *Writer) XXX_DiscardUnknown() {
	xxx_messageInfo_Writer.DiscardUnknown(m)
}

var xxx_messageInfo_Writer proto.InternalMessageInfo

func (m *Writer) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Writer) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Writer) GetNanoTimestamp() int64 {
	if m != nil {
		return m.NanoTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Writer)(nil), "medibloc.panaceacore.aol.Writer")
}

func init() { proto.RegisterFile("aol/writer.proto", fileDescriptor_e3eb7018e147f9d3) }

var fileDescriptor_e3eb7018e147f9d3 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcc, 0xcf, 0xd1,
	0x2f, 0x2f, 0xca, 0x2c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xc8, 0x4d,
	0x4d, 0xc9, 0x4c, 0xca, 0xc9, 0x4f, 0xd6, 0x2b, 0x48, 0xcc, 0x4b, 0x4c, 0x4e, 0x4d, 0x4c, 0xce,
	0x2f, 0x4a, 0xd5, 0x4b, 0xcc, 0xcf, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd2, 0x07,
	0xb1, 0x20, 0xea, 0x95, 0xb2, 0xb8, 0xd8, 0xc2, 0xc1, 0xfa, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0xf3,
	0xf3, 0x32, 0xb3, 0x53, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x05,
	0x2e, 0xee, 0x94, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x26, 0xb0,
	0x2c, 0xb2, 0x90, 0x90, 0x0a, 0x17, 0x6f, 0x5e, 0x62, 0x5e, 0x7e, 0x48, 0x66, 0x6e, 0x6a, 0x71,
	0x49, 0x62, 0x6e, 0x81, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0xaa, 0xa0, 0x93, 0xdb, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0x3c, 0xa0, 0x0f, 0xf5, 0x80, 0x2e, 0xc8, 0x07, 0xfa, 0x15,
	0xfa, 0x20, 0xaf, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x9d, 0x6e, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0x3a, 0x41, 0x1e, 0xfe, 0x00, 0x00, 0x00,
}

func (m *Writer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Writer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Writer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NanoTimestamp != 0 {
		i = encodeVarintWriter(dAtA, i, uint64(m.NanoTimestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintWriter(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintWriter(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWriter(dAtA []byte, offset int, v uint64) int {
	offset -= sovWriter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Writer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovWriter(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovWriter(uint64(l))
	}
	if m.NanoTimestamp != 0 {
		n += 1 + sovWriter(uint64(m.NanoTimestamp))
	}
	return n
}

func sovWriter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWriter(x uint64) (n int) {
	return sovWriter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Writer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWriter
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
			return fmt.Errorf("proto: Writer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Writer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWriter
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
				return ErrInvalidLengthWriter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWriter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWriter
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
				return ErrInvalidLengthWriter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWriter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NanoTimestamp", wireType)
			}
			m.NanoTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWriter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NanoTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWriter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWriter
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
func skipWriter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWriter
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
					return 0, ErrIntOverflowWriter
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
					return 0, ErrIntOverflowWriter
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
				return 0, ErrInvalidLengthWriter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWriter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWriter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWriter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWriter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWriter = fmt.Errorf("proto: unexpected end of group")
)
