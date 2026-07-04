// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: P0AlphaSet-r17 (line 16336).

var p0AlphaSetR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "p0-r17", Optional: true},
		{Name: "alpha-r17", Optional: true},
		{Name: "closedLoopIndex-r17"},
	},
}

var p0AlphaSetR17P0R17Constraints = per.Constrained(-16, 15)

const (
	P0AlphaSet_r17_ClosedLoopIndex_r17_I0 = 0
	P0AlphaSet_r17_ClosedLoopIndex_r17_I1 = 1
)

var p0AlphaSetR17ClosedLoopIndexR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{P0AlphaSet_r17_ClosedLoopIndex_r17_I0, P0AlphaSet_r17_ClosedLoopIndex_r17_I1},
}

type P0AlphaSet_r17 struct {
	P0_r17              *int64
	Alpha_r17           *Alpha
	ClosedLoopIndex_r17 int64
}

func (ie *P0AlphaSet_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(p0AlphaSetR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.P0_r17 != nil, ie.Alpha_r17 != nil}); err != nil {
		return err
	}

	// 2. p0-r17: integer
	{
		if ie.P0_r17 != nil {
			if err := e.EncodeInteger(*ie.P0_r17, p0AlphaSetR17P0R17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. alpha-r17: ref
	{
		if ie.Alpha_r17 != nil {
			if err := ie.Alpha_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. closedLoopIndex-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.ClosedLoopIndex_r17, p0AlphaSetR17ClosedLoopIndexR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *P0AlphaSet_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(p0AlphaSetR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. p0-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(p0AlphaSetR17P0R17Constraints)
			if err != nil {
				return err
			}
			ie.P0_r17 = &v
		}
	}

	// 3. alpha-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Alpha_r17 = new(Alpha)
			if err := ie.Alpha_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. closedLoopIndex-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(p0AlphaSetR17ClosedLoopIndexR17Constraints)
		if err != nil {
			return err
		}
		ie.ClosedLoopIndex_r17 = v2
	}

	return nil
}
