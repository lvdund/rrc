// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SINR-Range (line 15010).
// SINR-Range ::=                      INTEGER(0..127)

var sINRRangeConstraints = per.Constrained(0, 127)

type SINR_Range struct {
	Value int64
}

func (v *SINR_Range) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sINRRangeConstraints)
}

func (v *SINR_Range) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sINRRangeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
