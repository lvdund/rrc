// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC-v1700 (line 20562).

var featureSetUplinkPerCCV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedMinBandwidthUL-r17", Optional: true},
		{Name: "mTRP-PUSCH-RepetitionTypeB-r17", Optional: true},
		{Name: "mTRP-PUSCH-TypeB-CB-r17", Optional: true},
		{Name: "supportedBandwidthUL-v1710", Optional: true},
	},
}

const (
	FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N1 = 0
	FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N2 = 1
	FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N3 = 2
	FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N4 = 3
)

var featureSetUplinkPerCCV1700MTRPPUSCHRepetitionTypeBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N1, FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N2, FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N3, FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_RepetitionTypeB_r17_N4},
}

const (
	FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_TypeB_CB_r17_N1 = 0
	FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_TypeB_CB_r17_N2 = 1
	FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_TypeB_CB_r17_N4 = 2
)

var featureSetUplinkPerCCV1700MTRPPUSCHTypeBCBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_TypeB_CB_r17_N1, FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_TypeB_CB_r17_N2, FeatureSetUplinkPerCC_v1700_MTRP_PUSCH_TypeB_CB_r17_N4},
}

type FeatureSetUplinkPerCC_v1700 struct {
	SupportedMinBandwidthUL_r17    *SupportedBandwidth_v1700
	MTRP_PUSCH_RepetitionTypeB_r17 *int64
	MTRP_PUSCH_TypeB_CB_r17        *int64
	SupportedBandwidthUL_v1710     *SupportedBandwidth_v1700
}

func (ie *FeatureSetUplinkPerCC_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedMinBandwidthUL_r17 != nil, ie.MTRP_PUSCH_RepetitionTypeB_r17 != nil, ie.MTRP_PUSCH_TypeB_CB_r17 != nil, ie.SupportedBandwidthUL_v1710 != nil}); err != nil {
		return err
	}

	// 2. supportedMinBandwidthUL-r17: ref
	{
		if ie.SupportedMinBandwidthUL_r17 != nil {
			if err := ie.SupportedMinBandwidthUL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mTRP-PUSCH-RepetitionTypeB-r17: enumerated
	{
		if ie.MTRP_PUSCH_RepetitionTypeB_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MTRP_PUSCH_RepetitionTypeB_r17, featureSetUplinkPerCCV1700MTRPPUSCHRepetitionTypeBR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. mTRP-PUSCH-TypeB-CB-r17: enumerated
	{
		if ie.MTRP_PUSCH_TypeB_CB_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MTRP_PUSCH_TypeB_CB_r17, featureSetUplinkPerCCV1700MTRPPUSCHTypeBCBR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. supportedBandwidthUL-v1710: ref
	{
		if ie.SupportedBandwidthUL_v1710 != nil {
			if err := ie.SupportedBandwidthUL_v1710.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedMinBandwidthUL-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedMinBandwidthUL_r17 = new(SupportedBandwidth_v1700)
			if err := ie.SupportedMinBandwidthUL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mTRP-PUSCH-RepetitionTypeB-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCV1700MTRPPUSCHRepetitionTypeBR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_RepetitionTypeB_r17 = &idx
		}
	}

	// 4. mTRP-PUSCH-TypeB-CB-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCV1700MTRPPUSCHTypeBCBR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_TypeB_CB_r17 = &idx
		}
	}

	// 5. supportedBandwidthUL-v1710: ref
	{
		if seq.IsComponentPresent(3) {
			ie.SupportedBandwidthUL_v1710 = new(SupportedBandwidth_v1700)
			if err := ie.SupportedBandwidthUL_v1710.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
