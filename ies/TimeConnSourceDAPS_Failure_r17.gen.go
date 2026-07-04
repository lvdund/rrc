// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeConnSourceDAPS-Failure-r17 (line 3558).
// TimeConnSourceDAPS-Failure-r17 ::= INTEGER (0..1023)

var timeConnSourceDAPSFailureR17Constraints = per.Constrained(0, 1023)

type TimeConnSourceDAPS_Failure_r17 struct {
	Value int64
}

func (v *TimeConnSourceDAPS_Failure_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeConnSourceDAPSFailureR17Constraints)
}

func (v *TimeConnSourceDAPS_Failure_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeConnSourceDAPSFailureR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
