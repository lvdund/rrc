// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDCCH-ConfigCommon (line 11044).

var pDCCHConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "controlResourceSetZero", Optional: true},
		{Name: "commonControlResourceSet", Optional: true},
		{Name: "searchSpaceZero", Optional: true},
		{Name: "commonSearchSpaceList", Optional: true},
		{Name: "searchSpaceSIB1", Optional: true},
		{Name: "searchSpaceOtherSystemInformation", Optional: true},
		{Name: "pagingSearchSpace", Optional: true},
		{Name: "ra-SearchSpace", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
	},
}

var pDCCHConfigCommonCommonSearchSpaceListConstraints = per.SizeRange(1, 4)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS15KHZoneT"},
		{Name: "sCS30KHZoneT-SCS15KHZhalfT"},
		{Name: "sCS60KHZoneT-SCS30KHZhalfT-SCS15KHZquarterT"},
		{Name: "sCS120KHZoneT-SCS60KHZhalfT-SCS30KHZquarterT-SCS15KHZoneEighthT"},
		{Name: "sCS120KHZhalfT-SCS60KHZquarterT-SCS30KHZoneEighthT-SCS15KHZoneSixteenthT"},
		{Name: "sCS120KHZquarterT-SCS60KHZoneEighthT-SCS30KHZoneSixteenthT"},
		{Name: "sCS120KHZoneEighthT-SCS60KHZoneSixteenthT"},
		{Name: "sCS120KHZoneSixteenthT"},
	},
}

const (
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS15KHZoneT                                                             = 0
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS30KHZoneT_SCS15KHZhalfT                                               = 1
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              = 2
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          = 3
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = 4
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT               = 5
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                                = 6
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneSixteenthT                                                   = 7
)

var pDCCHConfigCommonExtCommonSearchSpaceListExtR16Constraints = per.SizeRange(1, 4)

var pDCCHConfigCommonExtSdtSearchSpaceR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "newSearchSpace"},
		{Name: "existingSearchSpace"},
	},
}

const (
	PDCCH_ConfigCommon_Ext_Sdt_SearchSpace_r17_NewSearchSpace      = 0
	PDCCH_ConfigCommon_Ext_Sdt_SearchSpace_r17_ExistingSearchSpace = 1
)

var pDCCHConfigCommonExtCommonSearchSpaceListExt2R17Constraints = per.SizeRange(1, 4)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS480KHZoneEighthT"},
		{Name: "sCS480KHZoneSixteenthT"},
	},
}

const (
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneEighthT    = 0
	PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneSixteenthT = 1
)

const (
	PDCCH_ConfigCommon_Ext_FollowUnifiedTCI_State_v1720_Enabled = 0
)

var pDCCHConfigCommonExtFollowUnifiedTCIStateV1720Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_ConfigCommon_Ext_FollowUnifiedTCI_State_v1720_Enabled},
}

const (
	PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_First  = 0
	PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_Second = 1
	PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_Both   = 2
	PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_None   = 3
)

var pDCCHConfigCommonExtApplyIndicatedTCIStateR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_First, PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_Second, PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_Both, PDCCH_ConfigCommon_Ext_ApplyIndicatedTCI_State_r18_None},
}

var pDCCHConfigCommonExtCommonSearchSpaceListExtR18Constraints = per.SizeRange(1, 4)

var pDCCHConfigCommonExtCommonSearchSpaceListExtR19Constraints = per.SizeRange(1, 4)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS15KHZoneT"},
		{Name: "sCS30KHZoneT-SCS15KHZhalfT"},
		{Name: "sCS60KHZoneT-SCS30KHZhalfT-SCS15KHZquarterT"},
		{Name: "sCS120KHZoneT-SCS60KHZhalfT-SCS30KHZquarterT-SCS15KHZoneEighthT"},
		{Name: "sCS120KHZhalfT-SCS60KHZquarterT-SCS30KHZoneEighthT-SCS15KHZoneSixteenthT"},
		{Name: "sCS480KHZoneT-SCS120KHZquarterT-SCS60KHZoneEighthT-SCS30KHZoneSixteenthT-SCS15KHZoneThirtySecondT"},
		{Name: "sCS480KHZhalfT-SCS120KHZoneEighthT-SCS60KHZoneSixteenthT-SCS30KHZoneThirtySecondT"},
		{Name: "sCS480KHZquarterT-SCS120KHZoneSixteenthT-SCS60KHZoneThirtySecondT"},
		{Name: "sCS480KHZoneEighthT-sCS120KHZoneThirtySecondT"},
		{Name: "sCS480KHZoneSixteenthT"},
		{Name: "sCS480KHZoneThirtySecondT"},
	},
}

const (
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS15KHZoneT                                                                                      = 0
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS30KHZoneT_SCS15KHZhalfT                                                                        = 1
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       = 2
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   = 3
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          = 4
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT = 5
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 = 6
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 = 7
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     = 8
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneSixteenthT                                                                            = 9
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneThirtySecondT                                                                         = 10
)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS15KHZoneT"},
		{Name: "sCS30KHZoneT-SCS15KHZhalfT"},
		{Name: "sCS60KHZoneT-SCS30KHZhalfT-SCS15KHZquarterT"},
		{Name: "sCS120KHZoneT-SCS60KHZhalfT-SCS30KHZquarterT-SCS15KHZoneEighthT"},
		{Name: "sCS120KHZhalfT-SCS60KHZquarterT-SCS30KHZoneEighthT-SCS15KHZoneSixteenthT"},
		{Name: "sCS480KHZoneT-SCS120KHZquarterT-SCS60KHZoneEighthT-SCS30KHZoneSixteenthT-SCS15KHZoneThirtySecondT"},
		{Name: "sCS480KHZhalfT-SCS120KHZoneEighthT-SCS60KHZoneSixteenthT-SCS30KHZoneThirtySecondT"},
		{Name: "sCS480KHZquarterT-SCS120KHZoneSixteenthT-SCS60KHZoneThirtySecondT"},
		{Name: "sCS480KHZoneEighthT-sCS120KHZoneThirtySecondT"},
		{Name: "sCS480KHZoneSixteenthT"},
		{Name: "sCS480KHZoneThirtySecondT"},
	},
}

const (
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS15KHZoneT                                                                                      = 0
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS30KHZoneT_SCS15KHZhalfT                                                                        = 1
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       = 2
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   = 3
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          = 4
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT = 5
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 = 6
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 = 7
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     = 8
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneSixteenthT                                                                            = 9
	PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneThirtySecondT                                                                         = 10
)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS15KHZoneTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS30KHZoneTSCS15KHZhalfTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneEighthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS15KHZoneT"},
		{Name: "sCS30KHZoneT-SCS15KHZhalfT"},
		{Name: "sCS60KHZoneT-SCS30KHZhalfT-SCS15KHZquarterT"},
		{Name: "sCS120KHZoneT-SCS60KHZhalfT-SCS30KHZquarterT-SCS15KHZoneEighthT"},
		{Name: "sCS120KHZhalfT-SCS60KHZquarterT-SCS30KHZoneEighthT-SCS15KHZoneSixteenthT"},
		{Name: "sCS480KHZoneT-SCS120KHZquarterT-SCS60KHZoneEighthT-SCS30KHZoneSixteenthT"},
		{Name: "sCS480KHZhalfT-SCS120KHZoneEighthT-SCS60KHZoneSixteenthT"},
		{Name: "sCS480KHZquarterT-SCS120KHZoneSixteenthT"},
		{Name: "sCS480KHZoneEighthT"},
		{Name: "sCS480KHZoneSixteenthT"},
	},
}

const (
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS15KHZoneT                                                             = 0
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS30KHZoneT_SCS15KHZhalfT                                               = 1
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              = 2
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          = 3
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = 4
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT = 5
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                 = 6
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZquarterT_SCS120KHZoneSixteenthT                                 = 7
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneEighthT                                                      = 8
	PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneSixteenthT                                                   = 9
)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS15KHZoneTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS30KHZoneTSCS15KHZhalfTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZquarterTSCS120KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneEighthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r17)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS15KHZoneTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS30KHZoneTSCS15KHZhalfTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS15KHZoneTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS30KHZoneTSCS15KHZhalfTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

var pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxPEI_PerPF_r19)

type PDCCH_ConfigCommon struct {
	ControlResourceSetZero            *ControlResourceSetZero
	CommonControlResourceSet          *ControlResourceSet
	SearchSpaceZero                   *SearchSpaceZero
	CommonSearchSpaceList             []SearchSpace
	SearchSpaceSIB1                   *SearchSpaceId
	SearchSpaceOtherSystemInformation *SearchSpaceId
	PagingSearchSpace                 *SearchSpaceId
	Ra_SearchSpace                    *SearchSpaceId
	FirstPDCCH_MonitoringOccasionOfPO *struct {
		Choice                                                                   int
		SCS15KHZoneT                                                             *[]int64
		SCS30KHZoneT_SCS15KHZhalfT                                               *[]int64
		SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              *[]int64
		SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          *[]int64
		SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT *[]int64
		SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT               *[]int64
		SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                                *[]int64
		SCS120KHZoneSixteenthT                                                   *[]int64
	}
	CommonSearchSpaceListExt_r16 []SearchSpaceExt_r16
	Sdt_SearchSpace_r17          *struct {
		Choice              int
		NewSearchSpace      *SearchSpace
		ExistingSearchSpace *SearchSpaceId
	}
	SearchSpaceMCCH_r17                     *SearchSpaceId
	SearchSpaceMTCH_r17                     *SearchSpaceId
	CommonSearchSpaceListExt2_r17           []SearchSpaceExt_v1700
	FirstPDCCH_MonitoringOccasionOfPO_v1710 *struct {
		Choice                 int
		SCS480KHZoneEighthT    *[]int64
		SCS480KHZoneSixteenthT *[]int64
	}
	Pei_ConfigBWP_r17 *struct {
		Pei_SearchSpace_r17                      SearchSpaceId
		FirstPDCCH_MonitoringOccasionOfPEI_O_r17 struct {
			Choice                                                                   int
			SCS15KHZoneT                                                             *[]int64
			SCS30KHZoneT_SCS15KHZhalfT                                               *[]int64
			SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              *[]int64
			SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          *[]int64
			SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT *[]int64
			SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT *[]int64
			SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                 *[]int64
			SCS480KHZquarterT_SCS120KHZoneSixteenthT                                 *[]int64
			SCS480KHZoneEighthT                                                      *[]int64
			SCS480KHZoneSixteenthT                                                   *[]int64
		}
	}
	FollowUnifiedTCI_State_v1720                     *int64
	ApplyIndicatedTCI_State_r18                      *int64
	CommonSearchSpaceListExt_r18                     []SearchSpaceExt_v1800
	SearchSpaceMulticastMCCH_r18                     *SearchSpaceId
	SearchSpaceMulticastMTCH_r18                     *SearchSpaceId
	CommonSearchSpaceListExt_r19                     []SearchSpaceExt_v1800
	PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 *struct {
		Choice                                                                                            int
		SCS15KHZoneT                                                                                      *[]int64
		SCS30KHZoneT_SCS15KHZhalfT                                                                        *[]int64
		SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       *[]int64
		SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   *[]int64
		SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          *[]int64
		SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT *[]int64
		SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 *[]int64
		SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 *[]int64
		SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     *[]int64
		SCS480KHZoneSixteenthT                                                                            *[]int64
		SCS480KHZoneThirtySecondT                                                                         *[]int64
	}
	PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19 *struct {
		Choice                                                                                            int
		SCS15KHZoneT                                                                                      *[]int64
		SCS30KHZoneT_SCS15KHZhalfT                                                                        *[]int64
		SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       *[]int64
		SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   *[]int64
		SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          *[]int64
		SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT *[]int64
		SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 *[]int64
		SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 *[]int64
		SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     *[]int64
		SCS480KHZoneSixteenthT                                                                            *[]int64
		SCS480KHZoneThirtySecondT                                                                         *[]int64
	}
	PagingAdaptPagingSearchSpace *SearchSpaceId
}

func (ie *PDCCH_ConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHConfigCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.FirstPDCCH_MonitoringOccasionOfPO != nil
	hasExtGroup1 := ie.CommonSearchSpaceListExt_r16 != nil
	hasExtGroup2 := ie.Sdt_SearchSpace_r17 != nil || ie.SearchSpaceMCCH_r17 != nil || ie.SearchSpaceMTCH_r17 != nil || ie.CommonSearchSpaceListExt2_r17 != nil || ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 != nil || ie.Pei_ConfigBWP_r17 != nil
	hasExtGroup3 := ie.FollowUnifiedTCI_State_v1720 != nil
	hasExtGroup4 := ie.ApplyIndicatedTCI_State_r18 != nil || ie.CommonSearchSpaceListExt_r18 != nil || ie.SearchSpaceMulticastMCCH_r18 != nil || ie.SearchSpaceMulticastMTCH_r18 != nil
	hasExtGroup5 := ie.CommonSearchSpaceListExt_r19 != nil || ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 != nil || ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19 != nil || ie.PagingAdaptPagingSearchSpace != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ControlResourceSetZero != nil, ie.CommonControlResourceSet != nil, ie.SearchSpaceZero != nil, ie.CommonSearchSpaceList != nil, ie.SearchSpaceSIB1 != nil, ie.SearchSpaceOtherSystemInformation != nil, ie.PagingSearchSpace != nil, ie.Ra_SearchSpace != nil}); err != nil {
		return err
	}

	// 3. controlResourceSetZero: ref
	{
		if ie.ControlResourceSetZero != nil {
			if err := ie.ControlResourceSetZero.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. commonControlResourceSet: ref
	{
		if ie.CommonControlResourceSet != nil {
			if err := ie.CommonControlResourceSet.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. searchSpaceZero: ref
	{
		if ie.SearchSpaceZero != nil {
			if err := ie.SearchSpaceZero.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. commonSearchSpaceList: sequence-of
	{
		if ie.CommonSearchSpaceList != nil {
			seqOf := e.NewSequenceOfEncoder(pDCCHConfigCommonCommonSearchSpaceListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.CommonSearchSpaceList))); err != nil {
				return err
			}
			for i := range ie.CommonSearchSpaceList {
				if err := ie.CommonSearchSpaceList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. searchSpaceSIB1: ref
	{
		if ie.SearchSpaceSIB1 != nil {
			if err := ie.SearchSpaceSIB1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. searchSpaceOtherSystemInformation: ref
	{
		if ie.SearchSpaceOtherSystemInformation != nil {
			if err := ie.SearchSpaceOtherSystemInformation.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. pagingSearchSpace: ref
	{
		if ie.PagingSearchSpace != nil {
			if err := ie.PagingSearchSpace.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. ra-SearchSpace: ref
	{
		if ie.Ra_SearchSpace != nil {
			if err := ie.Ra_SearchSpace.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "firstPDCCH-MonitoringOccasionOfPO", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FirstPDCCH_MonitoringOccasionOfPO != nil}); err != nil {
				return err
			}

			if ie.FirstPDCCH_MonitoringOccasionOfPO != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOConstraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.FirstPDCCH_MonitoringOccasionOfPO).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.FirstPDCCH_MonitoringOccasionOfPO).Choice {
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS15KHZoneT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS15KHZoneTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT)[i], per.Constrained(0, 139)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS30KHZoneT_SCS15KHZhalfT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS30KHZoneTSCS15KHZhalfTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT)[i], per.Constrained(0, 279)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i], per.Constrained(0, 559)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i], per.Constrained(0, 1119)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i], per.Constrained(0, 2239)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)[i], per.Constrained(0, 4479)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneEighthT_SCS60KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)[i], per.Constrained(0, 8959)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneSixteenthT)[i], per.Constrained(0, 17919)); err != nil {
							return err
						}
					}
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
					{Name: "commonSearchSpaceListExt-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CommonSearchSpaceListExt_r16 != nil}); err != nil {
				return err
			}

			if ie.CommonSearchSpaceListExt_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtCommonSearchSpaceListExtR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CommonSearchSpaceListExt_r16))); err != nil {
					return err
				}
				for i := range ie.CommonSearchSpaceListExt_r16 {
					if err := ie.CommonSearchSpaceListExt_r16[i].Encode(ex); err != nil {
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
					{Name: "sdt-SearchSpace-r17", Optional: true},
					{Name: "searchSpaceMCCH-r17", Optional: true},
					{Name: "searchSpaceMTCH-r17", Optional: true},
					{Name: "commonSearchSpaceListExt2-r17", Optional: true},
					{Name: "firstPDCCH-MonitoringOccasionOfPO-v1710", Optional: true},
					{Name: "pei-ConfigBWP-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sdt_SearchSpace_r17 != nil, ie.SearchSpaceMCCH_r17 != nil, ie.SearchSpaceMTCH_r17 != nil, ie.CommonSearchSpaceListExt2_r17 != nil, ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 != nil, ie.Pei_ConfigBWP_r17 != nil}); err != nil {
				return err
			}

			if ie.Sdt_SearchSpace_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCCHConfigCommonExtSdtSearchSpaceR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sdt_SearchSpace_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sdt_SearchSpace_r17).Choice {
				case PDCCH_ConfigCommon_Ext_Sdt_SearchSpace_r17_NewSearchSpace:
					if err := (*ie.Sdt_SearchSpace_r17).NewSearchSpace.Encode(ex); err != nil {
						return err
					}
				case PDCCH_ConfigCommon_Ext_Sdt_SearchSpace_r17_ExistingSearchSpace:
					if err := (*ie.Sdt_SearchSpace_r17).ExistingSearchSpace.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SearchSpaceMCCH_r17 != nil {
				if err := ie.SearchSpaceMCCH_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SearchSpaceMTCH_r17 != nil {
				if err := ie.SearchSpaceMTCH_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CommonSearchSpaceListExt2_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtCommonSearchSpaceListExt2R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CommonSearchSpaceListExt2_r17))); err != nil {
					return err
				}
				for i := range ie.CommonSearchSpaceListExt2_r17 {
					if err := ie.CommonSearchSpaceListExt2_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).Choice {
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneEighthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneEighthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT)[i], per.Constrained(0, 35839)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT)[i], per.Constrained(0, 71679)); err != nil {
							return err
						}
					}
				}
			}

			if ie.Pei_ConfigBWP_r17 != nil {
				c := ie.Pei_ConfigBWP_r17
				if err := c.Pei_SearchSpace_r17.Encode(ex); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.Choice), false, nil); err != nil {
						return err
					}
					switch c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.Choice {
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS15KHZoneT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS15KHZoneTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS15KHZoneT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS15KHZoneT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS15KHZoneT)[i], per.Constrained(0, 139)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS30KHZoneT_SCS15KHZhalfT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS30KHZoneTSCS15KHZhalfTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS30KHZoneT_SCS15KHZhalfT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS30KHZoneT_SCS15KHZhalfT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS30KHZoneT_SCS15KHZhalfT)[i], per.Constrained(0, 279)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i], per.Constrained(0, 559)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i], per.Constrained(0, 1119)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i], per.Constrained(0, 2239)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)[i], per.Constrained(0, 4479)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)[i], per.Constrained(0, 8959)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZquarterT_SCS120KHZoneSixteenthT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZquarterTSCS120KHZoneSixteenthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZquarterT_SCS120KHZoneSixteenthT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZquarterT_SCS120KHZoneSixteenthT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZquarterT_SCS120KHZoneSixteenthT)[i], per.Constrained(0, 17919)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneEighthT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneEighthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneEighthT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneEighthT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneEighthT)[i], per.Constrained(0, 35839)); err != nil {
								return err
							}
						}
					case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneSixteenthT:
						seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneSixteenthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneSixteenthT)))); err != nil {
							return err
						}
						for i := range *c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneSixteenthT {
							if err := ex.EncodeInteger((*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneSixteenthT)[i], per.Constrained(0, 71679)); err != nil {
								return err
							}
						}
					}
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
					{Name: "followUnifiedTCI-State-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FollowUnifiedTCI_State_v1720 != nil}); err != nil {
				return err
			}

			if ie.FollowUnifiedTCI_State_v1720 != nil {
				if err := ex.EncodeEnumerated(*ie.FollowUnifiedTCI_State_v1720, pDCCHConfigCommonExtFollowUnifiedTCIStateV1720Constraints); err != nil {
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
					{Name: "applyIndicatedTCI-State-r18", Optional: true},
					{Name: "commonSearchSpaceListExt-r18", Optional: true},
					{Name: "searchSpaceMulticastMCCH-r18", Optional: true},
					{Name: "searchSpaceMulticastMTCH-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ApplyIndicatedTCI_State_r18 != nil, ie.CommonSearchSpaceListExt_r18 != nil, ie.SearchSpaceMulticastMCCH_r18 != nil, ie.SearchSpaceMulticastMTCH_r18 != nil}); err != nil {
				return err
			}

			if ie.ApplyIndicatedTCI_State_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ApplyIndicatedTCI_State_r18, pDCCHConfigCommonExtApplyIndicatedTCIStateR18Constraints); err != nil {
					return err
				}
			}

			if ie.CommonSearchSpaceListExt_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtCommonSearchSpaceListExtR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CommonSearchSpaceListExt_r18))); err != nil {
					return err
				}
				for i := range ie.CommonSearchSpaceListExt_r18 {
					if err := ie.CommonSearchSpaceListExt_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SearchSpaceMulticastMCCH_r18 != nil {
				if err := ie.SearchSpaceMulticastMCCH_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SearchSpaceMulticastMTCH_r18 != nil {
				if err := ie.SearchSpaceMulticastMTCH_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "commonSearchSpaceListExt-r19", Optional: true},
					{Name: "pagingAdaptFirstPDCCH-MonitoringOccasionOfPO-r19", Optional: true},
					{Name: "pagingAdaptFirstPDCCH-MonitoringOccasionOfPEI-O-r19", Optional: true},
					{Name: "pagingAdaptPagingSearchSpace", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CommonSearchSpaceListExt_r19 != nil, ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 != nil, ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19 != nil, ie.PagingAdaptPagingSearchSpace != nil}); err != nil {
				return err
			}

			if ie.CommonSearchSpaceListExt_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtCommonSearchSpaceListExtR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CommonSearchSpaceListExt_r19))); err != nil {
					return err
				}
				for i := range ie.CommonSearchSpaceListExt_r19 {
					if err := ie.CommonSearchSpaceListExt_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).Choice {
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS15KHZoneT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS15KHZoneTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT)[i], per.Constrained(0, 139)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS30KHZoneT_SCS15KHZhalfT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS30KHZoneTSCS15KHZhalfTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT)[i], per.Constrained(0, 279)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i], per.Constrained(0, 559)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i], per.Constrained(0, 1119)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i], per.Constrained(0, 2239)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)[i], per.Constrained(0, 4479)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)[i], per.Constrained(0, 8959)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)[i], per.Constrained(0, 17919)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)[i], per.Constrained(0, 35839)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT)[i], per.Constrained(0, 71679)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT)[i], per.Constrained(0, 143359)); err != nil {
							return err
						}
					}
				}
			}

			if ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).Choice {
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS15KHZoneT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS15KHZoneTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS15KHZoneT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS15KHZoneT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS15KHZoneT)[i], per.Constrained(0, 139)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS30KHZoneT_SCS15KHZhalfT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS30KHZoneTSCS15KHZhalfTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS30KHZoneT_SCS15KHZhalfT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS30KHZoneT_SCS15KHZhalfT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS30KHZoneT_SCS15KHZhalfT)[i], per.Constrained(0, 279)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i], per.Constrained(0, 559)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i], per.Constrained(0, 1119)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i], per.Constrained(0, 2239)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)[i], per.Constrained(0, 4479)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)[i], per.Constrained(0, 8959)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)[i], per.Constrained(0, 17919)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)[i], per.Constrained(0, 35839)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneSixteenthT)[i], per.Constrained(0, 71679)); err != nil {
							return err
						}
					}
				case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneThirtySecondT:
					seqOf := ex.NewSequenceOfEncoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneThirtySecondTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneThirtySecondT)))); err != nil {
						return err
					}
					for i := range *(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneThirtySecondT {
						if err := ex.EncodeInteger((*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneThirtySecondT)[i], per.Constrained(0, 143359)); err != nil {
							return err
						}
					}
				}
			}

			if ie.PagingAdaptPagingSearchSpace != nil {
				if err := ie.PagingAdaptPagingSearchSpace.Encode(ex); err != nil {
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

func (ie *PDCCH_ConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. controlResourceSetZero: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ControlResourceSetZero = new(ControlResourceSetZero)
			if err := ie.ControlResourceSetZero.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. commonControlResourceSet: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CommonControlResourceSet = new(ControlResourceSet)
			if err := ie.CommonControlResourceSet.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. searchSpaceZero: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SearchSpaceZero = new(SearchSpaceZero)
			if err := ie.SearchSpaceZero.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. commonSearchSpaceList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(pDCCHConfigCommonCommonSearchSpaceListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CommonSearchSpaceList = make([]SearchSpace, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CommonSearchSpaceList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. searchSpaceSIB1: ref
	{
		if seq.IsComponentPresent(4) {
			ie.SearchSpaceSIB1 = new(SearchSpaceId)
			if err := ie.SearchSpaceSIB1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. searchSpaceOtherSystemInformation: ref
	{
		if seq.IsComponentPresent(5) {
			ie.SearchSpaceOtherSystemInformation = new(SearchSpaceId)
			if err := ie.SearchSpaceOtherSystemInformation.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. pagingSearchSpace: ref
	{
		if seq.IsComponentPresent(6) {
			ie.PagingSearchSpace = new(SearchSpaceId)
			if err := ie.PagingSearchSpace.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. ra-SearchSpace: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Ra_SearchSpace = new(SearchSpaceId)
			if err := ie.Ra_SearchSpace.Decode(d); err != nil {
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
				{Name: "firstPDCCH-MonitoringOccasionOfPO", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.FirstPDCCH_MonitoringOccasionOfPO = &struct {
				Choice                                                                   int
				SCS15KHZoneT                                                             *[]int64
				SCS30KHZoneT_SCS15KHZhalfT                                               *[]int64
				SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              *[]int64
				SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          *[]int64
				SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT *[]int64
				SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT               *[]int64
				SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                                *[]int64
				SCS120KHZoneSixteenthT                                                   *[]int64
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.FirstPDCCH_MonitoringOccasionOfPO).Choice = int(index)
			switch index {
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS15KHZoneT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS15KHZoneTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 139))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS30KHZoneT_SCS15KHZhalfT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS30KHZoneTSCS15KHZhalfTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 279))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 559))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 1119))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 2239))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 4479))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneEighthT_SCS60KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneEighthT_SCS60KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 8959))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 17919))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneSixteenthT)[i] = v
				}
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "commonSearchSpaceListExt-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtCommonSearchSpaceListExtR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CommonSearchSpaceListExt_r16 = make([]SearchSpaceExt_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CommonSearchSpaceListExt_r16[i].Decode(dx); err != nil {
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
				{Name: "sdt-SearchSpace-r17", Optional: true},
				{Name: "searchSpaceMCCH-r17", Optional: true},
				{Name: "searchSpaceMTCH-r17", Optional: true},
				{Name: "commonSearchSpaceListExt2-r17", Optional: true},
				{Name: "firstPDCCH-MonitoringOccasionOfPO-v1710", Optional: true},
				{Name: "pei-ConfigBWP-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sdt_SearchSpace_r17 = &struct {
				Choice              int
				NewSearchSpace      *SearchSpace
				ExistingSearchSpace *SearchSpaceId
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCCHConfigCommonExtSdtSearchSpaceR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sdt_SearchSpace_r17).Choice = int(index)
			switch index {
			case PDCCH_ConfigCommon_Ext_Sdt_SearchSpace_r17_NewSearchSpace:
				(*ie.Sdt_SearchSpace_r17).NewSearchSpace = new(SearchSpace)
				if err := (*ie.Sdt_SearchSpace_r17).NewSearchSpace.Decode(dx); err != nil {
					return err
				}
			case PDCCH_ConfigCommon_Ext_Sdt_SearchSpace_r17_ExistingSearchSpace:
				(*ie.Sdt_SearchSpace_r17).ExistingSearchSpace = new(SearchSpaceId)
				if err := (*ie.Sdt_SearchSpace_r17).ExistingSearchSpace.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SearchSpaceMCCH_r17 = new(SearchSpaceId)
			if err := ie.SearchSpaceMCCH_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SearchSpaceMTCH_r17 = new(SearchSpaceId)
			if err := ie.SearchSpaceMTCH_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtCommonSearchSpaceListExt2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CommonSearchSpaceListExt2_r17 = make([]SearchSpaceExt_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CommonSearchSpaceListExt2_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 = &struct {
				Choice                 int
				SCS480KHZoneEighthT    *[]int64
				SCS480KHZoneSixteenthT *[]int64
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).Choice = int(index)
			switch index {
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneEighthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneEighthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 35839))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 71679))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT)[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Pei_ConfigBWP_r17 = &struct {
				Pei_SearchSpace_r17                      SearchSpaceId
				FirstPDCCH_MonitoringOccasionOfPEI_O_r17 struct {
					Choice                                                                   int
					SCS15KHZoneT                                                             *[]int64
					SCS30KHZoneT_SCS15KHZhalfT                                               *[]int64
					SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              *[]int64
					SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          *[]int64
					SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT *[]int64
					SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT *[]int64
					SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                 *[]int64
					SCS480KHZquarterT_SCS120KHZoneSixteenthT                                 *[]int64
					SCS480KHZoneEighthT                                                      *[]int64
					SCS480KHZoneSixteenthT                                                   *[]int64
				}
			}{}
			c := ie.Pei_ConfigBWP_r17
			{
				if err := c.Pei_SearchSpace_r17.Decode(dx); err != nil {
					return err
				}
			}
			{
				choiceDec := dx.NewChoiceDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.Choice = int(index)
				switch index {
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS15KHZoneT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS15KHZoneT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS15KHZoneTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS15KHZoneT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 139))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS15KHZoneT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS30KHZoneT_SCS15KHZhalfT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS30KHZoneT_SCS15KHZhalfT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS30KHZoneTSCS15KHZhalfTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS30KHZoneT_SCS15KHZhalfT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 279))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS30KHZoneT_SCS15KHZhalfT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 559))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 1119))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 2239))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 4479))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 8959))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZquarterT_SCS120KHZoneSixteenthT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZquarterT_SCS120KHZoneSixteenthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZquarterTSCS120KHZoneSixteenthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZquarterT_SCS120KHZoneSixteenthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 17919))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZquarterT_SCS120KHZoneSixteenthT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneEighthT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneEighthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneEighthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneEighthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 35839))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneEighthT)[i] = v
					}
				case PDCCH_ConfigCommon_Ext_Pei_ConfigBWP_r17_FirstPDCCH_MonitoringOccasionOfPEI_O_r17_SCS480KHZoneSixteenthT:
					c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneSixteenthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPeiConfigBWPR17FirstPDCCHMonitoringOccasionOfPEIOR17SCS480KHZoneSixteenthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneSixteenthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 71679))
						if err != nil {
							return err
						}
						(*c.FirstPDCCH_MonitoringOccasionOfPEI_O_r17.SCS480KHZoneSixteenthT)[i] = v
					}
				}
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "followUnifiedTCI-State-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCCHConfigCommonExtFollowUnifiedTCIStateV1720Constraints)
			if err != nil {
				return err
			}
			ie.FollowUnifiedTCI_State_v1720 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "applyIndicatedTCI-State-r18", Optional: true},
				{Name: "commonSearchSpaceListExt-r18", Optional: true},
				{Name: "searchSpaceMulticastMCCH-r18", Optional: true},
				{Name: "searchSpaceMulticastMTCH-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCCHConfigCommonExtApplyIndicatedTCIStateR18Constraints)
			if err != nil {
				return err
			}
			ie.ApplyIndicatedTCI_State_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtCommonSearchSpaceListExtR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CommonSearchSpaceListExt_r18 = make([]SearchSpaceExt_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CommonSearchSpaceListExt_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SearchSpaceMulticastMCCH_r18 = new(SearchSpaceId)
			if err := ie.SearchSpaceMulticastMCCH_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.SearchSpaceMulticastMTCH_r18 = new(SearchSpaceId)
			if err := ie.SearchSpaceMulticastMTCH_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "commonSearchSpaceListExt-r19", Optional: true},
				{Name: "pagingAdaptFirstPDCCH-MonitoringOccasionOfPO-r19", Optional: true},
				{Name: "pagingAdaptFirstPDCCH-MonitoringOccasionOfPEI-O-r19", Optional: true},
				{Name: "pagingAdaptPagingSearchSpace", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtCommonSearchSpaceListExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CommonSearchSpaceListExt_r19 = make([]SearchSpaceExt_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CommonSearchSpaceListExt_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 = &struct {
				Choice                                                                                            int
				SCS15KHZoneT                                                                                      *[]int64
				SCS30KHZoneT_SCS15KHZhalfT                                                                        *[]int64
				SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       *[]int64
				SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   *[]int64
				SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          *[]int64
				SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT *[]int64
				SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 *[]int64
				SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 *[]int64
				SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     *[]int64
				SCS480KHZoneSixteenthT                                                                            *[]int64
				SCS480KHZoneThirtySecondT                                                                         *[]int64
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).Choice = int(index)
			switch index {
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS15KHZoneT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS15KHZoneTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 139))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS30KHZoneT_SCS15KHZhalfT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS30KHZoneTSCS15KHZhalfTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 279))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 559))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 1119))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 2239))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 4479))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 8959))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 17919))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 35839))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneSixteenthT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 71679))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 143359))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT)[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19 = &struct {
				Choice                                                                                            int
				SCS15KHZoneT                                                                                      *[]int64
				SCS30KHZoneT_SCS15KHZhalfT                                                                        *[]int64
				SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       *[]int64
				SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   *[]int64
				SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          *[]int64
				SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT *[]int64
				SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 *[]int64
				SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 *[]int64
				SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     *[]int64
				SCS480KHZoneSixteenthT                                                                            *[]int64
				SCS480KHZoneThirtySecondT                                                                         *[]int64
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).Choice = int(index)
			switch index {
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS15KHZoneT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS15KHZoneT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS15KHZoneTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS15KHZoneT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 139))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS15KHZoneT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS30KHZoneT_SCS15KHZhalfT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS30KHZoneT_SCS15KHZhalfT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS30KHZoneTSCS15KHZhalfTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS30KHZoneT_SCS15KHZhalfT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 279))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS30KHZoneT_SCS15KHZhalfT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 559))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 1119))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 2239))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 4479))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 8959))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 17919))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 35839))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneSixteenthT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 71679))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneSixteenthT)[i] = v
				}
			case PDCCH_ConfigCommon_Ext_PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19_SCS480KHZoneThirtySecondT:
				(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneThirtySecondT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pDCCHConfigCommonExtPagingAdaptFirstPDCCHMonitoringOccasionOfPEIOR19SCS480KHZoneThirtySecondTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneThirtySecondT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 143359))
					if err != nil {
						return err
					}
					(*(*ie.PagingAdaptFirstPDCCH_MonitoringOccasionOfPEI_O_r19).SCS480KHZoneThirtySecondT)[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.PagingAdaptPagingSearchSpace = new(SearchSpaceId)
			if err := ie.PagingAdaptPagingSearchSpace.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
