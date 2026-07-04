// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSRP-Range (line 14206).
// RSRP-Range ::=                      INTEGER(0..127)

var rSRPRangeConstraints = per.Constrained(0, 127)

type RSRP_Range struct {
	Value int64
}

func (v *RSRP_Range) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rSRPRangeConstraints)
}

func (v *RSRP_Range) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rSRPRangeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
