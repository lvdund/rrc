// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SK-CounterConfiguration-r18 (line 6517).

var sKCounterConfigurationR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sk-CounterConfigToReleaseList-r18", Optional: true},
		{Name: "sk-CounterConfigToAddModList-r18", Optional: true},
	},
}

var sKCounterConfigurationR18SkCounterConfigToReleaseListR18Constraints = per.SizeRange(1, common.MaxSecurityCellSet_r18)

var sKCounterConfigurationR18SkCounterConfigToAddModListR18Constraints = per.SizeRange(1, common.MaxSecurityCellSet_r18)

type SK_CounterConfiguration_r18 struct {
	Sk_CounterConfigToReleaseList_r18 []SecurityCellSetId_r18
	Sk_CounterConfigToAddModList_r18  []SK_CounterConfig_r18
}

func (ie *SK_CounterConfiguration_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sKCounterConfigurationR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sk_CounterConfigToReleaseList_r18 != nil, ie.Sk_CounterConfigToAddModList_r18 != nil}); err != nil {
		return err
	}

	// 2. sk-CounterConfigToReleaseList-r18: sequence-of
	{
		if ie.Sk_CounterConfigToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sKCounterConfigurationR18SkCounterConfigToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sk_CounterConfigToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Sk_CounterConfigToReleaseList_r18 {
				if err := ie.Sk_CounterConfigToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sk-CounterConfigToAddModList-r18: sequence-of
	{
		if ie.Sk_CounterConfigToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sKCounterConfigurationR18SkCounterConfigToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sk_CounterConfigToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Sk_CounterConfigToAddModList_r18 {
				if err := ie.Sk_CounterConfigToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SK_CounterConfiguration_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sKCounterConfigurationR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sk-CounterConfigToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sKCounterConfigurationR18SkCounterConfigToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sk_CounterConfigToReleaseList_r18 = make([]SecurityCellSetId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sk_CounterConfigToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sk-CounterConfigToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sKCounterConfigurationR18SkCounterConfigToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sk_CounterConfigToAddModList_r18 = make([]SK_CounterConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sk_CounterConfigToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
