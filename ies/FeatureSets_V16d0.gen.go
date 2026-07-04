// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSets-v16d0 (line 20073).

var featureSetsV16d0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSetsUplink-v16d0", Optional: true},
	},
}

var featureSetsV16d0FeatureSetsUplinkV16d0Constraints = per.SizeRange(1, common.MaxUplinkFeatureSets)

type FeatureSets_V16d0 struct {
	FeatureSetsUplink_V16d0 []FeatureSetUplink_V16d0
}

func (ie *FeatureSets_V16d0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetsV16d0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSetsUplink_V16d0 != nil}); err != nil {
		return err
	}

	// 2. featureSetsUplink-v16d0: sequence-of
	{
		if ie.FeatureSetsUplink_V16d0 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetsV16d0FeatureSetsUplinkV16d0Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsUplink_V16d0))); err != nil {
				return err
			}
			for i := range ie.FeatureSetsUplink_V16d0 {
				if err := ie.FeatureSetsUplink_V16d0[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSets_V16d0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetsV16d0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSetsUplink-v16d0: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(featureSetsV16d0FeatureSetsUplinkV16d0Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsUplink_V16d0 = make([]FeatureSetUplink_V16d0, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsUplink_V16d0[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
