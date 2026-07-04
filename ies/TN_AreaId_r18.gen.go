// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TN-AreaId-r18 (line 16163).
// TN-AreaId-r18 ::=                  INTEGER (1..maxTN-AreaInfo-r18)

var tNAreaIdR18Constraints = per.Constrained(1, common.MaxTN_AreaInfo_r18)

type TN_AreaId_r18 struct {
	Value int64
}

func (v *TN_AreaId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, tNAreaIdR18Constraints)
}

func (v *TN_AreaId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(tNAreaIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
