// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-PositionQCL-Relation-r16 (line 15916).
// SSB-PositionQCL-Relation-r16 ::=  ENUMERATED {n1,n2,n4,n8}

const (
	SSB_PositionQCL_Relation_r16_N1 = 0
	SSB_PositionQCL_Relation_r16_N2 = 1
	SSB_PositionQCL_Relation_r16_N4 = 2
	SSB_PositionQCL_Relation_r16_N8 = 3
)

var sSBPositionQCLRelationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_PositionQCL_Relation_r16_N1, SSB_PositionQCL_Relation_r16_N2, SSB_PositionQCL_Relation_r16_N4, SSB_PositionQCL_Relation_r16_N8},
}

type SSB_PositionQCL_Relation_r16 struct {
	Value int64
}

func (v *SSB_PositionQCL_Relation_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, sSBPositionQCLRelationR16Constraints)
}

func (v *SSB_PositionQCL_Relation_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(sSBPositionQCLRelationR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
