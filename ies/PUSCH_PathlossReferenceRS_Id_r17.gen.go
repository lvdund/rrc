// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-PathlossReferenceRS-Id-r17 (line 12617).
// PUSCH-PathlossReferenceRS-Id-r17 ::= INTEGER (0..maxNrofPUSCH-PathlossReferenceRSs-1-r16)

var pUSCHPathlossReferenceRSIdR17Constraints = per.Constrained(0, common.MaxNrofPUSCH_PathlossReferenceRSs_1_r16)

type PUSCH_PathlossReferenceRS_Id_r17 struct {
	Value int64
}

func (v *PUSCH_PathlossReferenceRS_Id_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUSCHPathlossReferenceRSIdR17Constraints)
}

func (v *PUSCH_PathlossReferenceRS_Id_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUSCHPathlossReferenceRSIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
