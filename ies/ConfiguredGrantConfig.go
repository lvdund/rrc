package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfiguredGrantConfig struct {
	FrequencyHopping                  *ConfiguredGrantConfig_frequencyHopping           `optional`
	Cg_DMRS_Configuration             DMRS_UplinkConfig                                 `madatory`
	Mcs_Table                         *ConfiguredGrantConfig_mcs_Table                  `optional`
	Mcs_TableTransformPrecoder        *ConfiguredGrantConfig_mcs_TableTransformPrecoder `optional`
	Uci_OnPUSCH                       *CG_UCI_OnPUSCH                                   `optional,setuprelease`
	ResourceAllocation                ConfiguredGrantConfig_resourceAllocation          `madatory`
	Rbg_Size                          *ConfiguredGrantConfig_rbg_Size                   `optional`
	PowerControlLoopToUse             ConfiguredGrantConfig_powerControlLoopToUse       `madatory`
	P0_PUSCH_Alpha                    P0_PUSCH_AlphaSetId                               `madatory`
	TransformPrecoder                 *ConfiguredGrantConfig_transformPrecoder          `optional`
	NrofHARQ_Processes                int64                                             `lb:1,ub:16,madatory`
	RepK                              ConfiguredGrantConfig_repK                        `madatory`
	RepK_RV                           *ConfiguredGrantConfig_repK_RV                    `optional`
	Periodicity                       ConfiguredGrantConfig_periodicity                 `madatory`
	ConfiguredGrantTimer              *int64                                            `lb:1,ub:64,optional`
	Rrc_ConfiguredUplinkGrant         *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant  `lb:18,ub:18,optional`
	Cg_RetransmissionTimer_r16        *int64                                            `lb:1,ub:64,optional,ext-3`
	Cg_minDFI_Delay_r16               *ConfiguredGrantConfig_cg_minDFI_Delay_r16        `optional,ext-3`
	Cg_nrofPUSCH_InSlot_r16           *int64                                            `lb:1,ub:7,optional,ext-3`
	Cg_nrofSlots_r16                  *int64                                            `lb:1,ub:40,optional,ext-3`
	Cg_StartingOffsets_r16            *CG_StartingOffsets_r16                           `optional,ext-3`
	Cg_UCI_Multiplexing_r16           *ConfiguredGrantConfig_cg_UCI_Multiplexing_r16    `optional,ext-3`
	Cg_COT_SharingOffset_r16          *int64                                            `lb:1,ub:39,optional,ext-3`
	BetaOffsetCG_UCI_r16              *int64                                            `lb:0,ub:31,optional,ext-3`
	Cg_COT_SharingList_r16            []CG_COT_Sharing_r16                              `lb:1,ub:1709,optional,ext-3`
	Harq_ProcID_Offset_r16            *int64                                            `lb:0,ub:15,optional,ext-3`
	Harq_ProcID_Offset2_r16           *int64                                            `lb:0,ub:15,optional,ext-3`
	ConfiguredGrantConfigIndex_r16    *ConfiguredGrantConfigIndex_r16                   `optional,ext-3`
	ConfiguredGrantConfigIndexMAC_r16 *ConfiguredGrantConfigIndexMAC_r16                `optional,ext-3`
	PeriodicityExt_r16                *int64                                            `lb:1,ub:5120,optional,ext-3`
	StartingFromRV0_r16               *ConfiguredGrantConfig_startingFromRV0_r16        `optional,ext-3`
	Phy_PriorityIndex_r16             *ConfiguredGrantConfig_phy_PriorityIndex_r16      `optional,ext-3`
	AutonomousTx_r16                  *ConfiguredGrantConfig_autonomousTx_r16           `optional,ext-3`
	Cg_betaOffsetsCrossPri0_r17       *BetaOffsetsCrossPriSelCG_r17                     `optional,ext-4,setuprelease`
	Cg_betaOffsetsCrossPri1_r17       *BetaOffsetsCrossPriSelCG_r17                     `optional,ext-4,setuprelease`
	MappingPattern_r17                *ConfiguredGrantConfig_mappingPattern_r17         `optional,ext-4`
	SequenceOffsetForRV_r17           *int64                                            `lb:0,ub:3,optional,ext-4`
	P0_PUSCH_Alpha2_r17               *P0_PUSCH_AlphaSetId                              `optional,ext-4`
	PowerControlLoopToUse2_r17        *ConfiguredGrantConfig_powerControlLoopToUse2_r17 `optional,ext-4`
	Cg_COT_SharingList_r17            []CG_COT_Sharing_r17                              `lb:1,ub:50722,optional,ext-4`
	PeriodicityExt_r17                *int64                                            `lb:1,ub:40960,optional,ext-4`
	RepK_v1710                        *ConfiguredGrantConfig_repK_v1710                 `optional,ext-4`
	NrofHARQ_Processes_v1700          *int64                                            `lb:17,ub:32,optional,ext-4`
	Harq_ProcID_Offset2_v1700         *int64                                            `lb:16,ub:31,optional,ext-4`
	ConfiguredGrantTimer_v1700        *int64                                            `lb:33,ub:288,optional,ext-4`
	Cg_minDFI_Delay_v1710             *int64                                            `lb:238,ub:3584,optional,ext-4`
	Harq_ProcID_Offset_v1730          *int64                                            `lb:16,ub:31,optional,ext-5`
	Cg_nrofSlots_r17                  *int64                                            `lb:1,ub:320,optional,ext-5`
}

func (ie *ConfiguredGrantConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Cg_RetransmissionTimer_r16 != nil || ie.Cg_minDFI_Delay_r16 != nil || ie.Cg_nrofPUSCH_InSlot_r16 != nil || ie.Cg_nrofSlots_r16 != nil || ie.Cg_StartingOffsets_r16 != nil || ie.Cg_UCI_Multiplexing_r16 != nil || ie.Cg_COT_SharingOffset_r16 != nil || ie.BetaOffsetCG_UCI_r16 != nil || len(ie.Cg_COT_SharingList_r16) > 0 || ie.Harq_ProcID_Offset_r16 != nil || ie.Harq_ProcID_Offset2_r16 != nil || ie.ConfiguredGrantConfigIndex_r16 != nil || ie.ConfiguredGrantConfigIndexMAC_r16 != nil || ie.PeriodicityExt_r16 != nil || ie.StartingFromRV0_r16 != nil || ie.Phy_PriorityIndex_r16 != nil || ie.AutonomousTx_r16 != nil || ie.Cg_betaOffsetsCrossPri0_r17 != nil || ie.Cg_betaOffsetsCrossPri1_r17 != nil || ie.MappingPattern_r17 != nil || ie.SequenceOffsetForRV_r17 != nil || ie.P0_PUSCH_Alpha2_r17 != nil || ie.PowerControlLoopToUse2_r17 != nil || len(ie.Cg_COT_SharingList_r17) > 0 || ie.PeriodicityExt_r17 != nil || ie.RepK_v1710 != nil || ie.NrofHARQ_Processes_v1700 != nil || ie.Harq_ProcID_Offset2_v1700 != nil || ie.ConfiguredGrantTimer_v1700 != nil || ie.Cg_minDFI_Delay_v1710 != nil || ie.Harq_ProcID_Offset_v1730 != nil || ie.Cg_nrofSlots_r17 != nil
	preambleBits := []bool{hasExtensions, ie.FrequencyHopping != nil, ie.Mcs_Table != nil, ie.Mcs_TableTransformPrecoder != nil, ie.Uci_OnPUSCH != nil, ie.Rbg_Size != nil, ie.TransformPrecoder != nil, ie.RepK_RV != nil, ie.ConfiguredGrantTimer != nil, ie.Rrc_ConfiguredUplinkGrant != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FrequencyHopping != nil {
		if err = ie.FrequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyHopping", err)
		}
	}
	if err = ie.Cg_DMRS_Configuration.Encode(w); err != nil {
		return utils.WrapError("Encode Cg_DMRS_Configuration", err)
	}
	if ie.Mcs_Table != nil {
		if err = ie.Mcs_Table.Encode(w); err != nil {
			return utils.WrapError("Encode Mcs_Table", err)
		}
	}
	if ie.Mcs_TableTransformPrecoder != nil {
		if err = ie.Mcs_TableTransformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode Mcs_TableTransformPrecoder", err)
		}
	}
	if ie.Uci_OnPUSCH != nil {
		tmp_Uci_OnPUSCH := utils.SetupRelease[*CG_UCI_OnPUSCH]{
			Setup: ie.Uci_OnPUSCH,
		}
		if err = tmp_Uci_OnPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Uci_OnPUSCH", err)
		}
	}
	if err = ie.ResourceAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceAllocation", err)
	}
	if ie.Rbg_Size != nil {
		if err = ie.Rbg_Size.Encode(w); err != nil {
			return utils.WrapError("Encode Rbg_Size", err)
		}
	}
	if err = ie.PowerControlLoopToUse.Encode(w); err != nil {
		return utils.WrapError("Encode PowerControlLoopToUse", err)
	}
	if err = ie.P0_PUSCH_Alpha.Encode(w); err != nil {
		return utils.WrapError("Encode P0_PUSCH_Alpha", err)
	}
	if ie.TransformPrecoder != nil {
		if err = ie.TransformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode TransformPrecoder", err)
		}
	}
	if err = w.WriteInteger(ie.NrofHARQ_Processes, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger NrofHARQ_Processes", err)
	}
	if err = ie.RepK.Encode(w); err != nil {
		return utils.WrapError("Encode RepK", err)
	}
	if ie.RepK_RV != nil {
		if err = ie.RepK_RV.Encode(w); err != nil {
			return utils.WrapError("Encode RepK_RV", err)
		}
	}
	if err = ie.Periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode Periodicity", err)
	}
	if ie.ConfiguredGrantTimer != nil {
		if err = w.WriteInteger(*ie.ConfiguredGrantTimer, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode ConfiguredGrantTimer", err)
		}
	}
	if ie.Rrc_ConfiguredUplinkGrant != nil {
		if err = ie.Rrc_ConfiguredUplinkGrant.Encode(w); err != nil {
			return utils.WrapError("Encode Rrc_ConfiguredUplinkGrant", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.Cg_RetransmissionTimer_r16 != nil || ie.Cg_minDFI_Delay_r16 != nil || ie.Cg_nrofPUSCH_InSlot_r16 != nil || ie.Cg_nrofSlots_r16 != nil || ie.Cg_StartingOffsets_r16 != nil || ie.Cg_UCI_Multiplexing_r16 != nil || ie.Cg_COT_SharingOffset_r16 != nil || ie.BetaOffsetCG_UCI_r16 != nil || len(ie.Cg_COT_SharingList_r16) > 0 || ie.Harq_ProcID_Offset_r16 != nil || ie.Harq_ProcID_Offset2_r16 != nil || ie.ConfiguredGrantConfigIndex_r16 != nil || ie.ConfiguredGrantConfigIndexMAC_r16 != nil || ie.PeriodicityExt_r16 != nil || ie.StartingFromRV0_r16 != nil || ie.Phy_PriorityIndex_r16 != nil || ie.AutonomousTx_r16 != nil, ie.Cg_betaOffsetsCrossPri0_r17 != nil || ie.Cg_betaOffsetsCrossPri1_r17 != nil || ie.MappingPattern_r17 != nil || ie.SequenceOffsetForRV_r17 != nil || ie.P0_PUSCH_Alpha2_r17 != nil || ie.PowerControlLoopToUse2_r17 != nil || len(ie.Cg_COT_SharingList_r17) > 0 || ie.PeriodicityExt_r17 != nil || ie.RepK_v1710 != nil || ie.NrofHARQ_Processes_v1700 != nil || ie.Harq_ProcID_Offset2_v1700 != nil || ie.ConfiguredGrantTimer_v1700 != nil || ie.Cg_minDFI_Delay_v1710 != nil, ie.Harq_ProcID_Offset_v1730 != nil || ie.Cg_nrofSlots_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ConfiguredGrantConfig", err)
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.Cg_RetransmissionTimer_r16 != nil, ie.Cg_minDFI_Delay_r16 != nil, ie.Cg_nrofPUSCH_InSlot_r16 != nil, ie.Cg_nrofSlots_r16 != nil, ie.Cg_StartingOffsets_r16 != nil, ie.Cg_UCI_Multiplexing_r16 != nil, ie.Cg_COT_SharingOffset_r16 != nil, ie.BetaOffsetCG_UCI_r16 != nil, len(ie.Cg_COT_SharingList_r16) > 0, ie.Harq_ProcID_Offset_r16 != nil, ie.Harq_ProcID_Offset2_r16 != nil, ie.ConfiguredGrantConfigIndex_r16 != nil, ie.ConfiguredGrantConfigIndexMAC_r16 != nil, ie.PeriodicityExt_r16 != nil, ie.StartingFromRV0_r16 != nil, ie.Phy_PriorityIndex_r16 != nil, ie.AutonomousTx_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Cg_RetransmissionTimer_r16 optional
			if ie.Cg_RetransmissionTimer_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Cg_RetransmissionTimer_r16, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
					return utils.WrapError("Encode Cg_RetransmissionTimer_r16", err)
				}
			}
			// encode Cg_minDFI_Delay_r16 optional
			if ie.Cg_minDFI_Delay_r16 != nil {
				if err = ie.Cg_minDFI_Delay_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_minDFI_Delay_r16", err)
				}
			}
			// encode Cg_nrofPUSCH_InSlot_r16 optional
			if ie.Cg_nrofPUSCH_InSlot_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Cg_nrofPUSCH_InSlot_r16, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
					return utils.WrapError("Encode Cg_nrofPUSCH_InSlot_r16", err)
				}
			}
			// encode Cg_nrofSlots_r16 optional
			if ie.Cg_nrofSlots_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Cg_nrofSlots_r16, &uper.Constraint{Lb: 1, Ub: 40}, false); err != nil {
					return utils.WrapError("Encode Cg_nrofSlots_r16", err)
				}
			}
			// encode Cg_StartingOffsets_r16 optional
			if ie.Cg_StartingOffsets_r16 != nil {
				if err = ie.Cg_StartingOffsets_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_StartingOffsets_r16", err)
				}
			}
			// encode Cg_UCI_Multiplexing_r16 optional
			if ie.Cg_UCI_Multiplexing_r16 != nil {
				if err = ie.Cg_UCI_Multiplexing_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_UCI_Multiplexing_r16", err)
				}
			}
			// encode Cg_COT_SharingOffset_r16 optional
			if ie.Cg_COT_SharingOffset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Cg_COT_SharingOffset_r16, &uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
					return utils.WrapError("Encode Cg_COT_SharingOffset_r16", err)
				}
			}
			// encode BetaOffsetCG_UCI_r16 optional
			if ie.BetaOffsetCG_UCI_r16 != nil {
				if err = extWriter.WriteInteger(*ie.BetaOffsetCG_UCI_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode BetaOffsetCG_UCI_r16", err)
				}
			}
			// encode Cg_COT_SharingList_r16 optional
			if len(ie.Cg_COT_SharingList_r16) > 0 {
				tmp_Cg_COT_SharingList_r16 := utils.NewSequence[*CG_COT_Sharing_r16]([]*CG_COT_Sharing_r16{}, uper.Constraint{Lb: 1, Ub: 1709}, false)
				for _, i := range ie.Cg_COT_SharingList_r16 {
					tmp_Cg_COT_SharingList_r16.Value = append(tmp_Cg_COT_SharingList_r16.Value, &i)
				}
				if err = tmp_Cg_COT_SharingList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_COT_SharingList_r16", err)
				}
			}
			// encode Harq_ProcID_Offset_r16 optional
			if ie.Harq_ProcID_Offset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcID_Offset_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcID_Offset_r16", err)
				}
			}
			// encode Harq_ProcID_Offset2_r16 optional
			if ie.Harq_ProcID_Offset2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcID_Offset2_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcID_Offset2_r16", err)
				}
			}
			// encode ConfiguredGrantConfigIndex_r16 optional
			if ie.ConfiguredGrantConfigIndex_r16 != nil {
				if err = ie.ConfiguredGrantConfigIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfiguredGrantConfigIndex_r16", err)
				}
			}
			// encode ConfiguredGrantConfigIndexMAC_r16 optional
			if ie.ConfiguredGrantConfigIndexMAC_r16 != nil {
				if err = ie.ConfiguredGrantConfigIndexMAC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfiguredGrantConfigIndexMAC_r16", err)
				}
			}
			// encode PeriodicityExt_r16 optional
			if ie.PeriodicityExt_r16 != nil {
				if err = extWriter.WriteInteger(*ie.PeriodicityExt_r16, &uper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Encode PeriodicityExt_r16", err)
				}
			}
			// encode StartingFromRV0_r16 optional
			if ie.StartingFromRV0_r16 != nil {
				if err = ie.StartingFromRV0_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode StartingFromRV0_r16", err)
				}
			}
			// encode Phy_PriorityIndex_r16 optional
			if ie.Phy_PriorityIndex_r16 != nil {
				if err = ie.Phy_PriorityIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Phy_PriorityIndex_r16", err)
				}
			}
			// encode AutonomousTx_r16 optional
			if ie.AutonomousTx_r16 != nil {
				if err = ie.AutonomousTx_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AutonomousTx_r16", err)
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
			optionals_ext_4 := []bool{ie.Cg_betaOffsetsCrossPri0_r17 != nil, ie.Cg_betaOffsetsCrossPri1_r17 != nil, ie.MappingPattern_r17 != nil, ie.SequenceOffsetForRV_r17 != nil, ie.P0_PUSCH_Alpha2_r17 != nil, ie.PowerControlLoopToUse2_r17 != nil, len(ie.Cg_COT_SharingList_r17) > 0, ie.PeriodicityExt_r17 != nil, ie.RepK_v1710 != nil, ie.NrofHARQ_Processes_v1700 != nil, ie.Harq_ProcID_Offset2_v1700 != nil, ie.ConfiguredGrantTimer_v1700 != nil, ie.Cg_minDFI_Delay_v1710 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Cg_betaOffsetsCrossPri0_r17 optional
			if ie.Cg_betaOffsetsCrossPri0_r17 != nil {
				tmp_Cg_betaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{
					Setup: ie.Cg_betaOffsetsCrossPri0_r17,
				}
				if err = tmp_Cg_betaOffsetsCrossPri0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_betaOffsetsCrossPri0_r17", err)
				}
			}
			// encode Cg_betaOffsetsCrossPri1_r17 optional
			if ie.Cg_betaOffsetsCrossPri1_r17 != nil {
				tmp_Cg_betaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{
					Setup: ie.Cg_betaOffsetsCrossPri1_r17,
				}
				if err = tmp_Cg_betaOffsetsCrossPri1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_betaOffsetsCrossPri1_r17", err)
				}
			}
			// encode MappingPattern_r17 optional
			if ie.MappingPattern_r17 != nil {
				if err = ie.MappingPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MappingPattern_r17", err)
				}
			}
			// encode SequenceOffsetForRV_r17 optional
			if ie.SequenceOffsetForRV_r17 != nil {
				if err = extWriter.WriteInteger(*ie.SequenceOffsetForRV_r17, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode SequenceOffsetForRV_r17", err)
				}
			}
			// encode P0_PUSCH_Alpha2_r17 optional
			if ie.P0_PUSCH_Alpha2_r17 != nil {
				if err = ie.P0_PUSCH_Alpha2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode P0_PUSCH_Alpha2_r17", err)
				}
			}
			// encode PowerControlLoopToUse2_r17 optional
			if ie.PowerControlLoopToUse2_r17 != nil {
				if err = ie.PowerControlLoopToUse2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PowerControlLoopToUse2_r17", err)
				}
			}
			// encode Cg_COT_SharingList_r17 optional
			if len(ie.Cg_COT_SharingList_r17) > 0 {
				tmp_Cg_COT_SharingList_r17 := utils.NewSequence[*CG_COT_Sharing_r17]([]*CG_COT_Sharing_r17{}, uper.Constraint{Lb: 1, Ub: 50722}, false)
				for _, i := range ie.Cg_COT_SharingList_r17 {
					tmp_Cg_COT_SharingList_r17.Value = append(tmp_Cg_COT_SharingList_r17.Value, &i)
				}
				if err = tmp_Cg_COT_SharingList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_COT_SharingList_r17", err)
				}
			}
			// encode PeriodicityExt_r17 optional
			if ie.PeriodicityExt_r17 != nil {
				if err = extWriter.WriteInteger(*ie.PeriodicityExt_r17, &uper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Encode PeriodicityExt_r17", err)
				}
			}
			// encode RepK_v1710 optional
			if ie.RepK_v1710 != nil {
				if err = ie.RepK_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RepK_v1710", err)
				}
			}
			// encode NrofHARQ_Processes_v1700 optional
			if ie.NrofHARQ_Processes_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.NrofHARQ_Processes_v1700, &uper.Constraint{Lb: 17, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode NrofHARQ_Processes_v1700", err)
				}
			}
			// encode Harq_ProcID_Offset2_v1700 optional
			if ie.Harq_ProcID_Offset2_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcID_Offset2_v1700, &uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcID_Offset2_v1700", err)
				}
			}
			// encode ConfiguredGrantTimer_v1700 optional
			if ie.ConfiguredGrantTimer_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.ConfiguredGrantTimer_v1700, &uper.Constraint{Lb: 33, Ub: 288}, false); err != nil {
					return utils.WrapError("Encode ConfiguredGrantTimer_v1700", err)
				}
			}
			// encode Cg_minDFI_Delay_v1710 optional
			if ie.Cg_minDFI_Delay_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.Cg_minDFI_Delay_v1710, &uper.Constraint{Lb: 238, Ub: 3584}, false); err != nil {
					return utils.WrapError("Encode Cg_minDFI_Delay_v1710", err)
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
			optionals_ext_5 := []bool{ie.Harq_ProcID_Offset_v1730 != nil, ie.Cg_nrofSlots_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Harq_ProcID_Offset_v1730 optional
			if ie.Harq_ProcID_Offset_v1730 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcID_Offset_v1730, &uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcID_Offset_v1730", err)
				}
			}
			// encode Cg_nrofSlots_r17 optional
			if ie.Cg_nrofSlots_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Cg_nrofSlots_r17, &uper.Constraint{Lb: 1, Ub: 320}, false); err != nil {
					return utils.WrapError("Encode Cg_nrofSlots_r17", err)
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

func (ie *ConfiguredGrantConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyHoppingPresent bool
	if FrequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mcs_TablePresent bool
	if Mcs_TablePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mcs_TableTransformPrecoderPresent bool
	if Mcs_TableTransformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Uci_OnPUSCHPresent bool
	if Uci_OnPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rbg_SizePresent bool
	if Rbg_SizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TransformPrecoderPresent bool
	if TransformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RepK_RVPresent bool
	if RepK_RVPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredGrantTimerPresent bool
	if ConfiguredGrantTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rrc_ConfiguredUplinkGrantPresent bool
	if Rrc_ConfiguredUplinkGrantPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyHoppingPresent {
		ie.FrequencyHopping = new(ConfiguredGrantConfig_frequencyHopping)
		if err = ie.FrequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyHopping", err)
		}
	}
	if err = ie.Cg_DMRS_Configuration.Decode(r); err != nil {
		return utils.WrapError("Decode Cg_DMRS_Configuration", err)
	}
	if Mcs_TablePresent {
		ie.Mcs_Table = new(ConfiguredGrantConfig_mcs_Table)
		if err = ie.Mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode Mcs_Table", err)
		}
	}
	if Mcs_TableTransformPrecoderPresent {
		ie.Mcs_TableTransformPrecoder = new(ConfiguredGrantConfig_mcs_TableTransformPrecoder)
		if err = ie.Mcs_TableTransformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode Mcs_TableTransformPrecoder", err)
		}
	}
	if Uci_OnPUSCHPresent {
		tmp_Uci_OnPUSCH := utils.SetupRelease[*CG_UCI_OnPUSCH]{}
		if err = tmp_Uci_OnPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Uci_OnPUSCH", err)
		}
		ie.Uci_OnPUSCH = tmp_Uci_OnPUSCH.Setup
	}
	if err = ie.ResourceAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceAllocation", err)
	}
	if Rbg_SizePresent {
		ie.Rbg_Size = new(ConfiguredGrantConfig_rbg_Size)
		if err = ie.Rbg_Size.Decode(r); err != nil {
			return utils.WrapError("Decode Rbg_Size", err)
		}
	}
	if err = ie.PowerControlLoopToUse.Decode(r); err != nil {
		return utils.WrapError("Decode PowerControlLoopToUse", err)
	}
	if err = ie.P0_PUSCH_Alpha.Decode(r); err != nil {
		return utils.WrapError("Decode P0_PUSCH_Alpha", err)
	}
	if TransformPrecoderPresent {
		ie.TransformPrecoder = new(ConfiguredGrantConfig_transformPrecoder)
		if err = ie.TransformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode TransformPrecoder", err)
		}
	}
	var tmp_int_NrofHARQ_Processes int64
	if tmp_int_NrofHARQ_Processes, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger NrofHARQ_Processes", err)
	}
	ie.NrofHARQ_Processes = tmp_int_NrofHARQ_Processes
	if err = ie.RepK.Decode(r); err != nil {
		return utils.WrapError("Decode RepK", err)
	}
	if RepK_RVPresent {
		ie.RepK_RV = new(ConfiguredGrantConfig_repK_RV)
		if err = ie.RepK_RV.Decode(r); err != nil {
			return utils.WrapError("Decode RepK_RV", err)
		}
	}
	if err = ie.Periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode Periodicity", err)
	}
	if ConfiguredGrantTimerPresent {
		var tmp_int_ConfiguredGrantTimer int64
		if tmp_int_ConfiguredGrantTimer, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode ConfiguredGrantTimer", err)
		}
		ie.ConfiguredGrantTimer = &tmp_int_ConfiguredGrantTimer
	}
	if Rrc_ConfiguredUplinkGrantPresent {
		ie.Rrc_ConfiguredUplinkGrant = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant)
		if err = ie.Rrc_ConfiguredUplinkGrant.Decode(r); err != nil {
			return utils.WrapError("Decode Rrc_ConfiguredUplinkGrant", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Cg_RetransmissionTimer_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_minDFI_Delay_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_nrofPUSCH_InSlot_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_nrofSlots_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_StartingOffsets_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_UCI_Multiplexing_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_COT_SharingOffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BetaOffsetCG_UCI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_COT_SharingList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcID_Offset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcID_Offset2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredGrantConfigIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredGrantConfigIndexMAC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PeriodicityExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			StartingFromRV0_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Phy_PriorityIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AutonomousTx_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Cg_RetransmissionTimer_r16 optional
			if Cg_RetransmissionTimer_r16Present {
				var tmp_int_Cg_RetransmissionTimer_r16 int64
				if tmp_int_Cg_RetransmissionTimer_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
					return utils.WrapError("Decode Cg_RetransmissionTimer_r16", err)
				}
				ie.Cg_RetransmissionTimer_r16 = &tmp_int_Cg_RetransmissionTimer_r16
			}
			// decode Cg_minDFI_Delay_r16 optional
			if Cg_minDFI_Delay_r16Present {
				ie.Cg_minDFI_Delay_r16 = new(ConfiguredGrantConfig_cg_minDFI_Delay_r16)
				if err = ie.Cg_minDFI_Delay_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cg_minDFI_Delay_r16", err)
				}
			}
			// decode Cg_nrofPUSCH_InSlot_r16 optional
			if Cg_nrofPUSCH_InSlot_r16Present {
				var tmp_int_Cg_nrofPUSCH_InSlot_r16 int64
				if tmp_int_Cg_nrofPUSCH_InSlot_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
					return utils.WrapError("Decode Cg_nrofPUSCH_InSlot_r16", err)
				}
				ie.Cg_nrofPUSCH_InSlot_r16 = &tmp_int_Cg_nrofPUSCH_InSlot_r16
			}
			// decode Cg_nrofSlots_r16 optional
			if Cg_nrofSlots_r16Present {
				var tmp_int_Cg_nrofSlots_r16 int64
				if tmp_int_Cg_nrofSlots_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 40}, false); err != nil {
					return utils.WrapError("Decode Cg_nrofSlots_r16", err)
				}
				ie.Cg_nrofSlots_r16 = &tmp_int_Cg_nrofSlots_r16
			}
			// decode Cg_StartingOffsets_r16 optional
			if Cg_StartingOffsets_r16Present {
				ie.Cg_StartingOffsets_r16 = new(CG_StartingOffsets_r16)
				if err = ie.Cg_StartingOffsets_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cg_StartingOffsets_r16", err)
				}
			}
			// decode Cg_UCI_Multiplexing_r16 optional
			if Cg_UCI_Multiplexing_r16Present {
				ie.Cg_UCI_Multiplexing_r16 = new(ConfiguredGrantConfig_cg_UCI_Multiplexing_r16)
				if err = ie.Cg_UCI_Multiplexing_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cg_UCI_Multiplexing_r16", err)
				}
			}
			// decode Cg_COT_SharingOffset_r16 optional
			if Cg_COT_SharingOffset_r16Present {
				var tmp_int_Cg_COT_SharingOffset_r16 int64
				if tmp_int_Cg_COT_SharingOffset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
					return utils.WrapError("Decode Cg_COT_SharingOffset_r16", err)
				}
				ie.Cg_COT_SharingOffset_r16 = &tmp_int_Cg_COT_SharingOffset_r16
			}
			// decode BetaOffsetCG_UCI_r16 optional
			if BetaOffsetCG_UCI_r16Present {
				var tmp_int_BetaOffsetCG_UCI_r16 int64
				if tmp_int_BetaOffsetCG_UCI_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode BetaOffsetCG_UCI_r16", err)
				}
				ie.BetaOffsetCG_UCI_r16 = &tmp_int_BetaOffsetCG_UCI_r16
			}
			// decode Cg_COT_SharingList_r16 optional
			if Cg_COT_SharingList_r16Present {
				tmp_Cg_COT_SharingList_r16 := utils.NewSequence[*CG_COT_Sharing_r16]([]*CG_COT_Sharing_r16{}, uper.Constraint{Lb: 1, Ub: 1709}, false)
				fn_Cg_COT_SharingList_r16 := func() *CG_COT_Sharing_r16 {
					return new(CG_COT_Sharing_r16)
				}
				if err = tmp_Cg_COT_SharingList_r16.Decode(extReader, fn_Cg_COT_SharingList_r16); err != nil {
					return utils.WrapError("Decode Cg_COT_SharingList_r16", err)
				}
				ie.Cg_COT_SharingList_r16 = []CG_COT_Sharing_r16{}
				for _, i := range tmp_Cg_COT_SharingList_r16.Value {
					ie.Cg_COT_SharingList_r16 = append(ie.Cg_COT_SharingList_r16, *i)
				}
			}
			// decode Harq_ProcID_Offset_r16 optional
			if Harq_ProcID_Offset_r16Present {
				var tmp_int_Harq_ProcID_Offset_r16 int64
				if tmp_int_Harq_ProcID_Offset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcID_Offset_r16", err)
				}
				ie.Harq_ProcID_Offset_r16 = &tmp_int_Harq_ProcID_Offset_r16
			}
			// decode Harq_ProcID_Offset2_r16 optional
			if Harq_ProcID_Offset2_r16Present {
				var tmp_int_Harq_ProcID_Offset2_r16 int64
				if tmp_int_Harq_ProcID_Offset2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcID_Offset2_r16", err)
				}
				ie.Harq_ProcID_Offset2_r16 = &tmp_int_Harq_ProcID_Offset2_r16
			}
			// decode ConfiguredGrantConfigIndex_r16 optional
			if ConfiguredGrantConfigIndex_r16Present {
				ie.ConfiguredGrantConfigIndex_r16 = new(ConfiguredGrantConfigIndex_r16)
				if err = ie.ConfiguredGrantConfigIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfiguredGrantConfigIndex_r16", err)
				}
			}
			// decode ConfiguredGrantConfigIndexMAC_r16 optional
			if ConfiguredGrantConfigIndexMAC_r16Present {
				ie.ConfiguredGrantConfigIndexMAC_r16 = new(ConfiguredGrantConfigIndexMAC_r16)
				if err = ie.ConfiguredGrantConfigIndexMAC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfiguredGrantConfigIndexMAC_r16", err)
				}
			}
			// decode PeriodicityExt_r16 optional
			if PeriodicityExt_r16Present {
				var tmp_int_PeriodicityExt_r16 int64
				if tmp_int_PeriodicityExt_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Decode PeriodicityExt_r16", err)
				}
				ie.PeriodicityExt_r16 = &tmp_int_PeriodicityExt_r16
			}
			// decode StartingFromRV0_r16 optional
			if StartingFromRV0_r16Present {
				ie.StartingFromRV0_r16 = new(ConfiguredGrantConfig_startingFromRV0_r16)
				if err = ie.StartingFromRV0_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode StartingFromRV0_r16", err)
				}
			}
			// decode Phy_PriorityIndex_r16 optional
			if Phy_PriorityIndex_r16Present {
				ie.Phy_PriorityIndex_r16 = new(ConfiguredGrantConfig_phy_PriorityIndex_r16)
				if err = ie.Phy_PriorityIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Phy_PriorityIndex_r16", err)
				}
			}
			// decode AutonomousTx_r16 optional
			if AutonomousTx_r16Present {
				ie.AutonomousTx_r16 = new(ConfiguredGrantConfig_autonomousTx_r16)
				if err = ie.AutonomousTx_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AutonomousTx_r16", err)
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

			Cg_betaOffsetsCrossPri0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_betaOffsetsCrossPri1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MappingPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SequenceOffsetForRV_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			P0_PUSCH_Alpha2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PowerControlLoopToUse2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_COT_SharingList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PeriodicityExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RepK_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NrofHARQ_Processes_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcID_Offset2_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredGrantTimer_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_minDFI_Delay_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Cg_betaOffsetsCrossPri0_r17 optional
			if Cg_betaOffsetsCrossPri0_r17Present {
				tmp_Cg_betaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{}
				if err = tmp_Cg_betaOffsetsCrossPri0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cg_betaOffsetsCrossPri0_r17", err)
				}
				ie.Cg_betaOffsetsCrossPri0_r17 = tmp_Cg_betaOffsetsCrossPri0_r17.Setup
			}
			// decode Cg_betaOffsetsCrossPri1_r17 optional
			if Cg_betaOffsetsCrossPri1_r17Present {
				tmp_Cg_betaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{}
				if err = tmp_Cg_betaOffsetsCrossPri1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cg_betaOffsetsCrossPri1_r17", err)
				}
				ie.Cg_betaOffsetsCrossPri1_r17 = tmp_Cg_betaOffsetsCrossPri1_r17.Setup
			}
			// decode MappingPattern_r17 optional
			if MappingPattern_r17Present {
				ie.MappingPattern_r17 = new(ConfiguredGrantConfig_mappingPattern_r17)
				if err = ie.MappingPattern_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MappingPattern_r17", err)
				}
			}
			// decode SequenceOffsetForRV_r17 optional
			if SequenceOffsetForRV_r17Present {
				var tmp_int_SequenceOffsetForRV_r17 int64
				if tmp_int_SequenceOffsetForRV_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Decode SequenceOffsetForRV_r17", err)
				}
				ie.SequenceOffsetForRV_r17 = &tmp_int_SequenceOffsetForRV_r17
			}
			// decode P0_PUSCH_Alpha2_r17 optional
			if P0_PUSCH_Alpha2_r17Present {
				ie.P0_PUSCH_Alpha2_r17 = new(P0_PUSCH_AlphaSetId)
				if err = ie.P0_PUSCH_Alpha2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode P0_PUSCH_Alpha2_r17", err)
				}
			}
			// decode PowerControlLoopToUse2_r17 optional
			if PowerControlLoopToUse2_r17Present {
				ie.PowerControlLoopToUse2_r17 = new(ConfiguredGrantConfig_powerControlLoopToUse2_r17)
				if err = ie.PowerControlLoopToUse2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PowerControlLoopToUse2_r17", err)
				}
			}
			// decode Cg_COT_SharingList_r17 optional
			if Cg_COT_SharingList_r17Present {
				tmp_Cg_COT_SharingList_r17 := utils.NewSequence[*CG_COT_Sharing_r17]([]*CG_COT_Sharing_r17{}, uper.Constraint{Lb: 1, Ub: 50722}, false)
				fn_Cg_COT_SharingList_r17 := func() *CG_COT_Sharing_r17 {
					return new(CG_COT_Sharing_r17)
				}
				if err = tmp_Cg_COT_SharingList_r17.Decode(extReader, fn_Cg_COT_SharingList_r17); err != nil {
					return utils.WrapError("Decode Cg_COT_SharingList_r17", err)
				}
				ie.Cg_COT_SharingList_r17 = []CG_COT_Sharing_r17{}
				for _, i := range tmp_Cg_COT_SharingList_r17.Value {
					ie.Cg_COT_SharingList_r17 = append(ie.Cg_COT_SharingList_r17, *i)
				}
			}
			// decode PeriodicityExt_r17 optional
			if PeriodicityExt_r17Present {
				var tmp_int_PeriodicityExt_r17 int64
				if tmp_int_PeriodicityExt_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Decode PeriodicityExt_r17", err)
				}
				ie.PeriodicityExt_r17 = &tmp_int_PeriodicityExt_r17
			}
			// decode RepK_v1710 optional
			if RepK_v1710Present {
				ie.RepK_v1710 = new(ConfiguredGrantConfig_repK_v1710)
				if err = ie.RepK_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode RepK_v1710", err)
				}
			}
			// decode NrofHARQ_Processes_v1700 optional
			if NrofHARQ_Processes_v1700Present {
				var tmp_int_NrofHARQ_Processes_v1700 int64
				if tmp_int_NrofHARQ_Processes_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 17, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode NrofHARQ_Processes_v1700", err)
				}
				ie.NrofHARQ_Processes_v1700 = &tmp_int_NrofHARQ_Processes_v1700
			}
			// decode Harq_ProcID_Offset2_v1700 optional
			if Harq_ProcID_Offset2_v1700Present {
				var tmp_int_Harq_ProcID_Offset2_v1700 int64
				if tmp_int_Harq_ProcID_Offset2_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcID_Offset2_v1700", err)
				}
				ie.Harq_ProcID_Offset2_v1700 = &tmp_int_Harq_ProcID_Offset2_v1700
			}
			// decode ConfiguredGrantTimer_v1700 optional
			if ConfiguredGrantTimer_v1700Present {
				var tmp_int_ConfiguredGrantTimer_v1700 int64
				if tmp_int_ConfiguredGrantTimer_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 33, Ub: 288}, false); err != nil {
					return utils.WrapError("Decode ConfiguredGrantTimer_v1700", err)
				}
				ie.ConfiguredGrantTimer_v1700 = &tmp_int_ConfiguredGrantTimer_v1700
			}
			// decode Cg_minDFI_Delay_v1710 optional
			if Cg_minDFI_Delay_v1710Present {
				var tmp_int_Cg_minDFI_Delay_v1710 int64
				if tmp_int_Cg_minDFI_Delay_v1710, err = extReader.ReadInteger(&uper.Constraint{Lb: 238, Ub: 3584}, false); err != nil {
					return utils.WrapError("Decode Cg_minDFI_Delay_v1710", err)
				}
				ie.Cg_minDFI_Delay_v1710 = &tmp_int_Cg_minDFI_Delay_v1710
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Harq_ProcID_Offset_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_nrofSlots_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Harq_ProcID_Offset_v1730 optional
			if Harq_ProcID_Offset_v1730Present {
				var tmp_int_Harq_ProcID_Offset_v1730 int64
				if tmp_int_Harq_ProcID_Offset_v1730, err = extReader.ReadInteger(&uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcID_Offset_v1730", err)
				}
				ie.Harq_ProcID_Offset_v1730 = &tmp_int_Harq_ProcID_Offset_v1730
			}
			// decode Cg_nrofSlots_r17 optional
			if Cg_nrofSlots_r17Present {
				var tmp_int_Cg_nrofSlots_r17 int64
				if tmp_int_Cg_nrofSlots_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 320}, false); err != nil {
					return utils.WrapError("Decode Cg_nrofSlots_r17", err)
				}
				ie.Cg_nrofSlots_r17 = &tmp_int_Cg_nrofSlots_r17
			}
		}
	}
	return nil
}
