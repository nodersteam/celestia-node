// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fraud/pb/proof.proto

package fraud_pb

import (
	fmt "fmt"
	pb "github.com/celestiaorg/celestia-node/ipld/pb"
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

type BadEncoding struct {
	Height uint64      `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
	Shares []*pb.Share `protobuf:"bytes,2,rep,name=Shares,proto3" json:"Shares,omitempty"`
	Index  uint32      `protobuf:"varint,3,opt,name=Index,proto3" json:"Index,omitempty"`
	IsRow  bool        `protobuf:"varint,4,opt,name=isRow,proto3" json:"isRow,omitempty"`
}

func (m *BadEncoding) Reset()         { *m = BadEncoding{} }
func (m *BadEncoding) String() string { return proto.CompactTextString(m) }
func (*BadEncoding) ProtoMessage()    {}
func (*BadEncoding) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{0}
}
func (m *BadEncoding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadEncoding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadEncoding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadEncoding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadEncoding.Merge(m, src)
}
func (m *BadEncoding) XXX_Size() int {
	return m.Size()
}
func (m *BadEncoding) XXX_DiscardUnknown() {
	xxx_messageInfo_BadEncoding.DiscardUnknown(m)
}

var xxx_messageInfo_BadEncoding proto.InternalMessageInfo

func (m *BadEncoding) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BadEncoding) GetShares() []*pb.Share {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *BadEncoding) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BadEncoding) GetIsRow() bool {
	if m != nil {
		return m.IsRow
	}
	return false
}

func init() {
	proto.RegisterType((*BadEncoding)(nil), "fraud.pb.BadEncoding")
}

func init() { proto.RegisterFile("fraud/pb/proof.proto", fileDescriptor_318cb87a8bb2d394) }

var fileDescriptor_318cb87a8bb2d394 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2b, 0x4a, 0x2c,
	0x4d, 0xd1, 0x2f, 0x48, 0xd2, 0x2f, 0x28, 0xca, 0xcf, 0x4f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x00, 0x8b, 0xea, 0x15, 0x24, 0x49, 0x09, 0x67, 0x16, 0xe4, 0x80, 0xa5, 0x8b, 0x33,
	0x12, 0x8b, 0x52, 0x21, 0xd2, 0x4a, 0x95, 0x5c, 0xdc, 0x4e, 0x89, 0x29, 0xae, 0x79, 0xc9, 0xf9,
	0x29, 0x99, 0x79, 0xe9, 0x42, 0x62, 0x5c, 0x6c, 0x1e, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x90, 0x1a, 0x17, 0x5b, 0x30, 0x48, 0x57, 0xb1, 0x04,
	0x93, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0x9f, 0x1e, 0xc8, 0x30, 0xbd, 0x82, 0x24, 0x3d, 0xb0, 0x70,
	0x10, 0x54, 0x56, 0x48, 0x84, 0x8b, 0xd5, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x59, 0x81, 0x51,
	0x83, 0x37, 0x08, 0xc2, 0x01, 0x89, 0x66, 0x16, 0x07, 0xe5, 0x97, 0x4b, 0xb0, 0x28, 0x30, 0x6a,
	0x70, 0x04, 0x41, 0x38, 0x4e, 0x12, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90,
	0xc4, 0x06, 0x76, 0x9b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x49, 0xe7, 0xec, 0x82, 0xd2, 0x00,
	0x00, 0x00,
}

func (m *BadEncoding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadEncoding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadEncoding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsRow {
		i--
		if m.IsRow {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Index != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProof(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BadEncoding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovProof(uint64(m.Height))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	if m.Index != 0 {
		n += 1 + sovProof(uint64(m.Index))
	}
	if m.IsRow {
		n += 2
	}
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BadEncoding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: BadEncoding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadEncoding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, &pb.Share{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRow", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsRow = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)