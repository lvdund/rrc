// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SN-FieldLengthUM (line 14139).
// SN-FieldLengthUM ::=                ENUMERATED {size6, size12}

const (
	SN_FieldLengthUM_Size6  = 0
	SN_FieldLengthUM_Size12 = 1
)

var sNFieldLengthUMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SN_FieldLengthUM_Size6, SN_FieldLengthUM_Size12},
}

type SN_FieldLengthUM struct {
	Value int64
}

func (v *SN_FieldLengthUM) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sNFieldLengthUMConstraints)
}

func (v *SN_FieldLengthUM) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sNFieldLengthUMConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
