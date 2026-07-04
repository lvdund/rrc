// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PortIndex4 (line 7294).
// PortIndex4::=                       INTEGER (0..3)

var portIndex4Constraints = per.Constrained(0, 3)

type PortIndex4 struct {
	Value int64
}

func (v *PortIndex4) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, portIndex4Constraints)
}

func (v *PortIndex4) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(portIndex4Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
