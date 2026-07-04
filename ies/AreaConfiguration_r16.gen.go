// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AreaConfiguration-r16 (line 26049).

var areaConfigurationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "areaConfig-r16"},
		{Name: "interFreqTargetList-r16", Optional: true},
	},
}

var areaConfigurationR16InterFreqTargetListR16Constraints = per.SizeRange(1, common.MaxFreq)

type AreaConfiguration_r16 struct {
	AreaConfig_r16          AreaConfig_r16
	InterFreqTargetList_r16 []InterFreqTargetInfo_r16
}

func (ie *AreaConfiguration_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(areaConfigurationR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterFreqTargetList_r16 != nil}); err != nil {
		return err
	}

	// 2. areaConfig-r16: ref
	{
		if err := ie.AreaConfig_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. interFreqTargetList-r16: sequence-of
	{
		if ie.InterFreqTargetList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(areaConfigurationR16InterFreqTargetListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.InterFreqTargetList_r16))); err != nil {
				return err
			}
			for i := range ie.InterFreqTargetList_r16 {
				if err := ie.InterFreqTargetList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *AreaConfiguration_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(areaConfigurationR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. areaConfig-r16: ref
	{
		if err := ie.AreaConfig_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. interFreqTargetList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(areaConfigurationR16InterFreqTargetListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.InterFreqTargetList_r16 = make([]InterFreqTargetInfo_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.InterFreqTargetList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
