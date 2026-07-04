// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSets-v17d0 (line 20081).

var featureSetsV17d0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSetsDownlink-v17d0", Optional: true},
	},
}

var featureSetsV17d0FeatureSetsDownlinkV17d0Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

type FeatureSets_V17d0 struct {
	FeatureSetsDownlink_V17d0 []FeatureSetDownlink_V17d0
}

func (ie *FeatureSets_V17d0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetsV17d0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSetsDownlink_V17d0 != nil}); err != nil {
		return err
	}

	// 2. featureSetsDownlink-v17d0: sequence-of
	{
		if ie.FeatureSetsDownlink_V17d0 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetsV17d0FeatureSetsDownlinkV17d0Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_V17d0))); err != nil {
				return err
			}
			for i := range ie.FeatureSetsDownlink_V17d0 {
				if err := ie.FeatureSetsDownlink_V17d0[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSets_V17d0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetsV17d0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSetsDownlink-v17d0: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(featureSetsV17d0FeatureSetsDownlinkV17d0Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_V17d0 = make([]FeatureSetDownlink_V17d0, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_V17d0[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
