// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-MeasObjectId-r16 (line 27560).
// SL-MeasObjectId-r16 ::=                 INTEGER (1..maxNrofSL-ObjectId-r16)

var sLMeasObjectIdR16Constraints = per.Constrained(1, common.MaxNrofSL_ObjectId_r16)

type SL_MeasObjectId_r16 struct {
	Value int64
}

func (v *SL_MeasObjectId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLMeasObjectIdR16Constraints)
}

func (v *SL_MeasObjectId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLMeasObjectIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
