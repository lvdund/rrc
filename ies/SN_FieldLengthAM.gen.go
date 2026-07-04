// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SN-FieldLengthAM (line 14140).
// SN-FieldLengthAM ::=                ENUMERATED {size12, size18}

const (
	SN_FieldLengthAM_Size12 = 0
	SN_FieldLengthAM_Size18 = 1
)

var sNFieldLengthAMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SN_FieldLengthAM_Size12, SN_FieldLengthAM_Size18},
}

type SN_FieldLengthAM struct {
	Value int64
}

func (v *SN_FieldLengthAM) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sNFieldLengthAMConstraints)
}

func (v *SN_FieldLengthAM) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sNFieldLengthAMConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
