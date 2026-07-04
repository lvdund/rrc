// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1780 (line 17612).

var cAParametersNRV1780Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "parallelTxPUCCH-PUSCH-SamePriority-r17", Optional: true},
		{Name: "supportedAggBW-FR1-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1780_ParallelTxPUCCH_PUSCH_SamePriority_r17_Supported = 0
)

var cAParametersNRV1780ParallelTxPUCCHPUSCHSamePriorityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1780_ParallelTxPUCCH_PUSCH_SamePriority_r17_Supported},
}

var cAParametersNRV1780SupportedAggBWFR1R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scalingFactorSCS-r17", Optional: true},
		{Name: "supportedAggBW-FDD-DL-r17", Optional: true},
		{Name: "supportedAggBW-FDD-UL-r17", Optional: true},
		{Name: "supportedAggBW-TDD-DL-r17", Optional: true},
		{Name: "supportedAggBW-TDD-UL-r17", Optional: true},
		{Name: "supportedAggBW-TotalDL-r17", Optional: true},
		{Name: "supportedAggBW-TotalUL-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1780_SupportedAggBW_FR1_r17_ScalingFactorSCS_r17_True = 0
)

var cAParametersNRV1780SupportedAggBWFR1R17ScalingFactorSCSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1780_SupportedAggBW_FR1_r17_ScalingFactorSCS_r17_True},
}

type CA_ParametersNR_v1780 struct {
	ParallelTxPUCCH_PUSCH_SamePriority_r17 *int64
	SupportedAggBW_FR1_r17                 *struct {
		ScalingFactorSCS_r17       *int64
		SupportedAggBW_FDD_DL_r17  *SupportedAggBandwidth_r17
		SupportedAggBW_FDD_UL_r17  *SupportedAggBandwidth_r17
		SupportedAggBW_TDD_DL_r17  *SupportedAggBandwidth_r17
		SupportedAggBW_TDD_UL_r17  *SupportedAggBandwidth_r17
		SupportedAggBW_TotalDL_r17 *SupportedAggBandwidth_r17
		SupportedAggBW_TotalUL_r17 *SupportedAggBandwidth_r17
	}
}

func (ie *CA_ParametersNR_v1780) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1780Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ParallelTxPUCCH_PUSCH_SamePriority_r17 != nil, ie.SupportedAggBW_FR1_r17 != nil}); err != nil {
		return err
	}

	// 2. parallelTxPUCCH-PUSCH-SamePriority-r17: enumerated
	{
		if ie.ParallelTxPUCCH_PUSCH_SamePriority_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxPUCCH_PUSCH_SamePriority_r17, cAParametersNRV1780ParallelTxPUCCHPUSCHSamePriorityR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. supportedAggBW-FR1-r17: sequence
	{
		if ie.SupportedAggBW_FR1_r17 != nil {
			c := ie.SupportedAggBW_FR1_r17
			cAParametersNRV1780SupportedAggBWFR1R17Seq := e.NewSequenceEncoder(cAParametersNRV1780SupportedAggBWFR1R17Constraints)
			if err := cAParametersNRV1780SupportedAggBWFR1R17Seq.EncodePreamble([]bool{c.ScalingFactorSCS_r17 != nil, c.SupportedAggBW_FDD_DL_r17 != nil, c.SupportedAggBW_FDD_UL_r17 != nil, c.SupportedAggBW_TDD_DL_r17 != nil, c.SupportedAggBW_TDD_UL_r17 != nil, c.SupportedAggBW_TotalDL_r17 != nil, c.SupportedAggBW_TotalUL_r17 != nil}); err != nil {
				return err
			}
			if c.ScalingFactorSCS_r17 != nil {
				if err := e.EncodeEnumerated((*c.ScalingFactorSCS_r17), cAParametersNRV1780SupportedAggBWFR1R17ScalingFactorSCSR17Constraints); err != nil {
					return err
				}
			}
			if c.SupportedAggBW_FDD_DL_r17 != nil {
				if err := c.SupportedAggBW_FDD_DL_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SupportedAggBW_FDD_UL_r17 != nil {
				if err := c.SupportedAggBW_FDD_UL_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SupportedAggBW_TDD_DL_r17 != nil {
				if err := c.SupportedAggBW_TDD_DL_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SupportedAggBW_TDD_UL_r17 != nil {
				if err := c.SupportedAggBW_TDD_UL_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SupportedAggBW_TotalDL_r17 != nil {
				if err := c.SupportedAggBW_TotalDL_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SupportedAggBW_TotalUL_r17 != nil {
				if err := c.SupportedAggBW_TotalUL_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1780) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1780Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. parallelTxPUCCH-PUSCH-SamePriority-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1780ParallelTxPUCCHPUSCHSamePriorityR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelTxPUCCH_PUSCH_SamePriority_r17 = &idx
		}
	}

	// 3. supportedAggBW-FR1-r17: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedAggBW_FR1_r17 = &struct {
				ScalingFactorSCS_r17       *int64
				SupportedAggBW_FDD_DL_r17  *SupportedAggBandwidth_r17
				SupportedAggBW_FDD_UL_r17  *SupportedAggBandwidth_r17
				SupportedAggBW_TDD_DL_r17  *SupportedAggBandwidth_r17
				SupportedAggBW_TDD_UL_r17  *SupportedAggBandwidth_r17
				SupportedAggBW_TotalDL_r17 *SupportedAggBandwidth_r17
				SupportedAggBW_TotalUL_r17 *SupportedAggBandwidth_r17
			}{}
			c := ie.SupportedAggBW_FR1_r17
			cAParametersNRV1780SupportedAggBWFR1R17Seq := d.NewSequenceDecoder(cAParametersNRV1780SupportedAggBWFR1R17Constraints)
			if err := cAParametersNRV1780SupportedAggBWFR1R17Seq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1780SupportedAggBWFR1R17Seq.IsComponentPresent(0) {
				c.ScalingFactorSCS_r17 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1780SupportedAggBWFR1R17ScalingFactorSCSR17Constraints)
				if err != nil {
					return err
				}
				(*c.ScalingFactorSCS_r17) = v
			}
			if cAParametersNRV1780SupportedAggBWFR1R17Seq.IsComponentPresent(1) {
				c.SupportedAggBW_FDD_DL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_FDD_DL_r17).Decode(d); err != nil {
					return err
				}
			}
			if cAParametersNRV1780SupportedAggBWFR1R17Seq.IsComponentPresent(2) {
				c.SupportedAggBW_FDD_UL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_FDD_UL_r17).Decode(d); err != nil {
					return err
				}
			}
			if cAParametersNRV1780SupportedAggBWFR1R17Seq.IsComponentPresent(3) {
				c.SupportedAggBW_TDD_DL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_TDD_DL_r17).Decode(d); err != nil {
					return err
				}
			}
			if cAParametersNRV1780SupportedAggBWFR1R17Seq.IsComponentPresent(4) {
				c.SupportedAggBW_TDD_UL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_TDD_UL_r17).Decode(d); err != nil {
					return err
				}
			}
			if cAParametersNRV1780SupportedAggBWFR1R17Seq.IsComponentPresent(5) {
				c.SupportedAggBW_TotalDL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_TotalDL_r17).Decode(d); err != nil {
					return err
				}
			}
			if cAParametersNRV1780SupportedAggBWFR1R17Seq.IsComponentPresent(6) {
				c.SupportedAggBW_TotalUL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_TotalUL_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
