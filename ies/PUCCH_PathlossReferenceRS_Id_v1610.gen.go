// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-PathlossReferenceRS-Id-v1610 (line 12228).
// PUCCH-PathlossReferenceRS-Id-v1610 ::=      INTEGER (maxNrofPUCCH-PathlossReferenceRSs..maxNrofPUCCH-PathlossReferenceRSs-1-r16)

var pUCCHPathlossReferenceRSIdV1610Constraints = per.Constrained(common.MaxNrofPUCCH_PathlossReferenceRSs, common.MaxNrofPUCCH_PathlossReferenceRSs_1_r16)

type PUCCH_PathlossReferenceRS_Id_v1610 struct {
	Value int64
}

func (v *PUCCH_PathlossReferenceRS_Id_v1610) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHPathlossReferenceRSIdV1610Constraints)
}

func (v *PUCCH_PathlossReferenceRS_Id_v1610) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHPathlossReferenceRSIdV1610Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
