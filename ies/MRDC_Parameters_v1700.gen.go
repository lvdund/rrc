// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1700 (line 22608).

var mRDCParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "condPSCellAdditionENDC-r17", Optional: true},
		{Name: "scg-ActivationDeactivationENDC-r17", Optional: true},
		{Name: "scg-ActivationDeactivationResumeENDC-r17", Optional: true},
	},
}

const (
	MRDC_Parameters_v1700_CondPSCellAdditionENDC_r17_Supported = 0
)

var mRDCParametersV1700CondPSCellAdditionENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1700_CondPSCellAdditionENDC_r17_Supported},
}

const (
	MRDC_Parameters_v1700_Scg_ActivationDeactivationENDC_r17_Supported = 0
)

var mRDCParametersV1700ScgActivationDeactivationENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1700_Scg_ActivationDeactivationENDC_r17_Supported},
}

const (
	MRDC_Parameters_v1700_Scg_ActivationDeactivationResumeENDC_r17_Supported = 0
)

var mRDCParametersV1700ScgActivationDeactivationResumeENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1700_Scg_ActivationDeactivationResumeENDC_r17_Supported},
}

type MRDC_Parameters_v1700 struct {
	CondPSCellAdditionENDC_r17               *int64
	Scg_ActivationDeactivationENDC_r17       *int64
	Scg_ActivationDeactivationResumeENDC_r17 *int64
}

func (ie *MRDC_Parameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CondPSCellAdditionENDC_r17 != nil, ie.Scg_ActivationDeactivationENDC_r17 != nil, ie.Scg_ActivationDeactivationResumeENDC_r17 != nil}); err != nil {
		return err
	}

	// 2. condPSCellAdditionENDC-r17: enumerated
	{
		if ie.CondPSCellAdditionENDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.CondPSCellAdditionENDC_r17, mRDCParametersV1700CondPSCellAdditionENDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. scg-ActivationDeactivationENDC-r17: enumerated
	{
		if ie.Scg_ActivationDeactivationENDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_ActivationDeactivationENDC_r17, mRDCParametersV1700ScgActivationDeactivationENDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. scg-ActivationDeactivationResumeENDC-r17: enumerated
	{
		if ie.Scg_ActivationDeactivationResumeENDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_ActivationDeactivationResumeENDC_r17, mRDCParametersV1700ScgActivationDeactivationResumeENDCR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. condPSCellAdditionENDC-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1700CondPSCellAdditionENDCR17Constraints)
			if err != nil {
				return err
			}
			ie.CondPSCellAdditionENDC_r17 = &idx
		}
	}

	// 3. scg-ActivationDeactivationENDC-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1700ScgActivationDeactivationENDCR17Constraints)
			if err != nil {
				return err
			}
			ie.Scg_ActivationDeactivationENDC_r17 = &idx
		}
	}

	// 4. scg-ActivationDeactivationResumeENDC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1700ScgActivationDeactivationResumeENDCR17Constraints)
			if err != nil {
				return err
			}
			ie.Scg_ActivationDeactivationResumeENDC_r17 = &idx
		}
	}

	return nil
}
