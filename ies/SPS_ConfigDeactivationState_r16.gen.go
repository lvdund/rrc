// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SPS-ConfigDeactivationState-r16 (line 5297).
// SPS-ConfigDeactivationState-r16 ::=     SEQUENCE (SIZE (1..maxNrofSPS-Config-r16)) OF SPS-ConfigIndex-r16

var sPSConfigDeactivationStateR16SizeConstraints = per.SizeRange(1, common.MaxNrofSPS_Config_r16)

type SPS_ConfigDeactivationState_r16 []SPS_ConfigIndex_r16

func (ie *SPS_ConfigDeactivationState_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sPSConfigDeactivationStateR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SPS_ConfigDeactivationState_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sPSConfigDeactivationStateR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SPS_ConfigDeactivationState_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
