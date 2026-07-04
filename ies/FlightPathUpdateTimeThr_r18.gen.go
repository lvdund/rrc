// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FlightPathUpdateTimeThr-r18 (line 4973).
// FlightPathUpdateTimeThr-r18 ::=       INTEGER (0..16383)

var flightPathUpdateTimeThrR18Constraints = per.Constrained(0, 16383)

type FlightPathUpdateTimeThr_r18 struct {
	Value int64
}

func (v *FlightPathUpdateTimeThr_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, flightPathUpdateTimeThrR18Constraints)
}

func (v *FlightPathUpdateTimeThr_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(flightPathUpdateTimeThrR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
