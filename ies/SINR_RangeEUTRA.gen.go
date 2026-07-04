// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SINR-RangeEUTRA (line 10064).
// SINR-RangeEUTRA ::=                 INTEGER (0..127)

var sINRRangeEUTRAConstraints = per.Constrained(0, 127)

type SINR_RangeEUTRA struct {
	Value int64
}

func (v *SINR_RangeEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sINRRangeEUTRAConstraints)
}

func (v *SINR_RangeEUTRA) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sINRRangeEUTRAConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
