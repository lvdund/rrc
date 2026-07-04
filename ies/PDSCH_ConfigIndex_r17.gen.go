// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDSCH-ConfigIndex-r17 (line 28518).
// PDSCH-ConfigIndex-r17  ::=           INTEGER (0..maxNrofPDSCH-ConfigPTM-1-r17)

var pDSCHConfigIndexR17Constraints = per.Constrained(0, common.MaxNrofPDSCH_ConfigPTM_1_r17)

type PDSCH_ConfigIndex_r17 struct {
	Value int64
}

func (v *PDSCH_ConfigIndex_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pDSCHConfigIndexR17Constraints)
}

func (v *PDSCH_ConfigIndex_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pDSCHConfigIndexR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
