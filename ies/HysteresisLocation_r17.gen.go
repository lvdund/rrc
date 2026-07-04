// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: HysteresisLocation-r17 (line 8510).
// HysteresisLocation-r17 ::=          INTEGER (0..32768)

var hysteresisLocationR17Constraints = per.Constrained(0, 32768)

type HysteresisLocation_r17 struct {
	Value int64
}

func (v *HysteresisLocation_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, hysteresisLocationR17Constraints)
}

func (v *HysteresisLocation_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(hysteresisLocationR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
