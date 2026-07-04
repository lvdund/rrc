// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-QoS-FlowIdentity-r16 (line 27744).
// SL-QoS-FlowIdentity-r16 ::=                    INTEGER (1..maxNrofSL-QFIs-r16)

var sLQoSFlowIdentityR16Constraints = per.Constrained(1, common.MaxNrofSL_QFIs_r16)

type SL_QoS_FlowIdentity_r16 struct {
	Value int64
}

func (v *SL_QoS_FlowIdentity_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLQoSFlowIdentityR16Constraints)
}

func (v *SL_QoS_FlowIdentity_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLQoSFlowIdentityR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
