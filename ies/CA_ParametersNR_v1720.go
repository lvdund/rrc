package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1720 struct {
	ParallelTxSRS_PUCCH_PUSCH_intraBand_r17           *CA_ParametersNR_v1720_parallelTxSRS_PUCCH_PUSCH_intraBand_r17           `optional`
	ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17     *CA_ParametersNR_v1720_parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17     `optional`
	SemiStaticPUCCH_CellSwitchSingleGroup_r17         *CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17         `optional`
	SemiStaticPUCCH_CellSwitchTwoGroups_r17           []TwoPUCCH_Grp_Configurations_r17                                        `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r17,optional`
	DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17  *CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17  `optional`
	DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17  *CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17  `optional`
	DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17    []TwoPUCCH_Grp_Configurations_r17                                        `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r17,optional`
	DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17    []TwoPUCCH_Grp_Configurations_r17                                        `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r17,optional`
	Ack_NACK_FeedbackForMulticast_r17                 *CA_ParametersNR_v1720_ack_NACK_FeedbackForMulticast_r17                 `optional`
	Ptp_Retx_Multicast_r17                            *CA_ParametersNR_v1720_ptp_Retx_Multicast_r17                            `optional`
	Nack_OnlyFeedbackForMulticast_r17                 *CA_ParametersNR_v1720_nack_OnlyFeedbackForMulticast_r17                 `optional`
	Nack_OnlyFeedbackSpecificResourceForMulticast_r17 *CA_ParametersNR_v1720_nack_OnlyFeedbackSpecificResourceForMulticast_r17 `optional`
	Ack_NACK_FeedbackForSPS_Multicast_r17             *CA_ParametersNR_v1720_ack_NACK_FeedbackForSPS_Multicast_r17             `optional`
	Ptp_Retx_SPS_Multicast_r17                        *CA_ParametersNR_v1720_ptp_Retx_SPS_Multicast_r17                        `optional`
	HigherPowerLimit_r17                              *CA_ParametersNR_v1720_higherPowerLimit_r17                              `optional`
	ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17      *CA_ParametersNR_v1720_parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17      `optional`
	Pdcch_MonitoringCA_r17                            *int64                                                                   `lb:4,ub:16,optional`
	Pdcch_BlindDetectionMCG_SCG_List_r17              []PDCCH_BlindDetectionMCG_SCG_r17                                        `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
	Pdcch_BlindDetectionMixedList1_r17                []PDCCH_BlindDetectionMixed_r17                                          `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
	Pdcch_BlindDetectionMixedList2_r17                []PDCCH_BlindDetectionMixed_r17                                          `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
	Pdcch_BlindDetectionMixedList3_r17                []PDCCH_BlindDetectionMixed1_r17                                         `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
}

func (ie *CA_ParametersNR_v1720) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ParallelTxSRS_PUCCH_PUSCH_intraBand_r17 != nil, ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17 != nil, ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17 != nil, len(ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17) > 0, ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 != nil, ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 != nil, len(ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17) > 0, len(ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17) > 0, ie.Ack_NACK_FeedbackForMulticast_r17 != nil, ie.Ptp_Retx_Multicast_r17 != nil, ie.Nack_OnlyFeedbackForMulticast_r17 != nil, ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17 != nil, ie.Ack_NACK_FeedbackForSPS_Multicast_r17 != nil, ie.Ptp_Retx_SPS_Multicast_r17 != nil, ie.HigherPowerLimit_r17 != nil, ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17 != nil, ie.Pdcch_MonitoringCA_r17 != nil, len(ie.Pdcch_BlindDetectionMCG_SCG_List_r17) > 0, len(ie.Pdcch_BlindDetectionMixedList1_r17) > 0, len(ie.Pdcch_BlindDetectionMixedList2_r17) > 0, len(ie.Pdcch_BlindDetectionMixedList3_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ParallelTxSRS_PUCCH_PUSCH_intraBand_r17 != nil {
		if err = ie.ParallelTxSRS_PUCCH_PUSCH_intraBand_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ParallelTxSRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17 != nil {
		if err = ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17 != nil {
		if err = ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SemiStaticPUCCH_CellSwitchSingleGroup_r17", err)
		}
	}
	if len(ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17) > 0 {
		tmp_SemiStaticPUCCH_CellSwitchTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		for _, i := range ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17 {
			tmp_SemiStaticPUCCH_CellSwitchTwoGroups_r17.Value = append(tmp_SemiStaticPUCCH_CellSwitchTwoGroups_r17.Value, &i)
		}
		if err = tmp_SemiStaticPUCCH_CellSwitchTwoGroups_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SemiStaticPUCCH_CellSwitchTwoGroups_r17", err)
		}
	}
	if ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 != nil {
		if err = ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17", err)
		}
	}
	if ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 != nil {
		if err = ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17", err)
		}
	}
	if len(ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17) > 0 {
		tmp_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		for _, i := range ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 {
			tmp_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Value = append(tmp_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Value, &i)
		}
		if err = tmp_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17", err)
		}
	}
	if len(ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17) > 0 {
		tmp_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		for _, i := range ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 {
			tmp_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Value = append(tmp_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Value, &i)
		}
		if err = tmp_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17", err)
		}
	}
	if ie.Ack_NACK_FeedbackForMulticast_r17 != nil {
		if err = ie.Ack_NACK_FeedbackForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ack_NACK_FeedbackForMulticast_r17", err)
		}
	}
	if ie.Ptp_Retx_Multicast_r17 != nil {
		if err = ie.Ptp_Retx_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ptp_Retx_Multicast_r17", err)
		}
	}
	if ie.Nack_OnlyFeedbackForMulticast_r17 != nil {
		if err = ie.Nack_OnlyFeedbackForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Nack_OnlyFeedbackForMulticast_r17", err)
		}
	}
	if ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17 != nil {
		if err = ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Nack_OnlyFeedbackSpecificResourceForMulticast_r17", err)
		}
	}
	if ie.Ack_NACK_FeedbackForSPS_Multicast_r17 != nil {
		if err = ie.Ack_NACK_FeedbackForSPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ack_NACK_FeedbackForSPS_Multicast_r17", err)
		}
	}
	if ie.Ptp_Retx_SPS_Multicast_r17 != nil {
		if err = ie.Ptp_Retx_SPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ptp_Retx_SPS_Multicast_r17", err)
		}
	}
	if ie.HigherPowerLimit_r17 != nil {
		if err = ie.HigherPowerLimit_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HigherPowerLimit_r17", err)
		}
	}
	if ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17 != nil {
		if err = ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if ie.Pdcch_MonitoringCA_r17 != nil {
		if err = w.WriteInteger(*ie.Pdcch_MonitoringCA_r17, &aper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Pdcch_MonitoringCA_r17", err)
		}
	}
	if len(ie.Pdcch_BlindDetectionMCG_SCG_List_r17) > 0 {
		tmp_Pdcch_BlindDetectionMCG_SCG_List_r17 := utils.NewSequence[*PDCCH_BlindDetectionMCG_SCG_r17]([]*PDCCH_BlindDetectionMCG_SCG_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.Pdcch_BlindDetectionMCG_SCG_List_r17 {
			tmp_Pdcch_BlindDetectionMCG_SCG_List_r17.Value = append(tmp_Pdcch_BlindDetectionMCG_SCG_List_r17.Value, &i)
		}
		if err = tmp_Pdcch_BlindDetectionMCG_SCG_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionMCG_SCG_List_r17", err)
		}
	}
	if len(ie.Pdcch_BlindDetectionMixedList1_r17) > 0 {
		tmp_Pdcch_BlindDetectionMixedList1_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.Pdcch_BlindDetectionMixedList1_r17 {
			tmp_Pdcch_BlindDetectionMixedList1_r17.Value = append(tmp_Pdcch_BlindDetectionMixedList1_r17.Value, &i)
		}
		if err = tmp_Pdcch_BlindDetectionMixedList1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionMixedList1_r17", err)
		}
	}
	if len(ie.Pdcch_BlindDetectionMixedList2_r17) > 0 {
		tmp_Pdcch_BlindDetectionMixedList2_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.Pdcch_BlindDetectionMixedList2_r17 {
			tmp_Pdcch_BlindDetectionMixedList2_r17.Value = append(tmp_Pdcch_BlindDetectionMixedList2_r17.Value, &i)
		}
		if err = tmp_Pdcch_BlindDetectionMixedList2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionMixedList2_r17", err)
		}
	}
	if len(ie.Pdcch_BlindDetectionMixedList3_r17) > 0 {
		tmp_Pdcch_BlindDetectionMixedList3_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed1_r17]([]*PDCCH_BlindDetectionMixed1_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.Pdcch_BlindDetectionMixedList3_r17 {
			tmp_Pdcch_BlindDetectionMixedList3_r17.Value = append(tmp_Pdcch_BlindDetectionMixedList3_r17.Value, &i)
		}
		if err = tmp_Pdcch_BlindDetectionMixedList3_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionMixedList3_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1720) Decode(r *aper.AperReader) error {
	var err error
	var ParallelTxSRS_PUCCH_PUSCH_intraBand_r17Present bool
	if ParallelTxSRS_PUCCH_PUSCH_intraBand_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17Present bool
	if ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SemiStaticPUCCH_CellSwitchSingleGroup_r17Present bool
	if SemiStaticPUCCH_CellSwitchSingleGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SemiStaticPUCCH_CellSwitchTwoGroups_r17Present bool
	if SemiStaticPUCCH_CellSwitchTwoGroups_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17Present bool
	if DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17Present bool
	if DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17Present bool
	if DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17Present bool
	if DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ack_NACK_FeedbackForMulticast_r17Present bool
	if Ack_NACK_FeedbackForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ptp_Retx_Multicast_r17Present bool
	if Ptp_Retx_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Nack_OnlyFeedbackForMulticast_r17Present bool
	if Nack_OnlyFeedbackForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Nack_OnlyFeedbackSpecificResourceForMulticast_r17Present bool
	if Nack_OnlyFeedbackSpecificResourceForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ack_NACK_FeedbackForSPS_Multicast_r17Present bool
	if Ack_NACK_FeedbackForSPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ptp_Retx_SPS_Multicast_r17Present bool
	if Ptp_Retx_SPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HigherPowerLimit_r17Present bool
	if HigherPowerLimit_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17Present bool
	if ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_MonitoringCA_r17Present bool
	if Pdcch_MonitoringCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionMCG_SCG_List_r17Present bool
	if Pdcch_BlindDetectionMCG_SCG_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionMixedList1_r17Present bool
	if Pdcch_BlindDetectionMixedList1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionMixedList2_r17Present bool
	if Pdcch_BlindDetectionMixedList2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionMixedList3_r17Present bool
	if Pdcch_BlindDetectionMixedList3_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ParallelTxSRS_PUCCH_PUSCH_intraBand_r17Present {
		ie.ParallelTxSRS_PUCCH_PUSCH_intraBand_r17 = new(CA_ParametersNR_v1720_parallelTxSRS_PUCCH_PUSCH_intraBand_r17)
		if err = ie.ParallelTxSRS_PUCCH_PUSCH_intraBand_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ParallelTxSRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17Present {
		ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17 = new(CA_ParametersNR_v1720_parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17)
		if err = ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ParallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if SemiStaticPUCCH_CellSwitchSingleGroup_r17Present {
		ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17 = new(CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17)
		if err = ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SemiStaticPUCCH_CellSwitchSingleGroup_r17", err)
		}
	}
	if SemiStaticPUCCH_CellSwitchTwoGroups_r17Present {
		tmp_SemiStaticPUCCH_CellSwitchTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		fn_SemiStaticPUCCH_CellSwitchTwoGroups_r17 := func() *TwoPUCCH_Grp_Configurations_r17 {
			return new(TwoPUCCH_Grp_Configurations_r17)
		}
		if err = tmp_SemiStaticPUCCH_CellSwitchTwoGroups_r17.Decode(r, fn_SemiStaticPUCCH_CellSwitchTwoGroups_r17); err != nil {
			return utils.WrapError("Decode SemiStaticPUCCH_CellSwitchTwoGroups_r17", err)
		}
		ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17 = []TwoPUCCH_Grp_Configurations_r17{}
		for _, i := range tmp_SemiStaticPUCCH_CellSwitchTwoGroups_r17.Value {
			ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17 = append(ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17, *i)
		}
	}
	if DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17Present {
		ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 = new(CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17)
		if err = ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17", err)
		}
	}
	if DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17Present {
		ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 = new(CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17)
		if err = ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17", err)
		}
	}
	if DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17Present {
		tmp_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		fn_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 := func() *TwoPUCCH_Grp_Configurations_r17 {
			return new(TwoPUCCH_Grp_Configurations_r17)
		}
		if err = tmp_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Decode(r, fn_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17); err != nil {
			return utils.WrapError("Decode DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17", err)
		}
		ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 = []TwoPUCCH_Grp_Configurations_r17{}
		for _, i := range tmp_DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Value {
			ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 = append(ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17, *i)
		}
	}
	if DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17Present {
		tmp_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, aper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		fn_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 := func() *TwoPUCCH_Grp_Configurations_r17 {
			return new(TwoPUCCH_Grp_Configurations_r17)
		}
		if err = tmp_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Decode(r, fn_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17); err != nil {
			return utils.WrapError("Decode DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17", err)
		}
		ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 = []TwoPUCCH_Grp_Configurations_r17{}
		for _, i := range tmp_DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Value {
			ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 = append(ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17, *i)
		}
	}
	if Ack_NACK_FeedbackForMulticast_r17Present {
		ie.Ack_NACK_FeedbackForMulticast_r17 = new(CA_ParametersNR_v1720_ack_NACK_FeedbackForMulticast_r17)
		if err = ie.Ack_NACK_FeedbackForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ack_NACK_FeedbackForMulticast_r17", err)
		}
	}
	if Ptp_Retx_Multicast_r17Present {
		ie.Ptp_Retx_Multicast_r17 = new(CA_ParametersNR_v1720_ptp_Retx_Multicast_r17)
		if err = ie.Ptp_Retx_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ptp_Retx_Multicast_r17", err)
		}
	}
	if Nack_OnlyFeedbackForMulticast_r17Present {
		ie.Nack_OnlyFeedbackForMulticast_r17 = new(CA_ParametersNR_v1720_nack_OnlyFeedbackForMulticast_r17)
		if err = ie.Nack_OnlyFeedbackForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nack_OnlyFeedbackForMulticast_r17", err)
		}
	}
	if Nack_OnlyFeedbackSpecificResourceForMulticast_r17Present {
		ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17 = new(CA_ParametersNR_v1720_nack_OnlyFeedbackSpecificResourceForMulticast_r17)
		if err = ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nack_OnlyFeedbackSpecificResourceForMulticast_r17", err)
		}
	}
	if Ack_NACK_FeedbackForSPS_Multicast_r17Present {
		ie.Ack_NACK_FeedbackForSPS_Multicast_r17 = new(CA_ParametersNR_v1720_ack_NACK_FeedbackForSPS_Multicast_r17)
		if err = ie.Ack_NACK_FeedbackForSPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ack_NACK_FeedbackForSPS_Multicast_r17", err)
		}
	}
	if Ptp_Retx_SPS_Multicast_r17Present {
		ie.Ptp_Retx_SPS_Multicast_r17 = new(CA_ParametersNR_v1720_ptp_Retx_SPS_Multicast_r17)
		if err = ie.Ptp_Retx_SPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ptp_Retx_SPS_Multicast_r17", err)
		}
	}
	if HigherPowerLimit_r17Present {
		ie.HigherPowerLimit_r17 = new(CA_ParametersNR_v1720_higherPowerLimit_r17)
		if err = ie.HigherPowerLimit_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HigherPowerLimit_r17", err)
		}
	}
	if ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17Present {
		ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17 = new(CA_ParametersNR_v1720_parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17)
		if err = ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ParallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if Pdcch_MonitoringCA_r17Present {
		var tmp_int_Pdcch_MonitoringCA_r17 int64
		if tmp_int_Pdcch_MonitoringCA_r17, err = r.ReadInteger(&aper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Pdcch_MonitoringCA_r17", err)
		}
		ie.Pdcch_MonitoringCA_r17 = &tmp_int_Pdcch_MonitoringCA_r17
	}
	if Pdcch_BlindDetectionMCG_SCG_List_r17Present {
		tmp_Pdcch_BlindDetectionMCG_SCG_List_r17 := utils.NewSequence[*PDCCH_BlindDetectionMCG_SCG_r17]([]*PDCCH_BlindDetectionMCG_SCG_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_Pdcch_BlindDetectionMCG_SCG_List_r17 := func() *PDCCH_BlindDetectionMCG_SCG_r17 {
			return new(PDCCH_BlindDetectionMCG_SCG_r17)
		}
		if err = tmp_Pdcch_BlindDetectionMCG_SCG_List_r17.Decode(r, fn_Pdcch_BlindDetectionMCG_SCG_List_r17); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionMCG_SCG_List_r17", err)
		}
		ie.Pdcch_BlindDetectionMCG_SCG_List_r17 = []PDCCH_BlindDetectionMCG_SCG_r17{}
		for _, i := range tmp_Pdcch_BlindDetectionMCG_SCG_List_r17.Value {
			ie.Pdcch_BlindDetectionMCG_SCG_List_r17 = append(ie.Pdcch_BlindDetectionMCG_SCG_List_r17, *i)
		}
	}
	if Pdcch_BlindDetectionMixedList1_r17Present {
		tmp_Pdcch_BlindDetectionMixedList1_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_Pdcch_BlindDetectionMixedList1_r17 := func() *PDCCH_BlindDetectionMixed_r17 {
			return new(PDCCH_BlindDetectionMixed_r17)
		}
		if err = tmp_Pdcch_BlindDetectionMixedList1_r17.Decode(r, fn_Pdcch_BlindDetectionMixedList1_r17); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionMixedList1_r17", err)
		}
		ie.Pdcch_BlindDetectionMixedList1_r17 = []PDCCH_BlindDetectionMixed_r17{}
		for _, i := range tmp_Pdcch_BlindDetectionMixedList1_r17.Value {
			ie.Pdcch_BlindDetectionMixedList1_r17 = append(ie.Pdcch_BlindDetectionMixedList1_r17, *i)
		}
	}
	if Pdcch_BlindDetectionMixedList2_r17Present {
		tmp_Pdcch_BlindDetectionMixedList2_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_Pdcch_BlindDetectionMixedList2_r17 := func() *PDCCH_BlindDetectionMixed_r17 {
			return new(PDCCH_BlindDetectionMixed_r17)
		}
		if err = tmp_Pdcch_BlindDetectionMixedList2_r17.Decode(r, fn_Pdcch_BlindDetectionMixedList2_r17); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionMixedList2_r17", err)
		}
		ie.Pdcch_BlindDetectionMixedList2_r17 = []PDCCH_BlindDetectionMixed_r17{}
		for _, i := range tmp_Pdcch_BlindDetectionMixedList2_r17.Value {
			ie.Pdcch_BlindDetectionMixedList2_r17 = append(ie.Pdcch_BlindDetectionMixedList2_r17, *i)
		}
	}
	if Pdcch_BlindDetectionMixedList3_r17Present {
		tmp_Pdcch_BlindDetectionMixedList3_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed1_r17]([]*PDCCH_BlindDetectionMixed1_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_Pdcch_BlindDetectionMixedList3_r17 := func() *PDCCH_BlindDetectionMixed1_r17 {
			return new(PDCCH_BlindDetectionMixed1_r17)
		}
		if err = tmp_Pdcch_BlindDetectionMixedList3_r17.Decode(r, fn_Pdcch_BlindDetectionMixedList3_r17); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionMixedList3_r17", err)
		}
		ie.Pdcch_BlindDetectionMixedList3_r17 = []PDCCH_BlindDetectionMixed1_r17{}
		for _, i := range tmp_Pdcch_BlindDetectionMixedList3_r17.Value {
			ie.Pdcch_BlindDetectionMixedList3_r17 = append(ie.Pdcch_BlindDetectionMixedList3_r17, *i)
		}
	}
	return nil
}
