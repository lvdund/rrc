// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: P0-PUSCH-r16 (line 12659).
// P0-PUSCH-r16 ::=                    INTEGER (-16..15)

var p0PUSCHR16Constraints = per.Constrained(-16, 15)

type P0_PUSCH_r16 struct {
	Value int64
}

func (v *P0_PUSCH_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, p0PUSCHR16Constraints)
}

func (v *P0_PUSCH_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(p0PUSCHR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
