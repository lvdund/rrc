// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetUplinkId (line 20539).
// FeatureSetUplinkId ::=                  INTEGER (0..maxUplinkFeatureSets)

var featureSetUplinkIdConstraints = per.Constrained(0, common.MaxUplinkFeatureSets)

type FeatureSetUplinkId struct {
	Value int64
}

func (v *FeatureSetUplinkId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featureSetUplinkIdConstraints)
}

func (v *FeatureSetUplinkId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featureSetUplinkIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
