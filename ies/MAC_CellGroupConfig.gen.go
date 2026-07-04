// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MAC-CellGroupConfig (line 8992).

var mACCellGroupConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-Config", Optional: true},
		{Name: "schedulingRequestConfig", Optional: true},
		{Name: "bsr-Config", Optional: true},
		{Name: "tag-Config", Optional: true},
		{Name: "phr-Config", Optional: true},
		{Name: "skipUplinkTxDynamic"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
		{Name: "extension-addition-7", Optional: true},
	},
}

var mAC_CellGroupConfigDrxConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Drx_Config_Release = 0
	MAC_CellGroupConfig_Drx_Config_Setup   = 1
)

var mAC_CellGroupConfigPhrConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Phr_Config_Release = 0
	MAC_CellGroupConfig_Phr_Config_Setup   = 1
)

var mACCellGroupConfigExtDataInactivityTimerConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Ext_DataInactivityTimer_Release = 0
	MAC_CellGroupConfig_Ext_DataInactivityTimer_Setup   = 1
)

const (
	MAC_CellGroupConfig_Ext_UsePreBSR_r16_True = 0
)

var mACCellGroupConfigExtUsePreBSRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_UsePreBSR_r16_True},
}

const (
	MAC_CellGroupConfig_Ext_Lch_BasedPrioritization_r16_Enabled = 0
)

var mACCellGroupConfigExtLchBasedPrioritizationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_Lch_BasedPrioritization_r16_Enabled},
}

var mACCellGroupConfigExtDrxConfigSecondaryGroupR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Ext_Drx_ConfigSecondaryGroup_r16_Release = 0
	MAC_CellGroupConfig_Ext_Drx_ConfigSecondaryGroup_r16_Setup   = 1
)

const (
	MAC_CellGroupConfig_Ext_EnhancedSkipUplinkTxDynamic_r16_True = 0
)

var mACCellGroupConfigExtEnhancedSkipUplinkTxDynamicR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_EnhancedSkipUplinkTxDynamic_r16_True},
}

const (
	MAC_CellGroupConfig_Ext_EnhancedSkipUplinkTxConfigured_r16_True = 0
)

var mACCellGroupConfigExtEnhancedSkipUplinkTxConfiguredR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_EnhancedSkipUplinkTxConfigured_r16_True},
}

const (
	MAC_CellGroupConfig_Ext_IntraCG_Prioritization_r17_Enabled = 0
)

var mACCellGroupConfigExtIntraCGPrioritizationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_IntraCG_Prioritization_r17_Enabled},
}

var mACCellGroupConfigExtDrxConfigSLR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Ext_Drx_ConfigSL_r17_Release = 0
	MAC_CellGroupConfig_Ext_Drx_ConfigSL_r17_Setup   = 1
)

var mACCellGroupConfigExtDrxConfigExtV1700Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Ext_Drx_ConfigExt_v1700_Release = 0
	MAC_CellGroupConfig_Ext_Drx_ConfigExt_v1700_Setup   = 1
)

var mACCellGroupConfigExtTarConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Ext_Tar_Config_r17_Release = 0
	MAC_CellGroupConfig_Ext_Tar_Config_r17_Setup   = 1
)

var mACCellGroupConfigExtGRNTIConfigToAddModListR17Constraints = per.SizeRange(1, common.MaxG_RNTI_r17)

var mACCellGroupConfigExtGRNTIConfigToReleaseListR17Constraints = per.SizeRange(1, common.MaxG_RNTI_r17)

var mACCellGroupConfigExtGCSRNTIConfigToAddModListR17Constraints = per.SizeRange(1, common.MaxG_CS_RNTI_r17)

var mACCellGroupConfigExtGCSRNTIConfigToReleaseListR17Constraints = per.SizeRange(1, common.MaxG_CS_RNTI_r17)

const (
	MAC_CellGroupConfig_Ext_Drx_LastTransmissionUL_r17_Enabled = 0
)

var mACCellGroupConfigExtDrxLastTransmissionULR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_Drx_LastTransmissionUL_r17_Enabled},
}

const (
	MAC_CellGroupConfig_Ext_PosMG_Request_r17_Enabled = 0
)

var mACCellGroupConfigExtPosMGRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_PosMG_Request_r17_Enabled},
}

var mACCellGroupConfigExtDrxConfigExt2V1800Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Ext_Drx_ConfigExt2_v1800_Release = 0
	MAC_CellGroupConfig_Ext_Drx_ConfigExt2_v1800_Setup   = 1
)

var mACCellGroupConfigAdditionalBSTableAllowedR18Constraints = per.FixedSize(common.MaxNrofLCGs_r18)

var mACCellGroupConfigExtDsrConfigToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofLCGs_r18)

var mACCellGroupConfigExtDsrConfigToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofLCGs_r18)

var mACCellGroupConfigExtTarConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	MAC_CellGroupConfig_Ext_Tar_Config_r18_Release = 0
	MAC_CellGroupConfig_Ext_Tar_Config_r18_Setup   = 1
)

var mACCellGroupConfigExtUlRateControlConfigListR19Constraints = per.SizeRange(1, common.MaxNrofRateCtrlQFIs_r19)

var mACCellGroupConfigExtUlRateQueryR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-RateQueryConfigList-r19"},
		{Name: "ul-RateQueryProhibitTimer-r19"},
	},
}

var mACCellGroupConfigExtUlRateQueryR19UlRateQueryConfigListR19Constraints = per.SizeRange(1, common.MaxNrofRateQueryQFIs_r19)

const (
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0     = 0
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0dot1 = 1
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0dot2 = 2
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0dot5 = 3
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S1     = 4
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S2     = 5
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S5     = 6
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S10    = 7
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S20    = 8
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S30    = 9
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S60    = 10
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S90    = 11
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S120   = 12
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S300   = 13
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S600   = 14
	MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_Spare1 = 15
)

var mACCellGroupConfigExtUlRateQueryR19UlRateQueryProhibitTimerR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0dot1, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0dot2, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S0dot5, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S1, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S2, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S5, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S10, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S20, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S30, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S60, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S90, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S120, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S300, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_S600, MAC_CellGroupConfig_Ext_Ul_RateQuery_r19_Ul_RateQueryProhibitTimer_r19_Spare1},
}

type MAC_CellGroupConfig struct {
	Drx_Config *struct {
		Choice  int
		Release *struct{}
		Setup   *DRX_Config
	}
	SchedulingRequestConfig *SchedulingRequestConfig
	Bsr_Config              *BSR_Config
	Tag_Config              *TAG_Config
	Phr_Config              *struct {
		Choice  int
		Release *struct{}
		Setup   *PHR_Config
	}
	SkipUplinkTxDynamic bool
	Csi_Mask            *bool
	DataInactivityTimer *struct {
		Choice  int
		Release *struct{}
		Setup   *DataInactivityTimer
	}
	UsePreBSR_r16                     *int64
	SchedulingRequestID_LBT_SCell_r16 *SchedulingRequestId
	Lch_BasedPrioritization_r16       *int64
	SchedulingRequestID_BFR_SCell_r16 *SchedulingRequestId
	Drx_ConfigSecondaryGroup_r16      *struct {
		Choice  int
		Release *struct{}
		Setup   *DRX_ConfigSecondaryGroup_r16
	}
	EnhancedSkipUplinkTxDynamic_r16    *int64
	EnhancedSkipUplinkTxConfigured_r16 *int64
	IntraCG_Prioritization_r17         *int64
	Drx_ConfigSL_r17                   *struct {
		Choice  int
		Release *struct{}
		Setup   *DRX_ConfigSL_r17
	}
	Drx_ConfigExt_v1700 *struct {
		Choice  int
		Release *struct{}
		Setup   *DRX_ConfigExt_v1700
	}
	SchedulingRequestID_BFR_r17   *SchedulingRequestId
	SchedulingRequestID_BFR2_r17  *SchedulingRequestId
	SchedulingRequestConfig_v1700 *SchedulingRequestConfig_v1700
	Tar_Config_r17                *struct {
		Choice  int
		Release *struct{}
		Setup   *TAR_Config_r17
	}
	G_RNTI_ConfigToAddModList_r17           []MBS_RNTI_SpecificConfig_r17
	G_RNTI_ConfigToReleaseList_r17          []MBS_RNTI_SpecificConfigId_r17
	G_CS_RNTI_ConfigToAddModList_r17        []MBS_RNTI_SpecificConfig_r17
	G_CS_RNTI_ConfigToReleaseList_r17       []MBS_RNTI_SpecificConfigId_r17
	AllowCSI_SRS_Tx_MulticastDRX_Active_r17 *bool
	SchedulingRequestID_PosMG_Request_r17   *SchedulingRequestId
	Drx_LastTransmissionUL_r17              *int64
	PosMG_Request_r17                       *int64
	Drx_ConfigExt2_v1800                    *struct {
		Choice  int
		Release *struct{}
		Setup   *DRX_ConfigExt2_v1800
	}
	AdditionalBS_TableAllowed_r18 *per.BitString
	Dsr_ConfigToAddModList_r18    []LCG_DSR_Config_r18
	Dsr_ConfigToReleaseList_r18   []LCG_Id_r18
	Tar_Config_r18                *struct {
		Choice  int
		Release *struct{}
		Setup   *TAR_Config_r18
	}
	Ul_RateQuery_r19 *struct {
		Ul_RateQueryConfigList_r19    []QoS_FlowIdentity_r19
		Ul_RateQueryProhibitTimer_r19 int64
	}
	Ul_RateControlConfigList_r19 []QoS_FlowIdentity_r19
	SchedulingRequestID_LTM_r19  *SchedulingRequestId
}

func (ie *MAC_CellGroupConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACCellGroupConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Csi_Mask != nil || ie.DataInactivityTimer != nil
	hasExtGroup1 := ie.UsePreBSR_r16 != nil || ie.SchedulingRequestID_LBT_SCell_r16 != nil || ie.Lch_BasedPrioritization_r16 != nil || ie.SchedulingRequestID_BFR_SCell_r16 != nil || ie.Drx_ConfigSecondaryGroup_r16 != nil
	hasExtGroup2 := ie.EnhancedSkipUplinkTxDynamic_r16 != nil || ie.EnhancedSkipUplinkTxConfigured_r16 != nil
	hasExtGroup3 := ie.IntraCG_Prioritization_r17 != nil || ie.Drx_ConfigSL_r17 != nil || ie.Drx_ConfigExt_v1700 != nil || ie.SchedulingRequestID_BFR_r17 != nil || ie.SchedulingRequestID_BFR2_r17 != nil || ie.SchedulingRequestConfig_v1700 != nil || ie.Tar_Config_r17 != nil || ie.G_RNTI_ConfigToAddModList_r17 != nil || ie.G_RNTI_ConfigToReleaseList_r17 != nil || ie.G_CS_RNTI_ConfigToAddModList_r17 != nil || ie.G_CS_RNTI_ConfigToReleaseList_r17 != nil || ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil
	hasExtGroup4 := ie.SchedulingRequestID_PosMG_Request_r17 != nil || ie.Drx_LastTransmissionUL_r17 != nil
	hasExtGroup5 := ie.PosMG_Request_r17 != nil
	hasExtGroup6 := ie.Drx_ConfigExt2_v1800 != nil || ie.AdditionalBS_TableAllowed_r18 != nil || ie.Dsr_ConfigToAddModList_r18 != nil || ie.Dsr_ConfigToReleaseList_r18 != nil || ie.Tar_Config_r18 != nil
	hasExtGroup7 := ie.Ul_RateQuery_r19 != nil || ie.Ul_RateControlConfigList_r19 != nil || ie.SchedulingRequestID_LTM_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Drx_Config != nil, ie.SchedulingRequestConfig != nil, ie.Bsr_Config != nil, ie.Tag_Config != nil, ie.Phr_Config != nil}); err != nil {
		return err
	}

	// 3. drx-Config: choice
	{
		if ie.Drx_Config != nil {
			choiceEnc := e.NewChoiceEncoder(mAC_CellGroupConfigDrxConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Drx_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Drx_Config).Choice {
			case MAC_CellGroupConfig_Drx_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Drx_Config_Setup:
				if err := (*ie.Drx_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Drx_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. schedulingRequestConfig: ref
	{
		if ie.SchedulingRequestConfig != nil {
			if err := ie.SchedulingRequestConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. bsr-Config: ref
	{
		if ie.Bsr_Config != nil {
			if err := ie.Bsr_Config.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. tag-Config: ref
	{
		if ie.Tag_Config != nil {
			if err := ie.Tag_Config.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. phr-Config: choice
	{
		if ie.Phr_Config != nil {
			choiceEnc := e.NewChoiceEncoder(mAC_CellGroupConfigPhrConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Phr_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Phr_Config).Choice {
			case MAC_CellGroupConfig_Phr_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Phr_Config_Setup:
				if err := (*ie.Phr_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Phr_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. skipUplinkTxDynamic: boolean
	{
		if err := e.EncodeBoolean(ie.SkipUplinkTxDynamic); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "csi-Mask", Optional: true},
					{Name: "dataInactivityTimer", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Csi_Mask != nil, ie.DataInactivityTimer != nil}); err != nil {
				return err
			}

			if ie.Csi_Mask != nil {
				if err := ex.EncodeBoolean(*ie.Csi_Mask); err != nil {
					return err
				}
			}

			if ie.DataInactivityTimer != nil {
				choiceEnc := ex.NewChoiceEncoder(mACCellGroupConfigExtDataInactivityTimerConstraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.DataInactivityTimer).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.DataInactivityTimer).Choice {
				case MAC_CellGroupConfig_Ext_DataInactivityTimer_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MAC_CellGroupConfig_Ext_DataInactivityTimer_Setup:
					if err := (*ie.DataInactivityTimer).Setup.Encode(ex); err != nil {
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
					{Name: "usePreBSR-r16", Optional: true},
					{Name: "schedulingRequestID-LBT-SCell-r16", Optional: true},
					{Name: "lch-BasedPrioritization-r16", Optional: true},
					{Name: "schedulingRequestID-BFR-SCell-r16", Optional: true},
					{Name: "drx-ConfigSecondaryGroup-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.UsePreBSR_r16 != nil, ie.SchedulingRequestID_LBT_SCell_r16 != nil, ie.Lch_BasedPrioritization_r16 != nil, ie.SchedulingRequestID_BFR_SCell_r16 != nil, ie.Drx_ConfigSecondaryGroup_r16 != nil}); err != nil {
				return err
			}

			if ie.UsePreBSR_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UsePreBSR_r16, mACCellGroupConfigExtUsePreBSRR16Constraints); err != nil {
					return err
				}
			}

			if ie.SchedulingRequestID_LBT_SCell_r16 != nil {
				if err := ie.SchedulingRequestID_LBT_SCell_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Lch_BasedPrioritization_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Lch_BasedPrioritization_r16, mACCellGroupConfigExtLchBasedPrioritizationR16Constraints); err != nil {
					return err
				}
			}

			if ie.SchedulingRequestID_BFR_SCell_r16 != nil {
				if err := ie.SchedulingRequestID_BFR_SCell_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Drx_ConfigSecondaryGroup_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(mACCellGroupConfigExtDrxConfigSecondaryGroupR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Drx_ConfigSecondaryGroup_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Drx_ConfigSecondaryGroup_r16).Choice {
				case MAC_CellGroupConfig_Ext_Drx_ConfigSecondaryGroup_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MAC_CellGroupConfig_Ext_Drx_ConfigSecondaryGroup_r16_Setup:
					if err := (*ie.Drx_ConfigSecondaryGroup_r16).Setup.Encode(ex); err != nil {
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
					{Name: "enhancedSkipUplinkTxDynamic-r16", Optional: true},
					{Name: "enhancedSkipUplinkTxConfigured-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnhancedSkipUplinkTxDynamic_r16 != nil, ie.EnhancedSkipUplinkTxConfigured_r16 != nil}); err != nil {
				return err
			}

			if ie.EnhancedSkipUplinkTxDynamic_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedSkipUplinkTxDynamic_r16, mACCellGroupConfigExtEnhancedSkipUplinkTxDynamicR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnhancedSkipUplinkTxConfigured_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedSkipUplinkTxConfigured_r16, mACCellGroupConfigExtEnhancedSkipUplinkTxConfiguredR16Constraints); err != nil {
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
					{Name: "intraCG-Prioritization-r17", Optional: true},
					{Name: "drx-ConfigSL-r17", Optional: true},
					{Name: "drx-ConfigExt-v1700", Optional: true},
					{Name: "schedulingRequestID-BFR-r17", Optional: true},
					{Name: "schedulingRequestID-BFR2-r17", Optional: true},
					{Name: "schedulingRequestConfig-v1700", Optional: true},
					{Name: "tar-Config-r17", Optional: true},
					{Name: "g-RNTI-ConfigToAddModList-r17", Optional: true},
					{Name: "g-RNTI-ConfigToReleaseList-r17", Optional: true},
					{Name: "g-CS-RNTI-ConfigToAddModList-r17", Optional: true},
					{Name: "g-CS-RNTI-ConfigToReleaseList-r17", Optional: true},
					{Name: "allowCSI-SRS-Tx-MulticastDRX-Active-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IntraCG_Prioritization_r17 != nil, ie.Drx_ConfigSL_r17 != nil, ie.Drx_ConfigExt_v1700 != nil, ie.SchedulingRequestID_BFR_r17 != nil, ie.SchedulingRequestID_BFR2_r17 != nil, ie.SchedulingRequestConfig_v1700 != nil, ie.Tar_Config_r17 != nil, ie.G_RNTI_ConfigToAddModList_r17 != nil, ie.G_RNTI_ConfigToReleaseList_r17 != nil, ie.G_CS_RNTI_ConfigToAddModList_r17 != nil, ie.G_CS_RNTI_ConfigToReleaseList_r17 != nil, ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil}); err != nil {
				return err
			}

			if ie.IntraCG_Prioritization_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.IntraCG_Prioritization_r17, mACCellGroupConfigExtIntraCGPrioritizationR17Constraints); err != nil {
					return err
				}
			}

			if ie.Drx_ConfigSL_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(mACCellGroupConfigExtDrxConfigSLR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Drx_ConfigSL_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Drx_ConfigSL_r17).Choice {
				case MAC_CellGroupConfig_Ext_Drx_ConfigSL_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MAC_CellGroupConfig_Ext_Drx_ConfigSL_r17_Setup:
					if err := (*ie.Drx_ConfigSL_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Drx_ConfigExt_v1700 != nil {
				choiceEnc := ex.NewChoiceEncoder(mACCellGroupConfigExtDrxConfigExtV1700Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Drx_ConfigExt_v1700).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Drx_ConfigExt_v1700).Choice {
				case MAC_CellGroupConfig_Ext_Drx_ConfigExt_v1700_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MAC_CellGroupConfig_Ext_Drx_ConfigExt_v1700_Setup:
					if err := (*ie.Drx_ConfigExt_v1700).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SchedulingRequestID_BFR_r17 != nil {
				if err := ie.SchedulingRequestID_BFR_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SchedulingRequestID_BFR2_r17 != nil {
				if err := ie.SchedulingRequestID_BFR2_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SchedulingRequestConfig_v1700 != nil {
				if err := ie.SchedulingRequestConfig_v1700.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Tar_Config_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(mACCellGroupConfigExtTarConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Tar_Config_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Tar_Config_r17).Choice {
				case MAC_CellGroupConfig_Ext_Tar_Config_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MAC_CellGroupConfig_Ext_Tar_Config_r17_Setup:
					if err := (*ie.Tar_Config_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.G_RNTI_ConfigToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtGRNTIConfigToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.G_RNTI_ConfigToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.G_RNTI_ConfigToAddModList_r17 {
					if err := ie.G_RNTI_ConfigToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.G_RNTI_ConfigToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtGRNTIConfigToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.G_RNTI_ConfigToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.G_RNTI_ConfigToReleaseList_r17 {
					if err := ie.G_RNTI_ConfigToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.G_CS_RNTI_ConfigToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtGCSRNTIConfigToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.G_CS_RNTI_ConfigToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.G_CS_RNTI_ConfigToAddModList_r17 {
					if err := ie.G_CS_RNTI_ConfigToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.G_CS_RNTI_ConfigToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtGCSRNTIConfigToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.G_CS_RNTI_ConfigToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.G_CS_RNTI_ConfigToReleaseList_r17 {
					if err := ie.G_CS_RNTI_ConfigToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil {
				if err := ex.EncodeBoolean(*ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17); err != nil {
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
					{Name: "schedulingRequestID-PosMG-Request-r17", Optional: true},
					{Name: "drx-LastTransmissionUL-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SchedulingRequestID_PosMG_Request_r17 != nil, ie.Drx_LastTransmissionUL_r17 != nil}); err != nil {
				return err
			}

			if ie.SchedulingRequestID_PosMG_Request_r17 != nil {
				if err := ie.SchedulingRequestID_PosMG_Request_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Drx_LastTransmissionUL_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Drx_LastTransmissionUL_r17, mACCellGroupConfigExtDrxLastTransmissionULR17Constraints); err != nil {
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
					{Name: "posMG-Request-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PosMG_Request_r17 != nil}); err != nil {
				return err
			}

			if ie.PosMG_Request_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PosMG_Request_r17, mACCellGroupConfigExtPosMGRequestR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 6:
		if hasExtGroup6 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "drx-ConfigExt2-v1800", Optional: true},
					{Name: "additionalBS-TableAllowed-r18", Optional: true},
					{Name: "dsr-ConfigToAddModList-r18", Optional: true},
					{Name: "dsr-ConfigToReleaseList-r18", Optional: true},
					{Name: "tar-Config-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Drx_ConfigExt2_v1800 != nil, ie.AdditionalBS_TableAllowed_r18 != nil, ie.Dsr_ConfigToAddModList_r18 != nil, ie.Dsr_ConfigToReleaseList_r18 != nil, ie.Tar_Config_r18 != nil}); err != nil {
				return err
			}

			if ie.Drx_ConfigExt2_v1800 != nil {
				choiceEnc := ex.NewChoiceEncoder(mACCellGroupConfigExtDrxConfigExt2V1800Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Drx_ConfigExt2_v1800).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Drx_ConfigExt2_v1800).Choice {
				case MAC_CellGroupConfig_Ext_Drx_ConfigExt2_v1800_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MAC_CellGroupConfig_Ext_Drx_ConfigExt2_v1800_Setup:
					if err := (*ie.Drx_ConfigExt2_v1800).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AdditionalBS_TableAllowed_r18 != nil {
				if err := ex.EncodeBitString(*ie.AdditionalBS_TableAllowed_r18, mACCellGroupConfigAdditionalBSTableAllowedR18Constraints); err != nil {
					return err
				}
			}

			if ie.Dsr_ConfigToAddModList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtDsrConfigToAddModListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Dsr_ConfigToAddModList_r18))); err != nil {
					return err
				}
				for i := range ie.Dsr_ConfigToAddModList_r18 {
					if err := ie.Dsr_ConfigToAddModList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dsr_ConfigToReleaseList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtDsrConfigToReleaseListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Dsr_ConfigToReleaseList_r18))); err != nil {
					return err
				}
				for i := range ie.Dsr_ConfigToReleaseList_r18 {
					if err := ie.Dsr_ConfigToReleaseList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Tar_Config_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(mACCellGroupConfigExtTarConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Tar_Config_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Tar_Config_r18).Choice {
				case MAC_CellGroupConfig_Ext_Tar_Config_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case MAC_CellGroupConfig_Ext_Tar_Config_r18_Setup:
					if err := (*ie.Tar_Config_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 7:
		if hasExtGroup7 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ul-RateQuery-r19", Optional: true},
					{Name: "ul-RateControlConfigList-r19", Optional: true},
					{Name: "schedulingRequestID-LTM-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ul_RateQuery_r19 != nil, ie.Ul_RateControlConfigList_r19 != nil, ie.SchedulingRequestID_LTM_r19 != nil}); err != nil {
				return err
			}

			if ie.Ul_RateQuery_r19 != nil {
				c := ie.Ul_RateQuery_r19
				mACCellGroupConfigExtUlRateQueryR19Seq := ex.NewSequenceEncoder(mACCellGroupConfigExtUlRateQueryR19Constraints)
				if err := mACCellGroupConfigExtUlRateQueryR19Seq.EncodeExtensionBit(false); err != nil {
					return err
				}
				{
					seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtUlRateQueryR19UlRateQueryConfigListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Ul_RateQueryConfigList_r19))); err != nil {
						return err
					}
					for i := range c.Ul_RateQueryConfigList_r19 {
						if err := c.Ul_RateQueryConfigList_r19[i].Encode(ex); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeEnumerated(c.Ul_RateQueryProhibitTimer_r19, mACCellGroupConfigExtUlRateQueryR19UlRateQueryProhibitTimerR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_RateControlConfigList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(mACCellGroupConfigExtUlRateControlConfigListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Ul_RateControlConfigList_r19))); err != nil {
					return err
				}
				for i := range ie.Ul_RateControlConfigList_r19 {
					if err := ie.Ul_RateControlConfigList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SchedulingRequestID_LTM_r19 != nil {
				if err := ie.SchedulingRequestID_LTM_r19.Encode(ex); err != nil {
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

func (ie *MAC_CellGroupConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACCellGroupConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. drx-Config: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Drx_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *DRX_Config
			}{}
			choiceDec := d.NewChoiceDecoder(mAC_CellGroupConfigDrxConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Drx_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MAC_CellGroupConfig_Drx_Config_Release:
				(*ie.Drx_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Drx_Config_Setup:
				(*ie.Drx_Config).Setup = new(DRX_Config)
				if err := (*ie.Drx_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. schedulingRequestConfig: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SchedulingRequestConfig = new(SchedulingRequestConfig)
			if err := ie.SchedulingRequestConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. bsr-Config: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Bsr_Config = new(BSR_Config)
			if err := ie.Bsr_Config.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. tag-Config: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Tag_Config = new(TAG_Config)
			if err := ie.Tag_Config.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. phr-Config: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Phr_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *PHR_Config
			}{}
			choiceDec := d.NewChoiceDecoder(mAC_CellGroupConfigPhrConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Phr_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MAC_CellGroupConfig_Phr_Config_Release:
				(*ie.Phr_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Phr_Config_Setup:
				(*ie.Phr_Config).Setup = new(PHR_Config)
				if err := (*ie.Phr_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. skipUplinkTxDynamic: boolean
	{
		v5, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.SkipUplinkTxDynamic = v5
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
				{Name: "csi-Mask", Optional: true},
				{Name: "dataInactivityTimer", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Csi_Mask = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.DataInactivityTimer = &struct {
				Choice  int
				Release *struct{}
				Setup   *DataInactivityTimer
			}{}
			choiceDec := dx.NewChoiceDecoder(mACCellGroupConfigExtDataInactivityTimerConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DataInactivityTimer).Choice = int(index)
			switch index {
			case MAC_CellGroupConfig_Ext_DataInactivityTimer_Release:
				(*ie.DataInactivityTimer).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Ext_DataInactivityTimer_Setup:
				(*ie.DataInactivityTimer).Setup = new(DataInactivityTimer)
				if err := (*ie.DataInactivityTimer).Setup.Decode(dx); err != nil {
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
				{Name: "usePreBSR-r16", Optional: true},
				{Name: "schedulingRequestID-LBT-SCell-r16", Optional: true},
				{Name: "lch-BasedPrioritization-r16", Optional: true},
				{Name: "schedulingRequestID-BFR-SCell-r16", Optional: true},
				{Name: "drx-ConfigSecondaryGroup-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACCellGroupConfigExtUsePreBSRR16Constraints)
			if err != nil {
				return err
			}
			ie.UsePreBSR_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SchedulingRequestID_LBT_SCell_r16 = new(SchedulingRequestId)
			if err := ie.SchedulingRequestID_LBT_SCell_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mACCellGroupConfigExtLchBasedPrioritizationR16Constraints)
			if err != nil {
				return err
			}
			ie.Lch_BasedPrioritization_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.SchedulingRequestID_BFR_SCell_r16 = new(SchedulingRequestId)
			if err := ie.SchedulingRequestID_BFR_SCell_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Drx_ConfigSecondaryGroup_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DRX_ConfigSecondaryGroup_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(mACCellGroupConfigExtDrxConfigSecondaryGroupR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Drx_ConfigSecondaryGroup_r16).Choice = int(index)
			switch index {
			case MAC_CellGroupConfig_Ext_Drx_ConfigSecondaryGroup_r16_Release:
				(*ie.Drx_ConfigSecondaryGroup_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Ext_Drx_ConfigSecondaryGroup_r16_Setup:
				(*ie.Drx_ConfigSecondaryGroup_r16).Setup = new(DRX_ConfigSecondaryGroup_r16)
				if err := (*ie.Drx_ConfigSecondaryGroup_r16).Setup.Decode(dx); err != nil {
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
				{Name: "enhancedSkipUplinkTxDynamic-r16", Optional: true},
				{Name: "enhancedSkipUplinkTxConfigured-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACCellGroupConfigExtEnhancedSkipUplinkTxDynamicR16Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedSkipUplinkTxDynamic_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACCellGroupConfigExtEnhancedSkipUplinkTxConfiguredR16Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedSkipUplinkTxConfigured_r16 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "intraCG-Prioritization-r17", Optional: true},
				{Name: "drx-ConfigSL-r17", Optional: true},
				{Name: "drx-ConfigExt-v1700", Optional: true},
				{Name: "schedulingRequestID-BFR-r17", Optional: true},
				{Name: "schedulingRequestID-BFR2-r17", Optional: true},
				{Name: "schedulingRequestConfig-v1700", Optional: true},
				{Name: "tar-Config-r17", Optional: true},
				{Name: "g-RNTI-ConfigToAddModList-r17", Optional: true},
				{Name: "g-RNTI-ConfigToReleaseList-r17", Optional: true},
				{Name: "g-CS-RNTI-ConfigToAddModList-r17", Optional: true},
				{Name: "g-CS-RNTI-ConfigToReleaseList-r17", Optional: true},
				{Name: "allowCSI-SRS-Tx-MulticastDRX-Active-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACCellGroupConfigExtIntraCGPrioritizationR17Constraints)
			if err != nil {
				return err
			}
			ie.IntraCG_Prioritization_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Drx_ConfigSL_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DRX_ConfigSL_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(mACCellGroupConfigExtDrxConfigSLR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Drx_ConfigSL_r17).Choice = int(index)
			switch index {
			case MAC_CellGroupConfig_Ext_Drx_ConfigSL_r17_Release:
				(*ie.Drx_ConfigSL_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Ext_Drx_ConfigSL_r17_Setup:
				(*ie.Drx_ConfigSL_r17).Setup = new(DRX_ConfigSL_r17)
				if err := (*ie.Drx_ConfigSL_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Drx_ConfigExt_v1700 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DRX_ConfigExt_v1700
			}{}
			choiceDec := dx.NewChoiceDecoder(mACCellGroupConfigExtDrxConfigExtV1700Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Drx_ConfigExt_v1700).Choice = int(index)
			switch index {
			case MAC_CellGroupConfig_Ext_Drx_ConfigExt_v1700_Release:
				(*ie.Drx_ConfigExt_v1700).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Ext_Drx_ConfigExt_v1700_Setup:
				(*ie.Drx_ConfigExt_v1700).Setup = new(DRX_ConfigExt_v1700)
				if err := (*ie.Drx_ConfigExt_v1700).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.SchedulingRequestID_BFR_r17 = new(SchedulingRequestId)
			if err := ie.SchedulingRequestID_BFR_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.SchedulingRequestID_BFR2_r17 = new(SchedulingRequestId)
			if err := ie.SchedulingRequestID_BFR2_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.SchedulingRequestConfig_v1700 = new(SchedulingRequestConfig_v1700)
			if err := ie.SchedulingRequestConfig_v1700.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Tar_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *TAR_Config_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(mACCellGroupConfigExtTarConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Tar_Config_r17).Choice = int(index)
			switch index {
			case MAC_CellGroupConfig_Ext_Tar_Config_r17_Release:
				(*ie.Tar_Config_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Ext_Tar_Config_r17_Setup:
				(*ie.Tar_Config_r17).Setup = new(TAR_Config_r17)
				if err := (*ie.Tar_Config_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtGRNTIConfigToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.G_RNTI_ConfigToAddModList_r17 = make([]MBS_RNTI_SpecificConfig_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.G_RNTI_ConfigToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtGRNTIConfigToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.G_RNTI_ConfigToReleaseList_r17 = make([]MBS_RNTI_SpecificConfigId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.G_RNTI_ConfigToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtGCSRNTIConfigToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.G_CS_RNTI_ConfigToAddModList_r17 = make([]MBS_RNTI_SpecificConfig_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.G_CS_RNTI_ConfigToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtGCSRNTIConfigToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.G_CS_RNTI_ConfigToReleaseList_r17 = make([]MBS_RNTI_SpecificConfigId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.G_CS_RNTI_ConfigToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "schedulingRequestID-PosMG-Request-r17", Optional: true},
				{Name: "drx-LastTransmissionUL-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SchedulingRequestID_PosMG_Request_r17 = new(SchedulingRequestId)
			if err := ie.SchedulingRequestID_PosMG_Request_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACCellGroupConfigExtDrxLastTransmissionULR17Constraints)
			if err != nil {
				return err
			}
			ie.Drx_LastTransmissionUL_r17 = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "posMG-Request-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACCellGroupConfigExtPosMGRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.PosMG_Request_r17 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "drx-ConfigExt2-v1800", Optional: true},
				{Name: "additionalBS-TableAllowed-r18", Optional: true},
				{Name: "dsr-ConfigToAddModList-r18", Optional: true},
				{Name: "dsr-ConfigToReleaseList-r18", Optional: true},
				{Name: "tar-Config-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Drx_ConfigExt2_v1800 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DRX_ConfigExt2_v1800
			}{}
			choiceDec := dx.NewChoiceDecoder(mACCellGroupConfigExtDrxConfigExt2V1800Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Drx_ConfigExt2_v1800).Choice = int(index)
			switch index {
			case MAC_CellGroupConfig_Ext_Drx_ConfigExt2_v1800_Release:
				(*ie.Drx_ConfigExt2_v1800).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Ext_Drx_ConfigExt2_v1800_Setup:
				(*ie.Drx_ConfigExt2_v1800).Setup = new(DRX_ConfigExt2_v1800)
				if err := (*ie.Drx_ConfigExt2_v1800).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeBitString(mACCellGroupConfigAdditionalBSTableAllowedR18Constraints)
			if err != nil {
				return err
			}
			ie.AdditionalBS_TableAllowed_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtDsrConfigToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dsr_ConfigToAddModList_r18 = make([]LCG_DSR_Config_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dsr_ConfigToAddModList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtDsrConfigToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dsr_ConfigToReleaseList_r18 = make([]LCG_Id_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dsr_ConfigToReleaseList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Tar_Config_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *TAR_Config_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(mACCellGroupConfigExtTarConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Tar_Config_r18).Choice = int(index)
			switch index {
			case MAC_CellGroupConfig_Ext_Tar_Config_r18_Release:
				(*ie.Tar_Config_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case MAC_CellGroupConfig_Ext_Tar_Config_r18_Setup:
				(*ie.Tar_Config_r18).Setup = new(TAR_Config_r18)
				if err := (*ie.Tar_Config_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ul-RateQuery-r19", Optional: true},
				{Name: "ul-RateControlConfigList-r19", Optional: true},
				{Name: "schedulingRequestID-LTM-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ul_RateQuery_r19 = &struct {
				Ul_RateQueryConfigList_r19    []QoS_FlowIdentity_r19
				Ul_RateQueryProhibitTimer_r19 int64
			}{}
			c := ie.Ul_RateQuery_r19
			mACCellGroupConfigExtUlRateQueryR19Seq := dx.NewSequenceDecoder(mACCellGroupConfigExtUlRateQueryR19Constraints)
			if err := mACCellGroupConfigExtUlRateQueryR19Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtUlRateQueryR19UlRateQueryConfigListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Ul_RateQueryConfigList_r19 = make([]QoS_FlowIdentity_r19, n)
				for i := int64(0); i < n; i++ {
					if err := c.Ul_RateQueryConfigList_r19[i].Decode(dx); err != nil {
						return err
					}
				}
			}
			{
				v, err := dx.DecodeEnumerated(mACCellGroupConfigExtUlRateQueryR19UlRateQueryProhibitTimerR19Constraints)
				if err != nil {
					return err
				}
				c.Ul_RateQueryProhibitTimer_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(mACCellGroupConfigExtUlRateControlConfigListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ul_RateControlConfigList_r19 = make([]QoS_FlowIdentity_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ul_RateControlConfigList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SchedulingRequestID_LTM_r19 = new(SchedulingRequestId)
			if err := ie.SchedulingRequestID_LTM_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
