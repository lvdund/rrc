// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB-Type-r18 (line 3211).

const (
	SIB_Type_r18_SibType15 = 0
	SIB_Type_r18_SibType16 = 1
	SIB_Type_r18_SibType17 = 2
	SIB_Type_r18_SibType18 = 3
	SIB_Type_r18_SibType19 = 4
	SIB_Type_r18_SibType20 = 5
	SIB_Type_r18_SibType21 = 6
	SIB_Type_r18_SibType22 = 7
	SIB_Type_r18_SibType23 = 8
	SIB_Type_r18_SibType24 = 9
	SIB_Type_r18_SibType25 = 10
	SIB_Type_r18_Spare5    = 11
	SIB_Type_r18_Spare4    = 12
	SIB_Type_r18_Spare3    = 13
	SIB_Type_r18_Spare2    = 14
	SIB_Type_r18_Spare1    = 15
)

var sIBTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB_Type_r18_SibType15, SIB_Type_r18_SibType16, SIB_Type_r18_SibType17, SIB_Type_r18_SibType18, SIB_Type_r18_SibType19, SIB_Type_r18_SibType20, SIB_Type_r18_SibType21, SIB_Type_r18_SibType22, SIB_Type_r18_SibType23, SIB_Type_r18_SibType24, SIB_Type_r18_SibType25, SIB_Type_r18_Spare5, SIB_Type_r18_Spare4, SIB_Type_r18_Spare3, SIB_Type_r18_Spare2, SIB_Type_r18_Spare1},
}

type SIB_Type_r18 struct {
	Value int64
}

func (v *SIB_Type_r18) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sIBTypeR18Constraints)
}

func (v *SIB_Type_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sIBTypeR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
