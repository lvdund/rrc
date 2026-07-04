// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PRS-ResourcePoolID-r18 (line 26873).
// SL-PRS-ResourcePoolID-r18 ::=     INTEGER (1.. maxNrofSL-PRS-TxPool-r18)

var sLPRSResourcePoolIDR18Constraints = per.Constrained(1, common.MaxNrofSL_PRS_TxPool_r18)

type SL_PRS_ResourcePoolID_r18 struct {
	Value int64
}

func (v *SL_PRS_ResourcePoolID_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLPRSResourcePoolIDR18Constraints)
}

func (v *SL_PRS_ResourcePoolID_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLPRSResourcePoolIDR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
