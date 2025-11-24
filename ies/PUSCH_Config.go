package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_Config struct {
	DataScramblingIdentityPUSCH                     *int64                                                      `lb:0,ub:1023,optional`
	TxConfig                                        *PUSCH_Config_txConfig                                      `optional`
	Dmrs_UplinkForPUSCH_MappingTypeA                *DMRS_UplinkConfig                                          `optional,setuprelease`
	Dmrs_UplinkForPUSCH_MappingTypeB                *DMRS_UplinkConfig                                          `optional,setuprelease`
	Pusch_PowerControl                              *PUSCH_PowerControl                                         `optional`
	FrequencyHopping                                *PUSCH_Config_frequencyHopping                              `optional`
	FrequencyHoppingOffsetLists                     []int64                                                     `lb:1,ub:4,e_lb:1,e_ub:maxNrofPhysicalResourceBlocks_1,optional`
	ResourceAllocation                              PUSCH_Config_resourceAllocation                             `madatory`
	Pusch_TimeDomainAllocationList                  *PUSCH_TimeDomainResourceAllocationList                     `optional,setuprelease`
	Pusch_AggregationFactor                         *PUSCH_Config_pusch_AggregationFactor                       `optional`
	Mcs_Table                                       *PUSCH_Config_mcs_Table                                     `optional`
	Mcs_TableTransformPrecoder                      *PUSCH_Config_mcs_TableTransformPrecoder                    `optional`
	TransformPrecoder                               *PUSCH_Config_transformPrecoder                             `optional`
	CodebookSubset                                  *PUSCH_Config_codebookSubset                                `optional`
	MaxRank                                         *int64                                                      `lb:1,ub:4,optional`
	Rbg_Size                                        *PUSCH_Config_rbg_Size                                      `optional`
	Uci_OnPUSCH                                     *UCI_OnPUSCH                                                `optional,setuprelease`
	Tp_pi2BPSK                                      *PUSCH_Config_tp_pi2BPSK                                    `optional`
	MinimumSchedulingOffsetK2_r16                   *MinSchedulingOffsetK2_Values_r16                           `optional,ext-1,setuprelease`
	Ul_AccessConfigListDCI_0_1_r16                  *UL_AccessConfigListDCI_0_1_r16                             `optional,ext-1,setuprelease`
	Harq_ProcessNumberSizeDCI_0_2_r16               *int64                                                      `lb:0,ub:4,optional,ext-1`
	Dmrs_SequenceInitializationDCI_0_2_r16          *PUSCH_Config_dmrs_SequenceInitializationDCI_0_2_r16        `optional,ext-1`
	NumberOfBitsForRV_DCI_0_2_r16                   *int64                                                      `lb:0,ub:2,optional,ext-1`
	AntennaPortsFieldPresenceDCI_0_2_r16            *PUSCH_Config_antennaPortsFieldPresenceDCI_0_2_r16          `optional,ext-1`
	Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16    *DMRS_UplinkConfig                                          `optional,ext-1,setuprelease`
	Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16    *DMRS_UplinkConfig                                          `optional,ext-1,setuprelease`
	FrequencyHoppingDCI_0_2_r16                     *PUSCH_Config_frequencyHoppingDCI_0_2_r16                   `optional,ext-1`
	FrequencyHoppingOffsetListsDCI_0_2_r16          *FrequencyHoppingOffsetListsDCI_0_2_r16                     `optional,ext-1,setuprelease`
	CodebookSubsetDCI_0_2_r16                       *PUSCH_Config_codebookSubsetDCI_0_2_r16                     `optional,ext-1`
	InvalidSymbolPatternIndicatorDCI_0_2_r16        *PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_2_r16      `optional,ext-1`
	MaxRankDCI_0_2_r16                              *int64                                                      `lb:1,ub:4,optional,ext-1`
	Mcs_TableDCI_0_2_r16                            *PUSCH_Config_mcs_TableDCI_0_2_r16                          `optional,ext-1`
	Mcs_TableTransformPrecoderDCI_0_2_r16           *PUSCH_Config_mcs_TableTransformPrecoderDCI_0_2_r16         `optional,ext-1`
	PriorityIndicatorDCI_0_2_r16                    *PUSCH_Config_priorityIndicatorDCI_0_2_r16                  `optional,ext-1`
	Pusch_RepTypeIndicatorDCI_0_2_r16               *PUSCH_Config_pusch_RepTypeIndicatorDCI_0_2_r16             `optional,ext-1`
	ResourceAllocationDCI_0_2_r16                   *PUSCH_Config_resourceAllocationDCI_0_2_r16                 `optional,ext-1`
	ResourceAllocationType1GranularityDCI_0_2_r16   *PUSCH_Config_resourceAllocationType1GranularityDCI_0_2_r16 `optional,ext-1`
	Uci_OnPUSCH_ListDCI_0_2_r16                     *UCI_OnPUSCH_ListDCI_0_2_r16                                `optional,ext-1,setuprelease`
	Pusch_TimeDomainAllocationListDCI_0_2_r16       *PUSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	Pusch_TimeDomainAllocationListDCI_0_1_r16       *PUSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	InvalidSymbolPatternIndicatorDCI_0_1_r16        *PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_1_r16      `optional,ext-1`
	PriorityIndicatorDCI_0_1_r16                    *PUSCH_Config_priorityIndicatorDCI_0_1_r16                  `optional,ext-1`
	Pusch_RepTypeIndicatorDCI_0_1_r16               *PUSCH_Config_pusch_RepTypeIndicatorDCI_0_1_r16             `optional,ext-1`
	FrequencyHoppingDCI_0_1_r16                     *PUSCH_Config_frequencyHoppingDCI_0_1_r16                   `optional,ext-1`
	Uci_OnPUSCH_ListDCI_0_1_r16                     *UCI_OnPUSCH_ListDCI_0_1_r16                                `optional,ext-1,setuprelease`
	InvalidSymbolPattern_r16                        *InvalidSymbolPattern_r16                                   `optional,ext-1`
	Pusch_PowerControl_v1610                        *PUSCH_PowerControl_v1610                                   `optional,ext-1,setuprelease`
	Ul_FullPowerTransmission_r16                    *PUSCH_Config_ul_FullPowerTransmission_r16                  `optional,ext-1`
	Pusch_TimeDomainAllocationListForMultiPUSCH_r16 *PUSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	NumberOfInvalidSymbolsForDL_UL_Switching_r16    *int64                                                      `lb:1,ub:4,optional,ext-1`
	Ul_AccessConfigListDCI_0_2_r17                  *UL_AccessConfigListDCI_0_2_r17                             `optional,ext-2,setuprelease`
	BetaOffsetsCrossPri0_r17                        *BetaOffsetsCrossPriSel_r17                                 `optional,ext-2,setuprelease`
	BetaOffsetsCrossPri1_r17                        *BetaOffsetsCrossPriSel_r17                                 `optional,ext-2,setuprelease`
	BetaOffsetsCrossPri0DCI_0_2_r17                 *BetaOffsetsCrossPriSelDCI_0_2_r17                          `optional,ext-2,setuprelease`
	BetaOffsetsCrossPri1DCI_0_2_r17                 *BetaOffsetsCrossPriSelDCI_0_2_r17                          `optional,ext-2,setuprelease`
	MappingPattern_r17                              *PUSCH_Config_mappingPattern_r17                            `optional,ext-2`
	SecondTPCFieldDCI_0_1_r17                       *PUSCH_Config_secondTPCFieldDCI_0_1_r17                     `optional,ext-2`
	SecondTPCFieldDCI_0_2_r17                       *PUSCH_Config_secondTPCFieldDCI_0_2_r17                     `optional,ext-2`
	SequenceOffsetForRV_r17                         *int64                                                      `lb:0,ub:3,optional,ext-2`
	Ul_AccessConfigListDCI_0_1_r17                  *UL_AccessConfigListDCI_0_1_r17                             `optional,ext-2,setuprelease`
	MinimumSchedulingOffsetK2_r17                   *MinSchedulingOffsetK2_Values_r17                           `optional,ext-2,setuprelease`
	AvailableSlotCounting_r17                       *PUSCH_Config_availableSlotCounting_r17                     `optional,ext-2`
	Dmrs_BundlingPUSCH_Config_r17                   *DMRS_BundlingPUSCH_Config_r17                              `optional,ext-2,setuprelease`
	Harq_ProcessNumberSizeDCI_0_2_v1700             *int64                                                      `lb:5,ub:5,optional,ext-2`
	Harq_ProcessNumberSizeDCI_0_1_r17               *int64                                                      `lb:5,ub:5,optional,ext-2`
	Mpe_ResourcePoolToAddModList_r17                []MPE_Resource_r17                                          `lb:1,ub:maxMPE_Resources_r17,optional,ext-2`
	Mpe_ResourcePoolToReleaseList_r17               []MPE_ResourceId_r17                                        `lb:1,ub:maxMPE_Resources_r17,optional,ext-2`
}

func (ie *PUSCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.MinimumSchedulingOffsetK2_r16 != nil || ie.Ul_AccessConfigListDCI_0_1_r16 != nil || ie.Harq_ProcessNumberSizeDCI_0_2_r16 != nil || ie.Dmrs_SequenceInitializationDCI_0_2_r16 != nil || ie.NumberOfBitsForRV_DCI_0_2_r16 != nil || ie.AntennaPortsFieldPresenceDCI_0_2_r16 != nil || ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil || ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil || ie.FrequencyHoppingDCI_0_2_r16 != nil || ie.FrequencyHoppingOffsetListsDCI_0_2_r16 != nil || ie.CodebookSubsetDCI_0_2_r16 != nil || ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 != nil || ie.MaxRankDCI_0_2_r16 != nil || ie.Mcs_TableDCI_0_2_r16 != nil || ie.Mcs_TableTransformPrecoderDCI_0_2_r16 != nil || ie.PriorityIndicatorDCI_0_2_r16 != nil || ie.Pusch_RepTypeIndicatorDCI_0_2_r16 != nil || ie.ResourceAllocationDCI_0_2_r16 != nil || ie.ResourceAllocationType1GranularityDCI_0_2_r16 != nil || ie.Uci_OnPUSCH_ListDCI_0_2_r16 != nil || ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 != nil || ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 != nil || ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 != nil || ie.PriorityIndicatorDCI_0_1_r16 != nil || ie.Pusch_RepTypeIndicatorDCI_0_1_r16 != nil || ie.FrequencyHoppingDCI_0_1_r16 != nil || ie.Uci_OnPUSCH_ListDCI_0_1_r16 != nil || ie.InvalidSymbolPattern_r16 != nil || ie.Pusch_PowerControl_v1610 != nil || ie.Ul_FullPowerTransmission_r16 != nil || ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil || ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 != nil || ie.Ul_AccessConfigListDCI_0_2_r17 != nil || ie.BetaOffsetsCrossPri0_r17 != nil || ie.BetaOffsetsCrossPri1_r17 != nil || ie.BetaOffsetsCrossPri0DCI_0_2_r17 != nil || ie.BetaOffsetsCrossPri1DCI_0_2_r17 != nil || ie.MappingPattern_r17 != nil || ie.SecondTPCFieldDCI_0_1_r17 != nil || ie.SecondTPCFieldDCI_0_2_r17 != nil || ie.SequenceOffsetForRV_r17 != nil || ie.Ul_AccessConfigListDCI_0_1_r17 != nil || ie.MinimumSchedulingOffsetK2_r17 != nil || ie.AvailableSlotCounting_r17 != nil || ie.Dmrs_BundlingPUSCH_Config_r17 != nil || ie.Harq_ProcessNumberSizeDCI_0_2_v1700 != nil || ie.Harq_ProcessNumberSizeDCI_0_1_r17 != nil || len(ie.Mpe_ResourcePoolToAddModList_r17) > 0 || len(ie.Mpe_ResourcePoolToReleaseList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.DataScramblingIdentityPUSCH != nil, ie.TxConfig != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeA != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeB != nil, ie.Pusch_PowerControl != nil, ie.FrequencyHopping != nil, len(ie.FrequencyHoppingOffsetLists) > 0, ie.Pusch_TimeDomainAllocationList != nil, ie.Pusch_AggregationFactor != nil, ie.Mcs_Table != nil, ie.Mcs_TableTransformPrecoder != nil, ie.TransformPrecoder != nil, ie.CodebookSubset != nil, ie.MaxRank != nil, ie.Rbg_Size != nil, ie.Uci_OnPUSCH != nil, ie.Tp_pi2BPSK != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DataScramblingIdentityPUSCH != nil {
		if err = w.WriteInteger(*ie.DataScramblingIdentityPUSCH, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode DataScramblingIdentityPUSCH", err)
		}
	}
	if ie.TxConfig != nil {
		if err = ie.TxConfig.Encode(w); err != nil {
			return utils.WrapError("Encode TxConfig", err)
		}
	}
	if ie.Dmrs_UplinkForPUSCH_MappingTypeA != nil {
		tmp_Dmrs_UplinkForPUSCH_MappingTypeA := utils.SetupRelease[*DMRS_UplinkConfig]{
			Setup: ie.Dmrs_UplinkForPUSCH_MappingTypeA,
		}
		if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeA.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_UplinkForPUSCH_MappingTypeA", err)
		}
	}
	if ie.Dmrs_UplinkForPUSCH_MappingTypeB != nil {
		tmp_Dmrs_UplinkForPUSCH_MappingTypeB := utils.SetupRelease[*DMRS_UplinkConfig]{
			Setup: ie.Dmrs_UplinkForPUSCH_MappingTypeB,
		}
		if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeB.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_UplinkForPUSCH_MappingTypeB", err)
		}
	}
	if ie.Pusch_PowerControl != nil {
		if err = ie.Pusch_PowerControl.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_PowerControl", err)
		}
	}
	if ie.FrequencyHopping != nil {
		if err = ie.FrequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyHopping", err)
		}
	}
	if len(ie.FrequencyHoppingOffsetLists) > 0 {
		tmp_FrequencyHoppingOffsetLists := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.FrequencyHoppingOffsetLists {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false)
			tmp_FrequencyHoppingOffsetLists.Value = append(tmp_FrequencyHoppingOffsetLists.Value, &tmp_ie)
		}
		if err = tmp_FrequencyHoppingOffsetLists.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyHoppingOffsetLists", err)
		}
	}
	if err = ie.ResourceAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceAllocation", err)
	}
	if ie.Pusch_TimeDomainAllocationList != nil {
		tmp_Pusch_TimeDomainAllocationList := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList]{
			Setup: ie.Pusch_TimeDomainAllocationList,
		}
		if err = tmp_Pusch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_TimeDomainAllocationList", err)
		}
	}
	if ie.Pusch_AggregationFactor != nil {
		if err = ie.Pusch_AggregationFactor.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_AggregationFactor", err)
		}
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
	if ie.TransformPrecoder != nil {
		if err = ie.TransformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode TransformPrecoder", err)
		}
	}
	if ie.CodebookSubset != nil {
		if err = ie.CodebookSubset.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookSubset", err)
		}
	}
	if ie.MaxRank != nil {
		if err = w.WriteInteger(*ie.MaxRank, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode MaxRank", err)
		}
	}
	if ie.Rbg_Size != nil {
		if err = ie.Rbg_Size.Encode(w); err != nil {
			return utils.WrapError("Encode Rbg_Size", err)
		}
	}
	if ie.Uci_OnPUSCH != nil {
		tmp_Uci_OnPUSCH := utils.SetupRelease[*UCI_OnPUSCH]{
			Setup: ie.Uci_OnPUSCH,
		}
		if err = tmp_Uci_OnPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Uci_OnPUSCH", err)
		}
	}
	if ie.Tp_pi2BPSK != nil {
		if err = ie.Tp_pi2BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode Tp_pi2BPSK", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.MinimumSchedulingOffsetK2_r16 != nil || ie.Ul_AccessConfigListDCI_0_1_r16 != nil || ie.Harq_ProcessNumberSizeDCI_0_2_r16 != nil || ie.Dmrs_SequenceInitializationDCI_0_2_r16 != nil || ie.NumberOfBitsForRV_DCI_0_2_r16 != nil || ie.AntennaPortsFieldPresenceDCI_0_2_r16 != nil || ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil || ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil || ie.FrequencyHoppingDCI_0_2_r16 != nil || ie.FrequencyHoppingOffsetListsDCI_0_2_r16 != nil || ie.CodebookSubsetDCI_0_2_r16 != nil || ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 != nil || ie.MaxRankDCI_0_2_r16 != nil || ie.Mcs_TableDCI_0_2_r16 != nil || ie.Mcs_TableTransformPrecoderDCI_0_2_r16 != nil || ie.PriorityIndicatorDCI_0_2_r16 != nil || ie.Pusch_RepTypeIndicatorDCI_0_2_r16 != nil || ie.ResourceAllocationDCI_0_2_r16 != nil || ie.ResourceAllocationType1GranularityDCI_0_2_r16 != nil || ie.Uci_OnPUSCH_ListDCI_0_2_r16 != nil || ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 != nil || ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 != nil || ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 != nil || ie.PriorityIndicatorDCI_0_1_r16 != nil || ie.Pusch_RepTypeIndicatorDCI_0_1_r16 != nil || ie.FrequencyHoppingDCI_0_1_r16 != nil || ie.Uci_OnPUSCH_ListDCI_0_1_r16 != nil || ie.InvalidSymbolPattern_r16 != nil || ie.Pusch_PowerControl_v1610 != nil || ie.Ul_FullPowerTransmission_r16 != nil || ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil || ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 != nil, ie.Ul_AccessConfigListDCI_0_2_r17 != nil || ie.BetaOffsetsCrossPri0_r17 != nil || ie.BetaOffsetsCrossPri1_r17 != nil || ie.BetaOffsetsCrossPri0DCI_0_2_r17 != nil || ie.BetaOffsetsCrossPri1DCI_0_2_r17 != nil || ie.MappingPattern_r17 != nil || ie.SecondTPCFieldDCI_0_1_r17 != nil || ie.SecondTPCFieldDCI_0_2_r17 != nil || ie.SequenceOffsetForRV_r17 != nil || ie.Ul_AccessConfigListDCI_0_1_r17 != nil || ie.MinimumSchedulingOffsetK2_r17 != nil || ie.AvailableSlotCounting_r17 != nil || ie.Dmrs_BundlingPUSCH_Config_r17 != nil || ie.Harq_ProcessNumberSizeDCI_0_2_v1700 != nil || ie.Harq_ProcessNumberSizeDCI_0_1_r17 != nil || len(ie.Mpe_ResourcePoolToAddModList_r17) > 0 || len(ie.Mpe_ResourcePoolToReleaseList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MinimumSchedulingOffsetK2_r16 != nil, ie.Ul_AccessConfigListDCI_0_1_r16 != nil, ie.Harq_ProcessNumberSizeDCI_0_2_r16 != nil, ie.Dmrs_SequenceInitializationDCI_0_2_r16 != nil, ie.NumberOfBitsForRV_DCI_0_2_r16 != nil, ie.AntennaPortsFieldPresenceDCI_0_2_r16 != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil, ie.FrequencyHoppingDCI_0_2_r16 != nil, ie.FrequencyHoppingOffsetListsDCI_0_2_r16 != nil, ie.CodebookSubsetDCI_0_2_r16 != nil, ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 != nil, ie.MaxRankDCI_0_2_r16 != nil, ie.Mcs_TableDCI_0_2_r16 != nil, ie.Mcs_TableTransformPrecoderDCI_0_2_r16 != nil, ie.PriorityIndicatorDCI_0_2_r16 != nil, ie.Pusch_RepTypeIndicatorDCI_0_2_r16 != nil, ie.ResourceAllocationDCI_0_2_r16 != nil, ie.ResourceAllocationType1GranularityDCI_0_2_r16 != nil, ie.Uci_OnPUSCH_ListDCI_0_2_r16 != nil, ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 != nil, ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 != nil, ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 != nil, ie.PriorityIndicatorDCI_0_1_r16 != nil, ie.Pusch_RepTypeIndicatorDCI_0_1_r16 != nil, ie.FrequencyHoppingDCI_0_1_r16 != nil, ie.Uci_OnPUSCH_ListDCI_0_1_r16 != nil, ie.InvalidSymbolPattern_r16 != nil, ie.Pusch_PowerControl_v1610 != nil, ie.Ul_FullPowerTransmission_r16 != nil, ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil, ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MinimumSchedulingOffsetK2_r16 optional
			if ie.MinimumSchedulingOffsetK2_r16 != nil {
				tmp_MinimumSchedulingOffsetK2_r16 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r16]{
					Setup: ie.MinimumSchedulingOffsetK2_r16,
				}
				if err = tmp_MinimumSchedulingOffsetK2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MinimumSchedulingOffsetK2_r16", err)
				}
			}
			// encode Ul_AccessConfigListDCI_0_1_r16 optional
			if ie.Ul_AccessConfigListDCI_0_1_r16 != nil {
				tmp_Ul_AccessConfigListDCI_0_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r16]{
					Setup: ie.Ul_AccessConfigListDCI_0_1_r16,
				}
				if err = tmp_Ul_AccessConfigListDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_AccessConfigListDCI_0_1_r16", err)
				}
			}
			// encode Harq_ProcessNumberSizeDCI_0_2_r16 optional
			if ie.Harq_ProcessNumberSizeDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcessNumberSizeDCI_0_2_r16, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcessNumberSizeDCI_0_2_r16", err)
				}
			}
			// encode Dmrs_SequenceInitializationDCI_0_2_r16 optional
			if ie.Dmrs_SequenceInitializationDCI_0_2_r16 != nil {
				if err = ie.Dmrs_SequenceInitializationDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_SequenceInitializationDCI_0_2_r16", err)
				}
			}
			// encode NumberOfBitsForRV_DCI_0_2_r16 optional
			if ie.NumberOfBitsForRV_DCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.NumberOfBitsForRV_DCI_0_2_r16, &uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode NumberOfBitsForRV_DCI_0_2_r16", err)
				}
			}
			// encode AntennaPortsFieldPresenceDCI_0_2_r16 optional
			if ie.AntennaPortsFieldPresenceDCI_0_2_r16 != nil {
				if err = ie.AntennaPortsFieldPresenceDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AntennaPortsFieldPresenceDCI_0_2_r16", err)
				}
			}
			// encode Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 optional
			if ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil {
				tmp_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{
					Setup: ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16,
				}
				if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16", err)
				}
			}
			// encode Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 optional
			if ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil {
				tmp_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{
					Setup: ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16,
				}
				if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16", err)
				}
			}
			// encode FrequencyHoppingDCI_0_2_r16 optional
			if ie.FrequencyHoppingDCI_0_2_r16 != nil {
				if err = ie.FrequencyHoppingDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FrequencyHoppingDCI_0_2_r16", err)
				}
			}
			// encode FrequencyHoppingOffsetListsDCI_0_2_r16 optional
			if ie.FrequencyHoppingOffsetListsDCI_0_2_r16 != nil {
				tmp_FrequencyHoppingOffsetListsDCI_0_2_r16 := utils.SetupRelease[*FrequencyHoppingOffsetListsDCI_0_2_r16]{
					Setup: ie.FrequencyHoppingOffsetListsDCI_0_2_r16,
				}
				if err = tmp_FrequencyHoppingOffsetListsDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FrequencyHoppingOffsetListsDCI_0_2_r16", err)
				}
			}
			// encode CodebookSubsetDCI_0_2_r16 optional
			if ie.CodebookSubsetDCI_0_2_r16 != nil {
				if err = ie.CodebookSubsetDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookSubsetDCI_0_2_r16", err)
				}
			}
			// encode InvalidSymbolPatternIndicatorDCI_0_2_r16 optional
			if ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 != nil {
				if err = ie.InvalidSymbolPatternIndicatorDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InvalidSymbolPatternIndicatorDCI_0_2_r16", err)
				}
			}
			// encode MaxRankDCI_0_2_r16 optional
			if ie.MaxRankDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.MaxRankDCI_0_2_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode MaxRankDCI_0_2_r16", err)
				}
			}
			// encode Mcs_TableDCI_0_2_r16 optional
			if ie.Mcs_TableDCI_0_2_r16 != nil {
				if err = ie.Mcs_TableDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mcs_TableDCI_0_2_r16", err)
				}
			}
			// encode Mcs_TableTransformPrecoderDCI_0_2_r16 optional
			if ie.Mcs_TableTransformPrecoderDCI_0_2_r16 != nil {
				if err = ie.Mcs_TableTransformPrecoderDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mcs_TableTransformPrecoderDCI_0_2_r16", err)
				}
			}
			// encode PriorityIndicatorDCI_0_2_r16 optional
			if ie.PriorityIndicatorDCI_0_2_r16 != nil {
				if err = ie.PriorityIndicatorDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PriorityIndicatorDCI_0_2_r16", err)
				}
			}
			// encode Pusch_RepTypeIndicatorDCI_0_2_r16 optional
			if ie.Pusch_RepTypeIndicatorDCI_0_2_r16 != nil {
				if err = ie.Pusch_RepTypeIndicatorDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_RepTypeIndicatorDCI_0_2_r16", err)
				}
			}
			// encode ResourceAllocationDCI_0_2_r16 optional
			if ie.ResourceAllocationDCI_0_2_r16 != nil {
				if err = ie.ResourceAllocationDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceAllocationDCI_0_2_r16", err)
				}
			}
			// encode ResourceAllocationType1GranularityDCI_0_2_r16 optional
			if ie.ResourceAllocationType1GranularityDCI_0_2_r16 != nil {
				if err = ie.ResourceAllocationType1GranularityDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceAllocationType1GranularityDCI_0_2_r16", err)
				}
			}
			// encode Uci_OnPUSCH_ListDCI_0_2_r16 optional
			if ie.Uci_OnPUSCH_ListDCI_0_2_r16 != nil {
				tmp_Uci_OnPUSCH_ListDCI_0_2_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_2_r16]{
					Setup: ie.Uci_OnPUSCH_ListDCI_0_2_r16,
				}
				if err = tmp_Uci_OnPUSCH_ListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uci_OnPUSCH_ListDCI_0_2_r16", err)
				}
			}
			// encode Pusch_TimeDomainAllocationListDCI_0_2_r16 optional
			if ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 != nil {
				tmp_Pusch_TimeDomainAllocationListDCI_0_2_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.Pusch_TimeDomainAllocationListDCI_0_2_r16,
				}
				if err = tmp_Pusch_TimeDomainAllocationListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_TimeDomainAllocationListDCI_0_2_r16", err)
				}
			}
			// encode Pusch_TimeDomainAllocationListDCI_0_1_r16 optional
			if ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 != nil {
				tmp_Pusch_TimeDomainAllocationListDCI_0_1_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.Pusch_TimeDomainAllocationListDCI_0_1_r16,
				}
				if err = tmp_Pusch_TimeDomainAllocationListDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_TimeDomainAllocationListDCI_0_1_r16", err)
				}
			}
			// encode InvalidSymbolPatternIndicatorDCI_0_1_r16 optional
			if ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 != nil {
				if err = ie.InvalidSymbolPatternIndicatorDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InvalidSymbolPatternIndicatorDCI_0_1_r16", err)
				}
			}
			// encode PriorityIndicatorDCI_0_1_r16 optional
			if ie.PriorityIndicatorDCI_0_1_r16 != nil {
				if err = ie.PriorityIndicatorDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PriorityIndicatorDCI_0_1_r16", err)
				}
			}
			// encode Pusch_RepTypeIndicatorDCI_0_1_r16 optional
			if ie.Pusch_RepTypeIndicatorDCI_0_1_r16 != nil {
				if err = ie.Pusch_RepTypeIndicatorDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_RepTypeIndicatorDCI_0_1_r16", err)
				}
			}
			// encode FrequencyHoppingDCI_0_1_r16 optional
			if ie.FrequencyHoppingDCI_0_1_r16 != nil {
				if err = ie.FrequencyHoppingDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FrequencyHoppingDCI_0_1_r16", err)
				}
			}
			// encode Uci_OnPUSCH_ListDCI_0_1_r16 optional
			if ie.Uci_OnPUSCH_ListDCI_0_1_r16 != nil {
				tmp_Uci_OnPUSCH_ListDCI_0_1_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_1_r16]{
					Setup: ie.Uci_OnPUSCH_ListDCI_0_1_r16,
				}
				if err = tmp_Uci_OnPUSCH_ListDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uci_OnPUSCH_ListDCI_0_1_r16", err)
				}
			}
			// encode InvalidSymbolPattern_r16 optional
			if ie.InvalidSymbolPattern_r16 != nil {
				if err = ie.InvalidSymbolPattern_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InvalidSymbolPattern_r16", err)
				}
			}
			// encode Pusch_PowerControl_v1610 optional
			if ie.Pusch_PowerControl_v1610 != nil {
				tmp_Pusch_PowerControl_v1610 := utils.SetupRelease[*PUSCH_PowerControl_v1610]{
					Setup: ie.Pusch_PowerControl_v1610,
				}
				if err = tmp_Pusch_PowerControl_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_PowerControl_v1610", err)
				}
			}
			// encode Ul_FullPowerTransmission_r16 optional
			if ie.Ul_FullPowerTransmission_r16 != nil {
				if err = ie.Ul_FullPowerTransmission_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_FullPowerTransmission_r16", err)
				}
			}
			// encode Pusch_TimeDomainAllocationListForMultiPUSCH_r16 optional
			if ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil {
				tmp_Pusch_TimeDomainAllocationListForMultiPUSCH_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16,
				}
				if err = tmp_Pusch_TimeDomainAllocationListForMultiPUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_TimeDomainAllocationListForMultiPUSCH_r16", err)
				}
			}
			// encode NumberOfInvalidSymbolsForDL_UL_Switching_r16 optional
			if ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 != nil {
				if err = extWriter.WriteInteger(*ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode NumberOfInvalidSymbolsForDL_UL_Switching_r16", err)
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
			optionals_ext_2 := []bool{ie.Ul_AccessConfigListDCI_0_2_r17 != nil, ie.BetaOffsetsCrossPri0_r17 != nil, ie.BetaOffsetsCrossPri1_r17 != nil, ie.BetaOffsetsCrossPri0DCI_0_2_r17 != nil, ie.BetaOffsetsCrossPri1DCI_0_2_r17 != nil, ie.MappingPattern_r17 != nil, ie.SecondTPCFieldDCI_0_1_r17 != nil, ie.SecondTPCFieldDCI_0_2_r17 != nil, ie.SequenceOffsetForRV_r17 != nil, ie.Ul_AccessConfigListDCI_0_1_r17 != nil, ie.MinimumSchedulingOffsetK2_r17 != nil, ie.AvailableSlotCounting_r17 != nil, ie.Dmrs_BundlingPUSCH_Config_r17 != nil, ie.Harq_ProcessNumberSizeDCI_0_2_v1700 != nil, ie.Harq_ProcessNumberSizeDCI_0_1_r17 != nil, len(ie.Mpe_ResourcePoolToAddModList_r17) > 0, len(ie.Mpe_ResourcePoolToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ul_AccessConfigListDCI_0_2_r17 optional
			if ie.Ul_AccessConfigListDCI_0_2_r17 != nil {
				tmp_Ul_AccessConfigListDCI_0_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_2_r17]{
					Setup: ie.Ul_AccessConfigListDCI_0_2_r17,
				}
				if err = tmp_Ul_AccessConfigListDCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_AccessConfigListDCI_0_2_r17", err)
				}
			}
			// encode BetaOffsetsCrossPri0_r17 optional
			if ie.BetaOffsetsCrossPri0_r17 != nil {
				tmp_BetaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{
					Setup: ie.BetaOffsetsCrossPri0_r17,
				}
				if err = tmp_BetaOffsetsCrossPri0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BetaOffsetsCrossPri0_r17", err)
				}
			}
			// encode BetaOffsetsCrossPri1_r17 optional
			if ie.BetaOffsetsCrossPri1_r17 != nil {
				tmp_BetaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{
					Setup: ie.BetaOffsetsCrossPri1_r17,
				}
				if err = tmp_BetaOffsetsCrossPri1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BetaOffsetsCrossPri1_r17", err)
				}
			}
			// encode BetaOffsetsCrossPri0DCI_0_2_r17 optional
			if ie.BetaOffsetsCrossPri0DCI_0_2_r17 != nil {
				tmp_BetaOffsetsCrossPri0DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{
					Setup: ie.BetaOffsetsCrossPri0DCI_0_2_r17,
				}
				if err = tmp_BetaOffsetsCrossPri0DCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BetaOffsetsCrossPri0DCI_0_2_r17", err)
				}
			}
			// encode BetaOffsetsCrossPri1DCI_0_2_r17 optional
			if ie.BetaOffsetsCrossPri1DCI_0_2_r17 != nil {
				tmp_BetaOffsetsCrossPri1DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{
					Setup: ie.BetaOffsetsCrossPri1DCI_0_2_r17,
				}
				if err = tmp_BetaOffsetsCrossPri1DCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BetaOffsetsCrossPri1DCI_0_2_r17", err)
				}
			}
			// encode MappingPattern_r17 optional
			if ie.MappingPattern_r17 != nil {
				if err = ie.MappingPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MappingPattern_r17", err)
				}
			}
			// encode SecondTPCFieldDCI_0_1_r17 optional
			if ie.SecondTPCFieldDCI_0_1_r17 != nil {
				if err = ie.SecondTPCFieldDCI_0_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SecondTPCFieldDCI_0_1_r17", err)
				}
			}
			// encode SecondTPCFieldDCI_0_2_r17 optional
			if ie.SecondTPCFieldDCI_0_2_r17 != nil {
				if err = ie.SecondTPCFieldDCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SecondTPCFieldDCI_0_2_r17", err)
				}
			}
			// encode SequenceOffsetForRV_r17 optional
			if ie.SequenceOffsetForRV_r17 != nil {
				if err = extWriter.WriteInteger(*ie.SequenceOffsetForRV_r17, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode SequenceOffsetForRV_r17", err)
				}
			}
			// encode Ul_AccessConfigListDCI_0_1_r17 optional
			if ie.Ul_AccessConfigListDCI_0_1_r17 != nil {
				tmp_Ul_AccessConfigListDCI_0_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r17]{
					Setup: ie.Ul_AccessConfigListDCI_0_1_r17,
				}
				if err = tmp_Ul_AccessConfigListDCI_0_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_AccessConfigListDCI_0_1_r17", err)
				}
			}
			// encode MinimumSchedulingOffsetK2_r17 optional
			if ie.MinimumSchedulingOffsetK2_r17 != nil {
				tmp_MinimumSchedulingOffsetK2_r17 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r17]{
					Setup: ie.MinimumSchedulingOffsetK2_r17,
				}
				if err = tmp_MinimumSchedulingOffsetK2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MinimumSchedulingOffsetK2_r17", err)
				}
			}
			// encode AvailableSlotCounting_r17 optional
			if ie.AvailableSlotCounting_r17 != nil {
				if err = ie.AvailableSlotCounting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AvailableSlotCounting_r17", err)
				}
			}
			// encode Dmrs_BundlingPUSCH_Config_r17 optional
			if ie.Dmrs_BundlingPUSCH_Config_r17 != nil {
				tmp_Dmrs_BundlingPUSCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUSCH_Config_r17]{
					Setup: ie.Dmrs_BundlingPUSCH_Config_r17,
				}
				if err = tmp_Dmrs_BundlingPUSCH_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingPUSCH_Config_r17", err)
				}
			}
			// encode Harq_ProcessNumberSizeDCI_0_2_v1700 optional
			if ie.Harq_ProcessNumberSizeDCI_0_2_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcessNumberSizeDCI_0_2_v1700, &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcessNumberSizeDCI_0_2_v1700", err)
				}
			}
			// encode Harq_ProcessNumberSizeDCI_0_1_r17 optional
			if ie.Harq_ProcessNumberSizeDCI_0_1_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcessNumberSizeDCI_0_1_r17, &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcessNumberSizeDCI_0_1_r17", err)
				}
			}
			// encode Mpe_ResourcePoolToAddModList_r17 optional
			if len(ie.Mpe_ResourcePoolToAddModList_r17) > 0 {
				tmp_Mpe_ResourcePoolToAddModList_r17 := utils.NewSequence[*MPE_Resource_r17]([]*MPE_Resource_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				for _, i := range ie.Mpe_ResourcePoolToAddModList_r17 {
					tmp_Mpe_ResourcePoolToAddModList_r17.Value = append(tmp_Mpe_ResourcePoolToAddModList_r17.Value, &i)
				}
				if err = tmp_Mpe_ResourcePoolToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mpe_ResourcePoolToAddModList_r17", err)
				}
			}
			// encode Mpe_ResourcePoolToReleaseList_r17 optional
			if len(ie.Mpe_ResourcePoolToReleaseList_r17) > 0 {
				tmp_Mpe_ResourcePoolToReleaseList_r17 := utils.NewSequence[*MPE_ResourceId_r17]([]*MPE_ResourceId_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				for _, i := range ie.Mpe_ResourcePoolToReleaseList_r17 {
					tmp_Mpe_ResourcePoolToReleaseList_r17.Value = append(tmp_Mpe_ResourcePoolToReleaseList_r17.Value, &i)
				}
				if err = tmp_Mpe_ResourcePoolToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mpe_ResourcePoolToReleaseList_r17", err)
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

func (ie *PUSCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var DataScramblingIdentityPUSCHPresent bool
	if DataScramblingIdentityPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TxConfigPresent bool
	if TxConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_UplinkForPUSCH_MappingTypeAPresent bool
	if Dmrs_UplinkForPUSCH_MappingTypeAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_UplinkForPUSCH_MappingTypeBPresent bool
	if Dmrs_UplinkForPUSCH_MappingTypeBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_PowerControlPresent bool
	if Pusch_PowerControlPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyHoppingPresent bool
	if FrequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyHoppingOffsetListsPresent bool
	if FrequencyHoppingOffsetListsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_TimeDomainAllocationListPresent bool
	if Pusch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_AggregationFactorPresent bool
	if Pusch_AggregationFactorPresent, err = r.ReadBool(); err != nil {
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
	var TransformPrecoderPresent bool
	if TransformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CodebookSubsetPresent bool
	if CodebookSubsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxRankPresent bool
	if MaxRankPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rbg_SizePresent bool
	if Rbg_SizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Uci_OnPUSCHPresent bool
	if Uci_OnPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tp_pi2BPSKPresent bool
	if Tp_pi2BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DataScramblingIdentityPUSCHPresent {
		var tmp_int_DataScramblingIdentityPUSCH int64
		if tmp_int_DataScramblingIdentityPUSCH, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode DataScramblingIdentityPUSCH", err)
		}
		ie.DataScramblingIdentityPUSCH = &tmp_int_DataScramblingIdentityPUSCH
	}
	if TxConfigPresent {
		ie.TxConfig = new(PUSCH_Config_txConfig)
		if err = ie.TxConfig.Decode(r); err != nil {
			return utils.WrapError("Decode TxConfig", err)
		}
	}
	if Dmrs_UplinkForPUSCH_MappingTypeAPresent {
		tmp_Dmrs_UplinkForPUSCH_MappingTypeA := utils.SetupRelease[*DMRS_UplinkConfig]{}
		if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeA.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_UplinkForPUSCH_MappingTypeA", err)
		}
		ie.Dmrs_UplinkForPUSCH_MappingTypeA = tmp_Dmrs_UplinkForPUSCH_MappingTypeA.Setup
	}
	if Dmrs_UplinkForPUSCH_MappingTypeBPresent {
		tmp_Dmrs_UplinkForPUSCH_MappingTypeB := utils.SetupRelease[*DMRS_UplinkConfig]{}
		if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeB.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_UplinkForPUSCH_MappingTypeB", err)
		}
		ie.Dmrs_UplinkForPUSCH_MappingTypeB = tmp_Dmrs_UplinkForPUSCH_MappingTypeB.Setup
	}
	if Pusch_PowerControlPresent {
		ie.Pusch_PowerControl = new(PUSCH_PowerControl)
		if err = ie.Pusch_PowerControl.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_PowerControl", err)
		}
	}
	if FrequencyHoppingPresent {
		ie.FrequencyHopping = new(PUSCH_Config_frequencyHopping)
		if err = ie.FrequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyHopping", err)
		}
	}
	if FrequencyHoppingOffsetListsPresent {
		tmp_FrequencyHoppingOffsetLists := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		fn_FrequencyHoppingOffsetLists := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false)
			return &ie
		}
		if err = tmp_FrequencyHoppingOffsetLists.Decode(r, fn_FrequencyHoppingOffsetLists); err != nil {
			return utils.WrapError("Decode FrequencyHoppingOffsetLists", err)
		}
		ie.FrequencyHoppingOffsetLists = []int64{}
		for _, i := range tmp_FrequencyHoppingOffsetLists.Value {
			ie.FrequencyHoppingOffsetLists = append(ie.FrequencyHoppingOffsetLists, int64(i.Value))
		}
	}
	if err = ie.ResourceAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceAllocation", err)
	}
	if Pusch_TimeDomainAllocationListPresent {
		tmp_Pusch_TimeDomainAllocationList := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList]{}
		if err = tmp_Pusch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_TimeDomainAllocationList", err)
		}
		ie.Pusch_TimeDomainAllocationList = tmp_Pusch_TimeDomainAllocationList.Setup
	}
	if Pusch_AggregationFactorPresent {
		ie.Pusch_AggregationFactor = new(PUSCH_Config_pusch_AggregationFactor)
		if err = ie.Pusch_AggregationFactor.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_AggregationFactor", err)
		}
	}
	if Mcs_TablePresent {
		ie.Mcs_Table = new(PUSCH_Config_mcs_Table)
		if err = ie.Mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode Mcs_Table", err)
		}
	}
	if Mcs_TableTransformPrecoderPresent {
		ie.Mcs_TableTransformPrecoder = new(PUSCH_Config_mcs_TableTransformPrecoder)
		if err = ie.Mcs_TableTransformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode Mcs_TableTransformPrecoder", err)
		}
	}
	if TransformPrecoderPresent {
		ie.TransformPrecoder = new(PUSCH_Config_transformPrecoder)
		if err = ie.TransformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode TransformPrecoder", err)
		}
	}
	if CodebookSubsetPresent {
		ie.CodebookSubset = new(PUSCH_Config_codebookSubset)
		if err = ie.CodebookSubset.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookSubset", err)
		}
	}
	if MaxRankPresent {
		var tmp_int_MaxRank int64
		if tmp_int_MaxRank, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode MaxRank", err)
		}
		ie.MaxRank = &tmp_int_MaxRank
	}
	if Rbg_SizePresent {
		ie.Rbg_Size = new(PUSCH_Config_rbg_Size)
		if err = ie.Rbg_Size.Decode(r); err != nil {
			return utils.WrapError("Decode Rbg_Size", err)
		}
	}
	if Uci_OnPUSCHPresent {
		tmp_Uci_OnPUSCH := utils.SetupRelease[*UCI_OnPUSCH]{}
		if err = tmp_Uci_OnPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Uci_OnPUSCH", err)
		}
		ie.Uci_OnPUSCH = tmp_Uci_OnPUSCH.Setup
	}
	if Tp_pi2BPSKPresent {
		ie.Tp_pi2BPSK = new(PUSCH_Config_tp_pi2BPSK)
		if err = ie.Tp_pi2BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode Tp_pi2BPSK", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			MinimumSchedulingOffsetK2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_AccessConfigListDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcessNumberSizeDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_SequenceInitializationDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NumberOfBitsForRV_DCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AntennaPortsFieldPresenceDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FrequencyHoppingDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FrequencyHoppingOffsetListsDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookSubsetDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InvalidSymbolPatternIndicatorDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxRankDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mcs_TableDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mcs_TableTransformPrecoderDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PriorityIndicatorDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_RepTypeIndicatorDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ResourceAllocationDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ResourceAllocationType1GranularityDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uci_OnPUSCH_ListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_TimeDomainAllocationListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_TimeDomainAllocationListDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InvalidSymbolPatternIndicatorDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PriorityIndicatorDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_RepTypeIndicatorDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FrequencyHoppingDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uci_OnPUSCH_ListDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InvalidSymbolPattern_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_PowerControl_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_FullPowerTransmission_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_TimeDomainAllocationListForMultiPUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NumberOfInvalidSymbolsForDL_UL_Switching_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MinimumSchedulingOffsetK2_r16 optional
			if MinimumSchedulingOffsetK2_r16Present {
				tmp_MinimumSchedulingOffsetK2_r16 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r16]{}
				if err = tmp_MinimumSchedulingOffsetK2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MinimumSchedulingOffsetK2_r16", err)
				}
				ie.MinimumSchedulingOffsetK2_r16 = tmp_MinimumSchedulingOffsetK2_r16.Setup
			}
			// decode Ul_AccessConfigListDCI_0_1_r16 optional
			if Ul_AccessConfigListDCI_0_1_r16Present {
				tmp_Ul_AccessConfigListDCI_0_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r16]{}
				if err = tmp_Ul_AccessConfigListDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_AccessConfigListDCI_0_1_r16", err)
				}
				ie.Ul_AccessConfigListDCI_0_1_r16 = tmp_Ul_AccessConfigListDCI_0_1_r16.Setup
			}
			// decode Harq_ProcessNumberSizeDCI_0_2_r16 optional
			if Harq_ProcessNumberSizeDCI_0_2_r16Present {
				var tmp_int_Harq_ProcessNumberSizeDCI_0_2_r16 int64
				if tmp_int_Harq_ProcessNumberSizeDCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcessNumberSizeDCI_0_2_r16", err)
				}
				ie.Harq_ProcessNumberSizeDCI_0_2_r16 = &tmp_int_Harq_ProcessNumberSizeDCI_0_2_r16
			}
			// decode Dmrs_SequenceInitializationDCI_0_2_r16 optional
			if Dmrs_SequenceInitializationDCI_0_2_r16Present {
				ie.Dmrs_SequenceInitializationDCI_0_2_r16 = new(PUSCH_Config_dmrs_SequenceInitializationDCI_0_2_r16)
				if err = ie.Dmrs_SequenceInitializationDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_SequenceInitializationDCI_0_2_r16", err)
				}
			}
			// decode NumberOfBitsForRV_DCI_0_2_r16 optional
			if NumberOfBitsForRV_DCI_0_2_r16Present {
				var tmp_int_NumberOfBitsForRV_DCI_0_2_r16 int64
				if tmp_int_NumberOfBitsForRV_DCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode NumberOfBitsForRV_DCI_0_2_r16", err)
				}
				ie.NumberOfBitsForRV_DCI_0_2_r16 = &tmp_int_NumberOfBitsForRV_DCI_0_2_r16
			}
			// decode AntennaPortsFieldPresenceDCI_0_2_r16 optional
			if AntennaPortsFieldPresenceDCI_0_2_r16Present {
				ie.AntennaPortsFieldPresenceDCI_0_2_r16 = new(PUSCH_Config_antennaPortsFieldPresenceDCI_0_2_r16)
				if err = ie.AntennaPortsFieldPresenceDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AntennaPortsFieldPresenceDCI_0_2_r16", err)
				}
			}
			// decode Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 optional
			if Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16Present {
				tmp_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{}
				if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16", err)
				}
				ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 = tmp_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16.Setup
			}
			// decode Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 optional
			if Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16Present {
				tmp_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{}
				if err = tmp_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16", err)
				}
				ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 = tmp_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16.Setup
			}
			// decode FrequencyHoppingDCI_0_2_r16 optional
			if FrequencyHoppingDCI_0_2_r16Present {
				ie.FrequencyHoppingDCI_0_2_r16 = new(PUSCH_Config_frequencyHoppingDCI_0_2_r16)
				if err = ie.FrequencyHoppingDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode FrequencyHoppingDCI_0_2_r16", err)
				}
			}
			// decode FrequencyHoppingOffsetListsDCI_0_2_r16 optional
			if FrequencyHoppingOffsetListsDCI_0_2_r16Present {
				tmp_FrequencyHoppingOffsetListsDCI_0_2_r16 := utils.SetupRelease[*FrequencyHoppingOffsetListsDCI_0_2_r16]{}
				if err = tmp_FrequencyHoppingOffsetListsDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode FrequencyHoppingOffsetListsDCI_0_2_r16", err)
				}
				ie.FrequencyHoppingOffsetListsDCI_0_2_r16 = tmp_FrequencyHoppingOffsetListsDCI_0_2_r16.Setup
			}
			// decode CodebookSubsetDCI_0_2_r16 optional
			if CodebookSubsetDCI_0_2_r16Present {
				ie.CodebookSubsetDCI_0_2_r16 = new(PUSCH_Config_codebookSubsetDCI_0_2_r16)
				if err = ie.CodebookSubsetDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookSubsetDCI_0_2_r16", err)
				}
			}
			// decode InvalidSymbolPatternIndicatorDCI_0_2_r16 optional
			if InvalidSymbolPatternIndicatorDCI_0_2_r16Present {
				ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 = new(PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_2_r16)
				if err = ie.InvalidSymbolPatternIndicatorDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode InvalidSymbolPatternIndicatorDCI_0_2_r16", err)
				}
			}
			// decode MaxRankDCI_0_2_r16 optional
			if MaxRankDCI_0_2_r16Present {
				var tmp_int_MaxRankDCI_0_2_r16 int64
				if tmp_int_MaxRankDCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode MaxRankDCI_0_2_r16", err)
				}
				ie.MaxRankDCI_0_2_r16 = &tmp_int_MaxRankDCI_0_2_r16
			}
			// decode Mcs_TableDCI_0_2_r16 optional
			if Mcs_TableDCI_0_2_r16Present {
				ie.Mcs_TableDCI_0_2_r16 = new(PUSCH_Config_mcs_TableDCI_0_2_r16)
				if err = ie.Mcs_TableDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mcs_TableDCI_0_2_r16", err)
				}
			}
			// decode Mcs_TableTransformPrecoderDCI_0_2_r16 optional
			if Mcs_TableTransformPrecoderDCI_0_2_r16Present {
				ie.Mcs_TableTransformPrecoderDCI_0_2_r16 = new(PUSCH_Config_mcs_TableTransformPrecoderDCI_0_2_r16)
				if err = ie.Mcs_TableTransformPrecoderDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mcs_TableTransformPrecoderDCI_0_2_r16", err)
				}
			}
			// decode PriorityIndicatorDCI_0_2_r16 optional
			if PriorityIndicatorDCI_0_2_r16Present {
				ie.PriorityIndicatorDCI_0_2_r16 = new(PUSCH_Config_priorityIndicatorDCI_0_2_r16)
				if err = ie.PriorityIndicatorDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PriorityIndicatorDCI_0_2_r16", err)
				}
			}
			// decode Pusch_RepTypeIndicatorDCI_0_2_r16 optional
			if Pusch_RepTypeIndicatorDCI_0_2_r16Present {
				ie.Pusch_RepTypeIndicatorDCI_0_2_r16 = new(PUSCH_Config_pusch_RepTypeIndicatorDCI_0_2_r16)
				if err = ie.Pusch_RepTypeIndicatorDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_RepTypeIndicatorDCI_0_2_r16", err)
				}
			}
			// decode ResourceAllocationDCI_0_2_r16 optional
			if ResourceAllocationDCI_0_2_r16Present {
				ie.ResourceAllocationDCI_0_2_r16 = new(PUSCH_Config_resourceAllocationDCI_0_2_r16)
				if err = ie.ResourceAllocationDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ResourceAllocationDCI_0_2_r16", err)
				}
			}
			// decode ResourceAllocationType1GranularityDCI_0_2_r16 optional
			if ResourceAllocationType1GranularityDCI_0_2_r16Present {
				ie.ResourceAllocationType1GranularityDCI_0_2_r16 = new(PUSCH_Config_resourceAllocationType1GranularityDCI_0_2_r16)
				if err = ie.ResourceAllocationType1GranularityDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ResourceAllocationType1GranularityDCI_0_2_r16", err)
				}
			}
			// decode Uci_OnPUSCH_ListDCI_0_2_r16 optional
			if Uci_OnPUSCH_ListDCI_0_2_r16Present {
				tmp_Uci_OnPUSCH_ListDCI_0_2_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_2_r16]{}
				if err = tmp_Uci_OnPUSCH_ListDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Uci_OnPUSCH_ListDCI_0_2_r16", err)
				}
				ie.Uci_OnPUSCH_ListDCI_0_2_r16 = tmp_Uci_OnPUSCH_ListDCI_0_2_r16.Setup
			}
			// decode Pusch_TimeDomainAllocationListDCI_0_2_r16 optional
			if Pusch_TimeDomainAllocationListDCI_0_2_r16Present {
				tmp_Pusch_TimeDomainAllocationListDCI_0_2_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_Pusch_TimeDomainAllocationListDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_TimeDomainAllocationListDCI_0_2_r16", err)
				}
				ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 = tmp_Pusch_TimeDomainAllocationListDCI_0_2_r16.Setup
			}
			// decode Pusch_TimeDomainAllocationListDCI_0_1_r16 optional
			if Pusch_TimeDomainAllocationListDCI_0_1_r16Present {
				tmp_Pusch_TimeDomainAllocationListDCI_0_1_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_Pusch_TimeDomainAllocationListDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_TimeDomainAllocationListDCI_0_1_r16", err)
				}
				ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 = tmp_Pusch_TimeDomainAllocationListDCI_0_1_r16.Setup
			}
			// decode InvalidSymbolPatternIndicatorDCI_0_1_r16 optional
			if InvalidSymbolPatternIndicatorDCI_0_1_r16Present {
				ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 = new(PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_1_r16)
				if err = ie.InvalidSymbolPatternIndicatorDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode InvalidSymbolPatternIndicatorDCI_0_1_r16", err)
				}
			}
			// decode PriorityIndicatorDCI_0_1_r16 optional
			if PriorityIndicatorDCI_0_1_r16Present {
				ie.PriorityIndicatorDCI_0_1_r16 = new(PUSCH_Config_priorityIndicatorDCI_0_1_r16)
				if err = ie.PriorityIndicatorDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PriorityIndicatorDCI_0_1_r16", err)
				}
			}
			// decode Pusch_RepTypeIndicatorDCI_0_1_r16 optional
			if Pusch_RepTypeIndicatorDCI_0_1_r16Present {
				ie.Pusch_RepTypeIndicatorDCI_0_1_r16 = new(PUSCH_Config_pusch_RepTypeIndicatorDCI_0_1_r16)
				if err = ie.Pusch_RepTypeIndicatorDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_RepTypeIndicatorDCI_0_1_r16", err)
				}
			}
			// decode FrequencyHoppingDCI_0_1_r16 optional
			if FrequencyHoppingDCI_0_1_r16Present {
				ie.FrequencyHoppingDCI_0_1_r16 = new(PUSCH_Config_frequencyHoppingDCI_0_1_r16)
				if err = ie.FrequencyHoppingDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode FrequencyHoppingDCI_0_1_r16", err)
				}
			}
			// decode Uci_OnPUSCH_ListDCI_0_1_r16 optional
			if Uci_OnPUSCH_ListDCI_0_1_r16Present {
				tmp_Uci_OnPUSCH_ListDCI_0_1_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_1_r16]{}
				if err = tmp_Uci_OnPUSCH_ListDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Uci_OnPUSCH_ListDCI_0_1_r16", err)
				}
				ie.Uci_OnPUSCH_ListDCI_0_1_r16 = tmp_Uci_OnPUSCH_ListDCI_0_1_r16.Setup
			}
			// decode InvalidSymbolPattern_r16 optional
			if InvalidSymbolPattern_r16Present {
				ie.InvalidSymbolPattern_r16 = new(InvalidSymbolPattern_r16)
				if err = ie.InvalidSymbolPattern_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode InvalidSymbolPattern_r16", err)
				}
			}
			// decode Pusch_PowerControl_v1610 optional
			if Pusch_PowerControl_v1610Present {
				tmp_Pusch_PowerControl_v1610 := utils.SetupRelease[*PUSCH_PowerControl_v1610]{}
				if err = tmp_Pusch_PowerControl_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_PowerControl_v1610", err)
				}
				ie.Pusch_PowerControl_v1610 = tmp_Pusch_PowerControl_v1610.Setup
			}
			// decode Ul_FullPowerTransmission_r16 optional
			if Ul_FullPowerTransmission_r16Present {
				ie.Ul_FullPowerTransmission_r16 = new(PUSCH_Config_ul_FullPowerTransmission_r16)
				if err = ie.Ul_FullPowerTransmission_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_FullPowerTransmission_r16", err)
				}
			}
			// decode Pusch_TimeDomainAllocationListForMultiPUSCH_r16 optional
			if Pusch_TimeDomainAllocationListForMultiPUSCH_r16Present {
				tmp_Pusch_TimeDomainAllocationListForMultiPUSCH_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_Pusch_TimeDomainAllocationListForMultiPUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_TimeDomainAllocationListForMultiPUSCH_r16", err)
				}
				ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 = tmp_Pusch_TimeDomainAllocationListForMultiPUSCH_r16.Setup
			}
			// decode NumberOfInvalidSymbolsForDL_UL_Switching_r16 optional
			if NumberOfInvalidSymbolsForDL_UL_Switching_r16Present {
				var tmp_int_NumberOfInvalidSymbolsForDL_UL_Switching_r16 int64
				if tmp_int_NumberOfInvalidSymbolsForDL_UL_Switching_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode NumberOfInvalidSymbolsForDL_UL_Switching_r16", err)
				}
				ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 = &tmp_int_NumberOfInvalidSymbolsForDL_UL_Switching_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Ul_AccessConfigListDCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BetaOffsetsCrossPri0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BetaOffsetsCrossPri1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BetaOffsetsCrossPri0DCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BetaOffsetsCrossPri1DCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MappingPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SecondTPCFieldDCI_0_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SecondTPCFieldDCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SequenceOffsetForRV_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_AccessConfigListDCI_0_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MinimumSchedulingOffsetK2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AvailableSlotCounting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_BundlingPUSCH_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcessNumberSizeDCI_0_2_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcessNumberSizeDCI_0_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mpe_ResourcePoolToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mpe_ResourcePoolToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ul_AccessConfigListDCI_0_2_r17 optional
			if Ul_AccessConfigListDCI_0_2_r17Present {
				tmp_Ul_AccessConfigListDCI_0_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_2_r17]{}
				if err = tmp_Ul_AccessConfigListDCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_AccessConfigListDCI_0_2_r17", err)
				}
				ie.Ul_AccessConfigListDCI_0_2_r17 = tmp_Ul_AccessConfigListDCI_0_2_r17.Setup
			}
			// decode BetaOffsetsCrossPri0_r17 optional
			if BetaOffsetsCrossPri0_r17Present {
				tmp_BetaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{}
				if err = tmp_BetaOffsetsCrossPri0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode BetaOffsetsCrossPri0_r17", err)
				}
				ie.BetaOffsetsCrossPri0_r17 = tmp_BetaOffsetsCrossPri0_r17.Setup
			}
			// decode BetaOffsetsCrossPri1_r17 optional
			if BetaOffsetsCrossPri1_r17Present {
				tmp_BetaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{}
				if err = tmp_BetaOffsetsCrossPri1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode BetaOffsetsCrossPri1_r17", err)
				}
				ie.BetaOffsetsCrossPri1_r17 = tmp_BetaOffsetsCrossPri1_r17.Setup
			}
			// decode BetaOffsetsCrossPri0DCI_0_2_r17 optional
			if BetaOffsetsCrossPri0DCI_0_2_r17Present {
				tmp_BetaOffsetsCrossPri0DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{}
				if err = tmp_BetaOffsetsCrossPri0DCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode BetaOffsetsCrossPri0DCI_0_2_r17", err)
				}
				ie.BetaOffsetsCrossPri0DCI_0_2_r17 = tmp_BetaOffsetsCrossPri0DCI_0_2_r17.Setup
			}
			// decode BetaOffsetsCrossPri1DCI_0_2_r17 optional
			if BetaOffsetsCrossPri1DCI_0_2_r17Present {
				tmp_BetaOffsetsCrossPri1DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{}
				if err = tmp_BetaOffsetsCrossPri1DCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode BetaOffsetsCrossPri1DCI_0_2_r17", err)
				}
				ie.BetaOffsetsCrossPri1DCI_0_2_r17 = tmp_BetaOffsetsCrossPri1DCI_0_2_r17.Setup
			}
			// decode MappingPattern_r17 optional
			if MappingPattern_r17Present {
				ie.MappingPattern_r17 = new(PUSCH_Config_mappingPattern_r17)
				if err = ie.MappingPattern_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MappingPattern_r17", err)
				}
			}
			// decode SecondTPCFieldDCI_0_1_r17 optional
			if SecondTPCFieldDCI_0_1_r17Present {
				ie.SecondTPCFieldDCI_0_1_r17 = new(PUSCH_Config_secondTPCFieldDCI_0_1_r17)
				if err = ie.SecondTPCFieldDCI_0_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SecondTPCFieldDCI_0_1_r17", err)
				}
			}
			// decode SecondTPCFieldDCI_0_2_r17 optional
			if SecondTPCFieldDCI_0_2_r17Present {
				ie.SecondTPCFieldDCI_0_2_r17 = new(PUSCH_Config_secondTPCFieldDCI_0_2_r17)
				if err = ie.SecondTPCFieldDCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SecondTPCFieldDCI_0_2_r17", err)
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
			// decode Ul_AccessConfigListDCI_0_1_r17 optional
			if Ul_AccessConfigListDCI_0_1_r17Present {
				tmp_Ul_AccessConfigListDCI_0_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r17]{}
				if err = tmp_Ul_AccessConfigListDCI_0_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_AccessConfigListDCI_0_1_r17", err)
				}
				ie.Ul_AccessConfigListDCI_0_1_r17 = tmp_Ul_AccessConfigListDCI_0_1_r17.Setup
			}
			// decode MinimumSchedulingOffsetK2_r17 optional
			if MinimumSchedulingOffsetK2_r17Present {
				tmp_MinimumSchedulingOffsetK2_r17 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r17]{}
				if err = tmp_MinimumSchedulingOffsetK2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MinimumSchedulingOffsetK2_r17", err)
				}
				ie.MinimumSchedulingOffsetK2_r17 = tmp_MinimumSchedulingOffsetK2_r17.Setup
			}
			// decode AvailableSlotCounting_r17 optional
			if AvailableSlotCounting_r17Present {
				ie.AvailableSlotCounting_r17 = new(PUSCH_Config_availableSlotCounting_r17)
				if err = ie.AvailableSlotCounting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AvailableSlotCounting_r17", err)
				}
			}
			// decode Dmrs_BundlingPUSCH_Config_r17 optional
			if Dmrs_BundlingPUSCH_Config_r17Present {
				tmp_Dmrs_BundlingPUSCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUSCH_Config_r17]{}
				if err = tmp_Dmrs_BundlingPUSCH_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingPUSCH_Config_r17", err)
				}
				ie.Dmrs_BundlingPUSCH_Config_r17 = tmp_Dmrs_BundlingPUSCH_Config_r17.Setup
			}
			// decode Harq_ProcessNumberSizeDCI_0_2_v1700 optional
			if Harq_ProcessNumberSizeDCI_0_2_v1700Present {
				var tmp_int_Harq_ProcessNumberSizeDCI_0_2_v1700 int64
				if tmp_int_Harq_ProcessNumberSizeDCI_0_2_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcessNumberSizeDCI_0_2_v1700", err)
				}
				ie.Harq_ProcessNumberSizeDCI_0_2_v1700 = &tmp_int_Harq_ProcessNumberSizeDCI_0_2_v1700
			}
			// decode Harq_ProcessNumberSizeDCI_0_1_r17 optional
			if Harq_ProcessNumberSizeDCI_0_1_r17Present {
				var tmp_int_Harq_ProcessNumberSizeDCI_0_1_r17 int64
				if tmp_int_Harq_ProcessNumberSizeDCI_0_1_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcessNumberSizeDCI_0_1_r17", err)
				}
				ie.Harq_ProcessNumberSizeDCI_0_1_r17 = &tmp_int_Harq_ProcessNumberSizeDCI_0_1_r17
			}
			// decode Mpe_ResourcePoolToAddModList_r17 optional
			if Mpe_ResourcePoolToAddModList_r17Present {
				tmp_Mpe_ResourcePoolToAddModList_r17 := utils.NewSequence[*MPE_Resource_r17]([]*MPE_Resource_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				fn_Mpe_ResourcePoolToAddModList_r17 := func() *MPE_Resource_r17 {
					return new(MPE_Resource_r17)
				}
				if err = tmp_Mpe_ResourcePoolToAddModList_r17.Decode(extReader, fn_Mpe_ResourcePoolToAddModList_r17); err != nil {
					return utils.WrapError("Decode Mpe_ResourcePoolToAddModList_r17", err)
				}
				ie.Mpe_ResourcePoolToAddModList_r17 = []MPE_Resource_r17{}
				for _, i := range tmp_Mpe_ResourcePoolToAddModList_r17.Value {
					ie.Mpe_ResourcePoolToAddModList_r17 = append(ie.Mpe_ResourcePoolToAddModList_r17, *i)
				}
			}
			// decode Mpe_ResourcePoolToReleaseList_r17 optional
			if Mpe_ResourcePoolToReleaseList_r17Present {
				tmp_Mpe_ResourcePoolToReleaseList_r17 := utils.NewSequence[*MPE_ResourceId_r17]([]*MPE_ResourceId_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				fn_Mpe_ResourcePoolToReleaseList_r17 := func() *MPE_ResourceId_r17 {
					return new(MPE_ResourceId_r17)
				}
				if err = tmp_Mpe_ResourcePoolToReleaseList_r17.Decode(extReader, fn_Mpe_ResourcePoolToReleaseList_r17); err != nil {
					return utils.WrapError("Decode Mpe_ResourcePoolToReleaseList_r17", err)
				}
				ie.Mpe_ResourcePoolToReleaseList_r17 = []MPE_ResourceId_r17{}
				for _, i := range tmp_Mpe_ResourcePoolToReleaseList_r17.Value {
					ie.Mpe_ResourcePoolToReleaseList_r17 = append(ie.Mpe_ResourcePoolToReleaseList_r17, *i)
				}
			}
		}
	}
	return nil
}
