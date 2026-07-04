// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ConfiguredGrantConfigList-r16 (line 28215).

var sLConfiguredGrantConfigListR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfiguredGrantConfigToReleaseList-r16", Optional: true},
		{Name: "sl-ConfiguredGrantConfigToAddModList-r16", Optional: true},
	},
}

var sLConfiguredGrantConfigListR16SlConfiguredGrantConfigToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofCG_SL_r16)

var sLConfiguredGrantConfigListR16SlConfiguredGrantConfigToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofCG_SL_r16)

type SL_ConfiguredGrantConfigList_r16 struct {
	Sl_ConfiguredGrantConfigToReleaseList_r16 []SL_ConfigIndexCG_r16
	Sl_ConfiguredGrantConfigToAddModList_r16  []SL_ConfiguredGrantConfig_r16
}

func (ie *SL_ConfiguredGrantConfigList_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfiguredGrantConfigListR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfiguredGrantConfigToReleaseList_r16 != nil, ie.Sl_ConfiguredGrantConfigToAddModList_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-ConfiguredGrantConfigToReleaseList-r16: sequence-of
	{
		if ie.Sl_ConfiguredGrantConfigToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfiguredGrantConfigListR16SlConfiguredGrantConfigToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ConfiguredGrantConfigToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_ConfiguredGrantConfigToReleaseList_r16 {
				if err := ie.Sl_ConfiguredGrantConfigToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-ConfiguredGrantConfigToAddModList-r16: sequence-of
	{
		if ie.Sl_ConfiguredGrantConfigToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfiguredGrantConfigListR16SlConfiguredGrantConfigToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ConfiguredGrantConfigToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_ConfiguredGrantConfigToAddModList_r16 {
				if err := ie.Sl_ConfiguredGrantConfigToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_ConfiguredGrantConfigList_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfiguredGrantConfigListR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ConfiguredGrantConfigToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLConfiguredGrantConfigListR16SlConfiguredGrantConfigToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ConfiguredGrantConfigToReleaseList_r16 = make([]SL_ConfigIndexCG_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ConfiguredGrantConfigToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-ConfiguredGrantConfigToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLConfiguredGrantConfigListR16SlConfiguredGrantConfigToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ConfiguredGrantConfigToAddModList_r16 = make([]SL_ConfiguredGrantConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ConfiguredGrantConfigToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
