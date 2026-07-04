// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntraBandPowerClass-r16 (line 17125).
// IntraBandPowerClass-r16 ::=         ENUMERATED {pc2, pc3, spare6, spare5, spare4, spare3, spare2, spare1}

const (
	IntraBandPowerClass_r16_Pc2    = 0
	IntraBandPowerClass_r16_Pc3    = 1
	IntraBandPowerClass_r16_Spare6 = 2
	IntraBandPowerClass_r16_Spare5 = 3
	IntraBandPowerClass_r16_Spare4 = 4
	IntraBandPowerClass_r16_Spare3 = 5
	IntraBandPowerClass_r16_Spare2 = 6
	IntraBandPowerClass_r16_Spare1 = 7
)

var intraBandPowerClassR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IntraBandPowerClass_r16_Pc2, IntraBandPowerClass_r16_Pc3, IntraBandPowerClass_r16_Spare6, IntraBandPowerClass_r16_Spare5, IntraBandPowerClass_r16_Spare4, IntraBandPowerClass_r16_Spare3, IntraBandPowerClass_r16_Spare2, IntraBandPowerClass_r16_Spare1},
}

type IntraBandPowerClass_r16 struct {
	Value int64
}

func (v *IntraBandPowerClass_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, intraBandPowerClassR16Constraints)
}

func (v *IntraBandPowerClass_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(intraBandPowerClassR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
