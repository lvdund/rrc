// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PathlossReferenceRS-Id-r17 (line 10936).
// PathlossReferenceRS-Id-r17 ::= INTEGER (0..maxNrofPathlossReferenceRSs-1-r17)

var pathlossReferenceRSIdR17Constraints = per.Constrained(0, common.MaxNrofPathlossReferenceRSs_1_r17)

type PathlossReferenceRS_Id_r17 struct {
	Value int64
}

func (v *PathlossReferenceRS_Id_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pathlossReferenceRSIdR17Constraints)
}

func (v *PathlossReferenceRS_Id_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pathlossReferenceRSIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
