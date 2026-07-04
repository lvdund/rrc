// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-SpatialRelationInfoId (line 12314).
// PUCCH-SpatialRelationInfoId ::=         INTEGER (1..maxNrofSpatialRelationInfos)

var pUCCHSpatialRelationInfoIdConstraints = per.Constrained(1, common.MaxNrofSpatialRelationInfos)

type PUCCH_SpatialRelationInfoId struct {
	Value int64
}

func (v *PUCCH_SpatialRelationInfoId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHSpatialRelationInfoIdConstraints)
}

func (v *PUCCH_SpatialRelationInfoId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHSpatialRelationInfoIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
