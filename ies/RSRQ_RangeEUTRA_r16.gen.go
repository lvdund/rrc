// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSRQ-RangeEUTRA-r16 (line 9321).
// RSRQ-RangeEUTRA-r16 ::=   INTEGER (-30..46)

var rSRQRangeEUTRAR16Constraints = per.Constrained(-30, 46)

type RSRQ_RangeEUTRA_r16 struct {
	Value int64
}

func (v *RSRQ_RangeEUTRA_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rSRQRangeEUTRAR16Constraints)
}

func (v *RSRQ_RangeEUTRA_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rSRQRangeEUTRAR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
