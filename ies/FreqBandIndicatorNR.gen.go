// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqBandIndicatorNR (line 8351).
// FreqBandIndicatorNR ::=             INTEGER (1..1024)

var freqBandIndicatorNRConstraints = per.Constrained(1, 1024)

type FreqBandIndicatorNR struct {
	Value int64
}

func (v *FreqBandIndicatorNR) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, freqBandIndicatorNRConstraints)
}

func (v *FreqBandIndicatorNR) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(freqBandIndicatorNRConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
