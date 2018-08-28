package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/annchain/OG/common/math"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (t *Tx) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = t.TxBase.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = t.From.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = t.To.DecodeMsg(dc)
	if err != nil {
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			return
		}
		t.Value = nil
	} else {
		if t.Value == nil {
			t.Value = new(math.BigInt)
		}
		err = t.Value.DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (t *Tx) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = t.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	err = t.From.EncodeMsg(en)
	if err != nil {
		return
	}
	err = t.To.EncodeMsg(en)
	if err != nil {
		return
	}
	if t.Value == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = t.Value.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (t *Tx) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, t.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = t.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = t.From.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = t.To.MarshalMsg(o)
	if err != nil {
		return
	}
	if t.Value == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = t.Value.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (t *Tx) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = t.TxBase.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = t.From.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = t.To.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		t.Value = nil
	} else {
		if t.Value == nil {
			t.Value = new(math.BigInt)
		}
		bts, err = t.Value.UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (t *Tx) Msgsize() (s int) {
	s = 1 + t.TxBase.Msgsize() + t.From.Msgsize() + t.To.Msgsize()
	if t.Value == nil {
		s += msgp.NilSize
	} else {
		s += t.Value.Msgsize()
	}
	return
}
