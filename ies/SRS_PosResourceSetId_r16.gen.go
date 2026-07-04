// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosResourceSetId-r16 (line 15432).
// SRS-PosResourceSetId-r16 ::=            INTEGER (0..maxNrofSRS-PosResourceSets-1-r16)

var sRSPosResourceSetIdR16Constraints = per.Constrained(0, common.MaxNrofSRS_PosResourceSets_1_r16)

type SRS_PosResourceSetId_r16 struct {
	Value int64
}

func (v *SRS_PosResourceSetId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSPosResourceSetIdR16Constraints)
}

func (v *SRS_PosResourceSetId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSPosResourceSetIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
