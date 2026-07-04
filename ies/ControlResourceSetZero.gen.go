// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ControlResourceSetZero (line 6805).
// ControlResourceSetZero ::=                  INTEGER (0..15)

var controlResourceSetZeroConstraints = per.Constrained(0, 15)

type ControlResourceSetZero struct {
	Value int64
}

func (v *ControlResourceSetZero) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, controlResourceSetZeroConstraints)
}

func (v *ControlResourceSetZero) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(controlResourceSetZeroConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
