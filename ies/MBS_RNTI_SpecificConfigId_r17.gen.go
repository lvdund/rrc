// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-RNTI-SpecificConfigId-r17 (line 9069).
// MBS-RNTI-SpecificConfigId-r17 ::= INTEGER (0..maxG-RNTI-1-r17)

var mBSRNTISpecificConfigIdR17Constraints = per.Constrained(0, common.MaxG_RNTI_1_r17)

type MBS_RNTI_SpecificConfigId_r17 struct {
	Value int64
}

func (v *MBS_RNTI_SpecificConfigId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mBSRNTISpecificConfigIdR17Constraints)
}

func (v *MBS_RNTI_SpecificConfigId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mBSRNTISpecificConfigIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
