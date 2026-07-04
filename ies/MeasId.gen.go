// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasId (line 9232).
// MeasId ::=                          INTEGER (1..maxNrofMeasId)

var measIdConstraints = per.Constrained(1, common.MaxNrofMeasId)

type MeasId struct {
	Value int64
}

func (v *MeasId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, measIdConstraints)
}

func (v *MeasId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(measIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
