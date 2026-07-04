// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GapOccasionRatioPerGapConfig-r19 (line 2817).

var gapOccasionRatioPerGapConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measGapId-r19"},
		{Name: "gapOccasionRatio-r19"},
	},
}

type GapOccasionRatioPerGapConfig_r19 struct {
	MeasGapId_r19        MeasGapId_r17
	GapOccasionRatio_r19 GapOccasionRatio_r19
}

func (ie *GapOccasionRatioPerGapConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gapOccasionRatioPerGapConfigR19Constraints)
	_ = seq

	// 1. measGapId-r19: ref
	{
		if err := ie.MeasGapId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. gapOccasionRatio-r19: ref
	{
		if err := ie.GapOccasionRatio_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *GapOccasionRatioPerGapConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gapOccasionRatioPerGapConfigR19Constraints)
	_ = seq

	// 1. measGapId-r19: ref
	{
		if err := ie.MeasGapId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. gapOccasionRatio-r19: ref
	{
		if err := ie.GapOccasionRatio_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
