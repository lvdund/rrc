// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MaxMIMO-LayersDCI-0-2-r16 (line 12694).
// MaxMIMO-LayersDCI-0-2-r16 ::=           INTEGER (1..4)

var maxMIMOLayersDCI02R16Constraints = per.Constrained(1, 4)

type MaxMIMO_LayersDCI_0_2_r16 struct {
	Value int64
}

func (v *MaxMIMO_LayersDCI_0_2_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, maxMIMOLayersDCI02R16Constraints)
}

func (v *MaxMIMO_LayersDCI_0_2_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(maxMIMOLayersDCI02R16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
