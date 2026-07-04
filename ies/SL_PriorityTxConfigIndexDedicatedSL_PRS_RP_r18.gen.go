// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PriorityTxConfigIndexDedicatedSL-PRS-RP-r18 (line 27684).

var sLPriorityTxConfigIndexDedicatedSLPRSRPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PriorityThresholdDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-DefaultTxConfigIndexDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-CBR-ConfigIndexDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-PRS-TxConfigIndexList-r18", Optional: true},
	},
}

var sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlPriorityThresholdDedicatedSLPRSRPR18Constraints = per.Constrained(1, 8)

var sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlDefaultTxConfigIndexDedicatedSLPRSRPR18Constraints = per.Constrained(0, common.MaxCBR_LevelDedSL_PRS_1_r18)

var sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlCBRConfigIndexDedicatedSLPRSRPR18Constraints = per.Constrained(0, common.MaxCBR_ConfigDedSL_PRS_1_r18)

var sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlPRSTxConfigIndexListR18Constraints = per.SizeRange(1, common.MaxCBR_LevelDedSL_PRS_1_r18)

type SL_PriorityTxConfigIndexDedicatedSL_PRS_RP_r18 struct {
	Sl_PriorityThresholdDedicatedSL_PRS_RP_r18    *int64
	Sl_DefaultTxConfigIndexDedicatedSL_PRS_RP_r18 *int64
	Sl_CBR_ConfigIndexDedicatedSL_PRS_RP_r18      *int64
	Sl_PRS_TxConfigIndexList_r18                  []SL_PRS_TxConfigIndex_r18
}

func (ie *SL_PriorityTxConfigIndexDedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPriorityTxConfigIndexDedicatedSLPRSRPR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PriorityThresholdDedicatedSL_PRS_RP_r18 != nil, ie.Sl_DefaultTxConfigIndexDedicatedSL_PRS_RP_r18 != nil, ie.Sl_CBR_ConfigIndexDedicatedSL_PRS_RP_r18 != nil, ie.Sl_PRS_TxConfigIndexList_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PriorityThresholdDedicatedSL-PRS-RP-r18: integer
	{
		if ie.Sl_PriorityThresholdDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityThresholdDedicatedSL_PRS_RP_r18, sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlPriorityThresholdDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-DefaultTxConfigIndexDedicatedSL-PRS-RP-r18: integer
	{
		if ie.Sl_DefaultTxConfigIndexDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_DefaultTxConfigIndexDedicatedSL_PRS_RP_r18, sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlDefaultTxConfigIndexDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-CBR-ConfigIndexDedicatedSL-PRS-RP-r18: integer
	{
		if ie.Sl_CBR_ConfigIndexDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_CBR_ConfigIndexDedicatedSL_PRS_RP_r18, sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlCBRConfigIndexDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-PRS-TxConfigIndexList-r18: sequence-of
	{
		if ie.Sl_PRS_TxConfigIndexList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlPRSTxConfigIndexListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PRS_TxConfigIndexList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PRS_TxConfigIndexList_r18 {
				if err := ie.Sl_PRS_TxConfigIndexList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_PriorityTxConfigIndexDedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPriorityTxConfigIndexDedicatedSLPRSRPR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PriorityThresholdDedicatedSL-PRS-RP-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlPriorityThresholdDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityThresholdDedicatedSL_PRS_RP_r18 = &v
		}
	}

	// 3. sl-DefaultTxConfigIndexDedicatedSL-PRS-RP-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlDefaultTxConfigIndexDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DefaultTxConfigIndexDedicatedSL_PRS_RP_r18 = &v
		}
	}

	// 4. sl-CBR-ConfigIndexDedicatedSL-PRS-RP-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlCBRConfigIndexDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CBR_ConfigIndexDedicatedSL_PRS_RP_r18 = &v
		}
	}

	// 5. sl-PRS-TxConfigIndexList-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLPriorityTxConfigIndexDedicatedSLPRSRPR18SlPRSTxConfigIndexListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PRS_TxConfigIndexList_r18 = make([]SL_PRS_TxConfigIndex_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PRS_TxConfigIndexList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
