// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-PathlossReferenceRS-Id-v1610 (line 12615).
// PUSCH-PathlossReferenceRS-Id-v1610 ::= INTEGER (maxNrofPUSCH-PathlossReferenceRSs..maxNrofPUSCH-PathlossReferenceRSs-1-r16)

var pUSCHPathlossReferenceRSIdV1610Constraints = per.Constrained(common.MaxNrofPUSCH_PathlossReferenceRSs, common.MaxNrofPUSCH_PathlossReferenceRSs_1_r16)

type PUSCH_PathlossReferenceRS_Id_v1610 struct {
	Value int64
}

func (v *PUSCH_PathlossReferenceRS_Id_v1610) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUSCHPathlossReferenceRSIdV1610Constraints)
}

func (v *PUSCH_PathlossReferenceRS_Id_v1610) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUSCHPathlossReferenceRSIdV1610Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
