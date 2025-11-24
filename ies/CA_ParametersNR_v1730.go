package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1730 struct {
	Dmrs_BundlingPUSCH_RepTypeAPerBC_r17                  *CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeAPerBC_r17                  `optional`
	Dmrs_BundlingPUSCH_RepTypeBPerBC_r17                  *CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeBPerBC_r17                  `optional`
	Dmrs_BundlingPUSCH_multiSlotPerBC_r17                 *CA_ParametersNR_v1730_dmrs_BundlingPUSCH_multiSlotPerBC_r17                 `optional`
	Dmrs_BundlingPUCCH_RepPerBC_r17                       *CA_ParametersNR_v1730_dmrs_BundlingPUCCH_RepPerBC_r17                       `optional`
	Dmrs_BundlingRestartPerBC_r17                         *CA_ParametersNR_v1730_dmrs_BundlingRestartPerBC_r17                         `optional`
	Dmrs_BundlingNonBackToBackTX_PerBC_r17                *CA_ParametersNR_v1730_dmrs_BundlingNonBackToBackTX_PerBC_r17                `optional`
	StayOnTargetCC_SRS_CarrierSwitch_r17                  *CA_ParametersNR_v1730_stayOnTargetCC_SRS_CarrierSwitch_r17                  `optional`
	Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17       *CA_ParametersNR_v1730_fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17       `optional`
	Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 *CA_ParametersNR_v1730_mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 `optional`
	Mode1_ForType1_CodebookGeneration_r17                 *CA_ParametersNR_v1730_mode1_ForType1_CodebookGeneration_r17                 `optional`
	Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 *CA_ParametersNR_v1730_nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 `optional`
	MultiPUCCH_ConfigForMulticast_r17                     *CA_ParametersNR_v1730_multiPUCCH_ConfigForMulticast_r17                     `optional`
	Pucch_ConfigForSPS_Multicast_r17                      *CA_ParametersNR_v1730_pucch_ConfigForSPS_Multicast_r17                      `optional`
	MaxNumberG_RNTI_HARQ_ACK_Codebook_r17                 *int64                                                                       `lb:1,ub:4,optional`
	Mux_HARQ_ACK_UnicastMulticast_r17                     *CA_ParametersNR_v1730_mux_HARQ_ACK_UnicastMulticast_r17                     `optional`
}

func (ie *CA_ParametersNR_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17 != nil, ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17 != nil, ie.Dmrs_BundlingPUSCH_multiSlotPerBC_r17 != nil, ie.Dmrs_BundlingPUCCH_RepPerBC_r17 != nil, ie.Dmrs_BundlingRestartPerBC_r17 != nil, ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17 != nil, ie.StayOnTargetCC_SRS_CarrierSwitch_r17 != nil, ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil, ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil, ie.Mode1_ForType1_CodebookGeneration_r17 != nil, ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 != nil, ie.MultiPUCCH_ConfigForMulticast_r17 != nil, ie.Pucch_ConfigForSPS_Multicast_r17 != nil, ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17 != nil, ie.Mux_HARQ_ACK_UnicastMulticast_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17 != nil {
		if err = ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_BundlingPUSCH_RepTypeAPerBC_r17", err)
		}
	}
	if ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17 != nil {
		if err = ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_BundlingPUSCH_RepTypeBPerBC_r17", err)
		}
	}
	if ie.Dmrs_BundlingPUSCH_multiSlotPerBC_r17 != nil {
		if err = ie.Dmrs_BundlingPUSCH_multiSlotPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_BundlingPUSCH_multiSlotPerBC_r17", err)
		}
	}
	if ie.Dmrs_BundlingPUCCH_RepPerBC_r17 != nil {
		if err = ie.Dmrs_BundlingPUCCH_RepPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_BundlingPUCCH_RepPerBC_r17", err)
		}
	}
	if ie.Dmrs_BundlingRestartPerBC_r17 != nil {
		if err = ie.Dmrs_BundlingRestartPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_BundlingRestartPerBC_r17", err)
		}
	}
	if ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17 != nil {
		if err = ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_BundlingNonBackToBackTX_PerBC_r17", err)
		}
	}
	if ie.StayOnTargetCC_SRS_CarrierSwitch_r17 != nil {
		if err = ie.StayOnTargetCC_SRS_CarrierSwitch_r17.Encode(w); err != nil {
			return utils.WrapError("Encode StayOnTargetCC_SRS_CarrierSwitch_r17", err)
		}
	}
	if ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil {
		if err = ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil {
		if err = ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if ie.Mode1_ForType1_CodebookGeneration_r17 != nil {
		if err = ie.Mode1_ForType1_CodebookGeneration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mode1_ForType1_CodebookGeneration_r17", err)
		}
	}
	if ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 != nil {
		if err = ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17", err)
		}
	}
	if ie.MultiPUCCH_ConfigForMulticast_r17 != nil {
		if err = ie.MultiPUCCH_ConfigForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MultiPUCCH_ConfigForMulticast_r17", err)
		}
	}
	if ie.Pucch_ConfigForSPS_Multicast_r17 != nil {
		if err = ie.Pucch_ConfigForSPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_ConfigForSPS_Multicast_r17", err)
		}
	}
	if ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17 != nil {
		if err = w.WriteInteger(*ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode MaxNumberG_RNTI_HARQ_ACK_Codebook_r17", err)
		}
	}
	if ie.Mux_HARQ_ACK_UnicastMulticast_r17 != nil {
		if err = ie.Mux_HARQ_ACK_UnicastMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_HARQ_ACK_UnicastMulticast_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1730) Decode(r *uper.UperReader) error {
	var err error
	var Dmrs_BundlingPUSCH_RepTypeAPerBC_r17Present bool
	if Dmrs_BundlingPUSCH_RepTypeAPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_BundlingPUSCH_RepTypeBPerBC_r17Present bool
	if Dmrs_BundlingPUSCH_RepTypeBPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_BundlingPUSCH_multiSlotPerBC_r17Present bool
	if Dmrs_BundlingPUSCH_multiSlotPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_BundlingPUCCH_RepPerBC_r17Present bool
	if Dmrs_BundlingPUCCH_RepPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_BundlingRestartPerBC_r17Present bool
	if Dmrs_BundlingRestartPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_BundlingNonBackToBackTX_PerBC_r17Present bool
	if Dmrs_BundlingNonBackToBackTX_PerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var StayOnTargetCC_SRS_CarrierSwitch_r17Present bool
	if StayOnTargetCC_SRS_CarrierSwitch_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present bool
	if Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present bool
	if Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mode1_ForType1_CodebookGeneration_r17Present bool
	if Mode1_ForType1_CodebookGeneration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17Present bool
	if Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiPUCCH_ConfigForMulticast_r17Present bool
	if MultiPUCCH_ConfigForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_ConfigForSPS_Multicast_r17Present bool
	if Pucch_ConfigForSPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberG_RNTI_HARQ_ACK_Codebook_r17Present bool
	if MaxNumberG_RNTI_HARQ_ACK_Codebook_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_HARQ_ACK_UnicastMulticast_r17Present bool
	if Mux_HARQ_ACK_UnicastMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dmrs_BundlingPUSCH_RepTypeAPerBC_r17Present {
		ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeAPerBC_r17)
		if err = ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_BundlingPUSCH_RepTypeAPerBC_r17", err)
		}
	}
	if Dmrs_BundlingPUSCH_RepTypeBPerBC_r17Present {
		ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeBPerBC_r17)
		if err = ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_BundlingPUSCH_RepTypeBPerBC_r17", err)
		}
	}
	if Dmrs_BundlingPUSCH_multiSlotPerBC_r17Present {
		ie.Dmrs_BundlingPUSCH_multiSlotPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUSCH_multiSlotPerBC_r17)
		if err = ie.Dmrs_BundlingPUSCH_multiSlotPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_BundlingPUSCH_multiSlotPerBC_r17", err)
		}
	}
	if Dmrs_BundlingPUCCH_RepPerBC_r17Present {
		ie.Dmrs_BundlingPUCCH_RepPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUCCH_RepPerBC_r17)
		if err = ie.Dmrs_BundlingPUCCH_RepPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_BundlingPUCCH_RepPerBC_r17", err)
		}
	}
	if Dmrs_BundlingRestartPerBC_r17Present {
		ie.Dmrs_BundlingRestartPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingRestartPerBC_r17)
		if err = ie.Dmrs_BundlingRestartPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_BundlingRestartPerBC_r17", err)
		}
	}
	if Dmrs_BundlingNonBackToBackTX_PerBC_r17Present {
		ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingNonBackToBackTX_PerBC_r17)
		if err = ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_BundlingNonBackToBackTX_PerBC_r17", err)
		}
	}
	if StayOnTargetCC_SRS_CarrierSwitch_r17Present {
		ie.StayOnTargetCC_SRS_CarrierSwitch_r17 = new(CA_ParametersNR_v1730_stayOnTargetCC_SRS_CarrierSwitch_r17)
		if err = ie.StayOnTargetCC_SRS_CarrierSwitch_r17.Decode(r); err != nil {
			return utils.WrapError("Decode StayOnTargetCC_SRS_CarrierSwitch_r17", err)
		}
	}
	if Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present {
		ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 = new(CA_ParametersNR_v1730_fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17)
		if err = ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present {
		ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 = new(CA_ParametersNR_v1730_mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17)
		if err = ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if Mode1_ForType1_CodebookGeneration_r17Present {
		ie.Mode1_ForType1_CodebookGeneration_r17 = new(CA_ParametersNR_v1730_mode1_ForType1_CodebookGeneration_r17)
		if err = ie.Mode1_ForType1_CodebookGeneration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mode1_ForType1_CodebookGeneration_r17", err)
		}
	}
	if Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17Present {
		ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 = new(CA_ParametersNR_v1730_nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17)
		if err = ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17", err)
		}
	}
	if MultiPUCCH_ConfigForMulticast_r17Present {
		ie.MultiPUCCH_ConfigForMulticast_r17 = new(CA_ParametersNR_v1730_multiPUCCH_ConfigForMulticast_r17)
		if err = ie.MultiPUCCH_ConfigForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MultiPUCCH_ConfigForMulticast_r17", err)
		}
	}
	if Pucch_ConfigForSPS_Multicast_r17Present {
		ie.Pucch_ConfigForSPS_Multicast_r17 = new(CA_ParametersNR_v1730_pucch_ConfigForSPS_Multicast_r17)
		if err = ie.Pucch_ConfigForSPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_ConfigForSPS_Multicast_r17", err)
		}
	}
	if MaxNumberG_RNTI_HARQ_ACK_Codebook_r17Present {
		var tmp_int_MaxNumberG_RNTI_HARQ_ACK_Codebook_r17 int64
		if tmp_int_MaxNumberG_RNTI_HARQ_ACK_Codebook_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode MaxNumberG_RNTI_HARQ_ACK_Codebook_r17", err)
		}
		ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17 = &tmp_int_MaxNumberG_RNTI_HARQ_ACK_Codebook_r17
	}
	if Mux_HARQ_ACK_UnicastMulticast_r17Present {
		ie.Mux_HARQ_ACK_UnicastMulticast_r17 = new(CA_ParametersNR_v1730_mux_HARQ_ACK_UnicastMulticast_r17)
		if err = ie.Mux_HARQ_ACK_UnicastMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_HARQ_ACK_UnicastMulticast_r17", err)
		}
	}
	return nil
}
