// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeSinceCPAC-Reconfig-r18 (line 3556).
// TimeSinceCPAC-Reconfig-r18 ::= INTEGER (0.. 1023)

var timeSinceCPACReconfigR18Constraints = per.Constrained(0, 1023)

type TimeSinceCPAC_Reconfig_r18 struct {
	Value int64
}

func (v *TimeSinceCPAC_Reconfig_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeSinceCPACReconfigR18Constraints)
}

func (v *TimeSinceCPAC_Reconfig_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeSinceCPACReconfigR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
