// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC-v1840 (line 20673).

var featureSetUplinkPerCCV1840Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthUL-v1840", Optional: true},
		{Name: "supportedMinBandwidthUL-v1840", Optional: true},
	},
}

type FeatureSetUplinkPerCC_v1840 struct {
	SupportedBandwidthUL_v1840    *SupportedBandwidth_v1840
	SupportedMinBandwidthUL_v1840 *SupportedBandwidth_v1840
}

func (ie *FeatureSetUplinkPerCC_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1840Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthUL_v1840 != nil, ie.SupportedMinBandwidthUL_v1840 != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthUL-v1840: ref
	{
		if ie.SupportedBandwidthUL_v1840 != nil {
			if err := ie.SupportedBandwidthUL_v1840.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthUL-v1840: ref
	{
		if ie.SupportedMinBandwidthUL_v1840 != nil {
			if err := ie.SupportedMinBandwidthUL_v1840.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1840Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthUL-v1840: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandwidthUL_v1840 = new(SupportedBandwidth_v1840)
			if err := ie.SupportedBandwidthUL_v1840.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthUL-v1840: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedMinBandwidthUL_v1840 = new(SupportedBandwidth_v1840)
			if err := ie.SupportedMinBandwidthUL_v1840.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
