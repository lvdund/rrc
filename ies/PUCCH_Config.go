package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Config struct {
	ResourceSetToAddModList                           []PUCCH_ResourceSet                                  `lb:1,ub:maxNrofPUCCH_ResourceSets,optional`
	ResourceSetToReleaseList                          []PUCCH_ResourceSetId                                `lb:1,ub:maxNrofPUCCH_ResourceSets,optional`
	ResourceToAddModList                              []PUCCH_Resource                                     `lb:1,ub:maxNrofPUCCH_Resources,optional`
	ResourceToReleaseList                             []PUCCH_ResourceId                                   `lb:1,ub:maxNrofPUCCH_Resources,optional`
	Format1                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	Format2                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	Format3                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	Format4                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	SchedulingRequestResourceToAddModList             []SchedulingRequestResourceConfig                    `lb:1,ub:maxNrofSR_Resources,optional`
	SchedulingRequestResourceToReleaseList            []SchedulingRequestResourceId                        `lb:1,ub:maxNrofSR_Resources,optional`
	Multi_CSI_PUCCH_ResourceList                      []PUCCH_ResourceId                                   `lb:1,ub:2,optional`
	Dl_DataToUL_ACK                                   []int64                                              `lb:1,ub:8,e_lb:0,e_ub:15,optional`
	SpatialRelationInfoToAddModList                   []PUCCH_SpatialRelationInfo                          `lb:1,ub:maxNrofSpatialRelationInfos,optional`
	SpatialRelationInfoToReleaseList                  []PUCCH_SpatialRelationInfoId                        `lb:1,ub:maxNrofSpatialRelationInfos,optional`
	Pucch_PowerControl                                *PUCCH_PowerControl                                  `optional`
	ResourceToAddModListExt_v1610                     []PUCCH_ResourceExt_v1610                            `lb:1,ub:maxNrofPUCCH_Resources,optional,ext-1`
	Dl_DataToUL_ACK_r16                               *DL_DataToUL_ACK_r16                                 `optional,ext-1,setuprelease`
	Ul_AccessConfigListDCI_1_1_r16                    *UL_AccessConfigListDCI_1_1_r16                      `optional,ext-1,setuprelease`
	SubslotLengthForPUCCH_r16                         *PUCCH_Config_subslotLengthForPUCCH_r16              `optional,ext-1`
	Dl_DataToUL_ACK_DCI_1_2_r16                       *DL_DataToUL_ACK_DCI_1_2_r16                         `optional,ext-1,setuprelease`
	NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 *int64                                               `lb:0,ub:3,optional,ext-1`
	Dmrs_UplinkTransformPrecodingPUCCH_r16            *PUCCH_Config_dmrs_UplinkTransformPrecodingPUCCH_r16 `optional,ext-1`
	SpatialRelationInfoToAddModListSizeExt_v1610      []PUCCH_SpatialRelationInfo                          `lb:1,ub:maxNrofSpatialRelationInfosDiff_r16,optional,ext-1`
	SpatialRelationInfoToReleaseListSizeExt_v1610     []PUCCH_SpatialRelationInfoId                        `lb:1,ub:maxNrofSpatialRelationInfosDiff_r16,optional,ext-1`
	SpatialRelationInfoToAddModListExt_v1610          []PUCCH_SpatialRelationInfoExt_r16                   `lb:1,ub:maxNrofSpatialRelationInfos_r16,optional,ext-1`
	SpatialRelationInfoToReleaseListExt_v1610         []PUCCH_SpatialRelationInfoId_r16                    `lb:1,ub:maxNrofSpatialRelationInfos_r16,optional,ext-1`
	ResourceGroupToAddModList_r16                     []PUCCH_ResourceGroup_r16                            `lb:1,ub:maxNrofPUCCH_ResourceGroups_r16,optional,ext-1`
	ResourceGroupToReleaseList_r16                    []PUCCH_ResourceGroupId_r16                          `lb:1,ub:maxNrofPUCCH_ResourceGroups_r16,optional,ext-1`
	Sps_PUCCH_AN_List_r16                             *SPS_PUCCH_AN_List_r16                               `optional,ext-1,setuprelease`
	SchedulingRequestResourceToAddModListExt_v1610    []SchedulingRequestResourceConfigExt_v1610           `lb:1,ub:maxNrofSR_Resources,optional,ext-1`
	Format0_r17                                       *PUCCH_FormatConfig                                  `optional,ext-2,setuprelease`
	Format2Ext_r17                                    *PUCCH_FormatConfigExt_r17                           `optional,ext-2,setuprelease`
	Format3Ext_r17                                    *PUCCH_FormatConfigExt_r17                           `optional,ext-2,setuprelease`
	Format4Ext_r17                                    *PUCCH_FormatConfigExt_r17                           `optional,ext-2,setuprelease`
	Ul_AccessConfigListDCI_1_2_r17                    *UL_AccessConfigListDCI_1_2_r17                      `optional,ext-2,setuprelease`
	MappingPattern_r17                                *PUCCH_Config_mappingPattern_r17                     `optional,ext-2`
	PowerControlSetInfoToAddModList_r17               []PUCCH_PowerControlSetInfo_r17                      `lb:1,ub:maxNrofPowerControlSetInfos_r17,optional,ext-2`
	PowerControlSetInfoToReleaseList_r17              []PUCCH_PowerControlSetInfoId_r17                    `lb:1,ub:maxNrofPowerControlSetInfos_r17,optional,ext-2`
	SecondTPCFieldDCI_1_1_r17                         *PUCCH_Config_secondTPCFieldDCI_1_1_r17              `optional,ext-2`
	SecondTPCFieldDCI_1_2_r17                         *PUCCH_Config_secondTPCFieldDCI_1_2_r17              `optional,ext-2`
	Dl_DataToUL_ACK_r17                               *DL_DataToUL_ACK_r17                                 `optional,ext-2,setuprelease`
	Dl_DataToUL_ACK_DCI_1_2_r17                       *DL_DataToUL_ACK_DCI_1_2_r17                         `optional,ext-2,setuprelease`
	Ul_AccessConfigListDCI_1_1_r17                    *UL_AccessConfigListDCI_1_1_r17                      `optional,ext-2,setuprelease`
	SchedulingRequestResourceToAddModListExt_v1700    []SchedulingRequestResourceConfigExt_v1700           `lb:1,ub:maxNrofSR_Resources,optional,ext-2`
	Dmrs_BundlingPUCCH_Config_r17                     *DMRS_BundlingPUCCH_Config_r17                       `optional,ext-2,setuprelease`
	Dl_DataToUL_ACK_v1700                             *DL_DataToUL_ACK_v1700                               `optional,ext-2,setuprelease`
	Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17        *DL_DataToUL_ACK_MulticastDCI_Format4_1_r17          `optional,ext-2,setuprelease`
	Sps_PUCCH_AN_ListMulticast_r17                    *SPS_PUCCH_AN_List_r16                               `optional,ext-2,setuprelease`
}

func (ie *PUCCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.ResourceToAddModListExt_v1610) > 0 || ie.Dl_DataToUL_ACK_r16 != nil || ie.Ul_AccessConfigListDCI_1_1_r16 != nil || ie.SubslotLengthForPUCCH_r16 != nil || ie.Dl_DataToUL_ACK_DCI_1_2_r16 != nil || ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil || ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 != nil || len(ie.SpatialRelationInfoToAddModListSizeExt_v1610) > 0 || len(ie.SpatialRelationInfoToReleaseListSizeExt_v1610) > 0 || len(ie.SpatialRelationInfoToAddModListExt_v1610) > 0 || len(ie.SpatialRelationInfoToReleaseListExt_v1610) > 0 || len(ie.ResourceGroupToAddModList_r16) > 0 || len(ie.ResourceGroupToReleaseList_r16) > 0 || ie.Sps_PUCCH_AN_List_r16 != nil || len(ie.SchedulingRequestResourceToAddModListExt_v1610) > 0 || ie.Format0_r17 != nil || ie.Format2Ext_r17 != nil || ie.Format3Ext_r17 != nil || ie.Format4Ext_r17 != nil || ie.Ul_AccessConfigListDCI_1_2_r17 != nil || ie.MappingPattern_r17 != nil || len(ie.PowerControlSetInfoToAddModList_r17) > 0 || len(ie.PowerControlSetInfoToReleaseList_r17) > 0 || ie.SecondTPCFieldDCI_1_1_r17 != nil || ie.SecondTPCFieldDCI_1_2_r17 != nil || ie.Dl_DataToUL_ACK_r17 != nil || ie.Dl_DataToUL_ACK_DCI_1_2_r17 != nil || ie.Ul_AccessConfigListDCI_1_1_r17 != nil || len(ie.SchedulingRequestResourceToAddModListExt_v1700) > 0 || ie.Dmrs_BundlingPUCCH_Config_r17 != nil || ie.Dl_DataToUL_ACK_v1700 != nil || ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil || ie.Sps_PUCCH_AN_ListMulticast_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.ResourceSetToAddModList) > 0, len(ie.ResourceSetToReleaseList) > 0, len(ie.ResourceToAddModList) > 0, len(ie.ResourceToReleaseList) > 0, ie.Format1 != nil, ie.Format2 != nil, ie.Format3 != nil, ie.Format4 != nil, len(ie.SchedulingRequestResourceToAddModList) > 0, len(ie.SchedulingRequestResourceToReleaseList) > 0, len(ie.Multi_CSI_PUCCH_ResourceList) > 0, len(ie.Dl_DataToUL_ACK) > 0, len(ie.SpatialRelationInfoToAddModList) > 0, len(ie.SpatialRelationInfoToReleaseList) > 0, ie.Pucch_PowerControl != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.ResourceSetToAddModList) > 0 {
		tmp_ResourceSetToAddModList := utils.NewSequence[*PUCCH_ResourceSet]([]*PUCCH_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		for _, i := range ie.ResourceSetToAddModList {
			tmp_ResourceSetToAddModList.Value = append(tmp_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceSetToAddModList", err)
		}
	}
	if len(ie.ResourceSetToReleaseList) > 0 {
		tmp_ResourceSetToReleaseList := utils.NewSequence[*PUCCH_ResourceSetId]([]*PUCCH_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		for _, i := range ie.ResourceSetToReleaseList {
			tmp_ResourceSetToReleaseList.Value = append(tmp_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceSetToReleaseList", err)
		}
	}
	if len(ie.ResourceToAddModList) > 0 {
		tmp_ResourceToAddModList := utils.NewSequence[*PUCCH_Resource]([]*PUCCH_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		for _, i := range ie.ResourceToAddModList {
			tmp_ResourceToAddModList.Value = append(tmp_ResourceToAddModList.Value, &i)
		}
		if err = tmp_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceToAddModList", err)
		}
	}
	if len(ie.ResourceToReleaseList) > 0 {
		tmp_ResourceToReleaseList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		for _, i := range ie.ResourceToReleaseList {
			tmp_ResourceToReleaseList.Value = append(tmp_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceToReleaseList", err)
		}
	}
	if ie.Format1 != nil {
		tmp_Format1 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.Format1,
		}
		if err = tmp_Format1.Encode(w); err != nil {
			return utils.WrapError("Encode Format1", err)
		}
	}
	if ie.Format2 != nil {
		tmp_Format2 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.Format2,
		}
		if err = tmp_Format2.Encode(w); err != nil {
			return utils.WrapError("Encode Format2", err)
		}
	}
	if ie.Format3 != nil {
		tmp_Format3 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.Format3,
		}
		if err = tmp_Format3.Encode(w); err != nil {
			return utils.WrapError("Encode Format3", err)
		}
	}
	if ie.Format4 != nil {
		tmp_Format4 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.Format4,
		}
		if err = tmp_Format4.Encode(w); err != nil {
			return utils.WrapError("Encode Format4", err)
		}
	}
	if len(ie.SchedulingRequestResourceToAddModList) > 0 {
		tmp_SchedulingRequestResourceToAddModList := utils.NewSequence[*SchedulingRequestResourceConfig]([]*SchedulingRequestResourceConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		for _, i := range ie.SchedulingRequestResourceToAddModList {
			tmp_SchedulingRequestResourceToAddModList.Value = append(tmp_SchedulingRequestResourceToAddModList.Value, &i)
		}
		if err = tmp_SchedulingRequestResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode SchedulingRequestResourceToAddModList", err)
		}
	}
	if len(ie.SchedulingRequestResourceToReleaseList) > 0 {
		tmp_SchedulingRequestResourceToReleaseList := utils.NewSequence[*SchedulingRequestResourceId]([]*SchedulingRequestResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		for _, i := range ie.SchedulingRequestResourceToReleaseList {
			tmp_SchedulingRequestResourceToReleaseList.Value = append(tmp_SchedulingRequestResourceToReleaseList.Value, &i)
		}
		if err = tmp_SchedulingRequestResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode SchedulingRequestResourceToReleaseList", err)
		}
	}
	if len(ie.Multi_CSI_PUCCH_ResourceList) > 0 {
		tmp_Multi_CSI_PUCCH_ResourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: 2}, false)
		for _, i := range ie.Multi_CSI_PUCCH_ResourceList {
			tmp_Multi_CSI_PUCCH_ResourceList.Value = append(tmp_Multi_CSI_PUCCH_ResourceList.Value, &i)
		}
		if err = tmp_Multi_CSI_PUCCH_ResourceList.Encode(w); err != nil {
			return utils.WrapError("Encode Multi_CSI_PUCCH_ResourceList", err)
		}
	}
	if len(ie.Dl_DataToUL_ACK) > 0 {
		tmp_Dl_DataToUL_ACK := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.Dl_DataToUL_ACK {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 15}, false)
			tmp_Dl_DataToUL_ACK.Value = append(tmp_Dl_DataToUL_ACK.Value, &tmp_ie)
		}
		if err = tmp_Dl_DataToUL_ACK.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_DataToUL_ACK", err)
		}
	}
	if len(ie.SpatialRelationInfoToAddModList) > 0 {
		tmp_SpatialRelationInfoToAddModList := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		for _, i := range ie.SpatialRelationInfoToAddModList {
			tmp_SpatialRelationInfoToAddModList.Value = append(tmp_SpatialRelationInfoToAddModList.Value, &i)
		}
		if err = tmp_SpatialRelationInfoToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelationInfoToAddModList", err)
		}
	}
	if len(ie.SpatialRelationInfoToReleaseList) > 0 {
		tmp_SpatialRelationInfoToReleaseList := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		for _, i := range ie.SpatialRelationInfoToReleaseList {
			tmp_SpatialRelationInfoToReleaseList.Value = append(tmp_SpatialRelationInfoToReleaseList.Value, &i)
		}
		if err = tmp_SpatialRelationInfoToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelationInfoToReleaseList", err)
		}
	}
	if ie.Pucch_PowerControl != nil {
		if err = ie.Pucch_PowerControl.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_PowerControl", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{len(ie.ResourceToAddModListExt_v1610) > 0 || ie.Dl_DataToUL_ACK_r16 != nil || ie.Ul_AccessConfigListDCI_1_1_r16 != nil || ie.SubslotLengthForPUCCH_r16 != nil || ie.Dl_DataToUL_ACK_DCI_1_2_r16 != nil || ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil || ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 != nil || len(ie.SpatialRelationInfoToAddModListSizeExt_v1610) > 0 || len(ie.SpatialRelationInfoToReleaseListSizeExt_v1610) > 0 || len(ie.SpatialRelationInfoToAddModListExt_v1610) > 0 || len(ie.SpatialRelationInfoToReleaseListExt_v1610) > 0 || len(ie.ResourceGroupToAddModList_r16) > 0 || len(ie.ResourceGroupToReleaseList_r16) > 0 || ie.Sps_PUCCH_AN_List_r16 != nil || len(ie.SchedulingRequestResourceToAddModListExt_v1610) > 0, ie.Format0_r17 != nil || ie.Format2Ext_r17 != nil || ie.Format3Ext_r17 != nil || ie.Format4Ext_r17 != nil || ie.Ul_AccessConfigListDCI_1_2_r17 != nil || ie.MappingPattern_r17 != nil || len(ie.PowerControlSetInfoToAddModList_r17) > 0 || len(ie.PowerControlSetInfoToReleaseList_r17) > 0 || ie.SecondTPCFieldDCI_1_1_r17 != nil || ie.SecondTPCFieldDCI_1_2_r17 != nil || ie.Dl_DataToUL_ACK_r17 != nil || ie.Dl_DataToUL_ACK_DCI_1_2_r17 != nil || ie.Ul_AccessConfigListDCI_1_1_r17 != nil || len(ie.SchedulingRequestResourceToAddModListExt_v1700) > 0 || ie.Dmrs_BundlingPUCCH_Config_r17 != nil || ie.Dl_DataToUL_ACK_v1700 != nil || ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil || ie.Sps_PUCCH_AN_ListMulticast_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.ResourceToAddModListExt_v1610) > 0, ie.Dl_DataToUL_ACK_r16 != nil, ie.Ul_AccessConfigListDCI_1_1_r16 != nil, ie.SubslotLengthForPUCCH_r16 != nil, ie.Dl_DataToUL_ACK_DCI_1_2_r16 != nil, ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil, ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 != nil, len(ie.SpatialRelationInfoToAddModListSizeExt_v1610) > 0, len(ie.SpatialRelationInfoToReleaseListSizeExt_v1610) > 0, len(ie.SpatialRelationInfoToAddModListExt_v1610) > 0, len(ie.SpatialRelationInfoToReleaseListExt_v1610) > 0, len(ie.ResourceGroupToAddModList_r16) > 0, len(ie.ResourceGroupToReleaseList_r16) > 0, ie.Sps_PUCCH_AN_List_r16 != nil, len(ie.SchedulingRequestResourceToAddModListExt_v1610) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ResourceToAddModListExt_v1610 optional
			if len(ie.ResourceToAddModListExt_v1610) > 0 {
				tmp_ResourceToAddModListExt_v1610 := utils.NewSequence[*PUCCH_ResourceExt_v1610]([]*PUCCH_ResourceExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
				for _, i := range ie.ResourceToAddModListExt_v1610 {
					tmp_ResourceToAddModListExt_v1610.Value = append(tmp_ResourceToAddModListExt_v1610.Value, &i)
				}
				if err = tmp_ResourceToAddModListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceToAddModListExt_v1610", err)
				}
			}
			// encode Dl_DataToUL_ACK_r16 optional
			if ie.Dl_DataToUL_ACK_r16 != nil {
				tmp_Dl_DataToUL_ACK_r16 := utils.SetupRelease[*DL_DataToUL_ACK_r16]{
					Setup: ie.Dl_DataToUL_ACK_r16,
				}
				if err = tmp_Dl_DataToUL_ACK_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_DataToUL_ACK_r16", err)
				}
			}
			// encode Ul_AccessConfigListDCI_1_1_r16 optional
			if ie.Ul_AccessConfigListDCI_1_1_r16 != nil {
				tmp_Ul_AccessConfigListDCI_1_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r16]{
					Setup: ie.Ul_AccessConfigListDCI_1_1_r16,
				}
				if err = tmp_Ul_AccessConfigListDCI_1_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_AccessConfigListDCI_1_1_r16", err)
				}
			}
			// encode SubslotLengthForPUCCH_r16 optional
			if ie.SubslotLengthForPUCCH_r16 != nil {
				if err = ie.SubslotLengthForPUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SubslotLengthForPUCCH_r16", err)
				}
			}
			// encode Dl_DataToUL_ACK_DCI_1_2_r16 optional
			if ie.Dl_DataToUL_ACK_DCI_1_2_r16 != nil {
				tmp_Dl_DataToUL_ACK_DCI_1_2_r16 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r16]{
					Setup: ie.Dl_DataToUL_ACK_DCI_1_2_r16,
				}
				if err = tmp_Dl_DataToUL_ACK_DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_DataToUL_ACK_DCI_1_2_r16", err)
				}
			}
			// encode NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 optional
			if ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16", err)
				}
			}
			// encode Dmrs_UplinkTransformPrecodingPUCCH_r16 optional
			if ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 != nil {
				if err = ie.Dmrs_UplinkTransformPrecodingPUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_UplinkTransformPrecodingPUCCH_r16", err)
				}
			}
			// encode SpatialRelationInfoToAddModListSizeExt_v1610 optional
			if len(ie.SpatialRelationInfoToAddModListSizeExt_v1610) > 0 {
				tmp_SpatialRelationInfoToAddModListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				for _, i := range ie.SpatialRelationInfoToAddModListSizeExt_v1610 {
					tmp_SpatialRelationInfoToAddModListSizeExt_v1610.Value = append(tmp_SpatialRelationInfoToAddModListSizeExt_v1610.Value, &i)
				}
				if err = tmp_SpatialRelationInfoToAddModListSizeExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationInfoToAddModListSizeExt_v1610", err)
				}
			}
			// encode SpatialRelationInfoToReleaseListSizeExt_v1610 optional
			if len(ie.SpatialRelationInfoToReleaseListSizeExt_v1610) > 0 {
				tmp_SpatialRelationInfoToReleaseListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				for _, i := range ie.SpatialRelationInfoToReleaseListSizeExt_v1610 {
					tmp_SpatialRelationInfoToReleaseListSizeExt_v1610.Value = append(tmp_SpatialRelationInfoToReleaseListSizeExt_v1610.Value, &i)
				}
				if err = tmp_SpatialRelationInfoToReleaseListSizeExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationInfoToReleaseListSizeExt_v1610", err)
				}
			}
			// encode SpatialRelationInfoToAddModListExt_v1610 optional
			if len(ie.SpatialRelationInfoToAddModListExt_v1610) > 0 {
				tmp_SpatialRelationInfoToAddModListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoExt_r16]([]*PUCCH_SpatialRelationInfoExt_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				for _, i := range ie.SpatialRelationInfoToAddModListExt_v1610 {
					tmp_SpatialRelationInfoToAddModListExt_v1610.Value = append(tmp_SpatialRelationInfoToAddModListExt_v1610.Value, &i)
				}
				if err = tmp_SpatialRelationInfoToAddModListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationInfoToAddModListExt_v1610", err)
				}
			}
			// encode SpatialRelationInfoToReleaseListExt_v1610 optional
			if len(ie.SpatialRelationInfoToReleaseListExt_v1610) > 0 {
				tmp_SpatialRelationInfoToReleaseListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId_r16]([]*PUCCH_SpatialRelationInfoId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				for _, i := range ie.SpatialRelationInfoToReleaseListExt_v1610 {
					tmp_SpatialRelationInfoToReleaseListExt_v1610.Value = append(tmp_SpatialRelationInfoToReleaseListExt_v1610.Value, &i)
				}
				if err = tmp_SpatialRelationInfoToReleaseListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationInfoToReleaseListExt_v1610", err)
				}
			}
			// encode ResourceGroupToAddModList_r16 optional
			if len(ie.ResourceGroupToAddModList_r16) > 0 {
				tmp_ResourceGroupToAddModList_r16 := utils.NewSequence[*PUCCH_ResourceGroup_r16]([]*PUCCH_ResourceGroup_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				for _, i := range ie.ResourceGroupToAddModList_r16 {
					tmp_ResourceGroupToAddModList_r16.Value = append(tmp_ResourceGroupToAddModList_r16.Value, &i)
				}
				if err = tmp_ResourceGroupToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceGroupToAddModList_r16", err)
				}
			}
			// encode ResourceGroupToReleaseList_r16 optional
			if len(ie.ResourceGroupToReleaseList_r16) > 0 {
				tmp_ResourceGroupToReleaseList_r16 := utils.NewSequence[*PUCCH_ResourceGroupId_r16]([]*PUCCH_ResourceGroupId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				for _, i := range ie.ResourceGroupToReleaseList_r16 {
					tmp_ResourceGroupToReleaseList_r16.Value = append(tmp_ResourceGroupToReleaseList_r16.Value, &i)
				}
				if err = tmp_ResourceGroupToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourceGroupToReleaseList_r16", err)
				}
			}
			// encode Sps_PUCCH_AN_List_r16 optional
			if ie.Sps_PUCCH_AN_List_r16 != nil {
				tmp_Sps_PUCCH_AN_List_r16 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{
					Setup: ie.Sps_PUCCH_AN_List_r16,
				}
				if err = tmp_Sps_PUCCH_AN_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_PUCCH_AN_List_r16", err)
				}
			}
			// encode SchedulingRequestResourceToAddModListExt_v1610 optional
			if len(ie.SchedulingRequestResourceToAddModListExt_v1610) > 0 {
				tmp_SchedulingRequestResourceToAddModListExt_v1610 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1610]([]*SchedulingRequestResourceConfigExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				for _, i := range ie.SchedulingRequestResourceToAddModListExt_v1610 {
					tmp_SchedulingRequestResourceToAddModListExt_v1610.Value = append(tmp_SchedulingRequestResourceToAddModListExt_v1610.Value, &i)
				}
				if err = tmp_SchedulingRequestResourceToAddModListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestResourceToAddModListExt_v1610", err)
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
			optionals_ext_2 := []bool{ie.Format0_r17 != nil, ie.Format2Ext_r17 != nil, ie.Format3Ext_r17 != nil, ie.Format4Ext_r17 != nil, ie.Ul_AccessConfigListDCI_1_2_r17 != nil, ie.MappingPattern_r17 != nil, len(ie.PowerControlSetInfoToAddModList_r17) > 0, len(ie.PowerControlSetInfoToReleaseList_r17) > 0, ie.SecondTPCFieldDCI_1_1_r17 != nil, ie.SecondTPCFieldDCI_1_2_r17 != nil, ie.Dl_DataToUL_ACK_r17 != nil, ie.Dl_DataToUL_ACK_DCI_1_2_r17 != nil, ie.Ul_AccessConfigListDCI_1_1_r17 != nil, len(ie.SchedulingRequestResourceToAddModListExt_v1700) > 0, ie.Dmrs_BundlingPUCCH_Config_r17 != nil, ie.Dl_DataToUL_ACK_v1700 != nil, ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil, ie.Sps_PUCCH_AN_ListMulticast_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Format0_r17 optional
			if ie.Format0_r17 != nil {
				tmp_Format0_r17 := utils.SetupRelease[*PUCCH_FormatConfig]{
					Setup: ie.Format0_r17,
				}
				if err = tmp_Format0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Format0_r17", err)
				}
			}
			// encode Format2Ext_r17 optional
			if ie.Format2Ext_r17 != nil {
				tmp_Format2Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{
					Setup: ie.Format2Ext_r17,
				}
				if err = tmp_Format2Ext_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Format2Ext_r17", err)
				}
			}
			// encode Format3Ext_r17 optional
			if ie.Format3Ext_r17 != nil {
				tmp_Format3Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{
					Setup: ie.Format3Ext_r17,
				}
				if err = tmp_Format3Ext_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Format3Ext_r17", err)
				}
			}
			// encode Format4Ext_r17 optional
			if ie.Format4Ext_r17 != nil {
				tmp_Format4Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{
					Setup: ie.Format4Ext_r17,
				}
				if err = tmp_Format4Ext_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Format4Ext_r17", err)
				}
			}
			// encode Ul_AccessConfigListDCI_1_2_r17 optional
			if ie.Ul_AccessConfigListDCI_1_2_r17 != nil {
				tmp_Ul_AccessConfigListDCI_1_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_2_r17]{
					Setup: ie.Ul_AccessConfigListDCI_1_2_r17,
				}
				if err = tmp_Ul_AccessConfigListDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_AccessConfigListDCI_1_2_r17", err)
				}
			}
			// encode MappingPattern_r17 optional
			if ie.MappingPattern_r17 != nil {
				if err = ie.MappingPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MappingPattern_r17", err)
				}
			}
			// encode PowerControlSetInfoToAddModList_r17 optional
			if len(ie.PowerControlSetInfoToAddModList_r17) > 0 {
				tmp_PowerControlSetInfoToAddModList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfo_r17]([]*PUCCH_PowerControlSetInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				for _, i := range ie.PowerControlSetInfoToAddModList_r17 {
					tmp_PowerControlSetInfoToAddModList_r17.Value = append(tmp_PowerControlSetInfoToAddModList_r17.Value, &i)
				}
				if err = tmp_PowerControlSetInfoToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PowerControlSetInfoToAddModList_r17", err)
				}
			}
			// encode PowerControlSetInfoToReleaseList_r17 optional
			if len(ie.PowerControlSetInfoToReleaseList_r17) > 0 {
				tmp_PowerControlSetInfoToReleaseList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfoId_r17]([]*PUCCH_PowerControlSetInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				for _, i := range ie.PowerControlSetInfoToReleaseList_r17 {
					tmp_PowerControlSetInfoToReleaseList_r17.Value = append(tmp_PowerControlSetInfoToReleaseList_r17.Value, &i)
				}
				if err = tmp_PowerControlSetInfoToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PowerControlSetInfoToReleaseList_r17", err)
				}
			}
			// encode SecondTPCFieldDCI_1_1_r17 optional
			if ie.SecondTPCFieldDCI_1_1_r17 != nil {
				if err = ie.SecondTPCFieldDCI_1_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SecondTPCFieldDCI_1_1_r17", err)
				}
			}
			// encode SecondTPCFieldDCI_1_2_r17 optional
			if ie.SecondTPCFieldDCI_1_2_r17 != nil {
				if err = ie.SecondTPCFieldDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SecondTPCFieldDCI_1_2_r17", err)
				}
			}
			// encode Dl_DataToUL_ACK_r17 optional
			if ie.Dl_DataToUL_ACK_r17 != nil {
				tmp_Dl_DataToUL_ACK_r17 := utils.SetupRelease[*DL_DataToUL_ACK_r17]{
					Setup: ie.Dl_DataToUL_ACK_r17,
				}
				if err = tmp_Dl_DataToUL_ACK_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_DataToUL_ACK_r17", err)
				}
			}
			// encode Dl_DataToUL_ACK_DCI_1_2_r17 optional
			if ie.Dl_DataToUL_ACK_DCI_1_2_r17 != nil {
				tmp_Dl_DataToUL_ACK_DCI_1_2_r17 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r17]{
					Setup: ie.Dl_DataToUL_ACK_DCI_1_2_r17,
				}
				if err = tmp_Dl_DataToUL_ACK_DCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_DataToUL_ACK_DCI_1_2_r17", err)
				}
			}
			// encode Ul_AccessConfigListDCI_1_1_r17 optional
			if ie.Ul_AccessConfigListDCI_1_1_r17 != nil {
				tmp_Ul_AccessConfigListDCI_1_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r17]{
					Setup: ie.Ul_AccessConfigListDCI_1_1_r17,
				}
				if err = tmp_Ul_AccessConfigListDCI_1_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_AccessConfigListDCI_1_1_r17", err)
				}
			}
			// encode SchedulingRequestResourceToAddModListExt_v1700 optional
			if len(ie.SchedulingRequestResourceToAddModListExt_v1700) > 0 {
				tmp_SchedulingRequestResourceToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1700]([]*SchedulingRequestResourceConfigExt_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				for _, i := range ie.SchedulingRequestResourceToAddModListExt_v1700 {
					tmp_SchedulingRequestResourceToAddModListExt_v1700.Value = append(tmp_SchedulingRequestResourceToAddModListExt_v1700.Value, &i)
				}
				if err = tmp_SchedulingRequestResourceToAddModListExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestResourceToAddModListExt_v1700", err)
				}
			}
			// encode Dmrs_BundlingPUCCH_Config_r17 optional
			if ie.Dmrs_BundlingPUCCH_Config_r17 != nil {
				tmp_Dmrs_BundlingPUCCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUCCH_Config_r17]{
					Setup: ie.Dmrs_BundlingPUCCH_Config_r17,
				}
				if err = tmp_Dmrs_BundlingPUCCH_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingPUCCH_Config_r17", err)
				}
			}
			// encode Dl_DataToUL_ACK_v1700 optional
			if ie.Dl_DataToUL_ACK_v1700 != nil {
				tmp_Dl_DataToUL_ACK_v1700 := utils.SetupRelease[*DL_DataToUL_ACK_v1700]{
					Setup: ie.Dl_DataToUL_ACK_v1700,
				}
				if err = tmp_Dl_DataToUL_ACK_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_DataToUL_ACK_v1700", err)
				}
			}
			// encode Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 optional
			if ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil {
				tmp_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 := utils.SetupRelease[*DL_DataToUL_ACK_MulticastDCI_Format4_1_r17]{
					Setup: ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17,
				}
				if err = tmp_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17", err)
				}
			}
			// encode Sps_PUCCH_AN_ListMulticast_r17 optional
			if ie.Sps_PUCCH_AN_ListMulticast_r17 != nil {
				tmp_Sps_PUCCH_AN_ListMulticast_r17 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{
					Setup: ie.Sps_PUCCH_AN_ListMulticast_r17,
				}
				if err = tmp_Sps_PUCCH_AN_ListMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_PUCCH_AN_ListMulticast_r17", err)
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

func (ie *PUCCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceSetToAddModListPresent bool
	if ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceSetToReleaseListPresent bool
	if ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceToAddModListPresent bool
	if ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceToReleaseListPresent bool
	if ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Format1Present bool
	if Format1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Format2Present bool
	if Format2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Format3Present bool
	if Format3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Format4Present bool
	if Format4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SchedulingRequestResourceToAddModListPresent bool
	if SchedulingRequestResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SchedulingRequestResourceToReleaseListPresent bool
	if SchedulingRequestResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Multi_CSI_PUCCH_ResourceListPresent bool
	if Multi_CSI_PUCCH_ResourceListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_DataToUL_ACKPresent bool
	if Dl_DataToUL_ACKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialRelationInfoToAddModListPresent bool
	if SpatialRelationInfoToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialRelationInfoToReleaseListPresent bool
	if SpatialRelationInfoToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_PowerControlPresent bool
	if Pucch_PowerControlPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ResourceSetToAddModListPresent {
		tmp_ResourceSetToAddModList := utils.NewSequence[*PUCCH_ResourceSet]([]*PUCCH_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		fn_ResourceSetToAddModList := func() *PUCCH_ResourceSet {
			return new(PUCCH_ResourceSet)
		}
		if err = tmp_ResourceSetToAddModList.Decode(r, fn_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode ResourceSetToAddModList", err)
		}
		ie.ResourceSetToAddModList = []PUCCH_ResourceSet{}
		for _, i := range tmp_ResourceSetToAddModList.Value {
			ie.ResourceSetToAddModList = append(ie.ResourceSetToAddModList, *i)
		}
	}
	if ResourceSetToReleaseListPresent {
		tmp_ResourceSetToReleaseList := utils.NewSequence[*PUCCH_ResourceSetId]([]*PUCCH_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		fn_ResourceSetToReleaseList := func() *PUCCH_ResourceSetId {
			return new(PUCCH_ResourceSetId)
		}
		if err = tmp_ResourceSetToReleaseList.Decode(r, fn_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode ResourceSetToReleaseList", err)
		}
		ie.ResourceSetToReleaseList = []PUCCH_ResourceSetId{}
		for _, i := range tmp_ResourceSetToReleaseList.Value {
			ie.ResourceSetToReleaseList = append(ie.ResourceSetToReleaseList, *i)
		}
	}
	if ResourceToAddModListPresent {
		tmp_ResourceToAddModList := utils.NewSequence[*PUCCH_Resource]([]*PUCCH_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		fn_ResourceToAddModList := func() *PUCCH_Resource {
			return new(PUCCH_Resource)
		}
		if err = tmp_ResourceToAddModList.Decode(r, fn_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode ResourceToAddModList", err)
		}
		ie.ResourceToAddModList = []PUCCH_Resource{}
		for _, i := range tmp_ResourceToAddModList.Value {
			ie.ResourceToAddModList = append(ie.ResourceToAddModList, *i)
		}
	}
	if ResourceToReleaseListPresent {
		tmp_ResourceToReleaseList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		fn_ResourceToReleaseList := func() *PUCCH_ResourceId {
			return new(PUCCH_ResourceId)
		}
		if err = tmp_ResourceToReleaseList.Decode(r, fn_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode ResourceToReleaseList", err)
		}
		ie.ResourceToReleaseList = []PUCCH_ResourceId{}
		for _, i := range tmp_ResourceToReleaseList.Value {
			ie.ResourceToReleaseList = append(ie.ResourceToReleaseList, *i)
		}
	}
	if Format1Present {
		tmp_Format1 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_Format1.Decode(r); err != nil {
			return utils.WrapError("Decode Format1", err)
		}
		ie.Format1 = tmp_Format1.Setup
	}
	if Format2Present {
		tmp_Format2 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_Format2.Decode(r); err != nil {
			return utils.WrapError("Decode Format2", err)
		}
		ie.Format2 = tmp_Format2.Setup
	}
	if Format3Present {
		tmp_Format3 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_Format3.Decode(r); err != nil {
			return utils.WrapError("Decode Format3", err)
		}
		ie.Format3 = tmp_Format3.Setup
	}
	if Format4Present {
		tmp_Format4 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_Format4.Decode(r); err != nil {
			return utils.WrapError("Decode Format4", err)
		}
		ie.Format4 = tmp_Format4.Setup
	}
	if SchedulingRequestResourceToAddModListPresent {
		tmp_SchedulingRequestResourceToAddModList := utils.NewSequence[*SchedulingRequestResourceConfig]([]*SchedulingRequestResourceConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		fn_SchedulingRequestResourceToAddModList := func() *SchedulingRequestResourceConfig {
			return new(SchedulingRequestResourceConfig)
		}
		if err = tmp_SchedulingRequestResourceToAddModList.Decode(r, fn_SchedulingRequestResourceToAddModList); err != nil {
			return utils.WrapError("Decode SchedulingRequestResourceToAddModList", err)
		}
		ie.SchedulingRequestResourceToAddModList = []SchedulingRequestResourceConfig{}
		for _, i := range tmp_SchedulingRequestResourceToAddModList.Value {
			ie.SchedulingRequestResourceToAddModList = append(ie.SchedulingRequestResourceToAddModList, *i)
		}
	}
	if SchedulingRequestResourceToReleaseListPresent {
		tmp_SchedulingRequestResourceToReleaseList := utils.NewSequence[*SchedulingRequestResourceId]([]*SchedulingRequestResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		fn_SchedulingRequestResourceToReleaseList := func() *SchedulingRequestResourceId {
			return new(SchedulingRequestResourceId)
		}
		if err = tmp_SchedulingRequestResourceToReleaseList.Decode(r, fn_SchedulingRequestResourceToReleaseList); err != nil {
			return utils.WrapError("Decode SchedulingRequestResourceToReleaseList", err)
		}
		ie.SchedulingRequestResourceToReleaseList = []SchedulingRequestResourceId{}
		for _, i := range tmp_SchedulingRequestResourceToReleaseList.Value {
			ie.SchedulingRequestResourceToReleaseList = append(ie.SchedulingRequestResourceToReleaseList, *i)
		}
	}
	if Multi_CSI_PUCCH_ResourceListPresent {
		tmp_Multi_CSI_PUCCH_ResourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: 2}, false)
		fn_Multi_CSI_PUCCH_ResourceList := func() *PUCCH_ResourceId {
			return new(PUCCH_ResourceId)
		}
		if err = tmp_Multi_CSI_PUCCH_ResourceList.Decode(r, fn_Multi_CSI_PUCCH_ResourceList); err != nil {
			return utils.WrapError("Decode Multi_CSI_PUCCH_ResourceList", err)
		}
		ie.Multi_CSI_PUCCH_ResourceList = []PUCCH_ResourceId{}
		for _, i := range tmp_Multi_CSI_PUCCH_ResourceList.Value {
			ie.Multi_CSI_PUCCH_ResourceList = append(ie.Multi_CSI_PUCCH_ResourceList, *i)
		}
	}
	if Dl_DataToUL_ACKPresent {
		tmp_Dl_DataToUL_ACK := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_Dl_DataToUL_ACK := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 15}, false)
			return &ie
		}
		if err = tmp_Dl_DataToUL_ACK.Decode(r, fn_Dl_DataToUL_ACK); err != nil {
			return utils.WrapError("Decode Dl_DataToUL_ACK", err)
		}
		ie.Dl_DataToUL_ACK = []int64{}
		for _, i := range tmp_Dl_DataToUL_ACK.Value {
			ie.Dl_DataToUL_ACK = append(ie.Dl_DataToUL_ACK, int64(i.Value))
		}
	}
	if SpatialRelationInfoToAddModListPresent {
		tmp_SpatialRelationInfoToAddModList := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		fn_SpatialRelationInfoToAddModList := func() *PUCCH_SpatialRelationInfo {
			return new(PUCCH_SpatialRelationInfo)
		}
		if err = tmp_SpatialRelationInfoToAddModList.Decode(r, fn_SpatialRelationInfoToAddModList); err != nil {
			return utils.WrapError("Decode SpatialRelationInfoToAddModList", err)
		}
		ie.SpatialRelationInfoToAddModList = []PUCCH_SpatialRelationInfo{}
		for _, i := range tmp_SpatialRelationInfoToAddModList.Value {
			ie.SpatialRelationInfoToAddModList = append(ie.SpatialRelationInfoToAddModList, *i)
		}
	}
	if SpatialRelationInfoToReleaseListPresent {
		tmp_SpatialRelationInfoToReleaseList := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		fn_SpatialRelationInfoToReleaseList := func() *PUCCH_SpatialRelationInfoId {
			return new(PUCCH_SpatialRelationInfoId)
		}
		if err = tmp_SpatialRelationInfoToReleaseList.Decode(r, fn_SpatialRelationInfoToReleaseList); err != nil {
			return utils.WrapError("Decode SpatialRelationInfoToReleaseList", err)
		}
		ie.SpatialRelationInfoToReleaseList = []PUCCH_SpatialRelationInfoId{}
		for _, i := range tmp_SpatialRelationInfoToReleaseList.Value {
			ie.SpatialRelationInfoToReleaseList = append(ie.SpatialRelationInfoToReleaseList, *i)
		}
	}
	if Pucch_PowerControlPresent {
		ie.Pucch_PowerControl = new(PUCCH_PowerControl)
		if err = ie.Pucch_PowerControl.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_PowerControl", err)
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

			ResourceToAddModListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_DataToUL_ACK_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_AccessConfigListDCI_1_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SubslotLengthForPUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_DataToUL_ACK_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_UplinkTransformPrecodingPUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationInfoToAddModListSizeExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationInfoToReleaseListSizeExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationInfoToAddModListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationInfoToReleaseListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ResourceGroupToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ResourceGroupToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_PUCCH_AN_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SchedulingRequestResourceToAddModListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ResourceToAddModListExt_v1610 optional
			if ResourceToAddModListExt_v1610Present {
				tmp_ResourceToAddModListExt_v1610 := utils.NewSequence[*PUCCH_ResourceExt_v1610]([]*PUCCH_ResourceExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
				fn_ResourceToAddModListExt_v1610 := func() *PUCCH_ResourceExt_v1610 {
					return new(PUCCH_ResourceExt_v1610)
				}
				if err = tmp_ResourceToAddModListExt_v1610.Decode(extReader, fn_ResourceToAddModListExt_v1610); err != nil {
					return utils.WrapError("Decode ResourceToAddModListExt_v1610", err)
				}
				ie.ResourceToAddModListExt_v1610 = []PUCCH_ResourceExt_v1610{}
				for _, i := range tmp_ResourceToAddModListExt_v1610.Value {
					ie.ResourceToAddModListExt_v1610 = append(ie.ResourceToAddModListExt_v1610, *i)
				}
			}
			// decode Dl_DataToUL_ACK_r16 optional
			if Dl_DataToUL_ACK_r16Present {
				tmp_Dl_DataToUL_ACK_r16 := utils.SetupRelease[*DL_DataToUL_ACK_r16]{}
				if err = tmp_Dl_DataToUL_ACK_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_DataToUL_ACK_r16", err)
				}
				ie.Dl_DataToUL_ACK_r16 = tmp_Dl_DataToUL_ACK_r16.Setup
			}
			// decode Ul_AccessConfigListDCI_1_1_r16 optional
			if Ul_AccessConfigListDCI_1_1_r16Present {
				tmp_Ul_AccessConfigListDCI_1_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r16]{}
				if err = tmp_Ul_AccessConfigListDCI_1_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_AccessConfigListDCI_1_1_r16", err)
				}
				ie.Ul_AccessConfigListDCI_1_1_r16 = tmp_Ul_AccessConfigListDCI_1_1_r16.Setup
			}
			// decode SubslotLengthForPUCCH_r16 optional
			if SubslotLengthForPUCCH_r16Present {
				ie.SubslotLengthForPUCCH_r16 = new(PUCCH_Config_subslotLengthForPUCCH_r16)
				if err = ie.SubslotLengthForPUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SubslotLengthForPUCCH_r16", err)
				}
			}
			// decode Dl_DataToUL_ACK_DCI_1_2_r16 optional
			if Dl_DataToUL_ACK_DCI_1_2_r16Present {
				tmp_Dl_DataToUL_ACK_DCI_1_2_r16 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r16]{}
				if err = tmp_Dl_DataToUL_ACK_DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_DataToUL_ACK_DCI_1_2_r16", err)
				}
				ie.Dl_DataToUL_ACK_DCI_1_2_r16 = tmp_Dl_DataToUL_ACK_DCI_1_2_r16.Setup
			}
			// decode NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 optional
			if NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16Present {
				var tmp_int_NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 int64
				if tmp_int_NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Decode NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16", err)
				}
				ie.NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 = &tmp_int_NumberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16
			}
			// decode Dmrs_UplinkTransformPrecodingPUCCH_r16 optional
			if Dmrs_UplinkTransformPrecodingPUCCH_r16Present {
				ie.Dmrs_UplinkTransformPrecodingPUCCH_r16 = new(PUCCH_Config_dmrs_UplinkTransformPrecodingPUCCH_r16)
				if err = ie.Dmrs_UplinkTransformPrecodingPUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_UplinkTransformPrecodingPUCCH_r16", err)
				}
			}
			// decode SpatialRelationInfoToAddModListSizeExt_v1610 optional
			if SpatialRelationInfoToAddModListSizeExt_v1610Present {
				tmp_SpatialRelationInfoToAddModListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				fn_SpatialRelationInfoToAddModListSizeExt_v1610 := func() *PUCCH_SpatialRelationInfo {
					return new(PUCCH_SpatialRelationInfo)
				}
				if err = tmp_SpatialRelationInfoToAddModListSizeExt_v1610.Decode(extReader, fn_SpatialRelationInfoToAddModListSizeExt_v1610); err != nil {
					return utils.WrapError("Decode SpatialRelationInfoToAddModListSizeExt_v1610", err)
				}
				ie.SpatialRelationInfoToAddModListSizeExt_v1610 = []PUCCH_SpatialRelationInfo{}
				for _, i := range tmp_SpatialRelationInfoToAddModListSizeExt_v1610.Value {
					ie.SpatialRelationInfoToAddModListSizeExt_v1610 = append(ie.SpatialRelationInfoToAddModListSizeExt_v1610, *i)
				}
			}
			// decode SpatialRelationInfoToReleaseListSizeExt_v1610 optional
			if SpatialRelationInfoToReleaseListSizeExt_v1610Present {
				tmp_SpatialRelationInfoToReleaseListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				fn_SpatialRelationInfoToReleaseListSizeExt_v1610 := func() *PUCCH_SpatialRelationInfoId {
					return new(PUCCH_SpatialRelationInfoId)
				}
				if err = tmp_SpatialRelationInfoToReleaseListSizeExt_v1610.Decode(extReader, fn_SpatialRelationInfoToReleaseListSizeExt_v1610); err != nil {
					return utils.WrapError("Decode SpatialRelationInfoToReleaseListSizeExt_v1610", err)
				}
				ie.SpatialRelationInfoToReleaseListSizeExt_v1610 = []PUCCH_SpatialRelationInfoId{}
				for _, i := range tmp_SpatialRelationInfoToReleaseListSizeExt_v1610.Value {
					ie.SpatialRelationInfoToReleaseListSizeExt_v1610 = append(ie.SpatialRelationInfoToReleaseListSizeExt_v1610, *i)
				}
			}
			// decode SpatialRelationInfoToAddModListExt_v1610 optional
			if SpatialRelationInfoToAddModListExt_v1610Present {
				tmp_SpatialRelationInfoToAddModListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoExt_r16]([]*PUCCH_SpatialRelationInfoExt_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				fn_SpatialRelationInfoToAddModListExt_v1610 := func() *PUCCH_SpatialRelationInfoExt_r16 {
					return new(PUCCH_SpatialRelationInfoExt_r16)
				}
				if err = tmp_SpatialRelationInfoToAddModListExt_v1610.Decode(extReader, fn_SpatialRelationInfoToAddModListExt_v1610); err != nil {
					return utils.WrapError("Decode SpatialRelationInfoToAddModListExt_v1610", err)
				}
				ie.SpatialRelationInfoToAddModListExt_v1610 = []PUCCH_SpatialRelationInfoExt_r16{}
				for _, i := range tmp_SpatialRelationInfoToAddModListExt_v1610.Value {
					ie.SpatialRelationInfoToAddModListExt_v1610 = append(ie.SpatialRelationInfoToAddModListExt_v1610, *i)
				}
			}
			// decode SpatialRelationInfoToReleaseListExt_v1610 optional
			if SpatialRelationInfoToReleaseListExt_v1610Present {
				tmp_SpatialRelationInfoToReleaseListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId_r16]([]*PUCCH_SpatialRelationInfoId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				fn_SpatialRelationInfoToReleaseListExt_v1610 := func() *PUCCH_SpatialRelationInfoId_r16 {
					return new(PUCCH_SpatialRelationInfoId_r16)
				}
				if err = tmp_SpatialRelationInfoToReleaseListExt_v1610.Decode(extReader, fn_SpatialRelationInfoToReleaseListExt_v1610); err != nil {
					return utils.WrapError("Decode SpatialRelationInfoToReleaseListExt_v1610", err)
				}
				ie.SpatialRelationInfoToReleaseListExt_v1610 = []PUCCH_SpatialRelationInfoId_r16{}
				for _, i := range tmp_SpatialRelationInfoToReleaseListExt_v1610.Value {
					ie.SpatialRelationInfoToReleaseListExt_v1610 = append(ie.SpatialRelationInfoToReleaseListExt_v1610, *i)
				}
			}
			// decode ResourceGroupToAddModList_r16 optional
			if ResourceGroupToAddModList_r16Present {
				tmp_ResourceGroupToAddModList_r16 := utils.NewSequence[*PUCCH_ResourceGroup_r16]([]*PUCCH_ResourceGroup_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				fn_ResourceGroupToAddModList_r16 := func() *PUCCH_ResourceGroup_r16 {
					return new(PUCCH_ResourceGroup_r16)
				}
				if err = tmp_ResourceGroupToAddModList_r16.Decode(extReader, fn_ResourceGroupToAddModList_r16); err != nil {
					return utils.WrapError("Decode ResourceGroupToAddModList_r16", err)
				}
				ie.ResourceGroupToAddModList_r16 = []PUCCH_ResourceGroup_r16{}
				for _, i := range tmp_ResourceGroupToAddModList_r16.Value {
					ie.ResourceGroupToAddModList_r16 = append(ie.ResourceGroupToAddModList_r16, *i)
				}
			}
			// decode ResourceGroupToReleaseList_r16 optional
			if ResourceGroupToReleaseList_r16Present {
				tmp_ResourceGroupToReleaseList_r16 := utils.NewSequence[*PUCCH_ResourceGroupId_r16]([]*PUCCH_ResourceGroupId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				fn_ResourceGroupToReleaseList_r16 := func() *PUCCH_ResourceGroupId_r16 {
					return new(PUCCH_ResourceGroupId_r16)
				}
				if err = tmp_ResourceGroupToReleaseList_r16.Decode(extReader, fn_ResourceGroupToReleaseList_r16); err != nil {
					return utils.WrapError("Decode ResourceGroupToReleaseList_r16", err)
				}
				ie.ResourceGroupToReleaseList_r16 = []PUCCH_ResourceGroupId_r16{}
				for _, i := range tmp_ResourceGroupToReleaseList_r16.Value {
					ie.ResourceGroupToReleaseList_r16 = append(ie.ResourceGroupToReleaseList_r16, *i)
				}
			}
			// decode Sps_PUCCH_AN_List_r16 optional
			if Sps_PUCCH_AN_List_r16Present {
				tmp_Sps_PUCCH_AN_List_r16 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{}
				if err = tmp_Sps_PUCCH_AN_List_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_PUCCH_AN_List_r16", err)
				}
				ie.Sps_PUCCH_AN_List_r16 = tmp_Sps_PUCCH_AN_List_r16.Setup
			}
			// decode SchedulingRequestResourceToAddModListExt_v1610 optional
			if SchedulingRequestResourceToAddModListExt_v1610Present {
				tmp_SchedulingRequestResourceToAddModListExt_v1610 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1610]([]*SchedulingRequestResourceConfigExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				fn_SchedulingRequestResourceToAddModListExt_v1610 := func() *SchedulingRequestResourceConfigExt_v1610 {
					return new(SchedulingRequestResourceConfigExt_v1610)
				}
				if err = tmp_SchedulingRequestResourceToAddModListExt_v1610.Decode(extReader, fn_SchedulingRequestResourceToAddModListExt_v1610); err != nil {
					return utils.WrapError("Decode SchedulingRequestResourceToAddModListExt_v1610", err)
				}
				ie.SchedulingRequestResourceToAddModListExt_v1610 = []SchedulingRequestResourceConfigExt_v1610{}
				for _, i := range tmp_SchedulingRequestResourceToAddModListExt_v1610.Value {
					ie.SchedulingRequestResourceToAddModListExt_v1610 = append(ie.SchedulingRequestResourceToAddModListExt_v1610, *i)
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

			Format0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Format2Ext_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Format3Ext_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Format4Ext_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_AccessConfigListDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MappingPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PowerControlSetInfoToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PowerControlSetInfoToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SecondTPCFieldDCI_1_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SecondTPCFieldDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_DataToUL_ACK_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_DataToUL_ACK_DCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_AccessConfigListDCI_1_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SchedulingRequestResourceToAddModListExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_BundlingPUCCH_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_DataToUL_ACK_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_PUCCH_AN_ListMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Format0_r17 optional
			if Format0_r17Present {
				tmp_Format0_r17 := utils.SetupRelease[*PUCCH_FormatConfig]{}
				if err = tmp_Format0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Format0_r17", err)
				}
				ie.Format0_r17 = tmp_Format0_r17.Setup
			}
			// decode Format2Ext_r17 optional
			if Format2Ext_r17Present {
				tmp_Format2Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{}
				if err = tmp_Format2Ext_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Format2Ext_r17", err)
				}
				ie.Format2Ext_r17 = tmp_Format2Ext_r17.Setup
			}
			// decode Format3Ext_r17 optional
			if Format3Ext_r17Present {
				tmp_Format3Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{}
				if err = tmp_Format3Ext_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Format3Ext_r17", err)
				}
				ie.Format3Ext_r17 = tmp_Format3Ext_r17.Setup
			}
			// decode Format4Ext_r17 optional
			if Format4Ext_r17Present {
				tmp_Format4Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{}
				if err = tmp_Format4Ext_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Format4Ext_r17", err)
				}
				ie.Format4Ext_r17 = tmp_Format4Ext_r17.Setup
			}
			// decode Ul_AccessConfigListDCI_1_2_r17 optional
			if Ul_AccessConfigListDCI_1_2_r17Present {
				tmp_Ul_AccessConfigListDCI_1_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_2_r17]{}
				if err = tmp_Ul_AccessConfigListDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_AccessConfigListDCI_1_2_r17", err)
				}
				ie.Ul_AccessConfigListDCI_1_2_r17 = tmp_Ul_AccessConfigListDCI_1_2_r17.Setup
			}
			// decode MappingPattern_r17 optional
			if MappingPattern_r17Present {
				ie.MappingPattern_r17 = new(PUCCH_Config_mappingPattern_r17)
				if err = ie.MappingPattern_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MappingPattern_r17", err)
				}
			}
			// decode PowerControlSetInfoToAddModList_r17 optional
			if PowerControlSetInfoToAddModList_r17Present {
				tmp_PowerControlSetInfoToAddModList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfo_r17]([]*PUCCH_PowerControlSetInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				fn_PowerControlSetInfoToAddModList_r17 := func() *PUCCH_PowerControlSetInfo_r17 {
					return new(PUCCH_PowerControlSetInfo_r17)
				}
				if err = tmp_PowerControlSetInfoToAddModList_r17.Decode(extReader, fn_PowerControlSetInfoToAddModList_r17); err != nil {
					return utils.WrapError("Decode PowerControlSetInfoToAddModList_r17", err)
				}
				ie.PowerControlSetInfoToAddModList_r17 = []PUCCH_PowerControlSetInfo_r17{}
				for _, i := range tmp_PowerControlSetInfoToAddModList_r17.Value {
					ie.PowerControlSetInfoToAddModList_r17 = append(ie.PowerControlSetInfoToAddModList_r17, *i)
				}
			}
			// decode PowerControlSetInfoToReleaseList_r17 optional
			if PowerControlSetInfoToReleaseList_r17Present {
				tmp_PowerControlSetInfoToReleaseList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfoId_r17]([]*PUCCH_PowerControlSetInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				fn_PowerControlSetInfoToReleaseList_r17 := func() *PUCCH_PowerControlSetInfoId_r17 {
					return new(PUCCH_PowerControlSetInfoId_r17)
				}
				if err = tmp_PowerControlSetInfoToReleaseList_r17.Decode(extReader, fn_PowerControlSetInfoToReleaseList_r17); err != nil {
					return utils.WrapError("Decode PowerControlSetInfoToReleaseList_r17", err)
				}
				ie.PowerControlSetInfoToReleaseList_r17 = []PUCCH_PowerControlSetInfoId_r17{}
				for _, i := range tmp_PowerControlSetInfoToReleaseList_r17.Value {
					ie.PowerControlSetInfoToReleaseList_r17 = append(ie.PowerControlSetInfoToReleaseList_r17, *i)
				}
			}
			// decode SecondTPCFieldDCI_1_1_r17 optional
			if SecondTPCFieldDCI_1_1_r17Present {
				ie.SecondTPCFieldDCI_1_1_r17 = new(PUCCH_Config_secondTPCFieldDCI_1_1_r17)
				if err = ie.SecondTPCFieldDCI_1_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SecondTPCFieldDCI_1_1_r17", err)
				}
			}
			// decode SecondTPCFieldDCI_1_2_r17 optional
			if SecondTPCFieldDCI_1_2_r17Present {
				ie.SecondTPCFieldDCI_1_2_r17 = new(PUCCH_Config_secondTPCFieldDCI_1_2_r17)
				if err = ie.SecondTPCFieldDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SecondTPCFieldDCI_1_2_r17", err)
				}
			}
			// decode Dl_DataToUL_ACK_r17 optional
			if Dl_DataToUL_ACK_r17Present {
				tmp_Dl_DataToUL_ACK_r17 := utils.SetupRelease[*DL_DataToUL_ACK_r17]{}
				if err = tmp_Dl_DataToUL_ACK_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_DataToUL_ACK_r17", err)
				}
				ie.Dl_DataToUL_ACK_r17 = tmp_Dl_DataToUL_ACK_r17.Setup
			}
			// decode Dl_DataToUL_ACK_DCI_1_2_r17 optional
			if Dl_DataToUL_ACK_DCI_1_2_r17Present {
				tmp_Dl_DataToUL_ACK_DCI_1_2_r17 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r17]{}
				if err = tmp_Dl_DataToUL_ACK_DCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_DataToUL_ACK_DCI_1_2_r17", err)
				}
				ie.Dl_DataToUL_ACK_DCI_1_2_r17 = tmp_Dl_DataToUL_ACK_DCI_1_2_r17.Setup
			}
			// decode Ul_AccessConfigListDCI_1_1_r17 optional
			if Ul_AccessConfigListDCI_1_1_r17Present {
				tmp_Ul_AccessConfigListDCI_1_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r17]{}
				if err = tmp_Ul_AccessConfigListDCI_1_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_AccessConfigListDCI_1_1_r17", err)
				}
				ie.Ul_AccessConfigListDCI_1_1_r17 = tmp_Ul_AccessConfigListDCI_1_1_r17.Setup
			}
			// decode SchedulingRequestResourceToAddModListExt_v1700 optional
			if SchedulingRequestResourceToAddModListExt_v1700Present {
				tmp_SchedulingRequestResourceToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1700]([]*SchedulingRequestResourceConfigExt_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				fn_SchedulingRequestResourceToAddModListExt_v1700 := func() *SchedulingRequestResourceConfigExt_v1700 {
					return new(SchedulingRequestResourceConfigExt_v1700)
				}
				if err = tmp_SchedulingRequestResourceToAddModListExt_v1700.Decode(extReader, fn_SchedulingRequestResourceToAddModListExt_v1700); err != nil {
					return utils.WrapError("Decode SchedulingRequestResourceToAddModListExt_v1700", err)
				}
				ie.SchedulingRequestResourceToAddModListExt_v1700 = []SchedulingRequestResourceConfigExt_v1700{}
				for _, i := range tmp_SchedulingRequestResourceToAddModListExt_v1700.Value {
					ie.SchedulingRequestResourceToAddModListExt_v1700 = append(ie.SchedulingRequestResourceToAddModListExt_v1700, *i)
				}
			}
			// decode Dmrs_BundlingPUCCH_Config_r17 optional
			if Dmrs_BundlingPUCCH_Config_r17Present {
				tmp_Dmrs_BundlingPUCCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUCCH_Config_r17]{}
				if err = tmp_Dmrs_BundlingPUCCH_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingPUCCH_Config_r17", err)
				}
				ie.Dmrs_BundlingPUCCH_Config_r17 = tmp_Dmrs_BundlingPUCCH_Config_r17.Setup
			}
			// decode Dl_DataToUL_ACK_v1700 optional
			if Dl_DataToUL_ACK_v1700Present {
				tmp_Dl_DataToUL_ACK_v1700 := utils.SetupRelease[*DL_DataToUL_ACK_v1700]{}
				if err = tmp_Dl_DataToUL_ACK_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_DataToUL_ACK_v1700", err)
				}
				ie.Dl_DataToUL_ACK_v1700 = tmp_Dl_DataToUL_ACK_v1700.Setup
			}
			// decode Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 optional
			if Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17Present {
				tmp_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 := utils.SetupRelease[*DL_DataToUL_ACK_MulticastDCI_Format4_1_r17]{}
				if err = tmp_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17", err)
				}
				ie.Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 = tmp_Dl_DataToUL_ACK_MulticastDCI_Format4_1_r17.Setup
			}
			// decode Sps_PUCCH_AN_ListMulticast_r17 optional
			if Sps_PUCCH_AN_ListMulticast_r17Present {
				tmp_Sps_PUCCH_AN_ListMulticast_r17 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{}
				if err = tmp_Sps_PUCCH_AN_ListMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_PUCCH_AN_ListMulticast_r17", err)
				}
				ie.Sps_PUCCH_AN_ListMulticast_r17 = tmp_Sps_PUCCH_AN_ListMulticast_r17.Setup
			}
		}
	}
	return nil
}
