// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-NoSecurityChangeId-r19 (line 8741).
// LTM-NoSecurityChangeId-r19 ::= INTEGER (1..maxNrofLTM-Configs-plus1-r18)

var lTMNoSecurityChangeIdR19Constraints = per.Constrained(1, common.MaxNrofLTM_Configs_Plus1_r18)

type LTM_NoSecurityChangeId_r19 struct {
	Value int64
}

func (v *LTM_NoSecurityChangeId_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, lTMNoSecurityChangeIdR19Constraints)
}

func (v *LTM_NoSecurityChangeId_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(lTMNoSecurityChangeIdR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
