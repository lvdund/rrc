package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfig struct {
	Tdd_UL_DL_ConfigurationDedicated              *TDD_UL_DL_ConfigDedicated                                       `optional`
	InitialDownlinkBWP                            *BWP_DownlinkDedicated                                           `optional`
	DownlinkBWP_ToReleaseList                     []BWP_Id                                                         `lb:1,ub:maxNrofBWPs,optional`
	DownlinkBWP_ToAddModList                      []BWP_Downlink                                                   `lb:1,ub:maxNrofBWPs,optional`
	FirstActiveDownlinkBWP_Id                     *BWP_Id                                                          `optional`
	Bwp_InactivityTimer                           *ServingCellConfig_bwp_InactivityTimer                           `optional`
	DefaultDownlinkBWP_Id                         *BWP_Id                                                          `optional`
	UplinkConfig                                  *UplinkConfig                                                    `optional`
	SupplementaryUplink                           *UplinkConfig                                                    `optional`
	Pdcch_ServingCellConfig                       *PDCCH_ServingCellConfig                                         `optional,setuprelease`
	Pdsch_ServingCellConfig                       *PDSCH_ServingCellConfig                                         `optional,setuprelease`
	Csi_MeasConfig                                *CSI_MeasConfig                                                  `optional,setuprelease`
	SCellDeactivationTimer                        *ServingCellConfig_sCellDeactivationTimer                        `optional`
	CrossCarrierSchedulingConfig                  *CrossCarrierSchedulingConfig                                    `optional`
	Tag_Id                                        TAG_Id                                                           `madatory`
	Dummy1                                        *ServingCellConfig_dummy1                                        `optional`
	PathlossReferenceLinking                      *ServingCellConfig_pathlossReferenceLinking                      `optional`
	ServingCellMO                                 *MeasObjectId                                                    `optional`
	Lte_CRS_ToMatchAround                         *RateMatchPatternLTE_CRS                                         `optional,ext-1,setuprelease`
	RateMatchPatternToAddModList                  []RateMatchPattern                                               `lb:1,ub:maxNrofRateMatchPatterns,optional,ext-1`
	RateMatchPatternToReleaseList                 []RateMatchPatternId                                             `lb:1,ub:maxNrofRateMatchPatterns,optional,ext-1`
	DownlinkChannelBW_PerSCS_List                 []SCS_SpecificCarrier                                            `lb:1,ub:maxSCSs,optional,ext-1`
	SupplementaryUplinkRelease_r16                *ServingCellConfig_supplementaryUplinkRelease_r16                `optional,ext-2`
	Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16   *TDD_UL_DL_ConfigDedicated_IAB_MT_r16                            `optional,ext-2`
	DormantBWP_Config_r16                         *DormantBWP_Config_r16                                           `optional,ext-2,setuprelease`
	Ca_SlotOffset_r16                             *ServingCellConfig_ca_SlotOffset_r16                             `lb:-2,ub:2,optional,ext-2`
	Dummy2                                        *DummyJ                                                          `optional,ext-2,setuprelease`
	IntraCellGuardBandsDL_List_r16                []IntraCellGuardBandsPerSCS_r16                                  `lb:1,ub:maxSCSs,optional,ext-2`
	IntraCellGuardBandsUL_List_r16                []IntraCellGuardBandsPerSCS_r16                                  `lb:1,ub:maxSCSs,optional,ext-2`
	Csi_RS_ValidationWithDCI_r16                  *ServingCellConfig_csi_RS_ValidationWithDCI_r16                  `optional,ext-2`
	Lte_CRS_PatternList1_r16                      *LTE_CRS_PatternList_r16                                         `optional,ext-2,setuprelease`
	Lte_CRS_PatternList2_r16                      *LTE_CRS_PatternList_r16                                         `optional,ext-2,setuprelease`
	Crs_RateMatch_PerCORESETPoolIndex_r16         *ServingCellConfig_crs_RateMatch_PerCORESETPoolIndex_r16         `optional,ext-2`
	EnableTwoDefaultTCI_States_r16                *ServingCellConfig_enableTwoDefaultTCI_States_r16                `optional,ext-2`
	EnableDefaultTCI_StatePerCoresetPoolIndex_r16 *ServingCellConfig_enableDefaultTCI_StatePerCoresetPoolIndex_r16 `optional,ext-2`
	EnableBeamSwitchTiming_r16                    *ServingCellConfig_enableBeamSwitchTiming_r16                    `optional,ext-2`
	Cbg_TxDiffTBsProcessingType1_r16              *ServingCellConfig_cbg_TxDiffTBsProcessingType1_r16              `optional,ext-2`
	Cbg_TxDiffTBsProcessingType2_r16              *ServingCellConfig_cbg_TxDiffTBsProcessingType2_r16              `optional,ext-2`
	DirectionalCollisionHandling_r16              *ServingCellConfig_directionalCollisionHandling_r16              `optional,ext-3`
	ChannelAccessConfig_r16                       *ChannelAccessConfig_r16                                         `optional,ext-3,setuprelease`
	Nr_dl_PRS_PDC_Info_r17                        *NR_DL_PRS_PDC_Info_r17                                          `optional,ext-4,setuprelease`
	SemiStaticChannelAccessConfigUE_r17           *SemiStaticChannelAccessConfigUE_r17                             `optional,ext-4,setuprelease`
	MimoParam_r17                                 *MIMOParam_r17                                                   `optional,ext-4,setuprelease`
	ChannelAccessMode2_r17                        *ServingCellConfig_channelAccessMode2_r17                        `optional,ext-4`
	TimeDomainHARQ_BundlingType1_r17              *ServingCellConfig_timeDomainHARQ_BundlingType1_r17              `optional,ext-4`
	NrofHARQ_BundlingGroups_r17                   *ServingCellConfig_nrofHARQ_BundlingGroups_r17                   `optional,ext-4`
	Fdmed_ReceptionMulticast_r17                  *ServingCellConfig_fdmed_ReceptionMulticast_r17                  `optional,ext-4`
	MoreThanOneNackOnlyMode_r17                   *ServingCellConfig_moreThanOneNackOnlyMode_r17                   `optional,ext-4`
	Tci_ActivatedConfig_r17                       *TCI_ActivatedConfig_r17                                         `optional,ext-4`
	DirectionalCollisionHandling_DC_r17           *ServingCellConfig_directionalCollisionHandling_DC_r17           `optional,ext-4`
	Lte_NeighCellsCRS_AssistInfoList_r17          *LTE_NeighCellsCRS_AssistInfoList_r17                            `optional,ext-4,setuprelease`
	Lte_NeighCellsCRS_Assumptions_r17             *ServingCellConfig_lte_NeighCellsCRS_Assumptions_r17             `optional,ext-5`
}

func (ie *ServingCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Lte_CRS_ToMatchAround != nil || len(ie.RateMatchPatternToAddModList) > 0 || len(ie.RateMatchPatternToReleaseList) > 0 || len(ie.DownlinkChannelBW_PerSCS_List) > 0 || ie.SupplementaryUplinkRelease_r16 != nil || ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil || ie.DormantBWP_Config_r16 != nil || ie.Ca_SlotOffset_r16 != nil || ie.Dummy2 != nil || len(ie.IntraCellGuardBandsDL_List_r16) > 0 || len(ie.IntraCellGuardBandsUL_List_r16) > 0 || ie.Csi_RS_ValidationWithDCI_r16 != nil || ie.Lte_CRS_PatternList1_r16 != nil || ie.Lte_CRS_PatternList2_r16 != nil || ie.Crs_RateMatch_PerCORESETPoolIndex_r16 != nil || ie.EnableTwoDefaultTCI_States_r16 != nil || ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil || ie.EnableBeamSwitchTiming_r16 != nil || ie.Cbg_TxDiffTBsProcessingType1_r16 != nil || ie.Cbg_TxDiffTBsProcessingType2_r16 != nil || ie.DirectionalCollisionHandling_r16 != nil || ie.ChannelAccessConfig_r16 != nil || ie.Nr_dl_PRS_PDC_Info_r17 != nil || ie.SemiStaticChannelAccessConfigUE_r17 != nil || ie.MimoParam_r17 != nil || ie.ChannelAccessMode2_r17 != nil || ie.TimeDomainHARQ_BundlingType1_r17 != nil || ie.NrofHARQ_BundlingGroups_r17 != nil || ie.Fdmed_ReceptionMulticast_r17 != nil || ie.MoreThanOneNackOnlyMode_r17 != nil || ie.Tci_ActivatedConfig_r17 != nil || ie.DirectionalCollisionHandling_DC_r17 != nil || ie.Lte_NeighCellsCRS_AssistInfoList_r17 != nil || ie.Lte_NeighCellsCRS_Assumptions_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Tdd_UL_DL_ConfigurationDedicated != nil, ie.InitialDownlinkBWP != nil, len(ie.DownlinkBWP_ToReleaseList) > 0, len(ie.DownlinkBWP_ToAddModList) > 0, ie.FirstActiveDownlinkBWP_Id != nil, ie.Bwp_InactivityTimer != nil, ie.DefaultDownlinkBWP_Id != nil, ie.UplinkConfig != nil, ie.SupplementaryUplink != nil, ie.Pdcch_ServingCellConfig != nil, ie.Pdsch_ServingCellConfig != nil, ie.Csi_MeasConfig != nil, ie.SCellDeactivationTimer != nil, ie.CrossCarrierSchedulingConfig != nil, ie.Dummy1 != nil, ie.PathlossReferenceLinking != nil, ie.ServingCellMO != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Tdd_UL_DL_ConfigurationDedicated != nil {
		if err = ie.Tdd_UL_DL_ConfigurationDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_UL_DL_ConfigurationDedicated", err)
		}
	}
	if ie.InitialDownlinkBWP != nil {
		if err = ie.InitialDownlinkBWP.Encode(w); err != nil {
			return utils.WrapError("Encode InitialDownlinkBWP", err)
		}
	}
	if len(ie.DownlinkBWP_ToReleaseList) > 0 {
		tmp_DownlinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.DownlinkBWP_ToReleaseList {
			tmp_DownlinkBWP_ToReleaseList.Value = append(tmp_DownlinkBWP_ToReleaseList.Value, &i)
		}
		if err = tmp_DownlinkBWP_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode DownlinkBWP_ToReleaseList", err)
		}
	}
	if len(ie.DownlinkBWP_ToAddModList) > 0 {
		tmp_DownlinkBWP_ToAddModList := utils.NewSequence[*BWP_Downlink]([]*BWP_Downlink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.DownlinkBWP_ToAddModList {
			tmp_DownlinkBWP_ToAddModList.Value = append(tmp_DownlinkBWP_ToAddModList.Value, &i)
		}
		if err = tmp_DownlinkBWP_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode DownlinkBWP_ToAddModList", err)
		}
	}
	if ie.FirstActiveDownlinkBWP_Id != nil {
		if err = ie.FirstActiveDownlinkBWP_Id.Encode(w); err != nil {
			return utils.WrapError("Encode FirstActiveDownlinkBWP_Id", err)
		}
	}
	if ie.Bwp_InactivityTimer != nil {
		if err = ie.Bwp_InactivityTimer.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_InactivityTimer", err)
		}
	}
	if ie.DefaultDownlinkBWP_Id != nil {
		if err = ie.DefaultDownlinkBWP_Id.Encode(w); err != nil {
			return utils.WrapError("Encode DefaultDownlinkBWP_Id", err)
		}
	}
	if ie.UplinkConfig != nil {
		if err = ie.UplinkConfig.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkConfig", err)
		}
	}
	if ie.SupplementaryUplink != nil {
		if err = ie.SupplementaryUplink.Encode(w); err != nil {
			return utils.WrapError("Encode SupplementaryUplink", err)
		}
	}
	if ie.Pdcch_ServingCellConfig != nil {
		tmp_Pdcch_ServingCellConfig := utils.SetupRelease[*PDCCH_ServingCellConfig]{
			Setup: ie.Pdcch_ServingCellConfig,
		}
		if err = tmp_Pdcch_ServingCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_ServingCellConfig", err)
		}
	}
	if ie.Pdsch_ServingCellConfig != nil {
		tmp_Pdsch_ServingCellConfig := utils.SetupRelease[*PDSCH_ServingCellConfig]{
			Setup: ie.Pdsch_ServingCellConfig,
		}
		if err = tmp_Pdsch_ServingCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ServingCellConfig", err)
		}
	}
	if ie.Csi_MeasConfig != nil {
		tmp_Csi_MeasConfig := utils.SetupRelease[*CSI_MeasConfig]{
			Setup: ie.Csi_MeasConfig,
		}
		if err = tmp_Csi_MeasConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_MeasConfig", err)
		}
	}
	if ie.SCellDeactivationTimer != nil {
		if err = ie.SCellDeactivationTimer.Encode(w); err != nil {
			return utils.WrapError("Encode SCellDeactivationTimer", err)
		}
	}
	if ie.CrossCarrierSchedulingConfig != nil {
		if err = ie.CrossCarrierSchedulingConfig.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierSchedulingConfig", err)
		}
	}
	if err = ie.Tag_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Tag_Id", err)
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.PathlossReferenceLinking != nil {
		if err = ie.PathlossReferenceLinking.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceLinking", err)
		}
	}
	if ie.ServingCellMO != nil {
		if err = ie.ServingCellMO.Encode(w); err != nil {
			return utils.WrapError("Encode ServingCellMO", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.Lte_CRS_ToMatchAround != nil || len(ie.RateMatchPatternToAddModList) > 0 || len(ie.RateMatchPatternToReleaseList) > 0 || len(ie.DownlinkChannelBW_PerSCS_List) > 0, ie.SupplementaryUplinkRelease_r16 != nil || ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil || ie.DormantBWP_Config_r16 != nil || ie.Ca_SlotOffset_r16 != nil || ie.Dummy2 != nil || len(ie.IntraCellGuardBandsDL_List_r16) > 0 || len(ie.IntraCellGuardBandsUL_List_r16) > 0 || ie.Csi_RS_ValidationWithDCI_r16 != nil || ie.Lte_CRS_PatternList1_r16 != nil || ie.Lte_CRS_PatternList2_r16 != nil || ie.Crs_RateMatch_PerCORESETPoolIndex_r16 != nil || ie.EnableTwoDefaultTCI_States_r16 != nil || ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil || ie.EnableBeamSwitchTiming_r16 != nil || ie.Cbg_TxDiffTBsProcessingType1_r16 != nil || ie.Cbg_TxDiffTBsProcessingType2_r16 != nil, ie.DirectionalCollisionHandling_r16 != nil || ie.ChannelAccessConfig_r16 != nil, ie.Nr_dl_PRS_PDC_Info_r17 != nil || ie.SemiStaticChannelAccessConfigUE_r17 != nil || ie.MimoParam_r17 != nil || ie.ChannelAccessMode2_r17 != nil || ie.TimeDomainHARQ_BundlingType1_r17 != nil || ie.NrofHARQ_BundlingGroups_r17 != nil || ie.Fdmed_ReceptionMulticast_r17 != nil || ie.MoreThanOneNackOnlyMode_r17 != nil || ie.Tci_ActivatedConfig_r17 != nil || ie.DirectionalCollisionHandling_DC_r17 != nil || ie.Lte_NeighCellsCRS_AssistInfoList_r17 != nil, ie.Lte_NeighCellsCRS_Assumptions_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ServingCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Lte_CRS_ToMatchAround != nil, len(ie.RateMatchPatternToAddModList) > 0, len(ie.RateMatchPatternToReleaseList) > 0, len(ie.DownlinkChannelBW_PerSCS_List) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Lte_CRS_ToMatchAround optional
			if ie.Lte_CRS_ToMatchAround != nil {
				tmp_Lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{
					Setup: ie.Lte_CRS_ToMatchAround,
				}
				if err = tmp_Lte_CRS_ToMatchAround.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lte_CRS_ToMatchAround", err)
				}
			}
			// encode RateMatchPatternToAddModList optional
			if len(ie.RateMatchPatternToAddModList) > 0 {
				tmp_RateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				for _, i := range ie.RateMatchPatternToAddModList {
					tmp_RateMatchPatternToAddModList.Value = append(tmp_RateMatchPatternToAddModList.Value, &i)
				}
				if err = tmp_RateMatchPatternToAddModList.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RateMatchPatternToAddModList", err)
				}
			}
			// encode RateMatchPatternToReleaseList optional
			if len(ie.RateMatchPatternToReleaseList) > 0 {
				tmp_RateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				for _, i := range ie.RateMatchPatternToReleaseList {
					tmp_RateMatchPatternToReleaseList.Value = append(tmp_RateMatchPatternToReleaseList.Value, &i)
				}
				if err = tmp_RateMatchPatternToReleaseList.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RateMatchPatternToReleaseList", err)
				}
			}
			// encode DownlinkChannelBW_PerSCS_List optional
			if len(ie.DownlinkChannelBW_PerSCS_List) > 0 {
				tmp_DownlinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.DownlinkChannelBW_PerSCS_List {
					tmp_DownlinkChannelBW_PerSCS_List.Value = append(tmp_DownlinkChannelBW_PerSCS_List.Value, &i)
				}
				if err = tmp_DownlinkChannelBW_PerSCS_List.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DownlinkChannelBW_PerSCS_List", err)
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
			optionals_ext_2 := []bool{ie.SupplementaryUplinkRelease_r16 != nil, ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil, ie.DormantBWP_Config_r16 != nil, ie.Ca_SlotOffset_r16 != nil, ie.Dummy2 != nil, len(ie.IntraCellGuardBandsDL_List_r16) > 0, len(ie.IntraCellGuardBandsUL_List_r16) > 0, ie.Csi_RS_ValidationWithDCI_r16 != nil, ie.Lte_CRS_PatternList1_r16 != nil, ie.Lte_CRS_PatternList2_r16 != nil, ie.Crs_RateMatch_PerCORESETPoolIndex_r16 != nil, ie.EnableTwoDefaultTCI_States_r16 != nil, ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil, ie.EnableBeamSwitchTiming_r16 != nil, ie.Cbg_TxDiffTBsProcessingType1_r16 != nil, ie.Cbg_TxDiffTBsProcessingType2_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupplementaryUplinkRelease_r16 optional
			if ie.SupplementaryUplinkRelease_r16 != nil {
				if err = ie.SupplementaryUplinkRelease_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupplementaryUplinkRelease_r16", err)
				}
			}
			// encode Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 optional
			if ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil {
				if err = ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16", err)
				}
			}
			// encode DormantBWP_Config_r16 optional
			if ie.DormantBWP_Config_r16 != nil {
				tmp_DormantBWP_Config_r16 := utils.SetupRelease[*DormantBWP_Config_r16]{
					Setup: ie.DormantBWP_Config_r16,
				}
				if err = tmp_DormantBWP_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DormantBWP_Config_r16", err)
				}
			}
			// encode Ca_SlotOffset_r16 optional
			if ie.Ca_SlotOffset_r16 != nil {
				if err = ie.Ca_SlotOffset_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ca_SlotOffset_r16", err)
				}
			}
			// encode Dummy2 optional
			if ie.Dummy2 != nil {
				tmp_Dummy2 := utils.SetupRelease[*DummyJ]{
					Setup: ie.Dummy2,
				}
				if err = tmp_Dummy2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy2", err)
				}
			}
			// encode IntraCellGuardBandsDL_List_r16 optional
			if len(ie.IntraCellGuardBandsDL_List_r16) > 0 {
				tmp_IntraCellGuardBandsDL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.IntraCellGuardBandsDL_List_r16 {
					tmp_IntraCellGuardBandsDL_List_r16.Value = append(tmp_IntraCellGuardBandsDL_List_r16.Value, &i)
				}
				if err = tmp_IntraCellGuardBandsDL_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraCellGuardBandsDL_List_r16", err)
				}
			}
			// encode IntraCellGuardBandsUL_List_r16 optional
			if len(ie.IntraCellGuardBandsUL_List_r16) > 0 {
				tmp_IntraCellGuardBandsUL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.IntraCellGuardBandsUL_List_r16 {
					tmp_IntraCellGuardBandsUL_List_r16.Value = append(tmp_IntraCellGuardBandsUL_List_r16.Value, &i)
				}
				if err = tmp_IntraCellGuardBandsUL_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraCellGuardBandsUL_List_r16", err)
				}
			}
			// encode Csi_RS_ValidationWithDCI_r16 optional
			if ie.Csi_RS_ValidationWithDCI_r16 != nil {
				if err = ie.Csi_RS_ValidationWithDCI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_RS_ValidationWithDCI_r16", err)
				}
			}
			// encode Lte_CRS_PatternList1_r16 optional
			if ie.Lte_CRS_PatternList1_r16 != nil {
				tmp_Lte_CRS_PatternList1_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{
					Setup: ie.Lte_CRS_PatternList1_r16,
				}
				if err = tmp_Lte_CRS_PatternList1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lte_CRS_PatternList1_r16", err)
				}
			}
			// encode Lte_CRS_PatternList2_r16 optional
			if ie.Lte_CRS_PatternList2_r16 != nil {
				tmp_Lte_CRS_PatternList2_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{
					Setup: ie.Lte_CRS_PatternList2_r16,
				}
				if err = tmp_Lte_CRS_PatternList2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lte_CRS_PatternList2_r16", err)
				}
			}
			// encode Crs_RateMatch_PerCORESETPoolIndex_r16 optional
			if ie.Crs_RateMatch_PerCORESETPoolIndex_r16 != nil {
				if err = ie.Crs_RateMatch_PerCORESETPoolIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Crs_RateMatch_PerCORESETPoolIndex_r16", err)
				}
			}
			// encode EnableTwoDefaultTCI_States_r16 optional
			if ie.EnableTwoDefaultTCI_States_r16 != nil {
				if err = ie.EnableTwoDefaultTCI_States_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnableTwoDefaultTCI_States_r16", err)
				}
			}
			// encode EnableDefaultTCI_StatePerCoresetPoolIndex_r16 optional
			if ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil {
				if err = ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnableDefaultTCI_StatePerCoresetPoolIndex_r16", err)
				}
			}
			// encode EnableBeamSwitchTiming_r16 optional
			if ie.EnableBeamSwitchTiming_r16 != nil {
				if err = ie.EnableBeamSwitchTiming_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnableBeamSwitchTiming_r16", err)
				}
			}
			// encode Cbg_TxDiffTBsProcessingType1_r16 optional
			if ie.Cbg_TxDiffTBsProcessingType1_r16 != nil {
				if err = ie.Cbg_TxDiffTBsProcessingType1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cbg_TxDiffTBsProcessingType1_r16", err)
				}
			}
			// encode Cbg_TxDiffTBsProcessingType2_r16 optional
			if ie.Cbg_TxDiffTBsProcessingType2_r16 != nil {
				if err = ie.Cbg_TxDiffTBsProcessingType2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cbg_TxDiffTBsProcessingType2_r16", err)
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
			optionals_ext_3 := []bool{ie.DirectionalCollisionHandling_r16 != nil, ie.ChannelAccessConfig_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode DirectionalCollisionHandling_r16 optional
			if ie.DirectionalCollisionHandling_r16 != nil {
				if err = ie.DirectionalCollisionHandling_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DirectionalCollisionHandling_r16", err)
				}
			}
			// encode ChannelAccessConfig_r16 optional
			if ie.ChannelAccessConfig_r16 != nil {
				tmp_ChannelAccessConfig_r16 := utils.SetupRelease[*ChannelAccessConfig_r16]{
					Setup: ie.ChannelAccessConfig_r16,
				}
				if err = tmp_ChannelAccessConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelAccessConfig_r16", err)
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
			optionals_ext_4 := []bool{ie.Nr_dl_PRS_PDC_Info_r17 != nil, ie.SemiStaticChannelAccessConfigUE_r17 != nil, ie.MimoParam_r17 != nil, ie.ChannelAccessMode2_r17 != nil, ie.TimeDomainHARQ_BundlingType1_r17 != nil, ie.NrofHARQ_BundlingGroups_r17 != nil, ie.Fdmed_ReceptionMulticast_r17 != nil, ie.MoreThanOneNackOnlyMode_r17 != nil, ie.Tci_ActivatedConfig_r17 != nil, ie.DirectionalCollisionHandling_DC_r17 != nil, ie.Lte_NeighCellsCRS_AssistInfoList_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Nr_dl_PRS_PDC_Info_r17 optional
			if ie.Nr_dl_PRS_PDC_Info_r17 != nil {
				tmp_Nr_dl_PRS_PDC_Info_r17 := utils.SetupRelease[*NR_DL_PRS_PDC_Info_r17]{
					Setup: ie.Nr_dl_PRS_PDC_Info_r17,
				}
				if err = tmp_Nr_dl_PRS_PDC_Info_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_dl_PRS_PDC_Info_r17", err)
				}
			}
			// encode SemiStaticChannelAccessConfigUE_r17 optional
			if ie.SemiStaticChannelAccessConfigUE_r17 != nil {
				tmp_SemiStaticChannelAccessConfigUE_r17 := utils.SetupRelease[*SemiStaticChannelAccessConfigUE_r17]{
					Setup: ie.SemiStaticChannelAccessConfigUE_r17,
				}
				if err = tmp_SemiStaticChannelAccessConfigUE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SemiStaticChannelAccessConfigUE_r17", err)
				}
			}
			// encode MimoParam_r17 optional
			if ie.MimoParam_r17 != nil {
				tmp_MimoParam_r17 := utils.SetupRelease[*MIMOParam_r17]{
					Setup: ie.MimoParam_r17,
				}
				if err = tmp_MimoParam_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MimoParam_r17", err)
				}
			}
			// encode ChannelAccessMode2_r17 optional
			if ie.ChannelAccessMode2_r17 != nil {
				if err = ie.ChannelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelAccessMode2_r17", err)
				}
			}
			// encode TimeDomainHARQ_BundlingType1_r17 optional
			if ie.TimeDomainHARQ_BundlingType1_r17 != nil {
				if err = ie.TimeDomainHARQ_BundlingType1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TimeDomainHARQ_BundlingType1_r17", err)
				}
			}
			// encode NrofHARQ_BundlingGroups_r17 optional
			if ie.NrofHARQ_BundlingGroups_r17 != nil {
				if err = ie.NrofHARQ_BundlingGroups_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NrofHARQ_BundlingGroups_r17", err)
				}
			}
			// encode Fdmed_ReceptionMulticast_r17 optional
			if ie.Fdmed_ReceptionMulticast_r17 != nil {
				if err = ie.Fdmed_ReceptionMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Fdmed_ReceptionMulticast_r17", err)
				}
			}
			// encode MoreThanOneNackOnlyMode_r17 optional
			if ie.MoreThanOneNackOnlyMode_r17 != nil {
				if err = ie.MoreThanOneNackOnlyMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MoreThanOneNackOnlyMode_r17", err)
				}
			}
			// encode Tci_ActivatedConfig_r17 optional
			if ie.Tci_ActivatedConfig_r17 != nil {
				if err = ie.Tci_ActivatedConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tci_ActivatedConfig_r17", err)
				}
			}
			// encode DirectionalCollisionHandling_DC_r17 optional
			if ie.DirectionalCollisionHandling_DC_r17 != nil {
				if err = ie.DirectionalCollisionHandling_DC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DirectionalCollisionHandling_DC_r17", err)
				}
			}
			// encode Lte_NeighCellsCRS_AssistInfoList_r17 optional
			if ie.Lte_NeighCellsCRS_AssistInfoList_r17 != nil {
				tmp_Lte_NeighCellsCRS_AssistInfoList_r17 := utils.SetupRelease[*LTE_NeighCellsCRS_AssistInfoList_r17]{
					Setup: ie.Lte_NeighCellsCRS_AssistInfoList_r17,
				}
				if err = tmp_Lte_NeighCellsCRS_AssistInfoList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lte_NeighCellsCRS_AssistInfoList_r17", err)
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
			optionals_ext_5 := []bool{ie.Lte_NeighCellsCRS_Assumptions_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Lte_NeighCellsCRS_Assumptions_r17 optional
			if ie.Lte_NeighCellsCRS_Assumptions_r17 != nil {
				if err = ie.Lte_NeighCellsCRS_Assumptions_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lte_NeighCellsCRS_Assumptions_r17", err)
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

func (ie *ServingCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_UL_DL_ConfigurationDedicatedPresent bool
	if Tdd_UL_DL_ConfigurationDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InitialDownlinkBWPPresent bool
	if InitialDownlinkBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DownlinkBWP_ToReleaseListPresent bool
	if DownlinkBWP_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DownlinkBWP_ToAddModListPresent bool
	if DownlinkBWP_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FirstActiveDownlinkBWP_IdPresent bool
	if FirstActiveDownlinkBWP_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_InactivityTimerPresent bool
	if Bwp_InactivityTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DefaultDownlinkBWP_IdPresent bool
	if DefaultDownlinkBWP_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkConfigPresent bool
	if UplinkConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupplementaryUplinkPresent bool
	if SupplementaryUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_ServingCellConfigPresent bool
	if Pdcch_ServingCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ServingCellConfigPresent bool
	if Pdsch_ServingCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_MeasConfigPresent bool
	if Csi_MeasConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SCellDeactivationTimerPresent bool
	if SCellDeactivationTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierSchedulingConfigPresent bool
	if CrossCarrierSchedulingConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceLinkingPresent bool
	if PathlossReferenceLinkingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ServingCellMOPresent bool
	if ServingCellMOPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Tdd_UL_DL_ConfigurationDedicatedPresent {
		ie.Tdd_UL_DL_ConfigurationDedicated = new(TDD_UL_DL_ConfigDedicated)
		if err = ie.Tdd_UL_DL_ConfigurationDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_UL_DL_ConfigurationDedicated", err)
		}
	}
	if InitialDownlinkBWPPresent {
		ie.InitialDownlinkBWP = new(BWP_DownlinkDedicated)
		if err = ie.InitialDownlinkBWP.Decode(r); err != nil {
			return utils.WrapError("Decode InitialDownlinkBWP", err)
		}
	}
	if DownlinkBWP_ToReleaseListPresent {
		tmp_DownlinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_DownlinkBWP_ToReleaseList := func() *BWP_Id {
			return new(BWP_Id)
		}
		if err = tmp_DownlinkBWP_ToReleaseList.Decode(r, fn_DownlinkBWP_ToReleaseList); err != nil {
			return utils.WrapError("Decode DownlinkBWP_ToReleaseList", err)
		}
		ie.DownlinkBWP_ToReleaseList = []BWP_Id{}
		for _, i := range tmp_DownlinkBWP_ToReleaseList.Value {
			ie.DownlinkBWP_ToReleaseList = append(ie.DownlinkBWP_ToReleaseList, *i)
		}
	}
	if DownlinkBWP_ToAddModListPresent {
		tmp_DownlinkBWP_ToAddModList := utils.NewSequence[*BWP_Downlink]([]*BWP_Downlink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_DownlinkBWP_ToAddModList := func() *BWP_Downlink {
			return new(BWP_Downlink)
		}
		if err = tmp_DownlinkBWP_ToAddModList.Decode(r, fn_DownlinkBWP_ToAddModList); err != nil {
			return utils.WrapError("Decode DownlinkBWP_ToAddModList", err)
		}
		ie.DownlinkBWP_ToAddModList = []BWP_Downlink{}
		for _, i := range tmp_DownlinkBWP_ToAddModList.Value {
			ie.DownlinkBWP_ToAddModList = append(ie.DownlinkBWP_ToAddModList, *i)
		}
	}
	if FirstActiveDownlinkBWP_IdPresent {
		ie.FirstActiveDownlinkBWP_Id = new(BWP_Id)
		if err = ie.FirstActiveDownlinkBWP_Id.Decode(r); err != nil {
			return utils.WrapError("Decode FirstActiveDownlinkBWP_Id", err)
		}
	}
	if Bwp_InactivityTimerPresent {
		ie.Bwp_InactivityTimer = new(ServingCellConfig_bwp_InactivityTimer)
		if err = ie.Bwp_InactivityTimer.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_InactivityTimer", err)
		}
	}
	if DefaultDownlinkBWP_IdPresent {
		ie.DefaultDownlinkBWP_Id = new(BWP_Id)
		if err = ie.DefaultDownlinkBWP_Id.Decode(r); err != nil {
			return utils.WrapError("Decode DefaultDownlinkBWP_Id", err)
		}
	}
	if UplinkConfigPresent {
		ie.UplinkConfig = new(UplinkConfig)
		if err = ie.UplinkConfig.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkConfig", err)
		}
	}
	if SupplementaryUplinkPresent {
		ie.SupplementaryUplink = new(UplinkConfig)
		if err = ie.SupplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode SupplementaryUplink", err)
		}
	}
	if Pdcch_ServingCellConfigPresent {
		tmp_Pdcch_ServingCellConfig := utils.SetupRelease[*PDCCH_ServingCellConfig]{}
		if err = tmp_Pdcch_ServingCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_ServingCellConfig", err)
		}
		ie.Pdcch_ServingCellConfig = tmp_Pdcch_ServingCellConfig.Setup
	}
	if Pdsch_ServingCellConfigPresent {
		tmp_Pdsch_ServingCellConfig := utils.SetupRelease[*PDSCH_ServingCellConfig]{}
		if err = tmp_Pdsch_ServingCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ServingCellConfig", err)
		}
		ie.Pdsch_ServingCellConfig = tmp_Pdsch_ServingCellConfig.Setup
	}
	if Csi_MeasConfigPresent {
		tmp_Csi_MeasConfig := utils.SetupRelease[*CSI_MeasConfig]{}
		if err = tmp_Csi_MeasConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_MeasConfig", err)
		}
		ie.Csi_MeasConfig = tmp_Csi_MeasConfig.Setup
	}
	if SCellDeactivationTimerPresent {
		ie.SCellDeactivationTimer = new(ServingCellConfig_sCellDeactivationTimer)
		if err = ie.SCellDeactivationTimer.Decode(r); err != nil {
			return utils.WrapError("Decode SCellDeactivationTimer", err)
		}
	}
	if CrossCarrierSchedulingConfigPresent {
		ie.CrossCarrierSchedulingConfig = new(CrossCarrierSchedulingConfig)
		if err = ie.CrossCarrierSchedulingConfig.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierSchedulingConfig", err)
		}
	}
	if err = ie.Tag_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Tag_Id", err)
	}
	if Dummy1Present {
		ie.Dummy1 = new(ServingCellConfig_dummy1)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if PathlossReferenceLinkingPresent {
		ie.PathlossReferenceLinking = new(ServingCellConfig_pathlossReferenceLinking)
		if err = ie.PathlossReferenceLinking.Decode(r); err != nil {
			return utils.WrapError("Decode PathlossReferenceLinking", err)
		}
	}
	if ServingCellMOPresent {
		ie.ServingCellMO = new(MeasObjectId)
		if err = ie.ServingCellMO.Decode(r); err != nil {
			return utils.WrapError("Decode ServingCellMO", err)
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

			Lte_CRS_ToMatchAroundPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RateMatchPatternToAddModListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RateMatchPatternToReleaseListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DownlinkChannelBW_PerSCS_ListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Lte_CRS_ToMatchAround optional
			if Lte_CRS_ToMatchAroundPresent {
				tmp_Lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{}
				if err = tmp_Lte_CRS_ToMatchAround.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lte_CRS_ToMatchAround", err)
				}
				ie.Lte_CRS_ToMatchAround = tmp_Lte_CRS_ToMatchAround.Setup
			}
			// decode RateMatchPatternToAddModList optional
			if RateMatchPatternToAddModListPresent {
				tmp_RateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				fn_RateMatchPatternToAddModList := func() *RateMatchPattern {
					return new(RateMatchPattern)
				}
				if err = tmp_RateMatchPatternToAddModList.Decode(extReader, fn_RateMatchPatternToAddModList); err != nil {
					return utils.WrapError("Decode RateMatchPatternToAddModList", err)
				}
				ie.RateMatchPatternToAddModList = []RateMatchPattern{}
				for _, i := range tmp_RateMatchPatternToAddModList.Value {
					ie.RateMatchPatternToAddModList = append(ie.RateMatchPatternToAddModList, *i)
				}
			}
			// decode RateMatchPatternToReleaseList optional
			if RateMatchPatternToReleaseListPresent {
				tmp_RateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				fn_RateMatchPatternToReleaseList := func() *RateMatchPatternId {
					return new(RateMatchPatternId)
				}
				if err = tmp_RateMatchPatternToReleaseList.Decode(extReader, fn_RateMatchPatternToReleaseList); err != nil {
					return utils.WrapError("Decode RateMatchPatternToReleaseList", err)
				}
				ie.RateMatchPatternToReleaseList = []RateMatchPatternId{}
				for _, i := range tmp_RateMatchPatternToReleaseList.Value {
					ie.RateMatchPatternToReleaseList = append(ie.RateMatchPatternToReleaseList, *i)
				}
			}
			// decode DownlinkChannelBW_PerSCS_List optional
			if DownlinkChannelBW_PerSCS_ListPresent {
				tmp_DownlinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_DownlinkChannelBW_PerSCS_List := func() *SCS_SpecificCarrier {
					return new(SCS_SpecificCarrier)
				}
				if err = tmp_DownlinkChannelBW_PerSCS_List.Decode(extReader, fn_DownlinkChannelBW_PerSCS_List); err != nil {
					return utils.WrapError("Decode DownlinkChannelBW_PerSCS_List", err)
				}
				ie.DownlinkChannelBW_PerSCS_List = []SCS_SpecificCarrier{}
				for _, i := range tmp_DownlinkChannelBW_PerSCS_List.Value {
					ie.DownlinkChannelBW_PerSCS_List = append(ie.DownlinkChannelBW_PerSCS_List, *i)
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

			SupplementaryUplinkRelease_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DormantBWP_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ca_SlotOffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dummy2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntraCellGuardBandsDL_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntraCellGuardBandsUL_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_RS_ValidationWithDCI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lte_CRS_PatternList1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lte_CRS_PatternList2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Crs_RateMatch_PerCORESETPoolIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnableTwoDefaultTCI_States_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnableDefaultTCI_StatePerCoresetPoolIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnableBeamSwitchTiming_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cbg_TxDiffTBsProcessingType1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cbg_TxDiffTBsProcessingType2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupplementaryUplinkRelease_r16 optional
			if SupplementaryUplinkRelease_r16Present {
				ie.SupplementaryUplinkRelease_r16 = new(ServingCellConfig_supplementaryUplinkRelease_r16)
				if err = ie.SupplementaryUplinkRelease_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupplementaryUplinkRelease_r16", err)
				}
			}
			// decode Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 optional
			if Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16Present {
				ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 = new(TDD_UL_DL_ConfigDedicated_IAB_MT_r16)
				if err = ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16", err)
				}
			}
			// decode DormantBWP_Config_r16 optional
			if DormantBWP_Config_r16Present {
				tmp_DormantBWP_Config_r16 := utils.SetupRelease[*DormantBWP_Config_r16]{}
				if err = tmp_DormantBWP_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DormantBWP_Config_r16", err)
				}
				ie.DormantBWP_Config_r16 = tmp_DormantBWP_Config_r16.Setup
			}
			// decode Ca_SlotOffset_r16 optional
			if Ca_SlotOffset_r16Present {
				ie.Ca_SlotOffset_r16 = new(ServingCellConfig_ca_SlotOffset_r16)
				if err = ie.Ca_SlotOffset_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ca_SlotOffset_r16", err)
				}
			}
			// decode Dummy2 optional
			if Dummy2Present {
				tmp_Dummy2 := utils.SetupRelease[*DummyJ]{}
				if err = tmp_Dummy2.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dummy2", err)
				}
				ie.Dummy2 = tmp_Dummy2.Setup
			}
			// decode IntraCellGuardBandsDL_List_r16 optional
			if IntraCellGuardBandsDL_List_r16Present {
				tmp_IntraCellGuardBandsDL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_IntraCellGuardBandsDL_List_r16 := func() *IntraCellGuardBandsPerSCS_r16 {
					return new(IntraCellGuardBandsPerSCS_r16)
				}
				if err = tmp_IntraCellGuardBandsDL_List_r16.Decode(extReader, fn_IntraCellGuardBandsDL_List_r16); err != nil {
					return utils.WrapError("Decode IntraCellGuardBandsDL_List_r16", err)
				}
				ie.IntraCellGuardBandsDL_List_r16 = []IntraCellGuardBandsPerSCS_r16{}
				for _, i := range tmp_IntraCellGuardBandsDL_List_r16.Value {
					ie.IntraCellGuardBandsDL_List_r16 = append(ie.IntraCellGuardBandsDL_List_r16, *i)
				}
			}
			// decode IntraCellGuardBandsUL_List_r16 optional
			if IntraCellGuardBandsUL_List_r16Present {
				tmp_IntraCellGuardBandsUL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_IntraCellGuardBandsUL_List_r16 := func() *IntraCellGuardBandsPerSCS_r16 {
					return new(IntraCellGuardBandsPerSCS_r16)
				}
				if err = tmp_IntraCellGuardBandsUL_List_r16.Decode(extReader, fn_IntraCellGuardBandsUL_List_r16); err != nil {
					return utils.WrapError("Decode IntraCellGuardBandsUL_List_r16", err)
				}
				ie.IntraCellGuardBandsUL_List_r16 = []IntraCellGuardBandsPerSCS_r16{}
				for _, i := range tmp_IntraCellGuardBandsUL_List_r16.Value {
					ie.IntraCellGuardBandsUL_List_r16 = append(ie.IntraCellGuardBandsUL_List_r16, *i)
				}
			}
			// decode Csi_RS_ValidationWithDCI_r16 optional
			if Csi_RS_ValidationWithDCI_r16Present {
				ie.Csi_RS_ValidationWithDCI_r16 = new(ServingCellConfig_csi_RS_ValidationWithDCI_r16)
				if err = ie.Csi_RS_ValidationWithDCI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_RS_ValidationWithDCI_r16", err)
				}
			}
			// decode Lte_CRS_PatternList1_r16 optional
			if Lte_CRS_PatternList1_r16Present {
				tmp_Lte_CRS_PatternList1_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{}
				if err = tmp_Lte_CRS_PatternList1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lte_CRS_PatternList1_r16", err)
				}
				ie.Lte_CRS_PatternList1_r16 = tmp_Lte_CRS_PatternList1_r16.Setup
			}
			// decode Lte_CRS_PatternList2_r16 optional
			if Lte_CRS_PatternList2_r16Present {
				tmp_Lte_CRS_PatternList2_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{}
				if err = tmp_Lte_CRS_PatternList2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lte_CRS_PatternList2_r16", err)
				}
				ie.Lte_CRS_PatternList2_r16 = tmp_Lte_CRS_PatternList2_r16.Setup
			}
			// decode Crs_RateMatch_PerCORESETPoolIndex_r16 optional
			if Crs_RateMatch_PerCORESETPoolIndex_r16Present {
				ie.Crs_RateMatch_PerCORESETPoolIndex_r16 = new(ServingCellConfig_crs_RateMatch_PerCORESETPoolIndex_r16)
				if err = ie.Crs_RateMatch_PerCORESETPoolIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Crs_RateMatch_PerCORESETPoolIndex_r16", err)
				}
			}
			// decode EnableTwoDefaultTCI_States_r16 optional
			if EnableTwoDefaultTCI_States_r16Present {
				ie.EnableTwoDefaultTCI_States_r16 = new(ServingCellConfig_enableTwoDefaultTCI_States_r16)
				if err = ie.EnableTwoDefaultTCI_States_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnableTwoDefaultTCI_States_r16", err)
				}
			}
			// decode EnableDefaultTCI_StatePerCoresetPoolIndex_r16 optional
			if EnableDefaultTCI_StatePerCoresetPoolIndex_r16Present {
				ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 = new(ServingCellConfig_enableDefaultTCI_StatePerCoresetPoolIndex_r16)
				if err = ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnableDefaultTCI_StatePerCoresetPoolIndex_r16", err)
				}
			}
			// decode EnableBeamSwitchTiming_r16 optional
			if EnableBeamSwitchTiming_r16Present {
				ie.EnableBeamSwitchTiming_r16 = new(ServingCellConfig_enableBeamSwitchTiming_r16)
				if err = ie.EnableBeamSwitchTiming_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnableBeamSwitchTiming_r16", err)
				}
			}
			// decode Cbg_TxDiffTBsProcessingType1_r16 optional
			if Cbg_TxDiffTBsProcessingType1_r16Present {
				ie.Cbg_TxDiffTBsProcessingType1_r16 = new(ServingCellConfig_cbg_TxDiffTBsProcessingType1_r16)
				if err = ie.Cbg_TxDiffTBsProcessingType1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cbg_TxDiffTBsProcessingType1_r16", err)
				}
			}
			// decode Cbg_TxDiffTBsProcessingType2_r16 optional
			if Cbg_TxDiffTBsProcessingType2_r16Present {
				ie.Cbg_TxDiffTBsProcessingType2_r16 = new(ServingCellConfig_cbg_TxDiffTBsProcessingType2_r16)
				if err = ie.Cbg_TxDiffTBsProcessingType2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cbg_TxDiffTBsProcessingType2_r16", err)
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

			DirectionalCollisionHandling_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelAccessConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode DirectionalCollisionHandling_r16 optional
			if DirectionalCollisionHandling_r16Present {
				ie.DirectionalCollisionHandling_r16 = new(ServingCellConfig_directionalCollisionHandling_r16)
				if err = ie.DirectionalCollisionHandling_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DirectionalCollisionHandling_r16", err)
				}
			}
			// decode ChannelAccessConfig_r16 optional
			if ChannelAccessConfig_r16Present {
				tmp_ChannelAccessConfig_r16 := utils.SetupRelease[*ChannelAccessConfig_r16]{}
				if err = tmp_ChannelAccessConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelAccessConfig_r16", err)
				}
				ie.ChannelAccessConfig_r16 = tmp_ChannelAccessConfig_r16.Setup
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Nr_dl_PRS_PDC_Info_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SemiStaticChannelAccessConfigUE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MimoParam_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TimeDomainHARQ_BundlingType1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NrofHARQ_BundlingGroups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Fdmed_ReceptionMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MoreThanOneNackOnlyMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tci_ActivatedConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DirectionalCollisionHandling_DC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lte_NeighCellsCRS_AssistInfoList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Nr_dl_PRS_PDC_Info_r17 optional
			if Nr_dl_PRS_PDC_Info_r17Present {
				tmp_Nr_dl_PRS_PDC_Info_r17 := utils.SetupRelease[*NR_DL_PRS_PDC_Info_r17]{}
				if err = tmp_Nr_dl_PRS_PDC_Info_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_dl_PRS_PDC_Info_r17", err)
				}
				ie.Nr_dl_PRS_PDC_Info_r17 = tmp_Nr_dl_PRS_PDC_Info_r17.Setup
			}
			// decode SemiStaticChannelAccessConfigUE_r17 optional
			if SemiStaticChannelAccessConfigUE_r17Present {
				tmp_SemiStaticChannelAccessConfigUE_r17 := utils.SetupRelease[*SemiStaticChannelAccessConfigUE_r17]{}
				if err = tmp_SemiStaticChannelAccessConfigUE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SemiStaticChannelAccessConfigUE_r17", err)
				}
				ie.SemiStaticChannelAccessConfigUE_r17 = tmp_SemiStaticChannelAccessConfigUE_r17.Setup
			}
			// decode MimoParam_r17 optional
			if MimoParam_r17Present {
				tmp_MimoParam_r17 := utils.SetupRelease[*MIMOParam_r17]{}
				if err = tmp_MimoParam_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MimoParam_r17", err)
				}
				ie.MimoParam_r17 = tmp_MimoParam_r17.Setup
			}
			// decode ChannelAccessMode2_r17 optional
			if ChannelAccessMode2_r17Present {
				ie.ChannelAccessMode2_r17 = new(ServingCellConfig_channelAccessMode2_r17)
				if err = ie.ChannelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelAccessMode2_r17", err)
				}
			}
			// decode TimeDomainHARQ_BundlingType1_r17 optional
			if TimeDomainHARQ_BundlingType1_r17Present {
				ie.TimeDomainHARQ_BundlingType1_r17 = new(ServingCellConfig_timeDomainHARQ_BundlingType1_r17)
				if err = ie.TimeDomainHARQ_BundlingType1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TimeDomainHARQ_BundlingType1_r17", err)
				}
			}
			// decode NrofHARQ_BundlingGroups_r17 optional
			if NrofHARQ_BundlingGroups_r17Present {
				ie.NrofHARQ_BundlingGroups_r17 = new(ServingCellConfig_nrofHARQ_BundlingGroups_r17)
				if err = ie.NrofHARQ_BundlingGroups_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NrofHARQ_BundlingGroups_r17", err)
				}
			}
			// decode Fdmed_ReceptionMulticast_r17 optional
			if Fdmed_ReceptionMulticast_r17Present {
				ie.Fdmed_ReceptionMulticast_r17 = new(ServingCellConfig_fdmed_ReceptionMulticast_r17)
				if err = ie.Fdmed_ReceptionMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Fdmed_ReceptionMulticast_r17", err)
				}
			}
			// decode MoreThanOneNackOnlyMode_r17 optional
			if MoreThanOneNackOnlyMode_r17Present {
				ie.MoreThanOneNackOnlyMode_r17 = new(ServingCellConfig_moreThanOneNackOnlyMode_r17)
				if err = ie.MoreThanOneNackOnlyMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MoreThanOneNackOnlyMode_r17", err)
				}
			}
			// decode Tci_ActivatedConfig_r17 optional
			if Tci_ActivatedConfig_r17Present {
				ie.Tci_ActivatedConfig_r17 = new(TCI_ActivatedConfig_r17)
				if err = ie.Tci_ActivatedConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tci_ActivatedConfig_r17", err)
				}
			}
			// decode DirectionalCollisionHandling_DC_r17 optional
			if DirectionalCollisionHandling_DC_r17Present {
				ie.DirectionalCollisionHandling_DC_r17 = new(ServingCellConfig_directionalCollisionHandling_DC_r17)
				if err = ie.DirectionalCollisionHandling_DC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DirectionalCollisionHandling_DC_r17", err)
				}
			}
			// decode Lte_NeighCellsCRS_AssistInfoList_r17 optional
			if Lte_NeighCellsCRS_AssistInfoList_r17Present {
				tmp_Lte_NeighCellsCRS_AssistInfoList_r17 := utils.SetupRelease[*LTE_NeighCellsCRS_AssistInfoList_r17]{}
				if err = tmp_Lte_NeighCellsCRS_AssistInfoList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lte_NeighCellsCRS_AssistInfoList_r17", err)
				}
				ie.Lte_NeighCellsCRS_AssistInfoList_r17 = tmp_Lte_NeighCellsCRS_AssistInfoList_r17.Setup
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Lte_NeighCellsCRS_Assumptions_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Lte_NeighCellsCRS_Assumptions_r17 optional
			if Lte_NeighCellsCRS_Assumptions_r17Present {
				ie.Lte_NeighCellsCRS_Assumptions_r17 = new(ServingCellConfig_lte_NeighCellsCRS_Assumptions_r17)
				if err = ie.Lte_NeighCellsCRS_Assumptions_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lte_NeighCellsCRS_Assumptions_r17", err)
				}
			}
		}
	}
	return nil
}
