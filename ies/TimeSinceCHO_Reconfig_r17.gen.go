// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeSinceCHO-Reconfig-r17 (line 3554).
// TimeSinceCHO-Reconfig-r17 ::= INTEGER (0..1023)

var timeSinceCHOReconfigR17Constraints = per.Constrained(0, 1023)

type TimeSinceCHO_Reconfig_r17 struct {
	Value int64
}

func (v *TimeSinceCHO_Reconfig_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeSinceCHOReconfigR17Constraints)
}

func (v *TimeSinceCHO_Reconfig_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeSinceCHOReconfigR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
