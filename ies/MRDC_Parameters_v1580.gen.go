// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1580 (line 22561).

var mRDCParametersV1580Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dynamicPowerSharingNEDC", Optional: true},
	},
}

const (
	MRDC_Parameters_v1580_DynamicPowerSharingNEDC_Supported = 0
)

var mRDCParametersV1580DynamicPowerSharingNEDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1580_DynamicPowerSharingNEDC_Supported},
}

type MRDC_Parameters_v1580 struct {
	DynamicPowerSharingNEDC *int64
}

func (ie *MRDC_Parameters_v1580) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1580Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DynamicPowerSharingNEDC != nil}); err != nil {
		return err
	}

	// 2. dynamicPowerSharingNEDC: enumerated
	{
		if ie.DynamicPowerSharingNEDC != nil {
			if err := e.EncodeEnumerated(*ie.DynamicPowerSharingNEDC, mRDCParametersV1580DynamicPowerSharingNEDCConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1580) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1580Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dynamicPowerSharingNEDC: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1580DynamicPowerSharingNEDCConstraints)
			if err != nil {
				return err
			}
			ie.DynamicPowerSharingNEDC = &idx
		}
	}

	return nil
}
