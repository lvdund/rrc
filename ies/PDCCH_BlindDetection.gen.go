// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetection (line 11726).
// PDCCH-BlindDetection ::=                INTEGER (1..15)

var pDCCHBlindDetectionConstraints = per.Constrained(1, 15)

type PDCCH_BlindDetection struct {
	Value int64
}

func (v *PDCCH_BlindDetection) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pDCCHBlindDetectionConstraints)
}

func (v *PDCCH_BlindDetection) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pDCCHBlindDetectionConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
