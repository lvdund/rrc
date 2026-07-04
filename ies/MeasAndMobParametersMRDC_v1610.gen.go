// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-v1610 (line 21381).

var measAndMobParametersMRDCV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-Common-v1610", Optional: true},
		{Name: "interNR-MeasEUTRA-IAB-r16", Optional: true},
	},
}

const (
	MeasAndMobParametersMRDC_v1610_InterNR_MeasEUTRA_IAB_r16_Supported = 0
)

var measAndMobParametersMRDCV1610InterNRMeasEUTRAIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersMRDC_v1610_InterNR_MeasEUTRA_IAB_r16_Supported},
}

type MeasAndMobParametersMRDC_v1610 struct {
	MeasAndMobParametersMRDC_Common_v1610 *MeasAndMobParametersMRDC_Common_v1610
	InterNR_MeasEUTRA_IAB_r16             *int64
}

func (ie *MeasAndMobParametersMRDC_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersMRDC_Common_v1610 != nil, ie.InterNR_MeasEUTRA_IAB_r16 != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1610: ref
	{
		if ie.MeasAndMobParametersMRDC_Common_v1610 != nil {
			if err := ie.MeasAndMobParametersMRDC_Common_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. interNR-MeasEUTRA-IAB-r16: enumerated
	{
		if ie.InterNR_MeasEUTRA_IAB_r16 != nil {
			if err := e.EncodeEnumerated(*ie.InterNR_MeasEUTRA_IAB_r16, measAndMobParametersMRDCV1610InterNRMeasEUTRAIABR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersMRDC-Common-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersMRDC_Common_v1610 = new(MeasAndMobParametersMRDC_Common_v1610)
			if err := ie.MeasAndMobParametersMRDC_Common_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. interNR-MeasEUTRA-IAB-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersMRDCV1610InterNRMeasEUTRAIABR16Constraints)
			if err != nil {
				return err
			}
			ie.InterNR_MeasEUTRA_IAB_r16 = &idx
		}
	}

	return nil
}
