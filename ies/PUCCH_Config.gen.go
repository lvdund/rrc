// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-Config (line 11968).

var pUCCHConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSetToAddModList", Optional: true},
		{Name: "resourceSetToReleaseList", Optional: true},
		{Name: "resourceToAddModList", Optional: true},
		{Name: "resourceToReleaseList", Optional: true},
		{Name: "format1", Optional: true},
		{Name: "format2", Optional: true},
		{Name: "format3", Optional: true},
		{Name: "format4", Optional: true},
		{Name: "schedulingRequestResourceToAddModList", Optional: true},
		{Name: "schedulingRequestResourceToReleaseList", Optional: true},
		{Name: "multi-CSI-PUCCH-ResourceList", Optional: true},
		{Name: "dl-DataToUL-ACK", Optional: true},
		{Name: "spatialRelationInfoToAddModList", Optional: true},
		{Name: "spatialRelationInfoToReleaseList", Optional: true},
		{Name: "pucch-PowerControl", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var pUCCHConfigResourceSetToAddModListConstraints = per.SizeRange(1, common.MaxNrofPUCCH_ResourceSets)

var pUCCHConfigResourceSetToReleaseListConstraints = per.SizeRange(1, common.MaxNrofPUCCH_ResourceSets)

var pUCCHConfigResourceToAddModListConstraints = per.SizeRange(1, common.MaxNrofPUCCH_Resources)

var pUCCHConfigResourceToReleaseListConstraints = per.SizeRange(1, common.MaxNrofPUCCH_Resources)

var pUCCH_ConfigFormat1Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Format1_Release = 0
	PUCCH_Config_Format1_Setup   = 1
)

var pUCCH_ConfigFormat2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Format2_Release = 0
	PUCCH_Config_Format2_Setup   = 1
)

var pUCCH_ConfigFormat3Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Format3_Release = 0
	PUCCH_Config_Format3_Setup   = 1
)

var pUCCH_ConfigFormat4Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Format4_Release = 0
	PUCCH_Config_Format4_Setup   = 1
)

var pUCCHConfigSchedulingRequestResourceToAddModListConstraints = per.SizeRange(1, common.MaxNrofSR_Resources)

var pUCCHConfigSchedulingRequestResourceToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSR_Resources)

var pUCCHConfigMultiCSIPUCCHResourceListConstraints = per.SizeRange(1, 2)

var pUCCHConfigDlDataToULACKConstraints = per.SizeRange(1, 8)

var pUCCHConfigSpatialRelationInfoToAddModListConstraints = per.SizeRange(1, common.MaxNrofSpatialRelationInfos)

var pUCCHConfigSpatialRelationInfoToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSpatialRelationInfos)

var pUCCHConfigExtResourceToAddModListExtV1610Constraints = per.SizeRange(1, common.MaxNrofPUCCH_Resources)

var pUCCHConfigExtDlDataToULACKR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Dl_DataToUL_ACK_r16_Release = 0
	PUCCH_Config_Ext_Dl_DataToUL_ACK_r16_Setup   = 1
)

var pUCCHConfigExtUlAccessConfigListDCI11R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r16_Release = 0
	PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r16_Setup   = 1
)

var pUCCHConfigExtSubslotLengthForPUCCHR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "normalCP-r16"},
		{Name: "extendedCP-r16"},
	},
}

const (
	PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_NormalCP_r16   = 0
	PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_ExtendedCP_r16 = 1
)

var pUCCHConfigExtDlDataToULACKDCI12R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r16_Release = 0
	PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r16_Setup   = 1
)

var pUCCHConfigNumberOfBitsForPUCCHResourceIndicatorDCI12R16Constraints = per.Constrained(0, 3)

const (
	PUCCH_Config_Ext_Dmrs_UplinkTransformPrecodingPUCCH_r16_Enabled = 0
)

var pUCCHConfigExtDmrsUplinkTransformPrecodingPUCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Config_Ext_Dmrs_UplinkTransformPrecodingPUCCH_r16_Enabled},
}

var pUCCHConfigExtSpatialRelationInfoToAddModListSizeExtV1610Constraints = per.SizeRange(1, common.MaxNrofSpatialRelationInfosDiff_r16)

var pUCCHConfigExtSpatialRelationInfoToReleaseListSizeExtV1610Constraints = per.SizeRange(1, common.MaxNrofSpatialRelationInfosDiff_r16)

var pUCCHConfigExtSpatialRelationInfoToAddModListExtV1610Constraints = per.SizeRange(1, common.MaxNrofSpatialRelationInfos_r16)

var pUCCHConfigExtSpatialRelationInfoToReleaseListExtV1610Constraints = per.SizeRange(1, common.MaxNrofSpatialRelationInfos_r16)

var pUCCHConfigExtResourceGroupToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofPUCCH_ResourceGroups_r16)

var pUCCHConfigExtResourceGroupToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofPUCCH_ResourceGroups_r16)

var pUCCHConfigExtSpsPUCCHANListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Sps_PUCCH_AN_List_r16_Release = 0
	PUCCH_Config_Ext_Sps_PUCCH_AN_List_r16_Setup   = 1
)

var pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1610Constraints = per.SizeRange(1, common.MaxNrofSR_Resources)

var pUCCHConfigExtFormat0R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Format0_r17_Release = 0
	PUCCH_Config_Ext_Format0_r17_Setup   = 1
)

var pUCCHConfigExtFormat2ExtR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Format2Ext_r17_Release = 0
	PUCCH_Config_Ext_Format2Ext_r17_Setup   = 1
)

var pUCCHConfigExtFormat3ExtR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Format3Ext_r17_Release = 0
	PUCCH_Config_Ext_Format3Ext_r17_Setup   = 1
)

var pUCCHConfigExtFormat4ExtR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Format4Ext_r17_Release = 0
	PUCCH_Config_Ext_Format4Ext_r17_Setup   = 1
)

var pUCCHConfigExtUlAccessConfigListDCI12R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_2_r17_Release = 0
	PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_2_r17_Setup   = 1
)

const (
	PUCCH_Config_Ext_MappingPattern_r17_CyclicMapping     = 0
	PUCCH_Config_Ext_MappingPattern_r17_SequentialMapping = 1
)

var pUCCHConfigExtMappingPatternR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Config_Ext_MappingPattern_r17_CyclicMapping, PUCCH_Config_Ext_MappingPattern_r17_SequentialMapping},
}

var pUCCHConfigExtPowerControlSetInfoToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofPowerControlSetInfos_r17)

var pUCCHConfigExtPowerControlSetInfoToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofPowerControlSetInfos_r17)

const (
	PUCCH_Config_Ext_SecondTPCFieldDCI_1_1_r17_Enabled = 0
)

var pUCCHConfigExtSecondTPCFieldDCI11R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Config_Ext_SecondTPCFieldDCI_1_1_r17_Enabled},
}

const (
	PUCCH_Config_Ext_SecondTPCFieldDCI_1_2_r17_Enabled = 0
)

var pUCCHConfigExtSecondTPCFieldDCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Config_Ext_SecondTPCFieldDCI_1_2_r17_Enabled},
}

var pUCCHConfigExtDlDataToULACKR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Dl_DataToUL_ACK_r17_Release = 0
	PUCCH_Config_Ext_Dl_DataToUL_ACK_r17_Setup   = 1
)

var pUCCHConfigExtDlDataToULACKDCI12R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r17_Release = 0
	PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r17_Setup   = 1
)

var pUCCHConfigExtUlAccessConfigListDCI11R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r17_Release = 0
	PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r17_Setup   = 1
)

var pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1700Constraints = per.SizeRange(1, common.MaxNrofSR_Resources)

var pUCCHConfigExtDmrsBundlingPUCCHConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Dmrs_BundlingPUCCH_Config_r17_Release = 0
	PUCCH_Config_Ext_Dmrs_BundlingPUCCH_Config_r17_Setup   = 1
)

var pUCCHConfigExtDlDataToULACKV1700Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Dl_DataToUL_ACK_v1700_Release = 0
	PUCCH_Config_Ext_Dl_DataToUL_ACK_v1700_Setup   = 1
)

var pUCCHConfigExtDlDataToULACKMulticastDCIFormat41R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17_Release = 0
	PUCCH_Config_Ext_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17_Setup   = 1
)

var pUCCHConfigExtSpsPUCCHANListMulticastR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_Config_Ext_Sps_PUCCH_AN_ListMulticast_r17_Release = 0
	PUCCH_Config_Ext_Sps_PUCCH_AN_ListMulticast_r17_Setup   = 1
)

var pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1900Constraints = per.SizeRange(1, common.MaxNrofSR_Resources)

var pUCCHConfigExtResourceToAddModListExtV1900Constraints = per.SizeRange(1, common.MaxNrofPUCCH_Resources)

const (
	PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_NormalCP_r16_N2 = 0
	PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_NormalCP_r16_N7 = 1
)

var pUCCHConfigExtSubslotLengthForPUCCHR16NormalCPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_NormalCP_r16_N2, PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_NormalCP_r16_N7},
}

const (
	PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_ExtendedCP_r16_N2 = 0
	PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_ExtendedCP_r16_N6 = 1
)

var pUCCHConfigExtSubslotLengthForPUCCHR16ExtendedCPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_ExtendedCP_r16_N2, PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_ExtendedCP_r16_N6},
}

type PUCCH_Config struct {
	ResourceSetToAddModList  []PUCCH_ResourceSet
	ResourceSetToReleaseList []PUCCH_ResourceSetId
	ResourceToAddModList     []PUCCH_Resource
	ResourceToReleaseList    []PUCCH_ResourceId
	Format1                  *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfig
	}
	Format2 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfig
	}
	Format3 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfig
	}
	Format4 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfig
	}
	SchedulingRequestResourceToAddModList  []SchedulingRequestResourceConfig
	SchedulingRequestResourceToReleaseList []SchedulingRequestResourceId
	Multi_CSI_PUCCH_ResourceList           []PUCCH_ResourceId
	Dl_DataToUL_ACK                        []int64
	SpatialRelationInfoToAddModList        []PUCCH_SpatialRelationInfo
	SpatialRelationInfoToReleaseList       []PUCCH_SpatialRelationInfoId
	Pucch_PowerControl                     *PUCCH_PowerControl
	ResourceToAddModListExt_v1610          []PUCCH_ResourceExt_v1610
	Dl_DataToUL_ACK_r16                    *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_r16
	}
	Ul_AccessConfigListDCI_1_1_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_AccessConfigListDCI_1_1_r16
	}
	SubslotLengthForPUCCH_r16 *struct {
		Choice         int
		NormalCP_r16   *int64
		ExtendedCP_r16 *int64
	}
	Dl_DataToUL_ACK_DCI_1_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_DCI_1_2_r16
	}
	NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 *int64
	Dmrs_UplinkTransformPrecodingPUCCH_r16            *int64
	SpatialRelationInfoToAddModListSizeExt_v1610      []PUCCH_SpatialRelationInfo
	SpatialRelationInfoToReleaseListSizeExt_v1610     []PUCCH_SpatialRelationInfoId
	SpatialRelationInfoToAddModListExt_v1610          []PUCCH_SpatialRelationInfoExt_r16
	SpatialRelationInfoToReleaseListExt_v1610         []PUCCH_SpatialRelationInfoId_r16
	ResourceGroupToAddModList_r16                     []PUCCH_ResourceGroup_r16
	ResourceGroupToReleaseList_r16                    []PUCCH_ResourceGroupId_r16
	Sps_PUCCH_AN_List_r16                             *struct {
		Choice  int
		Release *struct{}
		Setup   *SPS_PUCCH_AN_List_r16
	}
	SchedulingRequestResourceToAddModListExt_v1610 []SchedulingRequestResourceConfigExt_v1610
	Format0_r17                                    *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfig
	}
	Format2Ext_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfigExt_r17
	}
	Format3Ext_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfigExt_r17
	}
	Format4Ext_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_FormatConfigExt_r17
	}
	Ul_AccessConfigListDCI_1_2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_AccessConfigListDCI_1_2_r17
	}
	MappingPattern_r17                   *int64
	PowerControlSetInfoToAddModList_r17  []PUCCH_PowerControlSetInfo_r17
	PowerControlSetInfoToReleaseList_r17 []PUCCH_PowerControlSetInfoId_r17
	SecondTPCFieldDCI_1_1_r17            *int64
	SecondTPCFieldDCI_1_2_r17            *int64
	Dl_DataToUL_ACK_r17                  *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_r17
	}
	Dl_DataToUL_ACK_DCI_1_2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_DCI_1_2_r17
	}
	Ul_AccessConfigListDCI_1_1_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_AccessConfigListDCI_1_1_r17
	}
	SchedulingRequestResourceToAddModListExt_v1700 []SchedulingRequestResourceConfigExt_v1700
	Dmrs_BundlingPUCCH_Config_r17                  *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_BundlingPUCCH_Config_r17
	}
	Dl_DataToUL_ACK_v1700 *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_v1700
	}
	Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_MulticastDCI_Format4_1_r17
	}
	Sps_PUCCH_AN_ListMulticast_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SPS_PUCCH_AN_List_r16
	}
	SchedulingRequestResourceToAddModListExt_v1900 []SchedulingRequestResourceConfigExt_v1900
	ResourceToAddModListExt_v1900                  []PUCCH_ResourceExt_v1900
}

func (ie *PUCCH_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ResourceToAddModListExt_v1610 != nil || ie.Dl_DataToUL_ACK_r16 != nil || ie.Ul_AccessConfigListDCI_1_1_r16 != nil || ie.SubslotLengthForPUCCH_r16 != nil || ie.Dl_DataToUL_ACK_DCI_1_2_r16 != nil || ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil || ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 != nil || ie.SpatialRelationInfoToAddModListSizeExt_v1610 != nil || ie.SpatialRelationInfoToReleaseListSizeExt_v1610 != nil || ie.SpatialRelationInfoToAddModListExt_v1610 != nil || ie.SpatialRelationInfoToReleaseListExt_v1610 != nil || ie.ResourceGroupToAddModList_r16 != nil || ie.ResourceGroupToReleaseList_r16 != nil || ie.Sps_PUCCH_AN_List_r16 != nil || ie.SchedulingRequestResourceToAddModListExt_v1610 != nil
	hasExtGroup1 := ie.Format0_r17 != nil || ie.Format2Ext_r17 != nil || ie.Format3Ext_r17 != nil || ie.Format4Ext_r17 != nil || ie.Ul_AccessConfigListDCI_1_2_r17 != nil || ie.MappingPattern_r17 != nil || ie.PowerControlSetInfoToAddModList_r17 != nil || ie.PowerControlSetInfoToReleaseList_r17 != nil || ie.SecondTPCFieldDCI_1_1_r17 != nil || ie.SecondTPCFieldDCI_1_2_r17 != nil || ie.Dl_DataToUL_ACK_r17 != nil || ie.Dl_DataToUL_ACK_DCI_1_2_r17 != nil || ie.Ul_AccessConfigListDCI_1_1_r17 != nil || ie.SchedulingRequestResourceToAddModListExt_v1700 != nil || ie.Dmrs_BundlingPUCCH_Config_r17 != nil || ie.Dl_DataToUL_ACK_v1700 != nil || ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil || ie.Sps_PUCCH_AN_ListMulticast_r17 != nil
	hasExtGroup2 := ie.SchedulingRequestResourceToAddModListExt_v1900 != nil || ie.ResourceToAddModListExt_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResourceSetToAddModList != nil, ie.ResourceSetToReleaseList != nil, ie.ResourceToAddModList != nil, ie.ResourceToReleaseList != nil, ie.Format1 != nil, ie.Format2 != nil, ie.Format3 != nil, ie.Format4 != nil, ie.SchedulingRequestResourceToAddModList != nil, ie.SchedulingRequestResourceToReleaseList != nil, ie.Multi_CSI_PUCCH_ResourceList != nil, ie.Dl_DataToUL_ACK != nil, ie.SpatialRelationInfoToAddModList != nil, ie.SpatialRelationInfoToReleaseList != nil, ie.Pucch_PowerControl != nil}); err != nil {
		return err
	}

	// 3. resourceSetToAddModList: sequence-of
	{
		if ie.ResourceSetToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigResourceSetToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ResourceSetToAddModList))); err != nil {
				return err
			}
			for i := range ie.ResourceSetToAddModList {
				if err := ie.ResourceSetToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. resourceSetToReleaseList: sequence-of
	{
		if ie.ResourceSetToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigResourceSetToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ResourceSetToReleaseList))); err != nil {
				return err
			}
			for i := range ie.ResourceSetToReleaseList {
				if err := ie.ResourceSetToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. resourceToAddModList: sequence-of
	{
		if ie.ResourceToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigResourceToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ResourceToAddModList))); err != nil {
				return err
			}
			for i := range ie.ResourceToAddModList {
				if err := ie.ResourceToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. resourceToReleaseList: sequence-of
	{
		if ie.ResourceToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigResourceToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ResourceToReleaseList))); err != nil {
				return err
			}
			for i := range ie.ResourceToReleaseList {
				if err := ie.ResourceToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. format1: choice
	{
		if ie.Format1 != nil {
			choiceEnc := e.NewChoiceEncoder(pUCCH_ConfigFormat1Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Format1).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Format1).Choice {
			case PUCCH_Config_Format1_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format1_Setup:
				if err := (*ie.Format1).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Format1).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. format2: choice
	{
		if ie.Format2 != nil {
			choiceEnc := e.NewChoiceEncoder(pUCCH_ConfigFormat2Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Format2).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Format2).Choice {
			case PUCCH_Config_Format2_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format2_Setup:
				if err := (*ie.Format2).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Format2).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. format3: choice
	{
		if ie.Format3 != nil {
			choiceEnc := e.NewChoiceEncoder(pUCCH_ConfigFormat3Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Format3).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Format3).Choice {
			case PUCCH_Config_Format3_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format3_Setup:
				if err := (*ie.Format3).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Format3).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. format4: choice
	{
		if ie.Format4 != nil {
			choiceEnc := e.NewChoiceEncoder(pUCCH_ConfigFormat4Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Format4).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Format4).Choice {
			case PUCCH_Config_Format4_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format4_Setup:
				if err := (*ie.Format4).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Format4).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. schedulingRequestResourceToAddModList: sequence-of
	{
		if ie.SchedulingRequestResourceToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigSchedulingRequestResourceToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestResourceToAddModList))); err != nil {
				return err
			}
			for i := range ie.SchedulingRequestResourceToAddModList {
				if err := ie.SchedulingRequestResourceToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 12. schedulingRequestResourceToReleaseList: sequence-of
	{
		if ie.SchedulingRequestResourceToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigSchedulingRequestResourceToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestResourceToReleaseList))); err != nil {
				return err
			}
			for i := range ie.SchedulingRequestResourceToReleaseList {
				if err := ie.SchedulingRequestResourceToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. multi-CSI-PUCCH-ResourceList: sequence-of
	{
		if ie.Multi_CSI_PUCCH_ResourceList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigMultiCSIPUCCHResourceListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Multi_CSI_PUCCH_ResourceList))); err != nil {
				return err
			}
			for i := range ie.Multi_CSI_PUCCH_ResourceList {
				if err := ie.Multi_CSI_PUCCH_ResourceList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 14. dl-DataToUL-ACK: sequence-of
	{
		if ie.Dl_DataToUL_ACK != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigDlDataToULACKConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Dl_DataToUL_ACK))); err != nil {
				return err
			}
			for i := range ie.Dl_DataToUL_ACK {
				if err := e.EncodeInteger(ie.Dl_DataToUL_ACK[i], per.Constrained(0, 15)); err != nil {
					return err
				}
			}
		}
	}

	// 15. spatialRelationInfoToAddModList: sequence-of
	{
		if ie.SpatialRelationInfoToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigSpatialRelationInfoToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SpatialRelationInfoToAddModList))); err != nil {
				return err
			}
			for i := range ie.SpatialRelationInfoToAddModList {
				if err := ie.SpatialRelationInfoToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 16. spatialRelationInfoToReleaseList: sequence-of
	{
		if ie.SpatialRelationInfoToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHConfigSpatialRelationInfoToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SpatialRelationInfoToReleaseList))); err != nil {
				return err
			}
			for i := range ie.SpatialRelationInfoToReleaseList {
				if err := ie.SpatialRelationInfoToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 17. pucch-PowerControl: ref
	{
		if ie.Pucch_PowerControl != nil {
			if err := ie.Pucch_PowerControl.Encode(e); err != nil {
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
					{Name: "resourceToAddModListExt-v1610", Optional: true},
					{Name: "dl-DataToUL-ACK-r16", Optional: true},
					{Name: "ul-AccessConfigListDCI-1-1-r16", Optional: true},
					{Name: "subslotLengthForPUCCH-r16", Optional: true},
					{Name: "dl-DataToUL-ACK-DCI-1-2-r16", Optional: true},
					{Name: "numberOfBitsForPUCCH-ResourceIndicatorDCI-1-2-r16", Optional: true},
					{Name: "dmrs-UplinkTransformPrecodingPUCCH-r16", Optional: true},
					{Name: "spatialRelationInfoToAddModListSizeExt-v1610", Optional: true},
					{Name: "spatialRelationInfoToReleaseListSizeExt-v1610", Optional: true},
					{Name: "spatialRelationInfoToAddModListExt-v1610", Optional: true},
					{Name: "spatialRelationInfoToReleaseListExt-v1610", Optional: true},
					{Name: "resourceGroupToAddModList-r16", Optional: true},
					{Name: "resourceGroupToReleaseList-r16", Optional: true},
					{Name: "sps-PUCCH-AN-List-r16", Optional: true},
					{Name: "schedulingRequestResourceToAddModListExt-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ResourceToAddModListExt_v1610 != nil, ie.Dl_DataToUL_ACK_r16 != nil, ie.Ul_AccessConfigListDCI_1_1_r16 != nil, ie.SubslotLengthForPUCCH_r16 != nil, ie.Dl_DataToUL_ACK_DCI_1_2_r16 != nil, ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil, ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 != nil, ie.SpatialRelationInfoToAddModListSizeExt_v1610 != nil, ie.SpatialRelationInfoToReleaseListSizeExt_v1610 != nil, ie.SpatialRelationInfoToAddModListExt_v1610 != nil, ie.SpatialRelationInfoToReleaseListExt_v1610 != nil, ie.ResourceGroupToAddModList_r16 != nil, ie.ResourceGroupToReleaseList_r16 != nil, ie.Sps_PUCCH_AN_List_r16 != nil, ie.SchedulingRequestResourceToAddModListExt_v1610 != nil}); err != nil {
				return err
			}

			if ie.ResourceToAddModListExt_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtResourceToAddModListExtV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ResourceToAddModListExt_v1610))); err != nil {
					return err
				}
				for i := range ie.ResourceToAddModListExt_v1610 {
					if err := ie.ResourceToAddModListExt_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dl_DataToUL_ACK_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtDlDataToULACKR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_r16).Choice {
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_r16_Setup:
					if err := (*ie.Dl_DataToUL_ACK_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ul_AccessConfigListDCI_1_1_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtUlAccessConfigListDCI11R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_AccessConfigListDCI_1_1_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_AccessConfigListDCI_1_1_r16).Choice {
				case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r16_Setup:
					if err := (*ie.Ul_AccessConfigListDCI_1_1_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SubslotLengthForPUCCH_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtSubslotLengthForPUCCHR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.SubslotLengthForPUCCH_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.SubslotLengthForPUCCH_r16).Choice {
				case PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_NormalCP_r16:
					if err := ex.EncodeEnumerated((*(*ie.SubslotLengthForPUCCH_r16).NormalCP_r16), pUCCHConfigExtSubslotLengthForPUCCHR16NormalCPR16Constraints); err != nil {
						return err
					}
				case PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_ExtendedCP_r16:
					if err := ex.EncodeEnumerated((*(*ie.SubslotLengthForPUCCH_r16).ExtendedCP_r16), pUCCHConfigExtSubslotLengthForPUCCHR16ExtendedCPR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Dl_DataToUL_ACK_DCI_1_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtDlDataToULACKDCI12R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_DCI_1_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_DCI_1_2_r16).Choice {
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r16_Setup:
					if err := (*ie.Dl_DataToUL_ACK_DCI_1_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16, pUCCHConfigNumberOfBitsForPUCCHResourceIndicatorDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_UplinkTransformPrecodingPUCCH_r16, pUCCHConfigExtDmrsUplinkTransformPrecodingPUCCHR16Constraints); err != nil {
					return err
				}
			}

			if ie.SpatialRelationInfoToAddModListSizeExt_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtSpatialRelationInfoToAddModListSizeExtV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SpatialRelationInfoToAddModListSizeExt_v1610))); err != nil {
					return err
				}
				for i := range ie.SpatialRelationInfoToAddModListSizeExt_v1610 {
					if err := ie.SpatialRelationInfoToAddModListSizeExt_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SpatialRelationInfoToReleaseListSizeExt_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtSpatialRelationInfoToReleaseListSizeExtV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SpatialRelationInfoToReleaseListSizeExt_v1610))); err != nil {
					return err
				}
				for i := range ie.SpatialRelationInfoToReleaseListSizeExt_v1610 {
					if err := ie.SpatialRelationInfoToReleaseListSizeExt_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SpatialRelationInfoToAddModListExt_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtSpatialRelationInfoToAddModListExtV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SpatialRelationInfoToAddModListExt_v1610))); err != nil {
					return err
				}
				for i := range ie.SpatialRelationInfoToAddModListExt_v1610 {
					if err := ie.SpatialRelationInfoToAddModListExt_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SpatialRelationInfoToReleaseListExt_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtSpatialRelationInfoToReleaseListExtV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SpatialRelationInfoToReleaseListExt_v1610))); err != nil {
					return err
				}
				for i := range ie.SpatialRelationInfoToReleaseListExt_v1610 {
					if err := ie.SpatialRelationInfoToReleaseListExt_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ResourceGroupToAddModList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtResourceGroupToAddModListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ResourceGroupToAddModList_r16))); err != nil {
					return err
				}
				for i := range ie.ResourceGroupToAddModList_r16 {
					if err := ie.ResourceGroupToAddModList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ResourceGroupToReleaseList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtResourceGroupToReleaseListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ResourceGroupToReleaseList_r16))); err != nil {
					return err
				}
				for i := range ie.ResourceGroupToReleaseList_r16 {
					if err := ie.ResourceGroupToReleaseList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sps_PUCCH_AN_List_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtSpsPUCCHANListR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sps_PUCCH_AN_List_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sps_PUCCH_AN_List_r16).Choice {
				case PUCCH_Config_Ext_Sps_PUCCH_AN_List_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Sps_PUCCH_AN_List_r16_Setup:
					if err := (*ie.Sps_PUCCH_AN_List_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SchedulingRequestResourceToAddModListExt_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestResourceToAddModListExt_v1610))); err != nil {
					return err
				}
				for i := range ie.SchedulingRequestResourceToAddModListExt_v1610 {
					if err := ie.SchedulingRequestResourceToAddModListExt_v1610[i].Encode(ex); err != nil {
						return err
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
					{Name: "format0-r17", Optional: true},
					{Name: "format2Ext-r17", Optional: true},
					{Name: "format3Ext-r17", Optional: true},
					{Name: "format4Ext-r17", Optional: true},
					{Name: "ul-AccessConfigListDCI-1-2-r17", Optional: true},
					{Name: "mappingPattern-r17", Optional: true},
					{Name: "powerControlSetInfoToAddModList-r17", Optional: true},
					{Name: "powerControlSetInfoToReleaseList-r17", Optional: true},
					{Name: "secondTPCFieldDCI-1-1-r17", Optional: true},
					{Name: "secondTPCFieldDCI-1-2-r17", Optional: true},
					{Name: "dl-DataToUL-ACK-r17", Optional: true},
					{Name: "dl-DataToUL-ACK-DCI-1-2-r17", Optional: true},
					{Name: "ul-AccessConfigListDCI-1-1-r17", Optional: true},
					{Name: "schedulingRequestResourceToAddModListExt-v1700", Optional: true},
					{Name: "dmrs-BundlingPUCCH-Config-r17", Optional: true},
					{Name: "dl-DataToUL-ACK-v1700", Optional: true},
					{Name: "dl-DataToUL-ACK-MulticastDCI-Format4-1-r17", Optional: true},
					{Name: "sps-PUCCH-AN-ListMulticast-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Format0_r17 != nil, ie.Format2Ext_r17 != nil, ie.Format3Ext_r17 != nil, ie.Format4Ext_r17 != nil, ie.Ul_AccessConfigListDCI_1_2_r17 != nil, ie.MappingPattern_r17 != nil, ie.PowerControlSetInfoToAddModList_r17 != nil, ie.PowerControlSetInfoToReleaseList_r17 != nil, ie.SecondTPCFieldDCI_1_1_r17 != nil, ie.SecondTPCFieldDCI_1_2_r17 != nil, ie.Dl_DataToUL_ACK_r17 != nil, ie.Dl_DataToUL_ACK_DCI_1_2_r17 != nil, ie.Ul_AccessConfigListDCI_1_1_r17 != nil, ie.SchedulingRequestResourceToAddModListExt_v1700 != nil, ie.Dmrs_BundlingPUCCH_Config_r17 != nil, ie.Dl_DataToUL_ACK_v1700 != nil, ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil, ie.Sps_PUCCH_AN_ListMulticast_r17 != nil}); err != nil {
				return err
			}

			if ie.Format0_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtFormat0R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Format0_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Format0_r17).Choice {
				case PUCCH_Config_Ext_Format0_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Format0_r17_Setup:
					if err := (*ie.Format0_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Format2Ext_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtFormat2ExtR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Format2Ext_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Format2Ext_r17).Choice {
				case PUCCH_Config_Ext_Format2Ext_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Format2Ext_r17_Setup:
					if err := (*ie.Format2Ext_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Format3Ext_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtFormat3ExtR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Format3Ext_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Format3Ext_r17).Choice {
				case PUCCH_Config_Ext_Format3Ext_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Format3Ext_r17_Setup:
					if err := (*ie.Format3Ext_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Format4Ext_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtFormat4ExtR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Format4Ext_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Format4Ext_r17).Choice {
				case PUCCH_Config_Ext_Format4Ext_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Format4Ext_r17_Setup:
					if err := (*ie.Format4Ext_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ul_AccessConfigListDCI_1_2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtUlAccessConfigListDCI12R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_AccessConfigListDCI_1_2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_AccessConfigListDCI_1_2_r17).Choice {
				case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_2_r17_Setup:
					if err := (*ie.Ul_AccessConfigListDCI_1_2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MappingPattern_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MappingPattern_r17, pUCCHConfigExtMappingPatternR17Constraints); err != nil {
					return err
				}
			}

			if ie.PowerControlSetInfoToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtPowerControlSetInfoToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.PowerControlSetInfoToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.PowerControlSetInfoToAddModList_r17 {
					if err := ie.PowerControlSetInfoToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PowerControlSetInfoToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtPowerControlSetInfoToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.PowerControlSetInfoToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.PowerControlSetInfoToReleaseList_r17 {
					if err := ie.PowerControlSetInfoToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SecondTPCFieldDCI_1_1_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SecondTPCFieldDCI_1_1_r17, pUCCHConfigExtSecondTPCFieldDCI11R17Constraints); err != nil {
					return err
				}
			}

			if ie.SecondTPCFieldDCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SecondTPCFieldDCI_1_2_r17, pUCCHConfigExtSecondTPCFieldDCI12R17Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_DataToUL_ACK_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtDlDataToULACKR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_r17).Choice {
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_r17_Setup:
					if err := (*ie.Dl_DataToUL_ACK_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dl_DataToUL_ACK_DCI_1_2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtDlDataToULACKDCI12R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_DCI_1_2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_DCI_1_2_r17).Choice {
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r17_Setup:
					if err := (*ie.Dl_DataToUL_ACK_DCI_1_2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ul_AccessConfigListDCI_1_1_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtUlAccessConfigListDCI11R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_AccessConfigListDCI_1_1_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_AccessConfigListDCI_1_1_r17).Choice {
				case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r17_Setup:
					if err := (*ie.Ul_AccessConfigListDCI_1_1_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SchedulingRequestResourceToAddModListExt_v1700 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1700Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestResourceToAddModListExt_v1700))); err != nil {
					return err
				}
				for i := range ie.SchedulingRequestResourceToAddModListExt_v1700 {
					if err := ie.SchedulingRequestResourceToAddModListExt_v1700[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dmrs_BundlingPUCCH_Config_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtDmrsBundlingPUCCHConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_BundlingPUCCH_Config_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dmrs_BundlingPUCCH_Config_r17).Choice {
				case PUCCH_Config_Ext_Dmrs_BundlingPUCCH_Config_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Dmrs_BundlingPUCCH_Config_r17_Setup:
					if err := (*ie.Dmrs_BundlingPUCCH_Config_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dl_DataToUL_ACK_v1700 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtDlDataToULACKV1700Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_v1700).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_v1700).Choice {
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_v1700_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_v1700_Setup:
					if err := (*ie.Dl_DataToUL_ACK_v1700).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtDlDataToULACKMulticastDCIFormat41R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17).Choice {
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17_Setup:
					if err := (*ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sps_PUCCH_AN_ListMulticast_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHConfigExtSpsPUCCHANListMulticastR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sps_PUCCH_AN_ListMulticast_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sps_PUCCH_AN_ListMulticast_r17).Choice {
				case PUCCH_Config_Ext_Sps_PUCCH_AN_ListMulticast_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_Config_Ext_Sps_PUCCH_AN_ListMulticast_r17_Setup:
					if err := (*ie.Sps_PUCCH_AN_ListMulticast_r17).Setup.Encode(ex); err != nil {
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
					{Name: "schedulingRequestResourceToAddModListExt-v1900", Optional: true},
					{Name: "resourceToAddModListExt-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SchedulingRequestResourceToAddModListExt_v1900 != nil, ie.ResourceToAddModListExt_v1900 != nil}); err != nil {
				return err
			}

			if ie.SchedulingRequestResourceToAddModListExt_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestResourceToAddModListExt_v1900))); err != nil {
					return err
				}
				for i := range ie.SchedulingRequestResourceToAddModListExt_v1900 {
					if err := ie.SchedulingRequestResourceToAddModListExt_v1900[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ResourceToAddModListExt_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUCCHConfigExtResourceToAddModListExtV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ResourceToAddModListExt_v1900))); err != nil {
					return err
				}
				for i := range ie.ResourceToAddModListExt_v1900 {
					if err := ie.ResourceToAddModListExt_v1900[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *PUCCH_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. resourceSetToAddModList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigResourceSetToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceSetToAddModList = make([]PUCCH_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceSetToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. resourceSetToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigResourceSetToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceSetToReleaseList = make([]PUCCH_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceSetToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. resourceToAddModList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigResourceToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceToAddModList = make([]PUCCH_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. resourceToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigResourceToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceToReleaseList = make([]PUCCH_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. format1: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Format1 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pUCCH_ConfigFormat1Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format1).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUCCH_Config_Format1_Release:
				(*ie.Format1).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format1_Setup:
				(*ie.Format1).Setup = new(PUCCH_FormatConfig)
				if err := (*ie.Format1).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. format2: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Format2 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pUCCH_ConfigFormat2Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format2).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUCCH_Config_Format2_Release:
				(*ie.Format2).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format2_Setup:
				(*ie.Format2).Setup = new(PUCCH_FormatConfig)
				if err := (*ie.Format2).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. format3: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Format3 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pUCCH_ConfigFormat3Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format3).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUCCH_Config_Format3_Release:
				(*ie.Format3).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format3_Setup:
				(*ie.Format3).Setup = new(PUCCH_FormatConfig)
				if err := (*ie.Format3).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. format4: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Format4 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pUCCH_ConfigFormat4Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format4).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUCCH_Config_Format4_Release:
				(*ie.Format4).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Format4_Setup:
				(*ie.Format4).Setup = new(PUCCH_FormatConfig)
				if err := (*ie.Format4).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. schedulingRequestResourceToAddModList: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigSchedulingRequestResourceToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestResourceToAddModList = make([]SchedulingRequestResourceConfig, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestResourceToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. schedulingRequestResourceToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigSchedulingRequestResourceToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestResourceToReleaseList = make([]SchedulingRequestResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestResourceToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. multi-CSI-PUCCH-ResourceList: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigMultiCSIPUCCHResourceListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Multi_CSI_PUCCH_ResourceList = make([]PUCCH_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Multi_CSI_PUCCH_ResourceList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 14. dl-DataToUL-ACK: sequence-of
	{
		if seq.IsComponentPresent(11) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigDlDataToULACKConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dl_DataToUL_ACK = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				ie.Dl_DataToUL_ACK[i] = v
			}
		}
	}

	// 15. spatialRelationInfoToAddModList: sequence-of
	{
		if seq.IsComponentPresent(12) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigSpatialRelationInfoToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SpatialRelationInfoToAddModList = make([]PUCCH_SpatialRelationInfo, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SpatialRelationInfoToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 16. spatialRelationInfoToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(13) {
			seqOf := d.NewSequenceOfDecoder(pUCCHConfigSpatialRelationInfoToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SpatialRelationInfoToReleaseList = make([]PUCCH_SpatialRelationInfoId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SpatialRelationInfoToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 17. pucch-PowerControl: ref
	{
		if seq.IsComponentPresent(14) {
			ie.Pucch_PowerControl = new(PUCCH_PowerControl)
			if err := ie.Pucch_PowerControl.Decode(d); err != nil {
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
				{Name: "resourceToAddModListExt-v1610", Optional: true},
				{Name: "dl-DataToUL-ACK-r16", Optional: true},
				{Name: "ul-AccessConfigListDCI-1-1-r16", Optional: true},
				{Name: "subslotLengthForPUCCH-r16", Optional: true},
				{Name: "dl-DataToUL-ACK-DCI-1-2-r16", Optional: true},
				{Name: "numberOfBitsForPUCCH-ResourceIndicatorDCI-1-2-r16", Optional: true},
				{Name: "dmrs-UplinkTransformPrecodingPUCCH-r16", Optional: true},
				{Name: "spatialRelationInfoToAddModListSizeExt-v1610", Optional: true},
				{Name: "spatialRelationInfoToReleaseListSizeExt-v1610", Optional: true},
				{Name: "spatialRelationInfoToAddModListExt-v1610", Optional: true},
				{Name: "spatialRelationInfoToReleaseListExt-v1610", Optional: true},
				{Name: "resourceGroupToAddModList-r16", Optional: true},
				{Name: "resourceGroupToReleaseList-r16", Optional: true},
				{Name: "sps-PUCCH-AN-List-r16", Optional: true},
				{Name: "schedulingRequestResourceToAddModListExt-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtResourceToAddModListExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceToAddModListExt_v1610 = make([]PUCCH_ResourceExt_v1610, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceToAddModListExt_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Dl_DataToUL_ACK_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtDlDataToULACKR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_r16).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_r16_Release:
				(*ie.Dl_DataToUL_ACK_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_r16_Setup:
				(*ie.Dl_DataToUL_ACK_r16).Setup = new(DL_DataToUL_ACK_r16)
				if err := (*ie.Dl_DataToUL_ACK_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ul_AccessConfigListDCI_1_1_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_AccessConfigListDCI_1_1_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtUlAccessConfigListDCI11R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_AccessConfigListDCI_1_1_r16).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r16_Release:
				(*ie.Ul_AccessConfigListDCI_1_1_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r16_Setup:
				(*ie.Ul_AccessConfigListDCI_1_1_r16).Setup = new(UL_AccessConfigListDCI_1_1_r16)
				if err := (*ie.Ul_AccessConfigListDCI_1_1_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.SubslotLengthForPUCCH_r16 = &struct {
				Choice         int
				NormalCP_r16   *int64
				ExtendedCP_r16 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtSubslotLengthForPUCCHR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SubslotLengthForPUCCH_r16).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_NormalCP_r16:
				(*ie.SubslotLengthForPUCCH_r16).NormalCP_r16 = new(int64)
				v, err := dx.DecodeEnumerated(pUCCHConfigExtSubslotLengthForPUCCHR16NormalCPR16Constraints)
				if err != nil {
					return err
				}
				(*(*ie.SubslotLengthForPUCCH_r16).NormalCP_r16) = v
			case PUCCH_Config_Ext_SubslotLengthForPUCCH_r16_ExtendedCP_r16:
				(*ie.SubslotLengthForPUCCH_r16).ExtendedCP_r16 = new(int64)
				v, err := dx.DecodeEnumerated(pUCCHConfigExtSubslotLengthForPUCCHR16ExtendedCPR16Constraints)
				if err != nil {
					return err
				}
				(*(*ie.SubslotLengthForPUCCH_r16).ExtendedCP_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Dl_DataToUL_ACK_DCI_1_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_DCI_1_2_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtDlDataToULACKDCI12R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_DCI_1_2_r16).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r16_Release:
				(*ie.Dl_DataToUL_ACK_DCI_1_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r16_Setup:
				(*ie.Dl_DataToUL_ACK_DCI_1_2_r16).Setup = new(DL_DataToUL_ACK_DCI_1_2_r16)
				if err := (*ie.Dl_DataToUL_ACK_DCI_1_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeInteger(pUCCHConfigNumberOfBitsForPUCCHResourceIndicatorDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(pUCCHConfigExtDmrsUplinkTransformPrecodingPUCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtSpatialRelationInfoToAddModListSizeExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SpatialRelationInfoToAddModListSizeExt_v1610 = make([]PUCCH_SpatialRelationInfo, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SpatialRelationInfoToAddModListSizeExt_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtSpatialRelationInfoToReleaseListSizeExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SpatialRelationInfoToReleaseListSizeExt_v1610 = make([]PUCCH_SpatialRelationInfoId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SpatialRelationInfoToReleaseListSizeExt_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtSpatialRelationInfoToAddModListExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SpatialRelationInfoToAddModListExt_v1610 = make([]PUCCH_SpatialRelationInfoExt_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SpatialRelationInfoToAddModListExt_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtSpatialRelationInfoToReleaseListExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SpatialRelationInfoToReleaseListExt_v1610 = make([]PUCCH_SpatialRelationInfoId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SpatialRelationInfoToReleaseListExt_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(11) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtResourceGroupToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceGroupToAddModList_r16 = make([]PUCCH_ResourceGroup_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceGroupToAddModList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(12) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtResourceGroupToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceGroupToReleaseList_r16 = make([]PUCCH_ResourceGroupId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceGroupToReleaseList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(13) {
			ie.Sps_PUCCH_AN_List_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SPS_PUCCH_AN_List_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtSpsPUCCHANListR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sps_PUCCH_AN_List_r16).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Sps_PUCCH_AN_List_r16_Release:
				(*ie.Sps_PUCCH_AN_List_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Sps_PUCCH_AN_List_r16_Setup:
				(*ie.Sps_PUCCH_AN_List_r16).Setup = new(SPS_PUCCH_AN_List_r16)
				if err := (*ie.Sps_PUCCH_AN_List_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(14) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestResourceToAddModListExt_v1610 = make([]SchedulingRequestResourceConfigExt_v1610, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestResourceToAddModListExt_v1610[i].Decode(dx); err != nil {
					return err
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
				{Name: "format0-r17", Optional: true},
				{Name: "format2Ext-r17", Optional: true},
				{Name: "format3Ext-r17", Optional: true},
				{Name: "format4Ext-r17", Optional: true},
				{Name: "ul-AccessConfigListDCI-1-2-r17", Optional: true},
				{Name: "mappingPattern-r17", Optional: true},
				{Name: "powerControlSetInfoToAddModList-r17", Optional: true},
				{Name: "powerControlSetInfoToReleaseList-r17", Optional: true},
				{Name: "secondTPCFieldDCI-1-1-r17", Optional: true},
				{Name: "secondTPCFieldDCI-1-2-r17", Optional: true},
				{Name: "dl-DataToUL-ACK-r17", Optional: true},
				{Name: "dl-DataToUL-ACK-DCI-1-2-r17", Optional: true},
				{Name: "ul-AccessConfigListDCI-1-1-r17", Optional: true},
				{Name: "schedulingRequestResourceToAddModListExt-v1700", Optional: true},
				{Name: "dmrs-BundlingPUCCH-Config-r17", Optional: true},
				{Name: "dl-DataToUL-ACK-v1700", Optional: true},
				{Name: "dl-DataToUL-ACK-MulticastDCI-Format4-1-r17", Optional: true},
				{Name: "sps-PUCCH-AN-ListMulticast-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Format0_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfig
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtFormat0R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format0_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Format0_r17_Release:
				(*ie.Format0_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Format0_r17_Setup:
				(*ie.Format0_r17).Setup = new(PUCCH_FormatConfig)
				if err := (*ie.Format0_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Format2Ext_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfigExt_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtFormat2ExtR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format2Ext_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Format2Ext_r17_Release:
				(*ie.Format2Ext_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Format2Ext_r17_Setup:
				(*ie.Format2Ext_r17).Setup = new(PUCCH_FormatConfigExt_r17)
				if err := (*ie.Format2Ext_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Format3Ext_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfigExt_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtFormat3ExtR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format3Ext_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Format3Ext_r17_Release:
				(*ie.Format3Ext_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Format3Ext_r17_Setup:
				(*ie.Format3Ext_r17).Setup = new(PUCCH_FormatConfigExt_r17)
				if err := (*ie.Format3Ext_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Format4Ext_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_FormatConfigExt_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtFormat4ExtR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format4Ext_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Format4Ext_r17_Release:
				(*ie.Format4Ext_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Format4Ext_r17_Setup:
				(*ie.Format4Ext_r17).Setup = new(PUCCH_FormatConfigExt_r17)
				if err := (*ie.Format4Ext_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Ul_AccessConfigListDCI_1_2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_AccessConfigListDCI_1_2_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtUlAccessConfigListDCI12R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_AccessConfigListDCI_1_2_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_2_r17_Release:
				(*ie.Ul_AccessConfigListDCI_1_2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_2_r17_Setup:
				(*ie.Ul_AccessConfigListDCI_1_2_r17).Setup = new(UL_AccessConfigListDCI_1_2_r17)
				if err := (*ie.Ul_AccessConfigListDCI_1_2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(pUCCHConfigExtMappingPatternR17Constraints)
			if err != nil {
				return err
			}
			ie.MappingPattern_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtPowerControlSetInfoToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PowerControlSetInfoToAddModList_r17 = make([]PUCCH_PowerControlSetInfo_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PowerControlSetInfoToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtPowerControlSetInfoToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PowerControlSetInfoToReleaseList_r17 = make([]PUCCH_PowerControlSetInfoId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PowerControlSetInfoToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(pUCCHConfigExtSecondTPCFieldDCI11R17Constraints)
			if err != nil {
				return err
			}
			ie.SecondTPCFieldDCI_1_1_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(pUCCHConfigExtSecondTPCFieldDCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.SecondTPCFieldDCI_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			ie.Dl_DataToUL_ACK_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtDlDataToULACKR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_r17_Release:
				(*ie.Dl_DataToUL_ACK_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_r17_Setup:
				(*ie.Dl_DataToUL_ACK_r17).Setup = new(DL_DataToUL_ACK_r17)
				if err := (*ie.Dl_DataToUL_ACK_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(11) {
			ie.Dl_DataToUL_ACK_DCI_1_2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_DCI_1_2_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtDlDataToULACKDCI12R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_DCI_1_2_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r17_Release:
				(*ie.Dl_DataToUL_ACK_DCI_1_2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_DCI_1_2_r17_Setup:
				(*ie.Dl_DataToUL_ACK_DCI_1_2_r17).Setup = new(DL_DataToUL_ACK_DCI_1_2_r17)
				if err := (*ie.Dl_DataToUL_ACK_DCI_1_2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(12) {
			ie.Ul_AccessConfigListDCI_1_1_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_AccessConfigListDCI_1_1_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtUlAccessConfigListDCI11R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_AccessConfigListDCI_1_1_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r17_Release:
				(*ie.Ul_AccessConfigListDCI_1_1_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Ul_AccessConfigListDCI_1_1_r17_Setup:
				(*ie.Ul_AccessConfigListDCI_1_1_r17).Setup = new(UL_AccessConfigListDCI_1_1_r17)
				if err := (*ie.Ul_AccessConfigListDCI_1_1_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(13) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestResourceToAddModListExt_v1700 = make([]SchedulingRequestResourceConfigExt_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestResourceToAddModListExt_v1700[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(14) {
			ie.Dmrs_BundlingPUCCH_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_BundlingPUCCH_Config_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtDmrsBundlingPUCCHConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_BundlingPUCCH_Config_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Dmrs_BundlingPUCCH_Config_r17_Release:
				(*ie.Dmrs_BundlingPUCCH_Config_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Dmrs_BundlingPUCCH_Config_r17_Setup:
				(*ie.Dmrs_BundlingPUCCH_Config_r17).Setup = new(DMRS_BundlingPUCCH_Config_r17)
				if err := (*ie.Dmrs_BundlingPUCCH_Config_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(15) {
			ie.Dl_DataToUL_ACK_v1700 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_v1700
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtDlDataToULACKV1700Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_v1700).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_v1700_Release:
				(*ie.Dl_DataToUL_ACK_v1700).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_v1700_Setup:
				(*ie.Dl_DataToUL_ACK_v1700).Setup = new(DL_DataToUL_ACK_v1700)
				if err := (*ie.Dl_DataToUL_ACK_v1700).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(16) {
			ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_MulticastDCI_Format4_1_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtDlDataToULACKMulticastDCIFormat41R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17_Release:
				(*ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17_Setup:
				(*ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17).Setup = new(DL_DataToUL_ACK_MulticastDCI_Format4_1_r17)
				if err := (*ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(17) {
			ie.Sps_PUCCH_AN_ListMulticast_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SPS_PUCCH_AN_List_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHConfigExtSpsPUCCHANListMulticastR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sps_PUCCH_AN_ListMulticast_r17).Choice = int(index)
			switch index {
			case PUCCH_Config_Ext_Sps_PUCCH_AN_ListMulticast_r17_Release:
				(*ie.Sps_PUCCH_AN_ListMulticast_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_Config_Ext_Sps_PUCCH_AN_ListMulticast_r17_Setup:
				(*ie.Sps_PUCCH_AN_ListMulticast_r17).Setup = new(SPS_PUCCH_AN_List_r16)
				if err := (*ie.Sps_PUCCH_AN_ListMulticast_r17).Setup.Decode(dx); err != nil {
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
				{Name: "schedulingRequestResourceToAddModListExt-v1900", Optional: true},
				{Name: "resourceToAddModListExt-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtSchedulingRequestResourceToAddModListExtV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestResourceToAddModListExt_v1900 = make([]SchedulingRequestResourceConfigExt_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestResourceToAddModListExt_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(pUCCHConfigExtResourceToAddModListExtV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ResourceToAddModListExt_v1900 = make([]PUCCH_ResourceExt_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ResourceToAddModListExt_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
