// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetDownlinkPerCC-Id (line 19982).
// FeatureSetDownlinkPerCC-Id ::=      INTEGER (1..maxPerCC-FeatureSets)

var featureSetDownlinkPerCCIdConstraints = per.Constrained(1, common.MaxPerCC_FeatureSets)

type FeatureSetDownlinkPerCC_Id struct {
	Value int64
}

func (v *FeatureSetDownlinkPerCC_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featureSetDownlinkPerCCIdConstraints)
}

func (v *FeatureSetDownlinkPerCC_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featureSetDownlinkPerCCIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
