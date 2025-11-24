package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1610 struct {
	InDeviceCoexInd_r16                 *UE_NR_Capability_v1610_inDeviceCoexInd_r16                 `optional`
	Dl_DedicatedMessageSegmentation_r16 *UE_NR_Capability_v1610_dl_DedicatedMessageSegmentation_r16 `optional`
	Nrdc_Parameters_v1610               *NRDC_Parameters_v1610                                      `optional`
	PowSav_Parameters_r16               *PowSav_Parameters_r16                                      `optional`
	Fr1_Add_UE_NR_Capabilities_v1610    *UE_NR_CapabilityAddFRX_Mode_v1610                          `optional`
	Fr2_Add_UE_NR_Capabilities_v1610    *UE_NR_CapabilityAddFRX_Mode_v1610                          `optional`
	Bh_RLF_Indication_r16               *UE_NR_Capability_v1610_bh_RLF_Indication_r16               `optional`
	DirectSN_AdditionFirstRRC_IAB_r16   *UE_NR_Capability_v1610_directSN_AdditionFirstRRC_IAB_r16   `optional`
	Bap_Parameters_r16                  *BAP_Parameters_r16                                         `optional`
	ReferenceTimeProvision_r16          *UE_NR_Capability_v1610_referenceTimeProvision_r16          `optional`
	SidelinkParameters_r16              *SidelinkParameters_r16                                     `optional`
	HighSpeedParameters_r16             *HighSpeedParameters_r16                                    `optional`
	Mac_Parameters_v1610                *MAC_Parameters_v1610                                       `optional`
	McgRLF_RecoveryViaSCG_r16           *UE_NR_Capability_v1610_mcgRLF_RecoveryViaSCG_r16           `optional`
	ResumeWithStoredMCG_SCells_r16      *UE_NR_Capability_v1610_resumeWithStoredMCG_SCells_r16      `optional`
	ResumeWithStoredSCG_r16             *UE_NR_Capability_v1610_resumeWithStoredSCG_r16             `optional`
	ResumeWithSCG_Config_r16            *UE_NR_Capability_v1610_resumeWithSCG_Config_r16            `optional`
	Ue_BasedPerfMeas_Parameters_r16     *UE_BasedPerfMeas_Parameters_r16                            `optional`
	Son_Parameters_r16                  *SON_Parameters_r16                                         `optional`
	OnDemandSIB_Connected_r16           *UE_NR_Capability_v1610_onDemandSIB_Connected_r16           `optional`
	NonCriticalExtension                *UE_NR_Capability_v1640                                     `optional`
}

func (ie *UE_NR_Capability_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.InDeviceCoexInd_r16 != nil, ie.Dl_DedicatedMessageSegmentation_r16 != nil, ie.Nrdc_Parameters_v1610 != nil, ie.PowSav_Parameters_r16 != nil, ie.Fr1_Add_UE_NR_Capabilities_v1610 != nil, ie.Fr2_Add_UE_NR_Capabilities_v1610 != nil, ie.Bh_RLF_Indication_r16 != nil, ie.DirectSN_AdditionFirstRRC_IAB_r16 != nil, ie.Bap_Parameters_r16 != nil, ie.ReferenceTimeProvision_r16 != nil, ie.SidelinkParameters_r16 != nil, ie.HighSpeedParameters_r16 != nil, ie.Mac_Parameters_v1610 != nil, ie.McgRLF_RecoveryViaSCG_r16 != nil, ie.ResumeWithStoredMCG_SCells_r16 != nil, ie.ResumeWithStoredSCG_r16 != nil, ie.ResumeWithSCG_Config_r16 != nil, ie.Ue_BasedPerfMeas_Parameters_r16 != nil, ie.Son_Parameters_r16 != nil, ie.OnDemandSIB_Connected_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InDeviceCoexInd_r16 != nil {
		if err = ie.InDeviceCoexInd_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InDeviceCoexInd_r16", err)
		}
	}
	if ie.Dl_DedicatedMessageSegmentation_r16 != nil {
		if err = ie.Dl_DedicatedMessageSegmentation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_DedicatedMessageSegmentation_r16", err)
		}
	}
	if ie.Nrdc_Parameters_v1610 != nil {
		if err = ie.Nrdc_Parameters_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Nrdc_Parameters_v1610", err)
		}
	}
	if ie.PowSav_Parameters_r16 != nil {
		if err = ie.PowSav_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PowSav_Parameters_r16", err)
		}
	}
	if ie.Fr1_Add_UE_NR_Capabilities_v1610 != nil {
		if err = ie.Fr1_Add_UE_NR_Capabilities_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if ie.Fr2_Add_UE_NR_Capabilities_v1610 != nil {
		if err = ie.Fr2_Add_UE_NR_Capabilities_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Fr2_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if ie.Bh_RLF_Indication_r16 != nil {
		if err = ie.Bh_RLF_Indication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bh_RLF_Indication_r16", err)
		}
	}
	if ie.DirectSN_AdditionFirstRRC_IAB_r16 != nil {
		if err = ie.DirectSN_AdditionFirstRRC_IAB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DirectSN_AdditionFirstRRC_IAB_r16", err)
		}
	}
	if ie.Bap_Parameters_r16 != nil {
		if err = ie.Bap_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bap_Parameters_r16", err)
		}
	}
	if ie.ReferenceTimeProvision_r16 != nil {
		if err = ie.ReferenceTimeProvision_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReferenceTimeProvision_r16", err)
		}
	}
	if ie.SidelinkParameters_r16 != nil {
		if err = ie.SidelinkParameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SidelinkParameters_r16", err)
		}
	}
	if ie.HighSpeedParameters_r16 != nil {
		if err = ie.HighSpeedParameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedParameters_r16", err)
		}
	}
	if ie.Mac_Parameters_v1610 != nil {
		if err = ie.Mac_Parameters_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_Parameters_v1610", err)
		}
	}
	if ie.McgRLF_RecoveryViaSCG_r16 != nil {
		if err = ie.McgRLF_RecoveryViaSCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode McgRLF_RecoveryViaSCG_r16", err)
		}
	}
	if ie.ResumeWithStoredMCG_SCells_r16 != nil {
		if err = ie.ResumeWithStoredMCG_SCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResumeWithStoredMCG_SCells_r16", err)
		}
	}
	if ie.ResumeWithStoredSCG_r16 != nil {
		if err = ie.ResumeWithStoredSCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResumeWithStoredSCG_r16", err)
		}
	}
	if ie.ResumeWithSCG_Config_r16 != nil {
		if err = ie.ResumeWithSCG_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResumeWithSCG_Config_r16", err)
		}
	}
	if ie.Ue_BasedPerfMeas_Parameters_r16 != nil {
		if err = ie.Ue_BasedPerfMeas_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_BasedPerfMeas_Parameters_r16", err)
		}
	}
	if ie.Son_Parameters_r16 != nil {
		if err = ie.Son_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Son_Parameters_r16", err)
		}
	}
	if ie.OnDemandSIB_Connected_r16 != nil {
		if err = ie.OnDemandSIB_Connected_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OnDemandSIB_Connected_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1610) Decode(r *uper.UperReader) error {
	var err error
	var InDeviceCoexInd_r16Present bool
	if InDeviceCoexInd_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_DedicatedMessageSegmentation_r16Present bool
	if Dl_DedicatedMessageSegmentation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Nrdc_Parameters_v1610Present bool
	if Nrdc_Parameters_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PowSav_Parameters_r16Present bool
	if PowSav_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_Add_UE_NR_Capabilities_v1610Present bool
	if Fr1_Add_UE_NR_Capabilities_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_Add_UE_NR_Capabilities_v1610Present bool
	if Fr2_Add_UE_NR_Capabilities_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bh_RLF_Indication_r16Present bool
	if Bh_RLF_Indication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DirectSN_AdditionFirstRRC_IAB_r16Present bool
	if DirectSN_AdditionFirstRRC_IAB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bap_Parameters_r16Present bool
	if Bap_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReferenceTimeProvision_r16Present bool
	if ReferenceTimeProvision_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SidelinkParameters_r16Present bool
	if SidelinkParameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedParameters_r16Present bool
	if HighSpeedParameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_Parameters_v1610Present bool
	if Mac_Parameters_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var McgRLF_RecoveryViaSCG_r16Present bool
	if McgRLF_RecoveryViaSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResumeWithStoredMCG_SCells_r16Present bool
	if ResumeWithStoredMCG_SCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResumeWithStoredSCG_r16Present bool
	if ResumeWithStoredSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResumeWithSCG_Config_r16Present bool
	if ResumeWithSCG_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_BasedPerfMeas_Parameters_r16Present bool
	if Ue_BasedPerfMeas_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Son_Parameters_r16Present bool
	if Son_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OnDemandSIB_Connected_r16Present bool
	if OnDemandSIB_Connected_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if InDeviceCoexInd_r16Present {
		ie.InDeviceCoexInd_r16 = new(UE_NR_Capability_v1610_inDeviceCoexInd_r16)
		if err = ie.InDeviceCoexInd_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InDeviceCoexInd_r16", err)
		}
	}
	if Dl_DedicatedMessageSegmentation_r16Present {
		ie.Dl_DedicatedMessageSegmentation_r16 = new(UE_NR_Capability_v1610_dl_DedicatedMessageSegmentation_r16)
		if err = ie.Dl_DedicatedMessageSegmentation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_DedicatedMessageSegmentation_r16", err)
		}
	}
	if Nrdc_Parameters_v1610Present {
		ie.Nrdc_Parameters_v1610 = new(NRDC_Parameters_v1610)
		if err = ie.Nrdc_Parameters_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Nrdc_Parameters_v1610", err)
		}
	}
	if PowSav_Parameters_r16Present {
		ie.PowSav_Parameters_r16 = new(PowSav_Parameters_r16)
		if err = ie.PowSav_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PowSav_Parameters_r16", err)
		}
	}
	if Fr1_Add_UE_NR_Capabilities_v1610Present {
		ie.Fr1_Add_UE_NR_Capabilities_v1610 = new(UE_NR_CapabilityAddFRX_Mode_v1610)
		if err = ie.Fr1_Add_UE_NR_Capabilities_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if Fr2_Add_UE_NR_Capabilities_v1610Present {
		ie.Fr2_Add_UE_NR_Capabilities_v1610 = new(UE_NR_CapabilityAddFRX_Mode_v1610)
		if err = ie.Fr2_Add_UE_NR_Capabilities_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if Bh_RLF_Indication_r16Present {
		ie.Bh_RLF_Indication_r16 = new(UE_NR_Capability_v1610_bh_RLF_Indication_r16)
		if err = ie.Bh_RLF_Indication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bh_RLF_Indication_r16", err)
		}
	}
	if DirectSN_AdditionFirstRRC_IAB_r16Present {
		ie.DirectSN_AdditionFirstRRC_IAB_r16 = new(UE_NR_Capability_v1610_directSN_AdditionFirstRRC_IAB_r16)
		if err = ie.DirectSN_AdditionFirstRRC_IAB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DirectSN_AdditionFirstRRC_IAB_r16", err)
		}
	}
	if Bap_Parameters_r16Present {
		ie.Bap_Parameters_r16 = new(BAP_Parameters_r16)
		if err = ie.Bap_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bap_Parameters_r16", err)
		}
	}
	if ReferenceTimeProvision_r16Present {
		ie.ReferenceTimeProvision_r16 = new(UE_NR_Capability_v1610_referenceTimeProvision_r16)
		if err = ie.ReferenceTimeProvision_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReferenceTimeProvision_r16", err)
		}
	}
	if SidelinkParameters_r16Present {
		ie.SidelinkParameters_r16 = new(SidelinkParameters_r16)
		if err = ie.SidelinkParameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SidelinkParameters_r16", err)
		}
	}
	if HighSpeedParameters_r16Present {
		ie.HighSpeedParameters_r16 = new(HighSpeedParameters_r16)
		if err = ie.HighSpeedParameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedParameters_r16", err)
		}
	}
	if Mac_Parameters_v1610Present {
		ie.Mac_Parameters_v1610 = new(MAC_Parameters_v1610)
		if err = ie.Mac_Parameters_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_Parameters_v1610", err)
		}
	}
	if McgRLF_RecoveryViaSCG_r16Present {
		ie.McgRLF_RecoveryViaSCG_r16 = new(UE_NR_Capability_v1610_mcgRLF_RecoveryViaSCG_r16)
		if err = ie.McgRLF_RecoveryViaSCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode McgRLF_RecoveryViaSCG_r16", err)
		}
	}
	if ResumeWithStoredMCG_SCells_r16Present {
		ie.ResumeWithStoredMCG_SCells_r16 = new(UE_NR_Capability_v1610_resumeWithStoredMCG_SCells_r16)
		if err = ie.ResumeWithStoredMCG_SCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResumeWithStoredMCG_SCells_r16", err)
		}
	}
	if ResumeWithStoredSCG_r16Present {
		ie.ResumeWithStoredSCG_r16 = new(UE_NR_Capability_v1610_resumeWithStoredSCG_r16)
		if err = ie.ResumeWithStoredSCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResumeWithStoredSCG_r16", err)
		}
	}
	if ResumeWithSCG_Config_r16Present {
		ie.ResumeWithSCG_Config_r16 = new(UE_NR_Capability_v1610_resumeWithSCG_Config_r16)
		if err = ie.ResumeWithSCG_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResumeWithSCG_Config_r16", err)
		}
	}
	if Ue_BasedPerfMeas_Parameters_r16Present {
		ie.Ue_BasedPerfMeas_Parameters_r16 = new(UE_BasedPerfMeas_Parameters_r16)
		if err = ie.Ue_BasedPerfMeas_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_BasedPerfMeas_Parameters_r16", err)
		}
	}
	if Son_Parameters_r16Present {
		ie.Son_Parameters_r16 = new(SON_Parameters_r16)
		if err = ie.Son_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Son_Parameters_r16", err)
		}
	}
	if OnDemandSIB_Connected_r16Present {
		ie.OnDemandSIB_Connected_r16 = new(UE_NR_Capability_v1610_onDemandSIB_Connected_r16)
		if err = ie.OnDemandSIB_Connected_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OnDemandSIB_Connected_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1640)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
