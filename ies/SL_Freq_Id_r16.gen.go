// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-Freq-Id-r16 (line 27266).
// SL-Freq-Id-r16 ::=                     INTEGER (1.. maxNrofFreqSL-r16)

var sLFreqIdR16Constraints = per.Constrained(1, common.MaxNrofFreqSL_r16)

type SL_Freq_Id_r16 struct {
	Value int64
}

func (v *SL_Freq_Id_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLFreqIdR16Constraints)
}

func (v *SL_Freq_Id_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLFreqIdR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
