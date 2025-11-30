package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1700_IEs struct {
	OtherConfig_v1700                  *OtherConfig_v1700                          `optional`
	Sl_L2RelayUE_Config_r17            *SL_L2RelayUE_Config_r17                    `optional,setuprelease`
	Sl_L2RemoteUE_Config_r17           *SL_L2RemoteUE_Config_r17                   `optional,setuprelease`
	DedicatedPagingDelivery_r17        *[]byte                                     `optional`
	NeedForGapNCSG_ConfigNR_r17        *NeedForGapNCSG_ConfigNR_r17                `optional,setuprelease`
	NeedForGapNCSG_ConfigEUTRA_r17     *NeedForGapNCSG_ConfigEUTRA_r17             `optional,setuprelease`
	Musim_GapConfig_r17                *MUSIM_GapConfig_r17                        `optional,setuprelease`
	Ul_GapFR2_Config_r17               *UL_GapFR2_Config_r17                       `optional,setuprelease`
	Scg_State_r17                      *RRCReconfiguration_v1700_IEs_scg_State_r17 `optional`
	AppLayerMeasConfig_r17             *AppLayerMeasConfig_r17                     `optional`
	Ue_TxTEG_RequestUL_TDOA_Config_r17 *UE_TxTEG_RequestUL_TDOA_Config_r17         `optional,setuprelease`
	NonCriticalExtension               interface{}                                 `optional`
}

func (ie *RRCReconfiguration_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OtherConfig_v1700 != nil, ie.Sl_L2RelayUE_Config_r17 != nil, ie.Sl_L2RemoteUE_Config_r17 != nil, ie.DedicatedPagingDelivery_r17 != nil, ie.NeedForGapNCSG_ConfigNR_r17 != nil, ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil, ie.Musim_GapConfig_r17 != nil, ie.Ul_GapFR2_Config_r17 != nil, ie.Scg_State_r17 != nil, ie.AppLayerMeasConfig_r17 != nil, ie.Ue_TxTEG_RequestUL_TDOA_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OtherConfig_v1700 != nil {
		if err = ie.OtherConfig_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode OtherConfig_v1700", err)
		}
	}
	if ie.Sl_L2RelayUE_Config_r17 != nil {
		tmp_Sl_L2RelayUE_Config_r17 := utils.SetupRelease[*SL_L2RelayUE_Config_r17]{
			Setup: ie.Sl_L2RelayUE_Config_r17,
		}
		if err = tmp_Sl_L2RelayUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_L2RelayUE_Config_r17", err)
		}
	}
	if ie.Sl_L2RemoteUE_Config_r17 != nil {
		tmp_Sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{
			Setup: ie.Sl_L2RemoteUE_Config_r17,
		}
		if err = tmp_Sl_L2RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_L2RemoteUE_Config_r17", err)
		}
	}
	if ie.DedicatedPagingDelivery_r17 != nil {
		if err = w.WriteOctetString(*ie.DedicatedPagingDelivery_r17, nil, false); err != nil {
			return utils.WrapError("Encode DedicatedPagingDelivery_r17", err)
		}
	}
	if ie.NeedForGapNCSG_ConfigNR_r17 != nil {
		tmp_NeedForGapNCSG_ConfigNR_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigNR_r17]{
			Setup: ie.NeedForGapNCSG_ConfigNR_r17,
		}
		if err = tmp_NeedForGapNCSG_ConfigNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapNCSG_ConfigNR_r17", err)
		}
	}
	if ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil {
		tmp_NeedForGapNCSG_ConfigEUTRA_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigEUTRA_r17]{
			Setup: ie.NeedForGapNCSG_ConfigEUTRA_r17,
		}
		if err = tmp_NeedForGapNCSG_ConfigEUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapNCSG_ConfigEUTRA_r17", err)
		}
	}
	if ie.Musim_GapConfig_r17 != nil {
		tmp_Musim_GapConfig_r17 := utils.SetupRelease[*MUSIM_GapConfig_r17]{
			Setup: ie.Musim_GapConfig_r17,
		}
		if err = tmp_Musim_GapConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapConfig_r17", err)
		}
	}
	if ie.Ul_GapFR2_Config_r17 != nil {
		tmp_Ul_GapFR2_Config_r17 := utils.SetupRelease[*UL_GapFR2_Config_r17]{
			Setup: ie.Ul_GapFR2_Config_r17,
		}
		if err = tmp_Ul_GapFR2_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_GapFR2_Config_r17", err)
		}
	}
	if ie.Scg_State_r17 != nil {
		if err = ie.Scg_State_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_State_r17", err)
		}
	}
	if ie.AppLayerMeasConfig_r17 != nil {
		if err = ie.AppLayerMeasConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AppLayerMeasConfig_r17", err)
		}
	}
	if ie.Ue_TxTEG_RequestUL_TDOA_Config_r17 != nil {
		tmp_Ue_TxTEG_RequestUL_TDOA_Config_r17 := utils.SetupRelease[*UE_TxTEG_RequestUL_TDOA_Config_r17]{
			Setup: ie.Ue_TxTEG_RequestUL_TDOA_Config_r17,
		}
		if err = tmp_Ue_TxTEG_RequestUL_TDOA_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_TxTEG_RequestUL_TDOA_Config_r17", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var OtherConfig_v1700Present bool
	if OtherConfig_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_L2RelayUE_Config_r17Present bool
	if Sl_L2RelayUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_L2RemoteUE_Config_r17Present bool
	if Sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DedicatedPagingDelivery_r17Present bool
	if DedicatedPagingDelivery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapNCSG_ConfigNR_r17Present bool
	if NeedForGapNCSG_ConfigNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapNCSG_ConfigEUTRA_r17Present bool
	if NeedForGapNCSG_ConfigEUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_GapConfig_r17Present bool
	if Musim_GapConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_GapFR2_Config_r17Present bool
	if Ul_GapFR2_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_State_r17Present bool
	if Scg_State_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AppLayerMeasConfig_r17Present bool
	if AppLayerMeasConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_TxTEG_RequestUL_TDOA_Config_r17Present bool
	if Ue_TxTEG_RequestUL_TDOA_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OtherConfig_v1700Present {
		ie.OtherConfig_v1700 = new(OtherConfig_v1700)
		if err = ie.OtherConfig_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode OtherConfig_v1700", err)
		}
	}
	if Sl_L2RelayUE_Config_r17Present {
		tmp_Sl_L2RelayUE_Config_r17 := utils.SetupRelease[*SL_L2RelayUE_Config_r17]{}
		if err = tmp_Sl_L2RelayUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_L2RelayUE_Config_r17", err)
		}
		ie.Sl_L2RelayUE_Config_r17 = tmp_Sl_L2RelayUE_Config_r17.Setup
	}
	if Sl_L2RemoteUE_Config_r17Present {
		tmp_Sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{}
		if err = tmp_Sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_L2RemoteUE_Config_r17", err)
		}
		ie.Sl_L2RemoteUE_Config_r17 = tmp_Sl_L2RemoteUE_Config_r17.Setup
	}
	if DedicatedPagingDelivery_r17Present {
		var tmp_os_DedicatedPagingDelivery_r17 []byte
		if tmp_os_DedicatedPagingDelivery_r17, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode DedicatedPagingDelivery_r17", err)
		}
		ie.DedicatedPagingDelivery_r17 = &tmp_os_DedicatedPagingDelivery_r17
	}
	if NeedForGapNCSG_ConfigNR_r17Present {
		tmp_NeedForGapNCSG_ConfigNR_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigNR_r17]{}
		if err = tmp_NeedForGapNCSG_ConfigNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapNCSG_ConfigNR_r17", err)
		}
		ie.NeedForGapNCSG_ConfigNR_r17 = tmp_NeedForGapNCSG_ConfigNR_r17.Setup
	}
	if NeedForGapNCSG_ConfigEUTRA_r17Present {
		tmp_NeedForGapNCSG_ConfigEUTRA_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigEUTRA_r17]{}
		if err = tmp_NeedForGapNCSG_ConfigEUTRA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapNCSG_ConfigEUTRA_r17", err)
		}
		ie.NeedForGapNCSG_ConfigEUTRA_r17 = tmp_NeedForGapNCSG_ConfigEUTRA_r17.Setup
	}
	if Musim_GapConfig_r17Present {
		tmp_Musim_GapConfig_r17 := utils.SetupRelease[*MUSIM_GapConfig_r17]{}
		if err = tmp_Musim_GapConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_GapConfig_r17", err)
		}
		ie.Musim_GapConfig_r17 = tmp_Musim_GapConfig_r17.Setup
	}
	if Ul_GapFR2_Config_r17Present {
		tmp_Ul_GapFR2_Config_r17 := utils.SetupRelease[*UL_GapFR2_Config_r17]{}
		if err = tmp_Ul_GapFR2_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_GapFR2_Config_r17", err)
		}
		ie.Ul_GapFR2_Config_r17 = tmp_Ul_GapFR2_Config_r17.Setup
	}
	if Scg_State_r17Present {
		ie.Scg_State_r17 = new(RRCReconfiguration_v1700_IEs_scg_State_r17)
		if err = ie.Scg_State_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_State_r17", err)
		}
	}
	if AppLayerMeasConfig_r17Present {
		ie.AppLayerMeasConfig_r17 = new(AppLayerMeasConfig_r17)
		if err = ie.AppLayerMeasConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AppLayerMeasConfig_r17", err)
		}
	}
	if Ue_TxTEG_RequestUL_TDOA_Config_r17Present {
		tmp_Ue_TxTEG_RequestUL_TDOA_Config_r17 := utils.SetupRelease[*UE_TxTEG_RequestUL_TDOA_Config_r17]{}
		if err = tmp_Ue_TxTEG_RequestUL_TDOA_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_TxTEG_RequestUL_TDOA_Config_r17", err)
		}
		ie.Ue_TxTEG_RequestUL_TDOA_Config_r17 = tmp_Ue_TxTEG_RequestUL_TDOA_Config_r17.Setup
	}
	return nil
}
