// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TrackingAreaCode (line 16168).
// TrackingAreaCode ::= BIT STRING (SIZE (24))

var trackingAreaCodeConstraints = per.FixedSize(24)

type TrackingAreaCode struct {
	Value per.BitString
}

func (v *TrackingAreaCode) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, trackingAreaCodeConstraints)
}

func (v *TrackingAreaCode) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(trackingAreaCodeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
