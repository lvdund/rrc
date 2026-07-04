// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConfiguredGrantConfig (line 6532).

var configuredGrantConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyHopping", Optional: true},
		{Name: "cg-DMRS-Configuration"},
		{Name: "mcs-Table", Optional: true},
		{Name: "mcs-TableTransformPrecoder", Optional: true},
		{Name: "uci-OnPUSCH", Optional: true},
		{Name: "resourceAllocation"},
		{Name: "rbg-Size", Optional: true},
		{Name: "powerControlLoopToUse"},
		{Name: "p0-PUSCH-Alpha"},
		{Name: "transformPrecoder", Optional: true},
		{Name: "nrofHARQ-Processes"},
		{Name: "repK"},
		{Name: "repK-RV", Optional: true},
		{Name: "periodicity"},
		{Name: "configuredGrantTimer", Optional: true},
		{Name: "rrc-ConfiguredUplinkGrant", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	ConfiguredGrantConfig_FrequencyHopping_IntraSlot = 0
	ConfiguredGrantConfig_FrequencyHopping_InterSlot = 1
)

var configuredGrantConfigFrequencyHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_FrequencyHopping_IntraSlot, ConfiguredGrantConfig_FrequencyHopping_InterSlot},
}

const (
	ConfiguredGrantConfig_Mcs_Table_Qam256     = 0
	ConfiguredGrantConfig_Mcs_Table_Qam64LowSE = 1
)

var configuredGrantConfigMcsTableConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Mcs_Table_Qam256, ConfiguredGrantConfig_Mcs_Table_Qam64LowSE},
}

const (
	ConfiguredGrantConfig_Mcs_TableTransformPrecoder_Qam256     = 0
	ConfiguredGrantConfig_Mcs_TableTransformPrecoder_Qam64LowSE = 1
)

var configuredGrantConfigMcsTableTransformPrecoderConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Mcs_TableTransformPrecoder_Qam256, ConfiguredGrantConfig_Mcs_TableTransformPrecoder_Qam64LowSE},
}

var configuredGrantConfigUciOnPUSCHConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ConfiguredGrantConfig_Uci_OnPUSCH_Release = 0
	ConfiguredGrantConfig_Uci_OnPUSCH_Setup   = 1
)

const (
	ConfiguredGrantConfig_ResourceAllocation_ResourceAllocationType0 = 0
	ConfiguredGrantConfig_ResourceAllocation_ResourceAllocationType1 = 1
	ConfiguredGrantConfig_ResourceAllocation_DynamicSwitch           = 2
)

var configuredGrantConfigResourceAllocationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_ResourceAllocation_ResourceAllocationType0, ConfiguredGrantConfig_ResourceAllocation_ResourceAllocationType1, ConfiguredGrantConfig_ResourceAllocation_DynamicSwitch},
}

const (
	ConfiguredGrantConfig_Rbg_Size_Config2 = 0
)

var configuredGrantConfigRbgSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Rbg_Size_Config2},
}

const (
	ConfiguredGrantConfig_PowerControlLoopToUse_N0 = 0
	ConfiguredGrantConfig_PowerControlLoopToUse_N1 = 1
)

var configuredGrantConfigPowerControlLoopToUseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_PowerControlLoopToUse_N0, ConfiguredGrantConfig_PowerControlLoopToUse_N1},
}

const (
	ConfiguredGrantConfig_TransformPrecoder_Enabled  = 0
	ConfiguredGrantConfig_TransformPrecoder_Disabled = 1
)

var configuredGrantConfigTransformPrecoderConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_TransformPrecoder_Enabled, ConfiguredGrantConfig_TransformPrecoder_Disabled},
}

var configuredGrantConfigNrofHARQProcessesConstraints = per.Constrained(1, 16)

const (
	ConfiguredGrantConfig_RepK_N1 = 0
	ConfiguredGrantConfig_RepK_N2 = 1
	ConfiguredGrantConfig_RepK_N4 = 2
	ConfiguredGrantConfig_RepK_N8 = 3
)

var configuredGrantConfigRepKConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_RepK_N1, ConfiguredGrantConfig_RepK_N2, ConfiguredGrantConfig_RepK_N4, ConfiguredGrantConfig_RepK_N8},
}

const (
	ConfiguredGrantConfig_RepK_RV_S1_0231 = 0
	ConfiguredGrantConfig_RepK_RV_S2_0303 = 1
	ConfiguredGrantConfig_RepK_RV_S3_0000 = 2
)

var configuredGrantConfigRepKRVConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_RepK_RV_S1_0231, ConfiguredGrantConfig_RepK_RV_S2_0303, ConfiguredGrantConfig_RepK_RV_S3_0000},
}

const (
	ConfiguredGrantConfig_Periodicity_Sym2       = 0
	ConfiguredGrantConfig_Periodicity_Sym7       = 1
	ConfiguredGrantConfig_Periodicity_Sym1x14    = 2
	ConfiguredGrantConfig_Periodicity_Sym2x14    = 3
	ConfiguredGrantConfig_Periodicity_Sym4x14    = 4
	ConfiguredGrantConfig_Periodicity_Sym5x14    = 5
	ConfiguredGrantConfig_Periodicity_Sym8x14    = 6
	ConfiguredGrantConfig_Periodicity_Sym10x14   = 7
	ConfiguredGrantConfig_Periodicity_Sym16x14   = 8
	ConfiguredGrantConfig_Periodicity_Sym20x14   = 9
	ConfiguredGrantConfig_Periodicity_Sym32x14   = 10
	ConfiguredGrantConfig_Periodicity_Sym40x14   = 11
	ConfiguredGrantConfig_Periodicity_Sym64x14   = 12
	ConfiguredGrantConfig_Periodicity_Sym80x14   = 13
	ConfiguredGrantConfig_Periodicity_Sym128x14  = 14
	ConfiguredGrantConfig_Periodicity_Sym160x14  = 15
	ConfiguredGrantConfig_Periodicity_Sym256x14  = 16
	ConfiguredGrantConfig_Periodicity_Sym320x14  = 17
	ConfiguredGrantConfig_Periodicity_Sym512x14  = 18
	ConfiguredGrantConfig_Periodicity_Sym640x14  = 19
	ConfiguredGrantConfig_Periodicity_Sym1024x14 = 20
	ConfiguredGrantConfig_Periodicity_Sym1280x14 = 21
	ConfiguredGrantConfig_Periodicity_Sym2560x14 = 22
	ConfiguredGrantConfig_Periodicity_Sym5120x14 = 23
	ConfiguredGrantConfig_Periodicity_Sym6       = 24
	ConfiguredGrantConfig_Periodicity_Sym1x12    = 25
	ConfiguredGrantConfig_Periodicity_Sym2x12    = 26
	ConfiguredGrantConfig_Periodicity_Sym4x12    = 27
	ConfiguredGrantConfig_Periodicity_Sym5x12    = 28
	ConfiguredGrantConfig_Periodicity_Sym8x12    = 29
	ConfiguredGrantConfig_Periodicity_Sym10x12   = 30
	ConfiguredGrantConfig_Periodicity_Sym16x12   = 31
	ConfiguredGrantConfig_Periodicity_Sym20x12   = 32
	ConfiguredGrantConfig_Periodicity_Sym32x12   = 33
	ConfiguredGrantConfig_Periodicity_Sym40x12   = 34
	ConfiguredGrantConfig_Periodicity_Sym64x12   = 35
	ConfiguredGrantConfig_Periodicity_Sym80x12   = 36
	ConfiguredGrantConfig_Periodicity_Sym128x12  = 37
	ConfiguredGrantConfig_Periodicity_Sym160x12  = 38
	ConfiguredGrantConfig_Periodicity_Sym256x12  = 39
	ConfiguredGrantConfig_Periodicity_Sym320x12  = 40
	ConfiguredGrantConfig_Periodicity_Sym512x12  = 41
	ConfiguredGrantConfig_Periodicity_Sym640x12  = 42
	ConfiguredGrantConfig_Periodicity_Sym1280x12 = 43
	ConfiguredGrantConfig_Periodicity_Sym2560x12 = 44
)

var configuredGrantConfigPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Periodicity_Sym2, ConfiguredGrantConfig_Periodicity_Sym7, ConfiguredGrantConfig_Periodicity_Sym1x14, ConfiguredGrantConfig_Periodicity_Sym2x14, ConfiguredGrantConfig_Periodicity_Sym4x14, ConfiguredGrantConfig_Periodicity_Sym5x14, ConfiguredGrantConfig_Periodicity_Sym8x14, ConfiguredGrantConfig_Periodicity_Sym10x14, ConfiguredGrantConfig_Periodicity_Sym16x14, ConfiguredGrantConfig_Periodicity_Sym20x14, ConfiguredGrantConfig_Periodicity_Sym32x14, ConfiguredGrantConfig_Periodicity_Sym40x14, ConfiguredGrantConfig_Periodicity_Sym64x14, ConfiguredGrantConfig_Periodicity_Sym80x14, ConfiguredGrantConfig_Periodicity_Sym128x14, ConfiguredGrantConfig_Periodicity_Sym160x14, ConfiguredGrantConfig_Periodicity_Sym256x14, ConfiguredGrantConfig_Periodicity_Sym320x14, ConfiguredGrantConfig_Periodicity_Sym512x14, ConfiguredGrantConfig_Periodicity_Sym640x14, ConfiguredGrantConfig_Periodicity_Sym1024x14, ConfiguredGrantConfig_Periodicity_Sym1280x14, ConfiguredGrantConfig_Periodicity_Sym2560x14, ConfiguredGrantConfig_Periodicity_Sym5120x14, ConfiguredGrantConfig_Periodicity_Sym6, ConfiguredGrantConfig_Periodicity_Sym1x12, ConfiguredGrantConfig_Periodicity_Sym2x12, ConfiguredGrantConfig_Periodicity_Sym4x12, ConfiguredGrantConfig_Periodicity_Sym5x12, ConfiguredGrantConfig_Periodicity_Sym8x12, ConfiguredGrantConfig_Periodicity_Sym10x12, ConfiguredGrantConfig_Periodicity_Sym16x12, ConfiguredGrantConfig_Periodicity_Sym20x12, ConfiguredGrantConfig_Periodicity_Sym32x12, ConfiguredGrantConfig_Periodicity_Sym40x12, ConfiguredGrantConfig_Periodicity_Sym64x12, ConfiguredGrantConfig_Periodicity_Sym80x12, ConfiguredGrantConfig_Periodicity_Sym128x12, ConfiguredGrantConfig_Periodicity_Sym160x12, ConfiguredGrantConfig_Periodicity_Sym256x12, ConfiguredGrantConfig_Periodicity_Sym320x12, ConfiguredGrantConfig_Periodicity_Sym512x12, ConfiguredGrantConfig_Periodicity_Sym640x12, ConfiguredGrantConfig_Periodicity_Sym1280x12, ConfiguredGrantConfig_Periodicity_Sym2560x12},
}

var configuredGrantConfigConfiguredGrantTimerConstraints = per.Constrained(1, 64)

var configuredGrantConfigCgRetransmissionTimerR16Constraints = per.Constrained(1, 64)

const (
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym7     = 0
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym1x14  = 1
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym2x14  = 2
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym3x14  = 3
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym4x14  = 4
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym5x14  = 5
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym6x14  = 6
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym7x14  = 7
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym8x14  = 8
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym9x14  = 9
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym10x14 = 10
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym11x14 = 11
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym12x14 = 12
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym13x14 = 13
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym14x14 = 14
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym15x14 = 15
	ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym16x14 = 16
)

var configuredGrantConfigExtCgMinDFIDelayR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym7, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym1x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym2x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym3x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym4x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym5x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym6x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym7x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym8x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym9x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym10x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym11x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym12x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym13x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym14x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym15x14, ConfiguredGrantConfig_Ext_Cg_MinDFI_Delay_r16_Sym16x14},
}

var configuredGrantConfigCgNrofPUSCHInSlotR16Constraints = per.Constrained(1, 7)

var configuredGrantConfigCgNrofSlotsR16Constraints = per.Constrained(1, 40)

const (
	ConfiguredGrantConfig_Ext_Cg_UCI_Multiplexing_r16_Enabled = 0
)

var configuredGrantConfigExtCgUCIMultiplexingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_Cg_UCI_Multiplexing_r16_Enabled},
}

var configuredGrantConfigCgCOTSharingOffsetR16Constraints = per.Constrained(1, 39)

var configuredGrantConfigBetaOffsetCGUCIR16Constraints = per.Constrained(0, 31)

var configuredGrantConfigExtCgCOTSharingListR16Constraints = per.SizeRange(1, 1709)

var configuredGrantConfigHarqProcIDOffsetR16Constraints = per.Constrained(0, 15)

var configuredGrantConfigHarqProcIDOffset2R16Constraints = per.Constrained(0, 15)

var configuredGrantConfigPeriodicityExtR16Constraints = per.Constrained(1, 5120)

const (
	ConfiguredGrantConfig_Ext_StartingFromRV0_r16_On  = 0
	ConfiguredGrantConfig_Ext_StartingFromRV0_r16_Off = 1
)

var configuredGrantConfigExtStartingFromRV0R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_StartingFromRV0_r16_On, ConfiguredGrantConfig_Ext_StartingFromRV0_r16_Off},
}

const (
	ConfiguredGrantConfig_Ext_Phy_PriorityIndex_r16_P0 = 0
	ConfiguredGrantConfig_Ext_Phy_PriorityIndex_r16_P1 = 1
)

var configuredGrantConfigExtPhyPriorityIndexR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_Phy_PriorityIndex_r16_P0, ConfiguredGrantConfig_Ext_Phy_PriorityIndex_r16_P1},
}

const (
	ConfiguredGrantConfig_Ext_AutonomousTx_r16_Enabled = 0
)

var configuredGrantConfigExtAutonomousTxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_AutonomousTx_r16_Enabled},
}

var configuredGrantConfigExtCgBetaOffsetsCrossPri0R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri0_r17_Release = 0
	ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri0_r17_Setup   = 1
)

var configuredGrantConfigExtCgBetaOffsetsCrossPri1R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri1_r17_Release = 0
	ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri1_r17_Setup   = 1
)

const (
	ConfiguredGrantConfig_Ext_MappingPattern_r17_CyclicMapping     = 0
	ConfiguredGrantConfig_Ext_MappingPattern_r17_SequentialMapping = 1
)

var configuredGrantConfigExtMappingPatternR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_MappingPattern_r17_CyclicMapping, ConfiguredGrantConfig_Ext_MappingPattern_r17_SequentialMapping},
}

var configuredGrantConfigSequenceOffsetForRVR17Constraints = per.Constrained(0, 3)

const (
	ConfiguredGrantConfig_Ext_PowerControlLoopToUse2_r17_N0 = 0
	ConfiguredGrantConfig_Ext_PowerControlLoopToUse2_r17_N1 = 1
)

var configuredGrantConfigExtPowerControlLoopToUse2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_PowerControlLoopToUse2_r17_N0, ConfiguredGrantConfig_Ext_PowerControlLoopToUse2_r17_N1},
}

var configuredGrantConfigExtCgCOTSharingListR17Constraints = per.SizeRange(1, 50722)

var configuredGrantConfigPeriodicityExtR17Constraints = per.Constrained(1, 40960)

const (
	ConfiguredGrantConfig_Ext_RepK_v1710_N12 = 0
	ConfiguredGrantConfig_Ext_RepK_v1710_N16 = 1
	ConfiguredGrantConfig_Ext_RepK_v1710_N24 = 2
	ConfiguredGrantConfig_Ext_RepK_v1710_N32 = 3
)

var configuredGrantConfigExtRepKV1710Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_RepK_v1710_N12, ConfiguredGrantConfig_Ext_RepK_v1710_N16, ConfiguredGrantConfig_Ext_RepK_v1710_N24, ConfiguredGrantConfig_Ext_RepK_v1710_N32},
}

var configuredGrantConfigNrofHARQProcessesV1700Constraints = per.Constrained(17, 32)

var configuredGrantConfigHarqProcIDOffset2V1700Constraints = per.Constrained(16, 31)

var configuredGrantConfigConfiguredGrantTimerV1700Constraints = per.Constrained(33, 288)

var configuredGrantConfigCgMinDFIDelayV1710Constraints = per.Constrained(238, 3584)

var configuredGrantConfigHarqProcIDOffsetV1730Constraints = per.Constrained(16, 31)

var configuredGrantConfigCgNrofSlotsR17Constraints = per.Constrained(1, 320)

const (
	ConfiguredGrantConfig_Ext_DisableCG_RetransmissionMonitoring_r18_True = 0
)

var configuredGrantConfigExtDisableCGRetransmissionMonitoringR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConfiguredGrantConfig_Ext_DisableCG_RetransmissionMonitoring_r18_True},
}

var configuredGrantConfigNrofSlotsInCGPeriodR18Constraints = per.Constrained(2, 32)

var configuredGrantConfigPrecodingAndNumberOfLayersV1850Constraints = per.Constrained(64, 1023)

var configuredGrantConfigSrsResourceIndicatorV1850Constraints = per.Constrained(16, 255)

var configuredGrantConfigRrcConfiguredUplinkGrantConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "timeDomainOffset"},
		{Name: "timeDomainAllocation"},
		{Name: "frequencyDomainAllocation"},
		{Name: "antennaPort"},
		{Name: "dmrs-SeqInitialization", Optional: true},
		{Name: "precodingAndNumberOfLayers"},
		{Name: "srs-ResourceIndicator", Optional: true},
		{Name: "mcsAndTBS"},
		{Name: "frequencyHoppingOffset", Optional: true},
		{Name: "pathlossReferenceIndex"},
	},
}

var configuredGrantConfigExtUtoUCIConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofBitsInUTO-UCI-r18"},
		{Name: "betaOffsetUTO-UCI-r18"},
	},
}

type ConfiguredGrantConfig struct {
	FrequencyHopping           *int64
	Cg_DMRS_Configuration      DMRS_UplinkConfig
	Mcs_Table                  *int64
	Mcs_TableTransformPrecoder *int64
	Uci_OnPUSCH                *struct {
		Choice  int
		Release *struct{}
		Setup   *CG_UCI_OnPUSCH
	}
	ResourceAllocation        int64
	Rbg_Size                  *int64
	PowerControlLoopToUse     int64
	P0_PUSCH_Alpha            P0_PUSCH_AlphaSetId
	TransformPrecoder         *int64
	NrofHARQ_Processes        int64
	RepK                      int64
	RepK_RV                   *int64
	Periodicity               int64
	ConfiguredGrantTimer      *int64
	Rrc_ConfiguredUplinkGrant *struct {
		TimeDomainOffset           int64
		TimeDomainAllocation       int64
		FrequencyDomainAllocation  per.BitString
		AntennaPort                int64
		Dmrs_SeqInitialization     *int64
		PrecodingAndNumberOfLayers int64
		Srs_ResourceIndicator      *int64
		McsAndTBS                  int64
		FrequencyHoppingOffset     *int64
		PathlossReferenceIndex     int64
	}
	Cg_RetransmissionTimer_r16        *int64
	Cg_MinDFI_Delay_r16               *int64
	Cg_NrofPUSCH_InSlot_r16           *int64
	Cg_NrofSlots_r16                  *int64
	Cg_StartingOffsets_r16            *CG_StartingOffsets_r16
	Cg_UCI_Multiplexing_r16           *int64
	Cg_COT_SharingOffset_r16          *int64
	BetaOffsetCG_UCI_r16              *int64
	Cg_COT_SharingList_r16            []CG_COT_Sharing_r16
	Harq_ProcID_Offset_r16            *int64
	Harq_ProcID_Offset2_r16           *int64
	ConfiguredGrantConfigIndex_r16    *ConfiguredGrantConfigIndex_r16
	ConfiguredGrantConfigIndexMAC_r16 *ConfiguredGrantConfigIndexMAC_r16
	PeriodicityExt_r16                *int64
	StartingFromRV0_r16               *int64
	Phy_PriorityIndex_r16             *int64
	AutonomousTx_r16                  *int64
	Cg_BetaOffsetsCrossPri0_r17       *struct {
		Choice  int
		Release *struct{}
		Setup   *BetaOffsetsCrossPriSelCG_r17
	}
	Cg_BetaOffsetsCrossPri1_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BetaOffsetsCrossPriSelCG_r17
	}
	MappingPattern_r17                     *int64
	SequenceOffsetForRV_r17                *int64
	P0_PUSCH_Alpha2_r17                    *P0_PUSCH_AlphaSetId
	PowerControlLoopToUse2_r17             *int64
	Cg_COT_SharingList_r17                 []CG_COT_Sharing_r17
	PeriodicityExt_r17                     *int64
	RepK_v1710                             *int64
	NrofHARQ_Processes_v1700               *int64
	Harq_ProcID_Offset2_v1700              *int64
	ConfiguredGrantTimer_v1700             *int64
	Cg_MinDFI_Delay_v1710                  *int64
	Harq_ProcID_Offset_v1730               *int64
	Cg_NrofSlots_r17                       *int64
	DisableCG_RetransmissionMonitoring_r18 *int64
	NrofSlotsInCG_Period_r18               *int64
	Uto_UCI_Config_r18                     *struct {
		NrofBitsInUTO_UCI_r18 int64
		BetaOffsetUTO_UCI_r18 int64
	}
	PrecodingAndNumberOfLayers_v1850 *int64
	Srs_ResourceIndicator_v1850      *int64
}

func (ie *ConfiguredGrantConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(configuredGrantConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Cg_RetransmissionTimer_r16 != nil || ie.Cg_MinDFI_Delay_r16 != nil || ie.Cg_NrofPUSCH_InSlot_r16 != nil || ie.Cg_NrofSlots_r16 != nil || ie.Cg_StartingOffsets_r16 != nil || ie.Cg_UCI_Multiplexing_r16 != nil || ie.Cg_COT_SharingOffset_r16 != nil || ie.BetaOffsetCG_UCI_r16 != nil || ie.Cg_COT_SharingList_r16 != nil || ie.Harq_ProcID_Offset_r16 != nil || ie.Harq_ProcID_Offset2_r16 != nil || ie.ConfiguredGrantConfigIndex_r16 != nil || ie.ConfiguredGrantConfigIndexMAC_r16 != nil || ie.PeriodicityExt_r16 != nil || ie.StartingFromRV0_r16 != nil || ie.Phy_PriorityIndex_r16 != nil || ie.AutonomousTx_r16 != nil
	hasExtGroup1 := ie.Cg_BetaOffsetsCrossPri0_r17 != nil || ie.Cg_BetaOffsetsCrossPri1_r17 != nil || ie.MappingPattern_r17 != nil || ie.SequenceOffsetForRV_r17 != nil || ie.P0_PUSCH_Alpha2_r17 != nil || ie.PowerControlLoopToUse2_r17 != nil || ie.Cg_COT_SharingList_r17 != nil || ie.PeriodicityExt_r17 != nil || ie.RepK_v1710 != nil || ie.NrofHARQ_Processes_v1700 != nil || ie.Harq_ProcID_Offset2_v1700 != nil || ie.ConfiguredGrantTimer_v1700 != nil || ie.Cg_MinDFI_Delay_v1710 != nil
	hasExtGroup2 := ie.Harq_ProcID_Offset_v1730 != nil || ie.Cg_NrofSlots_r17 != nil
	hasExtGroup3 := ie.DisableCG_RetransmissionMonitoring_r18 != nil || ie.NrofSlotsInCG_Period_r18 != nil || ie.Uto_UCI_Config_r18 != nil
	hasExtGroup4 := ie.PrecodingAndNumberOfLayers_v1850 != nil || ie.Srs_ResourceIndicator_v1850 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyHopping != nil, ie.Mcs_Table != nil, ie.Mcs_TableTransformPrecoder != nil, ie.Uci_OnPUSCH != nil, ie.Rbg_Size != nil, ie.TransformPrecoder != nil, ie.RepK_RV != nil, ie.ConfiguredGrantTimer != nil, ie.Rrc_ConfiguredUplinkGrant != nil}); err != nil {
		return err
	}

	// 3. frequencyHopping: enumerated
	{
		if ie.FrequencyHopping != nil {
			if err := e.EncodeEnumerated(*ie.FrequencyHopping, configuredGrantConfigFrequencyHoppingConstraints); err != nil {
				return err
			}
		}
	}

	// 4. cg-DMRS-Configuration: ref
	{
		if err := ie.Cg_DMRS_Configuration.Encode(e); err != nil {
			return err
		}
	}

	// 5. mcs-Table: enumerated
	{
		if ie.Mcs_Table != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_Table, configuredGrantConfigMcsTableConstraints); err != nil {
				return err
			}
		}
	}

	// 6. mcs-TableTransformPrecoder: enumerated
	{
		if ie.Mcs_TableTransformPrecoder != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_TableTransformPrecoder, configuredGrantConfigMcsTableTransformPrecoderConstraints); err != nil {
				return err
			}
		}
	}

	// 7. uci-OnPUSCH: choice
	{
		if ie.Uci_OnPUSCH != nil {
			choiceEnc := e.NewChoiceEncoder(configuredGrantConfigUciOnPUSCHConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Uci_OnPUSCH).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Uci_OnPUSCH).Choice {
			case ConfiguredGrantConfig_Uci_OnPUSCH_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ConfiguredGrantConfig_Uci_OnPUSCH_Setup:
				if err := (*ie.Uci_OnPUSCH).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Uci_OnPUSCH).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. resourceAllocation: enumerated
	{
		if err := e.EncodeEnumerated(ie.ResourceAllocation, configuredGrantConfigResourceAllocationConstraints); err != nil {
			return err
		}
	}

	// 9. rbg-Size: enumerated
	{
		if ie.Rbg_Size != nil {
			if err := e.EncodeEnumerated(*ie.Rbg_Size, configuredGrantConfigRbgSizeConstraints); err != nil {
				return err
			}
		}
	}

	// 10. powerControlLoopToUse: enumerated
	{
		if err := e.EncodeEnumerated(ie.PowerControlLoopToUse, configuredGrantConfigPowerControlLoopToUseConstraints); err != nil {
			return err
		}
	}

	// 11. p0-PUSCH-Alpha: ref
	{
		if err := ie.P0_PUSCH_Alpha.Encode(e); err != nil {
			return err
		}
	}

	// 12. transformPrecoder: enumerated
	{
		if ie.TransformPrecoder != nil {
			if err := e.EncodeEnumerated(*ie.TransformPrecoder, configuredGrantConfigTransformPrecoderConstraints); err != nil {
				return err
			}
		}
	}

	// 13. nrofHARQ-Processes: integer
	{
		if err := e.EncodeInteger(ie.NrofHARQ_Processes, configuredGrantConfigNrofHARQProcessesConstraints); err != nil {
			return err
		}
	}

	// 14. repK: enumerated
	{
		if err := e.EncodeEnumerated(ie.RepK, configuredGrantConfigRepKConstraints); err != nil {
			return err
		}
	}

	// 15. repK-RV: enumerated
	{
		if ie.RepK_RV != nil {
			if err := e.EncodeEnumerated(*ie.RepK_RV, configuredGrantConfigRepKRVConstraints); err != nil {
				return err
			}
		}
	}

	// 16. periodicity: enumerated
	{
		if err := e.EncodeEnumerated(ie.Periodicity, configuredGrantConfigPeriodicityConstraints); err != nil {
			return err
		}
	}

	// 17. configuredGrantTimer: integer
	{
		if ie.ConfiguredGrantTimer != nil {
			if err := e.EncodeInteger(*ie.ConfiguredGrantTimer, configuredGrantConfigConfiguredGrantTimerConstraints); err != nil {
				return err
			}
		}
	}

	// 18. rrc-ConfiguredUplinkGrant: sequence
	{
		if ie.Rrc_ConfiguredUplinkGrant != nil {
			c := ie.Rrc_ConfiguredUplinkGrant
			configuredGrantConfigRrcConfiguredUplinkGrantSeq := e.NewSequenceEncoder(configuredGrantConfigRrcConfiguredUplinkGrantConstraints)
			if err := configuredGrantConfigRrcConfiguredUplinkGrantSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := configuredGrantConfigRrcConfiguredUplinkGrantSeq.EncodePreamble([]bool{c.Dmrs_SeqInitialization != nil, c.Srs_ResourceIndicator != nil, c.FrequencyHoppingOffset != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.TimeDomainOffset, per.Constrained(0, 5119)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.TimeDomainAllocation, per.Constrained(0, 15)); err != nil {
				return err
			}
			if err := e.EncodeBitString(c.FrequencyDomainAllocation, per.FixedSize(18)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.AntennaPort, per.Constrained(0, 31)); err != nil {
				return err
			}
			if c.Dmrs_SeqInitialization != nil {
				if err := e.EncodeInteger((*c.Dmrs_SeqInitialization), per.Constrained(0, 1)); err != nil {
					return err
				}
			}
			if err := e.EncodeInteger(c.PrecodingAndNumberOfLayers, per.Constrained(0, 63)); err != nil {
				return err
			}
			if c.Srs_ResourceIndicator != nil {
				if err := e.EncodeInteger((*c.Srs_ResourceIndicator), per.Constrained(0, 15)); err != nil {
					return err
				}
			}
			if err := e.EncodeInteger(c.McsAndTBS, per.Constrained(0, 31)); err != nil {
				return err
			}
			if c.FrequencyHoppingOffset != nil {
				if err := e.EncodeInteger((*c.FrequencyHoppingOffset), per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1)); err != nil {
					return err
				}
			}
			if err := e.EncodeInteger(c.PathlossReferenceIndex, per.Constrained(0, common.MaxNrofPUSCH_PathlossReferenceRSs_1)); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cg-RetransmissionTimer-r16", Optional: true},
					{Name: "cg-minDFI-Delay-r16", Optional: true},
					{Name: "cg-nrofPUSCH-InSlot-r16", Optional: true},
					{Name: "cg-nrofSlots-r16", Optional: true},
					{Name: "cg-StartingOffsets-r16", Optional: true},
					{Name: "cg-UCI-Multiplexing-r16", Optional: true},
					{Name: "cg-COT-SharingOffset-r16", Optional: true},
					{Name: "betaOffsetCG-UCI-r16", Optional: true},
					{Name: "cg-COT-SharingList-r16", Optional: true},
					{Name: "harq-ProcID-Offset-r16", Optional: true},
					{Name: "harq-ProcID-Offset2-r16", Optional: true},
					{Name: "configuredGrantConfigIndex-r16", Optional: true},
					{Name: "configuredGrantConfigIndexMAC-r16", Optional: true},
					{Name: "periodicityExt-r16", Optional: true},
					{Name: "startingFromRV0-r16", Optional: true},
					{Name: "phy-PriorityIndex-r16", Optional: true},
					{Name: "autonomousTx-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cg_RetransmissionTimer_r16 != nil, ie.Cg_MinDFI_Delay_r16 != nil, ie.Cg_NrofPUSCH_InSlot_r16 != nil, ie.Cg_NrofSlots_r16 != nil, ie.Cg_StartingOffsets_r16 != nil, ie.Cg_UCI_Multiplexing_r16 != nil, ie.Cg_COT_SharingOffset_r16 != nil, ie.BetaOffsetCG_UCI_r16 != nil, ie.Cg_COT_SharingList_r16 != nil, ie.Harq_ProcID_Offset_r16 != nil, ie.Harq_ProcID_Offset2_r16 != nil, ie.ConfiguredGrantConfigIndex_r16 != nil, ie.ConfiguredGrantConfigIndexMAC_r16 != nil, ie.PeriodicityExt_r16 != nil, ie.StartingFromRV0_r16 != nil, ie.Phy_PriorityIndex_r16 != nil, ie.AutonomousTx_r16 != nil}); err != nil {
				return err
			}

			if ie.Cg_RetransmissionTimer_r16 != nil {
				if err := ex.EncodeInteger(*ie.Cg_RetransmissionTimer_r16, configuredGrantConfigCgRetransmissionTimerR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_MinDFI_Delay_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_MinDFI_Delay_r16, configuredGrantConfigExtCgMinDFIDelayR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_NrofPUSCH_InSlot_r16 != nil {
				if err := ex.EncodeInteger(*ie.Cg_NrofPUSCH_InSlot_r16, configuredGrantConfigCgNrofPUSCHInSlotR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_NrofSlots_r16 != nil {
				if err := ex.EncodeInteger(*ie.Cg_NrofSlots_r16, configuredGrantConfigCgNrofSlotsR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_StartingOffsets_r16 != nil {
				if err := ie.Cg_StartingOffsets_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Cg_UCI_Multiplexing_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_UCI_Multiplexing_r16, configuredGrantConfigExtCgUCIMultiplexingR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_COT_SharingOffset_r16 != nil {
				if err := ex.EncodeInteger(*ie.Cg_COT_SharingOffset_r16, configuredGrantConfigCgCOTSharingOffsetR16Constraints); err != nil {
					return err
				}
			}

			if ie.BetaOffsetCG_UCI_r16 != nil {
				if err := ex.EncodeInteger(*ie.BetaOffsetCG_UCI_r16, configuredGrantConfigBetaOffsetCGUCIR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_COT_SharingList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(configuredGrantConfigExtCgCOTSharingListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Cg_COT_SharingList_r16))); err != nil {
					return err
				}
				for i := range ie.Cg_COT_SharingList_r16 {
					if err := ie.Cg_COT_SharingList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_ProcID_Offset_r16 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcID_Offset_r16, configuredGrantConfigHarqProcIDOffsetR16Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcID_Offset2_r16 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcID_Offset2_r16, configuredGrantConfigHarqProcIDOffset2R16Constraints); err != nil {
					return err
				}
			}

			if ie.ConfiguredGrantConfigIndex_r16 != nil {
				if err := ie.ConfiguredGrantConfigIndex_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ConfiguredGrantConfigIndexMAC_r16 != nil {
				if err := ie.ConfiguredGrantConfigIndexMAC_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PeriodicityExt_r16 != nil {
				if err := ex.EncodeInteger(*ie.PeriodicityExt_r16, configuredGrantConfigPeriodicityExtR16Constraints); err != nil {
					return err
				}
			}

			if ie.StartingFromRV0_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.StartingFromRV0_r16, configuredGrantConfigExtStartingFromRV0R16Constraints); err != nil {
					return err
				}
			}

			if ie.Phy_PriorityIndex_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Phy_PriorityIndex_r16, configuredGrantConfigExtPhyPriorityIndexR16Constraints); err != nil {
					return err
				}
			}

			if ie.AutonomousTx_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.AutonomousTx_r16, configuredGrantConfigExtAutonomousTxR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cg-betaOffsetsCrossPri0-r17", Optional: true},
					{Name: "cg-betaOffsetsCrossPri1-r17", Optional: true},
					{Name: "mappingPattern-r17", Optional: true},
					{Name: "sequenceOffsetForRV-r17", Optional: true},
					{Name: "p0-PUSCH-Alpha2-r17", Optional: true},
					{Name: "powerControlLoopToUse2-r17", Optional: true},
					{Name: "cg-COT-SharingList-r17", Optional: true},
					{Name: "periodicityExt-r17", Optional: true},
					{Name: "repK-v1710", Optional: true},
					{Name: "nrofHARQ-Processes-v1700", Optional: true},
					{Name: "harq-ProcID-Offset2-v1700", Optional: true},
					{Name: "configuredGrantTimer-v1700", Optional: true},
					{Name: "cg-minDFI-Delay-v1710", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cg_BetaOffsetsCrossPri0_r17 != nil, ie.Cg_BetaOffsetsCrossPri1_r17 != nil, ie.MappingPattern_r17 != nil, ie.SequenceOffsetForRV_r17 != nil, ie.P0_PUSCH_Alpha2_r17 != nil, ie.PowerControlLoopToUse2_r17 != nil, ie.Cg_COT_SharingList_r17 != nil, ie.PeriodicityExt_r17 != nil, ie.RepK_v1710 != nil, ie.NrofHARQ_Processes_v1700 != nil, ie.Harq_ProcID_Offset2_v1700 != nil, ie.ConfiguredGrantTimer_v1700 != nil, ie.Cg_MinDFI_Delay_v1710 != nil}); err != nil {
				return err
			}

			if ie.Cg_BetaOffsetsCrossPri0_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(configuredGrantConfigExtCgBetaOffsetsCrossPri0R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Cg_BetaOffsetsCrossPri0_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Cg_BetaOffsetsCrossPri0_r17).Choice {
				case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri0_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri0_r17_Setup:
					if err := (*ie.Cg_BetaOffsetsCrossPri0_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Cg_BetaOffsetsCrossPri1_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(configuredGrantConfigExtCgBetaOffsetsCrossPri1R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Cg_BetaOffsetsCrossPri1_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Cg_BetaOffsetsCrossPri1_r17).Choice {
				case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri1_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri1_r17_Setup:
					if err := (*ie.Cg_BetaOffsetsCrossPri1_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MappingPattern_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MappingPattern_r17, configuredGrantConfigExtMappingPatternR17Constraints); err != nil {
					return err
				}
			}

			if ie.SequenceOffsetForRV_r17 != nil {
				if err := ex.EncodeInteger(*ie.SequenceOffsetForRV_r17, configuredGrantConfigSequenceOffsetForRVR17Constraints); err != nil {
					return err
				}
			}

			if ie.P0_PUSCH_Alpha2_r17 != nil {
				if err := ie.P0_PUSCH_Alpha2_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PowerControlLoopToUse2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PowerControlLoopToUse2_r17, configuredGrantConfigExtPowerControlLoopToUse2R17Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_COT_SharingList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(configuredGrantConfigExtCgCOTSharingListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Cg_COT_SharingList_r17))); err != nil {
					return err
				}
				for i := range ie.Cg_COT_SharingList_r17 {
					if err := ie.Cg_COT_SharingList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PeriodicityExt_r17 != nil {
				if err := ex.EncodeInteger(*ie.PeriodicityExt_r17, configuredGrantConfigPeriodicityExtR17Constraints); err != nil {
					return err
				}
			}

			if ie.RepK_v1710 != nil {
				if err := ex.EncodeEnumerated(*ie.RepK_v1710, configuredGrantConfigExtRepKV1710Constraints); err != nil {
					return err
				}
			}

			if ie.NrofHARQ_Processes_v1700 != nil {
				if err := ex.EncodeInteger(*ie.NrofHARQ_Processes_v1700, configuredGrantConfigNrofHARQProcessesV1700Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcID_Offset2_v1700 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcID_Offset2_v1700, configuredGrantConfigHarqProcIDOffset2V1700Constraints); err != nil {
					return err
				}
			}

			if ie.ConfiguredGrantTimer_v1700 != nil {
				if err := ex.EncodeInteger(*ie.ConfiguredGrantTimer_v1700, configuredGrantConfigConfiguredGrantTimerV1700Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_MinDFI_Delay_v1710 != nil {
				if err := ex.EncodeInteger(*ie.Cg_MinDFI_Delay_v1710, configuredGrantConfigCgMinDFIDelayV1710Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "harq-ProcID-Offset-v1730", Optional: true},
					{Name: "cg-nrofSlots-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Harq_ProcID_Offset_v1730 != nil, ie.Cg_NrofSlots_r17 != nil}); err != nil {
				return err
			}

			if ie.Harq_ProcID_Offset_v1730 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcID_Offset_v1730, configuredGrantConfigHarqProcIDOffsetV1730Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_NrofSlots_r17 != nil {
				if err := ex.EncodeInteger(*ie.Cg_NrofSlots_r17, configuredGrantConfigCgNrofSlotsR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "disableCG-RetransmissionMonitoring-r18", Optional: true},
					{Name: "nrofSlotsInCG-Period-r18", Optional: true},
					{Name: "uto-UCI-Config-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DisableCG_RetransmissionMonitoring_r18 != nil, ie.NrofSlotsInCG_Period_r18 != nil, ie.Uto_UCI_Config_r18 != nil}); err != nil {
				return err
			}

			if ie.DisableCG_RetransmissionMonitoring_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DisableCG_RetransmissionMonitoring_r18, configuredGrantConfigExtDisableCGRetransmissionMonitoringR18Constraints); err != nil {
					return err
				}
			}

			if ie.NrofSlotsInCG_Period_r18 != nil {
				if err := ex.EncodeInteger(*ie.NrofSlotsInCG_Period_r18, configuredGrantConfigNrofSlotsInCGPeriodR18Constraints); err != nil {
					return err
				}
			}

			if ie.Uto_UCI_Config_r18 != nil {
				c := ie.Uto_UCI_Config_r18
				configuredGrantConfigExtUtoUCIConfigR18Seq := ex.NewSequenceEncoder(configuredGrantConfigExtUtoUCIConfigR18Constraints)
				if err := configuredGrantConfigExtUtoUCIConfigR18Seq.EncodeExtensionBit(false); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NrofBitsInUTO_UCI_r18, per.Constrained(3, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.BetaOffsetUTO_UCI_r18, per.Constrained(0, 31)); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "precodingAndNumberOfLayers-v1850", Optional: true},
					{Name: "srs-ResourceIndicator-v1850", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PrecodingAndNumberOfLayers_v1850 != nil, ie.Srs_ResourceIndicator_v1850 != nil}); err != nil {
				return err
			}

			if ie.PrecodingAndNumberOfLayers_v1850 != nil {
				if err := ex.EncodeInteger(*ie.PrecodingAndNumberOfLayers_v1850, configuredGrantConfigPrecodingAndNumberOfLayersV1850Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_ResourceIndicator_v1850 != nil {
				if err := ex.EncodeInteger(*ie.Srs_ResourceIndicator_v1850, configuredGrantConfigSrsResourceIndicatorV1850Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ConfiguredGrantConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(configuredGrantConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. frequencyHopping: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(configuredGrantConfigFrequencyHoppingConstraints)
			if err != nil {
				return err
			}
			ie.FrequencyHopping = &idx
		}
	}

	// 4. cg-DMRS-Configuration: ref
	{
		if err := ie.Cg_DMRS_Configuration.Decode(d); err != nil {
			return err
		}
	}

	// 5. mcs-Table: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(configuredGrantConfigMcsTableConstraints)
			if err != nil {
				return err
			}
			ie.Mcs_Table = &idx
		}
	}

	// 6. mcs-TableTransformPrecoder: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(configuredGrantConfigMcsTableTransformPrecoderConstraints)
			if err != nil {
				return err
			}
			ie.Mcs_TableTransformPrecoder = &idx
		}
	}

	// 7. uci-OnPUSCH: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Uci_OnPUSCH = &struct {
				Choice  int
				Release *struct{}
				Setup   *CG_UCI_OnPUSCH
			}{}
			choiceDec := d.NewChoiceDecoder(configuredGrantConfigUciOnPUSCHConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Uci_OnPUSCH).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ConfiguredGrantConfig_Uci_OnPUSCH_Release:
				(*ie.Uci_OnPUSCH).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ConfiguredGrantConfig_Uci_OnPUSCH_Setup:
				(*ie.Uci_OnPUSCH).Setup = new(CG_UCI_OnPUSCH)
				if err := (*ie.Uci_OnPUSCH).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. resourceAllocation: enumerated
	{
		v5, err := d.DecodeEnumerated(configuredGrantConfigResourceAllocationConstraints)
		if err != nil {
			return err
		}
		ie.ResourceAllocation = v5
	}

	// 9. rbg-Size: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(configuredGrantConfigRbgSizeConstraints)
			if err != nil {
				return err
			}
			ie.Rbg_Size = &idx
		}
	}

	// 10. powerControlLoopToUse: enumerated
	{
		v7, err := d.DecodeEnumerated(configuredGrantConfigPowerControlLoopToUseConstraints)
		if err != nil {
			return err
		}
		ie.PowerControlLoopToUse = v7
	}

	// 11. p0-PUSCH-Alpha: ref
	{
		if err := ie.P0_PUSCH_Alpha.Decode(d); err != nil {
			return err
		}
	}

	// 12. transformPrecoder: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(configuredGrantConfigTransformPrecoderConstraints)
			if err != nil {
				return err
			}
			ie.TransformPrecoder = &idx
		}
	}

	// 13. nrofHARQ-Processes: integer
	{
		v10, err := d.DecodeInteger(configuredGrantConfigNrofHARQProcessesConstraints)
		if err != nil {
			return err
		}
		ie.NrofHARQ_Processes = v10
	}

	// 14. repK: enumerated
	{
		v11, err := d.DecodeEnumerated(configuredGrantConfigRepKConstraints)
		if err != nil {
			return err
		}
		ie.RepK = v11
	}

	// 15. repK-RV: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(configuredGrantConfigRepKRVConstraints)
			if err != nil {
				return err
			}
			ie.RepK_RV = &idx
		}
	}

	// 16. periodicity: enumerated
	{
		v13, err := d.DecodeEnumerated(configuredGrantConfigPeriodicityConstraints)
		if err != nil {
			return err
		}
		ie.Periodicity = v13
	}

	// 17. configuredGrantTimer: integer
	{
		if seq.IsComponentPresent(14) {
			v, err := d.DecodeInteger(configuredGrantConfigConfiguredGrantTimerConstraints)
			if err != nil {
				return err
			}
			ie.ConfiguredGrantTimer = &v
		}
	}

	// 18. rrc-ConfiguredUplinkGrant: sequence
	{
		if seq.IsComponentPresent(15) {
			ie.Rrc_ConfiguredUplinkGrant = &struct {
				TimeDomainOffset           int64
				TimeDomainAllocation       int64
				FrequencyDomainAllocation  per.BitString
				AntennaPort                int64
				Dmrs_SeqInitialization     *int64
				PrecodingAndNumberOfLayers int64
				Srs_ResourceIndicator      *int64
				McsAndTBS                  int64
				FrequencyHoppingOffset     *int64
				PathlossReferenceIndex     int64
			}{}
			c := ie.Rrc_ConfiguredUplinkGrant
			configuredGrantConfigRrcConfiguredUplinkGrantSeq := d.NewSequenceDecoder(configuredGrantConfigRrcConfiguredUplinkGrantConstraints)
			if err := configuredGrantConfigRrcConfiguredUplinkGrantSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := configuredGrantConfigRrcConfiguredUplinkGrantSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 5119))
				if err != nil {
					return err
				}
				c.TimeDomainOffset = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.TimeDomainAllocation = v
			}
			{
				v, err := d.DecodeBitString(per.FixedSize(18))
				if err != nil {
					return err
				}
				c.FrequencyDomainAllocation = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 31))
				if err != nil {
					return err
				}
				c.AntennaPort = v
			}
			if configuredGrantConfigRrcConfiguredUplinkGrantSeq.IsComponentPresent(4) {
				c.Dmrs_SeqInitialization = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*c.Dmrs_SeqInitialization) = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 63))
				if err != nil {
					return err
				}
				c.PrecodingAndNumberOfLayers = v
			}
			if configuredGrantConfigRrcConfiguredUplinkGrantSeq.IsComponentPresent(6) {
				c.Srs_ResourceIndicator = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				(*c.Srs_ResourceIndicator) = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 31))
				if err != nil {
					return err
				}
				c.McsAndTBS = v
			}
			if configuredGrantConfigRrcConfiguredUplinkGrantSeq.IsComponentPresent(8) {
				c.FrequencyHoppingOffset = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1))
				if err != nil {
					return err
				}
				(*c.FrequencyHoppingOffset) = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofPUSCH_PathlossReferenceRSs_1))
				if err != nil {
					return err
				}
				c.PathlossReferenceIndex = v
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cg-RetransmissionTimer-r16", Optional: true},
				{Name: "cg-minDFI-Delay-r16", Optional: true},
				{Name: "cg-nrofPUSCH-InSlot-r16", Optional: true},
				{Name: "cg-nrofSlots-r16", Optional: true},
				{Name: "cg-StartingOffsets-r16", Optional: true},
				{Name: "cg-UCI-Multiplexing-r16", Optional: true},
				{Name: "cg-COT-SharingOffset-r16", Optional: true},
				{Name: "betaOffsetCG-UCI-r16", Optional: true},
				{Name: "cg-COT-SharingList-r16", Optional: true},
				{Name: "harq-ProcID-Offset-r16", Optional: true},
				{Name: "harq-ProcID-Offset2-r16", Optional: true},
				{Name: "configuredGrantConfigIndex-r16", Optional: true},
				{Name: "configuredGrantConfigIndexMAC-r16", Optional: true},
				{Name: "periodicityExt-r16", Optional: true},
				{Name: "startingFromRV0-r16", Optional: true},
				{Name: "phy-PriorityIndex-r16", Optional: true},
				{Name: "autonomousTx-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(configuredGrantConfigCgRetransmissionTimerR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_RetransmissionTimer_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtCgMinDFIDelayR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_MinDFI_Delay_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(configuredGrantConfigCgNrofPUSCHInSlotR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_NrofPUSCH_InSlot_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(configuredGrantConfigCgNrofSlotsR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_NrofSlots_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Cg_StartingOffsets_r16 = new(CG_StartingOffsets_r16)
			if err := ie.Cg_StartingOffsets_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtCgUCIMultiplexingR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_UCI_Multiplexing_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeInteger(configuredGrantConfigCgCOTSharingOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_COT_SharingOffset_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeInteger(configuredGrantConfigBetaOffsetCGUCIR16Constraints)
			if err != nil {
				return err
			}
			ie.BetaOffsetCG_UCI_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(configuredGrantConfigExtCgCOTSharingListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cg_COT_SharingList_r16 = make([]CG_COT_Sharing_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cg_COT_SharingList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeInteger(configuredGrantConfigHarqProcIDOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcID_Offset_r16 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeInteger(configuredGrantConfigHarqProcIDOffset2R16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcID_Offset2_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			ie.ConfiguredGrantConfigIndex_r16 = new(ConfiguredGrantConfigIndex_r16)
			if err := ie.ConfiguredGrantConfigIndex_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(12) {
			ie.ConfiguredGrantConfigIndexMAC_r16 = new(ConfiguredGrantConfigIndexMAC_r16)
			if err := ie.ConfiguredGrantConfigIndexMAC_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeInteger(configuredGrantConfigPeriodicityExtR16Constraints)
			if err != nil {
				return err
			}
			ie.PeriodicityExt_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtStartingFromRV0R16Constraints)
			if err != nil {
				return err
			}
			ie.StartingFromRV0_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtPhyPriorityIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.Phy_PriorityIndex_r16 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtAutonomousTxR16Constraints)
			if err != nil {
				return err
			}
			ie.AutonomousTx_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cg-betaOffsetsCrossPri0-r17", Optional: true},
				{Name: "cg-betaOffsetsCrossPri1-r17", Optional: true},
				{Name: "mappingPattern-r17", Optional: true},
				{Name: "sequenceOffsetForRV-r17", Optional: true},
				{Name: "p0-PUSCH-Alpha2-r17", Optional: true},
				{Name: "powerControlLoopToUse2-r17", Optional: true},
				{Name: "cg-COT-SharingList-r17", Optional: true},
				{Name: "periodicityExt-r17", Optional: true},
				{Name: "repK-v1710", Optional: true},
				{Name: "nrofHARQ-Processes-v1700", Optional: true},
				{Name: "harq-ProcID-Offset2-v1700", Optional: true},
				{Name: "configuredGrantTimer-v1700", Optional: true},
				{Name: "cg-minDFI-Delay-v1710", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cg_BetaOffsetsCrossPri0_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BetaOffsetsCrossPriSelCG_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(configuredGrantConfigExtCgBetaOffsetsCrossPri0R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cg_BetaOffsetsCrossPri0_r17).Choice = int(index)
			switch index {
			case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri0_r17_Release:
				(*ie.Cg_BetaOffsetsCrossPri0_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri0_r17_Setup:
				(*ie.Cg_BetaOffsetsCrossPri0_r17).Setup = new(BetaOffsetsCrossPriSelCG_r17)
				if err := (*ie.Cg_BetaOffsetsCrossPri0_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Cg_BetaOffsetsCrossPri1_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BetaOffsetsCrossPriSelCG_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(configuredGrantConfigExtCgBetaOffsetsCrossPri1R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cg_BetaOffsetsCrossPri1_r17).Choice = int(index)
			switch index {
			case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri1_r17_Release:
				(*ie.Cg_BetaOffsetsCrossPri1_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ConfiguredGrantConfig_Ext_Cg_BetaOffsetsCrossPri1_r17_Setup:
				(*ie.Cg_BetaOffsetsCrossPri1_r17).Setup = new(BetaOffsetsCrossPriSelCG_r17)
				if err := (*ie.Cg_BetaOffsetsCrossPri1_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtMappingPatternR17Constraints)
			if err != nil {
				return err
			}
			ie.MappingPattern_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(configuredGrantConfigSequenceOffsetForRVR17Constraints)
			if err != nil {
				return err
			}
			ie.SequenceOffsetForRV_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.P0_PUSCH_Alpha2_r17 = new(P0_PUSCH_AlphaSetId)
			if err := ie.P0_PUSCH_Alpha2_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtPowerControlLoopToUse2R17Constraints)
			if err != nil {
				return err
			}
			ie.PowerControlLoopToUse2_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			seqOf := dx.NewSequenceOfDecoder(configuredGrantConfigExtCgCOTSharingListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cg_COT_SharingList_r17 = make([]CG_COT_Sharing_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cg_COT_SharingList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeInteger(configuredGrantConfigPeriodicityExtR17Constraints)
			if err != nil {
				return err
			}
			ie.PeriodicityExt_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtRepKV1710Constraints)
			if err != nil {
				return err
			}
			ie.RepK_v1710 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeInteger(configuredGrantConfigNrofHARQProcessesV1700Constraints)
			if err != nil {
				return err
			}
			ie.NrofHARQ_Processes_v1700 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeInteger(configuredGrantConfigHarqProcIDOffset2V1700Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcID_Offset2_v1700 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeInteger(configuredGrantConfigConfiguredGrantTimerV1700Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredGrantTimer_v1700 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeInteger(configuredGrantConfigCgMinDFIDelayV1710Constraints)
			if err != nil {
				return err
			}
			ie.Cg_MinDFI_Delay_v1710 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "harq-ProcID-Offset-v1730", Optional: true},
				{Name: "cg-nrofSlots-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(configuredGrantConfigHarqProcIDOffsetV1730Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcID_Offset_v1730 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(configuredGrantConfigCgNrofSlotsR17Constraints)
			if err != nil {
				return err
			}
			ie.Cg_NrofSlots_r17 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "disableCG-RetransmissionMonitoring-r18", Optional: true},
				{Name: "nrofSlotsInCG-Period-r18", Optional: true},
				{Name: "uto-UCI-Config-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(configuredGrantConfigExtDisableCGRetransmissionMonitoringR18Constraints)
			if err != nil {
				return err
			}
			ie.DisableCG_RetransmissionMonitoring_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(configuredGrantConfigNrofSlotsInCGPeriodR18Constraints)
			if err != nil {
				return err
			}
			ie.NrofSlotsInCG_Period_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Uto_UCI_Config_r18 = &struct {
				NrofBitsInUTO_UCI_r18 int64
				BetaOffsetUTO_UCI_r18 int64
			}{}
			c := ie.Uto_UCI_Config_r18
			configuredGrantConfigExtUtoUCIConfigR18Seq := dx.NewSequenceDecoder(configuredGrantConfigExtUtoUCIConfigR18Constraints)
			if err := configuredGrantConfigExtUtoUCIConfigR18Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(3, 8))
				if err != nil {
					return err
				}
				c.NrofBitsInUTO_UCI_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 31))
				if err != nil {
					return err
				}
				c.BetaOffsetUTO_UCI_r18 = v
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "precodingAndNumberOfLayers-v1850", Optional: true},
				{Name: "srs-ResourceIndicator-v1850", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(configuredGrantConfigPrecodingAndNumberOfLayersV1850Constraints)
			if err != nil {
				return err
			}
			ie.PrecodingAndNumberOfLayers_v1850 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(configuredGrantConfigSrsResourceIndicatorV1850Constraints)
			if err != nil {
				return err
			}
			ie.Srs_ResourceIndicator_v1850 = &v
		}
	}

	return nil
}
