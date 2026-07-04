// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSRP-RangeEUTRA (line 10060).
// RSRP-RangeEUTRA ::=                 INTEGER (0..97)

var rSRPRangeEUTRAConstraints = per.Constrained(0, 97)

type RSRP_RangeEUTRA struct {
	Value int64
}

func (v *RSRP_RangeEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rSRPRangeEUTRAConstraints)
}

func (v *RSRP_RangeEUTRA) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rSRPRangeEUTRAConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
