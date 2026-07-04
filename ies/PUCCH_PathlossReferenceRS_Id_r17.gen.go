// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-PathlossReferenceRS-Id-r17 (line 12230).
// PUCCH-PathlossReferenceRS-Id-r17 ::=        INTEGER (0..maxNrofPUCCH-PathlossReferenceRSs-1-r17)

var pUCCHPathlossReferenceRSIdR17Constraints = per.Constrained(0, common.MaxNrofPUCCH_PathlossReferenceRSs_1_r17)

type PUCCH_PathlossReferenceRS_Id_r17 struct {
	Value int64
}

func (v *PUCCH_PathlossReferenceRS_Id_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHPathlossReferenceRSIdR17Constraints)
}

func (v *PUCCH_PathlossReferenceRS_Id_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHPathlossReferenceRSIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
