// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ARFCN-ValueUTRA-FDD-r16 (line 5029).
// ARFCN-ValueUTRA-FDD-r16 ::=                INTEGER (0..16383)

var aRFCNValueUTRAFDDR16Constraints = per.Constrained(0, 16383)

type ARFCN_ValueUTRA_FDD_r16 struct {
	Value int64
}

func (v *ARFCN_ValueUTRA_FDD_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, aRFCNValueUTRAFDDR16Constraints)
}

func (v *ARFCN_ValueUTRA_FDD_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(aRFCNValueUTRAFDDR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
