package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersCommon struct {
	Lcp_Restriction                           *MAC_ParametersCommon_lcp_Restriction                           `optional`
	Dummy                                     *MAC_ParametersCommon_dummy                                     `optional`
	Lch_ToSCellRestriction                    *MAC_ParametersCommon_lch_ToSCellRestriction                    `optional`
	RecommendedBitRate                        *MAC_ParametersCommon_recommendedBitRate                        `optional,ext-1`
	RecommendedBitRateQuery                   *MAC_ParametersCommon_recommendedBitRateQuery                   `optional,ext-1`
	RecommendedBitRateMultiplier_r16          *MAC_ParametersCommon_recommendedBitRateMultiplier_r16          `optional,ext-2`
	PreEmptiveBSR_r16                         *MAC_ParametersCommon_preEmptiveBSR_r16                         `optional,ext-2`
	AutonomousTransmission_r16                *MAC_ParametersCommon_autonomousTransmission_r16                `optional,ext-2`
	Lch_PriorityBasedPrioritization_r16       *MAC_ParametersCommon_lch_PriorityBasedPrioritization_r16       `optional,ext-2`
	Lch_ToConfiguredGrantMapping_r16          *MAC_ParametersCommon_lch_ToConfiguredGrantMapping_r16          `optional,ext-2`
	Lch_ToGrantPriorityRestriction_r16        *MAC_ParametersCommon_lch_ToGrantPriorityRestriction_r16        `optional,ext-2`
	SinglePHR_P_r16                           *MAC_ParametersCommon_singlePHR_P_r16                           `optional,ext-2`
	Ul_LBT_FailureDetectionRecovery_r16       *MAC_ParametersCommon_ul_LBT_FailureDetectionRecovery_r16       `optional,ext-2`
	Tdd_MPE_P_MPR_Reporting_r16               *MAC_ParametersCommon_tdd_MPE_P_MPR_Reporting_r16               `optional,ext-2`
	Lcid_ExtensionIAB_r16                     *MAC_ParametersCommon_lcid_ExtensionIAB_r16                     `optional,ext-2`
	SpCell_BFR_CBRA_r16                       *MAC_ParametersCommon_spCell_BFR_CBRA_r16                       `optional,ext-3`
	Srs_ResourceId_Ext_r16                    *MAC_ParametersCommon_srs_ResourceId_Ext_r16                    `optional,ext-4`
	EnhancedUuDRX_forSidelink_r17             *MAC_ParametersCommon_enhancedUuDRX_forSidelink_r17             `optional,ext-5`
	Mg_ActivationRequestPRS_Meas_r17          *MAC_ParametersCommon_mg_ActivationRequestPRS_Meas_r17          `optional,ext-5`
	Mg_ActivationCommPRS_Meas_r17             *MAC_ParametersCommon_mg_ActivationCommPRS_Meas_r17             `optional,ext-5`
	IntraCG_Prioritization_r17                *MAC_ParametersCommon_intraCG_Prioritization_r17                `optional,ext-5`
	JointPrioritizationCG_Retx_Timer_r17      *MAC_ParametersCommon_jointPrioritizationCG_Retx_Timer_r17      `optional,ext-5`
	SurvivalTime_r17                          *MAC_ParametersCommon_survivalTime_r17                          `optional,ext-5`
	Lcg_ExtensionIAB_r17                      *MAC_ParametersCommon_lcg_ExtensionIAB_r17                      `optional,ext-5`
	Harq_FeedbackDisabled_r17                 *MAC_ParametersCommon_harq_FeedbackDisabled_r17                 `optional,ext-5`
	Uplink_Harq_ModeB_r17                     *MAC_ParametersCommon_uplink_Harq_ModeB_r17                     `optional,ext-5`
	Sr_TriggeredBy_TA_Report_r17              *MAC_ParametersCommon_sr_TriggeredBy_TA_Report_r17              `optional,ext-5`
	ExtendedDRX_CycleInactive_r17             *MAC_ParametersCommon_extendedDRX_CycleInactive_r17             `optional,ext-5`
	SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 *MAC_ParametersCommon_simultaneousSR_PUSCH_DiffPUCCH_groups_r17 `optional,ext-5`
	LastTransmissionUL_r17                    *MAC_ParametersCommon_lastTransmissionUL_r17                    `optional,ext-5`
}

func (ie *MAC_ParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.RecommendedBitRate != nil || ie.RecommendedBitRateQuery != nil || ie.RecommendedBitRateMultiplier_r16 != nil || ie.PreEmptiveBSR_r16 != nil || ie.AutonomousTransmission_r16 != nil || ie.Lch_PriorityBasedPrioritization_r16 != nil || ie.Lch_ToConfiguredGrantMapping_r16 != nil || ie.Lch_ToGrantPriorityRestriction_r16 != nil || ie.SinglePHR_P_r16 != nil || ie.Ul_LBT_FailureDetectionRecovery_r16 != nil || ie.Tdd_MPE_P_MPR_Reporting_r16 != nil || ie.Lcid_ExtensionIAB_r16 != nil || ie.SpCell_BFR_CBRA_r16 != nil || ie.Srs_ResourceId_Ext_r16 != nil || ie.EnhancedUuDRX_forSidelink_r17 != nil || ie.Mg_ActivationRequestPRS_Meas_r17 != nil || ie.Mg_ActivationCommPRS_Meas_r17 != nil || ie.IntraCG_Prioritization_r17 != nil || ie.JointPrioritizationCG_Retx_Timer_r17 != nil || ie.SurvivalTime_r17 != nil || ie.Lcg_ExtensionIAB_r17 != nil || ie.Harq_FeedbackDisabled_r17 != nil || ie.Uplink_Harq_ModeB_r17 != nil || ie.Sr_TriggeredBy_TA_Report_r17 != nil || ie.ExtendedDRX_CycleInactive_r17 != nil || ie.SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil || ie.LastTransmissionUL_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Lcp_Restriction != nil, ie.Dummy != nil, ie.Lch_ToSCellRestriction != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Lcp_Restriction != nil {
		if err = ie.Lcp_Restriction.Encode(w); err != nil {
			return utils.WrapError("Encode Lcp_Restriction", err)
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.Lch_ToSCellRestriction != nil {
		if err = ie.Lch_ToSCellRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode Lch_ToSCellRestriction", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.RecommendedBitRate != nil || ie.RecommendedBitRateQuery != nil, ie.RecommendedBitRateMultiplier_r16 != nil || ie.PreEmptiveBSR_r16 != nil || ie.AutonomousTransmission_r16 != nil || ie.Lch_PriorityBasedPrioritization_r16 != nil || ie.Lch_ToConfiguredGrantMapping_r16 != nil || ie.Lch_ToGrantPriorityRestriction_r16 != nil || ie.SinglePHR_P_r16 != nil || ie.Ul_LBT_FailureDetectionRecovery_r16 != nil || ie.Tdd_MPE_P_MPR_Reporting_r16 != nil || ie.Lcid_ExtensionIAB_r16 != nil, ie.SpCell_BFR_CBRA_r16 != nil, ie.Srs_ResourceId_Ext_r16 != nil, ie.EnhancedUuDRX_forSidelink_r17 != nil || ie.Mg_ActivationRequestPRS_Meas_r17 != nil || ie.Mg_ActivationCommPRS_Meas_r17 != nil || ie.IntraCG_Prioritization_r17 != nil || ie.JointPrioritizationCG_Retx_Timer_r17 != nil || ie.SurvivalTime_r17 != nil || ie.Lcg_ExtensionIAB_r17 != nil || ie.Harq_FeedbackDisabled_r17 != nil || ie.Uplink_Harq_ModeB_r17 != nil || ie.Sr_TriggeredBy_TA_Report_r17 != nil || ie.ExtendedDRX_CycleInactive_r17 != nil || ie.SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil || ie.LastTransmissionUL_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.RecommendedBitRate != nil, ie.RecommendedBitRateQuery != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RecommendedBitRate optional
			if ie.RecommendedBitRate != nil {
				if err = ie.RecommendedBitRate.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RecommendedBitRate", err)
				}
			}
			// encode RecommendedBitRateQuery optional
			if ie.RecommendedBitRateQuery != nil {
				if err = ie.RecommendedBitRateQuery.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RecommendedBitRateQuery", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.RecommendedBitRateMultiplier_r16 != nil, ie.PreEmptiveBSR_r16 != nil, ie.AutonomousTransmission_r16 != nil, ie.Lch_PriorityBasedPrioritization_r16 != nil, ie.Lch_ToConfiguredGrantMapping_r16 != nil, ie.Lch_ToGrantPriorityRestriction_r16 != nil, ie.SinglePHR_P_r16 != nil, ie.Ul_LBT_FailureDetectionRecovery_r16 != nil, ie.Tdd_MPE_P_MPR_Reporting_r16 != nil, ie.Lcid_ExtensionIAB_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RecommendedBitRateMultiplier_r16 optional
			if ie.RecommendedBitRateMultiplier_r16 != nil {
				if err = ie.RecommendedBitRateMultiplier_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RecommendedBitRateMultiplier_r16", err)
				}
			}
			// encode PreEmptiveBSR_r16 optional
			if ie.PreEmptiveBSR_r16 != nil {
				if err = ie.PreEmptiveBSR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PreEmptiveBSR_r16", err)
				}
			}
			// encode AutonomousTransmission_r16 optional
			if ie.AutonomousTransmission_r16 != nil {
				if err = ie.AutonomousTransmission_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AutonomousTransmission_r16", err)
				}
			}
			// encode Lch_PriorityBasedPrioritization_r16 optional
			if ie.Lch_PriorityBasedPrioritization_r16 != nil {
				if err = ie.Lch_PriorityBasedPrioritization_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lch_PriorityBasedPrioritization_r16", err)
				}
			}
			// encode Lch_ToConfiguredGrantMapping_r16 optional
			if ie.Lch_ToConfiguredGrantMapping_r16 != nil {
				if err = ie.Lch_ToConfiguredGrantMapping_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lch_ToConfiguredGrantMapping_r16", err)
				}
			}
			// encode Lch_ToGrantPriorityRestriction_r16 optional
			if ie.Lch_ToGrantPriorityRestriction_r16 != nil {
				if err = ie.Lch_ToGrantPriorityRestriction_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lch_ToGrantPriorityRestriction_r16", err)
				}
			}
			// encode SinglePHR_P_r16 optional
			if ie.SinglePHR_P_r16 != nil {
				if err = ie.SinglePHR_P_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SinglePHR_P_r16", err)
				}
			}
			// encode Ul_LBT_FailureDetectionRecovery_r16 optional
			if ie.Ul_LBT_FailureDetectionRecovery_r16 != nil {
				if err = ie.Ul_LBT_FailureDetectionRecovery_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_LBT_FailureDetectionRecovery_r16", err)
				}
			}
			// encode Tdd_MPE_P_MPR_Reporting_r16 optional
			if ie.Tdd_MPE_P_MPR_Reporting_r16 != nil {
				if err = ie.Tdd_MPE_P_MPR_Reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tdd_MPE_P_MPR_Reporting_r16", err)
				}
			}
			// encode Lcid_ExtensionIAB_r16 optional
			if ie.Lcid_ExtensionIAB_r16 != nil {
				if err = ie.Lcid_ExtensionIAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lcid_ExtensionIAB_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.SpCell_BFR_CBRA_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SpCell_BFR_CBRA_r16 optional
			if ie.SpCell_BFR_CBRA_r16 != nil {
				if err = ie.SpCell_BFR_CBRA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpCell_BFR_CBRA_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.Srs_ResourceId_Ext_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Srs_ResourceId_Ext_r16 optional
			if ie.Srs_ResourceId_Ext_r16 != nil {
				if err = ie.Srs_ResourceId_Ext_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_ResourceId_Ext_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.EnhancedUuDRX_forSidelink_r17 != nil, ie.Mg_ActivationRequestPRS_Meas_r17 != nil, ie.Mg_ActivationCommPRS_Meas_r17 != nil, ie.IntraCG_Prioritization_r17 != nil, ie.JointPrioritizationCG_Retx_Timer_r17 != nil, ie.SurvivalTime_r17 != nil, ie.Lcg_ExtensionIAB_r17 != nil, ie.Harq_FeedbackDisabled_r17 != nil, ie.Uplink_Harq_ModeB_r17 != nil, ie.Sr_TriggeredBy_TA_Report_r17 != nil, ie.ExtendedDRX_CycleInactive_r17 != nil, ie.SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil, ie.LastTransmissionUL_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnhancedUuDRX_forSidelink_r17 optional
			if ie.EnhancedUuDRX_forSidelink_r17 != nil {
				if err = ie.EnhancedUuDRX_forSidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedUuDRX_forSidelink_r17", err)
				}
			}
			// encode Mg_ActivationRequestPRS_Meas_r17 optional
			if ie.Mg_ActivationRequestPRS_Meas_r17 != nil {
				if err = ie.Mg_ActivationRequestPRS_Meas_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mg_ActivationRequestPRS_Meas_r17", err)
				}
			}
			// encode Mg_ActivationCommPRS_Meas_r17 optional
			if ie.Mg_ActivationCommPRS_Meas_r17 != nil {
				if err = ie.Mg_ActivationCommPRS_Meas_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mg_ActivationCommPRS_Meas_r17", err)
				}
			}
			// encode IntraCG_Prioritization_r17 optional
			if ie.IntraCG_Prioritization_r17 != nil {
				if err = ie.IntraCG_Prioritization_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraCG_Prioritization_r17", err)
				}
			}
			// encode JointPrioritizationCG_Retx_Timer_r17 optional
			if ie.JointPrioritizationCG_Retx_Timer_r17 != nil {
				if err = ie.JointPrioritizationCG_Retx_Timer_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode JointPrioritizationCG_Retx_Timer_r17", err)
				}
			}
			// encode SurvivalTime_r17 optional
			if ie.SurvivalTime_r17 != nil {
				if err = ie.SurvivalTime_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SurvivalTime_r17", err)
				}
			}
			// encode Lcg_ExtensionIAB_r17 optional
			if ie.Lcg_ExtensionIAB_r17 != nil {
				if err = ie.Lcg_ExtensionIAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lcg_ExtensionIAB_r17", err)
				}
			}
			// encode Harq_FeedbackDisabled_r17 optional
			if ie.Harq_FeedbackDisabled_r17 != nil {
				if err = ie.Harq_FeedbackDisabled_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Harq_FeedbackDisabled_r17", err)
				}
			}
			// encode Uplink_Harq_ModeB_r17 optional
			if ie.Uplink_Harq_ModeB_r17 != nil {
				if err = ie.Uplink_Harq_ModeB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uplink_Harq_ModeB_r17", err)
				}
			}
			// encode Sr_TriggeredBy_TA_Report_r17 optional
			if ie.Sr_TriggeredBy_TA_Report_r17 != nil {
				if err = ie.Sr_TriggeredBy_TA_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sr_TriggeredBy_TA_Report_r17", err)
				}
			}
			// encode ExtendedDRX_CycleInactive_r17 optional
			if ie.ExtendedDRX_CycleInactive_r17 != nil {
				if err = ie.ExtendedDRX_CycleInactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedDRX_CycleInactive_r17", err)
				}
			}
			// encode SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 optional
			if ie.SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil {
				if err = ie.SimultaneousSR_PUSCH_DiffPUCCH_groups_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousSR_PUSCH_DiffPUCCH_groups_r17", err)
				}
			}
			// encode LastTransmissionUL_r17 optional
			if ie.LastTransmissionUL_r17 != nil {
				if err = ie.LastTransmissionUL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LastTransmissionUL_r17", err)
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

func (ie *MAC_ParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Lcp_RestrictionPresent bool
	if Lcp_RestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Lch_ToSCellRestrictionPresent bool
	if Lch_ToSCellRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Lcp_RestrictionPresent {
		ie.Lcp_Restriction = new(MAC_ParametersCommon_lcp_Restriction)
		if err = ie.Lcp_Restriction.Decode(r); err != nil {
			return utils.WrapError("Decode Lcp_Restriction", err)
		}
	}
	if DummyPresent {
		ie.Dummy = new(MAC_ParametersCommon_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	if Lch_ToSCellRestrictionPresent {
		ie.Lch_ToSCellRestriction = new(MAC_ParametersCommon_lch_ToSCellRestriction)
		if err = ie.Lch_ToSCellRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode Lch_ToSCellRestriction", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			RecommendedBitRatePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RecommendedBitRateQueryPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RecommendedBitRate optional
			if RecommendedBitRatePresent {
				ie.RecommendedBitRate = new(MAC_ParametersCommon_recommendedBitRate)
				if err = ie.RecommendedBitRate.Decode(extReader); err != nil {
					return utils.WrapError("Decode RecommendedBitRate", err)
				}
			}
			// decode RecommendedBitRateQuery optional
			if RecommendedBitRateQueryPresent {
				ie.RecommendedBitRateQuery = new(MAC_ParametersCommon_recommendedBitRateQuery)
				if err = ie.RecommendedBitRateQuery.Decode(extReader); err != nil {
					return utils.WrapError("Decode RecommendedBitRateQuery", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			RecommendedBitRateMultiplier_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PreEmptiveBSR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AutonomousTransmission_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lch_PriorityBasedPrioritization_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lch_ToConfiguredGrantMapping_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lch_ToGrantPriorityRestriction_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SinglePHR_P_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_LBT_FailureDetectionRecovery_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tdd_MPE_P_MPR_Reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lcid_ExtensionIAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RecommendedBitRateMultiplier_r16 optional
			if RecommendedBitRateMultiplier_r16Present {
				ie.RecommendedBitRateMultiplier_r16 = new(MAC_ParametersCommon_recommendedBitRateMultiplier_r16)
				if err = ie.RecommendedBitRateMultiplier_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RecommendedBitRateMultiplier_r16", err)
				}
			}
			// decode PreEmptiveBSR_r16 optional
			if PreEmptiveBSR_r16Present {
				ie.PreEmptiveBSR_r16 = new(MAC_ParametersCommon_preEmptiveBSR_r16)
				if err = ie.PreEmptiveBSR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PreEmptiveBSR_r16", err)
				}
			}
			// decode AutonomousTransmission_r16 optional
			if AutonomousTransmission_r16Present {
				ie.AutonomousTransmission_r16 = new(MAC_ParametersCommon_autonomousTransmission_r16)
				if err = ie.AutonomousTransmission_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AutonomousTransmission_r16", err)
				}
			}
			// decode Lch_PriorityBasedPrioritization_r16 optional
			if Lch_PriorityBasedPrioritization_r16Present {
				ie.Lch_PriorityBasedPrioritization_r16 = new(MAC_ParametersCommon_lch_PriorityBasedPrioritization_r16)
				if err = ie.Lch_PriorityBasedPrioritization_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lch_PriorityBasedPrioritization_r16", err)
				}
			}
			// decode Lch_ToConfiguredGrantMapping_r16 optional
			if Lch_ToConfiguredGrantMapping_r16Present {
				ie.Lch_ToConfiguredGrantMapping_r16 = new(MAC_ParametersCommon_lch_ToConfiguredGrantMapping_r16)
				if err = ie.Lch_ToConfiguredGrantMapping_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lch_ToConfiguredGrantMapping_r16", err)
				}
			}
			// decode Lch_ToGrantPriorityRestriction_r16 optional
			if Lch_ToGrantPriorityRestriction_r16Present {
				ie.Lch_ToGrantPriorityRestriction_r16 = new(MAC_ParametersCommon_lch_ToGrantPriorityRestriction_r16)
				if err = ie.Lch_ToGrantPriorityRestriction_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lch_ToGrantPriorityRestriction_r16", err)
				}
			}
			// decode SinglePHR_P_r16 optional
			if SinglePHR_P_r16Present {
				ie.SinglePHR_P_r16 = new(MAC_ParametersCommon_singlePHR_P_r16)
				if err = ie.SinglePHR_P_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SinglePHR_P_r16", err)
				}
			}
			// decode Ul_LBT_FailureDetectionRecovery_r16 optional
			if Ul_LBT_FailureDetectionRecovery_r16Present {
				ie.Ul_LBT_FailureDetectionRecovery_r16 = new(MAC_ParametersCommon_ul_LBT_FailureDetectionRecovery_r16)
				if err = ie.Ul_LBT_FailureDetectionRecovery_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_LBT_FailureDetectionRecovery_r16", err)
				}
			}
			// decode Tdd_MPE_P_MPR_Reporting_r16 optional
			if Tdd_MPE_P_MPR_Reporting_r16Present {
				ie.Tdd_MPE_P_MPR_Reporting_r16 = new(MAC_ParametersCommon_tdd_MPE_P_MPR_Reporting_r16)
				if err = ie.Tdd_MPE_P_MPR_Reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tdd_MPE_P_MPR_Reporting_r16", err)
				}
			}
			// decode Lcid_ExtensionIAB_r16 optional
			if Lcid_ExtensionIAB_r16Present {
				ie.Lcid_ExtensionIAB_r16 = new(MAC_ParametersCommon_lcid_ExtensionIAB_r16)
				if err = ie.Lcid_ExtensionIAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lcid_ExtensionIAB_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			SpCell_BFR_CBRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SpCell_BFR_CBRA_r16 optional
			if SpCell_BFR_CBRA_r16Present {
				ie.SpCell_BFR_CBRA_r16 = new(MAC_ParametersCommon_spCell_BFR_CBRA_r16)
				if err = ie.SpCell_BFR_CBRA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpCell_BFR_CBRA_r16", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Srs_ResourceId_Ext_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Srs_ResourceId_Ext_r16 optional
			if Srs_ResourceId_Ext_r16Present {
				ie.Srs_ResourceId_Ext_r16 = new(MAC_ParametersCommon_srs_ResourceId_Ext_r16)
				if err = ie.Srs_ResourceId_Ext_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_ResourceId_Ext_r16", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			EnhancedUuDRX_forSidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mg_ActivationRequestPRS_Meas_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mg_ActivationCommPRS_Meas_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntraCG_Prioritization_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			JointPrioritizationCG_Retx_Timer_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SurvivalTime_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lcg_ExtensionIAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_FeedbackDisabled_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uplink_Harq_ModeB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sr_TriggeredBy_TA_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExtendedDRX_CycleInactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousSR_PUSCH_DiffPUCCH_groups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LastTransmissionUL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnhancedUuDRX_forSidelink_r17 optional
			if EnhancedUuDRX_forSidelink_r17Present {
				ie.EnhancedUuDRX_forSidelink_r17 = new(MAC_ParametersCommon_enhancedUuDRX_forSidelink_r17)
				if err = ie.EnhancedUuDRX_forSidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedUuDRX_forSidelink_r17", err)
				}
			}
			// decode Mg_ActivationRequestPRS_Meas_r17 optional
			if Mg_ActivationRequestPRS_Meas_r17Present {
				ie.Mg_ActivationRequestPRS_Meas_r17 = new(MAC_ParametersCommon_mg_ActivationRequestPRS_Meas_r17)
				if err = ie.Mg_ActivationRequestPRS_Meas_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mg_ActivationRequestPRS_Meas_r17", err)
				}
			}
			// decode Mg_ActivationCommPRS_Meas_r17 optional
			if Mg_ActivationCommPRS_Meas_r17Present {
				ie.Mg_ActivationCommPRS_Meas_r17 = new(MAC_ParametersCommon_mg_ActivationCommPRS_Meas_r17)
				if err = ie.Mg_ActivationCommPRS_Meas_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mg_ActivationCommPRS_Meas_r17", err)
				}
			}
			// decode IntraCG_Prioritization_r17 optional
			if IntraCG_Prioritization_r17Present {
				ie.IntraCG_Prioritization_r17 = new(MAC_ParametersCommon_intraCG_Prioritization_r17)
				if err = ie.IntraCG_Prioritization_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraCG_Prioritization_r17", err)
				}
			}
			// decode JointPrioritizationCG_Retx_Timer_r17 optional
			if JointPrioritizationCG_Retx_Timer_r17Present {
				ie.JointPrioritizationCG_Retx_Timer_r17 = new(MAC_ParametersCommon_jointPrioritizationCG_Retx_Timer_r17)
				if err = ie.JointPrioritizationCG_Retx_Timer_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode JointPrioritizationCG_Retx_Timer_r17", err)
				}
			}
			// decode SurvivalTime_r17 optional
			if SurvivalTime_r17Present {
				ie.SurvivalTime_r17 = new(MAC_ParametersCommon_survivalTime_r17)
				if err = ie.SurvivalTime_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SurvivalTime_r17", err)
				}
			}
			// decode Lcg_ExtensionIAB_r17 optional
			if Lcg_ExtensionIAB_r17Present {
				ie.Lcg_ExtensionIAB_r17 = new(MAC_ParametersCommon_lcg_ExtensionIAB_r17)
				if err = ie.Lcg_ExtensionIAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lcg_ExtensionIAB_r17", err)
				}
			}
			// decode Harq_FeedbackDisabled_r17 optional
			if Harq_FeedbackDisabled_r17Present {
				ie.Harq_FeedbackDisabled_r17 = new(MAC_ParametersCommon_harq_FeedbackDisabled_r17)
				if err = ie.Harq_FeedbackDisabled_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Harq_FeedbackDisabled_r17", err)
				}
			}
			// decode Uplink_Harq_ModeB_r17 optional
			if Uplink_Harq_ModeB_r17Present {
				ie.Uplink_Harq_ModeB_r17 = new(MAC_ParametersCommon_uplink_Harq_ModeB_r17)
				if err = ie.Uplink_Harq_ModeB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Uplink_Harq_ModeB_r17", err)
				}
			}
			// decode Sr_TriggeredBy_TA_Report_r17 optional
			if Sr_TriggeredBy_TA_Report_r17Present {
				ie.Sr_TriggeredBy_TA_Report_r17 = new(MAC_ParametersCommon_sr_TriggeredBy_TA_Report_r17)
				if err = ie.Sr_TriggeredBy_TA_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sr_TriggeredBy_TA_Report_r17", err)
				}
			}
			// decode ExtendedDRX_CycleInactive_r17 optional
			if ExtendedDRX_CycleInactive_r17Present {
				ie.ExtendedDRX_CycleInactive_r17 = new(MAC_ParametersCommon_extendedDRX_CycleInactive_r17)
				if err = ie.ExtendedDRX_CycleInactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedDRX_CycleInactive_r17", err)
				}
			}
			// decode SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 optional
			if SimultaneousSR_PUSCH_DiffPUCCH_groups_r17Present {
				ie.SimultaneousSR_PUSCH_DiffPUCCH_groups_r17 = new(MAC_ParametersCommon_simultaneousSR_PUSCH_DiffPUCCH_groups_r17)
				if err = ie.SimultaneousSR_PUSCH_DiffPUCCH_groups_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousSR_PUSCH_DiffPUCCH_groups_r17", err)
				}
			}
			// decode LastTransmissionUL_r17 optional
			if LastTransmissionUL_r17Present {
				ie.LastTransmissionUL_r17 = new(MAC_ParametersCommon_lastTransmissionUL_r17)
				if err = ie.LastTransmissionUL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode LastTransmissionUL_r17", err)
				}
			}
		}
	}
	return nil
}
