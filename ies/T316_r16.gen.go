// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T316-r16 (line 1098).
// T316-r16 ::=         ENUMERATED {ms50, ms100, ms200, ms300, ms400, ms500, ms600, ms1000, ms1500, ms2000}

const (
	T316_r16_Ms50   = 0
	T316_r16_Ms100  = 1
	T316_r16_Ms200  = 2
	T316_r16_Ms300  = 3
	T316_r16_Ms400  = 4
	T316_r16_Ms500  = 5
	T316_r16_Ms600  = 6
	T316_r16_Ms1000 = 7
	T316_r16_Ms1500 = 8
	T316_r16_Ms2000 = 9
)

var t316R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T316_r16_Ms50, T316_r16_Ms100, T316_r16_Ms200, T316_r16_Ms300, T316_r16_Ms400, T316_r16_Ms500, T316_r16_Ms600, T316_r16_Ms1000, T316_r16_Ms1500, T316_r16_Ms2000},
}

type T316_r16 struct {
	Value int64
}

func (v *T316_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, t316R16Constraints)
}

func (v *T316_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(t316R16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
