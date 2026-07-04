// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-MeasId-r16 (line 27547).
// SL-MeasId-r16 ::=                   INTEGER (1..maxNrofSL-MeasId-r16)

var sLMeasIdR16Constraints = per.Constrained(1, common.MaxNrofSL_MeasId_r16)

type SL_MeasId_r16 struct {
	Value int64
}

func (v *SL_MeasId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLMeasIdR16Constraints)
}

func (v *SL_MeasId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLMeasIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
