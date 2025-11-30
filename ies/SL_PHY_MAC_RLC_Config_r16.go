package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PHY_MAC_RLC_Config_r16 struct {
	Sl_ScheduledConfig_r16         *SL_ScheduledConfig_r16                                `optional,setuprelease`
	Sl_UE_SelectedConfig_r16       *SL_UE_SelectedConfig_r16                              `optional,setuprelease`
	Sl_FreqInfoToReleaseList_r16   []SL_Freq_Id_r16                                       `lb:1,ub:maxNrofFreqSL_r16,optional`
	Sl_FreqInfoToAddModList_r16    []SL_FreqConfig_r16                                    `lb:1,ub:maxNrofFreqSL_r16,optional`
	Sl_RLC_BearerToReleaseList_r16 []SL_RLC_BearerConfigIndex_r16                         `lb:1,ub:maxSL_LCID_r16,optional`
	Sl_RLC_BearerToAddModList_r16  []SL_RLC_BearerConfig_r16                              `lb:1,ub:maxSL_LCID_r16,optional`
	Sl_MaxNumConsecutiveDTX_r16    *SL_PHY_MAC_RLC_Config_r16_sl_MaxNumConsecutiveDTX_r16 `optional`
	Sl_CSI_Acquisition_r16         *SL_PHY_MAC_RLC_Config_r16_sl_CSI_Acquisition_r16      `optional`
	Sl_CSI_SchedulingRequestId_r16 *SchedulingRequestId                                   `optional,setuprelease`
	Sl_SSB_PriorityNR_r16          *int64                                                 `lb:1,ub:8,optional`
	NetworkControlledSyncTx_r16    *SL_PHY_MAC_RLC_Config_r16_networkControlledSyncTx_r16 `optional`
}

func (ie *SL_PHY_MAC_RLC_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ScheduledConfig_r16 != nil, ie.Sl_UE_SelectedConfig_r16 != nil, len(ie.Sl_FreqInfoToReleaseList_r16) > 0, len(ie.Sl_FreqInfoToAddModList_r16) > 0, len(ie.Sl_RLC_BearerToReleaseList_r16) > 0, len(ie.Sl_RLC_BearerToAddModList_r16) > 0, ie.Sl_MaxNumConsecutiveDTX_r16 != nil, ie.Sl_CSI_Acquisition_r16 != nil, ie.Sl_CSI_SchedulingRequestId_r16 != nil, ie.Sl_SSB_PriorityNR_r16 != nil, ie.NetworkControlledSyncTx_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_ScheduledConfig_r16 != nil {
		tmp_Sl_ScheduledConfig_r16 := utils.SetupRelease[*SL_ScheduledConfig_r16]{
			Setup: ie.Sl_ScheduledConfig_r16,
		}
		if err = tmp_Sl_ScheduledConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ScheduledConfig_r16", err)
		}
	}
	if ie.Sl_UE_SelectedConfig_r16 != nil {
		tmp_Sl_UE_SelectedConfig_r16 := utils.SetupRelease[*SL_UE_SelectedConfig_r16]{
			Setup: ie.Sl_UE_SelectedConfig_r16,
		}
		if err = tmp_Sl_UE_SelectedConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_UE_SelectedConfig_r16", err)
		}
	}
	if len(ie.Sl_FreqInfoToReleaseList_r16) > 0 {
		tmp_Sl_FreqInfoToReleaseList_r16 := utils.NewSequence[*SL_Freq_Id_r16]([]*SL_Freq_Id_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.Sl_FreqInfoToReleaseList_r16 {
			tmp_Sl_FreqInfoToReleaseList_r16.Value = append(tmp_Sl_FreqInfoToReleaseList_r16.Value, &i)
		}
		if err = tmp_Sl_FreqInfoToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_FreqInfoToReleaseList_r16", err)
		}
	}
	if len(ie.Sl_FreqInfoToAddModList_r16) > 0 {
		tmp_Sl_FreqInfoToAddModList_r16 := utils.NewSequence[*SL_FreqConfig_r16]([]*SL_FreqConfig_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.Sl_FreqInfoToAddModList_r16 {
			tmp_Sl_FreqInfoToAddModList_r16.Value = append(tmp_Sl_FreqInfoToAddModList_r16.Value, &i)
		}
		if err = tmp_Sl_FreqInfoToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_FreqInfoToAddModList_r16", err)
		}
	}
	if len(ie.Sl_RLC_BearerToReleaseList_r16) > 0 {
		tmp_Sl_RLC_BearerToReleaseList_r16 := utils.NewSequence[*SL_RLC_BearerConfigIndex_r16]([]*SL_RLC_BearerConfigIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_BearerToReleaseList_r16 {
			tmp_Sl_RLC_BearerToReleaseList_r16.Value = append(tmp_Sl_RLC_BearerToReleaseList_r16.Value, &i)
		}
		if err = tmp_Sl_RLC_BearerToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_BearerToReleaseList_r16", err)
		}
	}
	if len(ie.Sl_RLC_BearerToAddModList_r16) > 0 {
		tmp_Sl_RLC_BearerToAddModList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, aper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.Sl_RLC_BearerToAddModList_r16 {
			tmp_Sl_RLC_BearerToAddModList_r16.Value = append(tmp_Sl_RLC_BearerToAddModList_r16.Value, &i)
		}
		if err = tmp_Sl_RLC_BearerToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RLC_BearerToAddModList_r16", err)
		}
	}
	if ie.Sl_MaxNumConsecutiveDTX_r16 != nil {
		if err = ie.Sl_MaxNumConsecutiveDTX_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if ie.Sl_CSI_Acquisition_r16 != nil {
		if err = ie.Sl_CSI_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CSI_Acquisition_r16", err)
		}
	}
	if ie.Sl_CSI_SchedulingRequestId_r16 != nil {
		tmp_Sl_CSI_SchedulingRequestId_r16 := utils.SetupRelease[*SchedulingRequestId]{
			Setup: ie.Sl_CSI_SchedulingRequestId_r16,
		}
		if err = tmp_Sl_CSI_SchedulingRequestId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CSI_SchedulingRequestId_r16", err)
		}
	}
	if ie.Sl_SSB_PriorityNR_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_SSB_PriorityNR_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_SSB_PriorityNR_r16", err)
		}
	}
	if ie.NetworkControlledSyncTx_r16 != nil {
		if err = ie.NetworkControlledSyncTx_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NetworkControlledSyncTx_r16", err)
		}
	}
	return nil
}

func (ie *SL_PHY_MAC_RLC_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_ScheduledConfig_r16Present bool
	if Sl_ScheduledConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_UE_SelectedConfig_r16Present bool
	if Sl_UE_SelectedConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_FreqInfoToReleaseList_r16Present bool
	if Sl_FreqInfoToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_FreqInfoToAddModList_r16Present bool
	if Sl_FreqInfoToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_BearerToReleaseList_r16Present bool
	if Sl_RLC_BearerToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RLC_BearerToAddModList_r16Present bool
	if Sl_RLC_BearerToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MaxNumConsecutiveDTX_r16Present bool
	if Sl_MaxNumConsecutiveDTX_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CSI_Acquisition_r16Present bool
	if Sl_CSI_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CSI_SchedulingRequestId_r16Present bool
	if Sl_CSI_SchedulingRequestId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SSB_PriorityNR_r16Present bool
	if Sl_SSB_PriorityNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NetworkControlledSyncTx_r16Present bool
	if NetworkControlledSyncTx_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_ScheduledConfig_r16Present {
		tmp_Sl_ScheduledConfig_r16 := utils.SetupRelease[*SL_ScheduledConfig_r16]{}
		if err = tmp_Sl_ScheduledConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ScheduledConfig_r16", err)
		}
		ie.Sl_ScheduledConfig_r16 = tmp_Sl_ScheduledConfig_r16.Setup
	}
	if Sl_UE_SelectedConfig_r16Present {
		tmp_Sl_UE_SelectedConfig_r16 := utils.SetupRelease[*SL_UE_SelectedConfig_r16]{}
		if err = tmp_Sl_UE_SelectedConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_UE_SelectedConfig_r16", err)
		}
		ie.Sl_UE_SelectedConfig_r16 = tmp_Sl_UE_SelectedConfig_r16.Setup
	}
	if Sl_FreqInfoToReleaseList_r16Present {
		tmp_Sl_FreqInfoToReleaseList_r16 := utils.NewSequence[*SL_Freq_Id_r16]([]*SL_Freq_Id_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_Sl_FreqInfoToReleaseList_r16 := func() *SL_Freq_Id_r16 {
			return new(SL_Freq_Id_r16)
		}
		if err = tmp_Sl_FreqInfoToReleaseList_r16.Decode(r, fn_Sl_FreqInfoToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Sl_FreqInfoToReleaseList_r16", err)
		}
		ie.Sl_FreqInfoToReleaseList_r16 = []SL_Freq_Id_r16{}
		for _, i := range tmp_Sl_FreqInfoToReleaseList_r16.Value {
			ie.Sl_FreqInfoToReleaseList_r16 = append(ie.Sl_FreqInfoToReleaseList_r16, *i)
		}
	}
	if Sl_FreqInfoToAddModList_r16Present {
		tmp_Sl_FreqInfoToAddModList_r16 := utils.NewSequence[*SL_FreqConfig_r16]([]*SL_FreqConfig_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_Sl_FreqInfoToAddModList_r16 := func() *SL_FreqConfig_r16 {
			return new(SL_FreqConfig_r16)
		}
		if err = tmp_Sl_FreqInfoToAddModList_r16.Decode(r, fn_Sl_FreqInfoToAddModList_r16); err != nil {
			return utils.WrapError("Decode Sl_FreqInfoToAddModList_r16", err)
		}
		ie.Sl_FreqInfoToAddModList_r16 = []SL_FreqConfig_r16{}
		for _, i := range tmp_Sl_FreqInfoToAddModList_r16.Value {
			ie.Sl_FreqInfoToAddModList_r16 = append(ie.Sl_FreqInfoToAddModList_r16, *i)
		}
	}
	if Sl_RLC_BearerToReleaseList_r16Present {
		tmp_Sl_RLC_BearerToReleaseList_r16 := utils.NewSequence[*SL_RLC_BearerConfigIndex_r16]([]*SL_RLC_BearerConfigIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_BearerToReleaseList_r16 := func() *SL_RLC_BearerConfigIndex_r16 {
			return new(SL_RLC_BearerConfigIndex_r16)
		}
		if err = tmp_Sl_RLC_BearerToReleaseList_r16.Decode(r, fn_Sl_RLC_BearerToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Sl_RLC_BearerToReleaseList_r16", err)
		}
		ie.Sl_RLC_BearerToReleaseList_r16 = []SL_RLC_BearerConfigIndex_r16{}
		for _, i := range tmp_Sl_RLC_BearerToReleaseList_r16.Value {
			ie.Sl_RLC_BearerToReleaseList_r16 = append(ie.Sl_RLC_BearerToReleaseList_r16, *i)
		}
	}
	if Sl_RLC_BearerToAddModList_r16Present {
		tmp_Sl_RLC_BearerToAddModList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, aper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_Sl_RLC_BearerToAddModList_r16 := func() *SL_RLC_BearerConfig_r16 {
			return new(SL_RLC_BearerConfig_r16)
		}
		if err = tmp_Sl_RLC_BearerToAddModList_r16.Decode(r, fn_Sl_RLC_BearerToAddModList_r16); err != nil {
			return utils.WrapError("Decode Sl_RLC_BearerToAddModList_r16", err)
		}
		ie.Sl_RLC_BearerToAddModList_r16 = []SL_RLC_BearerConfig_r16{}
		for _, i := range tmp_Sl_RLC_BearerToAddModList_r16.Value {
			ie.Sl_RLC_BearerToAddModList_r16 = append(ie.Sl_RLC_BearerToAddModList_r16, *i)
		}
	}
	if Sl_MaxNumConsecutiveDTX_r16Present {
		ie.Sl_MaxNumConsecutiveDTX_r16 = new(SL_PHY_MAC_RLC_Config_r16_sl_MaxNumConsecutiveDTX_r16)
		if err = ie.Sl_MaxNumConsecutiveDTX_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if Sl_CSI_Acquisition_r16Present {
		ie.Sl_CSI_Acquisition_r16 = new(SL_PHY_MAC_RLC_Config_r16_sl_CSI_Acquisition_r16)
		if err = ie.Sl_CSI_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CSI_Acquisition_r16", err)
		}
	}
	if Sl_CSI_SchedulingRequestId_r16Present {
		tmp_Sl_CSI_SchedulingRequestId_r16 := utils.SetupRelease[*SchedulingRequestId]{}
		if err = tmp_Sl_CSI_SchedulingRequestId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CSI_SchedulingRequestId_r16", err)
		}
		ie.Sl_CSI_SchedulingRequestId_r16 = tmp_Sl_CSI_SchedulingRequestId_r16.Setup
	}
	if Sl_SSB_PriorityNR_r16Present {
		var tmp_int_Sl_SSB_PriorityNR_r16 int64
		if tmp_int_Sl_SSB_PriorityNR_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_SSB_PriorityNR_r16", err)
		}
		ie.Sl_SSB_PriorityNR_r16 = &tmp_int_Sl_SSB_PriorityNR_r16
	}
	if NetworkControlledSyncTx_r16Present {
		ie.NetworkControlledSyncTx_r16 = new(SL_PHY_MAC_RLC_Config_r16_networkControlledSyncTx_r16)
		if err = ie.NetworkControlledSyncTx_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NetworkControlledSyncTx_r16", err)
		}
	}
	return nil
}
