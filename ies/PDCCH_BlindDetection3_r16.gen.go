// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetection3-r16 (line 11747).
// PDCCH-BlindDetection3-r16 ::=                INTEGER (1..15)

var pDCCHBlindDetection3R16Constraints = per.Constrained(1, 15)

type PDCCH_BlindDetection3_r16 struct {
	Value int64
}

func (v *PDCCH_BlindDetection3_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pDCCHBlindDetection3R16Constraints)
}

func (v *PDCCH_BlindDetection3_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pDCCHBlindDetection3R16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
