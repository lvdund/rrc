// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetUplinkPerCC-Id (line 20720).
// FeatureSetUplinkPerCC-Id ::=            INTEGER (1..maxPerCC-FeatureSets)

var featureSetUplinkPerCCIdConstraints = per.Constrained(1, common.MaxPerCC_FeatureSets)

type FeatureSetUplinkPerCC_Id struct {
	Value int64
}

func (v *FeatureSetUplinkPerCC_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featureSetUplinkPerCCIdConstraints)
}

func (v *FeatureSetUplinkPerCC_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featureSetUplinkPerCCIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
