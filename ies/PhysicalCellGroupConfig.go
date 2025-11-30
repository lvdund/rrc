package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PhysicalCellGroupConfig struct {
	Harq_ACK_SpatialBundlingPUCCH                           *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH                           `optional`
	Harq_ACK_SpatialBundlingPUSCH                           *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH                           `optional`
	P_NR_FR1                                                *P_Max                                                                           `optional`
	Pdsch_HARQ_ACK_Codebook                                 PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook                                  `madatory`
	Tpc_SRS_RNTI                                            *RNTI_Value                                                                      `optional`
	Tpc_PUCCH_RNTI                                          *RNTI_Value                                                                      `optional`
	Tpc_PUSCH_RNTI                                          *RNTI_Value                                                                      `optional`
	Sp_CSI_RNTI                                             *RNTI_Value                                                                      `optional`
	Cs_RNTI                                                 *RNTI_Value                                                                      `optional,setuprelease`
	Mcs_C_RNTI                                              *RNTI_Value                                                                      `optional,ext-1`
	P_UE_FR1                                                *P_Max                                                                           `optional,ext-1`
	XScale                                                  *PhysicalCellGroupConfig_xScale                                                  `optional,ext-2`
	Pdcch_BlindDetection                                    *PDCCH_BlindDetection                                                            `optional,ext-3,setuprelease`
	Dcp_Config_r16                                          *DCP_Config_r16                                                                  `optional,ext-4,setuprelease`
	Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16   *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16   `optional,ext-4`
	Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16   *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16   `optional,ext-4`
	Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16         *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16         `optional,ext-4`
	P_NR_FR2_r16                                            *P_Max                                                                           `optional,ext-4`
	P_UE_FR2_r16                                            *P_Max                                                                           `optional,ext-4`
	Nrdc_PCmode_FR1_r16                                     *PhysicalCellGroupConfig_nrdc_PCmode_FR1_r16                                     `optional,ext-4`
	Nrdc_PCmode_FR2_r16                                     *PhysicalCellGroupConfig_nrdc_PCmode_FR2_r16                                     `optional,ext-4`
	Pdsch_HARQ_ACK_Codebook_r16                             *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_r16                             `optional,ext-4`
	Nfi_TotalDAI_Included_r16                               *PhysicalCellGroupConfig_nfi_TotalDAI_Included_r16                               `optional,ext-4`
	Ul_TotalDAI_Included_r16                                *PhysicalCellGroupConfig_ul_TotalDAI_Included_r16                                `optional,ext-4`
	Pdsch_HARQ_ACK_OneShotFeedback_r16                      *PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedback_r16                      `optional,ext-4`
	Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16                   *PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackNDI_r16                   `optional,ext-4`
	Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16                   *PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackCBG_r16                   `optional,ext-4`
	DownlinkAssignmentIndexDCI_0_2_r16                      *PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_0_2_r16                      `optional,ext-4`
	DownlinkAssignmentIndexDCI_1_2_r16                      *PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16                      `optional,ext-4`
	Pdsch_HARQ_ACK_CodebookList_r16                         *PDSCH_HARQ_ACK_CodebookList_r16                                                 `optional,ext-4,setuprelease`
	AckNackFeedbackMode_r16                                 *PhysicalCellGroupConfig_ackNackFeedbackMode_r16                                 `optional,ext-4`
	Pdcch_BlindDetectionCA_CombIndicator_r16                *PDCCH_BlindDetectionCA_CombIndicator_r16                                        `optional,ext-4,setuprelease`
	Pdcch_BlindDetection2_r16                               *PDCCH_BlindDetection2_r16                                                       `optional,ext-4,setuprelease`
	Pdcch_BlindDetection3_r16                               *PDCCH_BlindDetection3_r16                                                       `optional,ext-4,setuprelease`
	BdFactorR_r16                                           *PhysicalCellGroupConfig_bdFactorR_r16                                           `optional,ext-4`
	Pdsch_HARQ_ACK_EnhType3ToAddModList_r17                 []PDSCH_HARQ_ACK_EnhType3_r17                                                    `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17                []PDSCH_HARQ_ACK_EnhType3Index_r17                                               `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17        []PDSCH_HARQ_ACK_EnhType3_r17                                                    `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17       []PDSCH_HARQ_ACK_EnhType3Index_r17                                               `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 *PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 `optional,ext-5`
	Pdsch_HARQ_ACK_EnhType3DCI_Field_r17                    *PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_Field_r17                    `optional,ext-5`
	Pdsch_HARQ_ACK_Retx_r17                                 *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Retx_r17                                 `optional,ext-5`
	Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17              *PhysicalCellGroupConfig_pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17              `optional,ext-5`
	Pucch_sSCell_r17                                        *SCellIndex                                                                      `optional,ext-5`
	Pucch_sSCellSecondaryPUCCHgroup_r17                     *SCellIndex                                                                      `optional,ext-5`
	Pucch_sSCellDyn_r17                                     *PhysicalCellGroupConfig_pucch_sSCellDyn_r17                                     `optional,ext-5`
	Pucch_sSCellDynSecondaryPUCCHgroup_r17                  *PhysicalCellGroupConfig_pucch_sSCellDynSecondaryPUCCHgroup_r17                  `optional,ext-5`
	Pucch_sSCellPattern_r17                                 []int64                                                                          `lb:1,ub:maxNrofSlots,e_lb:0,e_ub:1,optional,ext-5`
	Pucch_sSCellPatternSecondaryPUCCHgroup_r17              []int64                                                                          `lb:1,ub:maxNrofSlots,e_lb:0,e_ub:1,optional,ext-5`
	Uci_MuxWithDiffPrio_r17                                 *PhysicalCellGroupConfig_uci_MuxWithDiffPrio_r17                                 `optional,ext-5`
	Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17              *PhysicalCellGroupConfig_uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17              `optional,ext-5`
	SimultaneousPUCCH_PUSCH_r17                             *PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_r17                             `optional,ext-5`
	SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17         *PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17         `optional,ext-5`
	PrioLowDG_HighCG_r17                                    *PhysicalCellGroupConfig_prioLowDG_HighCG_r17                                    `optional,ext-5`
	PrioHighDG_LowCG_r17                                    *PhysicalCellGroupConfig_prioHighDG_LowCG_r17                                    `optional,ext-5`
	TwoQCLTypeDforPDCCHRepetition_r17                       *PhysicalCellGroupConfig_twoQCLTypeDforPDCCHRepetition_r17                       `optional,ext-5`
	MulticastConfig_r17                                     *MulticastConfig_r17                                                             `optional,ext-5,setuprelease`
	Pdcch_BlindDetectionCA_CombIndicator_r17                *PDCCH_BlindDetectionCA_CombIndicator_r17                                        `optional,ext-5,setuprelease`
	SimultaneousSR_PUSCH_diffPUCCH_Groups_r17               *PhysicalCellGroupConfig_simultaneousSR_PUSCH_diffPUCCH_Groups_r17               `optional,ext-6`
	IntraBandNC_PRACH_simulTx_r17                           *PhysicalCellGroupConfig_intraBandNC_PRACH_simulTx_r17                           `optional,ext-7`
}

func (ie *PhysicalCellGroupConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Mcs_C_RNTI != nil || ie.P_UE_FR1 != nil || ie.XScale != nil || ie.Pdcch_BlindDetection != nil || ie.Dcp_Config_r16 != nil || ie.Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil || ie.Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil || ie.Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil || ie.P_NR_FR2_r16 != nil || ie.P_UE_FR2_r16 != nil || ie.Nrdc_PCmode_FR1_r16 != nil || ie.Nrdc_PCmode_FR2_r16 != nil || ie.Pdsch_HARQ_ACK_Codebook_r16 != nil || ie.Nfi_TotalDAI_Included_r16 != nil || ie.Ul_TotalDAI_Included_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil || ie.DownlinkAssignmentIndexDCI_0_2_r16 != nil || ie.DownlinkAssignmentIndexDCI_1_2_r16 != nil || ie.Pdsch_HARQ_ACK_CodebookList_r16 != nil || ie.AckNackFeedbackMode_r16 != nil || ie.Pdcch_BlindDetectionCA_CombIndicator_r16 != nil || ie.Pdcch_BlindDetection2_r16 != nil || ie.Pdcch_BlindDetection3_r16 != nil || ie.BdFactorR_r16 != nil || len(ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0 || len(ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0 || len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0 || len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0 || ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil || ie.Pdsch_HARQ_ACK_Retx_r17 != nil || ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil || ie.Pucch_sSCell_r17 != nil || ie.Pucch_sSCellSecondaryPUCCHgroup_r17 != nil || ie.Pucch_sSCellDyn_r17 != nil || ie.Pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil || len(ie.Pucch_sSCellPattern_r17) > 0 || len(ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0 || ie.Uci_MuxWithDiffPrio_r17 != nil || ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil || ie.SimultaneousPUCCH_PUSCH_r17 != nil || ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil || ie.PrioLowDG_HighCG_r17 != nil || ie.PrioHighDG_LowCG_r17 != nil || ie.TwoQCLTypeDforPDCCHRepetition_r17 != nil || ie.MulticastConfig_r17 != nil || ie.Pdcch_BlindDetectionCA_CombIndicator_r17 != nil || ie.SimultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil || ie.IntraBandNC_PRACH_simulTx_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Harq_ACK_SpatialBundlingPUCCH != nil, ie.Harq_ACK_SpatialBundlingPUSCH != nil, ie.P_NR_FR1 != nil, ie.Tpc_SRS_RNTI != nil, ie.Tpc_PUCCH_RNTI != nil, ie.Tpc_PUSCH_RNTI != nil, ie.Sp_CSI_RNTI != nil, ie.Cs_RNTI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Harq_ACK_SpatialBundlingPUCCH != nil {
		if err = ie.Harq_ACK_SpatialBundlingPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode Harq_ACK_SpatialBundlingPUCCH", err)
		}
	}
	if ie.Harq_ACK_SpatialBundlingPUSCH != nil {
		if err = ie.Harq_ACK_SpatialBundlingPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Harq_ACK_SpatialBundlingPUSCH", err)
		}
	}
	if ie.P_NR_FR1 != nil {
		if err = ie.P_NR_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode P_NR_FR1", err)
		}
	}
	if err = ie.Pdsch_HARQ_ACK_Codebook.Encode(w); err != nil {
		return utils.WrapError("Encode Pdsch_HARQ_ACK_Codebook", err)
	}
	if ie.Tpc_SRS_RNTI != nil {
		if err = ie.Tpc_SRS_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_SRS_RNTI", err)
		}
	}
	if ie.Tpc_PUCCH_RNTI != nil {
		if err = ie.Tpc_PUCCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_PUCCH_RNTI", err)
		}
	}
	if ie.Tpc_PUSCH_RNTI != nil {
		if err = ie.Tpc_PUSCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_PUSCH_RNTI", err)
		}
	}
	if ie.Sp_CSI_RNTI != nil {
		if err = ie.Sp_CSI_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_RNTI", err)
		}
	}
	if ie.Cs_RNTI != nil {
		tmp_Cs_RNTI := utils.SetupRelease[*RNTI_Value]{
			Setup: ie.Cs_RNTI,
		}
		if err = tmp_Cs_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Cs_RNTI", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 7 bits for 7 extension groups
		extBitmap := []bool{ie.Mcs_C_RNTI != nil || ie.P_UE_FR1 != nil, ie.XScale != nil, ie.Pdcch_BlindDetection != nil, ie.Dcp_Config_r16 != nil || ie.Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil || ie.Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil || ie.Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil || ie.P_NR_FR2_r16 != nil || ie.P_UE_FR2_r16 != nil || ie.Nrdc_PCmode_FR1_r16 != nil || ie.Nrdc_PCmode_FR2_r16 != nil || ie.Pdsch_HARQ_ACK_Codebook_r16 != nil || ie.Nfi_TotalDAI_Included_r16 != nil || ie.Ul_TotalDAI_Included_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil || ie.DownlinkAssignmentIndexDCI_0_2_r16 != nil || ie.DownlinkAssignmentIndexDCI_1_2_r16 != nil || ie.Pdsch_HARQ_ACK_CodebookList_r16 != nil || ie.AckNackFeedbackMode_r16 != nil || ie.Pdcch_BlindDetectionCA_CombIndicator_r16 != nil || ie.Pdcch_BlindDetection2_r16 != nil || ie.Pdcch_BlindDetection3_r16 != nil || ie.BdFactorR_r16 != nil, len(ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0 || len(ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0 || len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0 || len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0 || ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil || ie.Pdsch_HARQ_ACK_Retx_r17 != nil || ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil || ie.Pucch_sSCell_r17 != nil || ie.Pucch_sSCellSecondaryPUCCHgroup_r17 != nil || ie.Pucch_sSCellDyn_r17 != nil || ie.Pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil || len(ie.Pucch_sSCellPattern_r17) > 0 || len(ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0 || ie.Uci_MuxWithDiffPrio_r17 != nil || ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil || ie.SimultaneousPUCCH_PUSCH_r17 != nil || ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil || ie.PrioLowDG_HighCG_r17 != nil || ie.PrioHighDG_LowCG_r17 != nil || ie.TwoQCLTypeDforPDCCHRepetition_r17 != nil || ie.MulticastConfig_r17 != nil || ie.Pdcch_BlindDetectionCA_CombIndicator_r17 != nil, ie.SimultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil, ie.IntraBandNC_PRACH_simulTx_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PhysicalCellGroupConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Mcs_C_RNTI != nil, ie.P_UE_FR1 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Mcs_C_RNTI optional
			if ie.Mcs_C_RNTI != nil {
				if err = ie.Mcs_C_RNTI.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mcs_C_RNTI", err)
				}
			}
			// encode P_UE_FR1 optional
			if ie.P_UE_FR1 != nil {
				if err = ie.P_UE_FR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode P_UE_FR1", err)
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
			optionals_ext_2 := []bool{ie.XScale != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode XScale optional
			if ie.XScale != nil {
				if err = ie.XScale.Encode(extWriter); err != nil {
					return utils.WrapError("Encode XScale", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.Pdcch_BlindDetection != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdcch_BlindDetection optional
			if ie.Pdcch_BlindDetection != nil {
				tmp_Pdcch_BlindDetection := utils.SetupRelease[*PDCCH_BlindDetection]{
					Setup: ie.Pdcch_BlindDetection,
				}
				if err = tmp_Pdcch_BlindDetection.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_BlindDetection", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.Dcp_Config_r16 != nil, ie.Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil, ie.Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil, ie.Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil, ie.P_NR_FR2_r16 != nil, ie.P_UE_FR2_r16 != nil, ie.Nrdc_PCmode_FR1_r16 != nil, ie.Nrdc_PCmode_FR2_r16 != nil, ie.Pdsch_HARQ_ACK_Codebook_r16 != nil, ie.Nfi_TotalDAI_Included_r16 != nil, ie.Ul_TotalDAI_Included_r16 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil, ie.DownlinkAssignmentIndexDCI_0_2_r16 != nil, ie.DownlinkAssignmentIndexDCI_1_2_r16 != nil, ie.Pdsch_HARQ_ACK_CodebookList_r16 != nil, ie.AckNackFeedbackMode_r16 != nil, ie.Pdcch_BlindDetectionCA_CombIndicator_r16 != nil, ie.Pdcch_BlindDetection2_r16 != nil, ie.Pdcch_BlindDetection3_r16 != nil, ie.BdFactorR_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dcp_Config_r16 optional
			if ie.Dcp_Config_r16 != nil {
				tmp_Dcp_Config_r16 := utils.SetupRelease[*DCP_Config_r16]{
					Setup: ie.Dcp_Config_r16,
				}
				if err = tmp_Dcp_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dcp_Config_r16", err)
				}
			}
			// encode Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 optional
			if ie.Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil {
				if err = ie.Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// encode Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 optional
			if ie.Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil {
				if err = ie.Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// encode Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 optional
			if ie.Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil {
				if err = ie.Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16", err)
				}
			}
			// encode P_NR_FR2_r16 optional
			if ie.P_NR_FR2_r16 != nil {
				if err = ie.P_NR_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode P_NR_FR2_r16", err)
				}
			}
			// encode P_UE_FR2_r16 optional
			if ie.P_UE_FR2_r16 != nil {
				if err = ie.P_UE_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode P_UE_FR2_r16", err)
				}
			}
			// encode Nrdc_PCmode_FR1_r16 optional
			if ie.Nrdc_PCmode_FR1_r16 != nil {
				if err = ie.Nrdc_PCmode_FR1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nrdc_PCmode_FR1_r16", err)
				}
			}
			// encode Nrdc_PCmode_FR2_r16 optional
			if ie.Nrdc_PCmode_FR2_r16 != nil {
				if err = ie.Nrdc_PCmode_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nrdc_PCmode_FR2_r16", err)
				}
			}
			// encode Pdsch_HARQ_ACK_Codebook_r16 optional
			if ie.Pdsch_HARQ_ACK_Codebook_r16 != nil {
				if err = ie.Pdsch_HARQ_ACK_Codebook_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_Codebook_r16", err)
				}
			}
			// encode Nfi_TotalDAI_Included_r16 optional
			if ie.Nfi_TotalDAI_Included_r16 != nil {
				if err = ie.Nfi_TotalDAI_Included_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nfi_TotalDAI_Included_r16", err)
				}
			}
			// encode Ul_TotalDAI_Included_r16 optional
			if ie.Ul_TotalDAI_Included_r16 != nil {
				if err = ie.Ul_TotalDAI_Included_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_TotalDAI_Included_r16", err)
				}
			}
			// encode Pdsch_HARQ_ACK_OneShotFeedback_r16 optional
			if ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 != nil {
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedback_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_OneShotFeedback_r16", err)
				}
			}
			// encode Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 optional
			if ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil {
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16", err)
				}
			}
			// encode Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 optional
			if ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil {
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16", err)
				}
			}
			// encode DownlinkAssignmentIndexDCI_0_2_r16 optional
			if ie.DownlinkAssignmentIndexDCI_0_2_r16 != nil {
				if err = ie.DownlinkAssignmentIndexDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DownlinkAssignmentIndexDCI_0_2_r16", err)
				}
			}
			// encode DownlinkAssignmentIndexDCI_1_2_r16 optional
			if ie.DownlinkAssignmentIndexDCI_1_2_r16 != nil {
				if err = ie.DownlinkAssignmentIndexDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DownlinkAssignmentIndexDCI_1_2_r16", err)
				}
			}
			// encode Pdsch_HARQ_ACK_CodebookList_r16 optional
			if ie.Pdsch_HARQ_ACK_CodebookList_r16 != nil {
				tmp_Pdsch_HARQ_ACK_CodebookList_r16 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{
					Setup: ie.Pdsch_HARQ_ACK_CodebookList_r16,
				}
				if err = tmp_Pdsch_HARQ_ACK_CodebookList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_CodebookList_r16", err)
				}
			}
			// encode AckNackFeedbackMode_r16 optional
			if ie.AckNackFeedbackMode_r16 != nil {
				if err = ie.AckNackFeedbackMode_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AckNackFeedbackMode_r16", err)
				}
			}
			// encode Pdcch_BlindDetectionCA_CombIndicator_r16 optional
			if ie.Pdcch_BlindDetectionCA_CombIndicator_r16 != nil {
				tmp_Pdcch_BlindDetectionCA_CombIndicator_r16 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r16]{
					Setup: ie.Pdcch_BlindDetectionCA_CombIndicator_r16,
				}
				if err = tmp_Pdcch_BlindDetectionCA_CombIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_BlindDetectionCA_CombIndicator_r16", err)
				}
			}
			// encode Pdcch_BlindDetection2_r16 optional
			if ie.Pdcch_BlindDetection2_r16 != nil {
				tmp_Pdcch_BlindDetection2_r16 := utils.SetupRelease[*PDCCH_BlindDetection2_r16]{
					Setup: ie.Pdcch_BlindDetection2_r16,
				}
				if err = tmp_Pdcch_BlindDetection2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_BlindDetection2_r16", err)
				}
			}
			// encode Pdcch_BlindDetection3_r16 optional
			if ie.Pdcch_BlindDetection3_r16 != nil {
				tmp_Pdcch_BlindDetection3_r16 := utils.SetupRelease[*PDCCH_BlindDetection3_r16]{
					Setup: ie.Pdcch_BlindDetection3_r16,
				}
				if err = tmp_Pdcch_BlindDetection3_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_BlindDetection3_r16", err)
				}
			}
			// encode BdFactorR_r16 optional
			if ie.BdFactorR_r16 != nil {
				if err = ie.BdFactorR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BdFactorR_r16", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{len(ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0, len(ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0, len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0, len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0, ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil, ie.Pdsch_HARQ_ACK_Retx_r17 != nil, ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil, ie.Pucch_sSCell_r17 != nil, ie.Pucch_sSCellSecondaryPUCCHgroup_r17 != nil, ie.Pucch_sSCellDyn_r17 != nil, ie.Pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil, len(ie.Pucch_sSCellPattern_r17) > 0, len(ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0, ie.Uci_MuxWithDiffPrio_r17 != nil, ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil, ie.SimultaneousPUCCH_PUSCH_r17 != nil, ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil, ie.PrioLowDG_HighCG_r17 != nil, ie.PrioHighDG_LowCG_r17 != nil, ie.TwoQCLTypeDforPDCCHRepetition_r17 != nil, ie.MulticastConfig_r17 != nil, ie.Pdcch_BlindDetectionCA_CombIndicator_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 optional
			if len(ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0 {
				tmp_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 {
					tmp_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Value = append(tmp_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Value, &i)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3ToAddModList_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 optional
			if len(ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0 {
				tmp_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 {
					tmp_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Value = append(tmp_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Value, &i)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 optional
			if len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0 {
				tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 {
					tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Value = append(tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Value, &i)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 optional
			if len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0 {
				tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 {
					tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Value = append(tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Value, &i)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 optional
			if ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 optional
			if ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3DCI_Field_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_Retx_r17 optional
			if ie.Pdsch_HARQ_ACK_Retx_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_Retx_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_Retx_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 optional
			if ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode Pucch_sSCell_r17 optional
			if ie.Pucch_sSCell_r17 != nil {
				if err = ie.Pucch_sSCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_sSCell_r17", err)
				}
			}
			// encode Pucch_sSCellSecondaryPUCCHgroup_r17 optional
			if ie.Pucch_sSCellSecondaryPUCCHgroup_r17 != nil {
				if err = ie.Pucch_sSCellSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_sSCellSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode Pucch_sSCellDyn_r17 optional
			if ie.Pucch_sSCellDyn_r17 != nil {
				if err = ie.Pucch_sSCellDyn_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_sSCellDyn_r17", err)
				}
			}
			// encode Pucch_sSCellDynSecondaryPUCCHgroup_r17 optional
			if ie.Pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil {
				if err = ie.Pucch_sSCellDynSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_sSCellDynSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode Pucch_sSCellPattern_r17 optional
			if len(ie.Pucch_sSCellPattern_r17) > 0 {
				tmp_Pucch_sSCellPattern_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				for _, i := range ie.Pucch_sSCellPattern_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 1}, false)
					tmp_Pucch_sSCellPattern_r17.Value = append(tmp_Pucch_sSCellPattern_r17.Value, &tmp_ie)
				}
				if err = tmp_Pucch_sSCellPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_sSCellPattern_r17", err)
				}
			}
			// encode Pucch_sSCellPatternSecondaryPUCCHgroup_r17 optional
			if len(ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0 {
				tmp_Pucch_sSCellPatternSecondaryPUCCHgroup_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				for _, i := range ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 1}, false)
					tmp_Pucch_sSCellPatternSecondaryPUCCHgroup_r17.Value = append(tmp_Pucch_sSCellPatternSecondaryPUCCHgroup_r17.Value, &tmp_ie)
				}
				if err = tmp_Pucch_sSCellPatternSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_sSCellPatternSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode Uci_MuxWithDiffPrio_r17 optional
			if ie.Uci_MuxWithDiffPrio_r17 != nil {
				if err = ie.Uci_MuxWithDiffPrio_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uci_MuxWithDiffPrio_r17", err)
				}
			}
			// encode Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 optional
			if ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil {
				if err = ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode SimultaneousPUCCH_PUSCH_r17 optional
			if ie.SimultaneousPUCCH_PUSCH_r17 != nil {
				if err = ie.SimultaneousPUCCH_PUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousPUCCH_PUSCH_r17", err)
				}
			}
			// encode SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 optional
			if ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil {
				if err = ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17", err)
				}
			}
			// encode PrioLowDG_HighCG_r17 optional
			if ie.PrioLowDG_HighCG_r17 != nil {
				if err = ie.PrioLowDG_HighCG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PrioLowDG_HighCG_r17", err)
				}
			}
			// encode PrioHighDG_LowCG_r17 optional
			if ie.PrioHighDG_LowCG_r17 != nil {
				if err = ie.PrioHighDG_LowCG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PrioHighDG_LowCG_r17", err)
				}
			}
			// encode TwoQCLTypeDforPDCCHRepetition_r17 optional
			if ie.TwoQCLTypeDforPDCCHRepetition_r17 != nil {
				if err = ie.TwoQCLTypeDforPDCCHRepetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoQCLTypeDforPDCCHRepetition_r17", err)
				}
			}
			// encode MulticastConfig_r17 optional
			if ie.MulticastConfig_r17 != nil {
				tmp_MulticastConfig_r17 := utils.SetupRelease[*MulticastConfig_r17]{
					Setup: ie.MulticastConfig_r17,
				}
				if err = tmp_MulticastConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MulticastConfig_r17", err)
				}
			}
			// encode Pdcch_BlindDetectionCA_CombIndicator_r17 optional
			if ie.Pdcch_BlindDetectionCA_CombIndicator_r17 != nil {
				tmp_Pdcch_BlindDetectionCA_CombIndicator_r17 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r17]{
					Setup: ie.Pdcch_BlindDetectionCA_CombIndicator_r17,
				}
				if err = tmp_Pdcch_BlindDetectionCA_CombIndicator_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_BlindDetectionCA_CombIndicator_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.SimultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SimultaneousSR_PUSCH_diffPUCCH_Groups_r17 optional
			if ie.SimultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil {
				if err = ie.SimultaneousSR_PUSCH_diffPUCCH_Groups_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousSR_PUSCH_diffPUCCH_Groups_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{ie.IntraBandNC_PRACH_simulTx_r17 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode IntraBandNC_PRACH_simulTx_r17 optional
			if ie.IntraBandNC_PRACH_simulTx_r17 != nil {
				if err = ie.IntraBandNC_PRACH_simulTx_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraBandNC_PRACH_simulTx_r17", err)
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

func (ie *PhysicalCellGroupConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Harq_ACK_SpatialBundlingPUCCHPresent bool
	if Harq_ACK_SpatialBundlingPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Harq_ACK_SpatialBundlingPUSCHPresent bool
	if Harq_ACK_SpatialBundlingPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P_NR_FR1Present bool
	if P_NR_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_SRS_RNTIPresent bool
	if Tpc_SRS_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_PUCCH_RNTIPresent bool
	if Tpc_PUCCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_PUSCH_RNTIPresent bool
	if Tpc_PUSCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_RNTIPresent bool
	if Sp_CSI_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Cs_RNTIPresent bool
	if Cs_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Harq_ACK_SpatialBundlingPUCCHPresent {
		ie.Harq_ACK_SpatialBundlingPUCCH = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH)
		if err = ie.Harq_ACK_SpatialBundlingPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode Harq_ACK_SpatialBundlingPUCCH", err)
		}
	}
	if Harq_ACK_SpatialBundlingPUSCHPresent {
		ie.Harq_ACK_SpatialBundlingPUSCH = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH)
		if err = ie.Harq_ACK_SpatialBundlingPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Harq_ACK_SpatialBundlingPUSCH", err)
		}
	}
	if P_NR_FR1Present {
		ie.P_NR_FR1 = new(P_Max)
		if err = ie.P_NR_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode P_NR_FR1", err)
		}
	}
	if err = ie.Pdsch_HARQ_ACK_Codebook.Decode(r); err != nil {
		return utils.WrapError("Decode Pdsch_HARQ_ACK_Codebook", err)
	}
	if Tpc_SRS_RNTIPresent {
		ie.Tpc_SRS_RNTI = new(RNTI_Value)
		if err = ie.Tpc_SRS_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_SRS_RNTI", err)
		}
	}
	if Tpc_PUCCH_RNTIPresent {
		ie.Tpc_PUCCH_RNTI = new(RNTI_Value)
		if err = ie.Tpc_PUCCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_PUCCH_RNTI", err)
		}
	}
	if Tpc_PUSCH_RNTIPresent {
		ie.Tpc_PUSCH_RNTI = new(RNTI_Value)
		if err = ie.Tpc_PUSCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_PUSCH_RNTI", err)
		}
	}
	if Sp_CSI_RNTIPresent {
		ie.Sp_CSI_RNTI = new(RNTI_Value)
		if err = ie.Sp_CSI_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_RNTI", err)
		}
	}
	if Cs_RNTIPresent {
		tmp_Cs_RNTI := utils.SetupRelease[*RNTI_Value]{}
		if err = tmp_Cs_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Cs_RNTI", err)
		}
		ie.Cs_RNTI = tmp_Cs_RNTI.Setup
	}

	if extensionBit {
		// Read extension bitmap: 7 bits for 7 extension groups
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

			Mcs_C_RNTIPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			P_UE_FR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Mcs_C_RNTI optional
			if Mcs_C_RNTIPresent {
				ie.Mcs_C_RNTI = new(RNTI_Value)
				if err = ie.Mcs_C_RNTI.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mcs_C_RNTI", err)
				}
			}
			// decode P_UE_FR1 optional
			if P_UE_FR1Present {
				ie.P_UE_FR1 = new(P_Max)
				if err = ie.P_UE_FR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode P_UE_FR1", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			XScalePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode XScale optional
			if XScalePresent {
				ie.XScale = new(PhysicalCellGroupConfig_xScale)
				if err = ie.XScale.Decode(extReader); err != nil {
					return utils.WrapError("Decode XScale", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Pdcch_BlindDetectionPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdcch_BlindDetection optional
			if Pdcch_BlindDetectionPresent {
				tmp_Pdcch_BlindDetection := utils.SetupRelease[*PDCCH_BlindDetection]{}
				if err = tmp_Pdcch_BlindDetection.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_BlindDetection", err)
				}
				ie.Pdcch_BlindDetection = tmp_Pdcch_BlindDetection.Setup
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Dcp_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			P_NR_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			P_UE_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nrdc_PCmode_FR1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nrdc_PCmode_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_Codebook_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nfi_TotalDAI_Included_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_TotalDAI_Included_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_OneShotFeedback_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DownlinkAssignmentIndexDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DownlinkAssignmentIndexDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_CodebookList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AckNackFeedbackMode_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_BlindDetectionCA_CombIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_BlindDetection2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_BlindDetection3_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BdFactorR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dcp_Config_r16 optional
			if Dcp_Config_r16Present {
				tmp_Dcp_Config_r16 := utils.SetupRelease[*DCP_Config_r16]{}
				if err = tmp_Dcp_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dcp_Config_r16", err)
				}
				ie.Dcp_Config_r16 = tmp_Dcp_Config_r16.Setup
			}
			// decode Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 optional
			if Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16Present {
				ie.Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16)
				if err = ie.Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// decode Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 optional
			if Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16Present {
				ie.Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16)
				if err = ie.Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// decode Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 optional
			if Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16Present {
				ie.Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16)
				if err = ie.Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16", err)
				}
			}
			// decode P_NR_FR2_r16 optional
			if P_NR_FR2_r16Present {
				ie.P_NR_FR2_r16 = new(P_Max)
				if err = ie.P_NR_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode P_NR_FR2_r16", err)
				}
			}
			// decode P_UE_FR2_r16 optional
			if P_UE_FR2_r16Present {
				ie.P_UE_FR2_r16 = new(P_Max)
				if err = ie.P_UE_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode P_UE_FR2_r16", err)
				}
			}
			// decode Nrdc_PCmode_FR1_r16 optional
			if Nrdc_PCmode_FR1_r16Present {
				ie.Nrdc_PCmode_FR1_r16 = new(PhysicalCellGroupConfig_nrdc_PCmode_FR1_r16)
				if err = ie.Nrdc_PCmode_FR1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nrdc_PCmode_FR1_r16", err)
				}
			}
			// decode Nrdc_PCmode_FR2_r16 optional
			if Nrdc_PCmode_FR2_r16Present {
				ie.Nrdc_PCmode_FR2_r16 = new(PhysicalCellGroupConfig_nrdc_PCmode_FR2_r16)
				if err = ie.Nrdc_PCmode_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nrdc_PCmode_FR2_r16", err)
				}
			}
			// decode Pdsch_HARQ_ACK_Codebook_r16 optional
			if Pdsch_HARQ_ACK_Codebook_r16Present {
				ie.Pdsch_HARQ_ACK_Codebook_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_r16)
				if err = ie.Pdsch_HARQ_ACK_Codebook_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_Codebook_r16", err)
				}
			}
			// decode Nfi_TotalDAI_Included_r16 optional
			if Nfi_TotalDAI_Included_r16Present {
				ie.Nfi_TotalDAI_Included_r16 = new(PhysicalCellGroupConfig_nfi_TotalDAI_Included_r16)
				if err = ie.Nfi_TotalDAI_Included_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nfi_TotalDAI_Included_r16", err)
				}
			}
			// decode Ul_TotalDAI_Included_r16 optional
			if Ul_TotalDAI_Included_r16Present {
				ie.Ul_TotalDAI_Included_r16 = new(PhysicalCellGroupConfig_ul_TotalDAI_Included_r16)
				if err = ie.Ul_TotalDAI_Included_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_TotalDAI_Included_r16", err)
				}
			}
			// decode Pdsch_HARQ_ACK_OneShotFeedback_r16 optional
			if Pdsch_HARQ_ACK_OneShotFeedback_r16Present {
				ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedback_r16)
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedback_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_OneShotFeedback_r16", err)
				}
			}
			// decode Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 optional
			if Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16Present {
				ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackNDI_r16)
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16", err)
				}
			}
			// decode Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 optional
			if Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16Present {
				ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackCBG_r16)
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16", err)
				}
			}
			// decode DownlinkAssignmentIndexDCI_0_2_r16 optional
			if DownlinkAssignmentIndexDCI_0_2_r16Present {
				ie.DownlinkAssignmentIndexDCI_0_2_r16 = new(PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_0_2_r16)
				if err = ie.DownlinkAssignmentIndexDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DownlinkAssignmentIndexDCI_0_2_r16", err)
				}
			}
			// decode DownlinkAssignmentIndexDCI_1_2_r16 optional
			if DownlinkAssignmentIndexDCI_1_2_r16Present {
				ie.DownlinkAssignmentIndexDCI_1_2_r16 = new(PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16)
				if err = ie.DownlinkAssignmentIndexDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DownlinkAssignmentIndexDCI_1_2_r16", err)
				}
			}
			// decode Pdsch_HARQ_ACK_CodebookList_r16 optional
			if Pdsch_HARQ_ACK_CodebookList_r16Present {
				tmp_Pdsch_HARQ_ACK_CodebookList_r16 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{}
				if err = tmp_Pdsch_HARQ_ACK_CodebookList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_CodebookList_r16", err)
				}
				ie.Pdsch_HARQ_ACK_CodebookList_r16 = tmp_Pdsch_HARQ_ACK_CodebookList_r16.Setup
			}
			// decode AckNackFeedbackMode_r16 optional
			if AckNackFeedbackMode_r16Present {
				ie.AckNackFeedbackMode_r16 = new(PhysicalCellGroupConfig_ackNackFeedbackMode_r16)
				if err = ie.AckNackFeedbackMode_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AckNackFeedbackMode_r16", err)
				}
			}
			// decode Pdcch_BlindDetectionCA_CombIndicator_r16 optional
			if Pdcch_BlindDetectionCA_CombIndicator_r16Present {
				tmp_Pdcch_BlindDetectionCA_CombIndicator_r16 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r16]{}
				if err = tmp_Pdcch_BlindDetectionCA_CombIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_BlindDetectionCA_CombIndicator_r16", err)
				}
				ie.Pdcch_BlindDetectionCA_CombIndicator_r16 = tmp_Pdcch_BlindDetectionCA_CombIndicator_r16.Setup
			}
			// decode Pdcch_BlindDetection2_r16 optional
			if Pdcch_BlindDetection2_r16Present {
				tmp_Pdcch_BlindDetection2_r16 := utils.SetupRelease[*PDCCH_BlindDetection2_r16]{}
				if err = tmp_Pdcch_BlindDetection2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_BlindDetection2_r16", err)
				}
				ie.Pdcch_BlindDetection2_r16 = tmp_Pdcch_BlindDetection2_r16.Setup
			}
			// decode Pdcch_BlindDetection3_r16 optional
			if Pdcch_BlindDetection3_r16Present {
				tmp_Pdcch_BlindDetection3_r16 := utils.SetupRelease[*PDCCH_BlindDetection3_r16]{}
				if err = tmp_Pdcch_BlindDetection3_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_BlindDetection3_r16", err)
				}
				ie.Pdcch_BlindDetection3_r16 = tmp_Pdcch_BlindDetection3_r16.Setup
			}
			// decode BdFactorR_r16 optional
			if BdFactorR_r16Present {
				ie.BdFactorR_r16 = new(PhysicalCellGroupConfig_bdFactorR_r16)
				if err = ie.BdFactorR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode BdFactorR_r16", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Pdsch_HARQ_ACK_EnhType3ToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_EnhType3DCI_Field_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_Retx_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_sSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_sSCellSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_sSCellDyn_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_sSCellDynSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_sSCellPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_sSCellPatternSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uci_MuxWithDiffPrio_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousPUCCH_PUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PrioLowDG_HighCG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PrioHighDG_LowCG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TwoQCLTypeDforPDCCHRepetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MulticastConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_BlindDetectionCA_CombIndicator_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 optional
			if Pdsch_HARQ_ACK_EnhType3ToAddModList_r17Present {
				tmp_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 := func() *PDSCH_HARQ_ACK_EnhType3_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3_r17)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Decode(extReader, fn_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3ToAddModList_r17", err)
				}
				ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 = []PDSCH_HARQ_ACK_EnhType3_r17{}
				for _, i := range tmp_Pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Value {
					ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 = append(ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17, *i)
				}
			}
			// decode Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 optional
			if Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17Present {
				tmp_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 := func() *PDSCH_HARQ_ACK_EnhType3Index_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3Index_r17)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Decode(extReader, fn_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17", err)
				}
				ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 = []PDSCH_HARQ_ACK_EnhType3Index_r17{}
				for _, i := range tmp_Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Value {
					ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 = append(ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17, *i)
				}
			}
			// decode Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 optional
			if Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17Present {
				tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 := func() *PDSCH_HARQ_ACK_EnhType3_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3_r17)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Decode(extReader, fn_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17", err)
				}
				ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 = []PDSCH_HARQ_ACK_EnhType3_r17{}
				for _, i := range tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Value {
					ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 = append(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17, *i)
				}
			}
			// decode Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 optional
			if Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17Present {
				tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 := func() *PDSCH_HARQ_ACK_EnhType3Index_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3Index_r17)
				}
				if err = tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Decode(extReader, fn_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17", err)
				}
				ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 = []PDSCH_HARQ_ACK_EnhType3Index_r17{}
				for _, i := range tmp_Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Value {
					ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 = append(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17, *i)
				}
			}
			// decode Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 optional
			if Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17Present {
				ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17)
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 optional
			if Pdsch_HARQ_ACK_EnhType3DCI_Field_r17Present {
				ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_Field_r17)
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3DCI_Field_r17", err)
				}
			}
			// decode Pdsch_HARQ_ACK_Retx_r17 optional
			if Pdsch_HARQ_ACK_Retx_r17Present {
				ie.Pdsch_HARQ_ACK_Retx_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_Retx_r17)
				if err = ie.Pdsch_HARQ_ACK_Retx_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_Retx_r17", err)
				}
			}
			// decode Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 optional
			if Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17Present {
				ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17)
				if err = ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode Pucch_sSCell_r17 optional
			if Pucch_sSCell_r17Present {
				ie.Pucch_sSCell_r17 = new(SCellIndex)
				if err = ie.Pucch_sSCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_sSCell_r17", err)
				}
			}
			// decode Pucch_sSCellSecondaryPUCCHgroup_r17 optional
			if Pucch_sSCellSecondaryPUCCHgroup_r17Present {
				ie.Pucch_sSCellSecondaryPUCCHgroup_r17 = new(SCellIndex)
				if err = ie.Pucch_sSCellSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_sSCellSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode Pucch_sSCellDyn_r17 optional
			if Pucch_sSCellDyn_r17Present {
				ie.Pucch_sSCellDyn_r17 = new(PhysicalCellGroupConfig_pucch_sSCellDyn_r17)
				if err = ie.Pucch_sSCellDyn_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_sSCellDyn_r17", err)
				}
			}
			// decode Pucch_sSCellDynSecondaryPUCCHgroup_r17 optional
			if Pucch_sSCellDynSecondaryPUCCHgroup_r17Present {
				ie.Pucch_sSCellDynSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_pucch_sSCellDynSecondaryPUCCHgroup_r17)
				if err = ie.Pucch_sSCellDynSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_sSCellDynSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode Pucch_sSCellPattern_r17 optional
			if Pucch_sSCellPattern_r17Present {
				tmp_Pucch_sSCellPattern_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				fn_Pucch_sSCellPattern_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 1}, false)
					return &ie
				}
				if err = tmp_Pucch_sSCellPattern_r17.Decode(extReader, fn_Pucch_sSCellPattern_r17); err != nil {
					return utils.WrapError("Decode Pucch_sSCellPattern_r17", err)
				}
				ie.Pucch_sSCellPattern_r17 = []int64{}
				for _, i := range tmp_Pucch_sSCellPattern_r17.Value {
					ie.Pucch_sSCellPattern_r17 = append(ie.Pucch_sSCellPattern_r17, int64(i.Value))
				}
			}
			// decode Pucch_sSCellPatternSecondaryPUCCHgroup_r17 optional
			if Pucch_sSCellPatternSecondaryPUCCHgroup_r17Present {
				tmp_Pucch_sSCellPatternSecondaryPUCCHgroup_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				fn_Pucch_sSCellPatternSecondaryPUCCHgroup_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 1}, false)
					return &ie
				}
				if err = tmp_Pucch_sSCellPatternSecondaryPUCCHgroup_r17.Decode(extReader, fn_Pucch_sSCellPatternSecondaryPUCCHgroup_r17); err != nil {
					return utils.WrapError("Decode Pucch_sSCellPatternSecondaryPUCCHgroup_r17", err)
				}
				ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17 = []int64{}
				for _, i := range tmp_Pucch_sSCellPatternSecondaryPUCCHgroup_r17.Value {
					ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17 = append(ie.Pucch_sSCellPatternSecondaryPUCCHgroup_r17, int64(i.Value))
				}
			}
			// decode Uci_MuxWithDiffPrio_r17 optional
			if Uci_MuxWithDiffPrio_r17Present {
				ie.Uci_MuxWithDiffPrio_r17 = new(PhysicalCellGroupConfig_uci_MuxWithDiffPrio_r17)
				if err = ie.Uci_MuxWithDiffPrio_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Uci_MuxWithDiffPrio_r17", err)
				}
			}
			// decode Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 optional
			if Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17Present {
				ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17)
				if err = ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode SimultaneousPUCCH_PUSCH_r17 optional
			if SimultaneousPUCCH_PUSCH_r17Present {
				ie.SimultaneousPUCCH_PUSCH_r17 = new(PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_r17)
				if err = ie.SimultaneousPUCCH_PUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousPUCCH_PUSCH_r17", err)
				}
			}
			// decode SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 optional
			if SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17Present {
				ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17)
				if err = ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17", err)
				}
			}
			// decode PrioLowDG_HighCG_r17 optional
			if PrioLowDG_HighCG_r17Present {
				ie.PrioLowDG_HighCG_r17 = new(PhysicalCellGroupConfig_prioLowDG_HighCG_r17)
				if err = ie.PrioLowDG_HighCG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PrioLowDG_HighCG_r17", err)
				}
			}
			// decode PrioHighDG_LowCG_r17 optional
			if PrioHighDG_LowCG_r17Present {
				ie.PrioHighDG_LowCG_r17 = new(PhysicalCellGroupConfig_prioHighDG_LowCG_r17)
				if err = ie.PrioHighDG_LowCG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PrioHighDG_LowCG_r17", err)
				}
			}
			// decode TwoQCLTypeDforPDCCHRepetition_r17 optional
			if TwoQCLTypeDforPDCCHRepetition_r17Present {
				ie.TwoQCLTypeDforPDCCHRepetition_r17 = new(PhysicalCellGroupConfig_twoQCLTypeDforPDCCHRepetition_r17)
				if err = ie.TwoQCLTypeDforPDCCHRepetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoQCLTypeDforPDCCHRepetition_r17", err)
				}
			}
			// decode MulticastConfig_r17 optional
			if MulticastConfig_r17Present {
				tmp_MulticastConfig_r17 := utils.SetupRelease[*MulticastConfig_r17]{}
				if err = tmp_MulticastConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MulticastConfig_r17", err)
				}
				ie.MulticastConfig_r17 = tmp_MulticastConfig_r17.Setup
			}
			// decode Pdcch_BlindDetectionCA_CombIndicator_r17 optional
			if Pdcch_BlindDetectionCA_CombIndicator_r17Present {
				tmp_Pdcch_BlindDetectionCA_CombIndicator_r17 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r17]{}
				if err = tmp_Pdcch_BlindDetectionCA_CombIndicator_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_BlindDetectionCA_CombIndicator_r17", err)
				}
				ie.Pdcch_BlindDetectionCA_CombIndicator_r17 = tmp_Pdcch_BlindDetectionCA_CombIndicator_r17.Setup
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SimultaneousSR_PUSCH_diffPUCCH_Groups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SimultaneousSR_PUSCH_diffPUCCH_Groups_r17 optional
			if SimultaneousSR_PUSCH_diffPUCCH_Groups_r17Present {
				ie.SimultaneousSR_PUSCH_diffPUCCH_Groups_r17 = new(PhysicalCellGroupConfig_simultaneousSR_PUSCH_diffPUCCH_Groups_r17)
				if err = ie.SimultaneousSR_PUSCH_diffPUCCH_Groups_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousSR_PUSCH_diffPUCCH_Groups_r17", err)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			IntraBandNC_PRACH_simulTx_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode IntraBandNC_PRACH_simulTx_r17 optional
			if IntraBandNC_PRACH_simulTx_r17Present {
				ie.IntraBandNC_PRACH_simulTx_r17 = new(PhysicalCellGroupConfig_intraBandNC_PRACH_simulTx_r17)
				if err = ie.IntraBandNC_PRACH_simulTx_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraBandNC_PRACH_simulTx_r17", err)
				}
			}
		}
	}
	return nil
}
