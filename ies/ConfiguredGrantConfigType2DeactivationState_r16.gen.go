// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConfiguredGrantConfigType2DeactivationState-r16 (line 5444).
// ConfiguredGrantConfigType2DeactivationState-r16 ::= SEQUENCE (SIZE (1..maxNrofConfiguredGrantConfig-r16)) OF ConfiguredGrantConfigIndex-r16

var configuredGrantConfigType2DeactivationStateR16SizeConstraints = per.SizeRange(1, common.MaxNrofConfiguredGrantConfig_r16)

type ConfiguredGrantConfigType2DeactivationState_r16 []ConfiguredGrantConfigIndex_r16

func (ie *ConfiguredGrantConfigType2DeactivationState_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(configuredGrantConfigType2DeactivationStateR16SizeConstraints)
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

func (ie *ConfiguredGrantConfigType2DeactivationState_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(configuredGrantConfigType2DeactivationStateR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ConfiguredGrantConfigType2DeactivationState_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
