// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CLI-RSSI-Range-r16 (line 6022).
// CLI-RSSI-Range-r16 ::=                      INTEGER(0..76)

var cLIRSSIRangeR16Constraints = per.Constrained(0, 76)

type CLI_RSSI_Range_r16 struct {
	Value int64
}

func (v *CLI_RSSI_Range_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, cLIRSSIRangeR16Constraints)
}

func (v *CLI_RSSI_Range_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(cLIRSSIRangeR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
