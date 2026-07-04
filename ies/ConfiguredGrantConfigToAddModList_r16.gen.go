// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConfiguredGrantConfigToAddModList-r16 (line 5440).
// ConfiguredGrantConfigToAddModList-r16    ::= SEQUENCE (SIZE (1..maxNrofConfiguredGrantConfig-r16)) OF ConfiguredGrantConfig

var configuredGrantConfigToAddModListR16SizeConstraints = per.SizeRange(1, common.MaxNrofConfiguredGrantConfig_r16)

type ConfiguredGrantConfigToAddModList_r16 []ConfiguredGrantConfig

func (ie *ConfiguredGrantConfigToAddModList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(configuredGrantConfigToAddModListR16SizeConstraints)
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

func (ie *ConfiguredGrantConfigToAddModList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(configuredGrantConfigToAddModListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ConfiguredGrantConfigToAddModList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
