// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v16d0 (line 20256).

var featureSetUplinkV16d0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-RepetitionTypeB-v16d0", Optional: true},
	},
}

const (
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N2  = 0
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N3  = 1
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N4  = 2
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N7  = 3
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N8  = 4
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N12 = 5
)

var featureSetUplinkV16d0PuschRepetitionTypeBV16d0MaxNumberPUSCHTxCap1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N2, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N3, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N4, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N7, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N8, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap1_r16_N12},
}

const (
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N2  = 0
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N3  = 1
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N4  = 2
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N7  = 3
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N8  = 4
	FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N12 = 5
)

var featureSetUplinkV16d0PuschRepetitionTypeBV16d0MaxNumberPUSCHTxCap2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N2, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N3, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N4, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N7, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N8, FeatureSetUplink_V16d0_Pusch_RepetitionTypeB_V16d0_MaxNumberPUSCH_Tx_Cap2_r16_N12},
}

type FeatureSetUplink_V16d0 struct {
	Pusch_RepetitionTypeB_V16d0 *struct {
		MaxNumberPUSCH_Tx_Cap1_r16 int64
		MaxNumberPUSCH_Tx_Cap2_r16 int64
	}
}

func (ie *FeatureSetUplink_V16d0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV16d0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pusch_RepetitionTypeB_V16d0 != nil}); err != nil {
		return err
	}

	// 2. pusch-RepetitionTypeB-v16d0: sequence
	{
		if ie.Pusch_RepetitionTypeB_V16d0 != nil {
			c := ie.Pusch_RepetitionTypeB_V16d0
			if err := e.EncodeEnumerated(c.MaxNumberPUSCH_Tx_Cap1_r16, featureSetUplinkV16d0PuschRepetitionTypeBV16d0MaxNumberPUSCHTxCap1R16Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberPUSCH_Tx_Cap2_r16, featureSetUplinkV16d0PuschRepetitionTypeBV16d0MaxNumberPUSCHTxCap2R16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_V16d0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV16d0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pusch-RepetitionTypeB-v16d0: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Pusch_RepetitionTypeB_V16d0 = &struct {
				MaxNumberPUSCH_Tx_Cap1_r16 int64
				MaxNumberPUSCH_Tx_Cap2_r16 int64
			}{}
			c := ie.Pusch_RepetitionTypeB_V16d0
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV16d0PuschRepetitionTypeBV16d0MaxNumberPUSCHTxCap1R16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberPUSCH_Tx_Cap1_r16 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV16d0PuschRepetitionTypeBV16d0MaxNumberPUSCHTxCap2R16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberPUSCH_Tx_Cap2_r16 = v
			}
		}
	}

	return nil
}
