// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PortIndex2 (line 7295).
// PortIndex2::=                       INTEGER (0..1)

var portIndex2Constraints = per.Constrained(0, 1)

type PortIndex2 struct {
	Value int64
}

func (v *PortIndex2) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, portIndex2Constraints)
}

func (v *PortIndex2) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(portIndex2Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
