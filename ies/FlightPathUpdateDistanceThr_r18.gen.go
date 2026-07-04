// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FlightPathUpdateDistanceThr-r18 (line 4971).
// FlightPathUpdateDistanceThr-r18 ::=   INTEGER (0..1023)

var flightPathUpdateDistanceThrR18Constraints = per.Constrained(0, 1023)

type FlightPathUpdateDistanceThr_r18 struct {
	Value int64
}

func (v *FlightPathUpdateDistanceThr_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, flightPathUpdateDistanceThrR18Constraints)
}

func (v *FlightPathUpdateDistanceThr_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(flightPathUpdateDistanceThrR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
