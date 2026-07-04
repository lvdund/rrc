// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v1860 (line 19760).

var featureSetDownlinkV1860Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-AntennaSwitching8T8R2SP-1Periodic-r18", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1860_Srs_AntennaSwitching8T8R2SP_1Periodic_r18_Supported = 0
)

var featureSetDownlinkV1860SrsAntennaSwitching8T8R2SP1PeriodicR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1860_Srs_AntennaSwitching8T8R2SP_1Periodic_r18_Supported},
}

type FeatureSetDownlink_v1860 struct {
	Srs_AntennaSwitching8T8R2SP_1Periodic_r18 *int64
}

func (ie *FeatureSetDownlink_v1860) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1860Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18 != nil}); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching8T8R2SP-1Periodic-r18: enumerated
	{
		if ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18, featureSetDownlinkV1860SrsAntennaSwitching8T8R2SP1PeriodicR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1860) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1860Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching8T8R2SP-1Periodic-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1860SrsAntennaSwitching8T8R2SP1PeriodicR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18 = &idx
		}
	}

	return nil
}
