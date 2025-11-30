package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1700 struct {
	InactiveStatePO_Determination_r17       *UE_NR_Capability_v1700_inactiveStatePO_Determination_r17       `optional`
	HighSpeedParameters_v1700               *HighSpeedParameters_v1700                                      `optional`
	PowSav_Parameters_v1700                 *PowSav_Parameters_v1700                                        `optional`
	Mac_Parameters_v1700                    *MAC_Parameters_v1700                                           `optional`
	Ims_Parameters_v1700                    *IMS_Parameters_v1700                                           `optional`
	MeasAndMobParameters_v1700              MeasAndMobParameters_v1700                                      `madatory`
	AppLayerMeasParameters_r17              *AppLayerMeasParameters_r17                                     `optional`
	RedCapParameters_r17                    *RedCapParameters_r17                                           `optional`
	Ra_SDT_r17                              *UE_NR_Capability_v1700_ra_SDT_r17                              `optional`
	Srb_SDT_r17                             *UE_NR_Capability_v1700_srb_SDT_r17                             `optional`
	GNB_SideRTT_BasedPDC_r17                *UE_NR_Capability_v1700_gNB_SideRTT_BasedPDC_r17                `optional`
	Bh_RLF_DetectionRecovery_Indication_r17 *UE_NR_Capability_v1700_bh_RLF_DetectionRecovery_Indication_r17 `optional`
	Nrdc_Parameters_v1700                   *NRDC_Parameters_v1700                                          `optional`
	Bap_Parameters_v1700                    *BAP_Parameters_v1700                                           `optional`
	Musim_GapPreference_r17                 *UE_NR_Capability_v1700_musim_GapPreference_r17                 `optional`
	MusimLeaveConnected_r17                 *UE_NR_Capability_v1700_musimLeaveConnected_r17                 `optional`
	Mbs_Parameters_r17                      MBS_Parameters_r17                                              `madatory`
	NonTerrestrialNetwork_r17               *UE_NR_Capability_v1700_nonTerrestrialNetwork_r17               `optional`
	Ntn_ScenarioSupport_r17                 *UE_NR_Capability_v1700_ntn_ScenarioSupport_r17                 `optional`
	SliceInfoforCellReselection_r17         *UE_NR_Capability_v1700_sliceInfoforCellReselection_r17         `optional`
	Ue_RadioPagingInfo_r17                  *UE_RadioPagingInfo_r17                                         `optional`
	Ul_GapFR2_Pattern_r17                   *aper.BitString                                                 `lb:4,ub:4,optional`
	Ntn_Parameters_r17                      *NTN_Parameters_r17                                             `optional`
	NonCriticalExtension                    interface{}                                                     `optional`
}

func (ie *UE_NR_Capability_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.InactiveStatePO_Determination_r17 != nil, ie.HighSpeedParameters_v1700 != nil, ie.PowSav_Parameters_v1700 != nil, ie.Mac_Parameters_v1700 != nil, ie.Ims_Parameters_v1700 != nil, ie.AppLayerMeasParameters_r17 != nil, ie.RedCapParameters_r17 != nil, ie.Ra_SDT_r17 != nil, ie.Srb_SDT_r17 != nil, ie.GNB_SideRTT_BasedPDC_r17 != nil, ie.Bh_RLF_DetectionRecovery_Indication_r17 != nil, ie.Nrdc_Parameters_v1700 != nil, ie.Bap_Parameters_v1700 != nil, ie.Musim_GapPreference_r17 != nil, ie.MusimLeaveConnected_r17 != nil, ie.NonTerrestrialNetwork_r17 != nil, ie.Ntn_ScenarioSupport_r17 != nil, ie.SliceInfoforCellReselection_r17 != nil, ie.Ue_RadioPagingInfo_r17 != nil, ie.Ul_GapFR2_Pattern_r17 != nil, ie.Ntn_Parameters_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InactiveStatePO_Determination_r17 != nil {
		if err = ie.InactiveStatePO_Determination_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InactiveStatePO_Determination_r17", err)
		}
	}
	if ie.HighSpeedParameters_v1700 != nil {
		if err = ie.HighSpeedParameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedParameters_v1700", err)
		}
	}
	if ie.PowSav_Parameters_v1700 != nil {
		if err = ie.PowSav_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode PowSav_Parameters_v1700", err)
		}
	}
	if ie.Mac_Parameters_v1700 != nil {
		if err = ie.Mac_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_Parameters_v1700", err)
		}
	}
	if ie.Ims_Parameters_v1700 != nil {
		if err = ie.Ims_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Ims_Parameters_v1700", err)
		}
	}
	if ie.AppLayerMeasParameters_r17 != nil {
		if err = ie.AppLayerMeasParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AppLayerMeasParameters_r17", err)
		}
	}
	if ie.RedCapParameters_r17 != nil {
		if err = ie.RedCapParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RedCapParameters_r17", err)
		}
	}
	if ie.Ra_SDT_r17 != nil {
		if err = ie.Ra_SDT_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_SDT_r17", err)
		}
	}
	if ie.Srb_SDT_r17 != nil {
		if err = ie.Srb_SDT_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srb_SDT_r17", err)
		}
	}
	if ie.GNB_SideRTT_BasedPDC_r17 != nil {
		if err = ie.GNB_SideRTT_BasedPDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode GNB_SideRTT_BasedPDC_r17", err)
		}
	}
	if ie.Bh_RLF_DetectionRecovery_Indication_r17 != nil {
		if err = ie.Bh_RLF_DetectionRecovery_Indication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Bh_RLF_DetectionRecovery_Indication_r17", err)
		}
	}
	if ie.Nrdc_Parameters_v1700 != nil {
		if err = ie.Nrdc_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Nrdc_Parameters_v1700", err)
		}
	}
	if ie.Bap_Parameters_v1700 != nil {
		if err = ie.Bap_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Bap_Parameters_v1700", err)
		}
	}
	if ie.Musim_GapPreference_r17 != nil {
		if err = ie.Musim_GapPreference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapPreference_r17", err)
		}
	}
	if ie.MusimLeaveConnected_r17 != nil {
		if err = ie.MusimLeaveConnected_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MusimLeaveConnected_r17", err)
		}
	}
	if ie.NonTerrestrialNetwork_r17 != nil {
		if err = ie.NonTerrestrialNetwork_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NonTerrestrialNetwork_r17", err)
		}
	}
	if ie.Ntn_ScenarioSupport_r17 != nil {
		if err = ie.Ntn_ScenarioSupport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_ScenarioSupport_r17", err)
		}
	}
	if ie.SliceInfoforCellReselection_r17 != nil {
		if err = ie.SliceInfoforCellReselection_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SliceInfoforCellReselection_r17", err)
		}
	}
	if ie.Ue_RadioPagingInfo_r17 != nil {
		if err = ie.Ue_RadioPagingInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_RadioPagingInfo_r17", err)
		}
	}
	if ie.Ul_GapFR2_Pattern_r17 != nil {
		if err = w.WriteBitString(ie.Ul_GapFR2_Pattern_r17.Bytes, uint(ie.Ul_GapFR2_Pattern_r17.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode Ul_GapFR2_Pattern_r17", err)
		}
	}
	if ie.Ntn_Parameters_r17 != nil {
		if err = ie.Ntn_Parameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_Parameters_r17", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1700) Decode(r *aper.AperReader) error {
	var err error
	var InactiveStatePO_Determination_r17Present bool
	if InactiveStatePO_Determination_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedParameters_v1700Present bool
	if HighSpeedParameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PowSav_Parameters_v1700Present bool
	if PowSav_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_Parameters_v1700Present bool
	if Mac_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ims_Parameters_v1700Present bool
	if Ims_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AppLayerMeasParameters_r17Present bool
	if AppLayerMeasParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RedCapParameters_r17Present bool
	if RedCapParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_SDT_r17Present bool
	if Ra_SDT_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srb_SDT_r17Present bool
	if Srb_SDT_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GNB_SideRTT_BasedPDC_r17Present bool
	if GNB_SideRTT_BasedPDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bh_RLF_DetectionRecovery_Indication_r17Present bool
	if Bh_RLF_DetectionRecovery_Indication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Nrdc_Parameters_v1700Present bool
	if Nrdc_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bap_Parameters_v1700Present bool
	if Bap_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_GapPreference_r17Present bool
	if Musim_GapPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MusimLeaveConnected_r17Present bool
	if MusimLeaveConnected_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonTerrestrialNetwork_r17Present bool
	if NonTerrestrialNetwork_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_ScenarioSupport_r17Present bool
	if Ntn_ScenarioSupport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SliceInfoforCellReselection_r17Present bool
	if SliceInfoforCellReselection_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_RadioPagingInfo_r17Present bool
	if Ue_RadioPagingInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_GapFR2_Pattern_r17Present bool
	if Ul_GapFR2_Pattern_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_Parameters_r17Present bool
	if Ntn_Parameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if InactiveStatePO_Determination_r17Present {
		ie.InactiveStatePO_Determination_r17 = new(UE_NR_Capability_v1700_inactiveStatePO_Determination_r17)
		if err = ie.InactiveStatePO_Determination_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InactiveStatePO_Determination_r17", err)
		}
	}
	if HighSpeedParameters_v1700Present {
		ie.HighSpeedParameters_v1700 = new(HighSpeedParameters_v1700)
		if err = ie.HighSpeedParameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedParameters_v1700", err)
		}
	}
	if PowSav_Parameters_v1700Present {
		ie.PowSav_Parameters_v1700 = new(PowSav_Parameters_v1700)
		if err = ie.PowSav_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode PowSav_Parameters_v1700", err)
		}
	}
	if Mac_Parameters_v1700Present {
		ie.Mac_Parameters_v1700 = new(MAC_Parameters_v1700)
		if err = ie.Mac_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_Parameters_v1700", err)
		}
	}
	if Ims_Parameters_v1700Present {
		ie.Ims_Parameters_v1700 = new(IMS_Parameters_v1700)
		if err = ie.Ims_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Ims_Parameters_v1700", err)
		}
	}
	if AppLayerMeasParameters_r17Present {
		ie.AppLayerMeasParameters_r17 = new(AppLayerMeasParameters_r17)
		if err = ie.AppLayerMeasParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AppLayerMeasParameters_r17", err)
		}
	}
	if RedCapParameters_r17Present {
		ie.RedCapParameters_r17 = new(RedCapParameters_r17)
		if err = ie.RedCapParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RedCapParameters_r17", err)
		}
	}
	if Ra_SDT_r17Present {
		ie.Ra_SDT_r17 = new(UE_NR_Capability_v1700_ra_SDT_r17)
		if err = ie.Ra_SDT_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_SDT_r17", err)
		}
	}
	if Srb_SDT_r17Present {
		ie.Srb_SDT_r17 = new(UE_NR_Capability_v1700_srb_SDT_r17)
		if err = ie.Srb_SDT_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Srb_SDT_r17", err)
		}
	}
	if GNB_SideRTT_BasedPDC_r17Present {
		ie.GNB_SideRTT_BasedPDC_r17 = new(UE_NR_Capability_v1700_gNB_SideRTT_BasedPDC_r17)
		if err = ie.GNB_SideRTT_BasedPDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode GNB_SideRTT_BasedPDC_r17", err)
		}
	}
	if Bh_RLF_DetectionRecovery_Indication_r17Present {
		ie.Bh_RLF_DetectionRecovery_Indication_r17 = new(UE_NR_Capability_v1700_bh_RLF_DetectionRecovery_Indication_r17)
		if err = ie.Bh_RLF_DetectionRecovery_Indication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Bh_RLF_DetectionRecovery_Indication_r17", err)
		}
	}
	if Nrdc_Parameters_v1700Present {
		ie.Nrdc_Parameters_v1700 = new(NRDC_Parameters_v1700)
		if err = ie.Nrdc_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Nrdc_Parameters_v1700", err)
		}
	}
	if Bap_Parameters_v1700Present {
		ie.Bap_Parameters_v1700 = new(BAP_Parameters_v1700)
		if err = ie.Bap_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Bap_Parameters_v1700", err)
		}
	}
	if Musim_GapPreference_r17Present {
		ie.Musim_GapPreference_r17 = new(UE_NR_Capability_v1700_musim_GapPreference_r17)
		if err = ie.Musim_GapPreference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_GapPreference_r17", err)
		}
	}
	if MusimLeaveConnected_r17Present {
		ie.MusimLeaveConnected_r17 = new(UE_NR_Capability_v1700_musimLeaveConnected_r17)
		if err = ie.MusimLeaveConnected_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MusimLeaveConnected_r17", err)
		}
	}
	if NonTerrestrialNetwork_r17Present {
		ie.NonTerrestrialNetwork_r17 = new(UE_NR_Capability_v1700_nonTerrestrialNetwork_r17)
		if err = ie.NonTerrestrialNetwork_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NonTerrestrialNetwork_r17", err)
		}
	}
	if Ntn_ScenarioSupport_r17Present {
		ie.Ntn_ScenarioSupport_r17 = new(UE_NR_Capability_v1700_ntn_ScenarioSupport_r17)
		if err = ie.Ntn_ScenarioSupport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_ScenarioSupport_r17", err)
		}
	}
	if SliceInfoforCellReselection_r17Present {
		ie.SliceInfoforCellReselection_r17 = new(UE_NR_Capability_v1700_sliceInfoforCellReselection_r17)
		if err = ie.SliceInfoforCellReselection_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SliceInfoforCellReselection_r17", err)
		}
	}
	if Ue_RadioPagingInfo_r17Present {
		ie.Ue_RadioPagingInfo_r17 = new(UE_RadioPagingInfo_r17)
		if err = ie.Ue_RadioPagingInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_RadioPagingInfo_r17", err)
		}
	}
	if Ul_GapFR2_Pattern_r17Present {
		var tmp_bs_Ul_GapFR2_Pattern_r17 []byte
		var n_Ul_GapFR2_Pattern_r17 uint
		if tmp_bs_Ul_GapFR2_Pattern_r17, n_Ul_GapFR2_Pattern_r17, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Ul_GapFR2_Pattern_r17", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Ul_GapFR2_Pattern_r17,
			NumBits: uint64(n_Ul_GapFR2_Pattern_r17),
		}
		ie.Ul_GapFR2_Pattern_r17 = &tmp_bitstring
	}
	if Ntn_Parameters_r17Present {
		ie.Ntn_Parameters_r17 = new(NTN_Parameters_r17)
		if err = ie.Ntn_Parameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_Parameters_r17", err)
		}
	}
	return nil
}
