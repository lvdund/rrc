// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-IntendedAreaID-r19 (line 4781).
// MBS-IntendedAreaID-r19  ::= INTEGER (1..maxNrofMBS-Area-r19)

var mBSIntendedAreaIDR19Constraints = per.Constrained(1, common.MaxNrofMBS_Area_r19)

type MBS_IntendedAreaID_r19 struct {
	Value int64
}

func (v *MBS_IntendedAreaID_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mBSIntendedAreaIDR19Constraints)
}

func (v *MBS_IntendedAreaID_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mBSIntendedAreaIDR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
