// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: OmniFlix/itc/v1/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_gogo_protobuf_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	MaxCampaignDuration time.Duration `protobuf:"bytes,1,opt,name=max_campaign_duration,json=maxCampaignDuration,proto3,stdduration" json:"max_campaign_duration" yaml:"max_campaign_duration"`
	CreationFee         types.Coin    `protobuf:"bytes,2,opt,name=creation_fee,json=creationFee,proto3" json:"creation_fee" yaml:"creation_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca4accf521d18f5, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "OmniFlix.itc.v1.Params")
}

func init() { proto.RegisterFile("OmniFlix/itc/v1/params.proto", fileDescriptor_dca4accf521d18f5) }

var fileDescriptor_dca4accf521d18f5 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xbd, 0x4e, 0xc3, 0x30,
	0x18, 0x8c, 0x19, 0x3a, 0xa4, 0x48, 0x48, 0x2d, 0x48, 0xa5, 0x54, 0x2e, 0xca, 0xd4, 0xc9, 0x56,
	0x40, 0x2c, 0x8c, 0x2d, 0xea, 0x84, 0x04, 0xea, 0x06, 0x4b, 0xe5, 0x18, 0xd7, 0xb5, 0x14, 0xe7,
	0x8b, 0x12, 0x27, 0xa4, 0x6f, 0xc1, 0xc8, 0x23, 0x75, 0xec, 0xc8, 0x80, 0x0a, 0x34, 0x6f, 0xc0,
	0x13, 0xa0, 0x26, 0x0e, 0x62, 0x60, 0xbb, 0xef, 0xe7, 0xee, 0x74, 0x3a, 0x77, 0x70, 0xa7, 0x23,
	0x35, 0x0d, 0x55, 0x41, 0x95, 0xe1, 0x34, 0xf7, 0x69, 0xcc, 0x12, 0xa6, 0x53, 0x12, 0x27, 0x60,
	0xa0, 0x73, 0xd4, 0x5c, 0x89, 0x32, 0x9c, 0xe4, 0x7e, 0x1f, 0x73, 0x48, 0x35, 0xa4, 0x34, 0x60,
	0xa9, 0xa0, 0xb9, 0x1f, 0x08, 0xc3, 0x7c, 0xca, 0x41, 0x45, 0x35, 0xa1, 0x7f, 0x2c, 0x41, 0x42,
	0x05, 0xe9, 0x1e, 0xd9, 0x2d, 0x96, 0x00, 0x32, 0x14, 0xb4, 0x9a, 0x82, 0x6c, 0x41, 0x9f, 0xb2,
	0x84, 0x19, 0x05, 0x96, 0xe5, 0xbd, 0x23, 0xb7, 0x75, 0x5f, 0xf9, 0x76, 0x9e, 0xdd, 0x13, 0xcd,
	0x8a, 0x39, 0x67, 0x3a, 0x66, 0x4a, 0x46, 0xf3, 0xe6, 0xb3, 0x87, 0xce, 0xd1, 0xa8, 0x7d, 0x71,
	0x4a, 0x6a, 0x29, 0xd2, 0x48, 0x91, 0x1b, 0xfb, 0x30, 0x1e, 0xad, 0xb7, 0x43, 0xe7, 0x7b, 0x3b,
	0x1c, 0xac, 0x98, 0x0e, 0xaf, 0xbd, 0x7f, 0x55, 0xbc, 0xd7, 0x8f, 0x21, 0x9a, 0x75, 0x35, 0x2b,
	0x26, 0xf6, 0xd4, 0xd0, 0x3b, 0x0f, 0xee, 0x21, 0x4f, 0x44, 0x85, 0xe7, 0x0b, 0x21, 0x7a, 0x07,
	0xd6, 0xaf, 0x0e, 0x4c, 0xf6, 0x81, 0x89, 0x0d, 0x4c, 0x26, 0xa0, 0xa2, 0xf1, 0x99, 0xf5, 0xeb,
	0xd6, 0x7e, 0x7f, 0xc9, 0xde, 0xac, 0xdd, 0x8c, 0x53, 0x21, 0xc6, 0xb7, 0xeb, 0x2f, 0xec, 0xac,
	0x77, 0x18, 0x6d, 0x76, 0x18, 0x7d, 0xee, 0x30, 0x7a, 0x29, 0xb1, 0xb3, 0x29, 0xb1, 0xf3, 0x56,
	0x62, 0xe7, 0x91, 0x48, 0x65, 0x96, 0x59, 0x40, 0x38, 0x68, 0xfa, 0x5b, 0x06, 0xe8, 0x48, 0x2d,
	0x42, 0x55, 0x2c, 0xb3, 0x80, 0xe6, 0x57, 0xb4, 0x6e, 0xc7, 0xac, 0x62, 0x91, 0x06, 0xad, 0x2a,
	0xfa, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x3f, 0x8e, 0xfd, 0xba, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CreationFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxCampaignDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxCampaignDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxCampaignDuration)
	n += 1 + l + sovParams(uint64(l))
	l = m.CreationFee.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCampaignDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxCampaignDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
