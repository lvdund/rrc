// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DRX-ConfigPTM-Index-r17 (line 28516).
// DRX-ConfigPTM-Index-r17 ::=          INTEGER (0..maxNrofDRX-ConfigPTM-1-r17)

var dRXConfigPTMIndexR17Constraints = per.Constrained(0, common.MaxNrofDRX_ConfigPTM_1_r17)

type DRX_ConfigPTM_Index_r17 struct {
	Value int64
}

func (v *DRX_ConfigPTM_Index_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, dRXConfigPTMIndexR17Constraints)
}

func (v *DRX_ConfigPTM_Index_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(dRXConfigPTMIndexR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
