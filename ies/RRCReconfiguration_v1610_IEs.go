package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1610_IEs struct {
	OtherConfig_v1610                   *OtherConfig_v1610                                   `optional`
	Bap_Config_r16                      *BAP_Config_r16                                      `optional,setuprelease`
	Iab_IP_AddressConfigurationList_r16 *IAB_IP_AddressConfigurationList_r16                 `optional`
	ConditionalReconfiguration_r16      *ConditionalReconfiguration_r16                      `optional`
	Daps_SourceRelease_r16              *RRCReconfiguration_v1610_IEs_daps_SourceRelease_r16 `optional`
	T316_r16                            *T316_r16                                            `optional,setuprelease`
	NeedForGapsConfigNR_r16             *NeedForGapsConfigNR_r16                             `optional,setuprelease`
	OnDemandSIB_Request_r16             *OnDemandSIB_Request_r16                             `optional,setuprelease`
	DedicatedPosSysInfoDelivery_r16     *[]byte                                              `optional`
	Sl_ConfigDedicatedNR_r16            *SL_ConfigDedicatedNR_r16                            `optional,setuprelease`
	Sl_ConfigDedicatedEUTRA_Info_r16    *SL_ConfigDedicatedEUTRA_Info_r16                    `optional,setuprelease`
	TargetCellSMTC_SCG_r16              *SSB_MTC                                             `optional`
	NonCriticalExtension                *RRCReconfiguration_v1700_IEs                        `optional`
}

func (ie *RRCReconfiguration_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.OtherConfig_v1610 != nil, ie.Bap_Config_r16 != nil, ie.Iab_IP_AddressConfigurationList_r16 != nil, ie.ConditionalReconfiguration_r16 != nil, ie.Daps_SourceRelease_r16 != nil, ie.T316_r16 != nil, ie.NeedForGapsConfigNR_r16 != nil, ie.OnDemandSIB_Request_r16 != nil, ie.DedicatedPosSysInfoDelivery_r16 != nil, ie.Sl_ConfigDedicatedNR_r16 != nil, ie.Sl_ConfigDedicatedEUTRA_Info_r16 != nil, ie.TargetCellSMTC_SCG_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OtherConfig_v1610 != nil {
		if err = ie.OtherConfig_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode OtherConfig_v1610", err)
		}
	}
	if ie.Bap_Config_r16 != nil {
		tmp_Bap_Config_r16 := utils.SetupRelease[*BAP_Config_r16]{
			Setup: ie.Bap_Config_r16,
		}
		if err = tmp_Bap_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bap_Config_r16", err)
		}
	}
	if ie.Iab_IP_AddressConfigurationList_r16 != nil {
		if err = ie.Iab_IP_AddressConfigurationList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Iab_IP_AddressConfigurationList_r16", err)
		}
	}
	if ie.ConditionalReconfiguration_r16 != nil {
		if err = ie.ConditionalReconfiguration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConditionalReconfiguration_r16", err)
		}
	}
	if ie.Daps_SourceRelease_r16 != nil {
		if err = ie.Daps_SourceRelease_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Daps_SourceRelease_r16", err)
		}
	}
	if ie.T316_r16 != nil {
		tmp_T316_r16 := utils.SetupRelease[*T316_r16]{
			Setup: ie.T316_r16,
		}
		if err = tmp_T316_r16.Encode(w); err != nil {
			return utils.WrapError("Encode T316_r16", err)
		}
	}
	if ie.NeedForGapsConfigNR_r16 != nil {
		tmp_NeedForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{
			Setup: ie.NeedForGapsConfigNR_r16,
		}
		if err = tmp_NeedForGapsConfigNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapsConfigNR_r16", err)
		}
	}
	if ie.OnDemandSIB_Request_r16 != nil {
		tmp_OnDemandSIB_Request_r16 := utils.SetupRelease[*OnDemandSIB_Request_r16]{
			Setup: ie.OnDemandSIB_Request_r16,
		}
		if err = tmp_OnDemandSIB_Request_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OnDemandSIB_Request_r16", err)
		}
	}
	if ie.DedicatedPosSysInfoDelivery_r16 != nil {
		if err = w.WriteOctetString(*ie.DedicatedPosSysInfoDelivery_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode DedicatedPosSysInfoDelivery_r16", err)
		}
	}
	if ie.Sl_ConfigDedicatedNR_r16 != nil {
		tmp_Sl_ConfigDedicatedNR_r16 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{
			Setup: ie.Sl_ConfigDedicatedNR_r16,
		}
		if err = tmp_Sl_ConfigDedicatedNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ConfigDedicatedNR_r16", err)
		}
	}
	if ie.Sl_ConfigDedicatedEUTRA_Info_r16 != nil {
		tmp_Sl_ConfigDedicatedEUTRA_Info_r16 := utils.SetupRelease[*SL_ConfigDedicatedEUTRA_Info_r16]{
			Setup: ie.Sl_ConfigDedicatedEUTRA_Info_r16,
		}
		if err = tmp_Sl_ConfigDedicatedEUTRA_Info_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ConfigDedicatedEUTRA_Info_r16", err)
		}
	}
	if ie.TargetCellSMTC_SCG_r16 != nil {
		if err = ie.TargetCellSMTC_SCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TargetCellSMTC_SCG_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var OtherConfig_v1610Present bool
	if OtherConfig_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bap_Config_r16Present bool
	if Bap_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Iab_IP_AddressConfigurationList_r16Present bool
	if Iab_IP_AddressConfigurationList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConditionalReconfiguration_r16Present bool
	if ConditionalReconfiguration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Daps_SourceRelease_r16Present bool
	if Daps_SourceRelease_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T316_r16Present bool
	if T316_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapsConfigNR_r16Present bool
	if NeedForGapsConfigNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OnDemandSIB_Request_r16Present bool
	if OnDemandSIB_Request_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DedicatedPosSysInfoDelivery_r16Present bool
	if DedicatedPosSysInfoDelivery_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ConfigDedicatedNR_r16Present bool
	if Sl_ConfigDedicatedNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ConfigDedicatedEUTRA_Info_r16Present bool
	if Sl_ConfigDedicatedEUTRA_Info_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TargetCellSMTC_SCG_r16Present bool
	if TargetCellSMTC_SCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OtherConfig_v1610Present {
		ie.OtherConfig_v1610 = new(OtherConfig_v1610)
		if err = ie.OtherConfig_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode OtherConfig_v1610", err)
		}
	}
	if Bap_Config_r16Present {
		tmp_Bap_Config_r16 := utils.SetupRelease[*BAP_Config_r16]{}
		if err = tmp_Bap_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bap_Config_r16", err)
		}
		ie.Bap_Config_r16 = tmp_Bap_Config_r16.Setup
	}
	if Iab_IP_AddressConfigurationList_r16Present {
		ie.Iab_IP_AddressConfigurationList_r16 = new(IAB_IP_AddressConfigurationList_r16)
		if err = ie.Iab_IP_AddressConfigurationList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Iab_IP_AddressConfigurationList_r16", err)
		}
	}
	if ConditionalReconfiguration_r16Present {
		ie.ConditionalReconfiguration_r16 = new(ConditionalReconfiguration_r16)
		if err = ie.ConditionalReconfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConditionalReconfiguration_r16", err)
		}
	}
	if Daps_SourceRelease_r16Present {
		ie.Daps_SourceRelease_r16 = new(RRCReconfiguration_v1610_IEs_daps_SourceRelease_r16)
		if err = ie.Daps_SourceRelease_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Daps_SourceRelease_r16", err)
		}
	}
	if T316_r16Present {
		tmp_T316_r16 := utils.SetupRelease[*T316_r16]{}
		if err = tmp_T316_r16.Decode(r); err != nil {
			return utils.WrapError("Decode T316_r16", err)
		}
		ie.T316_r16 = tmp_T316_r16.Setup
	}
	if NeedForGapsConfigNR_r16Present {
		tmp_NeedForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{}
		if err = tmp_NeedForGapsConfigNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapsConfigNR_r16", err)
		}
		ie.NeedForGapsConfigNR_r16 = tmp_NeedForGapsConfigNR_r16.Setup
	}
	if OnDemandSIB_Request_r16Present {
		tmp_OnDemandSIB_Request_r16 := utils.SetupRelease[*OnDemandSIB_Request_r16]{}
		if err = tmp_OnDemandSIB_Request_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OnDemandSIB_Request_r16", err)
		}
		ie.OnDemandSIB_Request_r16 = tmp_OnDemandSIB_Request_r16.Setup
	}
	if DedicatedPosSysInfoDelivery_r16Present {
		var tmp_os_DedicatedPosSysInfoDelivery_r16 []byte
		if tmp_os_DedicatedPosSysInfoDelivery_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode DedicatedPosSysInfoDelivery_r16", err)
		}
		ie.DedicatedPosSysInfoDelivery_r16 = &tmp_os_DedicatedPosSysInfoDelivery_r16
	}
	if Sl_ConfigDedicatedNR_r16Present {
		tmp_Sl_ConfigDedicatedNR_r16 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{}
		if err = tmp_Sl_ConfigDedicatedNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ConfigDedicatedNR_r16", err)
		}
		ie.Sl_ConfigDedicatedNR_r16 = tmp_Sl_ConfigDedicatedNR_r16.Setup
	}
	if Sl_ConfigDedicatedEUTRA_Info_r16Present {
		tmp_Sl_ConfigDedicatedEUTRA_Info_r16 := utils.SetupRelease[*SL_ConfigDedicatedEUTRA_Info_r16]{}
		if err = tmp_Sl_ConfigDedicatedEUTRA_Info_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ConfigDedicatedEUTRA_Info_r16", err)
		}
		ie.Sl_ConfigDedicatedEUTRA_Info_r16 = tmp_Sl_ConfigDedicatedEUTRA_Info_r16.Setup
	}
	if TargetCellSMTC_SCG_r16Present {
		ie.TargetCellSMTC_SCG_r16 = new(SSB_MTC)
		if err = ie.TargetCellSMTC_SCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TargetCellSMTC_SCG_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfiguration_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
