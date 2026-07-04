// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-SpatialRelationInfoId-r16 (line 12316).
// PUCCH-SpatialRelationInfoId-r16 ::=     INTEGER (1..maxNrofSpatialRelationInfos-r16)

var pUCCHSpatialRelationInfoIdR16Constraints = per.Constrained(1, common.MaxNrofSpatialRelationInfos_r16)

type PUCCH_SpatialRelationInfoId_r16 struct {
	Value int64
}

func (v *PUCCH_SpatialRelationInfoId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHSpatialRelationInfoIdR16Constraints)
}

func (v *PUCCH_SpatialRelationInfoId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHSpatialRelationInfoIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
