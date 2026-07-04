// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ControlResourceSetId (line 6796).
// ControlResourceSetId ::=                INTEGER (0..maxNrofControlResourceSets-1)

var controlResourceSetIdConstraints = per.Constrained(0, common.MaxNrofControlResourceSets_1)

type ControlResourceSetId struct {
	Value int64
}

func (v *ControlResourceSetId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, controlResourceSetIdConstraints)
}

func (v *ControlResourceSetId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(controlResourceSetIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
