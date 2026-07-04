// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-CandidateId-r18 (line 8659).
// LTM-CandidateId-r18 ::=                             INTEGER (1..maxNrofLTM-Configs-r18)

var lTMCandidateIdR18Constraints = per.Constrained(1, common.MaxNrofLTM_Configs_r18)

type LTM_CandidateId_r18 struct {
	Value int64
}

func (v *LTM_CandidateId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, lTMCandidateIdR18Constraints)
}

func (v *LTM_CandidateId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(lTMCandidateIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
