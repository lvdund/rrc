// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TimeSinceSdt-Execution-r19 (line 3568).
// TimeSinceSdt-Execution-r19 ::= INTEGER (0..172800)

var timeSinceSdtExecutionR19Constraints = per.Constrained(0, 172800)

type TimeSinceSdt_Execution_r19 struct {
	Value int64
}

func (v *TimeSinceSdt_Execution_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, timeSinceSdtExecutionR19Constraints)
}

func (v *TimeSinceSdt_Execution_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(timeSinceSdtExecutionR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
