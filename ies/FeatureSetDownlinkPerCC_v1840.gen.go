// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1840 (line 19947).

var featureSetDownlinkPerCCV1840Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthDL-v1840", Optional: true},
		{Name: "supportedMinBandwidthDL-v1840", Optional: true},
	},
}

type FeatureSetDownlinkPerCC_v1840 struct {
	SupportedBandwidthDL_v1840    *SupportedBandwidth_v1840
	SupportedMinBandwidthDL_v1840 *SupportedBandwidth_v1840
}

func (ie *FeatureSetDownlinkPerCC_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1840Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthDL_v1840 != nil, ie.SupportedMinBandwidthDL_v1840 != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthDL-v1840: ref
	{
		if ie.SupportedBandwidthDL_v1840 != nil {
			if err := ie.SupportedBandwidthDL_v1840.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthDL-v1840: ref
	{
		if ie.SupportedMinBandwidthDL_v1840 != nil {
			if err := ie.SupportedMinBandwidthDL_v1840.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1840Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthDL-v1840: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandwidthDL_v1840 = new(SupportedBandwidth_v1840)
			if err := ie.SupportedBandwidthDL_v1840.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthDL-v1840: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedMinBandwidthDL_v1840 = new(SupportedBandwidth_v1840)
			if err := ie.SupportedMinBandwidthDL_v1840.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
