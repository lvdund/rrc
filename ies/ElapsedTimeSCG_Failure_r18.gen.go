// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ElapsedTimeSCG-Failure-r18 (line 3564).
// ElapsedTimeSCG-Failure-r18 ::= INTEGER (0..1023)

var elapsedTimeSCGFailureR18Constraints = per.Constrained(0, 1023)

type ElapsedTimeSCG_Failure_r18 struct {
	Value int64
}

func (v *ElapsedTimeSCG_Failure_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, elapsedTimeSCGFailureR18Constraints)
}

func (v *ElapsedTimeSCG_Failure_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(elapsedTimeSCGFailureR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
