// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-PathlossReferenceRS-Id (line 12226).
// PUCCH-PathlossReferenceRS-Id ::=            INTEGER (0..maxNrofPUCCH-PathlossReferenceRSs-1)

var pUCCHPathlossReferenceRSIdConstraints = per.Constrained(0, common.MaxNrofPUCCH_PathlossReferenceRSs_1)

type PUCCH_PathlossReferenceRS_Id struct {
	Value int64
}

func (v *PUCCH_PathlossReferenceRS_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pUCCHPathlossReferenceRSIdConstraints)
}

func (v *PUCCH_PathlossReferenceRS_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pUCCHPathlossReferenceRSIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
