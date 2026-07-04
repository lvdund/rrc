// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GeneralParametersMRDC-v1610 (line 25634).

var generalParametersMRDCV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "f1c-OverEUTRA-r16", Optional: true},
	},
}

const (
	GeneralParametersMRDC_v1610_F1c_OverEUTRA_r16_Supported = 0
)

var generalParametersMRDCV1610F1cOverEUTRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GeneralParametersMRDC_v1610_F1c_OverEUTRA_r16_Supported},
}

type GeneralParametersMRDC_v1610 struct {
	F1c_OverEUTRA_r16 *int64
}

func (ie *GeneralParametersMRDC_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(generalParametersMRDCV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.F1c_OverEUTRA_r16 != nil}); err != nil {
		return err
	}

	// 2. f1c-OverEUTRA-r16: enumerated
	{
		if ie.F1c_OverEUTRA_r16 != nil {
			if err := e.EncodeEnumerated(*ie.F1c_OverEUTRA_r16, generalParametersMRDCV1610F1cOverEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *GeneralParametersMRDC_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(generalParametersMRDCV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. f1c-OverEUTRA-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(generalParametersMRDCV1610F1cOverEUTRAR16Constraints)
			if err != nil {
				return err
			}
			ie.F1c_OverEUTRA_r16 = &idx
		}
	}

	return nil
}
