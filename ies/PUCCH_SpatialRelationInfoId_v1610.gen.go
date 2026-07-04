// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-SpatialRelationInfoId-v1610 (line 12318).
// PUCCH-SpatialRelationInfoId-v1610::=    INTEGER (maxNrofSpatialRelationInfos-plus-1..maxNrofSpatialRelationInfos-r16)

var pUCCHSpatialRelationInfoIdV1610Constraints = per.Constrained(common.MaxNrofSpatialRelationInfos_Plus_1, common.MaxNrofSpatialRelationInfos_r16)

type PUCCH_SpatialRelationInfoId_v1610 struct {
	Value int64
}

func (v *PUCCH_SpatialRelationInfoId_v1610) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHSpatialRelationInfoIdV1610Constraints)
}

func (v *PUCCH_SpatialRelationInfoId_v1610) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHSpatialRelationInfoIdV1610Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
