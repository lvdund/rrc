// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC-v1780 (line 20571).

var featureSetUplinkPerCCV1780Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthUL-v1780", Optional: true},
	},
}

type FeatureSetUplinkPerCC_v1780 struct {
	SupportedBandwidthUL_v1780 *SupportedBandwidth_v1700
}

func (ie *FeatureSetUplinkPerCC_v1780) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1780Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthUL_v1780 != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthUL-v1780: ref
	{
		if ie.SupportedBandwidthUL_v1780 != nil {
			if err := ie.SupportedBandwidthUL_v1780.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC_v1780) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1780Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthUL-v1780: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandwidthUL_v1780 = new(SupportedBandwidth_v1700)
			if err := ie.SupportedBandwidthUL_v1780.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
