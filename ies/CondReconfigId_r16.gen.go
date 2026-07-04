// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CondReconfigId-r16 (line 6458).
// CondReconfigId-r16 ::=                    INTEGER (1.. maxNrofCondCells-r16)

var condReconfigIdR16Constraints = per.Constrained(1, common.MaxNrofCondCells_r16)

type CondReconfigId_r16 struct {
	Value int64
}

func (v *CondReconfigId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, condReconfigIdR16Constraints)
}

func (v *CondReconfigId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(condReconfigIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
