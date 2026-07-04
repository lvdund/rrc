// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NR-DL-PRS-ResourceID-r17 (line 10734).
// NR-DL-PRS-ResourceID-r17 ::= INTEGER (0..maxNrofPRS-ResourcesPerSet-1-r17)

var nRDLPRSResourceIDR17Constraints = per.Constrained(0, common.MaxNrofPRS_ResourcesPerSet_1_r17)

type NR_DL_PRS_ResourceID_r17 struct {
	Value int64
}

func (v *NR_DL_PRS_ResourceID_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, nRDLPRSResourceIDR17Constraints)
}

func (v *NR_DL_PRS_ResourceID_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(nRDLPRSResourceIDR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
