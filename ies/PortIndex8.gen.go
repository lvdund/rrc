// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PortIndex8 (line 7293).
// PortIndex8::=                       INTEGER (0..7)

var portIndex8Constraints = per.Constrained(0, 7)

type PortIndex8 struct {
	Value int64
}

func (v *PortIndex8) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, portIndex8Constraints)
}

func (v *PortIndex8) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(portIndex8Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
