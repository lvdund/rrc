// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SPS-ConfigDeactivationStateList-r16 (line 5299).
// SPS-ConfigDeactivationStateList-r16 ::= SEQUENCE (SIZE (1..maxNrofSPS-DeactivationState)) OF SPS-ConfigDeactivationState-r16

var sPSConfigDeactivationStateListR16SizeConstraints = per.SizeRange(1, common.MaxNrofSPS_DeactivationState)

type SPS_ConfigDeactivationStateList_r16 []SPS_ConfigDeactivationState_r16

func (ie *SPS_ConfigDeactivationStateList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sPSConfigDeactivationStateListR16SizeConstraints)
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

func (ie *SPS_ConfigDeactivationStateList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sPSConfigDeactivationStateListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SPS_ConfigDeactivationStateList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
