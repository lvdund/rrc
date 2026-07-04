// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ControlResourceSetId-r16 (line 6798).
// ControlResourceSetId-r16 ::=            INTEGER (0..maxNrofControlResourceSets-1-r16)

var controlResourceSetIdR16Constraints = per.Constrained(0, common.MaxNrofControlResourceSets_1_r16)

type ControlResourceSetId_r16 struct {
	Value int64
}

func (v *ControlResourceSetId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, controlResourceSetIdR16Constraints)
}

func (v *ControlResourceSetId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(controlResourceSetIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
