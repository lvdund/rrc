// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeSinceSHR-r18 (line 3566).
// TimeSinceSHR-r18 ::= INTEGER (0..172800)

var timeSinceSHRR18Constraints = per.Constrained(0, 172800)

type TimeSinceSHR_r18 struct {
	Value int64
}

func (v *TimeSinceSHR_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeSinceSHRR18Constraints)
}

func (v *TimeSinceSHR_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeSinceSHRR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
