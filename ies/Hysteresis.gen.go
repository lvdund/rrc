// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Hysteresis (line 8500).
// Hysteresis ::=                      INTEGER (0..30)

var hysteresisConstraints = per.Constrained(0, 30)

type Hysteresis struct {
	Value int64
}

func (v *Hysteresis) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, hysteresisConstraints)
}

func (v *Hysteresis) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(hysteresisConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
