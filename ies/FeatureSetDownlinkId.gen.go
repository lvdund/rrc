// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetDownlinkId (line 19875).
// FeatureSetDownlinkId ::=            INTEGER (0..maxDownlinkFeatureSets)

var featureSetDownlinkIdConstraints = per.Constrained(0, common.MaxDownlinkFeatureSets)

type FeatureSetDownlinkId struct {
	Value int64
}

func (v *FeatureSetDownlinkId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featureSetDownlinkIdConstraints)
}

func (v *FeatureSetDownlinkId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featureSetDownlinkIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
