// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IntraBandCC-CombinationReqList-r17 (line 5785).

var intraBandCCCombinationReqListR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellIndexList-r17"},
		{Name: "cc-CombinationList-r17"},
	},
}

var intraBandCCCombinationReqListR17ServCellIndexListR17Constraints = per.SizeRange(1, common.MaxNrofServingCells)

var intraBandCCCombinationReqListR17CcCombinationListR17Constraints = per.SizeRange(1, common.MaxNrofReqComDC_Location_r17)

type IntraBandCC_CombinationReqList_r17 struct {
	ServCellIndexList_r17  []ServCellIndex
	Cc_CombinationList_r17 []IntraBandCC_Combination_r17
}

func (ie *IntraBandCC_CombinationReqList_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(intraBandCCCombinationReqListR17Constraints)
	_ = seq

	// 1. servCellIndexList-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(intraBandCCCombinationReqListR17ServCellIndexListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.ServCellIndexList_r17))); err != nil {
			return err
		}
		for i := range ie.ServCellIndexList_r17 {
			if err := ie.ServCellIndexList_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. cc-CombinationList-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(intraBandCCCombinationReqListR17CcCombinationListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Cc_CombinationList_r17))); err != nil {
			return err
		}
		for i := range ie.Cc_CombinationList_r17 {
			if err := ie.Cc_CombinationList_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IntraBandCC_CombinationReqList_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(intraBandCCCombinationReqListR17Constraints)
	_ = seq

	// 1. servCellIndexList-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(intraBandCCCombinationReqListR17ServCellIndexListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.ServCellIndexList_r17 = make([]ServCellIndex, n)
		for i := int64(0); i < n; i++ {
			if err := ie.ServCellIndexList_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. cc-CombinationList-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(intraBandCCCombinationReqListR17CcCombinationListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Cc_CombinationList_r17 = make([]IntraBandCC_Combination_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Cc_CombinationList_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
