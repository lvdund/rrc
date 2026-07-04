// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConfiguredGrantConfigType2DeactivationStateList-r16 (line 5446).

var configuredGrantConfigType2DeactivationStateListR16SizeConstraints = per.SizeRange(1, common.MaxNrofCG_Type2DeactivationState)

type ConfiguredGrantConfigType2DeactivationStateList_r16 []ConfiguredGrantConfigType2DeactivationState_r16

func (ie *ConfiguredGrantConfigType2DeactivationStateList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(configuredGrantConfigType2DeactivationStateListR16SizeConstraints)
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

func (ie *ConfiguredGrantConfigType2DeactivationStateList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(configuredGrantConfigType2DeactivationStateListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ConfiguredGrantConfigType2DeactivationStateList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
