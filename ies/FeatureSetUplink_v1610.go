package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610 struct {
	Pusch_RepetitionTypeB_r16                        *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16                        `optional`
	Ul_CancellationSelfCarrier_r16                   *FeatureSetUplink_v1610_ul_CancellationSelfCarrier_r16                   `optional`
	Ul_CancellationCrossCarrier_r16                  *FeatureSetUplink_v1610_ul_CancellationCrossCarrier_r16                  `optional`
	Ul_FullPwrMode2_MaxSRS_ResInSet_r16              *FeatureSetUplink_v1610_ul_FullPwrMode2_MaxSRS_ResInSet_r16              `optional`
	CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 *FeatureSetUplink_v1610_cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 `optional`
	CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 *FeatureSetUplink_v1610_cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 `optional`
	SupportedSRS_PosResources_r16                    *SRS_AllPosResources_r16                                                 `optional`
	IntraFreqDAPS_UL_r16                             *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16                             `optional`
	IntraBandFreqSeparationUL_v1620                  *FreqSeparationClassUL_v1620                                             `optional`
	MultiPUCCH_r16                                   *FeatureSetUplink_v1610_multiPUCCH_r16                                   `optional`
	TwoPUCCH_Type1_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type1_r16                               `optional`
	TwoPUCCH_Type2_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type2_r16                               `optional`
	TwoPUCCH_Type3_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type3_r16                               `optional`
	TwoPUCCH_Type4_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type4_r16                               `optional`
	Mux_SR_HARQ_ACK_r16                              *FeatureSetUplink_v1610_mux_SR_HARQ_ACK_r16                              `optional`
	Dummy1                                           *FeatureSetUplink_v1610_dummy1                                           `optional`
	Dummy2                                           *FeatureSetUplink_v1610_dummy2                                           `optional`
	TwoPUCCH_Type5_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type5_r16                               `optional`
	TwoPUCCH_Type6_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type6_r16                               `optional`
	TwoPUCCH_Type7_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type7_r16                               `optional`
	TwoPUCCH_Type8_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type8_r16                               `optional`
	TwoPUCCH_Type9_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type9_r16                               `optional`
	TwoPUCCH_Type10_r16                              *FeatureSetUplink_v1610_twoPUCCH_Type10_r16                              `optional`
	TwoPUCCH_Type11_r16                              *FeatureSetUplink_v1610_twoPUCCH_Type11_r16                              `optional`
	Ul_IntraUE_Mux_r16                               *FeatureSetUplink_v1610_ul_IntraUE_Mux_r16                               `optional`
	Ul_FullPwrMode_r16                               *FeatureSetUplink_v1610_ul_FullPwrMode_r16                               `optional`
	CrossCarrierSchedulingProcessing_DiffSCS_r16     *CrossCarrierSchedulingProcessing_DiffSCS_r16                            `optional`
	Ul_FullPwrMode1_r16                              *FeatureSetUplink_v1610_ul_FullPwrMode1_r16                              `optional`
	Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16    *FeatureSetUplink_v1610_ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16    `optional`
	Ul_FullPwrMode2_TPMIGroup_r16                    *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16                    `lb:2,ub:2,optional`
}

func (ie *FeatureSetUplink_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pusch_RepetitionTypeB_r16 != nil, ie.Ul_CancellationSelfCarrier_r16 != nil, ie.Ul_CancellationCrossCarrier_r16 != nil, ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16 != nil, ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil, ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil, ie.SupportedSRS_PosResources_r16 != nil, ie.IntraFreqDAPS_UL_r16 != nil, ie.IntraBandFreqSeparationUL_v1620 != nil, ie.MultiPUCCH_r16 != nil, ie.TwoPUCCH_Type1_r16 != nil, ie.TwoPUCCH_Type2_r16 != nil, ie.TwoPUCCH_Type3_r16 != nil, ie.TwoPUCCH_Type4_r16 != nil, ie.Mux_SR_HARQ_ACK_r16 != nil, ie.Dummy1 != nil, ie.Dummy2 != nil, ie.TwoPUCCH_Type5_r16 != nil, ie.TwoPUCCH_Type6_r16 != nil, ie.TwoPUCCH_Type7_r16 != nil, ie.TwoPUCCH_Type8_r16 != nil, ie.TwoPUCCH_Type9_r16 != nil, ie.TwoPUCCH_Type10_r16 != nil, ie.TwoPUCCH_Type11_r16 != nil, ie.Ul_IntraUE_Mux_r16 != nil, ie.Ul_FullPwrMode_r16 != nil, ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil, ie.Ul_FullPwrMode1_r16 != nil, ie.Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16 != nil, ie.Ul_FullPwrMode2_TPMIGroup_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pusch_RepetitionTypeB_r16 != nil {
		if err = ie.Pusch_RepetitionTypeB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_RepetitionTypeB_r16", err)
		}
	}
	if ie.Ul_CancellationSelfCarrier_r16 != nil {
		if err = ie.Ul_CancellationSelfCarrier_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_CancellationSelfCarrier_r16", err)
		}
	}
	if ie.Ul_CancellationCrossCarrier_r16 != nil {
		if err = ie.Ul_CancellationCrossCarrier_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_CancellationCrossCarrier_r16", err)
		}
	}
	if ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16 != nil {
		if err = ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FullPwrMode2_MaxSRS_ResInSet_r16", err)
		}
	}
	if ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil {
		if err = ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil {
		if err = ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.SupportedSRS_PosResources_r16 != nil {
		if err = ie.SupportedSRS_PosResources_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedSRS_PosResources_r16", err)
		}
	}
	if ie.IntraFreqDAPS_UL_r16 != nil {
		if err = ie.IntraFreqDAPS_UL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqDAPS_UL_r16", err)
		}
	}
	if ie.IntraBandFreqSeparationUL_v1620 != nil {
		if err = ie.IntraBandFreqSeparationUL_v1620.Encode(w); err != nil {
			return utils.WrapError("Encode IntraBandFreqSeparationUL_v1620", err)
		}
	}
	if ie.MultiPUCCH_r16 != nil {
		if err = ie.MultiPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MultiPUCCH_r16", err)
		}
	}
	if ie.TwoPUCCH_Type1_r16 != nil {
		if err = ie.TwoPUCCH_Type1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type1_r16", err)
		}
	}
	if ie.TwoPUCCH_Type2_r16 != nil {
		if err = ie.TwoPUCCH_Type2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type2_r16", err)
		}
	}
	if ie.TwoPUCCH_Type3_r16 != nil {
		if err = ie.TwoPUCCH_Type3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type3_r16", err)
		}
	}
	if ie.TwoPUCCH_Type4_r16 != nil {
		if err = ie.TwoPUCCH_Type4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type4_r16", err)
		}
	}
	if ie.Mux_SR_HARQ_ACK_r16 != nil {
		if err = ie.Mux_SR_HARQ_ACK_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_SR_HARQ_ACK_r16", err)
		}
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.Dummy2 != nil {
		if err = ie.Dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy2", err)
		}
	}
	if ie.TwoPUCCH_Type5_r16 != nil {
		if err = ie.TwoPUCCH_Type5_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type5_r16", err)
		}
	}
	if ie.TwoPUCCH_Type6_r16 != nil {
		if err = ie.TwoPUCCH_Type6_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type6_r16", err)
		}
	}
	if ie.TwoPUCCH_Type7_r16 != nil {
		if err = ie.TwoPUCCH_Type7_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type7_r16", err)
		}
	}
	if ie.TwoPUCCH_Type8_r16 != nil {
		if err = ie.TwoPUCCH_Type8_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type8_r16", err)
		}
	}
	if ie.TwoPUCCH_Type9_r16 != nil {
		if err = ie.TwoPUCCH_Type9_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type9_r16", err)
		}
	}
	if ie.TwoPUCCH_Type10_r16 != nil {
		if err = ie.TwoPUCCH_Type10_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type10_r16", err)
		}
	}
	if ie.TwoPUCCH_Type11_r16 != nil {
		if err = ie.TwoPUCCH_Type11_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Type11_r16", err)
		}
	}
	if ie.Ul_IntraUE_Mux_r16 != nil {
		if err = ie.Ul_IntraUE_Mux_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_IntraUE_Mux_r16", err)
		}
	}
	if ie.Ul_FullPwrMode_r16 != nil {
		if err = ie.Ul_FullPwrMode_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FullPwrMode_r16", err)
		}
	}
	if ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil {
		if err = ie.CrossCarrierSchedulingProcessing_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if ie.Ul_FullPwrMode1_r16 != nil {
		if err = ie.Ul_FullPwrMode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FullPwrMode1_r16", err)
		}
	}
	if ie.Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16 != nil {
		if err = ie.Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16", err)
		}
	}
	if ie.Ul_FullPwrMode2_TPMIGroup_r16 != nil {
		if err = ie.Ul_FullPwrMode2_TPMIGroup_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_FullPwrMode2_TPMIGroup_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1610) Decode(r *aper.AperReader) error {
	var err error
	var Pusch_RepetitionTypeB_r16Present bool
	if Pusch_RepetitionTypeB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_CancellationSelfCarrier_r16Present bool
	if Ul_CancellationSelfCarrier_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_CancellationCrossCarrier_r16Present bool
	if Ul_CancellationCrossCarrier_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FullPwrMode2_MaxSRS_ResInSet_r16Present bool
	if Ul_FullPwrMode2_MaxSRS_ResInSet_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16Present bool
	if CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16Present bool
	if CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedSRS_PosResources_r16Present bool
	if SupportedSRS_PosResources_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFreqDAPS_UL_r16Present bool
	if IntraFreqDAPS_UL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraBandFreqSeparationUL_v1620Present bool
	if IntraBandFreqSeparationUL_v1620Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiPUCCH_r16Present bool
	if MultiPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type1_r16Present bool
	if TwoPUCCH_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type2_r16Present bool
	if TwoPUCCH_Type2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type3_r16Present bool
	if TwoPUCCH_Type3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type4_r16Present bool
	if TwoPUCCH_Type4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_SR_HARQ_ACK_r16Present bool
	if Mux_SR_HARQ_ACK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy2Present bool
	if Dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type5_r16Present bool
	if TwoPUCCH_Type5_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type6_r16Present bool
	if TwoPUCCH_Type6_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type7_r16Present bool
	if TwoPUCCH_Type7_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type8_r16Present bool
	if TwoPUCCH_Type8_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type9_r16Present bool
	if TwoPUCCH_Type9_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type10_r16Present bool
	if TwoPUCCH_Type10_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_Type11_r16Present bool
	if TwoPUCCH_Type11_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_IntraUE_Mux_r16Present bool
	if Ul_IntraUE_Mux_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FullPwrMode_r16Present bool
	if Ul_FullPwrMode_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingProcessing_DiffSCS_r16Present bool
	if CrossCarrierSchedulingProcessing_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FullPwrMode1_r16Present bool
	if Ul_FullPwrMode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16Present bool
	if Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_FullPwrMode2_TPMIGroup_r16Present bool
	if Ul_FullPwrMode2_TPMIGroup_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pusch_RepetitionTypeB_r16Present {
		ie.Pusch_RepetitionTypeB_r16 = new(FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16)
		if err = ie.Pusch_RepetitionTypeB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_RepetitionTypeB_r16", err)
		}
	}
	if Ul_CancellationSelfCarrier_r16Present {
		ie.Ul_CancellationSelfCarrier_r16 = new(FeatureSetUplink_v1610_ul_CancellationSelfCarrier_r16)
		if err = ie.Ul_CancellationSelfCarrier_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_CancellationSelfCarrier_r16", err)
		}
	}
	if Ul_CancellationCrossCarrier_r16Present {
		ie.Ul_CancellationCrossCarrier_r16 = new(FeatureSetUplink_v1610_ul_CancellationCrossCarrier_r16)
		if err = ie.Ul_CancellationCrossCarrier_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_CancellationCrossCarrier_r16", err)
		}
	}
	if Ul_FullPwrMode2_MaxSRS_ResInSet_r16Present {
		ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_MaxSRS_ResInSet_r16)
		if err = ie.Ul_FullPwrMode2_MaxSRS_ResInSet_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FullPwrMode2_MaxSRS_ResInSet_r16", err)
		}
	}
	if CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16Present {
		ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 = new(FeatureSetUplink_v1610_cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16)
		if err = ie.CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16Present {
		ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 = new(FeatureSetUplink_v1610_cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16)
		if err = ie.CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if SupportedSRS_PosResources_r16Present {
		ie.SupportedSRS_PosResources_r16 = new(SRS_AllPosResources_r16)
		if err = ie.SupportedSRS_PosResources_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedSRS_PosResources_r16", err)
		}
	}
	if IntraFreqDAPS_UL_r16Present {
		ie.IntraFreqDAPS_UL_r16 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16)
		if err = ie.IntraFreqDAPS_UL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqDAPS_UL_r16", err)
		}
	}
	if IntraBandFreqSeparationUL_v1620Present {
		ie.IntraBandFreqSeparationUL_v1620 = new(FreqSeparationClassUL_v1620)
		if err = ie.IntraBandFreqSeparationUL_v1620.Decode(r); err != nil {
			return utils.WrapError("Decode IntraBandFreqSeparationUL_v1620", err)
		}
	}
	if MultiPUCCH_r16Present {
		ie.MultiPUCCH_r16 = new(FeatureSetUplink_v1610_multiPUCCH_r16)
		if err = ie.MultiPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MultiPUCCH_r16", err)
		}
	}
	if TwoPUCCH_Type1_r16Present {
		ie.TwoPUCCH_Type1_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type1_r16)
		if err = ie.TwoPUCCH_Type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type1_r16", err)
		}
	}
	if TwoPUCCH_Type2_r16Present {
		ie.TwoPUCCH_Type2_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type2_r16)
		if err = ie.TwoPUCCH_Type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type2_r16", err)
		}
	}
	if TwoPUCCH_Type3_r16Present {
		ie.TwoPUCCH_Type3_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type3_r16)
		if err = ie.TwoPUCCH_Type3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type3_r16", err)
		}
	}
	if TwoPUCCH_Type4_r16Present {
		ie.TwoPUCCH_Type4_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type4_r16)
		if err = ie.TwoPUCCH_Type4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type4_r16", err)
		}
	}
	if Mux_SR_HARQ_ACK_r16Present {
		ie.Mux_SR_HARQ_ACK_r16 = new(FeatureSetUplink_v1610_mux_SR_HARQ_ACK_r16)
		if err = ie.Mux_SR_HARQ_ACK_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_SR_HARQ_ACK_r16", err)
		}
	}
	if Dummy1Present {
		ie.Dummy1 = new(FeatureSetUplink_v1610_dummy1)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if Dummy2Present {
		ie.Dummy2 = new(FeatureSetUplink_v1610_dummy2)
		if err = ie.Dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy2", err)
		}
	}
	if TwoPUCCH_Type5_r16Present {
		ie.TwoPUCCH_Type5_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type5_r16)
		if err = ie.TwoPUCCH_Type5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type5_r16", err)
		}
	}
	if TwoPUCCH_Type6_r16Present {
		ie.TwoPUCCH_Type6_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type6_r16)
		if err = ie.TwoPUCCH_Type6_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type6_r16", err)
		}
	}
	if TwoPUCCH_Type7_r16Present {
		ie.TwoPUCCH_Type7_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type7_r16)
		if err = ie.TwoPUCCH_Type7_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type7_r16", err)
		}
	}
	if TwoPUCCH_Type8_r16Present {
		ie.TwoPUCCH_Type8_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type8_r16)
		if err = ie.TwoPUCCH_Type8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type8_r16", err)
		}
	}
	if TwoPUCCH_Type9_r16Present {
		ie.TwoPUCCH_Type9_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type9_r16)
		if err = ie.TwoPUCCH_Type9_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type9_r16", err)
		}
	}
	if TwoPUCCH_Type10_r16Present {
		ie.TwoPUCCH_Type10_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type10_r16)
		if err = ie.TwoPUCCH_Type10_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type10_r16", err)
		}
	}
	if TwoPUCCH_Type11_r16Present {
		ie.TwoPUCCH_Type11_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type11_r16)
		if err = ie.TwoPUCCH_Type11_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Type11_r16", err)
		}
	}
	if Ul_IntraUE_Mux_r16Present {
		ie.Ul_IntraUE_Mux_r16 = new(FeatureSetUplink_v1610_ul_IntraUE_Mux_r16)
		if err = ie.Ul_IntraUE_Mux_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_IntraUE_Mux_r16", err)
		}
	}
	if Ul_FullPwrMode_r16Present {
		ie.Ul_FullPwrMode_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode_r16)
		if err = ie.Ul_FullPwrMode_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FullPwrMode_r16", err)
		}
	}
	if CrossCarrierSchedulingProcessing_DiffSCS_r16Present {
		ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 = new(CrossCarrierSchedulingProcessing_DiffSCS_r16)
		if err = ie.CrossCarrierSchedulingProcessing_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if Ul_FullPwrMode1_r16Present {
		ie.Ul_FullPwrMode1_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode1_r16)
		if err = ie.Ul_FullPwrMode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FullPwrMode1_r16", err)
		}
	}
	if Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16Present {
		ie.Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16)
		if err = ie.Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16", err)
		}
	}
	if Ul_FullPwrMode2_TPMIGroup_r16Present {
		ie.Ul_FullPwrMode2_TPMIGroup_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16)
		if err = ie.Ul_FullPwrMode2_TPMIGroup_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_FullPwrMode2_TPMIGroup_r16", err)
		}
	}
	return nil
}
