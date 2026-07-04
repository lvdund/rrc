// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ControlResourceSetId-v1610 (line 6800).
// ControlResourceSetId-v1610 ::=          INTEGER (maxNrofControlResourceSets..maxNrofControlResourceSets-1-r16)

var controlResourceSetIdV1610Constraints = per.Constrained(common.MaxNrofControlResourceSets, common.MaxNrofControlResourceSets_1_r16)

type ControlResourceSetId_v1610 struct {
	Value int64
}

func (v *ControlResourceSetId_v1610) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, controlResourceSetIdV1610Constraints)
}

func (v *ControlResourceSetId_v1610) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(controlResourceSetIdV1610Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
