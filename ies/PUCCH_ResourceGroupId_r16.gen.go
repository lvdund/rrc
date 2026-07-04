// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-ResourceGroupId-r16 (line 12161).
// PUCCH-ResourceGroupId-r16 ::=              INTEGER (0..maxNrofPUCCH-ResourceGroups-1-r16)

var pUCCHResourceGroupIdR16Constraints = per.Constrained(0, common.MaxNrofPUCCH_ResourceGroups_1_r16)

type PUCCH_ResourceGroupId_r16 struct {
	Value int64
}

func (v *PUCCH_ResourceGroupId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHResourceGroupIdR16Constraints)
}

func (v *PUCCH_ResourceGroupId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHResourceGroupIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
