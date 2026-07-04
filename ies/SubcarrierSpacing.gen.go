// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SubcarrierSpacing (line 15940).
// SubcarrierSpacing ::=               ENUMERATED {kHz15, kHz30, kHz60, kHz120, kHz240, kHz480-v1700, kHz960-v1700, spare1}

const (
	SubcarrierSpacing_KHz15        = 0
	SubcarrierSpacing_KHz30        = 1
	SubcarrierSpacing_KHz60        = 2
	SubcarrierSpacing_KHz120       = 3
	SubcarrierSpacing_KHz240       = 4
	SubcarrierSpacing_KHz480_v1700 = 5
	SubcarrierSpacing_KHz960_v1700 = 6
	SubcarrierSpacing_Spare1       = 7
)

var subcarrierSpacingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SubcarrierSpacing_KHz15, SubcarrierSpacing_KHz30, SubcarrierSpacing_KHz60, SubcarrierSpacing_KHz120, SubcarrierSpacing_KHz240, SubcarrierSpacing_KHz480_v1700, SubcarrierSpacing_KHz960_v1700, SubcarrierSpacing_Spare1},
}

type SubcarrierSpacing struct {
	Value int64
}

func (v *SubcarrierSpacing) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, subcarrierSpacingConstraints)
}

func (v *SubcarrierSpacing) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(subcarrierSpacingConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
