// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1590 (line 22565).

var mRDCParametersV1590Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "interBandContiguousMRDC", Optional: true},
	},
}

const (
	MRDC_Parameters_v1590_InterBandContiguousMRDC_Supported = 0
)

var mRDCParametersV1590InterBandContiguousMRDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1590_InterBandContiguousMRDC_Supported},
}

type MRDC_Parameters_v1590 struct {
	InterBandContiguousMRDC *int64
}

func (ie *MRDC_Parameters_v1590) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1590Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterBandContiguousMRDC != nil}); err != nil {
		return err
	}

	// 2. interBandContiguousMRDC: enumerated
	{
		if ie.InterBandContiguousMRDC != nil {
			if err := e.EncodeEnumerated(*ie.InterBandContiguousMRDC, mRDCParametersV1590InterBandContiguousMRDCConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1590) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1590Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. interBandContiguousMRDC: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1590InterBandContiguousMRDCConstraints)
			if err != nil {
				return err
			}
			ie.InterBandContiguousMRDC = &idx
		}
	}

	return nil
}
