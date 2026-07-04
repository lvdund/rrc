// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T312-r16 (line 9494).
// T312-r16 ::=                        ENUMERATED { ms0, ms50, ms100, ms200, ms300, ms400, ms500, ms1000}

const (
	T312_r16_Ms0    = 0
	T312_r16_Ms50   = 1
	T312_r16_Ms100  = 2
	T312_r16_Ms200  = 3
	T312_r16_Ms300  = 4
	T312_r16_Ms400  = 5
	T312_r16_Ms500  = 6
	T312_r16_Ms1000 = 7
)

var t312R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T312_r16_Ms0, T312_r16_Ms50, T312_r16_Ms100, T312_r16_Ms200, T312_r16_Ms300, T312_r16_Ms400, T312_r16_Ms500, T312_r16_Ms1000},
}

type T312_r16 struct {
	Value int64
}

func (v *T312_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, t312R16Constraints)
}

func (v *T312_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(t312R16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
