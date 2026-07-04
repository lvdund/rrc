// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NRDC-Parameters-v1700 (line 22669).

var nRDCParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "f1c-OverNR-RRC-r17", Optional: true},
		{Name: "measAndMobParametersNRDC-v1700"},
	},
}

const (
	NRDC_Parameters_v1700_F1c_OverNR_RRC_r17_Supported = 0
)

var nRDCParametersV1700F1cOverNRRRCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NRDC_Parameters_v1700_F1c_OverNR_RRC_r17_Supported},
}

type NRDC_Parameters_v1700 struct {
	F1c_OverNR_RRC_r17             *int64
	MeasAndMobParametersNRDC_v1700 MeasAndMobParametersMRDC_v1700
}

func (ie *NRDC_Parameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDCParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.F1c_OverNR_RRC_r17 != nil}); err != nil {
		return err
	}

	// 2. f1c-OverNR-RRC-r17: enumerated
	{
		if ie.F1c_OverNR_RRC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.F1c_OverNR_RRC_r17, nRDCParametersV1700F1cOverNRRRCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersNRDC-v1700: ref
	{
		if err := ie.MeasAndMobParametersNRDC_v1700.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NRDC_Parameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDCParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. f1c-OverNR-RRC-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(nRDCParametersV1700F1cOverNRRRCR17Constraints)
			if err != nil {
				return err
			}
			ie.F1c_OverNR_RRC_r17 = &idx
		}
	}

	// 3. measAndMobParametersNRDC-v1700: ref
	{
		if err := ie.MeasAndMobParametersNRDC_v1700.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
