// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersEUTRA (line 17249).

var cAParametersEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "multipleTimingAdvance", Optional: true},
		{Name: "simultaneousRx-Tx", Optional: true},
		{Name: "supportedNAICS-2CRS-AP", Optional: true},
		{Name: "additionalRx-Tx-PerformanceReq", Optional: true},
		{Name: "ue-CA-PowerClass-N", Optional: true},
		{Name: "supportedBandwidthCombinationSetEUTRA-v1530", Optional: true},
	},
}

const (
	CA_ParametersEUTRA_MultipleTimingAdvance_Supported = 0
)

var cAParametersEUTRAMultipleTimingAdvanceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersEUTRA_MultipleTimingAdvance_Supported},
}

const (
	CA_ParametersEUTRA_SimultaneousRx_Tx_Supported = 0
)

var cAParametersEUTRASimultaneousRxTxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersEUTRA_SimultaneousRx_Tx_Supported},
}

var cAParametersEUTRASupportedNAICS2CRSAPConstraints = per.SizeRange(1, 8)

const (
	CA_ParametersEUTRA_AdditionalRx_Tx_PerformanceReq_Supported = 0
)

var cAParametersEUTRAAdditionalRxTxPerformanceReqConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersEUTRA_AdditionalRx_Tx_PerformanceReq_Supported},
}

const (
	CA_ParametersEUTRA_Ue_CA_PowerClass_N_Class2 = 0
)

var cAParametersEUTRAUeCAPowerClassNConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersEUTRA_Ue_CA_PowerClass_N_Class2},
}

var cAParametersEUTRASupportedBandwidthCombinationSetEUTRAV1530Constraints = per.SizeRange(1, 32)

type CA_ParametersEUTRA struct {
	MultipleTimingAdvance                       *int64
	SimultaneousRx_Tx                           *int64
	SupportedNAICS_2CRS_AP                      *per.BitString
	AdditionalRx_Tx_PerformanceReq              *int64
	Ue_CA_PowerClass_N                          *int64
	SupportedBandwidthCombinationSetEUTRA_v1530 *per.BitString
}

func (ie *CA_ParametersEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersEUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MultipleTimingAdvance != nil, ie.SimultaneousRx_Tx != nil, ie.SupportedNAICS_2CRS_AP != nil, ie.AdditionalRx_Tx_PerformanceReq != nil, ie.Ue_CA_PowerClass_N != nil, ie.SupportedBandwidthCombinationSetEUTRA_v1530 != nil}); err != nil {
		return err
	}

	// 3. multipleTimingAdvance: enumerated
	{
		if ie.MultipleTimingAdvance != nil {
			if err := e.EncodeEnumerated(*ie.MultipleTimingAdvance, cAParametersEUTRAMultipleTimingAdvanceConstraints); err != nil {
				return err
			}
		}
	}

	// 4. simultaneousRx-Tx: enumerated
	{
		if ie.SimultaneousRx_Tx != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousRx_Tx, cAParametersEUTRASimultaneousRxTxConstraints); err != nil {
				return err
			}
		}
	}

	// 5. supportedNAICS-2CRS-AP: bit-string
	{
		if ie.SupportedNAICS_2CRS_AP != nil {
			if err := e.EncodeBitString(*ie.SupportedNAICS_2CRS_AP, cAParametersEUTRASupportedNAICS2CRSAPConstraints); err != nil {
				return err
			}
		}
	}

	// 6. additionalRx-Tx-PerformanceReq: enumerated
	{
		if ie.AdditionalRx_Tx_PerformanceReq != nil {
			if err := e.EncodeEnumerated(*ie.AdditionalRx_Tx_PerformanceReq, cAParametersEUTRAAdditionalRxTxPerformanceReqConstraints); err != nil {
				return err
			}
		}
	}

	// 7. ue-CA-PowerClass-N: enumerated
	{
		if ie.Ue_CA_PowerClass_N != nil {
			if err := e.EncodeEnumerated(*ie.Ue_CA_PowerClass_N, cAParametersEUTRAUeCAPowerClassNConstraints); err != nil {
				return err
			}
		}
	}

	// 8. supportedBandwidthCombinationSetEUTRA-v1530: bit-string
	{
		if ie.SupportedBandwidthCombinationSetEUTRA_v1530 != nil {
			if err := e.EncodeBitString(*ie.SupportedBandwidthCombinationSetEUTRA_v1530, cAParametersEUTRASupportedBandwidthCombinationSetEUTRAV1530Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersEUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. multipleTimingAdvance: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersEUTRAMultipleTimingAdvanceConstraints)
			if err != nil {
				return err
			}
			ie.MultipleTimingAdvance = &idx
		}
	}

	// 4. simultaneousRx-Tx: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersEUTRASimultaneousRxTxConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRx_Tx = &idx
		}
	}

	// 5. supportedNAICS-2CRS-AP: bit-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeBitString(cAParametersEUTRASupportedNAICS2CRSAPConstraints)
			if err != nil {
				return err
			}
			ie.SupportedNAICS_2CRS_AP = &v
		}
	}

	// 6. additionalRx-Tx-PerformanceReq: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersEUTRAAdditionalRxTxPerformanceReqConstraints)
			if err != nil {
				return err
			}
			ie.AdditionalRx_Tx_PerformanceReq = &idx
		}
	}

	// 7. ue-CA-PowerClass-N: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersEUTRAUeCAPowerClassNConstraints)
			if err != nil {
				return err
			}
			ie.Ue_CA_PowerClass_N = &idx
		}
	}

	// 8. supportedBandwidthCombinationSetEUTRA-v1530: bit-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeBitString(cAParametersEUTRASupportedBandwidthCombinationSetEUTRAV1530Constraints)
			if err != nil {
				return err
			}
			ie.SupportedBandwidthCombinationSetEUTRA_v1530 = &v
		}
	}

	return nil
}
