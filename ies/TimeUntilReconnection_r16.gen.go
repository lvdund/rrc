// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeUntilReconnection-r16 (line 3552).
// TimeUntilReconnection-r16 ::= INTEGER (0..172800)

var timeUntilReconnectionR16Constraints = per.Constrained(0, 172800)

type TimeUntilReconnection_r16 struct {
	Value int64
}

func (v *TimeUntilReconnection_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeUntilReconnectionR16Constraints)
}

func (v *TimeUntilReconnection_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeUntilReconnectionR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
