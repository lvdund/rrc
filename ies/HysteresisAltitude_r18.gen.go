// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HysteresisAltitude-r18 (line 8505).
// HysteresisAltitude-r18 ::=                      INTEGER (0..64)

var hysteresisAltitudeR18Constraints = per.Constrained(0, 64)

type HysteresisAltitude_r18 struct {
	Value int64
}

func (v *HysteresisAltitude_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, hysteresisAltitudeR18Constraints)
}

func (v *HysteresisAltitude_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(hysteresisAltitudeR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
