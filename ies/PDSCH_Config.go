package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config struct {
	DataScramblingIdentityPDSCH                             *int64                                                      `lb:0,ub:1023,optional`
	Dmrs_DownlinkForPDSCH_MappingTypeA                      *DMRS_DownlinkConfig                                        `optional,setuprelease`
	Dmrs_DownlinkForPDSCH_MappingTypeB                      *DMRS_DownlinkConfig                                        `optional,setuprelease`
	Tci_StatesToAddModList                                  []TCI_State                                                 `lb:1,ub:maxNrofTCI_States,optional`
	Tci_StatesToReleaseList                                 []TCI_StateId                                               `lb:1,ub:maxNrofTCI_States,optional`
	Vrb_ToPRB_Interleaver                                   *PDSCH_Config_vrb_ToPRB_Interleaver                         `optional`
	ResourceAllocation                                      PDSCH_Config_resourceAllocation                             `madatory`
	Pdsch_TimeDomainAllocationList                          *PDSCH_TimeDomainResourceAllocationList                     `optional,setuprelease`
	Pdsch_AggregationFactor                                 *PDSCH_Config_pdsch_AggregationFactor                       `optional`
	RateMatchPatternToAddModList                            []RateMatchPattern                                          `lb:1,ub:maxNrofRateMatchPatterns,optional`
	RateMatchPatternToReleaseList                           []RateMatchPatternId                                        `lb:1,ub:maxNrofRateMatchPatterns,optional`
	RateMatchPatternGroup1                                  *RateMatchPatternGroup                                      `optional`
	RateMatchPatternGroup2                                  *RateMatchPatternGroup                                      `optional`
	Rbg_Size                                                PDSCH_Config_rbg_Size                                       `madatory`
	Mcs_Table                                               *PDSCH_Config_mcs_Table                                     `optional`
	MaxNrofCodeWordsScheduledByDCI                          *PDSCH_Config_maxNrofCodeWordsScheduledByDCI                `optional`
	Prb_BundlingType                                        *PDSCH_Config_prb_BundlingType                              `optional`
	Zp_CSI_RS_ResourceToAddModList                          []ZP_CSI_RS_Resource                                        `lb:1,ub:maxNrofZP_CSI_RS_Resources,optional`
	Zp_CSI_RS_ResourceToReleaseList                         []ZP_CSI_RS_ResourceId                                      `lb:1,ub:maxNrofZP_CSI_RS_Resources,optional`
	Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList            []ZP_CSI_RS_ResourceSet                                     `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList           []ZP_CSI_RS_ResourceSetId                                   `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	Sp_ZP_CSI_RS_ResourceSetsToAddModList                   []ZP_CSI_RS_ResourceSet                                     `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	Sp_ZP_CSI_RS_ResourceSetsToReleaseList                  []ZP_CSI_RS_ResourceSetId                                   `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	P_ZP_CSI_RS_ResourceSet                                 *ZP_CSI_RS_ResourceSet                                      `optional,setuprelease`
	MaxMIMO_Layers_r16                                      *MaxMIMO_LayersDL_r16                                       `optional,ext-1,setuprelease`
	MinimumSchedulingOffsetK0_r16                           *MinSchedulingOffsetK0_Values_r16                           `optional,ext-1,setuprelease`
	AntennaPortsFieldPresenceDCI_1_2_r16                    *PDSCH_Config_antennaPortsFieldPresenceDCI_1_2_r16          `optional,ext-1`
	AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16  []ZP_CSI_RS_ResourceSet                                     `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional,ext-1`
	AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 []ZP_CSI_RS_ResourceSetId                                   `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional,ext-1`
	Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16          *DMRS_DownlinkConfig                                        `optional,ext-1,setuprelease`
	Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16          *DMRS_DownlinkConfig                                        `optional,ext-1,setuprelease`
	Dmrs_SequenceInitializationDCI_1_2_r16                  *PDSCH_Config_dmrs_SequenceInitializationDCI_1_2_r16        `optional,ext-1`
	Harq_ProcessNumberSizeDCI_1_2_r16                       *int64                                                      `lb:0,ub:4,optional,ext-1`
	Mcs_TableDCI_1_2_r16                                    *PDSCH_Config_mcs_TableDCI_1_2_r16                          `optional,ext-1`
	NumberOfBitsForRV_DCI_1_2_r16                           *int64                                                      `lb:0,ub:2,optional,ext-1`
	Pdsch_TimeDomainAllocationListDCI_1_2_r16               *PDSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	Prb_BundlingTypeDCI_1_2_r16                             *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16                   `optional,ext-1`
	PriorityIndicatorDCI_1_2_r16                            *PDSCH_Config_priorityIndicatorDCI_1_2_r16                  `optional,ext-1`
	RateMatchPatternGroup1DCI_1_2_r16                       *RateMatchPatternGroup                                      `optional,ext-1`
	RateMatchPatternGroup2DCI_1_2_r16                       *RateMatchPatternGroup                                      `optional,ext-1`
	ResourceAllocationType1GranularityDCI_1_2_r16           *PDSCH_Config_resourceAllocationType1GranularityDCI_1_2_r16 `optional,ext-1`
	Vrb_ToPRB_InterleaverDCI_1_2_r16                        *PDSCH_Config_vrb_ToPRB_InterleaverDCI_1_2_r16              `optional,ext-1`
	ReferenceOfSLIVDCI_1_2_r16                              *PDSCH_Config_referenceOfSLIVDCI_1_2_r16                    `optional,ext-1`
	ResourceAllocationDCI_1_2_r16                           *PDSCH_Config_resourceAllocationDCI_1_2_r16                 `optional,ext-1`
	PriorityIndicatorDCI_1_1_r16                            *PDSCH_Config_priorityIndicatorDCI_1_1_r16                  `optional,ext-1`
	DataScramblingIdentityPDSCH2_r16                        *int64                                                      `lb:0,ub:1023,optional,ext-1`
	Pdsch_TimeDomainAllocationList_r16                      *PDSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	RepetitionSchemeConfig_r16                              *RepetitionSchemeConfig_r16                                 `optional,ext-1,setuprelease`
	RepetitionSchemeConfig_v1630                            *RepetitionSchemeConfig_v1630                               `optional,ext-2,setuprelease`
	Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17               *PDSCH_Config_pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17     `optional,ext-3`
	Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17                      *PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_1_2_r17            `optional,ext-3`
	Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17                *PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17      `optional,ext-3`
	Pdsch_HARQ_ACK_RetxDCI_1_2_r17                          *PDSCH_Config_pdsch_HARQ_ACK_RetxDCI_1_2_r17                `optional,ext-3`
	Pucch_sSCellDynDCI_1_2_r17                              *PDSCH_Config_pucch_sSCellDynDCI_1_2_r17                    `optional,ext-3`
	Dl_OrJointTCI_StateList_r17                             *PDSCH_Config_dl_OrJointTCI_StateList_r17                   `lb:1,ub:maxNrofTCI_States,optional,ext-3`
	BeamAppTime_r17                                         *PDSCH_Config_beamAppTime_r17                               `optional,ext-3`
	Dummy                                                   *Dummy_TDRA_List                                            `optional,ext-3,setuprelease`
	Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17                  *PDSCH_Config_dmrs_FD_OCC_DisabledForRank1_PDSCH_r17        `optional,ext-3`
	MinimumSchedulingOffsetK0_r17                           *MinSchedulingOffsetK0_Values_r17                           `optional,ext-3,setuprelease`
	Harq_ProcessNumberSizeDCI_1_2_v1700                     *int64                                                      `lb:0,ub:5,optional,ext-3`
	Harq_ProcessNumberSizeDCI_1_1_r17                       *int64                                                      `lb:5,ub:5,optional,ext-3`
	Mcs_Table_r17                                           *PDSCH_Config_mcs_Table_r17                                 `optional,ext-3`
	Mcs_TableDCI_1_2_r17                                    *PDSCH_Config_mcs_TableDCI_1_2_r17                          `optional,ext-3`
	XOverheadMulticast_r17                                  *PDSCH_Config_xOverheadMulticast_r17                        `optional,ext-3`
	PriorityIndicatorDCI_4_2_r17                            *PDSCH_Config_priorityIndicatorDCI_4_2_r17                  `optional,ext-3`
	SizeDCI_4_2_r17                                         *int64                                                      `lb:20,ub:maxDCI_4_2_Size_r17,optional,ext-3`
	Pdsch_TimeDomainAllocationListForMultiPDSCH_r17         *MultiPDSCH_TDRA_List_r17                                   `optional,ext-4,setuprelease`
}

func (ie *PDSCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.MaxMIMO_Layers_r16 != nil || ie.MinimumSchedulingOffsetK0_r16 != nil || ie.AntennaPortsFieldPresenceDCI_1_2_r16 != nil || len(ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0 || len(ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0 || ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil || ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil || ie.Dmrs_SequenceInitializationDCI_1_2_r16 != nil || ie.Harq_ProcessNumberSizeDCI_1_2_r16 != nil || ie.Mcs_TableDCI_1_2_r16 != nil || ie.NumberOfBitsForRV_DCI_1_2_r16 != nil || ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil || ie.Prb_BundlingTypeDCI_1_2_r16 != nil || ie.PriorityIndicatorDCI_1_2_r16 != nil || ie.RateMatchPatternGroup1DCI_1_2_r16 != nil || ie.RateMatchPatternGroup2DCI_1_2_r16 != nil || ie.ResourceAllocationType1GranularityDCI_1_2_r16 != nil || ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 != nil || ie.ReferenceOfSLIVDCI_1_2_r16 != nil || ie.ResourceAllocationDCI_1_2_r16 != nil || ie.PriorityIndicatorDCI_1_1_r16 != nil || ie.DataScramblingIdentityPDSCH2_r16 != nil || ie.Pdsch_TimeDomainAllocationList_r16 != nil || ie.RepetitionSchemeConfig_r16 != nil || ie.RepetitionSchemeConfig_v1630 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil || ie.Pucch_sSCellDynDCI_1_2_r17 != nil || ie.Dl_OrJointTCI_StateList_r17 != nil || ie.BeamAppTime_r17 != nil || ie.Dummy != nil || ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil || ie.MinimumSchedulingOffsetK0_r17 != nil || ie.Harq_ProcessNumberSizeDCI_1_2_v1700 != nil || ie.Harq_ProcessNumberSizeDCI_1_1_r17 != nil || ie.Mcs_Table_r17 != nil || ie.Mcs_TableDCI_1_2_r17 != nil || ie.XOverheadMulticast_r17 != nil || ie.PriorityIndicatorDCI_4_2_r17 != nil || ie.SizeDCI_4_2_r17 != nil || ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil
	preambleBits := []bool{hasExtensions, ie.DataScramblingIdentityPDSCH != nil, ie.Dmrs_DownlinkForPDSCH_MappingTypeA != nil, ie.Dmrs_DownlinkForPDSCH_MappingTypeB != nil, len(ie.Tci_StatesToAddModList) > 0, len(ie.Tci_StatesToReleaseList) > 0, ie.Vrb_ToPRB_Interleaver != nil, ie.Pdsch_TimeDomainAllocationList != nil, ie.Pdsch_AggregationFactor != nil, len(ie.RateMatchPatternToAddModList) > 0, len(ie.RateMatchPatternToReleaseList) > 0, ie.RateMatchPatternGroup1 != nil, ie.RateMatchPatternGroup2 != nil, ie.Mcs_Table != nil, ie.MaxNrofCodeWordsScheduledByDCI != nil, ie.Prb_BundlingType != nil, len(ie.Zp_CSI_RS_ResourceToAddModList) > 0, len(ie.Zp_CSI_RS_ResourceToReleaseList) > 0, len(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList) > 0, len(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList) > 0, len(ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList) > 0, len(ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList) > 0, ie.P_ZP_CSI_RS_ResourceSet != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DataScramblingIdentityPDSCH != nil {
		if err = w.WriteInteger(*ie.DataScramblingIdentityPDSCH, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode DataScramblingIdentityPDSCH", err)
		}
	}
	if ie.Dmrs_DownlinkForPDSCH_MappingTypeA != nil {
		tmp_Dmrs_DownlinkForPDSCH_MappingTypeA := utils.SetupRelease[*DMRS_DownlinkConfig]{
			Setup: ie.Dmrs_DownlinkForPDSCH_MappingTypeA,
		}
		if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeA.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_DownlinkForPDSCH_MappingTypeA", err)
		}
	}
	if ie.Dmrs_DownlinkForPDSCH_MappingTypeB != nil {
		tmp_Dmrs_DownlinkForPDSCH_MappingTypeB := utils.SetupRelease[*DMRS_DownlinkConfig]{
			Setup: ie.Dmrs_DownlinkForPDSCH_MappingTypeB,
		}
		if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeB.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_DownlinkForPDSCH_MappingTypeB", err)
		}
	}
	if len(ie.Tci_StatesToAddModList) > 0 {
		tmp_Tci_StatesToAddModList := utils.NewSequence[*TCI_State]([]*TCI_State{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.Tci_StatesToAddModList {
			tmp_Tci_StatesToAddModList.Value = append(tmp_Tci_StatesToAddModList.Value, &i)
		}
		if err = tmp_Tci_StatesToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Tci_StatesToAddModList", err)
		}
	}
	if len(ie.Tci_StatesToReleaseList) > 0 {
		tmp_Tci_StatesToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.Tci_StatesToReleaseList {
			tmp_Tci_StatesToReleaseList.Value = append(tmp_Tci_StatesToReleaseList.Value, &i)
		}
		if err = tmp_Tci_StatesToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Tci_StatesToReleaseList", err)
		}
	}
	if ie.Vrb_ToPRB_Interleaver != nil {
		if err = ie.Vrb_ToPRB_Interleaver.Encode(w); err != nil {
			return utils.WrapError("Encode Vrb_ToPRB_Interleaver", err)
		}
	}
	if err = ie.ResourceAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceAllocation", err)
	}
	if ie.Pdsch_TimeDomainAllocationList != nil {
		tmp_Pdsch_TimeDomainAllocationList := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList]{
			Setup: ie.Pdsch_TimeDomainAllocationList,
		}
		if err = tmp_Pdsch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_TimeDomainAllocationList", err)
		}
	}
	if ie.Pdsch_AggregationFactor != nil {
		if err = ie.Pdsch_AggregationFactor.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_AggregationFactor", err)
		}
	}
	if len(ie.RateMatchPatternToAddModList) > 0 {
		tmp_RateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.RateMatchPatternToAddModList {
			tmp_RateMatchPatternToAddModList.Value = append(tmp_RateMatchPatternToAddModList.Value, &i)
		}
		if err = tmp_RateMatchPatternToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchPatternToAddModList", err)
		}
	}
	if len(ie.RateMatchPatternToReleaseList) > 0 {
		tmp_RateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.RateMatchPatternToReleaseList {
			tmp_RateMatchPatternToReleaseList.Value = append(tmp_RateMatchPatternToReleaseList.Value, &i)
		}
		if err = tmp_RateMatchPatternToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchPatternToReleaseList", err)
		}
	}
	if ie.RateMatchPatternGroup1 != nil {
		if err = ie.RateMatchPatternGroup1.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchPatternGroup1", err)
		}
	}
	if ie.RateMatchPatternGroup2 != nil {
		if err = ie.RateMatchPatternGroup2.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchPatternGroup2", err)
		}
	}
	if err = ie.Rbg_Size.Encode(w); err != nil {
		return utils.WrapError("Encode Rbg_Size", err)
	}
	if ie.Mcs_Table != nil {
		if err = ie.Mcs_Table.Encode(w); err != nil {
			return utils.WrapError("Encode Mcs_Table", err)
		}
	}
	if ie.MaxNrofCodeWordsScheduledByDCI != nil {
		if err = ie.MaxNrofCodeWordsScheduledByDCI.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNrofCodeWordsScheduledByDCI", err)
		}
	}
	if ie.Prb_BundlingType != nil {
		if err = ie.Prb_BundlingType.Encode(w); err != nil {
			return utils.WrapError("Encode Prb_BundlingType", err)
		}
	}
	if len(ie.Zp_CSI_RS_ResourceToAddModList) > 0 {
		tmp_Zp_CSI_RS_ResourceToAddModList := utils.NewSequence[*ZP_CSI_RS_Resource]([]*ZP_CSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		for _, i := range ie.Zp_CSI_RS_ResourceToAddModList {
			tmp_Zp_CSI_RS_ResourceToAddModList.Value = append(tmp_Zp_CSI_RS_ResourceToAddModList.Value, &i)
		}
		if err = tmp_Zp_CSI_RS_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Zp_CSI_RS_ResourceToAddModList", err)
		}
	}
	if len(ie.Zp_CSI_RS_ResourceToReleaseList) > 0 {
		tmp_Zp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceId]([]*ZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		for _, i := range ie.Zp_CSI_RS_ResourceToReleaseList {
			tmp_Zp_CSI_RS_ResourceToReleaseList.Value = append(tmp_Zp_CSI_RS_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_Zp_CSI_RS_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Zp_CSI_RS_ResourceToReleaseList", err)
		}
	}
	if len(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList) > 0 {
		tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList {
			tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Value = append(tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Value, &i)
		}
		if err = tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
	}
	if len(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList) > 0 {
		tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList {
			tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Value = append(tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Value, &i)
		}
		if err = tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
	}
	if len(ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList) > 0 {
		tmp_Sp_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList {
			tmp_Sp_ZP_CSI_RS_ResourceSetsToAddModList.Value = append(tmp_Sp_ZP_CSI_RS_ResourceSetsToAddModList.Value, &i)
		}
		if err = tmp_Sp_ZP_CSI_RS_ResourceSetsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
	}
	if len(ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList) > 0 {
		tmp_Sp_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList {
			tmp_Sp_ZP_CSI_RS_ResourceSetsToReleaseList.Value = append(tmp_Sp_ZP_CSI_RS_ResourceSetsToReleaseList.Value, &i)
		}
		if err = tmp_Sp_ZP_CSI_RS_ResourceSetsToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
	}
	if ie.P_ZP_CSI_RS_ResourceSet != nil {
		tmp_P_ZP_CSI_RS_ResourceSet := utils.SetupRelease[*ZP_CSI_RS_ResourceSet]{
			Setup: ie.P_ZP_CSI_RS_ResourceSet,
		}
		if err = tmp_P_ZP_CSI_RS_ResourceSet.Encode(w); err != nil {
			return utils.WrapError("Encode P_ZP_CSI_RS_ResourceSet", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.MaxMIMO_Layers_r16 != nil || ie.MinimumSchedulingOffsetK0_r16 != nil || ie.AntennaPortsFieldPresenceDCI_1_2_r16 != nil || len(ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0 || len(ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0 || ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil || ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil || ie.Dmrs_SequenceInitializationDCI_1_2_r16 != nil || ie.Harq_ProcessNumberSizeDCI_1_2_r16 != nil || ie.Mcs_TableDCI_1_2_r16 != nil || ie.NumberOfBitsForRV_DCI_1_2_r16 != nil || ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil || ie.Prb_BundlingTypeDCI_1_2_r16 != nil || ie.PriorityIndicatorDCI_1_2_r16 != nil || ie.RateMatchPatternGroup1DCI_1_2_r16 != nil || ie.RateMatchPatternGroup2DCI_1_2_r16 != nil || ie.ResourceAllocationType1GranularityDCI_1_2_r16 != nil || ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 != nil || ie.ReferenceOfSLIVDCI_1_2_r16 != nil || ie.ResourceAllocationDCI_1_2_r16 != nil || ie.PriorityIndicatorDCI_1_1_r16 != nil || ie.DataScramblingIdentityPDSCH2_r16 != nil || ie.Pdsch_TimeDomainAllocationList_r16 != nil || ie.RepetitionSchemeConfig_r16 != nil, ie.RepetitionSchemeConfig_v1630 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil || ie.Pucch_sSCellDynDCI_1_2_r17 != nil || ie.Dl_OrJointTCI_StateList_r17 != nil || ie.BeamAppTime_r17 != nil || ie.Dummy != nil || ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil || ie.MinimumSchedulingOffsetK0_r17 != nil || ie.Harq_ProcessNumberSizeDCI_1_2_v1700 != nil || ie.Harq_ProcessNumberSizeDCI_1_1_r17 != nil || ie.Mcs_Table_r17 != nil || ie.Mcs_TableDCI_1_2_r17 != nil || ie.XOverheadMulticast_r17 != nil || ie.PriorityIndicatorDCI_4_2_r17 != nil || ie.SizeDCI_4_2_r17 != nil, ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDSCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MaxMIMO_Layers_r16 != nil, ie.MinimumSchedulingOffsetK0_r16 != nil, ie.AntennaPortsFieldPresenceDCI_1_2_r16 != nil, len(ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0, len(ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0, ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil, ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil, ie.Dmrs_SequenceInitializationDCI_1_2_r16 != nil, ie.Harq_ProcessNumberSizeDCI_1_2_r16 != nil, ie.Mcs_TableDCI_1_2_r16 != nil, ie.NumberOfBitsForRV_DCI_1_2_r16 != nil, ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil, ie.Prb_BundlingTypeDCI_1_2_r16 != nil, ie.PriorityIndicatorDCI_1_2_r16 != nil, ie.RateMatchPatternGroup1DCI_1_2_r16 != nil, ie.RateMatchPatternGroup2DCI_1_2_r16 != nil, ie.ResourceAllocationType1GranularityDCI_1_2_r16 != nil, ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 != nil, ie.ReferenceOfSLIVDCI_1_2_r16 != nil, ie.ResourceAllocationDCI_1_2_r16 != nil, ie.PriorityIndicatorDCI_1_1_r16 != nil, ie.DataScramblingIdentityPDSCH2_r16 != nil, ie.Pdsch_TimeDomainAllocationList_r16 != nil, ie.RepetitionSchemeConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxMIMO_Layers_r16 optional
			if ie.MaxMIMO_Layers_r16 != nil {
				tmp_MaxMIMO_Layers_r16 := utils.SetupRelease[*MaxMIMO_LayersDL_r16]{
					Setup: ie.MaxMIMO_Layers_r16,
				}
				if err = tmp_MaxMIMO_Layers_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxMIMO_Layers_r16", err)
				}
			}
			// encode MinimumSchedulingOffsetK0_r16 optional
			if ie.MinimumSchedulingOffsetK0_r16 != nil {
				tmp_MinimumSchedulingOffsetK0_r16 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r16]{
					Setup: ie.MinimumSchedulingOffsetK0_r16,
				}
				if err = tmp_MinimumSchedulingOffsetK0_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MinimumSchedulingOffsetK0_r16", err)
				}
			}
			// encode AntennaPortsFieldPresenceDCI_1_2_r16 optional
			if ie.AntennaPortsFieldPresenceDCI_1_2_r16 != nil {
				if err = ie.AntennaPortsFieldPresenceDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AntennaPortsFieldPresenceDCI_1_2_r16", err)
				}
			}
			// encode AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 optional
			if len(ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0 {
				tmp_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				for _, i := range ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 {
					tmp_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Value = append(tmp_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Value, &i)
				}
				if err = tmp_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16", err)
				}
			}
			// encode AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 optional
			if len(ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0 {
				tmp_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				for _, i := range ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 {
					tmp_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Value = append(tmp_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Value, &i)
				}
				if err = tmp_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16", err)
				}
			}
			// encode Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 optional
			if ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil {
				tmp_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{
					Setup: ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16,
				}
				if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16", err)
				}
			}
			// encode Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 optional
			if ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil {
				tmp_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{
					Setup: ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16,
				}
				if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16", err)
				}
			}
			// encode Dmrs_SequenceInitializationDCI_1_2_r16 optional
			if ie.Dmrs_SequenceInitializationDCI_1_2_r16 != nil {
				if err = ie.Dmrs_SequenceInitializationDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_SequenceInitializationDCI_1_2_r16", err)
				}
			}
			// encode Harq_ProcessNumberSizeDCI_1_2_r16 optional
			if ie.Harq_ProcessNumberSizeDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcessNumberSizeDCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcessNumberSizeDCI_1_2_r16", err)
				}
			}
			// encode Mcs_TableDCI_1_2_r16 optional
			if ie.Mcs_TableDCI_1_2_r16 != nil {
				if err = ie.Mcs_TableDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mcs_TableDCI_1_2_r16", err)
				}
			}
			// encode NumberOfBitsForRV_DCI_1_2_r16 optional
			if ie.NumberOfBitsForRV_DCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.NumberOfBitsForRV_DCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode NumberOfBitsForRV_DCI_1_2_r16", err)
				}
			}
			// encode Pdsch_TimeDomainAllocationListDCI_1_2_r16 optional
			if ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil {
				tmp_Pdsch_TimeDomainAllocationListDCI_1_2_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16,
				}
				if err = tmp_Pdsch_TimeDomainAllocationListDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_TimeDomainAllocationListDCI_1_2_r16", err)
				}
			}
			// encode Prb_BundlingTypeDCI_1_2_r16 optional
			if ie.Prb_BundlingTypeDCI_1_2_r16 != nil {
				if err = ie.Prb_BundlingTypeDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prb_BundlingTypeDCI_1_2_r16", err)
				}
			}
			// encode PriorityIndicatorDCI_1_2_r16 optional
			if ie.PriorityIndicatorDCI_1_2_r16 != nil {
				if err = ie.PriorityIndicatorDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PriorityIndicatorDCI_1_2_r16", err)
				}
			}
			// encode RateMatchPatternGroup1DCI_1_2_r16 optional
			if ie.RateMatchPatternGroup1DCI_1_2_r16 != nil {
				if err = ie.RateMatchPatternGroup1DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RateMatchPatternGroup1DCI_1_2_r16", err)
				}
			}
			// encode RateMatchPatternGroup2DCI_1_2_r16 optional
			if ie.RateMatchPatternGroup2DCI_1_2_r16 != nil {
				if err = ie.RateMatchPatternGroup2DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RateMatchPatternGroup2DCI_1_2_r16", err)
				}
			}
			// encode ResourceAllocationType1GranularityDCI_1_2_r16 optional
			if ie.ResourceAllocationType1GranularityDCI_1_2_r16 != nil {
				if err = ie.ResourceAllocationType1GranularityDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceAllocationType1GranularityDCI_1_2_r16", err)
				}
			}
			// encode Vrb_ToPRB_InterleaverDCI_1_2_r16 optional
			if ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 != nil {
				if err = ie.Vrb_ToPRB_InterleaverDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Vrb_ToPRB_InterleaverDCI_1_2_r16", err)
				}
			}
			// encode ReferenceOfSLIVDCI_1_2_r16 optional
			if ie.ReferenceOfSLIVDCI_1_2_r16 != nil {
				if err = ie.ReferenceOfSLIVDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReferenceOfSLIVDCI_1_2_r16", err)
				}
			}
			// encode ResourceAllocationDCI_1_2_r16 optional
			if ie.ResourceAllocationDCI_1_2_r16 != nil {
				if err = ie.ResourceAllocationDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceAllocationDCI_1_2_r16", err)
				}
			}
			// encode PriorityIndicatorDCI_1_1_r16 optional
			if ie.PriorityIndicatorDCI_1_1_r16 != nil {
				if err = ie.PriorityIndicatorDCI_1_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PriorityIndicatorDCI_1_1_r16", err)
				}
			}
			// encode DataScramblingIdentityPDSCH2_r16 optional
			if ie.DataScramblingIdentityPDSCH2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.DataScramblingIdentityPDSCH2_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Encode DataScramblingIdentityPDSCH2_r16", err)
				}
			}
			// encode Pdsch_TimeDomainAllocationList_r16 optional
			if ie.Pdsch_TimeDomainAllocationList_r16 != nil {
				tmp_Pdsch_TimeDomainAllocationList_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.Pdsch_TimeDomainAllocationList_r16,
				}
				if err = tmp_Pdsch_TimeDomainAllocationList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_TimeDomainAllocationList_r16", err)
				}
			}
			// encode RepetitionSchemeConfig_r16 optional
			if ie.RepetitionSchemeConfig_r16 != nil {
				tmp_RepetitionSchemeConfig_r16 := utils.SetupRelease[*RepetitionSchemeConfig_r16]{
					Setup: ie.RepetitionSchemeConfig_r16,
				}
				if err = tmp_RepetitionSchemeConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RepetitionSchemeConfig_r16", err)
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
			optionals_ext_2 := []bool{ie.RepetitionSchemeConfig_v1630 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RepetitionSchemeConfig_v1630 optional
			if ie.RepetitionSchemeConfig_v1630 != nil {
				tmp_RepetitionSchemeConfig_v1630 := utils.SetupRelease[*RepetitionSchemeConfig_v1630]{
					Setup: ie.RepetitionSchemeConfig_v1630,
				}
				if err = tmp_RepetitionSchemeConfig_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RepetitionSchemeConfig_v1630", err)
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
			optionals_ext_3 := []bool{ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil, ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil, ie.Pucch_sSCellDynDCI_1_2_r17 != nil, ie.Dl_OrJointTCI_StateList_r17 != nil, ie.BeamAppTime_r17 != nil, ie.Dummy != nil, ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil, ie.MinimumSchedulingOffsetK0_r17 != nil, ie.Harq_ProcessNumberSizeDCI_1_2_v1700 != nil, ie.Harq_ProcessNumberSizeDCI_1_1_r17 != nil, ie.Mcs_Table_r17 != nil, ie.Mcs_TableDCI_1_2_r17 != nil, ie.XOverheadMulticast_r17 != nil, ie.PriorityIndicatorDCI_4_2_r17 != nil, ie.SizeDCI_4_2_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 optional
			if ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 optional
			if ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 optional
			if ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17", err)
				}
			}
			// encode Pdsch_HARQ_ACK_RetxDCI_1_2_r17 optional
			if ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil {
				if err = ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_HARQ_ACK_RetxDCI_1_2_r17", err)
				}
			}
			// encode Pucch_sSCellDynDCI_1_2_r17 optional
			if ie.Pucch_sSCellDynDCI_1_2_r17 != nil {
				if err = ie.Pucch_sSCellDynDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_sSCellDynDCI_1_2_r17", err)
				}
			}
			// encode Dl_OrJointTCI_StateList_r17 optional
			if ie.Dl_OrJointTCI_StateList_r17 != nil {
				if err = ie.Dl_OrJointTCI_StateList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_OrJointTCI_StateList_r17", err)
				}
			}
			// encode BeamAppTime_r17 optional
			if ie.BeamAppTime_r17 != nil {
				if err = ie.BeamAppTime_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamAppTime_r17", err)
				}
			}
			// encode Dummy optional
			if ie.Dummy != nil {
				tmp_Dummy := utils.SetupRelease[*Dummy_TDRA_List]{
					Setup: ie.Dummy,
				}
				if err = tmp_Dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy", err)
				}
			}
			// encode Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 optional
			if ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil {
				if err = ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17", err)
				}
			}
			// encode MinimumSchedulingOffsetK0_r17 optional
			if ie.MinimumSchedulingOffsetK0_r17 != nil {
				tmp_MinimumSchedulingOffsetK0_r17 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r17]{
					Setup: ie.MinimumSchedulingOffsetK0_r17,
				}
				if err = tmp_MinimumSchedulingOffsetK0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MinimumSchedulingOffsetK0_r17", err)
				}
			}
			// encode Harq_ProcessNumberSizeDCI_1_2_v1700 optional
			if ie.Harq_ProcessNumberSizeDCI_1_2_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcessNumberSizeDCI_1_2_v1700, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcessNumberSizeDCI_1_2_v1700", err)
				}
			}
			// encode Harq_ProcessNumberSizeDCI_1_1_r17 optional
			if ie.Harq_ProcessNumberSizeDCI_1_1_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcessNumberSizeDCI_1_1_r17, &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcessNumberSizeDCI_1_1_r17", err)
				}
			}
			// encode Mcs_Table_r17 optional
			if ie.Mcs_Table_r17 != nil {
				if err = ie.Mcs_Table_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mcs_Table_r17", err)
				}
			}
			// encode Mcs_TableDCI_1_2_r17 optional
			if ie.Mcs_TableDCI_1_2_r17 != nil {
				if err = ie.Mcs_TableDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mcs_TableDCI_1_2_r17", err)
				}
			}
			// encode XOverheadMulticast_r17 optional
			if ie.XOverheadMulticast_r17 != nil {
				if err = ie.XOverheadMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode XOverheadMulticast_r17", err)
				}
			}
			// encode PriorityIndicatorDCI_4_2_r17 optional
			if ie.PriorityIndicatorDCI_4_2_r17 != nil {
				if err = ie.PriorityIndicatorDCI_4_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PriorityIndicatorDCI_4_2_r17", err)
				}
			}
			// encode SizeDCI_4_2_r17 optional
			if ie.SizeDCI_4_2_r17 != nil {
				if err = extWriter.WriteInteger(*ie.SizeDCI_4_2_r17, &uper.Constraint{Lb: 20, Ub: maxDCI_4_2_Size_r17}, false); err != nil {
					return utils.WrapError("Encode SizeDCI_4_2_r17", err)
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
			optionals_ext_4 := []bool{ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 optional
			if ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil {
				tmp_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 := utils.SetupRelease[*MultiPDSCH_TDRA_List_r17]{
					Setup: ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17,
				}
				if err = tmp_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_TimeDomainAllocationListForMultiPDSCH_r17", err)
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

func (ie *PDSCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var DataScramblingIdentityPDSCHPresent bool
	if DataScramblingIdentityPDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_DownlinkForPDSCH_MappingTypeAPresent bool
	if Dmrs_DownlinkForPDSCH_MappingTypeAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_DownlinkForPDSCH_MappingTypeBPresent bool
	if Dmrs_DownlinkForPDSCH_MappingTypeBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tci_StatesToAddModListPresent bool
	if Tci_StatesToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tci_StatesToReleaseListPresent bool
	if Tci_StatesToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Vrb_ToPRB_InterleaverPresent bool
	if Vrb_ToPRB_InterleaverPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_TimeDomainAllocationListPresent bool
	if Pdsch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_AggregationFactorPresent bool
	if Pdsch_AggregationFactorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchPatternToAddModListPresent bool
	if RateMatchPatternToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchPatternToReleaseListPresent bool
	if RateMatchPatternToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchPatternGroup1Present bool
	if RateMatchPatternGroup1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchPatternGroup2Present bool
	if RateMatchPatternGroup2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mcs_TablePresent bool
	if Mcs_TablePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNrofCodeWordsScheduledByDCIPresent bool
	if MaxNrofCodeWordsScheduledByDCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Prb_BundlingTypePresent bool
	if Prb_BundlingTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Zp_CSI_RS_ResourceToAddModListPresent bool
	if Zp_CSI_RS_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Zp_CSI_RS_ResourceToReleaseListPresent bool
	if Zp_CSI_RS_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Aperiodic_ZP_CSI_RS_ResourceSetsToAddModListPresent bool
	if Aperiodic_ZP_CSI_RS_ResourceSetsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseListPresent bool
	if Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_ZP_CSI_RS_ResourceSetsToAddModListPresent bool
	if Sp_ZP_CSI_RS_ResourceSetsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_ZP_CSI_RS_ResourceSetsToReleaseListPresent bool
	if Sp_ZP_CSI_RS_ResourceSetsToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P_ZP_CSI_RS_ResourceSetPresent bool
	if P_ZP_CSI_RS_ResourceSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DataScramblingIdentityPDSCHPresent {
		var tmp_int_DataScramblingIdentityPDSCH int64
		if tmp_int_DataScramblingIdentityPDSCH, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode DataScramblingIdentityPDSCH", err)
		}
		ie.DataScramblingIdentityPDSCH = &tmp_int_DataScramblingIdentityPDSCH
	}
	if Dmrs_DownlinkForPDSCH_MappingTypeAPresent {
		tmp_Dmrs_DownlinkForPDSCH_MappingTypeA := utils.SetupRelease[*DMRS_DownlinkConfig]{}
		if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeA.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_DownlinkForPDSCH_MappingTypeA", err)
		}
		ie.Dmrs_DownlinkForPDSCH_MappingTypeA = tmp_Dmrs_DownlinkForPDSCH_MappingTypeA.Setup
	}
	if Dmrs_DownlinkForPDSCH_MappingTypeBPresent {
		tmp_Dmrs_DownlinkForPDSCH_MappingTypeB := utils.SetupRelease[*DMRS_DownlinkConfig]{}
		if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeB.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_DownlinkForPDSCH_MappingTypeB", err)
		}
		ie.Dmrs_DownlinkForPDSCH_MappingTypeB = tmp_Dmrs_DownlinkForPDSCH_MappingTypeB.Setup
	}
	if Tci_StatesToAddModListPresent {
		tmp_Tci_StatesToAddModList := utils.NewSequence[*TCI_State]([]*TCI_State{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_Tci_StatesToAddModList := func() *TCI_State {
			return new(TCI_State)
		}
		if err = tmp_Tci_StatesToAddModList.Decode(r, fn_Tci_StatesToAddModList); err != nil {
			return utils.WrapError("Decode Tci_StatesToAddModList", err)
		}
		ie.Tci_StatesToAddModList = []TCI_State{}
		for _, i := range tmp_Tci_StatesToAddModList.Value {
			ie.Tci_StatesToAddModList = append(ie.Tci_StatesToAddModList, *i)
		}
	}
	if Tci_StatesToReleaseListPresent {
		tmp_Tci_StatesToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_Tci_StatesToReleaseList := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_Tci_StatesToReleaseList.Decode(r, fn_Tci_StatesToReleaseList); err != nil {
			return utils.WrapError("Decode Tci_StatesToReleaseList", err)
		}
		ie.Tci_StatesToReleaseList = []TCI_StateId{}
		for _, i := range tmp_Tci_StatesToReleaseList.Value {
			ie.Tci_StatesToReleaseList = append(ie.Tci_StatesToReleaseList, *i)
		}
	}
	if Vrb_ToPRB_InterleaverPresent {
		ie.Vrb_ToPRB_Interleaver = new(PDSCH_Config_vrb_ToPRB_Interleaver)
		if err = ie.Vrb_ToPRB_Interleaver.Decode(r); err != nil {
			return utils.WrapError("Decode Vrb_ToPRB_Interleaver", err)
		}
	}
	if err = ie.ResourceAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceAllocation", err)
	}
	if Pdsch_TimeDomainAllocationListPresent {
		tmp_Pdsch_TimeDomainAllocationList := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList]{}
		if err = tmp_Pdsch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_TimeDomainAllocationList", err)
		}
		ie.Pdsch_TimeDomainAllocationList = tmp_Pdsch_TimeDomainAllocationList.Setup
	}
	if Pdsch_AggregationFactorPresent {
		ie.Pdsch_AggregationFactor = new(PDSCH_Config_pdsch_AggregationFactor)
		if err = ie.Pdsch_AggregationFactor.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_AggregationFactor", err)
		}
	}
	if RateMatchPatternToAddModListPresent {
		tmp_RateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_RateMatchPatternToAddModList := func() *RateMatchPattern {
			return new(RateMatchPattern)
		}
		if err = tmp_RateMatchPatternToAddModList.Decode(r, fn_RateMatchPatternToAddModList); err != nil {
			return utils.WrapError("Decode RateMatchPatternToAddModList", err)
		}
		ie.RateMatchPatternToAddModList = []RateMatchPattern{}
		for _, i := range tmp_RateMatchPatternToAddModList.Value {
			ie.RateMatchPatternToAddModList = append(ie.RateMatchPatternToAddModList, *i)
		}
	}
	if RateMatchPatternToReleaseListPresent {
		tmp_RateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_RateMatchPatternToReleaseList := func() *RateMatchPatternId {
			return new(RateMatchPatternId)
		}
		if err = tmp_RateMatchPatternToReleaseList.Decode(r, fn_RateMatchPatternToReleaseList); err != nil {
			return utils.WrapError("Decode RateMatchPatternToReleaseList", err)
		}
		ie.RateMatchPatternToReleaseList = []RateMatchPatternId{}
		for _, i := range tmp_RateMatchPatternToReleaseList.Value {
			ie.RateMatchPatternToReleaseList = append(ie.RateMatchPatternToReleaseList, *i)
		}
	}
	if RateMatchPatternGroup1Present {
		ie.RateMatchPatternGroup1 = new(RateMatchPatternGroup)
		if err = ie.RateMatchPatternGroup1.Decode(r); err != nil {
			return utils.WrapError("Decode RateMatchPatternGroup1", err)
		}
	}
	if RateMatchPatternGroup2Present {
		ie.RateMatchPatternGroup2 = new(RateMatchPatternGroup)
		if err = ie.RateMatchPatternGroup2.Decode(r); err != nil {
			return utils.WrapError("Decode RateMatchPatternGroup2", err)
		}
	}
	if err = ie.Rbg_Size.Decode(r); err != nil {
		return utils.WrapError("Decode Rbg_Size", err)
	}
	if Mcs_TablePresent {
		ie.Mcs_Table = new(PDSCH_Config_mcs_Table)
		if err = ie.Mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode Mcs_Table", err)
		}
	}
	if MaxNrofCodeWordsScheduledByDCIPresent {
		ie.MaxNrofCodeWordsScheduledByDCI = new(PDSCH_Config_maxNrofCodeWordsScheduledByDCI)
		if err = ie.MaxNrofCodeWordsScheduledByDCI.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNrofCodeWordsScheduledByDCI", err)
		}
	}
	if Prb_BundlingTypePresent {
		ie.Prb_BundlingType = new(PDSCH_Config_prb_BundlingType)
		if err = ie.Prb_BundlingType.Decode(r); err != nil {
			return utils.WrapError("Decode Prb_BundlingType", err)
		}
	}
	if Zp_CSI_RS_ResourceToAddModListPresent {
		tmp_Zp_CSI_RS_ResourceToAddModList := utils.NewSequence[*ZP_CSI_RS_Resource]([]*ZP_CSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		fn_Zp_CSI_RS_ResourceToAddModList := func() *ZP_CSI_RS_Resource {
			return new(ZP_CSI_RS_Resource)
		}
		if err = tmp_Zp_CSI_RS_ResourceToAddModList.Decode(r, fn_Zp_CSI_RS_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode Zp_CSI_RS_ResourceToAddModList", err)
		}
		ie.Zp_CSI_RS_ResourceToAddModList = []ZP_CSI_RS_Resource{}
		for _, i := range tmp_Zp_CSI_RS_ResourceToAddModList.Value {
			ie.Zp_CSI_RS_ResourceToAddModList = append(ie.Zp_CSI_RS_ResourceToAddModList, *i)
		}
	}
	if Zp_CSI_RS_ResourceToReleaseListPresent {
		tmp_Zp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceId]([]*ZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		fn_Zp_CSI_RS_ResourceToReleaseList := func() *ZP_CSI_RS_ResourceId {
			return new(ZP_CSI_RS_ResourceId)
		}
		if err = tmp_Zp_CSI_RS_ResourceToReleaseList.Decode(r, fn_Zp_CSI_RS_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode Zp_CSI_RS_ResourceToReleaseList", err)
		}
		ie.Zp_CSI_RS_ResourceToReleaseList = []ZP_CSI_RS_ResourceId{}
		for _, i := range tmp_Zp_CSI_RS_ResourceToReleaseList.Value {
			ie.Zp_CSI_RS_ResourceToReleaseList = append(ie.Zp_CSI_RS_ResourceToReleaseList, *i)
		}
	}
	if Aperiodic_ZP_CSI_RS_ResourceSetsToAddModListPresent {
		tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList := func() *ZP_CSI_RS_ResourceSet {
			return new(ZP_CSI_RS_ResourceSet)
		}
		if err = tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Decode(r, fn_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList); err != nil {
			return utils.WrapError("Decode Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
		ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList = []ZP_CSI_RS_ResourceSet{}
		for _, i := range tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Value {
			ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList = append(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList, *i)
		}
	}
	if Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseListPresent {
		tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList := func() *ZP_CSI_RS_ResourceSetId {
			return new(ZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Decode(r, fn_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList); err != nil {
			return utils.WrapError("Decode Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
		ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList = []ZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Value {
			ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList = append(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList, *i)
		}
	}
	if Sp_ZP_CSI_RS_ResourceSetsToAddModListPresent {
		tmp_Sp_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_Sp_ZP_CSI_RS_ResourceSetsToAddModList := func() *ZP_CSI_RS_ResourceSet {
			return new(ZP_CSI_RS_ResourceSet)
		}
		if err = tmp_Sp_ZP_CSI_RS_ResourceSetsToAddModList.Decode(r, fn_Sp_ZP_CSI_RS_ResourceSetsToAddModList); err != nil {
			return utils.WrapError("Decode Sp_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
		ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList = []ZP_CSI_RS_ResourceSet{}
		for _, i := range tmp_Sp_ZP_CSI_RS_ResourceSetsToAddModList.Value {
			ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList = append(ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList, *i)
		}
	}
	if Sp_ZP_CSI_RS_ResourceSetsToReleaseListPresent {
		tmp_Sp_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_Sp_ZP_CSI_RS_ResourceSetsToReleaseList := func() *ZP_CSI_RS_ResourceSetId {
			return new(ZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_Sp_ZP_CSI_RS_ResourceSetsToReleaseList.Decode(r, fn_Sp_ZP_CSI_RS_ResourceSetsToReleaseList); err != nil {
			return utils.WrapError("Decode Sp_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
		ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList = []ZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_Sp_ZP_CSI_RS_ResourceSetsToReleaseList.Value {
			ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList = append(ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList, *i)
		}
	}
	if P_ZP_CSI_RS_ResourceSetPresent {
		tmp_P_ZP_CSI_RS_ResourceSet := utils.SetupRelease[*ZP_CSI_RS_ResourceSet]{}
		if err = tmp_P_ZP_CSI_RS_ResourceSet.Decode(r); err != nil {
			return utils.WrapError("Decode P_ZP_CSI_RS_ResourceSet", err)
		}
		ie.P_ZP_CSI_RS_ResourceSet = tmp_P_ZP_CSI_RS_ResourceSet.Setup
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			MaxMIMO_Layers_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MinimumSchedulingOffsetK0_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AntennaPortsFieldPresenceDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_SequenceInitializationDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcessNumberSizeDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mcs_TableDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NumberOfBitsForRV_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_TimeDomainAllocationListDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prb_BundlingTypeDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PriorityIndicatorDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RateMatchPatternGroup1DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RateMatchPatternGroup2DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ResourceAllocationType1GranularityDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Vrb_ToPRB_InterleaverDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ReferenceOfSLIVDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ResourceAllocationDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PriorityIndicatorDCI_1_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DataScramblingIdentityPDSCH2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_TimeDomainAllocationList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RepetitionSchemeConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxMIMO_Layers_r16 optional
			if MaxMIMO_Layers_r16Present {
				tmp_MaxMIMO_Layers_r16 := utils.SetupRelease[*MaxMIMO_LayersDL_r16]{}
				if err = tmp_MaxMIMO_Layers_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxMIMO_Layers_r16", err)
				}
				ie.MaxMIMO_Layers_r16 = tmp_MaxMIMO_Layers_r16.Setup
			}
			// decode MinimumSchedulingOffsetK0_r16 optional
			if MinimumSchedulingOffsetK0_r16Present {
				tmp_MinimumSchedulingOffsetK0_r16 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r16]{}
				if err = tmp_MinimumSchedulingOffsetK0_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MinimumSchedulingOffsetK0_r16", err)
				}
				ie.MinimumSchedulingOffsetK0_r16 = tmp_MinimumSchedulingOffsetK0_r16.Setup
			}
			// decode AntennaPortsFieldPresenceDCI_1_2_r16 optional
			if AntennaPortsFieldPresenceDCI_1_2_r16Present {
				ie.AntennaPortsFieldPresenceDCI_1_2_r16 = new(PDSCH_Config_antennaPortsFieldPresenceDCI_1_2_r16)
				if err = ie.AntennaPortsFieldPresenceDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AntennaPortsFieldPresenceDCI_1_2_r16", err)
				}
			}
			// decode AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 optional
			if AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16Present {
				tmp_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				fn_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 := func() *ZP_CSI_RS_ResourceSet {
					return new(ZP_CSI_RS_ResourceSet)
				}
				if err = tmp_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Decode(extReader, fn_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16); err != nil {
					return utils.WrapError("Decode AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16", err)
				}
				ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 = []ZP_CSI_RS_ResourceSet{}
				for _, i := range tmp_AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Value {
					ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 = append(ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16, *i)
				}
			}
			// decode AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 optional
			if AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16Present {
				tmp_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				fn_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 := func() *ZP_CSI_RS_ResourceSetId {
					return new(ZP_CSI_RS_ResourceSetId)
				}
				if err = tmp_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Decode(extReader, fn_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16); err != nil {
					return utils.WrapError("Decode AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16", err)
				}
				ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 = []ZP_CSI_RS_ResourceSetId{}
				for _, i := range tmp_AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Value {
					ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 = append(ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16, *i)
				}
			}
			// decode Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 optional
			if Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16Present {
				tmp_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{}
				if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16", err)
				}
				ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 = tmp_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16.Setup
			}
			// decode Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 optional
			if Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16Present {
				tmp_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{}
				if err = tmp_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16", err)
				}
				ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 = tmp_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16.Setup
			}
			// decode Dmrs_SequenceInitializationDCI_1_2_r16 optional
			if Dmrs_SequenceInitializationDCI_1_2_r16Present {
				ie.Dmrs_SequenceInitializationDCI_1_2_r16 = new(PDSCH_Config_dmrs_SequenceInitializationDCI_1_2_r16)
				if err = ie.Dmrs_SequenceInitializationDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_SequenceInitializationDCI_1_2_r16", err)
				}
			}
			// decode Harq_ProcessNumberSizeDCI_1_2_r16 optional
			if Harq_ProcessNumberSizeDCI_1_2_r16Present {
				var tmp_int_Harq_ProcessNumberSizeDCI_1_2_r16 int64
				if tmp_int_Harq_ProcessNumberSizeDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcessNumberSizeDCI_1_2_r16", err)
				}
				ie.Harq_ProcessNumberSizeDCI_1_2_r16 = &tmp_int_Harq_ProcessNumberSizeDCI_1_2_r16
			}
			// decode Mcs_TableDCI_1_2_r16 optional
			if Mcs_TableDCI_1_2_r16Present {
				ie.Mcs_TableDCI_1_2_r16 = new(PDSCH_Config_mcs_TableDCI_1_2_r16)
				if err = ie.Mcs_TableDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mcs_TableDCI_1_2_r16", err)
				}
			}
			// decode NumberOfBitsForRV_DCI_1_2_r16 optional
			if NumberOfBitsForRV_DCI_1_2_r16Present {
				var tmp_int_NumberOfBitsForRV_DCI_1_2_r16 int64
				if tmp_int_NumberOfBitsForRV_DCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode NumberOfBitsForRV_DCI_1_2_r16", err)
				}
				ie.NumberOfBitsForRV_DCI_1_2_r16 = &tmp_int_NumberOfBitsForRV_DCI_1_2_r16
			}
			// decode Pdsch_TimeDomainAllocationListDCI_1_2_r16 optional
			if Pdsch_TimeDomainAllocationListDCI_1_2_r16Present {
				tmp_Pdsch_TimeDomainAllocationListDCI_1_2_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_Pdsch_TimeDomainAllocationListDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_TimeDomainAllocationListDCI_1_2_r16", err)
				}
				ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 = tmp_Pdsch_TimeDomainAllocationListDCI_1_2_r16.Setup
			}
			// decode Prb_BundlingTypeDCI_1_2_r16 optional
			if Prb_BundlingTypeDCI_1_2_r16Present {
				ie.Prb_BundlingTypeDCI_1_2_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16)
				if err = ie.Prb_BundlingTypeDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prb_BundlingTypeDCI_1_2_r16", err)
				}
			}
			// decode PriorityIndicatorDCI_1_2_r16 optional
			if PriorityIndicatorDCI_1_2_r16Present {
				ie.PriorityIndicatorDCI_1_2_r16 = new(PDSCH_Config_priorityIndicatorDCI_1_2_r16)
				if err = ie.PriorityIndicatorDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PriorityIndicatorDCI_1_2_r16", err)
				}
			}
			// decode RateMatchPatternGroup1DCI_1_2_r16 optional
			if RateMatchPatternGroup1DCI_1_2_r16Present {
				ie.RateMatchPatternGroup1DCI_1_2_r16 = new(RateMatchPatternGroup)
				if err = ie.RateMatchPatternGroup1DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RateMatchPatternGroup1DCI_1_2_r16", err)
				}
			}
			// decode RateMatchPatternGroup2DCI_1_2_r16 optional
			if RateMatchPatternGroup2DCI_1_2_r16Present {
				ie.RateMatchPatternGroup2DCI_1_2_r16 = new(RateMatchPatternGroup)
				if err = ie.RateMatchPatternGroup2DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RateMatchPatternGroup2DCI_1_2_r16", err)
				}
			}
			// decode ResourceAllocationType1GranularityDCI_1_2_r16 optional
			if ResourceAllocationType1GranularityDCI_1_2_r16Present {
				ie.ResourceAllocationType1GranularityDCI_1_2_r16 = new(PDSCH_Config_resourceAllocationType1GranularityDCI_1_2_r16)
				if err = ie.ResourceAllocationType1GranularityDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ResourceAllocationType1GranularityDCI_1_2_r16", err)
				}
			}
			// decode Vrb_ToPRB_InterleaverDCI_1_2_r16 optional
			if Vrb_ToPRB_InterleaverDCI_1_2_r16Present {
				ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 = new(PDSCH_Config_vrb_ToPRB_InterleaverDCI_1_2_r16)
				if err = ie.Vrb_ToPRB_InterleaverDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Vrb_ToPRB_InterleaverDCI_1_2_r16", err)
				}
			}
			// decode ReferenceOfSLIVDCI_1_2_r16 optional
			if ReferenceOfSLIVDCI_1_2_r16Present {
				ie.ReferenceOfSLIVDCI_1_2_r16 = new(PDSCH_Config_referenceOfSLIVDCI_1_2_r16)
				if err = ie.ReferenceOfSLIVDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReferenceOfSLIVDCI_1_2_r16", err)
				}
			}
			// decode ResourceAllocationDCI_1_2_r16 optional
			if ResourceAllocationDCI_1_2_r16Present {
				ie.ResourceAllocationDCI_1_2_r16 = new(PDSCH_Config_resourceAllocationDCI_1_2_r16)
				if err = ie.ResourceAllocationDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ResourceAllocationDCI_1_2_r16", err)
				}
			}
			// decode PriorityIndicatorDCI_1_1_r16 optional
			if PriorityIndicatorDCI_1_1_r16Present {
				ie.PriorityIndicatorDCI_1_1_r16 = new(PDSCH_Config_priorityIndicatorDCI_1_1_r16)
				if err = ie.PriorityIndicatorDCI_1_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PriorityIndicatorDCI_1_1_r16", err)
				}
			}
			// decode DataScramblingIdentityPDSCH2_r16 optional
			if DataScramblingIdentityPDSCH2_r16Present {
				var tmp_int_DataScramblingIdentityPDSCH2_r16 int64
				if tmp_int_DataScramblingIdentityPDSCH2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Decode DataScramblingIdentityPDSCH2_r16", err)
				}
				ie.DataScramblingIdentityPDSCH2_r16 = &tmp_int_DataScramblingIdentityPDSCH2_r16
			}
			// decode Pdsch_TimeDomainAllocationList_r16 optional
			if Pdsch_TimeDomainAllocationList_r16Present {
				tmp_Pdsch_TimeDomainAllocationList_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_Pdsch_TimeDomainAllocationList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_TimeDomainAllocationList_r16", err)
				}
				ie.Pdsch_TimeDomainAllocationList_r16 = tmp_Pdsch_TimeDomainAllocationList_r16.Setup
			}
			// decode RepetitionSchemeConfig_r16 optional
			if RepetitionSchemeConfig_r16Present {
				tmp_RepetitionSchemeConfig_r16 := utils.SetupRelease[*RepetitionSchemeConfig_r16]{}
				if err = tmp_RepetitionSchemeConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RepetitionSchemeConfig_r16", err)
				}
				ie.RepetitionSchemeConfig_r16 = tmp_RepetitionSchemeConfig_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			RepetitionSchemeConfig_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RepetitionSchemeConfig_v1630 optional
			if RepetitionSchemeConfig_v1630Present {
				tmp_RepetitionSchemeConfig_v1630 := utils.SetupRelease[*RepetitionSchemeConfig_v1630]{}
				if err = tmp_RepetitionSchemeConfig_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode RepetitionSchemeConfig_v1630", err)
				}
				ie.RepetitionSchemeConfig_v1630 = tmp_RepetitionSchemeConfig_v1630.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_HARQ_ACK_RetxDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_sSCellDynDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_OrJointTCI_StateList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamAppTime_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MinimumSchedulingOffsetK0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcessNumberSizeDCI_1_2_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcessNumberSizeDCI_1_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mcs_Table_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mcs_TableDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			XOverheadMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PriorityIndicatorDCI_4_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SizeDCI_4_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 optional
			if Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17Present {
				ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17)
				if err = ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17", err)
				}
			}
			// decode Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 optional
			if Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17Present {
				ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_1_2_r17)
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17", err)
				}
			}
			// decode Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 optional
			if Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17Present {
				ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17)
				if err = ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17", err)
				}
			}
			// decode Pdsch_HARQ_ACK_RetxDCI_1_2_r17 optional
			if Pdsch_HARQ_ACK_RetxDCI_1_2_r17Present {
				ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_RetxDCI_1_2_r17)
				if err = ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_HARQ_ACK_RetxDCI_1_2_r17", err)
				}
			}
			// decode Pucch_sSCellDynDCI_1_2_r17 optional
			if Pucch_sSCellDynDCI_1_2_r17Present {
				ie.Pucch_sSCellDynDCI_1_2_r17 = new(PDSCH_Config_pucch_sSCellDynDCI_1_2_r17)
				if err = ie.Pucch_sSCellDynDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_sSCellDynDCI_1_2_r17", err)
				}
			}
			// decode Dl_OrJointTCI_StateList_r17 optional
			if Dl_OrJointTCI_StateList_r17Present {
				ie.Dl_OrJointTCI_StateList_r17 = new(PDSCH_Config_dl_OrJointTCI_StateList_r17)
				if err = ie.Dl_OrJointTCI_StateList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_OrJointTCI_StateList_r17", err)
				}
			}
			// decode BeamAppTime_r17 optional
			if BeamAppTime_r17Present {
				ie.BeamAppTime_r17 = new(PDSCH_Config_beamAppTime_r17)
				if err = ie.BeamAppTime_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamAppTime_r17", err)
				}
			}
			// decode Dummy optional
			if DummyPresent {
				tmp_Dummy := utils.SetupRelease[*Dummy_TDRA_List]{}
				if err = tmp_Dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dummy", err)
				}
				ie.Dummy = tmp_Dummy.Setup
			}
			// decode Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 optional
			if Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17Present {
				ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 = new(PDSCH_Config_dmrs_FD_OCC_DisabledForRank1_PDSCH_r17)
				if err = ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17", err)
				}
			}
			// decode MinimumSchedulingOffsetK0_r17 optional
			if MinimumSchedulingOffsetK0_r17Present {
				tmp_MinimumSchedulingOffsetK0_r17 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r17]{}
				if err = tmp_MinimumSchedulingOffsetK0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MinimumSchedulingOffsetK0_r17", err)
				}
				ie.MinimumSchedulingOffsetK0_r17 = tmp_MinimumSchedulingOffsetK0_r17.Setup
			}
			// decode Harq_ProcessNumberSizeDCI_1_2_v1700 optional
			if Harq_ProcessNumberSizeDCI_1_2_v1700Present {
				var tmp_int_Harq_ProcessNumberSizeDCI_1_2_v1700 int64
				if tmp_int_Harq_ProcessNumberSizeDCI_1_2_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcessNumberSizeDCI_1_2_v1700", err)
				}
				ie.Harq_ProcessNumberSizeDCI_1_2_v1700 = &tmp_int_Harq_ProcessNumberSizeDCI_1_2_v1700
			}
			// decode Harq_ProcessNumberSizeDCI_1_1_r17 optional
			if Harq_ProcessNumberSizeDCI_1_1_r17Present {
				var tmp_int_Harq_ProcessNumberSizeDCI_1_1_r17 int64
				if tmp_int_Harq_ProcessNumberSizeDCI_1_1_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcessNumberSizeDCI_1_1_r17", err)
				}
				ie.Harq_ProcessNumberSizeDCI_1_1_r17 = &tmp_int_Harq_ProcessNumberSizeDCI_1_1_r17
			}
			// decode Mcs_Table_r17 optional
			if Mcs_Table_r17Present {
				ie.Mcs_Table_r17 = new(PDSCH_Config_mcs_Table_r17)
				if err = ie.Mcs_Table_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mcs_Table_r17", err)
				}
			}
			// decode Mcs_TableDCI_1_2_r17 optional
			if Mcs_TableDCI_1_2_r17Present {
				ie.Mcs_TableDCI_1_2_r17 = new(PDSCH_Config_mcs_TableDCI_1_2_r17)
				if err = ie.Mcs_TableDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mcs_TableDCI_1_2_r17", err)
				}
			}
			// decode XOverheadMulticast_r17 optional
			if XOverheadMulticast_r17Present {
				ie.XOverheadMulticast_r17 = new(PDSCH_Config_xOverheadMulticast_r17)
				if err = ie.XOverheadMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode XOverheadMulticast_r17", err)
				}
			}
			// decode PriorityIndicatorDCI_4_2_r17 optional
			if PriorityIndicatorDCI_4_2_r17Present {
				ie.PriorityIndicatorDCI_4_2_r17 = new(PDSCH_Config_priorityIndicatorDCI_4_2_r17)
				if err = ie.PriorityIndicatorDCI_4_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PriorityIndicatorDCI_4_2_r17", err)
				}
			}
			// decode SizeDCI_4_2_r17 optional
			if SizeDCI_4_2_r17Present {
				var tmp_int_SizeDCI_4_2_r17 int64
				if tmp_int_SizeDCI_4_2_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 20, Ub: maxDCI_4_2_Size_r17}, false); err != nil {
					return utils.WrapError("Decode SizeDCI_4_2_r17", err)
				}
				ie.SizeDCI_4_2_r17 = &tmp_int_SizeDCI_4_2_r17
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pdsch_TimeDomainAllocationListForMultiPDSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 optional
			if Pdsch_TimeDomainAllocationListForMultiPDSCH_r17Present {
				tmp_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 := utils.SetupRelease[*MultiPDSCH_TDRA_List_r17]{}
				if err = tmp_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_TimeDomainAllocationListForMultiPDSCH_r17", err)
				}
				ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 = tmp_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17.Setup
			}
		}
	}
	return nil
}
