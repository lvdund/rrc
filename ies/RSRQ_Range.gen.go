// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSRQ-Range (line 14211).
// RSRQ-Range ::=                      INTEGER(0..127)

var rSRQRangeConstraints = per.Constrained(0, 127)

type RSRQ_Range struct {
	Value int64
}

func (v *RSRQ_Range) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rSRQRangeConstraints)
}

func (v *RSRQ_Range) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rSRQRangeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
