// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetection4-r17 (line 11749).
// PDCCH-BlindDetection4-r17 ::=                INTEGER (1..15)

var pDCCHBlindDetection4R17Constraints = per.Constrained(1, 15)

type PDCCH_BlindDetection4_r17 struct {
	Value int64
}

func (v *PDCCH_BlindDetection4_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pDCCHBlindDetection4R17Constraints)
}

func (v *PDCCH_BlindDetection4_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pDCCHBlindDetection4R17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
