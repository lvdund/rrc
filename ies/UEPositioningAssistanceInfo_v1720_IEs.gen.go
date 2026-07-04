// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEPositioningAssistanceInfo-v1720-IEs (line 3586).

var uEPositioningAssistanceInfoV1720IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-TxTEG-TimingErrorMarginValue-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc0  = 0
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc2  = 1
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc4  = 2
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc6  = 3
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc8  = 4
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc12 = 5
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc16 = 6
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc20 = 7
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc24 = 8
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc32 = 9
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc40 = 10
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc48 = 11
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc56 = 12
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc64 = 13
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc72 = 14
	UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc80 = 15
)

var uEPositioningAssistanceInfoV1720IEsUeTxTEGTimingErrorMarginValueR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc0, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc2, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc4, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc6, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc8, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc12, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc16, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc20, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc24, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc32, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc40, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc48, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc56, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc64, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc72, UEPositioningAssistanceInfo_v1720_IEs_Ue_TxTEG_TimingErrorMarginValue_r17_Tc80},
}

type UEPositioningAssistanceInfo_v1720_IEs struct {
	Ue_TxTEG_TimingErrorMarginValue_r17 *int64
	NonCriticalExtension                *struct{}
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEPositioningAssistanceInfoV1720IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ue_TxTEG_TimingErrorMarginValue_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ue-TxTEG-TimingErrorMarginValue-r17: enumerated
	{
		if ie.Ue_TxTEG_TimingErrorMarginValue_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ue_TxTEG_TimingErrorMarginValue_r17, uEPositioningAssistanceInfoV1720IEsUeTxTEGTimingErrorMarginValueR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEPositioningAssistanceInfoV1720IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ue-TxTEG-TimingErrorMarginValue-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uEPositioningAssistanceInfoV1720IEsUeTxTEGTimingErrorMarginValueR17Constraints)
			if err != nil {
				return err
			}
			ie.Ue_TxTEG_TimingErrorMarginValue_r17 = &idx
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
