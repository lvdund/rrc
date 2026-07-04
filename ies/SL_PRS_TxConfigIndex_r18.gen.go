// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PRS-TxConfigIndex-r18 (line 27692).
// SL-PRS-TxConfigIndex-r18 ::=    INTEGER (0.. maxNrofSL-PRS-TxConfig-r18)

var sLPRSTxConfigIndexR18Constraints = per.Constrained(0, common.MaxNrofSL_PRS_TxConfig_r18)

type SL_PRS_TxConfigIndex_r18 struct {
	Value int64
}

func (v *SL_PRS_TxConfigIndex_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLPRSTxConfigIndexR18Constraints)
}

func (v *SL_PRS_TxConfigIndex_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLPRSTxConfigIndexR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
