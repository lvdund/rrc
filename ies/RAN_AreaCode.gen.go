// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RAN-AreaCode (line 13253).
// RAN-AreaCode ::=                INTEGER (0..255)

var rANAreaCodeConstraints = per.Constrained(0, 255)

type RAN_AreaCode struct {
	Value int64
}

func (v *RAN_AreaCode) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rANAreaCodeConstraints)
}

func (v *RAN_AreaCode) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rANAreaCodeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
