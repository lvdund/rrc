// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v15t0 (line 19505).

var featureSetDownlinkV15t0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "zeroSlotOffsetAperiodicSRS", Optional: true},
	},
}

const (
	FeatureSetDownlink_V15t0_ZeroSlotOffsetAperiodicSRS_Supported = 0
)

var featureSetDownlinkV15t0ZeroSlotOffsetAperiodicSRSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V15t0_ZeroSlotOffsetAperiodicSRS_Supported},
}

type FeatureSetDownlink_V15t0 struct {
	ZeroSlotOffsetAperiodicSRS *int64
}

func (ie *FeatureSetDownlink_V15t0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV15t0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ZeroSlotOffsetAperiodicSRS != nil}); err != nil {
		return err
	}

	// 2. zeroSlotOffsetAperiodicSRS: enumerated
	{
		if ie.ZeroSlotOffsetAperiodicSRS != nil {
			if err := e.EncodeEnumerated(*ie.ZeroSlotOffsetAperiodicSRS, featureSetDownlinkV15t0ZeroSlotOffsetAperiodicSRSConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_V15t0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV15t0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. zeroSlotOffsetAperiodicSRS: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV15t0ZeroSlotOffsetAperiodicSRSConstraints)
			if err != nil {
				return err
			}
			ie.ZeroSlotOffsetAperiodicSRS = &idx
		}
	}

	return nil
}
