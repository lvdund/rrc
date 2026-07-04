// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetEUTRA-DownlinkId (line 19987).
// FeatureSetEUTRA-DownlinkId ::=      INTEGER (0..maxEUTRA-DL-FeatureSets)

var featureSetEUTRADownlinkIdConstraints = per.Constrained(0, common.MaxEUTRA_DL_FeatureSets)

type FeatureSetEUTRA_DownlinkId struct {
	Value int64
}

func (v *FeatureSetEUTRA_DownlinkId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featureSetEUTRADownlinkIdConstraints)
}

func (v *FeatureSetEUTRA_DownlinkId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featureSetEUTRADownlinkIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
