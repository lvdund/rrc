// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDSCH-HARQ-ACK-EnhType3Index-r17 (line 11724).
// PDSCH-HARQ-ACK-EnhType3Index-r17 ::=    INTEGER (0..maxNrofEnhType3HARQ-ACK-1-r17)

var pDSCHHARQACKEnhType3IndexR17Constraints = per.Constrained(0, common.MaxNrofEnhType3HARQ_ACK_1_r17)

type PDSCH_HARQ_ACK_EnhType3Index_r17 struct {
	Value int64
}

func (v *PDSCH_HARQ_ACK_EnhType3Index_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pDSCHHARQACKEnhType3IndexR17Constraints)
}

func (v *PDSCH_HARQ_ACK_EnhType3Index_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pDSCHHARQACKEnhType3IndexR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
