// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CI-ConfigurationPerServingCell-r16 (line 16281).

var cIConfigurationPerServingCellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId"},
		{Name: "positionInDCI-r16"},
		{Name: "positionInDCI-ForSUL-r16", Optional: true},
		{Name: "ci-PayloadSize-r16"},
		{Name: "timeFrequencyRegion-r16"},
		{Name: "uplinkCancellationPriority-v1610", Optional: true},
	},
}

var cIConfigurationPerServingCellR16PositionInDCIR16Constraints = per.Constrained(0, common.MaxCI_DCI_PayloadSize_1_r16)

var cIConfigurationPerServingCellR16PositionInDCIForSULR16Constraints = per.Constrained(0, common.MaxCI_DCI_PayloadSize_1_r16)

const (
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N1   = 0
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N2   = 1
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N4   = 2
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N5   = 3
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N7   = 4
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N8   = 5
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N10  = 6
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N14  = 7
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N16  = 8
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N20  = 9
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N28  = 10
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N32  = 11
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N35  = 12
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N42  = 13
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N56  = 14
	CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N112 = 15
)

var cIConfigurationPerServingCellR16CiPayloadSizeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N1, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N2, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N4, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N5, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N7, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N8, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N10, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N14, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N16, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N20, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N28, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N32, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N35, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N42, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N56, CI_ConfigurationPerServingCell_r16_Ci_PayloadSize_r16_N112},
}

const (
	CI_ConfigurationPerServingCell_r16_UplinkCancellationPriority_v1610_Enabled = 0
)

var cIConfigurationPerServingCellR16UplinkCancellationPriorityV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CI_ConfigurationPerServingCell_r16_UplinkCancellationPriority_v1610_Enabled},
}

var cIConfigurationPerServingCellR16TimeFrequencyRegionR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "timeDurationForCI-r16", Optional: true},
		{Name: "timeGranularityForCI-r16"},
		{Name: "frequencyRegionForCI-r16"},
		{Name: "deltaOffset-r16"},
	},
}

const (
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N2  = 0
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N4  = 1
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N7  = 2
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N14 = 3
)

var cIConfigurationPerServingCellR16TimeFrequencyRegionR16TimeDurationForCIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N2, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N4, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N7, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeDurationForCI_r16_N14},
}

const (
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N1  = 0
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N2  = 1
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N4  = 2
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N7  = 3
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N14 = 4
	CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N28 = 5
)

var cIConfigurationPerServingCellR16TimeFrequencyRegionR16TimeGranularityForCIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N1, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N2, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N4, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N7, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N14, CI_ConfigurationPerServingCell_r16_TimeFrequencyRegion_r16_TimeGranularityForCI_r16_N28},
}

type CI_ConfigurationPerServingCell_r16 struct {
	ServingCellId            ServCellIndex
	PositionInDCI_r16        int64
	PositionInDCI_ForSUL_r16 *int64
	Ci_PayloadSize_r16       int64
	TimeFrequencyRegion_r16  struct {
		TimeDurationForCI_r16    *int64
		TimeGranularityForCI_r16 int64
		FrequencyRegionForCI_r16 int64
		DeltaOffset_r16          int64
	}
	UplinkCancellationPriority_v1610 *int64
}

func (ie *CI_ConfigurationPerServingCell_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cIConfigurationPerServingCellR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PositionInDCI_ForSUL_r16 != nil, ie.UplinkCancellationPriority_v1610 != nil}); err != nil {
		return err
	}

	// 2. servingCellId: ref
	{
		if err := ie.ServingCellId.Encode(e); err != nil {
			return err
		}
	}

	// 3. positionInDCI-r16: integer
	{
		if err := e.EncodeInteger(ie.PositionInDCI_r16, cIConfigurationPerServingCellR16PositionInDCIR16Constraints); err != nil {
			return err
		}
	}

	// 4. positionInDCI-ForSUL-r16: integer
	{
		if ie.PositionInDCI_ForSUL_r16 != nil {
			if err := e.EncodeInteger(*ie.PositionInDCI_ForSUL_r16, cIConfigurationPerServingCellR16PositionInDCIForSULR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ci-PayloadSize-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ci_PayloadSize_r16, cIConfigurationPerServingCellR16CiPayloadSizeR16Constraints); err != nil {
			return err
		}
	}

	// 6. timeFrequencyRegion-r16: sequence
	{
		{
			c := &ie.TimeFrequencyRegion_r16
			cIConfigurationPerServingCellR16TimeFrequencyRegionR16Seq := e.NewSequenceEncoder(cIConfigurationPerServingCellR16TimeFrequencyRegionR16Constraints)
			if err := cIConfigurationPerServingCellR16TimeFrequencyRegionR16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := cIConfigurationPerServingCellR16TimeFrequencyRegionR16Seq.EncodePreamble([]bool{c.TimeDurationForCI_r16 != nil}); err != nil {
				return err
			}
			if c.TimeDurationForCI_r16 != nil {
				if err := e.EncodeEnumerated((*c.TimeDurationForCI_r16), cIConfigurationPerServingCellR16TimeFrequencyRegionR16TimeDurationForCIR16Constraints); err != nil {
					return err
				}
			}
			if err := e.EncodeEnumerated(c.TimeGranularityForCI_r16, cIConfigurationPerServingCellR16TimeFrequencyRegionR16TimeGranularityForCIR16Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.FrequencyRegionForCI_r16, per.Constrained(0, 37949)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.DeltaOffset_r16, per.Constrained(0, 2)); err != nil {
				return err
			}
		}
	}

	// 7. uplinkCancellationPriority-v1610: enumerated
	{
		if ie.UplinkCancellationPriority_v1610 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkCancellationPriority_v1610, cIConfigurationPerServingCellR16UplinkCancellationPriorityV1610Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cIConfigurationPerServingCellR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. servingCellId: ref
	{
		if err := ie.ServingCellId.Decode(d); err != nil {
			return err
		}
	}

	// 3. positionInDCI-r16: integer
	{
		v1, err := d.DecodeInteger(cIConfigurationPerServingCellR16PositionInDCIR16Constraints)
		if err != nil {
			return err
		}
		ie.PositionInDCI_r16 = v1
	}

	// 4. positionInDCI-ForSUL-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(cIConfigurationPerServingCellR16PositionInDCIForSULR16Constraints)
			if err != nil {
				return err
			}
			ie.PositionInDCI_ForSUL_r16 = &v
		}
	}

	// 5. ci-PayloadSize-r16: enumerated
	{
		v3, err := d.DecodeEnumerated(cIConfigurationPerServingCellR16CiPayloadSizeR16Constraints)
		if err != nil {
			return err
		}
		ie.Ci_PayloadSize_r16 = v3
	}

	// 6. timeFrequencyRegion-r16: sequence
	{
		{
			c := &ie.TimeFrequencyRegion_r16
			cIConfigurationPerServingCellR16TimeFrequencyRegionR16Seq := d.NewSequenceDecoder(cIConfigurationPerServingCellR16TimeFrequencyRegionR16Constraints)
			if err := cIConfigurationPerServingCellR16TimeFrequencyRegionR16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := cIConfigurationPerServingCellR16TimeFrequencyRegionR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if cIConfigurationPerServingCellR16TimeFrequencyRegionR16Seq.IsComponentPresent(0) {
				c.TimeDurationForCI_r16 = new(int64)
				v, err := d.DecodeEnumerated(cIConfigurationPerServingCellR16TimeFrequencyRegionR16TimeDurationForCIR16Constraints)
				if err != nil {
					return err
				}
				(*c.TimeDurationForCI_r16) = v
			}
			{
				v, err := d.DecodeEnumerated(cIConfigurationPerServingCellR16TimeFrequencyRegionR16TimeGranularityForCIR16Constraints)
				if err != nil {
					return err
				}
				c.TimeGranularityForCI_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 37949))
				if err != nil {
					return err
				}
				c.FrequencyRegionForCI_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 2))
				if err != nil {
					return err
				}
				c.DeltaOffset_r16 = v
			}
		}
	}

	// 7. uplinkCancellationPriority-v1610: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cIConfigurationPerServingCellR16UplinkCancellationPriorityV1610Constraints)
			if err != nil {
				return err
			}
			ie.UplinkCancellationPriority_v1610 = &idx
		}
	}

	return nil
}
