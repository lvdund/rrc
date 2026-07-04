// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v15n0 (line 22573).

var mRDCParametersV15n0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraBandENDC-Support-UL", Optional: true},
	},
}

const (
	MRDC_Parameters_V15n0_IntraBandENDC_Support_UL_Non_Contiguous = 0
	MRDC_Parameters_V15n0_IntraBandENDC_Support_UL_Both           = 1
)

var mRDCParametersV15n0IntraBandENDCSupportULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_V15n0_IntraBandENDC_Support_UL_Non_Contiguous, MRDC_Parameters_V15n0_IntraBandENDC_Support_UL_Both},
}

type MRDC_Parameters_V15n0 struct {
	IntraBandENDC_Support_UL *int64
}

func (ie *MRDC_Parameters_V15n0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV15n0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraBandENDC_Support_UL != nil}); err != nil {
		return err
	}

	// 2. intraBandENDC-Support-UL: enumerated
	{
		if ie.IntraBandENDC_Support_UL != nil {
			if err := e.EncodeEnumerated(*ie.IntraBandENDC_Support_UL, mRDCParametersV15n0IntraBandENDCSupportULConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_V15n0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV15n0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intraBandENDC-Support-UL: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV15n0IntraBandENDCSupportULConstraints)
			if err != nil {
				return err
			}
			ie.IntraBandENDC_Support_UL = &idx
		}
	}

	return nil
}
