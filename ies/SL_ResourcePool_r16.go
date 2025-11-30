package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_ResourcePool_r16 struct {
	Sl_PSCCH_Config_r16               *SL_PSCCH_Config_r16                             `optional,setuprelease`
	Sl_PSSCH_Config_r16               *SL_PSSCH_Config_r16                             `optional,setuprelease`
	Sl_PSFCH_Config_r16               *SL_PSFCH_Config_r16                             `optional,setuprelease`
	Sl_SyncAllowed_r16                *SL_SyncAllowed_r16                              `optional`
	Sl_SubchannelSize_r16             *SL_ResourcePool_r16_sl_SubchannelSize_r16       `optional`
	Dummy                             *int64                                           `lb:10,ub:160,optional`
	Sl_StartRB_Subchannel_r16         *int64                                           `lb:0,ub:265,optional`
	Sl_NumSubchannel_r16              *int64                                           `lb:1,ub:27,optional`
	Sl_Additional_MCS_Table_r16       *SL_ResourcePool_r16_sl_Additional_MCS_Table_r16 `optional`
	Sl_ThreshS_RSSI_CBR_r16           *int64                                           `lb:0,ub:45,optional`
	Sl_TimeWindowSizeCBR_r16          *SL_ResourcePool_r16_sl_TimeWindowSizeCBR_r16    `optional`
	Sl_TimeWindowSizeCR_r16           *SL_ResourcePool_r16_sl_TimeWindowSizeCR_r16     `optional`
	Sl_PTRS_Config_r16                *SL_PTRS_Config_r16                              `optional`
	Sl_UE_SelectedConfigRP_r16        *SL_UE_SelectedConfigRP_r16                      `optional`
	Sl_RxParametersNcell_r16          *Sl_RxParametersNcell_r16                        `optional`
	Sl_ZoneConfigMCR_List_r16         []SL_ZoneConfigMCR_r16                           `lb:16,ub:16,optional`
	Sl_FilterCoefficient_r16          *FilterCoefficient                               `optional`
	Sl_RB_Number_r16                  *int64                                           `lb:10,ub:275,optional`
	Sl_PreemptionEnable_r16           *SL_ResourcePool_r16_sl_PreemptionEnable_r16     `optional`
	Sl_PriorityThreshold_UL_URLLC_r16 *int64                                           `lb:1,ub:9,optional`
	Sl_PriorityThreshold_r16          *int64                                           `lb:1,ub:9,optional`
	Sl_X_Overhead_r16                 *SL_ResourcePool_r16_sl_X_Overhead_r16           `optional`
	Sl_PowerControl_r16               *SL_PowerControl_r16                             `optional`
	Sl_TxPercentageList_r16           *SL_TxPercentageList_r16                         `optional`
	Sl_MinMaxMCS_List_r16             *SL_MinMaxMCS_List_r16                           `optional`
	Sl_TimeResource_r16               *aper.BitString                                  `lb:10,ub:160,optional,ext-1`
	Sl_PBPS_CPS_Config_r17            *SL_PBPS_CPS_Config_r17                          `optional,ext-2,setuprelease`
	Sl_InterUE_CoordinationConfig_r17 *SL_InterUE_CoordinationConfig_r17               `optional,ext-2,setuprelease`
}

func (ie *SL_ResourcePool_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sl_TimeResource_r16 != nil || ie.Sl_PBPS_CPS_Config_r17 != nil || ie.Sl_InterUE_CoordinationConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_PSCCH_Config_r16 != nil, ie.Sl_PSSCH_Config_r16 != nil, ie.Sl_PSFCH_Config_r16 != nil, ie.Sl_SyncAllowed_r16 != nil, ie.Sl_SubchannelSize_r16 != nil, ie.Dummy != nil, ie.Sl_StartRB_Subchannel_r16 != nil, ie.Sl_NumSubchannel_r16 != nil, ie.Sl_Additional_MCS_Table_r16 != nil, ie.Sl_ThreshS_RSSI_CBR_r16 != nil, ie.Sl_TimeWindowSizeCBR_r16 != nil, ie.Sl_TimeWindowSizeCR_r16 != nil, ie.Sl_PTRS_Config_r16 != nil, ie.Sl_UE_SelectedConfigRP_r16 != nil, ie.Sl_RxParametersNcell_r16 != nil, len(ie.Sl_ZoneConfigMCR_List_r16) > 0, ie.Sl_FilterCoefficient_r16 != nil, ie.Sl_RB_Number_r16 != nil, ie.Sl_PreemptionEnable_r16 != nil, ie.Sl_PriorityThreshold_UL_URLLC_r16 != nil, ie.Sl_PriorityThreshold_r16 != nil, ie.Sl_X_Overhead_r16 != nil, ie.Sl_PowerControl_r16 != nil, ie.Sl_TxPercentageList_r16 != nil, ie.Sl_MinMaxMCS_List_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_PSCCH_Config_r16 != nil {
		tmp_Sl_PSCCH_Config_r16 := utils.SetupRelease[*SL_PSCCH_Config_r16]{
			Setup: ie.Sl_PSCCH_Config_r16,
		}
		if err = tmp_Sl_PSCCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSCCH_Config_r16", err)
		}
	}
	if ie.Sl_PSSCH_Config_r16 != nil {
		tmp_Sl_PSSCH_Config_r16 := utils.SetupRelease[*SL_PSSCH_Config_r16]{
			Setup: ie.Sl_PSSCH_Config_r16,
		}
		if err = tmp_Sl_PSSCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSSCH_Config_r16", err)
		}
	}
	if ie.Sl_PSFCH_Config_r16 != nil {
		tmp_Sl_PSFCH_Config_r16 := utils.SetupRelease[*SL_PSFCH_Config_r16]{
			Setup: ie.Sl_PSFCH_Config_r16,
		}
		if err = tmp_Sl_PSFCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_Config_r16", err)
		}
	}
	if ie.Sl_SyncAllowed_r16 != nil {
		if err = ie.Sl_SyncAllowed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SyncAllowed_r16", err)
		}
	}
	if ie.Sl_SubchannelSize_r16 != nil {
		if err = ie.Sl_SubchannelSize_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SubchannelSize_r16", err)
		}
	}
	if ie.Dummy != nil {
		if err = w.WriteInteger(*ie.Dummy, &aper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.Sl_StartRB_Subchannel_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_StartRB_Subchannel_r16, &aper.Constraint{Lb: 0, Ub: 265}, false); err != nil {
			return utils.WrapError("Encode Sl_StartRB_Subchannel_r16", err)
		}
	}
	if ie.Sl_NumSubchannel_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_NumSubchannel_r16, &aper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Encode Sl_NumSubchannel_r16", err)
		}
	}
	if ie.Sl_Additional_MCS_Table_r16 != nil {
		if err = ie.Sl_Additional_MCS_Table_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Additional_MCS_Table_r16", err)
		}
	}
	if ie.Sl_ThreshS_RSSI_CBR_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_ThreshS_RSSI_CBR_r16, &aper.Constraint{Lb: 0, Ub: 45}, false); err != nil {
			return utils.WrapError("Encode Sl_ThreshS_RSSI_CBR_r16", err)
		}
	}
	if ie.Sl_TimeWindowSizeCBR_r16 != nil {
		if err = ie.Sl_TimeWindowSizeCBR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TimeWindowSizeCBR_r16", err)
		}
	}
	if ie.Sl_TimeWindowSizeCR_r16 != nil {
		if err = ie.Sl_TimeWindowSizeCR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TimeWindowSizeCR_r16", err)
		}
	}
	if ie.Sl_PTRS_Config_r16 != nil {
		if err = ie.Sl_PTRS_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PTRS_Config_r16", err)
		}
	}
	if ie.Sl_UE_SelectedConfigRP_r16 != nil {
		if err = ie.Sl_UE_SelectedConfigRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_UE_SelectedConfigRP_r16", err)
		}
	}
	if ie.Sl_RxParametersNcell_r16 != nil {
		if err = ie.Sl_RxParametersNcell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RxParametersNcell_r16", err)
		}
	}
	if len(ie.Sl_ZoneConfigMCR_List_r16) > 0 {
		tmp_Sl_ZoneConfigMCR_List_r16 := utils.NewSequence[*SL_ZoneConfigMCR_r16]([]*SL_ZoneConfigMCR_r16{}, aper.Constraint{Lb: 16, Ub: 16}, false)
		for _, i := range ie.Sl_ZoneConfigMCR_List_r16 {
			tmp_Sl_ZoneConfigMCR_List_r16.Value = append(tmp_Sl_ZoneConfigMCR_List_r16.Value, &i)
		}
		if err = tmp_Sl_ZoneConfigMCR_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ZoneConfigMCR_List_r16", err)
		}
	}
	if ie.Sl_FilterCoefficient_r16 != nil {
		if err = ie.Sl_FilterCoefficient_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_FilterCoefficient_r16", err)
		}
	}
	if ie.Sl_RB_Number_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_RB_Number_r16, &aper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Encode Sl_RB_Number_r16", err)
		}
	}
	if ie.Sl_PreemptionEnable_r16 != nil {
		if err = ie.Sl_PreemptionEnable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PreemptionEnable_r16", err)
		}
	}
	if ie.Sl_PriorityThreshold_UL_URLLC_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityThreshold_UL_URLLC_r16, &aper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityThreshold_UL_URLLC_r16", err)
		}
	}
	if ie.Sl_PriorityThreshold_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PriorityThreshold_r16, &aper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode Sl_PriorityThreshold_r16", err)
		}
	}
	if ie.Sl_X_Overhead_r16 != nil {
		if err = ie.Sl_X_Overhead_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_X_Overhead_r16", err)
		}
	}
	if ie.Sl_PowerControl_r16 != nil {
		if err = ie.Sl_PowerControl_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PowerControl_r16", err)
		}
	}
	if ie.Sl_TxPercentageList_r16 != nil {
		if err = ie.Sl_TxPercentageList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TxPercentageList_r16", err)
		}
	}
	if ie.Sl_MinMaxMCS_List_r16 != nil {
		if err = ie.Sl_MinMaxMCS_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MinMaxMCS_List_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Sl_TimeResource_r16 != nil, ie.Sl_PBPS_CPS_Config_r17 != nil || ie.Sl_InterUE_CoordinationConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_ResourcePool_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_TimeResource_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_TimeResource_r16 optional
			if ie.Sl_TimeResource_r16 != nil {
				if err = extWriter.WriteBitString(ie.Sl_TimeResource_r16.Bytes, uint(ie.Sl_TimeResource_r16.NumBits), &aper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
					return utils.WrapError("Encode Sl_TimeResource_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Sl_PBPS_CPS_Config_r17 != nil, ie.Sl_InterUE_CoordinationConfig_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_PBPS_CPS_Config_r17 optional
			if ie.Sl_PBPS_CPS_Config_r17 != nil {
				tmp_Sl_PBPS_CPS_Config_r17 := utils.SetupRelease[*SL_PBPS_CPS_Config_r17]{
					Setup: ie.Sl_PBPS_CPS_Config_r17,
				}
				if err = tmp_Sl_PBPS_CPS_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_PBPS_CPS_Config_r17", err)
				}
			}
			// encode Sl_InterUE_CoordinationConfig_r17 optional
			if ie.Sl_InterUE_CoordinationConfig_r17 != nil {
				tmp_Sl_InterUE_CoordinationConfig_r17 := utils.SetupRelease[*SL_InterUE_CoordinationConfig_r17]{
					Setup: ie.Sl_InterUE_CoordinationConfig_r17,
				}
				if err = tmp_Sl_InterUE_CoordinationConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_InterUE_CoordinationConfig_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SL_ResourcePool_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSCCH_Config_r16Present bool
	if Sl_PSCCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSSCH_Config_r16Present bool
	if Sl_PSSCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSFCH_Config_r16Present bool
	if Sl_PSFCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SyncAllowed_r16Present bool
	if Sl_SyncAllowed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SubchannelSize_r16Present bool
	if Sl_SubchannelSize_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_StartRB_Subchannel_r16Present bool
	if Sl_StartRB_Subchannel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_NumSubchannel_r16Present bool
	if Sl_NumSubchannel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Additional_MCS_Table_r16Present bool
	if Sl_Additional_MCS_Table_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ThreshS_RSSI_CBR_r16Present bool
	if Sl_ThreshS_RSSI_CBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TimeWindowSizeCBR_r16Present bool
	if Sl_TimeWindowSizeCBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TimeWindowSizeCR_r16Present bool
	if Sl_TimeWindowSizeCR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PTRS_Config_r16Present bool
	if Sl_PTRS_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_UE_SelectedConfigRP_r16Present bool
	if Sl_UE_SelectedConfigRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RxParametersNcell_r16Present bool
	if Sl_RxParametersNcell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ZoneConfigMCR_List_r16Present bool
	if Sl_ZoneConfigMCR_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_FilterCoefficient_r16Present bool
	if Sl_FilterCoefficient_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RB_Number_r16Present bool
	if Sl_RB_Number_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PreemptionEnable_r16Present bool
	if Sl_PreemptionEnable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PriorityThreshold_UL_URLLC_r16Present bool
	if Sl_PriorityThreshold_UL_URLLC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PriorityThreshold_r16Present bool
	if Sl_PriorityThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_X_Overhead_r16Present bool
	if Sl_X_Overhead_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PowerControl_r16Present bool
	if Sl_PowerControl_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxPercentageList_r16Present bool
	if Sl_TxPercentageList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MinMaxMCS_List_r16Present bool
	if Sl_MinMaxMCS_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PSCCH_Config_r16Present {
		tmp_Sl_PSCCH_Config_r16 := utils.SetupRelease[*SL_PSCCH_Config_r16]{}
		if err = tmp_Sl_PSCCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PSCCH_Config_r16", err)
		}
		ie.Sl_PSCCH_Config_r16 = tmp_Sl_PSCCH_Config_r16.Setup
	}
	if Sl_PSSCH_Config_r16Present {
		tmp_Sl_PSSCH_Config_r16 := utils.SetupRelease[*SL_PSSCH_Config_r16]{}
		if err = tmp_Sl_PSSCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PSSCH_Config_r16", err)
		}
		ie.Sl_PSSCH_Config_r16 = tmp_Sl_PSSCH_Config_r16.Setup
	}
	if Sl_PSFCH_Config_r16Present {
		tmp_Sl_PSFCH_Config_r16 := utils.SetupRelease[*SL_PSFCH_Config_r16]{}
		if err = tmp_Sl_PSFCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_Config_r16", err)
		}
		ie.Sl_PSFCH_Config_r16 = tmp_Sl_PSFCH_Config_r16.Setup
	}
	if Sl_SyncAllowed_r16Present {
		ie.Sl_SyncAllowed_r16 = new(SL_SyncAllowed_r16)
		if err = ie.Sl_SyncAllowed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SyncAllowed_r16", err)
		}
	}
	if Sl_SubchannelSize_r16Present {
		ie.Sl_SubchannelSize_r16 = new(SL_ResourcePool_r16_sl_SubchannelSize_r16)
		if err = ie.Sl_SubchannelSize_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SubchannelSize_r16", err)
		}
	}
	if DummyPresent {
		var tmp_int_Dummy int64
		if tmp_int_Dummy, err = r.ReadInteger(&aper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
		ie.Dummy = &tmp_int_Dummy
	}
	if Sl_StartRB_Subchannel_r16Present {
		var tmp_int_Sl_StartRB_Subchannel_r16 int64
		if tmp_int_Sl_StartRB_Subchannel_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 265}, false); err != nil {
			return utils.WrapError("Decode Sl_StartRB_Subchannel_r16", err)
		}
		ie.Sl_StartRB_Subchannel_r16 = &tmp_int_Sl_StartRB_Subchannel_r16
	}
	if Sl_NumSubchannel_r16Present {
		var tmp_int_Sl_NumSubchannel_r16 int64
		if tmp_int_Sl_NumSubchannel_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Decode Sl_NumSubchannel_r16", err)
		}
		ie.Sl_NumSubchannel_r16 = &tmp_int_Sl_NumSubchannel_r16
	}
	if Sl_Additional_MCS_Table_r16Present {
		ie.Sl_Additional_MCS_Table_r16 = new(SL_ResourcePool_r16_sl_Additional_MCS_Table_r16)
		if err = ie.Sl_Additional_MCS_Table_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Additional_MCS_Table_r16", err)
		}
	}
	if Sl_ThreshS_RSSI_CBR_r16Present {
		var tmp_int_Sl_ThreshS_RSSI_CBR_r16 int64
		if tmp_int_Sl_ThreshS_RSSI_CBR_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 45}, false); err != nil {
			return utils.WrapError("Decode Sl_ThreshS_RSSI_CBR_r16", err)
		}
		ie.Sl_ThreshS_RSSI_CBR_r16 = &tmp_int_Sl_ThreshS_RSSI_CBR_r16
	}
	if Sl_TimeWindowSizeCBR_r16Present {
		ie.Sl_TimeWindowSizeCBR_r16 = new(SL_ResourcePool_r16_sl_TimeWindowSizeCBR_r16)
		if err = ie.Sl_TimeWindowSizeCBR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TimeWindowSizeCBR_r16", err)
		}
	}
	if Sl_TimeWindowSizeCR_r16Present {
		ie.Sl_TimeWindowSizeCR_r16 = new(SL_ResourcePool_r16_sl_TimeWindowSizeCR_r16)
		if err = ie.Sl_TimeWindowSizeCR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TimeWindowSizeCR_r16", err)
		}
	}
	if Sl_PTRS_Config_r16Present {
		ie.Sl_PTRS_Config_r16 = new(SL_PTRS_Config_r16)
		if err = ie.Sl_PTRS_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PTRS_Config_r16", err)
		}
	}
	if Sl_UE_SelectedConfigRP_r16Present {
		ie.Sl_UE_SelectedConfigRP_r16 = new(SL_UE_SelectedConfigRP_r16)
		if err = ie.Sl_UE_SelectedConfigRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_UE_SelectedConfigRP_r16", err)
		}
	}
	if Sl_RxParametersNcell_r16Present {
		ie.Sl_RxParametersNcell_r16 = new(Sl_RxParametersNcell_r16)
		if err = ie.Sl_RxParametersNcell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RxParametersNcell_r16", err)
		}
	}
	if Sl_ZoneConfigMCR_List_r16Present {
		tmp_Sl_ZoneConfigMCR_List_r16 := utils.NewSequence[*SL_ZoneConfigMCR_r16]([]*SL_ZoneConfigMCR_r16{}, aper.Constraint{Lb: 16, Ub: 16}, false)
		fn_Sl_ZoneConfigMCR_List_r16 := func() *SL_ZoneConfigMCR_r16 {
			return new(SL_ZoneConfigMCR_r16)
		}
		if err = tmp_Sl_ZoneConfigMCR_List_r16.Decode(r, fn_Sl_ZoneConfigMCR_List_r16); err != nil {
			return utils.WrapError("Decode Sl_ZoneConfigMCR_List_r16", err)
		}
		ie.Sl_ZoneConfigMCR_List_r16 = []SL_ZoneConfigMCR_r16{}
		for _, i := range tmp_Sl_ZoneConfigMCR_List_r16.Value {
			ie.Sl_ZoneConfigMCR_List_r16 = append(ie.Sl_ZoneConfigMCR_List_r16, *i)
		}
	}
	if Sl_FilterCoefficient_r16Present {
		ie.Sl_FilterCoefficient_r16 = new(FilterCoefficient)
		if err = ie.Sl_FilterCoefficient_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_FilterCoefficient_r16", err)
		}
	}
	if Sl_RB_Number_r16Present {
		var tmp_int_Sl_RB_Number_r16 int64
		if tmp_int_Sl_RB_Number_r16, err = r.ReadInteger(&aper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Decode Sl_RB_Number_r16", err)
		}
		ie.Sl_RB_Number_r16 = &tmp_int_Sl_RB_Number_r16
	}
	if Sl_PreemptionEnable_r16Present {
		ie.Sl_PreemptionEnable_r16 = new(SL_ResourcePool_r16_sl_PreemptionEnable_r16)
		if err = ie.Sl_PreemptionEnable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PreemptionEnable_r16", err)
		}
	}
	if Sl_PriorityThreshold_UL_URLLC_r16Present {
		var tmp_int_Sl_PriorityThreshold_UL_URLLC_r16 int64
		if tmp_int_Sl_PriorityThreshold_UL_URLLC_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityThreshold_UL_URLLC_r16", err)
		}
		ie.Sl_PriorityThreshold_UL_URLLC_r16 = &tmp_int_Sl_PriorityThreshold_UL_URLLC_r16
	}
	if Sl_PriorityThreshold_r16Present {
		var tmp_int_Sl_PriorityThreshold_r16 int64
		if tmp_int_Sl_PriorityThreshold_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sl_PriorityThreshold_r16", err)
		}
		ie.Sl_PriorityThreshold_r16 = &tmp_int_Sl_PriorityThreshold_r16
	}
	if Sl_X_Overhead_r16Present {
		ie.Sl_X_Overhead_r16 = new(SL_ResourcePool_r16_sl_X_Overhead_r16)
		if err = ie.Sl_X_Overhead_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_X_Overhead_r16", err)
		}
	}
	if Sl_PowerControl_r16Present {
		ie.Sl_PowerControl_r16 = new(SL_PowerControl_r16)
		if err = ie.Sl_PowerControl_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PowerControl_r16", err)
		}
	}
	if Sl_TxPercentageList_r16Present {
		ie.Sl_TxPercentageList_r16 = new(SL_TxPercentageList_r16)
		if err = ie.Sl_TxPercentageList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TxPercentageList_r16", err)
		}
	}
	if Sl_MinMaxMCS_List_r16Present {
		ie.Sl_MinMaxMCS_List_r16 = new(SL_MinMaxMCS_List_r16)
		if err = ie.Sl_MinMaxMCS_List_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MinMaxMCS_List_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sl_TimeResource_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_TimeResource_r16 optional
			if Sl_TimeResource_r16Present {
				var tmp_bs_Sl_TimeResource_r16 []byte
				var n_Sl_TimeResource_r16 uint
				if tmp_bs_Sl_TimeResource_r16, n_Sl_TimeResource_r16, err = extReader.ReadBitString(&aper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
					return utils.WrapError("Decode Sl_TimeResource_r16", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_Sl_TimeResource_r16,
					NumBits: uint64(n_Sl_TimeResource_r16),
				}
				ie.Sl_TimeResource_r16 = &tmp_bitstring
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sl_PBPS_CPS_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_InterUE_CoordinationConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_PBPS_CPS_Config_r17 optional
			if Sl_PBPS_CPS_Config_r17Present {
				tmp_Sl_PBPS_CPS_Config_r17 := utils.SetupRelease[*SL_PBPS_CPS_Config_r17]{}
				if err = tmp_Sl_PBPS_CPS_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_PBPS_CPS_Config_r17", err)
				}
				ie.Sl_PBPS_CPS_Config_r17 = tmp_Sl_PBPS_CPS_Config_r17.Setup
			}
			// decode Sl_InterUE_CoordinationConfig_r17 optional
			if Sl_InterUE_CoordinationConfig_r17Present {
				tmp_Sl_InterUE_CoordinationConfig_r17 := utils.SetupRelease[*SL_InterUE_CoordinationConfig_r17]{}
				if err = tmp_Sl_InterUE_CoordinationConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_InterUE_CoordinationConfig_r17", err)
				}
				ie.Sl_InterUE_CoordinationConfig_r17 = tmp_Sl_InterUE_CoordinationConfig_r17.Setup
			}
		}
	}
	return nil
}
