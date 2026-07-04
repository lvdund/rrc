// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetCombinationId (line 19442).
// FeatureSetCombinationId ::=         INTEGER (0.. maxFeatureSetCombinations)

var featureSetCombinationIdConstraints = per.Constrained(0, common.MaxFeatureSetCombinations)

type FeatureSetCombinationId struct {
	Value int64
}

func (v *FeatureSetCombinationId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, featureSetCombinationIdConstraints)
}

func (v *FeatureSetCombinationId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(featureSetCombinationIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
