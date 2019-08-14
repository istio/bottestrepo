// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/model/v1beta1/check.proto

package v1beta1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

	rpc "istio.io/gogo-genproto/googleapis/google/rpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Expresses the result of a precondition check.
type CheckResult struct {
	// A status code of OK indicates preconditions were satisfied. Any other code indicates preconditions were not
	// satisfied and details describe why.
	Status rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status"`
	// The amount of time for which this result can be considered valid.
	ValidDuration time.Duration `protobuf:"bytes,2,opt,name=valid_duration,json=validDuration,proto3,stdduration" json:"valid_duration"`
	// The number of uses for which this result can be considered valid.
	ValidUseCount int32 `protobuf:"varint,3,opt,name=valid_use_count,json=validUseCount,proto3" json:"valid_use_count,omitempty"`
}

func (m *CheckResult) Reset()      { *m = CheckResult{} }
func (*CheckResult) ProtoMessage() {}
func (*CheckResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cd393dc2a446fd, []int{0}
}
func (m *CheckResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResult.Merge(m, src)
}
func (m *CheckResult) XXX_Size() int {
	return m.Size()
}
func (m *CheckResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResult.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CheckResult)(nil), "istio.mixer.adapter.model.v1beta1.CheckResult")
}

func init() {
	proto.RegisterFile("mixer/adapter/model/v1beta1/check.proto", fileDescriptor_b5cd393dc2a446fd)
}

var fileDescriptor_b5cd393dc2a446fd = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x86, 0x6d, 0x7e, 0x2a, 0x94, 0x0a, 0x90, 0x22, 0x24, 0x4a, 0x87, 0xaf, 0x85, 0x01, 0x3a,
	0xd9, 0x14, 0x6e, 0xd0, 0x32, 0x31, 0x16, 0xb1, 0xb0, 0x54, 0x4e, 0x62, 0x82, 0x45, 0x5a, 0x47,
	0x89, 0x53, 0x31, 0x72, 0x04, 0x46, 0x8e, 0xc0, 0xce, 0x25, 0x3a, 0x76, 0xec, 0x04, 0xc4, 0x59,
	0x18, 0x7b, 0x04, 0x14, 0xdb, 0x59, 0xd9, 0x6c, 0x3f, 0xcf, 0xfb, 0xf9, 0x95, 0xed, 0x5d, 0xcc,
	0xc4, 0x0b, 0xcf, 0x28, 0x8b, 0x58, 0xaa, 0x78, 0x46, 0x67, 0x32, 0xe2, 0x09, 0x5d, 0x0c, 0x03,
	0xae, 0xd8, 0x90, 0x86, 0x4f, 0x3c, 0x7c, 0x26, 0x69, 0x26, 0x95, 0xf4, 0x4f, 0x45, 0xae, 0x84,
	0x24, 0x46, 0x27, 0x4e, 0x27, 0x46, 0x27, 0x4e, 0xef, 0x1e, 0xc5, 0x32, 0x96, 0xc6, 0xa6, 0xf5,
	0xca, 0x06, 0xbb, 0x10, 0x4b, 0x19, 0x27, 0x9c, 0x9a, 0x5d, 0x50, 0x3c, 0xd2, 0xa8, 0xc8, 0x98,
	0x12, 0x72, 0xee, 0xf8, 0xb1, 0xe3, 0x59, 0x1a, 0xd2, 0x5c, 0x31, 0x55, 0xe4, 0x16, 0x9c, 0x7d,
	0x62, 0xaf, 0x3d, 0xae, 0x1b, 0x4c, 0x78, 0x5e, 0x24, 0xca, 0xbf, 0xf4, 0x5a, 0x96, 0x77, 0x70,
	0x1f, 0x0f, 0xda, 0x57, 0x3e, 0xb1, 0x49, 0x92, 0xa5, 0x21, 0xb9, 0x33, 0x64, 0xb4, 0xb3, 0xfc,
	0xea, 0xa1, 0x89, 0xf3, 0xfc, 0x5b, 0xef, 0x60, 0xc1, 0x12, 0x11, 0x4d, 0x9b, 0x2b, 0x3b, 0x5b,
	0x26, 0x79, 0xd2, 0x24, 0x9b, 0x4e, 0xe4, 0xc6, 0x09, 0xa3, 0xbd, 0x7a, 0xc0, 0xfb, 0x77, 0x0f,
	0x4f, 0xf6, 0x4d, 0xb4, 0x01, 0xfe, 0xb9, 0x77, 0x68, 0x67, 0x15, 0x39, 0x9f, 0x86, 0xb2, 0x98,
	0xab, 0xce, 0x76, 0x1f, 0x0f, 0x76, 0x9d, 0x77, 0x9f, 0xf3, 0x71, 0x7d, 0x38, 0x0a, 0x96, 0x25,
	0xa0, 0x55, 0x09, 0x68, 0x5d, 0x02, 0xda, 0x94, 0x80, 0x5e, 0x35, 0xe0, 0x0f, 0x0d, 0x68, 0xa9,
	0x01, 0xaf, 0x34, 0xe0, 0x1f, 0x0d, 0xf8, 0x57, 0x03, 0xda, 0x68, 0xc0, 0x6f, 0x15, 0xa0, 0x55,
	0x05, 0x68, 0x5d, 0x01, 0x7a, 0x18, 0xd8, 0x17, 0x16, 0x92, 0xb2, 0x54, 0xd0, 0x7f, 0xfe, 0x25,
	0x68, 0x99, 0xde, 0xd7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x5b, 0xc8, 0xb5, 0xbd, 0x01,
	0x00, 0x00,
}

func (m *CheckResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValidUseCount != 0 {
		i = encodeVarintCheck(dAtA, i, uint64(m.ValidUseCount))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ValidDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCheck(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCheck(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCheck(dAtA []byte, offset int, v uint64) int {
	offset -= sovCheck(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CheckResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Status.Size()
	n += 1 + l + sovCheck(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovCheck(uint64(l))
	if m.ValidUseCount != 0 {
		n += 1 + sovCheck(uint64(m.ValidUseCount))
	}
	return n
}

func sovCheck(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCheck(x uint64) (n int) {
	return sovCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CheckResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CheckResult{`,
		`Status:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "rpc.Status", 1), `&`, ``, 1) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ValidDuration), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`ValidUseCount:` + fmt.Sprintf("%v", this.ValidUseCount) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCheck(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CheckResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheck
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
			return fmt.Errorf("proto: CheckResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ValidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUseCount", wireType)
			}
			m.ValidUseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUseCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheck
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCheck
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
func skipCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCheck
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
					return 0, ErrIntOverflowCheck
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCheck
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
				return 0, ErrInvalidLengthCheck
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCheck
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCheck
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCheck(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCheck
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCheck = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCheck   = fmt.Errorf("proto: integer overflow")
)