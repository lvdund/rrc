// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BetaOffsets (line 5170).

var betaOffsetsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "betaOffsetACK-Index1", Optional: true},
		{Name: "betaOffsetACK-Index2", Optional: true},
		{Name: "betaOffsetACK-Index3", Optional: true},
		{Name: "betaOffsetCSI-Part1-Index1", Optional: true},
		{Name: "betaOffsetCSI-Part1-Index2", Optional: true},
		{Name: "betaOffsetCSI-Part2-Index1", Optional: true},
		{Name: "betaOffsetCSI-Part2-Index2", Optional: true},
	},
}

var betaOffsetsBetaOffsetACKIndex1Constraints = per.Constrained(0, 31)

var betaOffsetsBetaOffsetACKIndex2Constraints = per.Constrained(0, 31)

var betaOffsetsBetaOffsetACKIndex3Constraints = per.Constrained(0, 31)

var betaOffsetsBetaOffsetCSIPart1Index1Constraints = per.Constrained(0, 31)

var betaOffsetsBetaOffsetCSIPart1Index2Constraints = per.Constrained(0, 31)

var betaOffsetsBetaOffsetCSIPart2Index1Constraints = per.Constrained(0, 31)

var betaOffsetsBetaOffsetCSIPart2Index2Constraints = per.Constrained(0, 31)

type BetaOffsets struct {
	BetaOffsetACK_Index1       *int64
	BetaOffsetACK_Index2       *int64
	BetaOffsetACK_Index3       *int64
	BetaOffsetCSI_Part1_Index1 *int64
	BetaOffsetCSI_Part1_Index2 *int64
	BetaOffsetCSI_Part2_Index1 *int64
	BetaOffsetCSI_Part2_Index2 *int64
}

func (ie *BetaOffsets) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(betaOffsetsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BetaOffsetACK_Index1 != nil, ie.BetaOffsetACK_Index2 != nil, ie.BetaOffsetACK_Index3 != nil, ie.BetaOffsetCSI_Part1_Index1 != nil, ie.BetaOffsetCSI_Part1_Index2 != nil, ie.BetaOffsetCSI_Part2_Index1 != nil, ie.BetaOffsetCSI_Part2_Index2 != nil}); err != nil {
		return err
	}

	// 2. betaOffsetACK-Index1: integer
	{
		if ie.BetaOffsetACK_Index1 != nil {
			if err := e.EncodeInteger(*ie.BetaOffsetACK_Index1, betaOffsetsBetaOffsetACKIndex1Constraints); err != nil {
				return err
			}
		}
	}

	// 3. betaOffsetACK-Index2: integer
	{
		if ie.BetaOffsetACK_Index2 != nil {
			if err := e.EncodeInteger(*ie.BetaOffsetACK_Index2, betaOffsetsBetaOffsetACKIndex2Constraints); err != nil {
				return err
			}
		}
	}

	// 4. betaOffsetACK-Index3: integer
	{
		if ie.BetaOffsetACK_Index3 != nil {
			if err := e.EncodeInteger(*ie.BetaOffsetACK_Index3, betaOffsetsBetaOffsetACKIndex3Constraints); err != nil {
				return err
			}
		}
	}

	// 5. betaOffsetCSI-Part1-Index1: integer
	{
		if ie.BetaOffsetCSI_Part1_Index1 != nil {
			if err := e.EncodeInteger(*ie.BetaOffsetCSI_Part1_Index1, betaOffsetsBetaOffsetCSIPart1Index1Constraints); err != nil {
				return err
			}
		}
	}

	// 6. betaOffsetCSI-Part1-Index2: integer
	{
		if ie.BetaOffsetCSI_Part1_Index2 != nil {
			if err := e.EncodeInteger(*ie.BetaOffsetCSI_Part1_Index2, betaOffsetsBetaOffsetCSIPart1Index2Constraints); err != nil {
				return err
			}
		}
	}

	// 7. betaOffsetCSI-Part2-Index1: integer
	{
		if ie.BetaOffsetCSI_Part2_Index1 != nil {
			if err := e.EncodeInteger(*ie.BetaOffsetCSI_Part2_Index1, betaOffsetsBetaOffsetCSIPart2Index1Constraints); err != nil {
				return err
			}
		}
	}

	// 8. betaOffsetCSI-Part2-Index2: integer
	{
		if ie.BetaOffsetCSI_Part2_Index2 != nil {
			if err := e.EncodeInteger(*ie.BetaOffsetCSI_Part2_Index2, betaOffsetsBetaOffsetCSIPart2Index2Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BetaOffsets) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(betaOffsetsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. betaOffsetACK-Index1: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(betaOffsetsBetaOffsetACKIndex1Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetACK_Index1 = &v
		}
	}

	// 3. betaOffsetACK-Index2: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(betaOffsetsBetaOffsetACKIndex2Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetACK_Index2 = &v
		}
	}

	// 4. betaOffsetACK-Index3: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(betaOffsetsBetaOffsetACKIndex3Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetACK_Index3 = &v
		}
	}

	// 5. betaOffsetCSI-Part1-Index1: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(betaOffsetsBetaOffsetCSIPart1Index1Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetCSI_Part1_Index1 = &v
		}
	}

	// 6. betaOffsetCSI-Part1-Index2: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(betaOffsetsBetaOffsetCSIPart1Index2Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetCSI_Part1_Index2 = &v
		}
	}

	// 7. betaOffsetCSI-Part2-Index1: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(betaOffsetsBetaOffsetCSIPart2Index1Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetCSI_Part2_Index1 = &v
		}
	}

	// 8. betaOffsetCSI-Part2-Index2: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(betaOffsetsBetaOffsetCSIPart2Index2Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetCSI_Part2_Index2 = &v
		}
	}

	return nil
}
