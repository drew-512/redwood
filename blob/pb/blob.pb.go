// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blob.proto

package pb

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	pb "redwood.dev/types/pb"
	reflect "reflect"
	strings "strings"
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

type BlobID struct {
	HashAlg pb.HashAlg `protobuf:"varint,1,opt,name=hashAlg,proto3,enum=Redwood.types.HashAlg" json:"hashAlg,omitempty"`
	Hash    []byte     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *BlobID) Reset()      { *m = BlobID{} }
func (*BlobID) ProtoMessage() {}
func (*BlobID) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{0}
}
func (m *BlobID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlobID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlobID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlobID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlobID.Merge(m, src)
}
func (m *BlobID) XXX_Size() int {
	return m.Size()
}
func (m *BlobID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlobID.DiscardUnknown(m)
}

var xxx_messageInfo_BlobID proto.InternalMessageInfo

func (m *BlobID) GetHashAlg() pb.HashAlg {
	if m != nil {
		return m.HashAlg
	}
	return pb.Unknown
}

func (m *BlobID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Manifest struct {
	Size_      uint64   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	ChunkSHA3S [][]byte `protobuf:"bytes,2,rep,name=chunkSHA3s,proto3" json:"chunkSHA3s,omitempty"`
}

func (m *Manifest) Reset()      { *m = Manifest{} }
func (*Manifest) ProtoMessage() {}
func (*Manifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6903d1e8a20272e8, []int{1}
}
func (m *Manifest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Manifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Manifest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Manifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest.Merge(m, src)
}
func (m *Manifest) XXX_Size() int {
	return m.Size()
}
func (m *Manifest) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest proto.InternalMessageInfo

func (m *Manifest) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Manifest) GetChunkSHA3S() [][]byte {
	if m != nil {
		return m.ChunkSHA3S
	}
	return nil
}

func init() {
	proto.RegisterType((*BlobID)(nil), "Redwood.blob.BlobID")
	proto.RegisterType((*Manifest)(nil), "Redwood.blob.Manifest")
}

func init() { proto.RegisterFile("blob.proto", fileDescriptor_6903d1e8a20272e8) }

var fileDescriptor_6903d1e8a20272e8 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x2d, 0x4f, 0xc3, 0x40,
	0x18, 0xc7, 0xef, 0x19, 0xcb, 0x20, 0x97, 0x06, 0x51, 0x08, 0x69, 0x26, 0x9e, 0x34, 0x53, 0x35,
	0xf4, 0x08, 0xf3, 0x24, 0x5b, 0x10, 0x43, 0x80, 0x28, 0x0a, 0xdc, 0x6e, 0xed, 0xda, 0x86, 0xb2,
	0x6b, 0x76, 0x2d, 0x04, 0x14, 0x1f, 0x81, 0x8f, 0xc1, 0x47, 0x40, 0x22, 0x91, 0x95, 0x93, 0xbb,
	0xab, 0x41, 0x4e, 0x22, 0x49, 0xaf, 0x2c, 0x99, 0xfb, 0x3d, 0xff, 0x97, 0x7b, 0x79, 0x28, 0xe5,
	0x99, 0xe0, 0x7e, 0xbe, 0x14, 0x85, 0xb0, 0xad, 0x20, 0x0a, 0x9f, 0x85, 0x08, 0xfd, 0x46, 0xeb,
	0x9f, 0xc6, 0x69, 0x91, 0x94, 0xdc, 0x9f, 0x89, 0x47, 0x16, 0x8b, 0x58, 0x30, 0x13, 0xe2, 0xe5,
	0xdc, 0x4c, 0x66, 0x30, 0xd4, 0x96, 0xfb, 0xc7, 0xc5, 0x4b, 0x1e, 0x49, 0x96, 0x73, 0x66, 0xa0,
	0x55, 0x07, 0x37, 0xb4, 0x37, 0xce, 0x04, 0xbf, 0xba, 0xb4, 0xcf, 0xe8, 0x7e, 0x32, 0x95, 0xc9,
	0x28, 0x8b, 0x1d, 0x70, 0xc1, 0x3b, 0x3c, 0x3f, 0xf1, 0xb7, 0xd7, 0xb5, 0x85, 0x49, 0xeb, 0x06,
	0xdb, 0x98, 0x6d, 0xd3, 0x6e, 0x83, 0x4e, 0xc7, 0x05, 0xcf, 0x0a, 0x0c, 0x0f, 0x2e, 0xe8, 0xc1,
	0xf5, 0x74, 0x91, 0xce, 0x23, 0x59, 0x34, 0xbe, 0x4c, 0x5f, 0x23, 0x73, 0x5c, 0x37, 0x30, 0x6c,
	0x23, 0xa5, 0xb3, 0xa4, 0x5c, 0x3c, 0xdc, 0x4e, 0x46, 0x43, 0xe9, 0x74, 0xdc, 0x3d, 0xcf, 0x0a,
	0x76, 0x94, 0xf1, 0x5d, 0xa5, 0x90, 0xac, 0x14, 0x92, 0xb5, 0x42, 0xd8, 0x28, 0x84, 0x5f, 0x85,
	0xf0, 0xa6, 0x11, 0x3e, 0x34, 0xc2, 0xa7, 0x46, 0xf8, 0xd2, 0x08, 0xdf, 0x1a, 0xa1, 0xd2, 0x08,
	0x6b, 0x8d, 0xf0, 0xa3, 0x91, 0x6c, 0x34, 0xc2, 0x7b, 0x8d, 0xa4, 0xaa, 0x91, 0xac, 0x6a, 0x24,
	0xf7, 0x47, 0xcb, 0xff, 0x97, 0x87, 0xd1, 0x13, 0x6b, 0x96, 0xc5, 0x72, 0xce, 0x7b, 0xe6, 0xc7,
	0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x16, 0xc9, 0x8f, 0x52, 0x01, 0x00, 0x00,
}

func (this *BlobID) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*BlobID)
	if !ok {
		that2, ok := that.(BlobID)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *BlobID")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *BlobID but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *BlobID but is not nil && this == nil")
	}
	if this.HashAlg != that1.HashAlg {
		return fmt.Errorf("HashAlg this(%v) Not Equal that(%v)", this.HashAlg, that1.HashAlg)
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return fmt.Errorf("Hash this(%v) Not Equal that(%v)", this.Hash, that1.Hash)
	}
	return nil
}
func (this *BlobID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BlobID)
	if !ok {
		that2, ok := that.(BlobID)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.HashAlg != that1.HashAlg {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	return true
}
func (this *Manifest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Manifest)
	if !ok {
		that2, ok := that.(Manifest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Manifest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Manifest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Manifest but is not nil && this == nil")
	}
	if this.Size_ != that1.Size_ {
		return fmt.Errorf("Size_ this(%v) Not Equal that(%v)", this.Size_, that1.Size_)
	}
	if len(this.ChunkSHA3S) != len(that1.ChunkSHA3S) {
		return fmt.Errorf("ChunkSHA3S this(%v) Not Equal that(%v)", len(this.ChunkSHA3S), len(that1.ChunkSHA3S))
	}
	for i := range this.ChunkSHA3S {
		if !bytes.Equal(this.ChunkSHA3S[i], that1.ChunkSHA3S[i]) {
			return fmt.Errorf("ChunkSHA3S this[%v](%v) Not Equal that[%v](%v)", i, this.ChunkSHA3S[i], i, that1.ChunkSHA3S[i])
		}
	}
	return nil
}
func (this *Manifest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Manifest)
	if !ok {
		that2, ok := that.(Manifest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Size_ != that1.Size_ {
		return false
	}
	if len(this.ChunkSHA3S) != len(that1.ChunkSHA3S) {
		return false
	}
	for i := range this.ChunkSHA3S {
		if !bytes.Equal(this.ChunkSHA3S[i], that1.ChunkSHA3S[i]) {
			return false
		}
	}
	return true
}
func (this *BlobID) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.BlobID{")
	s = append(s, "HashAlg: "+fmt.Sprintf("%#v", this.HashAlg)+",\n")
	s = append(s, "Hash: "+fmt.Sprintf("%#v", this.Hash)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Manifest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.Manifest{")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "ChunkSHA3S: "+fmt.Sprintf("%#v", this.ChunkSHA3S)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringBlob(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *BlobID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlobID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlobID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintBlob(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.HashAlg != 0 {
		i = encodeVarintBlob(dAtA, i, uint64(m.HashAlg))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Manifest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Manifest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Manifest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChunkSHA3S) > 0 {
		for iNdEx := len(m.ChunkSHA3S) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChunkSHA3S[iNdEx])
			copy(dAtA[i:], m.ChunkSHA3S[iNdEx])
			i = encodeVarintBlob(dAtA, i, uint64(len(m.ChunkSHA3S[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Size_ != 0 {
		i = encodeVarintBlob(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlob(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlob(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedBlobID(r randyBlob, easy bool) *BlobID {
	this := &BlobID{}
	this.HashAlg = pb.HashAlg([]int32{0, 1, 2}[r.Intn(3)])
	v1 := r.Intn(100)
	this.Hash = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Hash[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedManifest(r randyBlob, easy bool) *Manifest {
	this := &Manifest{}
	this.Size_ = uint64(uint64(r.Uint32()))
	v2 := r.Intn(10)
	this.ChunkSHA3S = make([][]byte, v2)
	for i := 0; i < v2; i++ {
		v3 := r.Intn(100)
		this.ChunkSHA3S[i] = make([]byte, v3)
		for j := 0; j < v3; j++ {
			this.ChunkSHA3S[i][j] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyBlob interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneBlob(r randyBlob) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringBlob(r randyBlob) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneBlob(r)
	}
	return string(tmps)
}
func randUnrecognizedBlob(r randyBlob, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldBlob(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldBlob(dAtA []byte, r randyBlob, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateBlob(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateBlob(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateBlob(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateBlob(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateBlob(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateBlob(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateBlob(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *BlobID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HashAlg != 0 {
		n += 1 + sovBlob(uint64(m.HashAlg))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlob(uint64(l))
	}
	return n
}

func (m *Manifest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Size_ != 0 {
		n += 1 + sovBlob(uint64(m.Size_))
	}
	if len(m.ChunkSHA3S) > 0 {
		for _, b := range m.ChunkSHA3S {
			l = len(b)
			n += 1 + l + sovBlob(uint64(l))
		}
	}
	return n
}

func sovBlob(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlob(x uint64) (n int) {
	return sovBlob(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BlobID) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BlobID{`,
		`HashAlg:` + fmt.Sprintf("%v", this.HashAlg) + `,`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Manifest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Manifest{`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`ChunkSHA3S:` + fmt.Sprintf("%v", this.ChunkSHA3S) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlob(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BlobID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlob
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
			return fmt.Errorf("proto: BlobID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlobID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashAlg", wireType)
			}
			m.HashAlg = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HashAlg |= pb.HashAlg(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
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
				return ErrInvalidLengthBlob
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlob
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
func (m *Manifest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlob
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
			return fmt.Errorf("proto: Manifest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Manifest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkSHA3S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlob
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
				return ErrInvalidLengthBlob
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChunkSHA3S = append(m.ChunkSHA3S, make([]byte, postIndex-iNdEx))
			copy(m.ChunkSHA3S[len(m.ChunkSHA3S)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlob
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
func skipBlob(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlob
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
					return 0, ErrIntOverflowBlob
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
					return 0, ErrIntOverflowBlob
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
				return 0, ErrInvalidLengthBlob
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlob
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlob
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlob        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlob          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlob = fmt.Errorf("proto: unexpected end of group")
)
