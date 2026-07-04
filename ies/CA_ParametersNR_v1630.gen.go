// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1630 (line 17378).

var cAParametersNRV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "simulTX-SRS-AntSwitchingInterBandUL-CA-r16", Optional: true},
		{Name: "beamManagementType-r16", Optional: true},
		{Name: "intraBandFreqSeparationUL-AggBW-GapBW-r16", Optional: true},
		{Name: "interCA-NonAlignedFrame-B-r16", Optional: true},
	},
}

const (
	CA_ParametersNR_v1630_BeamManagementType_r16_Ibm   = 0
	CA_ParametersNR_v1630_BeamManagementType_r16_Dummy = 1
)

var cAParametersNRV1630BeamManagementTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1630_BeamManagementType_r16_Ibm, CA_ParametersNR_v1630_BeamManagementType_r16_Dummy},
}

const (
	CA_ParametersNR_v1630_IntraBandFreqSeparationUL_AggBW_GapBW_r16_ClassI   = 0
	CA_ParametersNR_v1630_IntraBandFreqSeparationUL_AggBW_GapBW_r16_ClassII  = 1
	CA_ParametersNR_v1630_IntraBandFreqSeparationUL_AggBW_GapBW_r16_ClassIII = 2
)

var cAParametersNRV1630IntraBandFreqSeparationULAggBWGapBWR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1630_IntraBandFreqSeparationUL_AggBW_GapBW_r16_ClassI, CA_ParametersNR_v1630_IntraBandFreqSeparationUL_AggBW_GapBW_r16_ClassII, CA_ParametersNR_v1630_IntraBandFreqSeparationUL_AggBW_GapBW_r16_ClassIII},
}

const (
	CA_ParametersNR_v1630_InterCA_NonAlignedFrame_B_r16_Supported = 0
)

var cAParametersNRV1630InterCANonAlignedFrameBR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1630_InterCA_NonAlignedFrame_B_r16_Supported},
}

type CA_ParametersNR_v1630 struct {
	SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 *SimulSRS_ForAntennaSwitching_r16
	BeamManagementType_r16                     *int64
	IntraBandFreqSeparationUL_AggBW_GapBW_r16  *int64
	InterCA_NonAlignedFrame_B_r16              *int64
}

func (ie *CA_ParametersNR_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1630Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 != nil, ie.BeamManagementType_r16 != nil, ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16 != nil, ie.InterCA_NonAlignedFrame_B_r16 != nil}); err != nil {
		return err
	}

	// 2. simulTX-SRS-AntSwitchingInterBandUL-CA-r16: ref
	{
		if ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 != nil {
			if err := ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. beamManagementType-r16: enumerated
	{
		if ie.BeamManagementType_r16 != nil {
			if err := e.EncodeEnumerated(*ie.BeamManagementType_r16, cAParametersNRV1630BeamManagementTypeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. intraBandFreqSeparationUL-AggBW-GapBW-r16: enumerated
	{
		if ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16, cAParametersNRV1630IntraBandFreqSeparationULAggBWGapBWR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. interCA-NonAlignedFrame-B-r16: enumerated
	{
		if ie.InterCA_NonAlignedFrame_B_r16 != nil {
			if err := e.EncodeEnumerated(*ie.InterCA_NonAlignedFrame_B_r16, cAParametersNRV1630InterCANonAlignedFrameBR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1630Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. simulTX-SRS-AntSwitchingInterBandUL-CA-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16 = new(SimulSRS_ForAntennaSwitching_r16)
			if err := ie.SimulTX_SRS_AntSwitchingInterBandUL_CA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. beamManagementType-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1630BeamManagementTypeR16Constraints)
			if err != nil {
				return err
			}
			ie.BeamManagementType_r16 = &idx
		}
	}

	// 4. intraBandFreqSeparationUL-AggBW-GapBW-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1630IntraBandFreqSeparationULAggBWGapBWR16Constraints)
			if err != nil {
				return err
			}
			ie.IntraBandFreqSeparationUL_AggBW_GapBW_r16 = &idx
		}
	}

	// 5. interCA-NonAlignedFrame-B-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1630InterCANonAlignedFrameBR16Constraints)
			if err != nil {
				return err
			}
			ie.InterCA_NonAlignedFrame_B_r16 = &idx
		}
	}

	return nil
}
