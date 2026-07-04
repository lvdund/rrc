// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T-ReassemblyExt-r17 (line 14179).
// T-ReassemblyExt-r17 ::=             ENUMERATED {ms210, ms220, ms340, ms350, ms550, ms1100, ms1650, ms2200}

const (
	T_ReassemblyExt_r17_Ms210  = 0
	T_ReassemblyExt_r17_Ms220  = 1
	T_ReassemblyExt_r17_Ms340  = 2
	T_ReassemblyExt_r17_Ms350  = 3
	T_ReassemblyExt_r17_Ms550  = 4
	T_ReassemblyExt_r17_Ms1100 = 5
	T_ReassemblyExt_r17_Ms1650 = 6
	T_ReassemblyExt_r17_Ms2200 = 7
)

var tReassemblyExtR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T_ReassemblyExt_r17_Ms210, T_ReassemblyExt_r17_Ms220, T_ReassemblyExt_r17_Ms340, T_ReassemblyExt_r17_Ms350, T_ReassemblyExt_r17_Ms550, T_ReassemblyExt_r17_Ms1100, T_ReassemblyExt_r17_Ms1650, T_ReassemblyExt_r17_Ms2200},
}

type T_ReassemblyExt_r17 struct {
	Value int64
}

func (v *T_ReassemblyExt_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, tReassemblyExtR17Constraints)
}

func (v *T_ReassemblyExt_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(tReassemblyExtR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
