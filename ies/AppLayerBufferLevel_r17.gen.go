// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AppLayerBufferLevel-r17 (line 791).
// AppLayerBufferLevel-r17 ::= INTEGER (0..30000)

var appLayerBufferLevelR17Constraints = per.Constrained(0, 30000)

type AppLayerBufferLevel_r17 struct {
	Value int64
}

func (v *AppLayerBufferLevel_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, appLayerBufferLevelR17Constraints)
}

func (v *AppLayerBufferLevel_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(appLayerBufferLevelR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
