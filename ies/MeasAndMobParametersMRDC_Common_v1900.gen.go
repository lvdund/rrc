// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-Common-v1900 (line 21451).

var measAndMobParametersMRDCCommonV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "threeCarrierMeasWithoutGap-r19", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_Common_v1900_ThreeCarrierMeasWithoutGap_r19_Supported = 0
)

var measAndMobParametersMRDCCommonV1900ThreeCarrierMeasWithoutGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_Common_v1900_ThreeCarrierMeasWithoutGap_r19_Supported},
}

type MeasAndMobParametersMRDC_Common_v1900 struct {
	ThreeCarrierMeasWithoutGap_r19 *int64
}

func (ie *MeasAndMobParametersMRDC_Common_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ThreeCarrierMeasWithoutGap_r19 != nil}); err != nil {
		return err
	}

	// 2. threeCarrierMeasWithoutGap-r19: enumerated
	{
		if ie.ThreeCarrierMeasWithoutGap_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ThreeCarrierMeasWithoutGap_r19, measAndMobParametersMRDCCommonV1900ThreeCarrierMeasWithoutGapR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. threeCarrierMeasWithoutGap-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCCommonV1900ThreeCarrierMeasWithoutGapR19Constraints)
			if err != nil {
				return err
			}
			ie.ThreeCarrierMeasWithoutGap_r19 = &idx
		}
	}

	return nil
}
