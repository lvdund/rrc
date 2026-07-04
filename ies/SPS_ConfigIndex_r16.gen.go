// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SPS-ConfigIndex-r16 (line 15258).
// SPS-ConfigIndex-r16             ::= INTEGER (0.. maxNrofSPS-Config-1-r16)

var sPSConfigIndexR16Constraints = per.Constrained(0, common.MaxNrofSPS_Config_1_r16)

type SPS_ConfigIndex_r16 struct {
	Value int64
}

func (v *SPS_ConfigIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sPSConfigIndexR16Constraints)
}

func (v *SPS_ConfigIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sPSConfigIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
