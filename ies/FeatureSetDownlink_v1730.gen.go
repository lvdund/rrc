// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v1730 (line 19644).

var featureSetDownlinkV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "prs-AsSpatialRelationRS-For-SRS-r17", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1730_Prs_AsSpatialRelationRS_For_SRS_r17_Supported = 0
)

var featureSetDownlinkV1730PrsAsSpatialRelationRSForSRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1730_Prs_AsSpatialRelationRS_For_SRS_r17_Supported},
}

type FeatureSetDownlink_v1730 struct {
	Prs_AsSpatialRelationRS_For_SRS_r17 *int64
}

func (ie *FeatureSetDownlink_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1730Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Prs_AsSpatialRelationRS_For_SRS_r17 != nil}); err != nil {
		return err
	}

	// 2. prs-AsSpatialRelationRS-For-SRS-r17: enumerated
	{
		if ie.Prs_AsSpatialRelationRS_For_SRS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Prs_AsSpatialRelationRS_For_SRS_r17, featureSetDownlinkV1730PrsAsSpatialRelationRSForSRSR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1730Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. prs-AsSpatialRelationRS-For-SRS-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1730PrsAsSpatialRelationRSForSRSR17Constraints)
			if err != nil {
				return err
			}
			ie.Prs_AsSpatialRelationRS_For_SRS_r17 = &idx
		}
	}

	return nil
}
