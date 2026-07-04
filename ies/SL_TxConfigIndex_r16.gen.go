// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxConfigIndex-r16 (line 26903).
// SL-TxConfigIndex-r16 ::=            INTEGER (0..maxTxConfig-1-r16)

var sLTxConfigIndexR16Constraints = per.Constrained(0, common.MaxTxConfig_1_r16)

type SL_TxConfigIndex_r16 struct {
	Value int64
}

func (v *SL_TxConfigIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLTxConfigIndexR16Constraints)
}

func (v *SL_TxConfigIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLTxConfigIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
