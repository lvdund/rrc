// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GapOccasionRatio-r19 (line 8467).
// GapOccasionRatio-r19 ::= ENUMERATED {pc0, pc20, pc40, pc60}

const (
	GapOccasionRatio_r19_Pc0  = 0
	GapOccasionRatio_r19_Pc20 = 1
	GapOccasionRatio_r19_Pc40 = 2
	GapOccasionRatio_r19_Pc60 = 3
)

var gapOccasionRatioR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapOccasionRatio_r19_Pc0, GapOccasionRatio_r19_Pc20, GapOccasionRatio_r19_Pc40, GapOccasionRatio_r19_Pc60},
}

type GapOccasionRatio_r19 struct {
	Value int64
}

func (v *GapOccasionRatio_r19) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, gapOccasionRatioR19Constraints)
}

func (v *GapOccasionRatio_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(gapOccasionRatioR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
