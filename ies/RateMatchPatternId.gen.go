// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RateMatchPatternId (line 13293).
// RateMatchPatternId ::=              INTEGER (0..maxNrofRateMatchPatterns-1)

var rateMatchPatternIdConstraints = per.Constrained(0, common.MaxNrofRateMatchPatterns_1)

type RateMatchPatternId struct {
	Value int64
}

func (v *RateMatchPatternId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rateMatchPatternIdConstraints)
}

func (v *RateMatchPatternId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rateMatchPatternIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
