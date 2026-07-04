// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1770 (line 22614).

var mRDCParametersV1770Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "higherPowerLimitMRDC-r17", Optional: true},
	},
}

const (
	MRDC_Parameters_v1770_HigherPowerLimitMRDC_r17_Supported = 0
)

var mRDCParametersV1770HigherPowerLimitMRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1770_HigherPowerLimitMRDC_r17_Supported},
}

type MRDC_Parameters_v1770 struct {
	HigherPowerLimitMRDC_r17 *int64
}

func (ie *MRDC_Parameters_v1770) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1770Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.HigherPowerLimitMRDC_r17 != nil}); err != nil {
		return err
	}

	// 2. higherPowerLimitMRDC-r17: enumerated
	{
		if ie.HigherPowerLimitMRDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HigherPowerLimitMRDC_r17, mRDCParametersV1770HigherPowerLimitMRDCR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1770) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1770Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. higherPowerLimitMRDC-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1770HigherPowerLimitMRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.HigherPowerLimitMRDC_r17 = &idx
		}
	}

	return nil
}
