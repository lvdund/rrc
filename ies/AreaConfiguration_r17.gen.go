// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AreaConfiguration-r17 (line 26054).

var areaConfigurationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "areaConfig-r17", Optional: true},
		{Name: "interFreqTargetList-r17", Optional: true},
	},
}

var areaConfigurationR17InterFreqTargetListR17Constraints = per.SizeRange(1, common.MaxFreq)

type AreaConfiguration_r17 struct {
	AreaConfig_r17          *AreaConfig_r16
	InterFreqTargetList_r17 []InterFreqTargetInfo_r16
}

func (ie *AreaConfiguration_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(areaConfigurationR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AreaConfig_r17 != nil, ie.InterFreqTargetList_r17 != nil}); err != nil {
		return err
	}

	// 2. areaConfig-r17: ref
	{
		if ie.AreaConfig_r17 != nil {
			if err := ie.AreaConfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. interFreqTargetList-r17: sequence-of
	{
		if ie.InterFreqTargetList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(areaConfigurationR17InterFreqTargetListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.InterFreqTargetList_r17))); err != nil {
				return err
			}
			for i := range ie.InterFreqTargetList_r17 {
				if err := ie.InterFreqTargetList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *AreaConfiguration_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(areaConfigurationR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. areaConfig-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AreaConfig_r17 = new(AreaConfig_r16)
			if err := ie.AreaConfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. interFreqTargetList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(areaConfigurationR17InterFreqTargetListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.InterFreqTargetList_r17 = make([]InterFreqTargetInfo_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.InterFreqTargetList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
