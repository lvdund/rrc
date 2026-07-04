// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-StartingOffsets-r16 (line 6692).

var cGStartingOffsetsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cg-StartingFullBW-InsideCOT-r16", Optional: true},
		{Name: "cg-StartingFullBW-OutsideCOT-r16", Optional: true},
		{Name: "cg-StartingPartialBW-InsideCOT-r16", Optional: true},
		{Name: "cg-StartingPartialBW-OutsideCOT-r16", Optional: true},
	},
}

var cGStartingOffsetsR16CgStartingFullBWInsideCOTR16Constraints = per.SizeRange(1, 7)

var cGStartingOffsetsR16CgStartingFullBWOutsideCOTR16Constraints = per.SizeRange(1, 7)

var cGStartingOffsetsR16CgStartingPartialBWInsideCOTR16Constraints = per.Constrained(0, 6)

var cGStartingOffsetsR16CgStartingPartialBWOutsideCOTR16Constraints = per.Constrained(0, 6)

type CG_StartingOffsets_r16 struct {
	Cg_StartingFullBW_InsideCOT_r16     []int64
	Cg_StartingFullBW_OutsideCOT_r16    []int64
	Cg_StartingPartialBW_InsideCOT_r16  *int64
	Cg_StartingPartialBW_OutsideCOT_r16 *int64
}

func (ie *CG_StartingOffsets_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGStartingOffsetsR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cg_StartingFullBW_InsideCOT_r16 != nil, ie.Cg_StartingFullBW_OutsideCOT_r16 != nil, ie.Cg_StartingPartialBW_InsideCOT_r16 != nil, ie.Cg_StartingPartialBW_OutsideCOT_r16 != nil}); err != nil {
		return err
	}

	// 2. cg-StartingFullBW-InsideCOT-r16: sequence-of
	{
		if ie.Cg_StartingFullBW_InsideCOT_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(cGStartingOffsetsR16CgStartingFullBWInsideCOTR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cg_StartingFullBW_InsideCOT_r16))); err != nil {
				return err
			}
			for i := range ie.Cg_StartingFullBW_InsideCOT_r16 {
				if err := e.EncodeInteger(ie.Cg_StartingFullBW_InsideCOT_r16[i], per.Constrained(0, 6)); err != nil {
					return err
				}
			}
		}
	}

	// 3. cg-StartingFullBW-OutsideCOT-r16: sequence-of
	{
		if ie.Cg_StartingFullBW_OutsideCOT_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(cGStartingOffsetsR16CgStartingFullBWOutsideCOTR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cg_StartingFullBW_OutsideCOT_r16))); err != nil {
				return err
			}
			for i := range ie.Cg_StartingFullBW_OutsideCOT_r16 {
				if err := e.EncodeInteger(ie.Cg_StartingFullBW_OutsideCOT_r16[i], per.Constrained(0, 6)); err != nil {
					return err
				}
			}
		}
	}

	// 4. cg-StartingPartialBW-InsideCOT-r16: integer
	{
		if ie.Cg_StartingPartialBW_InsideCOT_r16 != nil {
			if err := e.EncodeInteger(*ie.Cg_StartingPartialBW_InsideCOT_r16, cGStartingOffsetsR16CgStartingPartialBWInsideCOTR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. cg-StartingPartialBW-OutsideCOT-r16: integer
	{
		if ie.Cg_StartingPartialBW_OutsideCOT_r16 != nil {
			if err := e.EncodeInteger(*ie.Cg_StartingPartialBW_OutsideCOT_r16, cGStartingOffsetsR16CgStartingPartialBWOutsideCOTR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CG_StartingOffsets_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGStartingOffsetsR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cg-StartingFullBW-InsideCOT-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(cGStartingOffsetsR16CgStartingFullBWInsideCOTR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cg_StartingFullBW_InsideCOT_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 6))
				if err != nil {
					return err
				}
				ie.Cg_StartingFullBW_InsideCOT_r16[i] = v
			}
		}
	}

	// 3. cg-StartingFullBW-OutsideCOT-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(cGStartingOffsetsR16CgStartingFullBWOutsideCOTR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cg_StartingFullBW_OutsideCOT_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 6))
				if err != nil {
					return err
				}
				ie.Cg_StartingFullBW_OutsideCOT_r16[i] = v
			}
		}
	}

	// 4. cg-StartingPartialBW-InsideCOT-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(cGStartingOffsetsR16CgStartingPartialBWInsideCOTR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_StartingPartialBW_InsideCOT_r16 = &v
		}
	}

	// 5. cg-StartingPartialBW-OutsideCOT-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(cGStartingOffsetsR16CgStartingPartialBWOutsideCOTR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_StartingPartialBW_OutsideCOT_r16 = &v
		}
	}

	return nil
}
