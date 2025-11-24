package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfigCommonNR_r16 struct {
	Sl_FreqInfoList_r16                []SL_FreqConfigCommon_r16                          `lb:1,ub:maxNrofFreqSL_r16,optional`
	Sl_UE_SelectedConfig_r16           *SL_UE_SelectedConfig_r16                          `optional`
	Sl_NR_AnchorCarrierFreqList_r16    *SL_NR_AnchorCarrierFreqList_r16                   `optional`
	Sl_EUTRA_AnchorCarrierFreqList_r16 *SL_EUTRA_AnchorCarrierFreqList_r16                `optional`
	Sl_RadioBearerConfigList_r16       []SL_RadioBearerConfig_r16                         `lb:1,ub:maxNrofSLRB_r16,optional`
	Sl_RLC_BearerConfigList_r16        []SL_RLC_BearerConfig_r16                          `lb:1,ub:maxSL_LCID_r16,optional`
	Sl_MeasConfigCommon_r16            *SL_MeasConfigCommon_r16                           `optional`
	Sl_CSI_Acquisition_r16             *SL_ConfigCommonNR_r16_sl_CSI_Acquisition_r16      `optional`
	Sl_OffsetDFN_r16                   *int64                                             `lb:1,ub:1000,optional`
	T400_r16                           *SL_ConfigCommonNR_r16_t400_r16                    `optional`
	Sl_MaxNumConsecutiveDTX_r16        *SL_ConfigCommonNR_r16_sl_MaxNumConsecutiveDTX_r16 `optional`
	Sl_SSB_PriorityNR_r16              *int64                                             `lb:1,ub:8,optional`
}

func (ie *SL_ConfigCommonNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_FreqInfoList_r16) > 0, ie.Sl_UE_SelectedConfig_r16 != nil, ie.Sl_NR_AnchorCarrierFreqList_r16 != nil, ie.Sl_EUTRA_AnchorCarrierFreqList_r16 != nil, len(ie.Sl_RadioBearerConfigList_r16) > 0, len(ie.Sl_RLC_BearerConfigList_r16) > 0, ie.Sl_MeasConfigCommon_r16 != nil, ie.Sl_CSI_Acquisition_r16 != nil, ie.Sl_OffsetDFN_r16 != nil, ie.T400_r16 != nil, ie.Sl_MaxNumConsecutiveDTX_r16 != nil, ie.Sl_SSB_PriorityNR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_FreqInfoList_r16) > 0 {
		tmp_Sl_FreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.Sl_FreqInfoList_r16 {
			tmp_Sl_FreqInfoList_r16.Value = append(tmp_Sl_FreqInfoList_r16.Value, &i)
		}
		if err = tmp_Sl_FreqInfoList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_FreqInfoList_r16", err)
		}
	}
	if ie.Sl_UE_SelectedConfig_r16 != nil {
		if err = ie.Sl_UE_SelectedConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_UE_SelectedConfig_r16", err)
		}
	}
	if ie.Sl_NR_AnchorCarrierFreqList_r16 != nil {
		if err = ie.Sl_NR_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_NR_AnchorCarrierFreqList_r16", err)
		}
	}
	if ie.Sl_EUTRA_AnchorCarrierFreqList_r16 != nil {
		if err = ie.Sl_EUTRA_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_EUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if len(ie.Sl_RadioBearerConfigList_r16) > 0 {
		tmp_Sl_RadioBearerConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.Sl_RadioBearerConfigList_r16 {
			tmp_Sl_RadioBearerConfigList_r16.Value = append(tmp_Sl_RadioBearerConfigList_r16.Value, &i)
		}
		if err = tmp_Sl_RadioBearerConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RadioBearerConfigList_r16", err)
		}
	}
	if len(ie.Sl_RLC_BearerConfigList_r16) > 0 {
		tmp_Sl_RLC_BearerConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_BearerConfigList_r16 {
			tmp_Sl_RLC_BearerConfigList_r16.Value = append(tmp_Sl_RLC_BearerConfigList_r16.Value, &i)
		}
		if err = tmp_Sl_RLC_BearerConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_BearerConfigList_r16", err)
		}
	}
	if ie.Sl_MeasConfigCommon_r16 != nil {
		if err = ie.Sl_MeasConfigCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasConfigCommon_r16", err)
		}
	}
	if ie.Sl_CSI_Acquisition_r16 != nil {
		if err = ie.Sl_CSI_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CSI_Acquisition_r16", err)
		}
	}
	if ie.Sl_OffsetDFN_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_OffsetDFN_r16, &uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Encode Sl_OffsetDFN_r16", err)
		}
	}
	if ie.T400_r16 != nil {
		if err = ie.T400_r16.Encode(w); err != nil {
			return utils.WrapError("Encode T400_r16", err)
		}
	}
	if ie.Sl_MaxNumConsecutiveDTX_r16 != nil {
		if err = ie.Sl_MaxNumConsecutiveDTX_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if ie.Sl_SSB_PriorityNR_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_SSB_PriorityNR_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_SSB_PriorityNR_r16", err)
		}
	}
	return nil
}

func (ie *SL_ConfigCommonNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_FreqInfoList_r16Present bool
	if Sl_FreqInfoList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_UE_SelectedConfig_r16Present bool
	if Sl_UE_SelectedConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_NR_AnchorCarrierFreqList_r16Present bool
	if Sl_NR_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_EUTRA_AnchorCarrierFreqList_r16Present bool
	if Sl_EUTRA_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RadioBearerConfigList_r16Present bool
	if Sl_RadioBearerConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_BearerConfigList_r16Present bool
	if Sl_RLC_BearerConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasConfigCommon_r16Present bool
	if Sl_MeasConfigCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CSI_Acquisition_r16Present bool
	if Sl_CSI_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_OffsetDFN_r16Present bool
	if Sl_OffsetDFN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T400_r16Present bool
	if T400_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MaxNumConsecutiveDTX_r16Present bool
	if Sl_MaxNumConsecutiveDTX_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SSB_PriorityNR_r16Present bool
	if Sl_SSB_PriorityNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_FreqInfoList_r16Present {
		tmp_Sl_FreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_Sl_FreqInfoList_r16 := func() *SL_FreqConfigCommon_r16 {
			return new(SL_FreqConfigCommon_r16)
		}
		if err = tmp_Sl_FreqInfoList_r16.Decode(r, fn_Sl_FreqInfoList_r16); err != nil {
			return utils.WrapError("Decode Sl_FreqInfoList_r16", err)
		}
		ie.Sl_FreqInfoList_r16 = []SL_FreqConfigCommon_r16{}
		for _, i := range tmp_Sl_FreqInfoList_r16.Value {
			ie.Sl_FreqInfoList_r16 = append(ie.Sl_FreqInfoList_r16, *i)
		}
	}
	if Sl_UE_SelectedConfig_r16Present {
		ie.Sl_UE_SelectedConfig_r16 = new(SL_UE_SelectedConfig_r16)
		if err = ie.Sl_UE_SelectedConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_UE_SelectedConfig_r16", err)
		}
	}
	if Sl_NR_AnchorCarrierFreqList_r16Present {
		ie.Sl_NR_AnchorCarrierFreqList_r16 = new(SL_NR_AnchorCarrierFreqList_r16)
		if err = ie.Sl_NR_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_NR_AnchorCarrierFreqList_r16", err)
		}
	}
	if Sl_EUTRA_AnchorCarrierFreqList_r16Present {
		ie.Sl_EUTRA_AnchorCarrierFreqList_r16 = new(SL_EUTRA_AnchorCarrierFreqList_r16)
		if err = ie.Sl_EUTRA_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_EUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if Sl_RadioBearerConfigList_r16Present {
		tmp_Sl_RadioBearerConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_Sl_RadioBearerConfigList_r16 := func() *SL_RadioBearerConfig_r16 {
			return new(SL_RadioBearerConfig_r16)
		}
		if err = tmp_Sl_RadioBearerConfigList_r16.Decode(r, fn_Sl_RadioBearerConfigList_r16); err != nil {
			return utils.WrapError("Decode Sl_RadioBearerConfigList_r16", err)
		}
		ie.Sl_RadioBearerConfigList_r16 = []SL_RadioBearerConfig_r16{}
		for _, i := range tmp_Sl_RadioBearerConfigList_r16.Value {
			ie.Sl_RadioBearerConfigList_r16 = append(ie.Sl_RadioBearerConfigList_r16, *i)
		}
	}
	if Sl_RLC_BearerConfigList_r16Present {
		tmp_Sl_RLC_BearerConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_BearerConfigList_r16 := func() *SL_RLC_BearerConfig_r16 {
			return new(SL_RLC_BearerConfig_r16)
		}
		if err = tmp_Sl_RLC_BearerConfigList_r16.Decode(r, fn_Sl_RLC_BearerConfigList_r16); err != nil {
			return utils.WrapError("Decode Sl_RLC_BearerConfigList_r16", err)
		}
		ie.Sl_RLC_BearerConfigList_r16 = []SL_RLC_BearerConfig_r16{}
		for _, i := range tmp_Sl_RLC_BearerConfigList_r16.Value {
			ie.Sl_RLC_BearerConfigList_r16 = append(ie.Sl_RLC_BearerConfigList_r16, *i)
		}
	}
	if Sl_MeasConfigCommon_r16Present {
		ie.Sl_MeasConfigCommon_r16 = new(SL_MeasConfigCommon_r16)
		if err = ie.Sl_MeasConfigCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MeasConfigCommon_r16", err)
		}
	}
	if Sl_CSI_Acquisition_r16Present {
		ie.Sl_CSI_Acquisition_r16 = new(SL_ConfigCommonNR_r16_sl_CSI_Acquisition_r16)
		if err = ie.Sl_CSI_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CSI_Acquisition_r16", err)
		}
	}
	if Sl_OffsetDFN_r16Present {
		var tmp_int_Sl_OffsetDFN_r16 int64
		if tmp_int_Sl_OffsetDFN_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Decode Sl_OffsetDFN_r16", err)
		}
		ie.Sl_OffsetDFN_r16 = &tmp_int_Sl_OffsetDFN_r16
	}
	if T400_r16Present {
		ie.T400_r16 = new(SL_ConfigCommonNR_r16_t400_r16)
		if err = ie.T400_r16.Decode(r); err != nil {
			return utils.WrapError("Decode T400_r16", err)
		}
	}
	if Sl_MaxNumConsecutiveDTX_r16Present {
		ie.Sl_MaxNumConsecutiveDTX_r16 = new(SL_ConfigCommonNR_r16_sl_MaxNumConsecutiveDTX_r16)
		if err = ie.Sl_MaxNumConsecutiveDTX_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if Sl_SSB_PriorityNR_r16Present {
		var tmp_int_Sl_SSB_PriorityNR_r16 int64
		if tmp_int_Sl_SSB_PriorityNR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_SSB_PriorityNR_r16", err)
		}
		ie.Sl_SSB_PriorityNR_r16 = &tmp_int_Sl_SSB_PriorityNR_r16
	}
	return nil
}
