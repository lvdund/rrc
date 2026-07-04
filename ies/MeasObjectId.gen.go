// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasObjectId (line 9426).
// MeasObjectId ::=                    INTEGER (1..maxNrofObjectId)

var measObjectIdConstraints = per.Constrained(1, common.MaxNrofObjectId)

type MeasObjectId struct {
	Value int64
}

func (v *MeasObjectId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, measObjectIdConstraints)
}

func (v *MeasObjectId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(measObjectIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
