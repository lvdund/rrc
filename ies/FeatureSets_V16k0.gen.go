// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSets-v16k0 (line 20077).

var featureSetsV16k0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "featureSetsDownlink-v16k0", Optional: true},
	},
}

var featureSetsV16k0FeatureSetsDownlinkV16k0Constraints = per.SizeRange(1, common.MaxDownlinkFeatureSets)

type FeatureSets_V16k0 struct {
	FeatureSetsDownlink_V16k0 []FeatureSetDownlink_V16k0
}

func (ie *FeatureSets_V16k0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetsV16k0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FeatureSetsDownlink_V16k0 != nil}); err != nil {
		return err
	}

	// 2. featureSetsDownlink-v16k0: sequence-of
	{
		if ie.FeatureSetsDownlink_V16k0 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetsV16k0FeatureSetsDownlinkV16k0Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.FeatureSetsDownlink_V16k0))); err != nil {
				return err
			}
			for i := range ie.FeatureSetsDownlink_V16k0 {
				if err := ie.FeatureSetsDownlink_V16k0[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSets_V16k0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetsV16k0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. featureSetsDownlink-v16k0: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(featureSetsV16k0FeatureSetsDownlinkV16k0Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureSetsDownlink_V16k0 = make([]FeatureSetDownlink_V16k0, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureSetsDownlink_V16k0[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
