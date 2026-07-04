// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: P0-PUSCH-AlphaSet (line 12584).

var p0PUSCHAlphaSetConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "p0-PUSCH-AlphaSetId"},
		{Name: "p0", Optional: true},
		{Name: "alpha", Optional: true},
	},
}

var p0PUSCHAlphaSetP0Constraints = per.Constrained(-16, 15)

type P0_PUSCH_AlphaSet struct {
	P0_PUSCH_AlphaSetId P0_PUSCH_AlphaSetId
	P0                  *int64
	Alpha               *Alpha
}

func (ie *P0_PUSCH_AlphaSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(p0PUSCHAlphaSetConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.P0 != nil, ie.Alpha != nil}); err != nil {
		return err
	}

	// 2. p0-PUSCH-AlphaSetId: ref
	{
		if err := ie.P0_PUSCH_AlphaSetId.Encode(e); err != nil {
			return err
		}
	}

	// 3. p0: integer
	{
		if ie.P0 != nil {
			if err := e.EncodeInteger(*ie.P0, p0PUSCHAlphaSetP0Constraints); err != nil {
				return err
			}
		}
	}

	// 4. alpha: ref
	{
		if ie.Alpha != nil {
			if err := ie.Alpha.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *P0_PUSCH_AlphaSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(p0PUSCHAlphaSetConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. p0-PUSCH-AlphaSetId: ref
	{
		if err := ie.P0_PUSCH_AlphaSetId.Decode(d); err != nil {
			return err
		}
	}

	// 3. p0: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(p0PUSCHAlphaSetP0Constraints)
			if err != nil {
				return err
			}
			ie.P0 = &v
		}
	}

	// 4. alpha: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Alpha = new(Alpha)
			if err := ie.Alpha.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
