// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1840 (line 22624).

var mRDCParametersV1840Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraBandENDC-NominalSpacing-r18", Optional: true},
	},
}

const (
	MRDC_Parameters_v1840_IntraBandENDC_NominalSpacing_r18_Supported = 0
)

var mRDCParametersV1840IntraBandENDCNominalSpacingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1840_IntraBandENDC_NominalSpacing_r18_Supported},
}

type MRDC_Parameters_v1840 struct {
	IntraBandENDC_NominalSpacing_r18 *int64
}

func (ie *MRDC_Parameters_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1840Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraBandENDC_NominalSpacing_r18 != nil}); err != nil {
		return err
	}

	// 2. intraBandENDC-NominalSpacing-r18: enumerated
	{
		if ie.IntraBandENDC_NominalSpacing_r18 != nil {
			if err := e.EncodeEnumerated(*ie.IntraBandENDC_NominalSpacing_r18, mRDCParametersV1840IntraBandENDCNominalSpacingR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1840Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intraBandENDC-NominalSpacing-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1840IntraBandENDCNominalSpacingR18Constraints)
			if err != nil {
				return err
			}
			ie.IntraBandENDC_NominalSpacing_r18 = &idx
		}
	}

	return nil
}
