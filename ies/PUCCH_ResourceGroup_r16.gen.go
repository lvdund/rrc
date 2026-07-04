// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-ResourceGroup-r16 (line 12156).

var pUCCHResourceGroupR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-ResourceGroupId-r16"},
		{Name: "resourcePerGroupList-r16"},
	},
}

var pUCCHResourceGroupR16ResourcePerGroupListR16Constraints = per.SizeRange(1, common.MaxNrofPUCCH_ResourcesPerGroup_r16)

type PUCCH_ResourceGroup_r16 struct {
	Pucch_ResourceGroupId_r16 PUCCH_ResourceGroupId_r16
	ResourcePerGroupList_r16  []PUCCH_ResourceId
}

func (ie *PUCCH_ResourceGroup_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHResourceGroupR16Constraints)
	_ = seq

	// 1. pucch-ResourceGroupId-r16: ref
	{
		if err := ie.Pucch_ResourceGroupId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. resourcePerGroupList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pUCCHResourceGroupR16ResourcePerGroupListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.ResourcePerGroupList_r16))); err != nil {
			return err
		}
		for i := range ie.ResourcePerGroupList_r16 {
			if err := ie.ResourcePerGroupList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_ResourceGroup_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHResourceGroupR16Constraints)
	_ = seq

	// 1. pucch-ResourceGroupId-r16: ref
	{
		if err := ie.Pucch_ResourceGroupId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. resourcePerGroupList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pUCCHResourceGroupR16ResourcePerGroupListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.ResourcePerGroupList_r16 = make([]PUCCH_ResourceId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.ResourcePerGroupList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
