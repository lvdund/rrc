// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1780 (line 19927).

var featureSetDownlinkPerCCV1780Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthDL-v1780", Optional: true},
	},
}

type FeatureSetDownlinkPerCC_v1780 struct {
	SupportedBandwidthDL_v1780 *SupportedBandwidth_v1700
}

func (ie *FeatureSetDownlinkPerCC_v1780) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1780Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthDL_v1780 != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthDL-v1780: ref
	{
		if ie.SupportedBandwidthDL_v1780 != nil {
			if err := ie.SupportedBandwidthDL_v1780.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1780) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1780Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthDL-v1780: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandwidthDL_v1780 = new(SupportedBandwidth_v1700)
			if err := ie.SupportedBandwidthDL_v1780.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
