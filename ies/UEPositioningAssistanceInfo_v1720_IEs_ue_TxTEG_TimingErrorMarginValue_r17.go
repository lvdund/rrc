package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc0  aper.Enumerated = 0
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc2  aper.Enumerated = 1
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc4  aper.Enumerated = 2
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc6  aper.Enumerated = 3
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc8  aper.Enumerated = 4
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc12 aper.Enumerated = 5
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc16 aper.Enumerated = 6
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc20 aper.Enumerated = 7
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc24 aper.Enumerated = 8
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc32 aper.Enumerated = 9
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc40 aper.Enumerated = 10
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc48 aper.Enumerated = 11
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc56 aper.Enumerated = 12
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc64 aper.Enumerated = 13
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc72 aper.Enumerated = 14
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc80 aper.Enumerated = 15
)

type UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17", err)
	}
	return nil
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
