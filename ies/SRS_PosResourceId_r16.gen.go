// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosResourceId-r16 (line 15653).
// SRS-PosResourceId-r16 ::=               INTEGER (0..maxNrofSRS-PosResources-1-r16)

var sRSPosResourceIdR16Constraints = per.Constrained(0, common.MaxNrofSRS_PosResources_1_r16)

type SRS_PosResourceId_r16 struct {
	Value int64
}

func (v *SRS_PosResourceId_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRSPosResourceIdR16Constraints)
}

func (v *SRS_PosResourceId_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRSPosResourceIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
