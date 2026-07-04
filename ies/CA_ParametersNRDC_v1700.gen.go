// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v1700 (line 18418).

var cAParametersNRDCV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "simultaneousRxTx-IAB-MultipleParents-r17", Optional: true},
		{Name: "condPSCellAdditionNRDC-r17", Optional: true},
		{Name: "scg-ActivationDeactivationNRDC-r17", Optional: true},
		{Name: "scg-ActivationDeactivationResumeNRDC-r17", Optional: true},
		{Name: "beamManagementType-CBM-r17", Optional: true},
	},
}

const (
	CA_ParametersNRDC_v1700_SimultaneousRxTx_IAB_MultipleParents_r17_Supported = 0
)

var cAParametersNRDCV1700SimultaneousRxTxIABMultipleParentsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1700_SimultaneousRxTx_IAB_MultipleParents_r17_Supported},
}

const (
	CA_ParametersNRDC_v1700_CondPSCellAdditionNRDC_r17_Supported = 0
)

var cAParametersNRDCV1700CondPSCellAdditionNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1700_CondPSCellAdditionNRDC_r17_Supported},
}

const (
	CA_ParametersNRDC_v1700_Scg_ActivationDeactivationNRDC_r17_Supported = 0
)

var cAParametersNRDCV1700ScgActivationDeactivationNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1700_Scg_ActivationDeactivationNRDC_r17_Supported},
}

const (
	CA_ParametersNRDC_v1700_Scg_ActivationDeactivationResumeNRDC_r17_Supported = 0
)

var cAParametersNRDCV1700ScgActivationDeactivationResumeNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1700_Scg_ActivationDeactivationResumeNRDC_r17_Supported},
}

const (
	CA_ParametersNRDC_v1700_BeamManagementType_CBM_r17_Supported = 0
)

var cAParametersNRDCV1700BeamManagementTypeCBMR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNRDC_v1700_BeamManagementType_CBM_r17_Supported},
}

type CA_ParametersNRDC_v1700 struct {
	SimultaneousRxTx_IAB_MultipleParents_r17 *int64
	CondPSCellAdditionNRDC_r17               *int64
	Scg_ActivationDeactivationNRDC_r17       *int64
	Scg_ActivationDeactivationResumeNRDC_r17 *int64
	BeamManagementType_CBM_r17               *int64
}

func (ie *CA_ParametersNRDC_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SimultaneousRxTx_IAB_MultipleParents_r17 != nil, ie.CondPSCellAdditionNRDC_r17 != nil, ie.Scg_ActivationDeactivationNRDC_r17 != nil, ie.Scg_ActivationDeactivationResumeNRDC_r17 != nil, ie.BeamManagementType_CBM_r17 != nil}); err != nil {
		return err
	}

	// 2. simultaneousRxTx-IAB-MultipleParents-r17: enumerated
	{
		if ie.SimultaneousRxTx_IAB_MultipleParents_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousRxTx_IAB_MultipleParents_r17, cAParametersNRDCV1700SimultaneousRxTxIABMultipleParentsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. condPSCellAdditionNRDC-r17: enumerated
	{
		if ie.CondPSCellAdditionNRDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.CondPSCellAdditionNRDC_r17, cAParametersNRDCV1700CondPSCellAdditionNRDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. scg-ActivationDeactivationNRDC-r17: enumerated
	{
		if ie.Scg_ActivationDeactivationNRDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_ActivationDeactivationNRDC_r17, cAParametersNRDCV1700ScgActivationDeactivationNRDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. scg-ActivationDeactivationResumeNRDC-r17: enumerated
	{
		if ie.Scg_ActivationDeactivationResumeNRDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_ActivationDeactivationResumeNRDC_r17, cAParametersNRDCV1700ScgActivationDeactivationResumeNRDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. beamManagementType-CBM-r17: enumerated
	{
		if ie.BeamManagementType_CBM_r17 != nil {
			if err := e.EncodeEnumerated(*ie.BeamManagementType_CBM_r17, cAParametersNRDCV1700BeamManagementTypeCBMR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. simultaneousRxTx-IAB-MultipleParents-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1700SimultaneousRxTxIABMultipleParentsR17Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRxTx_IAB_MultipleParents_r17 = &idx
		}
	}

	// 3. condPSCellAdditionNRDC-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1700CondPSCellAdditionNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.CondPSCellAdditionNRDC_r17 = &idx
		}
	}

	// 4. scg-ActivationDeactivationNRDC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1700ScgActivationDeactivationNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.Scg_ActivationDeactivationNRDC_r17 = &idx
		}
	}

	// 5. scg-ActivationDeactivationResumeNRDC-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1700ScgActivationDeactivationResumeNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.Scg_ActivationDeactivationResumeNRDC_r17 = &idx
		}
	}

	// 6. beamManagementType-CBM-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersNRDCV1700BeamManagementTypeCBMR17Constraints)
			if err != nil {
				return err
			}
			ie.BeamManagementType_CBM_r17 = &idx
		}
	}

	return nil
}
