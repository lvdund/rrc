// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ResourcePool-r16 (line 27939).

var sLResourcePoolR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PSCCH-Config-r16", Optional: true},
		{Name: "sl-PSSCH-Config-r16", Optional: true},
		{Name: "sl-PSFCH-Config-r16", Optional: true},
		{Name: "sl-SyncAllowed-r16", Optional: true},
		{Name: "sl-SubchannelSize-r16", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "sl-StartRB-Subchannel-r16", Optional: true},
		{Name: "sl-NumSubchannel-r16", Optional: true},
		{Name: "sl-Additional-MCS-Table-r16", Optional: true},
		{Name: "sl-ThreshS-RSSI-CBR-r16", Optional: true},
		{Name: "sl-TimeWindowSizeCBR-r16", Optional: true},
		{Name: "sl-TimeWindowSizeCR-r16", Optional: true},
		{Name: "sl-PTRS-Config-r16", Optional: true},
		{Name: "sl-UE-SelectedConfigRP-r16", Optional: true},
		{Name: "sl-RxParametersNcell-r16", Optional: true},
		{Name: "sl-ZoneConfigMCR-List-r16", Optional: true},
		{Name: "sl-FilterCoefficient-r16", Optional: true},
		{Name: "sl-RB-Number-r16", Optional: true},
		{Name: "sl-PreemptionEnable-r16", Optional: true},
		{Name: "sl-PriorityThreshold-UL-URLLC-r16", Optional: true},
		{Name: "sl-PriorityThreshold-r16", Optional: true},
		{Name: "sl-X-Overhead-r16", Optional: true},
		{Name: "sl-PowerControl-r16", Optional: true},
		{Name: "sl-TxPercentageList-r16", Optional: true},
		{Name: "sl-MinMaxMCS-List-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var sL_ResourcePool_r16SlPSCCHConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ResourcePool_r16_Sl_PSCCH_Config_r16_Release = 0
	SL_ResourcePool_r16_Sl_PSCCH_Config_r16_Setup   = 1
)

var sL_ResourcePool_r16SlPSSCHConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ResourcePool_r16_Sl_PSSCH_Config_r16_Release = 0
	SL_ResourcePool_r16_Sl_PSSCH_Config_r16_Setup   = 1
)

var sL_ResourcePool_r16SlPSFCHConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ResourcePool_r16_Sl_PSFCH_Config_r16_Release = 0
	SL_ResourcePool_r16_Sl_PSFCH_Config_r16_Setup   = 1
)

const (
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N10  = 0
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N12  = 1
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N15  = 2
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N20  = 3
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N25  = 4
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N50  = 5
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N75  = 6
	SL_ResourcePool_r16_Sl_SubchannelSize_r16_N100 = 7
)

var sLResourcePoolR16SlSubchannelSizeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Sl_SubchannelSize_r16_N10, SL_ResourcePool_r16_Sl_SubchannelSize_r16_N12, SL_ResourcePool_r16_Sl_SubchannelSize_r16_N15, SL_ResourcePool_r16_Sl_SubchannelSize_r16_N20, SL_ResourcePool_r16_Sl_SubchannelSize_r16_N25, SL_ResourcePool_r16_Sl_SubchannelSize_r16_N50, SL_ResourcePool_r16_Sl_SubchannelSize_r16_N75, SL_ResourcePool_r16_Sl_SubchannelSize_r16_N100},
}

var sLResourcePoolR16DummyConstraints = per.Constrained(10, 160)

var sLResourcePoolR16SlStartRBSubchannelR16Constraints = per.Constrained(0, 265)

var sLResourcePoolR16SlNumSubchannelR16Constraints = per.Constrained(1, 27)

const (
	SL_ResourcePool_r16_Sl_Additional_MCS_Table_r16_Qam256            = 0
	SL_ResourcePool_r16_Sl_Additional_MCS_Table_r16_Qam64LowSE        = 1
	SL_ResourcePool_r16_Sl_Additional_MCS_Table_r16_Qam256_Qam64LowSE = 2
)

var sLResourcePoolR16SlAdditionalMCSTableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Sl_Additional_MCS_Table_r16_Qam256, SL_ResourcePool_r16_Sl_Additional_MCS_Table_r16_Qam64LowSE, SL_ResourcePool_r16_Sl_Additional_MCS_Table_r16_Qam256_Qam64LowSE},
}

var sLResourcePoolR16SlThreshSRSSICBRR16Constraints = per.Constrained(0, 45)

const (
	SL_ResourcePool_r16_Sl_TimeWindowSizeCBR_r16_Ms100   = 0
	SL_ResourcePool_r16_Sl_TimeWindowSizeCBR_r16_Slot100 = 1
)

var sLResourcePoolR16SlTimeWindowSizeCBRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Sl_TimeWindowSizeCBR_r16_Ms100, SL_ResourcePool_r16_Sl_TimeWindowSizeCBR_r16_Slot100},
}

const (
	SL_ResourcePool_r16_Sl_TimeWindowSizeCR_r16_Ms1000   = 0
	SL_ResourcePool_r16_Sl_TimeWindowSizeCR_r16_Slot1000 = 1
)

var sLResourcePoolR16SlTimeWindowSizeCRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Sl_TimeWindowSizeCR_r16_Ms1000, SL_ResourcePool_r16_Sl_TimeWindowSizeCR_r16_Slot1000},
}

var sLResourcePoolR16SlZoneConfigMCRListR16Constraints = per.FixedSize(16)

var sLResourcePoolR16SlRBNumberR16Constraints = per.Constrained(10, 275)

const (
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Enabled = 0
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl1     = 1
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl2     = 2
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl3     = 3
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl4     = 4
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl5     = 5
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl6     = 6
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl7     = 7
	SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl8     = 8
)

var sLResourcePoolR16SlPreemptionEnableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Enabled, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl1, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl2, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl3, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl4, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl5, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl6, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl7, SL_ResourcePool_r16_Sl_PreemptionEnable_r16_Pl8},
}

var sLResourcePoolR16SlPriorityThresholdULURLLCR16Constraints = per.Constrained(1, 9)

var sLResourcePoolR16SlPriorityThresholdR16Constraints = per.Constrained(1, 9)

const (
	SL_ResourcePool_r16_Sl_X_Overhead_r16_N0 = 0
	SL_ResourcePool_r16_Sl_X_Overhead_r16_N3 = 1
	SL_ResourcePool_r16_Sl_X_Overhead_r16_N6 = 2
	SL_ResourcePool_r16_Sl_X_Overhead_r16_N9 = 3
)

var sLResourcePoolR16SlXOverheadR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Sl_X_Overhead_r16_N0, SL_ResourcePool_r16_Sl_X_Overhead_r16_N3, SL_ResourcePool_r16_Sl_X_Overhead_r16_N6, SL_ResourcePool_r16_Sl_X_Overhead_r16_N9},
}

var sLResourcePoolR16SlTimeResourceR16Constraints = per.SizeRange(10, 160)

var sLResourcePoolR16ExtSlPBPSCPSConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ResourcePool_r16_Ext_Sl_PBPS_CPS_Config_r17_Release = 0
	SL_ResourcePool_r16_Ext_Sl_PBPS_CPS_Config_r17_Setup   = 1
)

var sLResourcePoolR16ExtSlInterUECoordinationConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ResourcePool_r16_Ext_Sl_InterUE_CoordinationConfig_r17_Release = 0
	SL_ResourcePool_r16_Ext_Sl_InterUE_CoordinationConfig_r17_Setup   = 1
)

var sLResourcePoolR16ExtSlCPEStartingPositionsPSCCHPSSCHInitiateCOTListR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18_Release = 0
	SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18_Setup   = 1
)

var sLResourcePoolR16SlCPEStartingPositionsPSCCHPSSCHInitiateCOTDefaultR18Constraints = per.Constrained(1, 9)

var sLResourcePoolR16ExtSlCPEStartingPositionsPSCCHPSSCHWithinCOTListR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18_Release = 0
	SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18_Setup   = 1
)

var sLResourcePoolR16SlCPEStartingPositionsPSCCHPSSCHWithinCOTDefaultR18Constraints = per.Constrained(1, 9)

const (
	SL_ResourcePool_r16_Ext_Sl_Type1_LBT_BlockingOption1_r18_Enabled = 0
)

var sLResourcePoolR16ExtSlType1LBTBlockingOption1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_Type1_LBT_BlockingOption1_r18_Enabled},
}

const (
	SL_ResourcePool_r16_Ext_Sl_Type1_LBT_BlockingOption2_r18_Enabled = 0
)

var sLResourcePoolR16ExtSlType1LBTBlockingOption2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_Type1_LBT_BlockingOption2_r18_Enabled},
}

const (
	SL_ResourcePool_r16_Ext_Sl_NumInterlacePerSubchannel_r18_Sc1 = 0
	SL_ResourcePool_r16_Ext_Sl_NumInterlacePerSubchannel_r18_Sc2 = 1
)

var sLResourcePoolR16ExtSlNumInterlacePerSubchannelR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_NumInterlacePerSubchannel_r18_Sc1, SL_ResourcePool_r16_Ext_Sl_NumInterlacePerSubchannel_r18_Sc2},
}

const (
	SL_ResourcePool_r16_Ext_Sl_NumReferencePRBs_OfInterlace_r18_Prb10 = 0
	SL_ResourcePool_r16_Ext_Sl_NumReferencePRBs_OfInterlace_r18_Prb11 = 1
)

var sLResourcePoolR16ExtSlNumReferencePRBsOfInterlaceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_NumReferencePRBs_OfInterlace_r18_Prb10, SL_ResourcePool_r16_Ext_Sl_NumReferencePRBs_OfInterlace_r18_Prb11},
}

const (
	SL_ResourcePool_r16_Ext_Sl_TransmissionStructureForPSFCH_r18_CommonInterlace    = 0
	SL_ResourcePool_r16_Ext_Sl_TransmissionStructureForPSFCH_r18_DedicatedInterlace = 1
)

var sLResourcePoolR16ExtSlTransmissionStructureForPSFCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_TransmissionStructureForPSFCH_r18_CommonInterlace, SL_ResourcePool_r16_Ext_Sl_TransmissionStructureForPSFCH_r18_DedicatedInterlace},
}

const (
	SL_ResourcePool_r16_Ext_Sl_NumDedicatedPRBs_ForPSFCH_r18_Prb1 = 0
	SL_ResourcePool_r16_Ext_Sl_NumDedicatedPRBs_ForPSFCH_r18_Prb2 = 1
	SL_ResourcePool_r16_Ext_Sl_NumDedicatedPRBs_ForPSFCH_r18_Prb5 = 2
)

var sLResourcePoolR16ExtSlNumDedicatedPRBsForPSFCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_NumDedicatedPRBs_ForPSFCH_r18_Prb1, SL_ResourcePool_r16_Ext_Sl_NumDedicatedPRBs_ForPSFCH_r18_Prb2, SL_ResourcePool_r16_Ext_Sl_NumDedicatedPRBs_ForPSFCH_r18_Prb5},
}

const (
	SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O1 = 0
	SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O2 = 1
	SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O3 = 2
	SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O4 = 3
)

var sLResourcePoolR16ExtSlNumPSFCHOccasionsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O1, SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O2, SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O3, SL_ResourcePool_r16_Ext_Sl_NumPSFCH_Occasions_r18_O4},
}

var sLResourcePoolR16SlPSFCHCommonInterlaceIndexR18Constraints = per.Constrained(0, 9)

var sLResourcePoolR16SlCPEStartingPositionPSFCHR18Constraints = per.Constrained(1, 9)

const (
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym7  = 0
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym8  = 1
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym9  = 2
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym10 = 3
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym11 = 4
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym12 = 5
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym13 = 6
	SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym14 = 7
)

var sLResourcePoolR16ExtSlNumRefSymbolLengthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym7, SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym8, SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym9, SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym10, SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym11, SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym12, SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym13, SL_ResourcePool_r16_Ext_Sl_NumRefSymbolLength_r18_Sym14},
}

var sLResourcePoolR16ExtSlPSFCHRBSetListR18Constraints = per.SizeRange(1, 4)

var sLResourcePoolR16ExtSlIUCRBSetListR18Constraints = per.SizeRange(1, 4)

var sLResourcePoolR16SlPSFCHPowerOffsetR18Constraints = per.Constrained(0, 10)

var sLResourcePoolR16ExtSlRBSetIndexOfResourcePoolR18Constraints = per.SizeRange(1, 5)

const (
	SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_Brid       = 0
	SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_Daa        = 1
	SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_BridAndDAA = 2
	SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_Spare1     = 3
)

var sLResourcePoolR16ExtSlA2XServiceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_Brid, SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_Daa, SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_BridAndDAA, SL_ResourcePool_r16_Ext_Sl_A2X_Service_r18_Spare1},
}

var sLResourcePoolR16ExtSlPRSResourcesSharedSLPRSRPR18Constraints = per.SizeRange(1, 17)

var sLResourcePoolR16NumSymSLPRS2ndStageSCIR18Constraints = per.Constrained(1, 4)

var sLResourcePoolR16SlRxParametersNcellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TDD-Configuration-r16", Optional: true},
		{Name: "sl-SyncConfigIndex-r16"},
	},
}

type SL_ResourcePool_r16 struct {
	Sl_PSCCH_Config_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_PSCCH_Config_r16
	}
	Sl_PSSCH_Config_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_PSSCH_Config_r16
	}
	Sl_PSFCH_Config_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_PSFCH_Config_r16
	}
	Sl_SyncAllowed_r16          *SL_SyncAllowed_r16
	Sl_SubchannelSize_r16       *int64
	Dummy                       *int64
	Sl_StartRB_Subchannel_r16   *int64
	Sl_NumSubchannel_r16        *int64
	Sl_Additional_MCS_Table_r16 *int64
	Sl_ThreshS_RSSI_CBR_r16     *int64
	Sl_TimeWindowSizeCBR_r16    *int64
	Sl_TimeWindowSizeCR_r16     *int64
	Sl_PTRS_Config_r16          *SL_PTRS_Config_r16
	Sl_UE_SelectedConfigRP_r16  *SL_UE_SelectedConfigRP_r16
	Sl_RxParametersNcell_r16    *struct {
		Sl_TDD_Configuration_r16 *TDD_UL_DL_ConfigCommon
		Sl_SyncConfigIndex_r16   int64
	}
	Sl_ZoneConfigMCR_List_r16         []SL_ZoneConfigMCR_r16
	Sl_FilterCoefficient_r16          *FilterCoefficient
	Sl_RB_Number_r16                  *int64
	Sl_PreemptionEnable_r16           *int64
	Sl_PriorityThreshold_UL_URLLC_r16 *int64
	Sl_PriorityThreshold_r16          *int64
	Sl_X_Overhead_r16                 *int64
	Sl_PowerControl_r16               *SL_PowerControl_r16
	Sl_TxPercentageList_r16           *SL_TxPercentageList_r16
	Sl_MinMaxMCS_List_r16             *SL_MinMaxMCS_List_r16
	Sl_TimeResource_r16               *per.BitString
	Sl_PBPS_CPS_Config_r17            *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_PBPS_CPS_Config_r17
	}
	Sl_InterUE_CoordinationConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_InterUE_CoordinationConfig_r17
	}
	Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18
	}
	Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_Default_r18 *int64
	Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18      *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18
	}
	Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_Default_r18 *int64
	Sl_Type1_LBT_BlockingOption1_r18                          *int64
	Sl_Type1_LBT_BlockingOption2_r18                          *int64
	Sl_NumInterlacePerSubchannel_r18                          *int64
	Sl_NumReferencePRBs_OfInterlace_r18                       *int64
	Sl_TransmissionStructureForPSFCH_r18                      *int64
	Sl_NumDedicatedPRBs_ForPSFCH_r18                          *int64
	Sl_NumPSFCH_Occasions_r18                                 *int64
	Sl_PSFCH_CommonInterlaceIndex_r18                         *int64
	Sl_CPE_StartingPositionPSFCH_r18                          *int64
	Sl_NumRefSymbolLength_r18                                 *int64
	Sl_PSFCH_RB_SetList_r18                                   []per.BitString
	Sl_IUC_RB_SetList_r18                                     []per.BitString
	Sl_PSFCH_PowerOffset_r18                                  *int64
	Sl_RBSetIndexOfResourcePool_r18                           []int64
	Sl_A2X_Service_r18                                        *int64
	Sl_PRS_ResourcesSharedSL_PRS_RP_r18                       []SL_PRS_ResourceSharedSL_PRS_RP_r18
	NumSym_SL_PRS_2ndStageSCI_r18                             *int64
	Sl_SCI_BasedSL_PRS_TxTriggerSCI2_D_r18                    *bool
}

func (ie *SL_ResourcePool_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLResourcePoolR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_TimeResource_r16 != nil
	hasExtGroup1 := ie.Sl_PBPS_CPS_Config_r17 != nil || ie.Sl_InterUE_CoordinationConfig_r17 != nil
	hasExtGroup2 := ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18 != nil || ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_Default_r18 != nil || ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18 != nil || ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_Default_r18 != nil || ie.Sl_Type1_LBT_BlockingOption1_r18 != nil || ie.Sl_Type1_LBT_BlockingOption2_r18 != nil || ie.Sl_NumInterlacePerSubchannel_r18 != nil || ie.Sl_NumReferencePRBs_OfInterlace_r18 != nil || ie.Sl_TransmissionStructureForPSFCH_r18 != nil || ie.Sl_NumDedicatedPRBs_ForPSFCH_r18 != nil || ie.Sl_NumPSFCH_Occasions_r18 != nil || ie.Sl_PSFCH_CommonInterlaceIndex_r18 != nil || ie.Sl_CPE_StartingPositionPSFCH_r18 != nil || ie.Sl_NumRefSymbolLength_r18 != nil || ie.Sl_PSFCH_RB_SetList_r18 != nil || ie.Sl_IUC_RB_SetList_r18 != nil || ie.Sl_PSFCH_PowerOffset_r18 != nil || ie.Sl_RBSetIndexOfResourcePool_r18 != nil || ie.Sl_A2X_Service_r18 != nil || ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18 != nil || ie.NumSym_SL_PRS_2ndStageSCI_r18 != nil || ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI2_D_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PSCCH_Config_r16 != nil, ie.Sl_PSSCH_Config_r16 != nil, ie.Sl_PSFCH_Config_r16 != nil, ie.Sl_SyncAllowed_r16 != nil, ie.Sl_SubchannelSize_r16 != nil, ie.Dummy != nil, ie.Sl_StartRB_Subchannel_r16 != nil, ie.Sl_NumSubchannel_r16 != nil, ie.Sl_Additional_MCS_Table_r16 != nil, ie.Sl_ThreshS_RSSI_CBR_r16 != nil, ie.Sl_TimeWindowSizeCBR_r16 != nil, ie.Sl_TimeWindowSizeCR_r16 != nil, ie.Sl_PTRS_Config_r16 != nil, ie.Sl_UE_SelectedConfigRP_r16 != nil, ie.Sl_RxParametersNcell_r16 != nil, ie.Sl_ZoneConfigMCR_List_r16 != nil, ie.Sl_FilterCoefficient_r16 != nil, ie.Sl_RB_Number_r16 != nil, ie.Sl_PreemptionEnable_r16 != nil, ie.Sl_PriorityThreshold_UL_URLLC_r16 != nil, ie.Sl_PriorityThreshold_r16 != nil, ie.Sl_X_Overhead_r16 != nil, ie.Sl_PowerControl_r16 != nil, ie.Sl_TxPercentageList_r16 != nil, ie.Sl_MinMaxMCS_List_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-PSCCH-Config-r16: choice
	{
		if ie.Sl_PSCCH_Config_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_ResourcePool_r16SlPSCCHConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PSCCH_Config_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_PSCCH_Config_r16).Choice {
			case SL_ResourcePool_r16_Sl_PSCCH_Config_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Sl_PSCCH_Config_r16_Setup:
				if err := (*ie.Sl_PSCCH_Config_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_PSCCH_Config_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. sl-PSSCH-Config-r16: choice
	{
		if ie.Sl_PSSCH_Config_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_ResourcePool_r16SlPSSCHConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PSSCH_Config_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_PSSCH_Config_r16).Choice {
			case SL_ResourcePool_r16_Sl_PSSCH_Config_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Sl_PSSCH_Config_r16_Setup:
				if err := (*ie.Sl_PSSCH_Config_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_PSSCH_Config_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. sl-PSFCH-Config-r16: choice
	{
		if ie.Sl_PSFCH_Config_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_ResourcePool_r16SlPSFCHConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PSFCH_Config_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_PSFCH_Config_r16).Choice {
			case SL_ResourcePool_r16_Sl_PSFCH_Config_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Sl_PSFCH_Config_r16_Setup:
				if err := (*ie.Sl_PSFCH_Config_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_PSFCH_Config_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. sl-SyncAllowed-r16: ref
	{
		if ie.Sl_SyncAllowed_r16 != nil {
			if err := ie.Sl_SyncAllowed_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. sl-SubchannelSize-r16: enumerated
	{
		if ie.Sl_SubchannelSize_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SubchannelSize_r16, sLResourcePoolR16SlSubchannelSizeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. dummy: integer
	{
		if ie.Dummy != nil {
			if err := e.EncodeInteger(*ie.Dummy, sLResourcePoolR16DummyConstraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-StartRB-Subchannel-r16: integer
	{
		if ie.Sl_StartRB_Subchannel_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_StartRB_Subchannel_r16, sLResourcePoolR16SlStartRBSubchannelR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-NumSubchannel-r16: integer
	{
		if ie.Sl_NumSubchannel_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumSubchannel_r16, sLResourcePoolR16SlNumSubchannelR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sl-Additional-MCS-Table-r16: enumerated
	{
		if ie.Sl_Additional_MCS_Table_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Additional_MCS_Table_r16, sLResourcePoolR16SlAdditionalMCSTableR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-ThreshS-RSSI-CBR-r16: integer
	{
		if ie.Sl_ThreshS_RSSI_CBR_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_ThreshS_RSSI_CBR_r16, sLResourcePoolR16SlThreshSRSSICBRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 13. sl-TimeWindowSizeCBR-r16: enumerated
	{
		if ie.Sl_TimeWindowSizeCBR_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TimeWindowSizeCBR_r16, sLResourcePoolR16SlTimeWindowSizeCBRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 14. sl-TimeWindowSizeCR-r16: enumerated
	{
		if ie.Sl_TimeWindowSizeCR_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TimeWindowSizeCR_r16, sLResourcePoolR16SlTimeWindowSizeCRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 15. sl-PTRS-Config-r16: ref
	{
		if ie.Sl_PTRS_Config_r16 != nil {
			if err := ie.Sl_PTRS_Config_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 16. sl-UE-SelectedConfigRP-r16: ref
	{
		if ie.Sl_UE_SelectedConfigRP_r16 != nil {
			if err := ie.Sl_UE_SelectedConfigRP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 17. sl-RxParametersNcell-r16: sequence
	{
		if ie.Sl_RxParametersNcell_r16 != nil {
			c := ie.Sl_RxParametersNcell_r16
			sLResourcePoolR16SlRxParametersNcellR16Seq := e.NewSequenceEncoder(sLResourcePoolR16SlRxParametersNcellR16Constraints)
			if err := sLResourcePoolR16SlRxParametersNcellR16Seq.EncodePreamble([]bool{c.Sl_TDD_Configuration_r16 != nil}); err != nil {
				return err
			}
			if c.Sl_TDD_Configuration_r16 != nil {
				if err := c.Sl_TDD_Configuration_r16.Encode(e); err != nil {
					return err
				}
			}
			if err := e.EncodeInteger(c.Sl_SyncConfigIndex_r16, per.Constrained(0, 15)); err != nil {
				return err
			}
		}
	}

	// 18. sl-ZoneConfigMCR-List-r16: sequence-of
	{
		if ie.Sl_ZoneConfigMCR_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLResourcePoolR16SlZoneConfigMCRListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ZoneConfigMCR_List_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_ZoneConfigMCR_List_r16 {
				if err := ie.Sl_ZoneConfigMCR_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 19. sl-FilterCoefficient-r16: ref
	{
		if ie.Sl_FilterCoefficient_r16 != nil {
			if err := ie.Sl_FilterCoefficient_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 20. sl-RB-Number-r16: integer
	{
		if ie.Sl_RB_Number_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_RB_Number_r16, sLResourcePoolR16SlRBNumberR16Constraints); err != nil {
				return err
			}
		}
	}

	// 21. sl-PreemptionEnable-r16: enumerated
	{
		if ie.Sl_PreemptionEnable_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PreemptionEnable_r16, sLResourcePoolR16SlPreemptionEnableR16Constraints); err != nil {
				return err
			}
		}
	}

	// 22. sl-PriorityThreshold-UL-URLLC-r16: integer
	{
		if ie.Sl_PriorityThreshold_UL_URLLC_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityThreshold_UL_URLLC_r16, sLResourcePoolR16SlPriorityThresholdULURLLCR16Constraints); err != nil {
				return err
			}
		}
	}

	// 23. sl-PriorityThreshold-r16: integer
	{
		if ie.Sl_PriorityThreshold_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityThreshold_r16, sLResourcePoolR16SlPriorityThresholdR16Constraints); err != nil {
				return err
			}
		}
	}

	// 24. sl-X-Overhead-r16: enumerated
	{
		if ie.Sl_X_Overhead_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_X_Overhead_r16, sLResourcePoolR16SlXOverheadR16Constraints); err != nil {
				return err
			}
		}
	}

	// 25. sl-PowerControl-r16: ref
	{
		if ie.Sl_PowerControl_r16 != nil {
			if err := ie.Sl_PowerControl_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 26. sl-TxPercentageList-r16: ref
	{
		if ie.Sl_TxPercentageList_r16 != nil {
			if err := ie.Sl_TxPercentageList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 27. sl-MinMaxMCS-List-r16: ref
	{
		if ie.Sl_MinMaxMCS_List_r16 != nil {
			if err := ie.Sl_MinMaxMCS_List_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-TimeResource-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_TimeResource_r16 != nil}); err != nil {
				return err
			}

			if ie.Sl_TimeResource_r16 != nil {
				if err := ex.EncodeBitString(*ie.Sl_TimeResource_r16, sLResourcePoolR16SlTimeResourceR16Constraints); err != nil {
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
					{Name: "sl-PBPS-CPS-Config-r17", Optional: true},
					{Name: "sl-InterUE-CoordinationConfig-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PBPS_CPS_Config_r17 != nil, ie.Sl_InterUE_CoordinationConfig_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_PBPS_CPS_Config_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLResourcePoolR16ExtSlPBPSCPSConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PBPS_CPS_Config_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_PBPS_CPS_Config_r17).Choice {
				case SL_ResourcePool_r16_Ext_Sl_PBPS_CPS_Config_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_ResourcePool_r16_Ext_Sl_PBPS_CPS_Config_r17_Setup:
					if err := (*ie.Sl_PBPS_CPS_Config_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_InterUE_CoordinationConfig_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLResourcePoolR16ExtSlInterUECoordinationConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_InterUE_CoordinationConfig_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_InterUE_CoordinationConfig_r17).Choice {
				case SL_ResourcePool_r16_Ext_Sl_InterUE_CoordinationConfig_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_ResourcePool_r16_Ext_Sl_InterUE_CoordinationConfig_r17_Setup:
					if err := (*ie.Sl_InterUE_CoordinationConfig_r17).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-InitiateCOT-List-r18", Optional: true},
					{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-InitiateCOT-Default-r18", Optional: true},
					{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-WithinCOT-List-r18", Optional: true},
					{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-WithinCOT-Default-r18", Optional: true},
					{Name: "sl-Type1-LBT-BlockingOption1-r18", Optional: true},
					{Name: "sl-Type1-LBT-BlockingOption2-r18", Optional: true},
					{Name: "sl-NumInterlacePerSubchannel-r18", Optional: true},
					{Name: "sl-NumReferencePRBs-OfInterlace-r18", Optional: true},
					{Name: "sl-TransmissionStructureForPSFCH-r18", Optional: true},
					{Name: "sl-NumDedicatedPRBs-ForPSFCH-r18", Optional: true},
					{Name: "sl-NumPSFCH-Occasions-r18", Optional: true},
					{Name: "sl-PSFCH-CommonInterlaceIndex-r18", Optional: true},
					{Name: "sl-CPE-StartingPositionPSFCH-r18", Optional: true},
					{Name: "sl-NumRefSymbolLength-r18", Optional: true},
					{Name: "sl-PSFCH-RB-SetList-r18", Optional: true},
					{Name: "sl-IUC-RB-SetList-r18", Optional: true},
					{Name: "sl-PSFCH-PowerOffset-r18", Optional: true},
					{Name: "sl-RBSetIndexOfResourcePool-r18", Optional: true},
					{Name: "sl-A2X-Service-r18", Optional: true},
					{Name: "sl-PRS-ResourcesSharedSL-PRS-RP-r18", Optional: true},
					{Name: "numSym-SL-PRS-2ndStageSCI-r18", Optional: true},
					{Name: "sl-SCI-basedSL-PRS-TxTriggerSCI2-D-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18 != nil, ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_Default_r18 != nil, ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18 != nil, ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_Default_r18 != nil, ie.Sl_Type1_LBT_BlockingOption1_r18 != nil, ie.Sl_Type1_LBT_BlockingOption2_r18 != nil, ie.Sl_NumInterlacePerSubchannel_r18 != nil, ie.Sl_NumReferencePRBs_OfInterlace_r18 != nil, ie.Sl_TransmissionStructureForPSFCH_r18 != nil, ie.Sl_NumDedicatedPRBs_ForPSFCH_r18 != nil, ie.Sl_NumPSFCH_Occasions_r18 != nil, ie.Sl_PSFCH_CommonInterlaceIndex_r18 != nil, ie.Sl_CPE_StartingPositionPSFCH_r18 != nil, ie.Sl_NumRefSymbolLength_r18 != nil, ie.Sl_PSFCH_RB_SetList_r18 != nil, ie.Sl_IUC_RB_SetList_r18 != nil, ie.Sl_PSFCH_PowerOffset_r18 != nil, ie.Sl_RBSetIndexOfResourcePool_r18 != nil, ie.Sl_A2X_Service_r18 != nil, ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18 != nil, ie.NumSym_SL_PRS_2ndStageSCI_r18 != nil, ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI2_D_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLResourcePoolR16ExtSlCPEStartingPositionsPSCCHPSSCHInitiateCOTListR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18).Choice {
				case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18_Setup:
					if err := (*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_Default_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_Default_r18, sLResourcePoolR16SlCPEStartingPositionsPSCCHPSSCHInitiateCOTDefaultR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLResourcePoolR16ExtSlCPEStartingPositionsPSCCHPSSCHWithinCOTListR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18).Choice {
				case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18_Setup:
					if err := (*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_Default_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_Default_r18, sLResourcePoolR16SlCPEStartingPositionsPSCCHPSSCHWithinCOTDefaultR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_Type1_LBT_BlockingOption1_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_Type1_LBT_BlockingOption1_r18, sLResourcePoolR16ExtSlType1LBTBlockingOption1R18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_Type1_LBT_BlockingOption2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_Type1_LBT_BlockingOption2_r18, sLResourcePoolR16ExtSlType1LBTBlockingOption2R18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_NumInterlacePerSubchannel_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_NumInterlacePerSubchannel_r18, sLResourcePoolR16ExtSlNumInterlacePerSubchannelR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_NumReferencePRBs_OfInterlace_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_NumReferencePRBs_OfInterlace_r18, sLResourcePoolR16ExtSlNumReferencePRBsOfInterlaceR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_TransmissionStructureForPSFCH_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_TransmissionStructureForPSFCH_r18, sLResourcePoolR16ExtSlTransmissionStructureForPSFCHR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_NumDedicatedPRBs_ForPSFCH_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_NumDedicatedPRBs_ForPSFCH_r18, sLResourcePoolR16ExtSlNumDedicatedPRBsForPSFCHR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_NumPSFCH_Occasions_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_NumPSFCH_Occasions_r18, sLResourcePoolR16ExtSlNumPSFCHOccasionsR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PSFCH_CommonInterlaceIndex_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_PSFCH_CommonInterlaceIndex_r18, sLResourcePoolR16SlPSFCHCommonInterlaceIndexR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_CPE_StartingPositionPSFCH_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_CPE_StartingPositionPSFCH_r18, sLResourcePoolR16SlCPEStartingPositionPSFCHR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_NumRefSymbolLength_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_NumRefSymbolLength_r18, sLResourcePoolR16ExtSlNumRefSymbolLengthR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PSFCH_RB_SetList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLResourcePoolR16ExtSlPSFCHRBSetListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_PSFCH_RB_SetList_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_PSFCH_RB_SetList_r18 {
					if err := ex.EncodeBitString(ie.Sl_PSFCH_RB_SetList_r18[i], per.SizeRange(10, 275)); err != nil {
						return err
					}
				}
			}

			if ie.Sl_IUC_RB_SetList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLResourcePoolR16ExtSlIUCRBSetListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_IUC_RB_SetList_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_IUC_RB_SetList_r18 {
					if err := ex.EncodeBitString(ie.Sl_IUC_RB_SetList_r18[i], per.SizeRange(10, 275)); err != nil {
						return err
					}
				}
			}

			if ie.Sl_PSFCH_PowerOffset_r18 != nil {
				if err := ex.EncodeInteger(*ie.Sl_PSFCH_PowerOffset_r18, sLResourcePoolR16SlPSFCHPowerOffsetR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_RBSetIndexOfResourcePool_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLResourcePoolR16ExtSlRBSetIndexOfResourcePoolR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_RBSetIndexOfResourcePool_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_RBSetIndexOfResourcePool_r18 {
					if err := ex.EncodeInteger(ie.Sl_RBSetIndexOfResourcePool_r18[i], per.Constrained(0, 4)); err != nil {
						return err
					}
				}
			}

			if ie.Sl_A2X_Service_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_A2X_Service_r18, sLResourcePoolR16ExtSlA2XServiceR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLResourcePoolR16ExtSlPRSResourcesSharedSLPRSRPR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18))); err != nil {
					return err
				}
				for i := range ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18 {
					if err := ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.NumSym_SL_PRS_2ndStageSCI_r18 != nil {
				if err := ex.EncodeInteger(*ie.NumSym_SL_PRS_2ndStageSCI_r18, sLResourcePoolR16NumSymSLPRS2ndStageSCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI2_D_r18 != nil {
				if err := ex.EncodeBoolean(*ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI2_D_r18); err != nil {
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

func (ie *SL_ResourcePool_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLResourcePoolR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PSCCH-Config-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_PSCCH_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_PSCCH_Config_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sL_ResourcePool_r16SlPSCCHConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PSCCH_Config_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_ResourcePool_r16_Sl_PSCCH_Config_r16_Release:
				(*ie.Sl_PSCCH_Config_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Sl_PSCCH_Config_r16_Setup:
				(*ie.Sl_PSCCH_Config_r16).Setup = new(SL_PSCCH_Config_r16)
				if err := (*ie.Sl_PSCCH_Config_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PSSCH-Config-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_PSSCH_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_PSSCH_Config_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sL_ResourcePool_r16SlPSSCHConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PSSCH_Config_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_ResourcePool_r16_Sl_PSSCH_Config_r16_Release:
				(*ie.Sl_PSSCH_Config_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Sl_PSSCH_Config_r16_Setup:
				(*ie.Sl_PSSCH_Config_r16).Setup = new(SL_PSSCH_Config_r16)
				if err := (*ie.Sl_PSSCH_Config_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-PSFCH-Config-r16: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_PSFCH_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_PSFCH_Config_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sL_ResourcePool_r16SlPSFCHConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PSFCH_Config_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_ResourcePool_r16_Sl_PSFCH_Config_r16_Release:
				(*ie.Sl_PSFCH_Config_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Sl_PSFCH_Config_r16_Setup:
				(*ie.Sl_PSFCH_Config_r16).Setup = new(SL_PSFCH_Config_r16)
				if err := (*ie.Sl_PSFCH_Config_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-SyncAllowed-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_SyncAllowed_r16 = new(SL_SyncAllowed_r16)
			if err := ie.Sl_SyncAllowed_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. sl-SubchannelSize-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sLResourcePoolR16SlSubchannelSizeR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SubchannelSize_r16 = &idx
		}
	}

	// 8. dummy: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(sLResourcePoolR16DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &v
		}
	}

	// 9. sl-StartRB-Subchannel-r16: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(sLResourcePoolR16SlStartRBSubchannelR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_StartRB_Subchannel_r16 = &v
		}
	}

	// 10. sl-NumSubchannel-r16: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(sLResourcePoolR16SlNumSubchannelR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumSubchannel_r16 = &v
		}
	}

	// 11. sl-Additional-MCS-Table-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(sLResourcePoolR16SlAdditionalMCSTableR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Additional_MCS_Table_r16 = &idx
		}
	}

	// 12. sl-ThreshS-RSSI-CBR-r16: integer
	{
		if seq.IsComponentPresent(9) {
			v, err := d.DecodeInteger(sLResourcePoolR16SlThreshSRSSICBRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ThreshS_RSSI_CBR_r16 = &v
		}
	}

	// 13. sl-TimeWindowSizeCBR-r16: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sLResourcePoolR16SlTimeWindowSizeCBRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeWindowSizeCBR_r16 = &idx
		}
	}

	// 14. sl-TimeWindowSizeCR-r16: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(sLResourcePoolR16SlTimeWindowSizeCRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeWindowSizeCR_r16 = &idx
		}
	}

	// 15. sl-PTRS-Config-r16: ref
	{
		if seq.IsComponentPresent(12) {
			ie.Sl_PTRS_Config_r16 = new(SL_PTRS_Config_r16)
			if err := ie.Sl_PTRS_Config_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 16. sl-UE-SelectedConfigRP-r16: ref
	{
		if seq.IsComponentPresent(13) {
			ie.Sl_UE_SelectedConfigRP_r16 = new(SL_UE_SelectedConfigRP_r16)
			if err := ie.Sl_UE_SelectedConfigRP_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 17. sl-RxParametersNcell-r16: sequence
	{
		if seq.IsComponentPresent(14) {
			ie.Sl_RxParametersNcell_r16 = &struct {
				Sl_TDD_Configuration_r16 *TDD_UL_DL_ConfigCommon
				Sl_SyncConfigIndex_r16   int64
			}{}
			c := ie.Sl_RxParametersNcell_r16
			sLResourcePoolR16SlRxParametersNcellR16Seq := d.NewSequenceDecoder(sLResourcePoolR16SlRxParametersNcellR16Constraints)
			if err := sLResourcePoolR16SlRxParametersNcellR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if sLResourcePoolR16SlRxParametersNcellR16Seq.IsComponentPresent(0) {
				c.Sl_TDD_Configuration_r16 = new(TDD_UL_DL_ConfigCommon)
				if err := (*c.Sl_TDD_Configuration_r16).Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.Sl_SyncConfigIndex_r16 = v
			}
		}
	}

	// 18. sl-ZoneConfigMCR-List-r16: sequence-of
	{
		if seq.IsComponentPresent(15) {
			seqOf := d.NewSequenceOfDecoder(sLResourcePoolR16SlZoneConfigMCRListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ZoneConfigMCR_List_r16 = make([]SL_ZoneConfigMCR_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ZoneConfigMCR_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 19. sl-FilterCoefficient-r16: ref
	{
		if seq.IsComponentPresent(16) {
			ie.Sl_FilterCoefficient_r16 = new(FilterCoefficient)
			if err := ie.Sl_FilterCoefficient_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 20. sl-RB-Number-r16: integer
	{
		if seq.IsComponentPresent(17) {
			v, err := d.DecodeInteger(sLResourcePoolR16SlRBNumberR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_RB_Number_r16 = &v
		}
	}

	// 21. sl-PreemptionEnable-r16: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(sLResourcePoolR16SlPreemptionEnableR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PreemptionEnable_r16 = &idx
		}
	}

	// 22. sl-PriorityThreshold-UL-URLLC-r16: integer
	{
		if seq.IsComponentPresent(19) {
			v, err := d.DecodeInteger(sLResourcePoolR16SlPriorityThresholdULURLLCR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityThreshold_UL_URLLC_r16 = &v
		}
	}

	// 23. sl-PriorityThreshold-r16: integer
	{
		if seq.IsComponentPresent(20) {
			v, err := d.DecodeInteger(sLResourcePoolR16SlPriorityThresholdR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityThreshold_r16 = &v
		}
	}

	// 24. sl-X-Overhead-r16: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(sLResourcePoolR16SlXOverheadR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_X_Overhead_r16 = &idx
		}
	}

	// 25. sl-PowerControl-r16: ref
	{
		if seq.IsComponentPresent(22) {
			ie.Sl_PowerControl_r16 = new(SL_PowerControl_r16)
			if err := ie.Sl_PowerControl_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 26. sl-TxPercentageList-r16: ref
	{
		if seq.IsComponentPresent(23) {
			ie.Sl_TxPercentageList_r16 = new(SL_TxPercentageList_r16)
			if err := ie.Sl_TxPercentageList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 27. sl-MinMaxMCS-List-r16: ref
	{
		if seq.IsComponentPresent(24) {
			ie.Sl_MinMaxMCS_List_r16 = new(SL_MinMaxMCS_List_r16)
			if err := ie.Sl_MinMaxMCS_List_r16.Decode(d); err != nil {
				return err
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
				{Name: "sl-TimeResource-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBitString(sLResourcePoolR16SlTimeResourceR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeResource_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-PBPS-CPS-Config-r17", Optional: true},
				{Name: "sl-InterUE-CoordinationConfig-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_PBPS_CPS_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_PBPS_CPS_Config_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(sLResourcePoolR16ExtSlPBPSCPSConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PBPS_CPS_Config_r17).Choice = int(index)
			switch index {
			case SL_ResourcePool_r16_Ext_Sl_PBPS_CPS_Config_r17_Release:
				(*ie.Sl_PBPS_CPS_Config_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Ext_Sl_PBPS_CPS_Config_r17_Setup:
				(*ie.Sl_PBPS_CPS_Config_r17).Setup = new(SL_PBPS_CPS_Config_r17)
				if err := (*ie.Sl_PBPS_CPS_Config_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_InterUE_CoordinationConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_InterUE_CoordinationConfig_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(sLResourcePoolR16ExtSlInterUECoordinationConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_InterUE_CoordinationConfig_r17).Choice = int(index)
			switch index {
			case SL_ResourcePool_r16_Ext_Sl_InterUE_CoordinationConfig_r17_Release:
				(*ie.Sl_InterUE_CoordinationConfig_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Ext_Sl_InterUE_CoordinationConfig_r17_Setup:
				(*ie.Sl_InterUE_CoordinationConfig_r17).Setup = new(SL_InterUE_CoordinationConfig_r17)
				if err := (*ie.Sl_InterUE_CoordinationConfig_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-InitiateCOT-List-r18", Optional: true},
				{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-InitiateCOT-Default-r18", Optional: true},
				{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-WithinCOT-List-r18", Optional: true},
				{Name: "sl-CPE-StartingPositionsPSCCH-PSSCH-WithinCOT-Default-r18", Optional: true},
				{Name: "sl-Type1-LBT-BlockingOption1-r18", Optional: true},
				{Name: "sl-Type1-LBT-BlockingOption2-r18", Optional: true},
				{Name: "sl-NumInterlacePerSubchannel-r18", Optional: true},
				{Name: "sl-NumReferencePRBs-OfInterlace-r18", Optional: true},
				{Name: "sl-TransmissionStructureForPSFCH-r18", Optional: true},
				{Name: "sl-NumDedicatedPRBs-ForPSFCH-r18", Optional: true},
				{Name: "sl-NumPSFCH-Occasions-r18", Optional: true},
				{Name: "sl-PSFCH-CommonInterlaceIndex-r18", Optional: true},
				{Name: "sl-CPE-StartingPositionPSFCH-r18", Optional: true},
				{Name: "sl-NumRefSymbolLength-r18", Optional: true},
				{Name: "sl-PSFCH-RB-SetList-r18", Optional: true},
				{Name: "sl-IUC-RB-SetList-r18", Optional: true},
				{Name: "sl-PSFCH-PowerOffset-r18", Optional: true},
				{Name: "sl-RBSetIndexOfResourcePool-r18", Optional: true},
				{Name: "sl-A2X-Service-r18", Optional: true},
				{Name: "sl-PRS-ResourcesSharedSL-PRS-RP-r18", Optional: true},
				{Name: "numSym-SL-PRS-2ndStageSCI-r18", Optional: true},
				{Name: "sl-SCI-basedSL-PRS-TxTriggerSCI2-D-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sLResourcePoolR16ExtSlCPEStartingPositionsPSCCHPSSCHInitiateCOTListR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18).Choice = int(index)
			switch index {
			case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18_Release:
				(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18_Setup:
				(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18).Setup = new(SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18)
				if err := (*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_List_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(sLResourcePoolR16SlCPEStartingPositionsPSCCHPSSCHInitiateCOTDefaultR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_InitiateCOT_Default_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sLResourcePoolR16ExtSlCPEStartingPositionsPSCCHPSSCHWithinCOTListR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18).Choice = int(index)
			switch index {
			case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18_Release:
				(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_ResourcePool_r16_Ext_Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18_Setup:
				(*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18).Setup = new(SL_CPE_StartingPositionsPSCCH_PSSCH_List_r18)
				if err := (*ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_List_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(sLResourcePoolR16SlCPEStartingPositionsPSCCHPSSCHWithinCOTDefaultR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CPE_StartingPositionsPSCCH_PSSCH_WithinCOT_Default_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlType1LBTBlockingOption1R18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Type1_LBT_BlockingOption1_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlType1LBTBlockingOption2R18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Type1_LBT_BlockingOption2_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlNumInterlacePerSubchannelR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumInterlacePerSubchannel_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlNumReferencePRBsOfInterlaceR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumReferencePRBs_OfInterlace_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlTransmissionStructureForPSFCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TransmissionStructureForPSFCH_r18 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlNumDedicatedPRBsForPSFCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumDedicatedPRBs_ForPSFCH_r18 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlNumPSFCHOccasionsR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumPSFCH_Occasions_r18 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeInteger(sLResourcePoolR16SlPSFCHCommonInterlaceIndexR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_CommonInterlaceIndex_r18 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeInteger(sLResourcePoolR16SlCPEStartingPositionPSFCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CPE_StartingPositionPSFCH_r18 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlNumRefSymbolLengthR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumRefSymbolLength_r18 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			seqOf := dx.NewSequenceOfDecoder(sLResourcePoolR16ExtSlPSFCHRBSetListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_RB_SetList_r18 = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeBitString(per.SizeRange(10, 275))
				if err != nil {
					return err
				}
				ie.Sl_PSFCH_RB_SetList_r18[i] = v
			}
		}

		if groupSeq.IsComponentPresent(15) {
			seqOf := dx.NewSequenceOfDecoder(sLResourcePoolR16ExtSlIUCRBSetListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_IUC_RB_SetList_r18 = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeBitString(per.SizeRange(10, 275))
				if err != nil {
					return err
				}
				ie.Sl_IUC_RB_SetList_r18[i] = v
			}
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeInteger(sLResourcePoolR16SlPSFCHPowerOffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_PowerOffset_r18 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			seqOf := dx.NewSequenceOfDecoder(sLResourcePoolR16ExtSlRBSetIndexOfResourcePoolR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RBSetIndexOfResourcePool_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				ie.Sl_RBSetIndexOfResourcePool_r18[i] = v
			}
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(sLResourcePoolR16ExtSlA2XServiceR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_A2X_Service_r18 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			seqOf := dx.NewSequenceOfDecoder(sLResourcePoolR16ExtSlPRSResourcesSharedSLPRSRPR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18 = make([]SL_PRS_ResourceSharedSL_PRS_RP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PRS_ResourcesSharedSL_PRS_RP_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeInteger(sLResourcePoolR16NumSymSLPRS2ndStageSCIR18Constraints)
			if err != nil {
				return err
			}
			ie.NumSym_SL_PRS_2ndStageSCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI2_D_r18 = &v
		}
	}

	return nil
}
