// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GoodServingCellEvaluation-r17 (line 5764).

var goodServingCellEvaluationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "offset-r17", Optional: true},
	},
}

const (
	GoodServingCellEvaluation_r17_Offset_r17_Db2 = 0
	GoodServingCellEvaluation_r17_Offset_r17_Db4 = 1
	GoodServingCellEvaluation_r17_Offset_r17_Db6 = 2
	GoodServingCellEvaluation_r17_Offset_r17_Db8 = 3
)

var goodServingCellEvaluationR17OffsetR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GoodServingCellEvaluation_r17_Offset_r17_Db2, GoodServingCellEvaluation_r17_Offset_r17_Db4, GoodServingCellEvaluation_r17_Offset_r17_Db6, GoodServingCellEvaluation_r17_Offset_r17_Db8},
}

type GoodServingCellEvaluation_r17 struct {
	Offset_r17 *int64
}

func (ie *GoodServingCellEvaluation_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(goodServingCellEvaluationR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Offset_r17 != nil}); err != nil {
		return err
	}

	// 2. offset-r17: enumerated
	{
		if ie.Offset_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Offset_r17, goodServingCellEvaluationR17OffsetR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *GoodServingCellEvaluation_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(goodServingCellEvaluationR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. offset-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(goodServingCellEvaluationR17OffsetR17Constraints)
			if err != nil {
				return err
			}
			ie.Offset_r17 = &idx
		}
	}

	return nil
}
