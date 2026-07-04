// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxMIMO-LayersDL-r16 (line 11447).
// MaxMIMO-LayersDL-r16 ::=                INTEGER (1..8)

var maxMIMOLayersDLR16Constraints = per.Constrained(1, 8)

type MaxMIMO_LayersDL_r16 struct {
	Value int64
}

func (v *MaxMIMO_LayersDL_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, maxMIMOLayersDLR16Constraints)
}

func (v *MaxMIMO_LayersDL_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(maxMIMOLayersDLR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
