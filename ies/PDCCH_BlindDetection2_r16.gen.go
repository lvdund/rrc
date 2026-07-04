// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetection2-r16 (line 11745).
// PDCCH-BlindDetection2-r16 ::=                INTEGER (1..15)

var pDCCHBlindDetection2R16Constraints = per.Constrained(1, 15)

type PDCCH_BlindDetection2_r16 struct {
	Value int64
}

func (v *PDCCH_BlindDetection2_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pDCCHBlindDetection2R16Constraints)
}

func (v *PDCCH_BlindDetection2_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pDCCHBlindDetection2R16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
