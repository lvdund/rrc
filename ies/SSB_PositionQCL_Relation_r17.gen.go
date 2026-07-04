// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-PositionQCL-Relation-r17 (line 15918).
// SSB-PositionQCL-Relation-r17 ::=  ENUMERATED {n32, n64}

const (
	SSB_PositionQCL_Relation_r17_N32 = 0
	SSB_PositionQCL_Relation_r17_N64 = 1
)

var sSBPositionQCLRelationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_PositionQCL_Relation_r17_N32, SSB_PositionQCL_Relation_r17_N64},
}

type SSB_PositionQCL_Relation_r17 struct {
	Value int64
}

func (v *SSB_PositionQCL_Relation_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sSBPositionQCLRelationR17Constraints)
}

func (v *SSB_PositionQCL_Relation_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sSBPositionQCLRelationR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
