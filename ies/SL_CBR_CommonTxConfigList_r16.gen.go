// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-CBR-CommonTxConfigList-r16 (line 26908).

var sLCBRCommonTxConfigListR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-CBR-RangeConfigList-r16", Optional: true},
		{Name: "sl-CBR-PSSCH-TxConfigList-r16", Optional: true},
	},
}

var sLCBRCommonTxConfigListR16SlCBRRangeConfigListR16Constraints = per.SizeRange(1, common.MaxCBR_Config_r16)

var sLCBRCommonTxConfigListR16SlCBRPSSCHTxConfigListR16Constraints = per.SizeRange(1, common.MaxTxConfig_r16)

type SL_CBR_CommonTxConfigList_r16 struct {
	Sl_CBR_RangeConfigList_r16    []SL_CBR_LevelsConfig_r16
	Sl_CBR_PSSCH_TxConfigList_r16 []SL_CBR_PSSCH_TxConfig_r16
}

func (ie *SL_CBR_CommonTxConfigList_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLCBRCommonTxConfigListR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_CBR_RangeConfigList_r16 != nil, ie.Sl_CBR_PSSCH_TxConfigList_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-CBR-RangeConfigList-r16: sequence-of
	{
		if ie.Sl_CBR_RangeConfigList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLCBRCommonTxConfigListR16SlCBRRangeConfigListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_CBR_RangeConfigList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_CBR_RangeConfigList_r16 {
				if err := ie.Sl_CBR_RangeConfigList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-CBR-PSSCH-TxConfigList-r16: sequence-of
	{
		if ie.Sl_CBR_PSSCH_TxConfigList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLCBRCommonTxConfigListR16SlCBRPSSCHTxConfigListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_CBR_PSSCH_TxConfigList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_CBR_PSSCH_TxConfigList_r16 {
				if err := ie.Sl_CBR_PSSCH_TxConfigList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_CBR_CommonTxConfigList_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLCBRCommonTxConfigListR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-CBR-RangeConfigList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLCBRCommonTxConfigListR16SlCBRRangeConfigListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_CBR_RangeConfigList_r16 = make([]SL_CBR_LevelsConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_CBR_RangeConfigList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-CBR-PSSCH-TxConfigList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLCBRCommonTxConfigListR16SlCBRPSSCHTxConfigListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_CBR_PSSCH_TxConfigList_r16 = make([]SL_CBR_PSSCH_TxConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_CBR_PSSCH_TxConfigList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
