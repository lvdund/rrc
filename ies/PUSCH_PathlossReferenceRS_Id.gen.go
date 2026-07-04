// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-PathlossReferenceRS-Id (line 12613).
// PUSCH-PathlossReferenceRS-Id ::=    INTEGER (0..maxNrofPUSCH-PathlossReferenceRSs-1)

var pUSCHPathlossReferenceRSIdConstraints = per.Constrained(0, common.MaxNrofPUSCH_PathlossReferenceRSs_1)

type PUSCH_PathlossReferenceRS_Id struct {
	Value int64
}

func (v *PUSCH_PathlossReferenceRS_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUSCHPathlossReferenceRSIdConstraints)
}

func (v *PUSCH_PathlossReferenceRS_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUSCHPathlossReferenceRSIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
