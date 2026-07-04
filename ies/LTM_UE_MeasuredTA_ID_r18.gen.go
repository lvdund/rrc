// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-UE-MeasuredTA-ID-r18 (line 8665).
// LTM-UE-MeasuredTA-ID-r18 ::= INTEGER (1..maxNrofLTM-Configs-plus1-r18)

var lTMUEMeasuredTAIDR18Constraints = per.Constrained(1, common.MaxNrofLTM_Configs_Plus1_r18)

type LTM_UE_MeasuredTA_ID_r18 struct {
	Value int64
}

func (v *LTM_UE_MeasuredTA_ID_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, lTMUEMeasuredTAIDR18Constraints)
}

func (v *LTM_UE_MeasuredTA_ID_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(lTMUEMeasuredTAIDR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
