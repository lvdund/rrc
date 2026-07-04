// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConfiguredGrantConfigToReleaseList-r16 (line 5442).
// ConfiguredGrantConfigToReleaseList-r16   ::= SEQUENCE (SIZE (1..maxNrofConfiguredGrantConfig-r16)) OF ConfiguredGrantConfigIndex-r16

var configuredGrantConfigToReleaseListR16SizeConstraints = per.SizeRange(1, common.MaxNrofConfiguredGrantConfig_r16)

type ConfiguredGrantConfigToReleaseList_r16 []ConfiguredGrantConfigIndex_r16

func (ie *ConfiguredGrantConfigToReleaseList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(configuredGrantConfigToReleaseListR16SizeConstraints)
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

func (ie *ConfiguredGrantConfigToReleaseList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(configuredGrantConfigToReleaseListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ConfiguredGrantConfigToReleaseList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
