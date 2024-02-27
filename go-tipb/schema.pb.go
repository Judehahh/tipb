// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TableInfo struct {
	TableId          int64         `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id"`
	Columns          []*ColumnInfo `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *TableInfo) Reset()                    { *m = TableInfo{} }
func (m *TableInfo) String() string            { return proto.CompactTextString(m) }
func (*TableInfo) ProtoMessage()               {}
func (*TableInfo) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{0} }

func (m *TableInfo) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *TableInfo) GetColumns() []*ColumnInfo {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ColumnInfo struct {
	ColumnId         int64    `protobuf:"varint,1,opt,name=column_id,json=columnId" json:"column_id"`
	Tp               int32    `protobuf:"varint,2,opt,name=tp" json:"tp"`
	Collation        int32    `protobuf:"varint,3,opt,name=collation" json:"collation"`
	ColumnLen        int32    `protobuf:"varint,4,opt,name=columnLen" json:"columnLen"`
	Decimal          int32    `protobuf:"varint,5,opt,name=decimal" json:"decimal"`
	Flag             int32    `protobuf:"varint,6,opt,name=flag" json:"flag"`
	Elems            []string `protobuf:"bytes,7,rep,name=elems" json:"elems,omitempty"`
	DefaultVal       []byte   `protobuf:"bytes,8,opt,name=default_val,json=defaultVal" json:"default_val,omitempty"`
	PkHandle         bool     `protobuf:"varint,21,opt,name=pk_handle,json=pkHandle" json:"pk_handle"`
	Array            bool     `protobuf:"varint,22,opt,name=array" json:"array"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ColumnInfo) Reset()                    { *m = ColumnInfo{} }
func (m *ColumnInfo) String() string            { return proto.CompactTextString(m) }
func (*ColumnInfo) ProtoMessage()               {}
func (*ColumnInfo) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{1} }

func (m *ColumnInfo) GetColumnId() int64 {
	if m != nil {
		return m.ColumnId
	}
	return 0
}

func (m *ColumnInfo) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

func (m *ColumnInfo) GetCollation() int32 {
	if m != nil {
		return m.Collation
	}
	return 0
}

func (m *ColumnInfo) GetColumnLen() int32 {
	if m != nil {
		return m.ColumnLen
	}
	return 0
}

func (m *ColumnInfo) GetDecimal() int32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *ColumnInfo) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *ColumnInfo) GetElems() []string {
	if m != nil {
		return m.Elems
	}
	return nil
}

func (m *ColumnInfo) GetDefaultVal() []byte {
	if m != nil {
		return m.DefaultVal
	}
	return nil
}

func (m *ColumnInfo) GetPkHandle() bool {
	if m != nil {
		return m.PkHandle
	}
	return false
}

func (m *ColumnInfo) GetArray() bool {
	if m != nil {
		return m.Array
	}
	return false
}

type IndexInfo struct {
	TableId          int64         `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id"`
	IndexId          int64         `protobuf:"varint,2,opt,name=index_id,json=indexId" json:"index_id"`
	Columns          []*ColumnInfo `protobuf:"bytes,3,rep,name=columns" json:"columns,omitempty"`
	Unique           bool          `protobuf:"varint,4,opt,name=unique" json:"unique"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *IndexInfo) Reset()                    { *m = IndexInfo{} }
func (m *IndexInfo) String() string            { return proto.CompactTextString(m) }
func (*IndexInfo) ProtoMessage()               {}
func (*IndexInfo) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{2} }

func (m *IndexInfo) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *IndexInfo) GetIndexId() int64 {
	if m != nil {
		return m.IndexId
	}
	return 0
}

func (m *IndexInfo) GetColumns() []*ColumnInfo {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *IndexInfo) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

// KeyRange is the encoded index key range, low is closed, high is open. (low <= x < high)
type KeyRange struct {
	Low              []byte `protobuf:"bytes,1,opt,name=low" json:"low,omitempty"`
	High             []byte `protobuf:"bytes,2,opt,name=high" json:"high,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{3} }

func (m *KeyRange) GetLow() []byte {
	if m != nil {
		return m.Low
	}
	return nil
}

func (m *KeyRange) GetHigh() []byte {
	if m != nil {
		return m.High
	}
	return nil
}

func init() {
	proto.RegisterType((*TableInfo)(nil), "tipb.TableInfo")
	proto.RegisterType((*ColumnInfo)(nil), "tipb.ColumnInfo")
	proto.RegisterType((*IndexInfo)(nil), "tipb.IndexInfo")
	proto.RegisterType((*KeyRange)(nil), "tipb.KeyRange")
}
func (m *TableInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.TableId))
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSchema(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ColumnInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ColumnInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.ColumnId))
	dAtA[i] = 0x10
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.Tp))
	dAtA[i] = 0x18
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.Collation))
	dAtA[i] = 0x20
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.ColumnLen))
	dAtA[i] = 0x28
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.Decimal))
	dAtA[i] = 0x30
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.Flag))
	if len(m.Elems) > 0 {
		for _, s := range m.Elems {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.DefaultVal != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.DefaultVal)))
		i += copy(dAtA[i:], m.DefaultVal)
	}
	dAtA[i] = 0xa8
	i++
	dAtA[i] = 0x1
	i++
	if m.PkHandle {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0xb0
	i++
	dAtA[i] = 0x1
	i++
	if m.Array {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *IndexInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.TableId))
	dAtA[i] = 0x10
	i++
	i = encodeVarintSchema(dAtA, i, uint64(m.IndexId))
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSchema(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x20
	i++
	if m.Unique {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *KeyRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyRange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Low != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Low)))
		i += copy(dAtA[i:], m.Low)
	}
	if m.High != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.High)))
		i += copy(dAtA[i:], m.High)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TableInfo) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.TableId))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ColumnInfo) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.ColumnId))
	n += 1 + sovSchema(uint64(m.Tp))
	n += 1 + sovSchema(uint64(m.Collation))
	n += 1 + sovSchema(uint64(m.ColumnLen))
	n += 1 + sovSchema(uint64(m.Decimal))
	n += 1 + sovSchema(uint64(m.Flag))
	if len(m.Elems) > 0 {
		for _, s := range m.Elems {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.DefaultVal != nil {
		l = len(m.DefaultVal)
		n += 1 + l + sovSchema(uint64(l))
	}
	n += 3
	n += 3
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IndexInfo) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.TableId))
	n += 1 + sovSchema(uint64(m.IndexId))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	n += 2
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *KeyRange) Size() (n int) {
	var l int
	_ = l
	if m.Low != nil {
		l = len(m.Low)
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.High != nil {
		l = len(m.High)
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSchema(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnInfo{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func (m *ColumnInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ColumnInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ColumnInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnId", wireType)
			}
			m.ColumnId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ColumnId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collation", wireType)
			}
			m.Collation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Collation |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnLen", wireType)
			}
			m.ColumnLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ColumnLen |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag", wireType)
			}
			m.Flag = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flag |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elems = append(m.Elems, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultVal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultVal = append(m.DefaultVal[:0], dAtA[iNdEx:postIndex]...)
			if m.DefaultVal == nil {
				m.DefaultVal = []byte{}
			}
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkHandle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PkHandle = bool(v != 0)
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Array", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Array = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func (m *IndexInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexId", wireType)
			}
			m.IndexId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnInfo{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unique", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Unique = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func (m *KeyRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Low = append(m.Low[:0], dAtA[iNdEx:postIndex]...)
			if m.Low == nil {
				m.Low = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.High = append(m.High[:0], dAtA[iNdEx:postIndex]...)
			if m.High == nil {
				m.High = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
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
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
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
					return 0, ErrIntOverflowSchema
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
					return 0, ErrIntOverflowSchema
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchema
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
				next, err := skipSchema(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthSchema = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("schema.proto", fileDescriptorSchema) }

var fileDescriptorSchema = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xd7, 0x49, 0xba, 0x4d, 0x67, 0x7b, 0x58, 0x59, 0x5d, 0x64, 0xad, 0x50, 0x1a, 0x72,
	0x0a, 0x1c, 0x02, 0xe2, 0x11, 0x96, 0x0b, 0x15, 0x1c, 0x50, 0x84, 0x10, 0xb7, 0xca, 0x8d, 0xdd,
	0xd4, 0x5a, 0xc7, 0x36, 0xad, 0x03, 0xec, 0x93, 0xc0, 0xcb, 0x70, 0xdf, 0x23, 0x4f, 0x80, 0x50,
	0x79, 0x11, 0x64, 0x27, 0xd1, 0x86, 0xbd, 0xb0, 0xb7, 0x99, 0xef, 0xff, 0x3d, 0xfa, 0xc7, 0x03,
	0xf3, 0x43, 0xb5, 0xe3, 0x0d, 0x2d, 0xcc, 0x5e, 0x5b, 0x8d, 0x23, 0x2b, 0xcc, 0xe6, 0x72, 0x51,
	0xeb, 0x5a, 0x7b, 0xf0, 0xdc, 0x55, 0x9d, 0x96, 0x7d, 0x84, 0xd9, 0x7b, 0xba, 0x91, 0x7c, 0xa5,
	0xb6, 0x1a, 0x2f, 0x21, 0xb6, 0xae, 0x59, 0x0b, 0x46, 0x50, 0x8a, 0xf2, 0xf0, 0x2a, 0xba, 0xfd,
	0xb5, 0x3c, 0x29, 0xa7, 0x9e, 0xae, 0x18, 0x7e, 0x06, 0xd3, 0x4a, 0xcb, 0xb6, 0x51, 0x07, 0x12,
	0xa4, 0x61, 0x7e, 0xf6, 0xf2, 0xbc, 0x70, 0xb3, 0x8b, 0x57, 0x1e, 0xba, 0x19, 0xe5, 0x60, 0xc8,
	0x7e, 0x04, 0x00, 0x77, 0x1c, 0x3f, 0x81, 0x59, 0xa7, 0xdc, 0x1f, 0x1e, 0x77, 0x78, 0xc5, 0xf0,
	0x02, 0x02, 0x6b, 0x48, 0x90, 0xa2, 0x7c, 0xd2, 0x6b, 0x81, 0x35, 0x38, 0xf3, 0x0f, 0x25, 0xb5,
	0x42, 0x2b, 0x12, 0x8e, 0xc4, 0x3b, 0xdc, 0x7b, 0xda, 0x46, 0xbd, 0xe5, 0x8a, 0x44, 0xf7, 0x3c,
	0x1d, 0xc6, 0x09, 0x4c, 0x19, 0xaf, 0x44, 0x43, 0x25, 0x99, 0x8c, 0x1c, 0x03, 0xc4, 0x04, 0xa2,
	0xad, 0xa4, 0x35, 0x39, 0x1d, 0x89, 0x9e, 0xe0, 0x05, 0x4c, 0xb8, 0xe4, 0xcd, 0x81, 0x4c, 0xd3,
	0x30, 0x9f, 0x95, 0x5d, 0x83, 0x97, 0x70, 0xc6, 0xf8, 0x96, 0xb6, 0xd2, 0xae, 0x3f, 0x53, 0x49,
	0xe2, 0x14, 0xe5, 0xf3, 0x12, 0x7a, 0xf4, 0x81, 0x4a, 0xb7, 0xb1, 0xb9, 0x5e, 0xef, 0xa8, 0x62,
	0x92, 0x93, 0x8b, 0x14, 0xe5, 0xf1, 0xb0, 0xb1, 0xb9, 0x7e, 0xed, 0x29, 0xbe, 0x84, 0x09, 0xdd,
	0xef, 0xe9, 0x0d, 0x79, 0x34, 0x92, 0x3b, 0x94, 0x7d, 0x43, 0x30, 0x5b, 0x29, 0xc6, 0xbf, 0x3e,
	0xec, 0x34, 0x4b, 0x88, 0x85, 0x73, 0x3b, 0x43, 0x30, 0x36, 0x78, 0xfa, 0xef, 0xed, 0xc2, 0xff,
	0xdc, 0x0e, 0x3f, 0x86, 0xd3, 0x56, 0x89, 0x4f, 0x2d, 0xf7, 0x9f, 0x39, 0x04, 0xeb, 0x59, 0xf6,
	0x02, 0xe2, 0x37, 0xfc, 0xa6, 0xa4, 0xaa, 0xe6, 0xf8, 0x1c, 0x42, 0xa9, 0xbf, 0xf8, 0x48, 0xf3,
	0xd2, 0x95, 0x18, 0x43, 0xb4, 0x13, 0xf5, 0xce, 0x87, 0x98, 0x97, 0xbe, 0xbe, 0x7a, 0x7a, 0x7b,
	0x4c, 0xd0, 0xcf, 0x63, 0x82, 0x7e, 0x1f, 0x13, 0xf4, 0xfd, 0x4f, 0x72, 0x02, 0x17, 0x95, 0x6e,
	0x0a, 0x23, 0x54, 0x5d, 0x51, 0x53, 0x58, 0xc1, 0x36, 0x3e, 0xcc, 0x3b, 0xf4, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x8b, 0xd2, 0x95, 0xbb, 0x02, 0x00, 0x00,
}
