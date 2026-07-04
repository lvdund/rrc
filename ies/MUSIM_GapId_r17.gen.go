// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-GapId-r17 (line 10216).
// MUSIM-GapId-r17 ::=                  INTEGER (0..2)

var mUSIMGapIdR17Constraints = per.Constrained(0, 2)

type MUSIM_GapId_r17 struct {
	Value int64
}

func (v *MUSIM_GapId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mUSIMGapIdR17Constraints)
}

func (v *MUSIM_GapId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mUSIMGapIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
