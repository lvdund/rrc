// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PriorityTxConfigIndex-v1650 (line 26899).

var sLPriorityTxConfigIndexV1650Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MCS-RangeList-r16", Optional: true},
	},
}

var sLPriorityTxConfigIndexV1650SlMCSRangeListR16Constraints = per.SizeRange(1, common.MaxCBR_Level_r16)

type SL_PriorityTxConfigIndex_v1650 struct {
	Sl_MCS_RangeList_r16 []SL_MinMaxMCS_List_r16
}

func (ie *SL_PriorityTxConfigIndex_v1650) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPriorityTxConfigIndexV1650Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_MCS_RangeList_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-MCS-RangeList-r16: sequence-of
	{
		if ie.Sl_MCS_RangeList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPriorityTxConfigIndexV1650SlMCSRangeListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MCS_RangeList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_MCS_RangeList_r16 {
				if err := ie.Sl_MCS_RangeList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_PriorityTxConfigIndex_v1650) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPriorityTxConfigIndexV1650Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-MCS-RangeList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLPriorityTxConfigIndexV1650SlMCSRangeListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MCS_RangeList_r16 = make([]SL_MinMaxMCS_List_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MCS_RangeList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
