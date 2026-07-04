// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB-Type-r17 (line 3208).

const (
	SIB_Type_r17_SibType2     = 0
	SIB_Type_r17_SibType3     = 1
	SIB_Type_r17_SibType4     = 2
	SIB_Type_r17_SibType5     = 3
	SIB_Type_r17_SibType9     = 4
	SIB_Type_r17_SibType10    = 5
	SIB_Type_r17_SibType11    = 6
	SIB_Type_r17_SibType12    = 7
	SIB_Type_r17_SibType13    = 8
	SIB_Type_r17_SibType14    = 9
	SIB_Type_r17_PosSIB_v1810 = 10
	SIB_Type_r17_Spare5       = 11
	SIB_Type_r17_Spare4       = 12
	SIB_Type_r17_Spare3       = 13
	SIB_Type_r17_Spare2       = 14
	SIB_Type_r17_Spare1       = 15
)

var sIBTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB_Type_r17_SibType2, SIB_Type_r17_SibType3, SIB_Type_r17_SibType4, SIB_Type_r17_SibType5, SIB_Type_r17_SibType9, SIB_Type_r17_SibType10, SIB_Type_r17_SibType11, SIB_Type_r17_SibType12, SIB_Type_r17_SibType13, SIB_Type_r17_SibType14, SIB_Type_r17_PosSIB_v1810, SIB_Type_r17_Spare5, SIB_Type_r17_Spare4, SIB_Type_r17_Spare3, SIB_Type_r17_Spare2, SIB_Type_r17_Spare1},
}

type SIB_Type_r17 struct {
	Value int64
}

func (v *SIB_Type_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sIBTypeR17Constraints)
}

func (v *SIB_Type_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sIBTypeR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
