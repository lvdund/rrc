// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ResourcePoolID-r16 (line 26842).
// SL-ResourcePoolID-r16 ::=        INTEGER (1..maxNrofPoolID-r16)

var sLResourcePoolIDR16Constraints = per.Constrained(1, common.MaxNrofPoolID_r16)

type SL_ResourcePoolID_r16 struct {
	Value int64
}

func (v *SL_ResourcePoolID_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLResourcePoolIDR16Constraints)
}

func (v *SL_ResourcePoolID_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLResourcePoolIDR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
