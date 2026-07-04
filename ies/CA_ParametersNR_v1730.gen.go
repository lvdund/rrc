// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1730 (line 17561).

var cAParametersNRV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dmrs-BundlingPUSCH-RepTypeAPerBC-r17", Optional: true},
		{Name: "dmrs-BundlingPUSCH-RepTypeBPerBC-r17", Optional: true},
		{Name: "dmrs-BundlingPUSCH-multiSlotPerBC-r17", Optional: true},
		{Name: "dmrs-BundlingPUCCH-RepPerBC-r17", Optional: true},
		{Name: "dmrs-BundlingRestartPerBC-r17", Optional: true},
		{Name: "dmrs-BundlingNonBackToBackTX-PerBC-r17", Optional: true},
		{Name: "stayOnTargetCC-SRS-CarrierSwitch-r17", Optional: true},
		{Name: "fdm-CodebookForMux-UnicastMulticastHARQ-ACK-r17", Optional: true},
		{Name: "mode2-TDM-CodebookForMux-UnicastMulticastHARQ-ACK-r17", Optional: true},
		{Name: "mode1-ForType1-CodebookGeneration-r17", Optional: true},
		{Name: "nack-OnlyFeedbackSpecificResourceForSPS-Multicast-r17", Optional: true},
		{Name: "multiPUCCH-ConfigForMulticast-r17", Optional: true},
		{Name: "pucch-ConfigForSPS-Multicast-r17", Optional: true},
		{Name: "maxNumberG-RNTI-HARQ-ACK-Codebook-r17", Optional: true},
		{Name: "mux-HARQ-ACK-UnicastMulticast-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1730_Dmrs_BundlingPUSCH_RepTypeAPerBC_r17_Supported = 0
)

var cAParametersNRV1730DmrsBundlingPUSCHRepTypeAPerBCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Dmrs_BundlingPUSCH_RepTypeAPerBC_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Dmrs_BundlingPUSCH_RepTypeBPerBC_r17_Supported = 0
)

var cAParametersNRV1730DmrsBundlingPUSCHRepTypeBPerBCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Dmrs_BundlingPUSCH_RepTypeBPerBC_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Dmrs_BundlingPUSCH_MultiSlotPerBC_r17_Supported = 0
)

var cAParametersNRV1730DmrsBundlingPUSCHMultiSlotPerBCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Dmrs_BundlingPUSCH_MultiSlotPerBC_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Dmrs_BundlingPUCCH_RepPerBC_r17_Supported = 0
)

var cAParametersNRV1730DmrsBundlingPUCCHRepPerBCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Dmrs_BundlingPUCCH_RepPerBC_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Dmrs_BundlingRestartPerBC_r17_Supported = 0
)

var cAParametersNRV1730DmrsBundlingRestartPerBCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Dmrs_BundlingRestartPerBC_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Dmrs_BundlingNonBackToBackTX_PerBC_r17_Supported = 0
)

var cAParametersNRV1730DmrsBundlingNonBackToBackTXPerBCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Dmrs_BundlingNonBackToBackTX_PerBC_r17_Supported},
}

const (
	CA_ParametersNR_v1730_StayOnTargetCC_SRS_CarrierSwitch_r17_Supported = 0
)

var cAParametersNRV1730StayOnTargetCCSRSCarrierSwitchR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_StayOnTargetCC_SRS_CarrierSwitch_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17_Supported = 0
)

var cAParametersNRV1730FdmCodebookForMuxUnicastMulticastHARQACKR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17_Supported = 0
)

var cAParametersNRV1730Mode2TDMCodebookForMuxUnicastMulticastHARQACKR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Mode1_ForType1_CodebookGeneration_r17_Supported = 0
)

var cAParametersNRV1730Mode1ForType1CodebookGenerationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Mode1_ForType1_CodebookGeneration_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17_Supported = 0
)

var cAParametersNRV1730NackOnlyFeedbackSpecificResourceForSPSMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17_Supported},
}

const (
	CA_ParametersNR_v1730_MultiPUCCH_ConfigForMulticast_r17_Supported = 0
)

var cAParametersNRV1730MultiPUCCHConfigForMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_MultiPUCCH_ConfigForMulticast_r17_Supported},
}

const (
	CA_ParametersNR_v1730_Pucch_ConfigForSPS_Multicast_r17_Supported = 0
)

var cAParametersNRV1730PucchConfigForSPSMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Pucch_ConfigForSPS_Multicast_r17_Supported},
}

var cAParametersNRV1730MaxNumberGRNTIHARQACKCodebookR17Constraints = per.Constrained(1, 4)

const (
	CA_ParametersNR_v1730_Mux_HARQ_ACK_UnicastMulticast_r17_Supported = 0
)

var cAParametersNRV1730MuxHARQACKUnicastMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1730_Mux_HARQ_ACK_UnicastMulticast_r17_Supported},
}

type CA_ParametersNR_v1730 struct {
	Dmrs_BundlingPUSCH_RepTypeAPerBC_r17                  *int64
	Dmrs_BundlingPUSCH_RepTypeBPerBC_r17                  *int64
	Dmrs_BundlingPUSCH_MultiSlotPerBC_r17                 *int64
	Dmrs_BundlingPUCCH_RepPerBC_r17                       *int64
	Dmrs_BundlingRestartPerBC_r17                         *int64
	Dmrs_BundlingNonBackToBackTX_PerBC_r17                *int64
	StayOnTargetCC_SRS_CarrierSwitch_r17                  *int64
	Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17       *int64
	Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 *int64
	Mode1_ForType1_CodebookGeneration_r17                 *int64
	Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 *int64
	MultiPUCCH_ConfigForMulticast_r17                     *int64
	Pucch_ConfigForSPS_Multicast_r17                      *int64
	MaxNumberG_RNTI_HARQ_ACK_Codebook_r17                 *int64
	Mux_HARQ_ACK_UnicastMulticast_r17                     *int64
}

func (ie *CA_ParametersNR_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1730Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17 != nil, ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17 != nil, ie.Dmrs_BundlingPUSCH_MultiSlotPerBC_r17 != nil, ie.Dmrs_BundlingPUCCH_RepPerBC_r17 != nil, ie.Dmrs_BundlingRestartPerBC_r17 != nil, ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17 != nil, ie.StayOnTargetCC_SRS_CarrierSwitch_r17 != nil, ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil, ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil, ie.Mode1_ForType1_CodebookGeneration_r17 != nil, ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 != nil, ie.MultiPUCCH_ConfigForMulticast_r17 != nil, ie.Pucch_ConfigForSPS_Multicast_r17 != nil, ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17 != nil, ie.Mux_HARQ_ACK_UnicastMulticast_r17 != nil}); err != nil {
		return err
	}

	// 2. dmrs-BundlingPUSCH-RepTypeAPerBC-r17: enumerated
	{
		if ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17, cAParametersNRV1730DmrsBundlingPUSCHRepTypeAPerBCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. dmrs-BundlingPUSCH-RepTypeBPerBC-r17: enumerated
	{
		if ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17, cAParametersNRV1730DmrsBundlingPUSCHRepTypeBPerBCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. dmrs-BundlingPUSCH-multiSlotPerBC-r17: enumerated
	{
		if ie.Dmrs_BundlingPUSCH_MultiSlotPerBC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_BundlingPUSCH_MultiSlotPerBC_r17, cAParametersNRV1730DmrsBundlingPUSCHMultiSlotPerBCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. dmrs-BundlingPUCCH-RepPerBC-r17: enumerated
	{
		if ie.Dmrs_BundlingPUCCH_RepPerBC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_BundlingPUCCH_RepPerBC_r17, cAParametersNRV1730DmrsBundlingPUCCHRepPerBCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. dmrs-BundlingRestartPerBC-r17: enumerated
	{
		if ie.Dmrs_BundlingRestartPerBC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_BundlingRestartPerBC_r17, cAParametersNRV1730DmrsBundlingRestartPerBCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. dmrs-BundlingNonBackToBackTX-PerBC-r17: enumerated
	{
		if ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17, cAParametersNRV1730DmrsBundlingNonBackToBackTXPerBCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. stayOnTargetCC-SRS-CarrierSwitch-r17: enumerated
	{
		if ie.StayOnTargetCC_SRS_CarrierSwitch_r17 != nil {
			if err := e.EncodeEnumerated(*ie.StayOnTargetCC_SRS_CarrierSwitch_r17, cAParametersNRV1730StayOnTargetCCSRSCarrierSwitchR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. fdm-CodebookForMux-UnicastMulticastHARQ-ACK-r17: enumerated
	{
		if ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17, cAParametersNRV1730FdmCodebookForMuxUnicastMulticastHARQACKR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. mode2-TDM-CodebookForMux-UnicastMulticastHARQ-ACK-r17: enumerated
	{
		if ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17, cAParametersNRV1730Mode2TDMCodebookForMuxUnicastMulticastHARQACKR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. mode1-ForType1-CodebookGeneration-r17: enumerated
	{
		if ie.Mode1_ForType1_CodebookGeneration_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Mode1_ForType1_CodebookGeneration_r17, cAParametersNRV1730Mode1ForType1CodebookGenerationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. nack-OnlyFeedbackSpecificResourceForSPS-Multicast-r17: enumerated
	{
		if ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17, cAParametersNRV1730NackOnlyFeedbackSpecificResourceForSPSMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. multiPUCCH-ConfigForMulticast-r17: enumerated
	{
		if ie.MultiPUCCH_ConfigForMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MultiPUCCH_ConfigForMulticast_r17, cAParametersNRV1730MultiPUCCHConfigForMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. pucch-ConfigForSPS-Multicast-r17: enumerated
	{
		if ie.Pucch_ConfigForSPS_Multicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_ConfigForSPS_Multicast_r17, cAParametersNRV1730PucchConfigForSPSMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 15. maxNumberG-RNTI-HARQ-ACK-Codebook-r17: integer
	{
		if ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17 != nil {
			if err := e.EncodeInteger(*ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17, cAParametersNRV1730MaxNumberGRNTIHARQACKCodebookR17Constraints); err != nil {
				return err
			}
		}
	}

	// 16. mux-HARQ-ACK-UnicastMulticast-r17: enumerated
	{
		if ie.Mux_HARQ_ACK_UnicastMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Mux_HARQ_ACK_UnicastMulticast_r17, cAParametersNRV1730MuxHARQACKUnicastMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1730Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dmrs-BundlingPUSCH-RepTypeAPerBC-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730DmrsBundlingPUSCHRepTypeAPerBCR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUSCH_RepTypeAPerBC_r17 = &idx
		}
	}

	// 3. dmrs-BundlingPUSCH-RepTypeBPerBC-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730DmrsBundlingPUSCHRepTypeBPerBCR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUSCH_RepTypeBPerBC_r17 = &idx
		}
	}

	// 4. dmrs-BundlingPUSCH-multiSlotPerBC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730DmrsBundlingPUSCHMultiSlotPerBCR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUSCH_MultiSlotPerBC_r17 = &idx
		}
	}

	// 5. dmrs-BundlingPUCCH-RepPerBC-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730DmrsBundlingPUCCHRepPerBCR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUCCH_RepPerBC_r17 = &idx
		}
	}

	// 6. dmrs-BundlingRestartPerBC-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730DmrsBundlingRestartPerBCR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingRestartPerBC_r17 = &idx
		}
	}

	// 7. dmrs-BundlingNonBackToBackTX-PerBC-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730DmrsBundlingNonBackToBackTXPerBCR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingNonBackToBackTX_PerBC_r17 = &idx
		}
	}

	// 8. stayOnTargetCC-SRS-CarrierSwitch-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730StayOnTargetCCSRSCarrierSwitchR17Constraints)
			if err != nil {
				return err
			}
			ie.StayOnTargetCC_SRS_CarrierSwitch_r17 = &idx
		}
	}

	// 9. fdm-CodebookForMux-UnicastMulticastHARQ-ACK-r17: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730FdmCodebookForMuxUnicastMulticastHARQACKR17Constraints)
			if err != nil {
				return err
			}
			ie.Fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 = &idx
		}
	}

	// 10. mode2-TDM-CodebookForMux-UnicastMulticastHARQ-ACK-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730Mode2TDMCodebookForMuxUnicastMulticastHARQACKR17Constraints)
			if err != nil {
				return err
			}
			ie.Mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 = &idx
		}
	}

	// 11. mode1-ForType1-CodebookGeneration-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730Mode1ForType1CodebookGenerationR17Constraints)
			if err != nil {
				return err
			}
			ie.Mode1_ForType1_CodebookGeneration_r17 = &idx
		}
	}

	// 12. nack-OnlyFeedbackSpecificResourceForSPS-Multicast-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730NackOnlyFeedbackSpecificResourceForSPSMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 = &idx
		}
	}

	// 13. multiPUCCH-ConfigForMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730MultiPUCCHConfigForMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUCCH_ConfigForMulticast_r17 = &idx
		}
	}

	// 14. pucch-ConfigForSPS-Multicast-r17: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730PucchConfigForSPSMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_ConfigForSPS_Multicast_r17 = &idx
		}
	}

	// 15. maxNumberG-RNTI-HARQ-ACK-Codebook-r17: integer
	{
		if seq.IsComponentPresent(13) {
			v, err := d.DecodeInteger(cAParametersNRV1730MaxNumberGRNTIHARQACKCodebookR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberG_RNTI_HARQ_ACK_Codebook_r17 = &v
		}
	}

	// 16. mux-HARQ-ACK-UnicastMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1730MuxHARQACKUnicastMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Mux_HARQ_ACK_UnicastMulticast_r17 = &idx
		}
	}

	return nil
}
