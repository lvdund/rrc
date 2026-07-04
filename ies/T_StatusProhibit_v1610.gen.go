// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T-StatusProhibit-v1610 (line 14177).
// T-StatusProhibit-v1610 ::=          ENUMERATED { ms1, ms2, ms3, ms4, spare4, spare3, spare2, spare1}

const (
	T_StatusProhibit_v1610_Ms1    = 0
	T_StatusProhibit_v1610_Ms2    = 1
	T_StatusProhibit_v1610_Ms3    = 2
	T_StatusProhibit_v1610_Ms4    = 3
	T_StatusProhibit_v1610_Spare4 = 4
	T_StatusProhibit_v1610_Spare3 = 5
	T_StatusProhibit_v1610_Spare2 = 6
	T_StatusProhibit_v1610_Spare1 = 7
)

var tStatusProhibitV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T_StatusProhibit_v1610_Ms1, T_StatusProhibit_v1610_Ms2, T_StatusProhibit_v1610_Ms3, T_StatusProhibit_v1610_Ms4, T_StatusProhibit_v1610_Spare4, T_StatusProhibit_v1610_Spare3, T_StatusProhibit_v1610_Spare2, T_StatusProhibit_v1610_Spare1},
}

type T_StatusProhibit_v1610 struct {
	Value int64
}

func (v *T_StatusProhibit_v1610) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, tStatusProhibitV1610Constraints)
}

func (v *T_StatusProhibit_v1610) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(tStatusProhibitV1610Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
