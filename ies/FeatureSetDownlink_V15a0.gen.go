// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v15a0 (line 19501).

var featureSetDownlinkV15a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedSRS-Resources", Optional: true},
	},
}

type FeatureSetDownlink_V15a0 struct {
	SupportedSRS_Resources *SRS_Resources
}

func (ie *FeatureSetDownlink_V15a0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV15a0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedSRS_Resources != nil}); err != nil {
		return err
	}

	// 2. supportedSRS-Resources: ref
	{
		if ie.SupportedSRS_Resources != nil {
			if err := ie.SupportedSRS_Resources.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_V15a0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV15a0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedSRS-Resources: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedSRS_Resources = new(SRS_Resources)
			if err := ie.SupportedSRS_Resources.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
