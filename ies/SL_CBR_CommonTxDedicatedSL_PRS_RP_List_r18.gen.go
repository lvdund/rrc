// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-CBR-CommonTxDedicatedSL-PRS-RP-List-r18 (line 26925).

var sLCBRCommonTxDedicatedSLPRSRPListR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-CBR-RangeDedicatedSL-PRS-RP-List-r18", Optional: true},
		{Name: "sl-CBR-SL-PRS-TxConfigList-r18", Optional: true},
	},
}

var sLCBRCommonTxDedicatedSLPRSRPListR18SlCBRRangeDedicatedSLPRSRPListR18Constraints = per.SizeRange(1, common.MaxCBR_ConfigDedSL_PRS_1_r18)

var sLCBRCommonTxDedicatedSLPRSRPListR18SlCBRSLPRSTxConfigListR18Constraints = per.SizeRange(1, common.MaxNrofSL_PRS_TxConfig_r18)

type SL_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 struct {
	Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18 []SL_CBR_LevelsDedicatedSL_PRS_RP_r18
	Sl_CBR_SL_PRS_TxConfigList_r18          []SL_CBR_SL_PRS_TxConfig_r18
}

func (ie *SL_CBR_CommonTxDedicatedSL_PRS_RP_List_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLCBRCommonTxDedicatedSLPRSRPListR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18 != nil, ie.Sl_CBR_SL_PRS_TxConfigList_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-CBR-RangeDedicatedSL-PRS-RP-List-r18: sequence-of
	{
		if ie.Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLCBRCommonTxDedicatedSLPRSRPListR18SlCBRRangeDedicatedSLPRSRPListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18 {
				if err := ie.Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-CBR-SL-PRS-TxConfigList-r18: sequence-of
	{
		if ie.Sl_CBR_SL_PRS_TxConfigList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLCBRCommonTxDedicatedSLPRSRPListR18SlCBRSLPRSTxConfigListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_CBR_SL_PRS_TxConfigList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_CBR_SL_PRS_TxConfigList_r18 {
				if err := ie.Sl_CBR_SL_PRS_TxConfigList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_CBR_CommonTxDedicatedSL_PRS_RP_List_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLCBRCommonTxDedicatedSLPRSRPListR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-CBR-RangeDedicatedSL-PRS-RP-List-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLCBRCommonTxDedicatedSLPRSRPListR18SlCBRRangeDedicatedSLPRSRPListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18 = make([]SL_CBR_LevelsDedicatedSL_PRS_RP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_CBR_RangeDedicatedSL_PRS_RP_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-CBR-SL-PRS-TxConfigList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLCBRCommonTxDedicatedSLPRSRPListR18SlCBRSLPRSTxConfigListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_CBR_SL_PRS_TxConfigList_r18 = make([]SL_CBR_SL_PRS_TxConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_CBR_SL_PRS_TxConfigList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
