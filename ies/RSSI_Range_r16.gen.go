// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSSI-Range-r16 (line 14216).
// RSSI-Range-r16 ::=                  INTEGER(0..76)

var rSSIRangeR16Constraints = per.Constrained(0, 76)

type RSSI_Range_r16 struct {
	Value int64
}

func (v *RSSI_Range_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rSSIRangeR16Constraints)
}

func (v *RSSI_Range_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rSSIRangeR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
