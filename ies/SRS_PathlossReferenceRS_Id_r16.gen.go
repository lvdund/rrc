// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PathlossReferenceRS-Id-r16 (line 15400).
// SRS-PathlossReferenceRS-Id-r16 ::=          INTEGER (0..maxNrofSRS-PathlossReferenceRS-1-r16)

var sRSPathlossReferenceRSIdR16Constraints = per.Constrained(0, common.MaxNrofSRS_PathlossReferenceRS_1_r16)

type SRS_PathlossReferenceRS_Id_r16 struct {
	Value int64
}

func (v *SRS_PathlossReferenceRS_Id_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSPathlossReferenceRSIdR16Constraints)
}

func (v *SRS_PathlossReferenceRS_Id_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSPathlossReferenceRSIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
