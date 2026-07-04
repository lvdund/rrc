// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSRQ-RangeEUTRA (line 10062).
// RSRQ-RangeEUTRA ::=                 INTEGER (0..34)

var rSRQRangeEUTRAConstraints = per.Constrained(0, 34)

type RSRQ_RangeEUTRA struct {
	Value int64
}

func (v *RSRQ_RangeEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rSRQRangeEUTRAConstraints)
}

func (v *RSRQ_RangeEUTRA) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rSRQRangeEUTRAConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
