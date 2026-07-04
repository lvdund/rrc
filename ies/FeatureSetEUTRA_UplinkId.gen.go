// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetEUTRA-UplinkId (line 19992).
// FeatureSetEUTRA-UplinkId ::=                    INTEGER (0..maxEUTRA-UL-FeatureSets)

var featureSetEUTRAUplinkIdConstraints = per.Constrained(0, common.MaxEUTRA_UL_FeatureSets)

type FeatureSetEUTRA_UplinkId struct {
	Value int64
}

func (v *FeatureSetEUTRA_UplinkId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featureSetEUTRAUplinkIdConstraints)
}

func (v *FeatureSetEUTRA_UplinkId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featureSetEUTRAUplinkIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
