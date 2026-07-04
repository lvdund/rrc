// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PriorityTxConfigIndex-r16 (line 26892).

var sLPriorityTxConfigIndexR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PriorityThreshold-r16", Optional: true},
		{Name: "sl-DefaultTxConfigIndex-r16", Optional: true},
		{Name: "sl-CBR-ConfigIndex-r16", Optional: true},
		{Name: "sl-Tx-ConfigIndexList-r16", Optional: true},
	},
}

var sLPriorityTxConfigIndexR16SlPriorityThresholdR16Constraints = per.Constrained(1, 8)

var sLPriorityTxConfigIndexR16SlDefaultTxConfigIndexR16Constraints = per.Constrained(0, common.MaxCBR_Level_1_r16)

var sLPriorityTxConfigIndexR16SlCBRConfigIndexR16Constraints = per.Constrained(0, common.MaxCBR_Config_1_r16)

var sLPriorityTxConfigIndexR16SlTxConfigIndexListR16Constraints = per.SizeRange(1, common.MaxCBR_Level_r16)

type SL_PriorityTxConfigIndex_r16 struct {
	Sl_PriorityThreshold_r16    *int64
	Sl_DefaultTxConfigIndex_r16 *int64
	Sl_CBR_ConfigIndex_r16      *int64
	Sl_Tx_ConfigIndexList_r16   []SL_TxConfigIndex_r16
}

func (ie *SL_PriorityTxConfigIndex_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPriorityTxConfigIndexR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PriorityThreshold_r16 != nil, ie.Sl_DefaultTxConfigIndex_r16 != nil, ie.Sl_CBR_ConfigIndex_r16 != nil, ie.Sl_Tx_ConfigIndexList_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-PriorityThreshold-r16: integer
	{
		if ie.Sl_PriorityThreshold_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityThreshold_r16, sLPriorityTxConfigIndexR16SlPriorityThresholdR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-DefaultTxConfigIndex-r16: integer
	{
		if ie.Sl_DefaultTxConfigIndex_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_DefaultTxConfigIndex_r16, sLPriorityTxConfigIndexR16SlDefaultTxConfigIndexR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-CBR-ConfigIndex-r16: integer
	{
		if ie.Sl_CBR_ConfigIndex_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_CBR_ConfigIndex_r16, sLPriorityTxConfigIndexR16SlCBRConfigIndexR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-Tx-ConfigIndexList-r16: sequence-of
	{
		if ie.Sl_Tx_ConfigIndexList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPriorityTxConfigIndexR16SlTxConfigIndexListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_Tx_ConfigIndexList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_Tx_ConfigIndexList_r16 {
				if err := ie.Sl_Tx_ConfigIndexList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_PriorityTxConfigIndex_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPriorityTxConfigIndexR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PriorityThreshold-r16: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLPriorityTxConfigIndexR16SlPriorityThresholdR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityThreshold_r16 = &v
		}
	}

	// 3. sl-DefaultTxConfigIndex-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLPriorityTxConfigIndexR16SlDefaultTxConfigIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DefaultTxConfigIndex_r16 = &v
		}
	}

	// 4. sl-CBR-ConfigIndex-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLPriorityTxConfigIndexR16SlCBRConfigIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CBR_ConfigIndex_r16 = &v
		}
	}

	// 5. sl-Tx-ConfigIndexList-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLPriorityTxConfigIndexR16SlTxConfigIndexListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_Tx_ConfigIndexList_r16 = make([]SL_TxConfigIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_Tx_ConfigIndexList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
