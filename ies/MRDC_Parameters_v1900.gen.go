// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1900 (line 22628).

var mRDCParametersV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "interBandMRDC-WithOverlapDL-Bands-r19", Optional: true},
	},
}

const (
	MRDC_Parameters_v1900_InterBandMRDC_WithOverlapDL_Bands_r19_Supported = 0
)

var mRDCParametersV1900InterBandMRDCWithOverlapDLBandsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1900_InterBandMRDC_WithOverlapDL_Bands_r19_Supported},
}

type MRDC_Parameters_v1900 struct {
	InterBandMRDC_WithOverlapDL_Bands_r19 *int64
}

func (ie *MRDC_Parameters_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterBandMRDC_WithOverlapDL_Bands_r19 != nil}); err != nil {
		return err
	}

	// 2. interBandMRDC-WithOverlapDL-Bands-r19: enumerated
	{
		if ie.InterBandMRDC_WithOverlapDL_Bands_r19 != nil {
			if err := e.EncodeEnumerated(*ie.InterBandMRDC_WithOverlapDL_Bands_r19, mRDCParametersV1900InterBandMRDCWithOverlapDLBandsR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. interBandMRDC-WithOverlapDL-Bands-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1900InterBandMRDCWithOverlapDLBandsR19Constraints)
			if err != nil {
				return err
			}
			ie.InterBandMRDC_WithOverlapDL_Bands_r19 = &idx
		}
	}

	return nil
}
