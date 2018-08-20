package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Foo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.Bar, err = dc.ReadString()
	if err != nil {
		return
	}
	z.Baz, err = dc.ReadFloat64()
	if err != nil {
		return
	}
	{
		var zb0002 []byte
		zb0002, err = dc.ReadBytes([]byte(z.Address))
		if err != nil {
			return
		}
		z.Address = Hash(zb0002)
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Parents) >= int(zb0003) {
		z.Parents = (z.Parents)[:zb0003]
	} else {
		z.Parents = make([]Hash, zb0003)
	}
	for za0001 := range z.Parents {
		{
			var zb0004 []byte
			zb0004, err = dc.ReadBytes([]byte(z.Parents[za0001]))
			if err != nil {
				return
			}
			z.Parents[za0001] = Hash(zb0004)
		}
	}
	var zb0005 uint32
	zb0005, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	if z.KV == nil {
		z.KV = make(map[string]float64, zb0005)
	} else if len(z.KV) > 0 {
		for key := range z.KV {
			delete(z.KV, key)
		}
	}
	for zb0005 > 0 {
		zb0005--
		var za0002 string
		var za0003 float64
		za0002, err = dc.ReadString()
		if err != nil {
			return
		}
		za0003, err = dc.ReadFloat64()
		if err != nil {
			return
		}
		z.KV[za0002] = za0003
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Foo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bar)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Baz)
	if err != nil {
		return
	}
	err = en.WriteBytes([]byte(z.Address))
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Parents)))
	if err != nil {
		return
	}
	for za0001 := range z.Parents {
		err = en.WriteBytes([]byte(z.Parents[za0001]))
		if err != nil {
			return
		}
	}
	err = en.WriteMapHeader(uint32(len(z.KV)))
	if err != nil {
		return
	}
	for za0002, za0003 := range z.KV {
		err = en.WriteString(za0002)
		if err != nil {
			return
		}
		err = en.WriteFloat64(za0003)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Foo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	o = msgp.AppendString(o, z.Bar)
	o = msgp.AppendFloat64(o, z.Baz)
	o = msgp.AppendBytes(o, []byte(z.Address))
	o = msgp.AppendArrayHeader(o, uint32(len(z.Parents)))
	for za0001 := range z.Parents {
		o = msgp.AppendBytes(o, []byte(z.Parents[za0001]))
	}
	o = msgp.AppendMapHeader(o, uint32(len(z.KV)))
	for za0002, za0003 := range z.KV {
		o = msgp.AppendString(o, za0002)
		o = msgp.AppendFloat64(o, za0003)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Foo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.Bar, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.Baz, bts, err = msgp.ReadFloat64Bytes(bts)
	if err != nil {
		return
	}
	{
		var zb0002 []byte
		zb0002, bts, err = msgp.ReadBytesBytes(bts, []byte(z.Address))
		if err != nil {
			return
		}
		z.Address = Hash(zb0002)
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Parents) >= int(zb0003) {
		z.Parents = (z.Parents)[:zb0003]
	} else {
		z.Parents = make([]Hash, zb0003)
	}
	for za0001 := range z.Parents {
		{
			var zb0004 []byte
			zb0004, bts, err = msgp.ReadBytesBytes(bts, []byte(z.Parents[za0001]))
			if err != nil {
				return
			}
			z.Parents[za0001] = Hash(zb0004)
		}
	}
	var zb0005 uint32
	zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	if z.KV == nil {
		z.KV = make(map[string]float64, zb0005)
	} else if len(z.KV) > 0 {
		for key := range z.KV {
			delete(z.KV, key)
		}
	}
	for zb0005 > 0 {
		var za0002 string
		var za0003 float64
		zb0005--
		za0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		za0003, bts, err = msgp.ReadFloat64Bytes(bts)
		if err != nil {
			return
		}
		z.KV[za0002] = za0003
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Foo) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Bar) + msgp.Float64Size + msgp.BytesPrefixSize + len([]byte(z.Address)) + msgp.ArrayHeaderSize
	for za0001 := range z.Parents {
		s += msgp.BytesPrefixSize + len([]byte(z.Parents[za0001]))
	}
	s += msgp.MapHeaderSize
	if z.KV != nil {
		for za0002, za0003 := range z.KV {
			_ = za0003
			s += msgp.StringPrefixSize + len(za0002) + msgp.Float64Size
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Hash) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 []byte
		zb0001, err = dc.ReadBytes([]byte((*z)))
		if err != nil {
			return
		}
		(*z) = Hash(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Hash) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes([]byte(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Hash) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, []byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Hash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 []byte
		zb0001, bts, err = msgp.ReadBytesBytes(bts, []byte((*z)))
		if err != nil {
			return
		}
		(*z) = Hash(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Hash) Msgsize() (s int) {
	s = msgp.BytesPrefixSize + len([]byte(z))
	return
}