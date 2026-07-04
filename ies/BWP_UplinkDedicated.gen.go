// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BWP-UplinkDedicated (line 5392).

var bWPUplinkDedicatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-Config", Optional: true},
		{Name: "pusch-Config", Optional: true},
		{Name: "configuredGrantConfig", Optional: true},
		{Name: "srs-Config", Optional: true},
		{Name: "beamFailureRecoveryConfig", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

var bWP_UplinkDedicatedPucchConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Pucch_Config_Release = 0
	BWP_UplinkDedicated_Pucch_Config_Setup   = 1
)

var bWP_UplinkDedicatedPuschConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Pusch_Config_Release = 0
	BWP_UplinkDedicated_Pusch_Config_Setup   = 1
)

var bWP_UplinkDedicatedConfiguredGrantConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_ConfiguredGrantConfig_Release = 0
	BWP_UplinkDedicated_ConfiguredGrantConfig_Setup   = 1
)

var bWP_UplinkDedicatedSrsConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Srs_Config_Release = 0
	BWP_UplinkDedicated_Srs_Config_Setup   = 1
)

var bWP_UplinkDedicatedBeamFailureRecoveryConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_BeamFailureRecoveryConfig_Release = 0
	BWP_UplinkDedicated_BeamFailureRecoveryConfig_Setup   = 1
)

var bWPUplinkDedicatedExtSlPUCCHConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Sl_PUCCH_Config_r16_Release = 0
	BWP_UplinkDedicated_Ext_Sl_PUCCH_Config_r16_Setup   = 1
)

var bWPUplinkDedicatedCpExtensionC2R16Constraints = per.Constrained(1, 28)

var bWPUplinkDedicatedCpExtensionC3R16Constraints = per.Constrained(1, 28)

const (
	BWP_UplinkDedicated_Ext_UseInterlacePUCCH_PUSCH_r16_Enabled = 0
)

var bWPUplinkDedicatedExtUseInterlacePUCCHPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkDedicated_Ext_UseInterlacePUCCH_PUSCH_r16_Enabled},
}

var bWPUplinkDedicatedExtPucchConfigurationListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Pucch_ConfigurationList_r16_Release = 0
	BWP_UplinkDedicated_Ext_Pucch_ConfigurationList_r16_Setup   = 1
)

var bWPUplinkDedicatedExtLbtFailureRecoveryConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Lbt_FailureRecoveryConfig_r16_Release = 0
	BWP_UplinkDedicated_Ext_Lbt_FailureRecoveryConfig_r16_Setup   = 1
)

var bWPUplinkDedicatedExtUlTCIStateListR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "explicitlist"},
		{Name: "unifiedTCI-StateRef-r17"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Ul_TCI_StateList_r17_Explicitlist            = 0
	BWP_UplinkDedicated_Ext_Ul_TCI_StateList_r17_UnifiedTCI_StateRef_r17 = 1
)

var bWPUplinkDedicatedExtPucchConfigurationListMulticast1R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast1_r17_Release = 0
	BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast1_r17_Setup   = 1
)

var bWPUplinkDedicatedExtPucchConfigurationListMulticast2R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast2_r17_Release = 0
	BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast2_r17_Setup   = 1
)

var bWPUplinkDedicatedExtPucchConfigMulticast1R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast1_r17_Release = 0
	BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast1_r17_Setup   = 1
)

var bWPUplinkDedicatedExtPucchConfigMulticast2R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast2_r17_Release = 0
	BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast2_r17_Setup   = 1
)

var bWPUplinkDedicatedExtPathlossReferenceRSToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofPathlossReferenceRSs_r17)

var bWPUplinkDedicatedExtPathlossReferenceRSToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofPathlossReferenceRSs_r17)

const (
	BWP_UplinkDedicated_Ext_Sbfd_Config2_Transmission_r19_Enabled = 0
)

var bWPUplinkDedicatedExtSbfdConfig2TransmissionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkDedicated_Ext_Sbfd_Config2_Transmission_r19_Enabled},
}

var bWPUplinkDedicatedSbfdConfig2PUSCHRBOffsetR19Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

const (
	BWP_UplinkDedicated_Ext_Ul_Muting_NonSBFD_Symbol_r19_Enabled = 0
)

var bWPUplinkDedicatedExtUlMutingNonSBFDSymbolR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkDedicated_Ext_Ul_Muting_NonSBFD_Symbol_r19_Enabled},
}

const (
	BWP_UplinkDedicated_Ext_TwoTA_Without_MultiDCI_MultiTRP_r19_Enabled = 0
)

var bWPUplinkDedicatedExtTwoTAWithoutMultiDCIMultiTRPR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkDedicated_Ext_TwoTA_Without_MultiDCI_MultiTRP_r19_Enabled},
}

var bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-TCI-ToAddModList-r17", Optional: true},
		{Name: "ul-TCI-ToReleaseList-r17", Optional: true},
	},
}

var bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistUlTCIToAddModListR17Constraints = per.SizeRange(1, common.MaxUL_TCI_r17)

var bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistUlTCIToReleaseListR17Constraints = per.SizeRange(1, common.MaxUL_TCI_r17)

type BWP_UplinkDedicated struct {
	Pucch_Config *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_Config
	}
	Pusch_Config *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_Config
	}
	ConfiguredGrantConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *ConfiguredGrantConfig
	}
	Srs_Config *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_Config
	}
	BeamFailureRecoveryConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *BeamFailureRecoveryConfig
	}
	Sl_PUCCH_Config_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_Config
	}
	Cp_ExtensionC2_r16          *int64
	Cp_ExtensionC3_r16          *int64
	UseInterlacePUCCH_PUSCH_r16 *int64
	Pucch_ConfigurationList_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_ConfigurationList_r16
	}
	Lbt_FailureRecoveryConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *LBT_FailureRecoveryConfig_r16
	}
	ConfiguredGrantConfigToAddModList_r16               *ConfiguredGrantConfigToAddModList_r16
	ConfiguredGrantConfigToReleaseList_r16              *ConfiguredGrantConfigToReleaseList_r16
	ConfiguredGrantConfigType2DeactivationStateList_r16 *ConfiguredGrantConfigType2DeactivationStateList_r16
	Ul_TCI_StateList_r17                                *struct {
		Choice       int
		Explicitlist *struct {
			Ul_TCI_ToAddModList_r17  []TCI_UL_State_r17
			Ul_TCI_ToReleaseList_r17 []TCI_UL_StateId_r17
		}
		UnifiedTCI_StateRef_r17 *ServingCellAndBWP_Id_r17
	}
	Ul_PowerControl_r17                   *Uplink_PowerControlId_r17
	Pucch_ConfigurationListMulticast1_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_ConfigurationList_r16
	}
	Pucch_ConfigurationListMulticast2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_ConfigurationList_r16
	}
	Pucch_ConfigMulticast1_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_Config
	}
	Pucch_ConfigMulticast2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_Config
	}
	PathlossReferenceRSToAddModList_r17  []PathlossReferenceRS_r17
	PathlossReferenceRSToReleaseList_r17 []PathlossReferenceRS_Id_r17
	Sbfd_Config2_Transmission_r19        *int64
	Sbfd_Config2_PUSCH_RB_Offset_r19     *int64
	Ul_Muting_NonSBFD_Symbol_r19         *int64
	TwoTA_Without_MultiDCI_MultiTRP_r19  *int64
}

func (ie *BWP_UplinkDedicated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPUplinkDedicatedConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_PUCCH_Config_r16 != nil || ie.Cp_ExtensionC2_r16 != nil || ie.Cp_ExtensionC3_r16 != nil || ie.UseInterlacePUCCH_PUSCH_r16 != nil || ie.Pucch_ConfigurationList_r16 != nil || ie.Lbt_FailureRecoveryConfig_r16 != nil || ie.ConfiguredGrantConfigToAddModList_r16 != nil || ie.ConfiguredGrantConfigToReleaseList_r16 != nil || ie.ConfiguredGrantConfigType2DeactivationStateList_r16 != nil
	hasExtGroup1 := ie.Ul_TCI_StateList_r17 != nil || ie.Ul_PowerControl_r17 != nil || ie.Pucch_ConfigurationListMulticast1_r17 != nil || ie.Pucch_ConfigurationListMulticast2_r17 != nil
	hasExtGroup2 := ie.Pucch_ConfigMulticast1_r17 != nil || ie.Pucch_ConfigMulticast2_r17 != nil
	hasExtGroup3 := ie.PathlossReferenceRSToAddModList_r17 != nil || ie.PathlossReferenceRSToReleaseList_r17 != nil
	hasExtGroup4 := ie.Sbfd_Config2_Transmission_r19 != nil || ie.Sbfd_Config2_PUSCH_RB_Offset_r19 != nil || ie.Ul_Muting_NonSBFD_Symbol_r19 != nil || ie.TwoTA_Without_MultiDCI_MultiTRP_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pucch_Config != nil, ie.Pusch_Config != nil, ie.ConfiguredGrantConfig != nil, ie.Srs_Config != nil, ie.BeamFailureRecoveryConfig != nil}); err != nil {
		return err
	}

	// 3. pucch-Config: choice
	{
		if ie.Pucch_Config != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkDedicatedPucchConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pucch_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pucch_Config).Choice {
			case BWP_UplinkDedicated_Pucch_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Pucch_Config_Setup:
				if err := (*ie.Pucch_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pucch_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. pusch-Config: choice
	{
		if ie.Pusch_Config != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkDedicatedPuschConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pusch_Config).Choice {
			case BWP_UplinkDedicated_Pusch_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Pusch_Config_Setup:
				if err := (*ie.Pusch_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pusch_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. configuredGrantConfig: choice
	{
		if ie.ConfiguredGrantConfig != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkDedicatedConfiguredGrantConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ConfiguredGrantConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ConfiguredGrantConfig).Choice {
			case BWP_UplinkDedicated_ConfiguredGrantConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_ConfiguredGrantConfig_Setup:
				if err := (*ie.ConfiguredGrantConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ConfiguredGrantConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. srs-Config: choice
	{
		if ie.Srs_Config != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkDedicatedSrsConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_Config).Choice {
			case BWP_UplinkDedicated_Srs_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Srs_Config_Setup:
				if err := (*ie.Srs_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. beamFailureRecoveryConfig: choice
	{
		if ie.BeamFailureRecoveryConfig != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkDedicatedBeamFailureRecoveryConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.BeamFailureRecoveryConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.BeamFailureRecoveryConfig).Choice {
			case BWP_UplinkDedicated_BeamFailureRecoveryConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_BeamFailureRecoveryConfig_Setup:
				if err := (*ie.BeamFailureRecoveryConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.BeamFailureRecoveryConfig).Choice), Constraint: "index out of range"}
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
					{Name: "sl-PUCCH-Config-r16", Optional: true},
					{Name: "cp-ExtensionC2-r16", Optional: true},
					{Name: "cp-ExtensionC3-r16", Optional: true},
					{Name: "useInterlacePUCCH-PUSCH-r16", Optional: true},
					{Name: "pucch-ConfigurationList-r16", Optional: true},
					{Name: "lbt-FailureRecoveryConfig-r16", Optional: true},
					{Name: "configuredGrantConfigToAddModList-r16", Optional: true},
					{Name: "configuredGrantConfigToReleaseList-r16", Optional: true},
					{Name: "configuredGrantConfigType2DeactivationStateList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PUCCH_Config_r16 != nil, ie.Cp_ExtensionC2_r16 != nil, ie.Cp_ExtensionC3_r16 != nil, ie.UseInterlacePUCCH_PUSCH_r16 != nil, ie.Pucch_ConfigurationList_r16 != nil, ie.Lbt_FailureRecoveryConfig_r16 != nil, ie.ConfiguredGrantConfigToAddModList_r16 != nil, ie.ConfiguredGrantConfigToReleaseList_r16 != nil, ie.ConfiguredGrantConfigType2DeactivationStateList_r16 != nil}); err != nil {
				return err
			}

			if ie.Sl_PUCCH_Config_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtSlPUCCHConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PUCCH_Config_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_PUCCH_Config_r16).Choice {
				case BWP_UplinkDedicated_Ext_Sl_PUCCH_Config_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkDedicated_Ext_Sl_PUCCH_Config_r16_Setup:
					if err := (*ie.Sl_PUCCH_Config_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Cp_ExtensionC2_r16 != nil {
				if err := ex.EncodeInteger(*ie.Cp_ExtensionC2_r16, bWPUplinkDedicatedCpExtensionC2R16Constraints); err != nil {
					return err
				}
			}

			if ie.Cp_ExtensionC3_r16 != nil {
				if err := ex.EncodeInteger(*ie.Cp_ExtensionC3_r16, bWPUplinkDedicatedCpExtensionC3R16Constraints); err != nil {
					return err
				}
			}

			if ie.UseInterlacePUCCH_PUSCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UseInterlacePUCCH_PUSCH_r16, bWPUplinkDedicatedExtUseInterlacePUCCHPUSCHR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_ConfigurationList_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtPucchConfigurationListR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pucch_ConfigurationList_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pucch_ConfigurationList_r16).Choice {
				case BWP_UplinkDedicated_Ext_Pucch_ConfigurationList_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkDedicated_Ext_Pucch_ConfigurationList_r16_Setup:
					if err := (*ie.Pucch_ConfigurationList_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Lbt_FailureRecoveryConfig_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtLbtFailureRecoveryConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lbt_FailureRecoveryConfig_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lbt_FailureRecoveryConfig_r16).Choice {
				case BWP_UplinkDedicated_Ext_Lbt_FailureRecoveryConfig_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkDedicated_Ext_Lbt_FailureRecoveryConfig_r16_Setup:
					if err := (*ie.Lbt_FailureRecoveryConfig_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ConfiguredGrantConfigToAddModList_r16 != nil {
				if err := ie.ConfiguredGrantConfigToAddModList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ConfiguredGrantConfigToReleaseList_r16 != nil {
				if err := ie.ConfiguredGrantConfigToReleaseList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ConfiguredGrantConfigType2DeactivationStateList_r16 != nil {
				if err := ie.ConfiguredGrantConfigType2DeactivationStateList_r16.Encode(ex); err != nil {
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
					{Name: "ul-TCI-StateList-r17", Optional: true},
					{Name: "ul-powerControl-r17", Optional: true},
					{Name: "pucch-ConfigurationListMulticast1-r17", Optional: true},
					{Name: "pucch-ConfigurationListMulticast2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ul_TCI_StateList_r17 != nil, ie.Ul_PowerControl_r17 != nil, ie.Pucch_ConfigurationListMulticast1_r17 != nil, ie.Pucch_ConfigurationListMulticast2_r17 != nil}); err != nil {
				return err
			}

			if ie.Ul_TCI_StateList_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtUlTCIStateListR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_TCI_StateList_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_TCI_StateList_r17).Choice {
				case BWP_UplinkDedicated_Ext_Ul_TCI_StateList_r17_Explicitlist:
					c := (*(*ie.Ul_TCI_StateList_r17).Explicitlist)
					bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistSeq := ex.NewSequenceEncoder(bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistConstraints)
					if err := bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistSeq.EncodePreamble([]bool{c.Ul_TCI_ToAddModList_r17 != nil, c.Ul_TCI_ToReleaseList_r17 != nil}); err != nil {
						return err
					}
					if c.Ul_TCI_ToAddModList_r17 != nil {
						seqOf := ex.NewSequenceOfEncoder(bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistUlTCIToAddModListR17Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Ul_TCI_ToAddModList_r17))); err != nil {
							return err
						}
						for i := range c.Ul_TCI_ToAddModList_r17 {
							if err := c.Ul_TCI_ToAddModList_r17[i].Encode(ex); err != nil {
								return err
							}
						}
					}
					if c.Ul_TCI_ToReleaseList_r17 != nil {
						seqOf := ex.NewSequenceOfEncoder(bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistUlTCIToReleaseListR17Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Ul_TCI_ToReleaseList_r17))); err != nil {
							return err
						}
						for i := range c.Ul_TCI_ToReleaseList_r17 {
							if err := c.Ul_TCI_ToReleaseList_r17[i].Encode(ex); err != nil {
								return err
							}
						}
					}
				case BWP_UplinkDedicated_Ext_Ul_TCI_StateList_r17_UnifiedTCI_StateRef_r17:
					if err := (*ie.Ul_TCI_StateList_r17).UnifiedTCI_StateRef_r17.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ul_PowerControl_r17 != nil {
				if err := ie.Ul_PowerControl_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Pucch_ConfigurationListMulticast1_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtPucchConfigurationListMulticast1R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pucch_ConfigurationListMulticast1_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pucch_ConfigurationListMulticast1_r17).Choice {
				case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast1_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast1_r17_Setup:
					if err := (*ie.Pucch_ConfigurationListMulticast1_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pucch_ConfigurationListMulticast2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtPucchConfigurationListMulticast2R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pucch_ConfigurationListMulticast2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pucch_ConfigurationListMulticast2_r17).Choice {
				case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast2_r17_Setup:
					if err := (*ie.Pucch_ConfigurationListMulticast2_r17).Setup.Encode(ex); err != nil {
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
					{Name: "pucch-ConfigMulticast1-r17", Optional: true},
					{Name: "pucch-ConfigMulticast2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pucch_ConfigMulticast1_r17 != nil, ie.Pucch_ConfigMulticast2_r17 != nil}); err != nil {
				return err
			}

			if ie.Pucch_ConfigMulticast1_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtPucchConfigMulticast1R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pucch_ConfigMulticast1_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pucch_ConfigMulticast1_r17).Choice {
				case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast1_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast1_r17_Setup:
					if err := (*ie.Pucch_ConfigMulticast1_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pucch_ConfigMulticast2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkDedicatedExtPucchConfigMulticast2R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pucch_ConfigMulticast2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pucch_ConfigMulticast2_r17).Choice {
				case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast2_r17_Setup:
					if err := (*ie.Pucch_ConfigMulticast2_r17).Setup.Encode(ex); err != nil {
						return err
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
					{Name: "pathlossReferenceRSToAddModList-r17", Optional: true},
					{Name: "pathlossReferenceRSToReleaseList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PathlossReferenceRSToAddModList_r17 != nil, ie.PathlossReferenceRSToReleaseList_r17 != nil}); err != nil {
				return err
			}

			if ie.PathlossReferenceRSToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(bWPUplinkDedicatedExtPathlossReferenceRSToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRSToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.PathlossReferenceRSToAddModList_r17 {
					if err := ie.PathlossReferenceRSToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PathlossReferenceRSToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(bWPUplinkDedicatedExtPathlossReferenceRSToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRSToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.PathlossReferenceRSToReleaseList_r17 {
					if err := ie.PathlossReferenceRSToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "sbfd-Config2-Transmission-r19", Optional: true},
					{Name: "sbfd-Config2-PUSCH-RB-Offset-r19", Optional: true},
					{Name: "ul-Muting-NonSBFD-Symbol-r19", Optional: true},
					{Name: "twoTA-Without-MultiDCI-MultiTRP-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sbfd_Config2_Transmission_r19 != nil, ie.Sbfd_Config2_PUSCH_RB_Offset_r19 != nil, ie.Ul_Muting_NonSBFD_Symbol_r19 != nil, ie.TwoTA_Without_MultiDCI_MultiTRP_r19 != nil}); err != nil {
				return err
			}

			if ie.Sbfd_Config2_Transmission_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Sbfd_Config2_Transmission_r19, bWPUplinkDedicatedExtSbfdConfig2TransmissionR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sbfd_Config2_PUSCH_RB_Offset_r19 != nil {
				if err := ex.EncodeInteger(*ie.Sbfd_Config2_PUSCH_RB_Offset_r19, bWPUplinkDedicatedSbfdConfig2PUSCHRBOffsetR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_Muting_NonSBFD_Symbol_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_Muting_NonSBFD_Symbol_r19, bWPUplinkDedicatedExtUlMutingNonSBFDSymbolR19Constraints); err != nil {
					return err
				}
			}

			if ie.TwoTA_Without_MultiDCI_MultiTRP_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoTA_Without_MultiDCI_MultiTRP_r19, bWPUplinkDedicatedExtTwoTAWithoutMultiDCIMultiTRPR19Constraints); err != nil {
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

func (ie *BWP_UplinkDedicated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPUplinkDedicatedConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pucch-Config: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Pucch_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkDedicatedPucchConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pucch_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkDedicated_Pucch_Config_Release:
				(*ie.Pucch_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Pucch_Config_Setup:
				(*ie.Pucch_Config).Setup = new(PUCCH_Config)
				if err := (*ie.Pucch_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. pusch-Config: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Pusch_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkDedicatedPuschConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkDedicated_Pusch_Config_Release:
				(*ie.Pusch_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Pusch_Config_Setup:
				(*ie.Pusch_Config).Setup = new(PUSCH_Config)
				if err := (*ie.Pusch_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. configuredGrantConfig: choice
	{
		if seq.IsComponentPresent(2) {
			ie.ConfiguredGrantConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *ConfiguredGrantConfig
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkDedicatedConfiguredGrantConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ConfiguredGrantConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkDedicated_ConfiguredGrantConfig_Release:
				(*ie.ConfiguredGrantConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_ConfiguredGrantConfig_Setup:
				(*ie.ConfiguredGrantConfig).Setup = new(ConfiguredGrantConfig)
				if err := (*ie.ConfiguredGrantConfig).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. srs-Config: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Srs_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkDedicatedSrsConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkDedicated_Srs_Config_Release:
				(*ie.Srs_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Srs_Config_Setup:
				(*ie.Srs_Config).Setup = new(SRS_Config)
				if err := (*ie.Srs_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. beamFailureRecoveryConfig: choice
	{
		if seq.IsComponentPresent(4) {
			ie.BeamFailureRecoveryConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *BeamFailureRecoveryConfig
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkDedicatedBeamFailureRecoveryConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BeamFailureRecoveryConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkDedicated_BeamFailureRecoveryConfig_Release:
				(*ie.BeamFailureRecoveryConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_BeamFailureRecoveryConfig_Setup:
				(*ie.BeamFailureRecoveryConfig).Setup = new(BeamFailureRecoveryConfig)
				if err := (*ie.BeamFailureRecoveryConfig).Setup.Decode(d); err != nil {
					return err
				}
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
				{Name: "sl-PUCCH-Config-r16", Optional: true},
				{Name: "cp-ExtensionC2-r16", Optional: true},
				{Name: "cp-ExtensionC3-r16", Optional: true},
				{Name: "useInterlacePUCCH-PUSCH-r16", Optional: true},
				{Name: "pucch-ConfigurationList-r16", Optional: true},
				{Name: "lbt-FailureRecoveryConfig-r16", Optional: true},
				{Name: "configuredGrantConfigToAddModList-r16", Optional: true},
				{Name: "configuredGrantConfigToReleaseList-r16", Optional: true},
				{Name: "configuredGrantConfigType2DeactivationStateList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_PUCCH_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_Config
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtSlPUCCHConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PUCCH_Config_r16).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Sl_PUCCH_Config_r16_Release:
				(*ie.Sl_PUCCH_Config_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Ext_Sl_PUCCH_Config_r16_Setup:
				(*ie.Sl_PUCCH_Config_r16).Setup = new(PUCCH_Config)
				if err := (*ie.Sl_PUCCH_Config_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(bWPUplinkDedicatedCpExtensionC2R16Constraints)
			if err != nil {
				return err
			}
			ie.Cp_ExtensionC2_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(bWPUplinkDedicatedCpExtensionC3R16Constraints)
			if err != nil {
				return err
			}
			ie.Cp_ExtensionC3_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bWPUplinkDedicatedExtUseInterlacePUCCHPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.UseInterlacePUCCH_PUSCH_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Pucch_ConfigurationList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_ConfigurationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtPucchConfigurationListR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pucch_ConfigurationList_r16).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Pucch_ConfigurationList_r16_Release:
				(*ie.Pucch_ConfigurationList_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Ext_Pucch_ConfigurationList_r16_Setup:
				(*ie.Pucch_ConfigurationList_r16).Setup = new(PUCCH_ConfigurationList_r16)
				if err := (*ie.Pucch_ConfigurationList_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Lbt_FailureRecoveryConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LBT_FailureRecoveryConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtLbtFailureRecoveryConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lbt_FailureRecoveryConfig_r16).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Lbt_FailureRecoveryConfig_r16_Release:
				(*ie.Lbt_FailureRecoveryConfig_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Ext_Lbt_FailureRecoveryConfig_r16_Setup:
				(*ie.Lbt_FailureRecoveryConfig_r16).Setup = new(LBT_FailureRecoveryConfig_r16)
				if err := (*ie.Lbt_FailureRecoveryConfig_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.ConfiguredGrantConfigToAddModList_r16 = new(ConfiguredGrantConfigToAddModList_r16)
			if err := ie.ConfiguredGrantConfigToAddModList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.ConfiguredGrantConfigToReleaseList_r16 = new(ConfiguredGrantConfigToReleaseList_r16)
			if err := ie.ConfiguredGrantConfigToReleaseList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(8) {
			ie.ConfiguredGrantConfigType2DeactivationStateList_r16 = new(ConfiguredGrantConfigType2DeactivationStateList_r16)
			if err := ie.ConfiguredGrantConfigType2DeactivationStateList_r16.Decode(dx); err != nil {
				return err
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
				{Name: "ul-TCI-StateList-r17", Optional: true},
				{Name: "ul-powerControl-r17", Optional: true},
				{Name: "pucch-ConfigurationListMulticast1-r17", Optional: true},
				{Name: "pucch-ConfigurationListMulticast2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ul_TCI_StateList_r17 = &struct {
				Choice       int
				Explicitlist *struct {
					Ul_TCI_ToAddModList_r17  []TCI_UL_State_r17
					Ul_TCI_ToReleaseList_r17 []TCI_UL_StateId_r17
				}
				UnifiedTCI_StateRef_r17 *ServingCellAndBWP_Id_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtUlTCIStateListR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_TCI_StateList_r17).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Ul_TCI_StateList_r17_Explicitlist:
				(*ie.Ul_TCI_StateList_r17).Explicitlist = &struct {
					Ul_TCI_ToAddModList_r17  []TCI_UL_State_r17
					Ul_TCI_ToReleaseList_r17 []TCI_UL_StateId_r17
				}{}
				c := (*(*ie.Ul_TCI_StateList_r17).Explicitlist)
				bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistSeq := dx.NewSequenceDecoder(bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistConstraints)
				if err := bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistSeq.DecodePreamble(); err != nil {
					return err
				}
				if bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistSeq.IsComponentPresent(0) {
					seqOf := dx.NewSequenceOfDecoder(bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistUlTCIToAddModListR17Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Ul_TCI_ToAddModList_r17 = make([]TCI_UL_State_r17, n)
					for i := int64(0); i < n; i++ {
						if err := c.Ul_TCI_ToAddModList_r17[i].Decode(dx); err != nil {
							return err
						}
					}
				}
				if bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistSeq.IsComponentPresent(1) {
					seqOf := dx.NewSequenceOfDecoder(bWPUplinkDedicatedExtUlTCIStateListR17ExplicitlistUlTCIToReleaseListR17Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Ul_TCI_ToReleaseList_r17 = make([]TCI_UL_StateId_r17, n)
					for i := int64(0); i < n; i++ {
						if err := c.Ul_TCI_ToReleaseList_r17[i].Decode(dx); err != nil {
							return err
						}
					}
				}
			case BWP_UplinkDedicated_Ext_Ul_TCI_StateList_r17_UnifiedTCI_StateRef_r17:
				(*ie.Ul_TCI_StateList_r17).UnifiedTCI_StateRef_r17 = new(ServingCellAndBWP_Id_r17)
				if err := (*ie.Ul_TCI_StateList_r17).UnifiedTCI_StateRef_r17.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ul_PowerControl_r17 = new(Uplink_PowerControlId_r17)
			if err := ie.Ul_PowerControl_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Pucch_ConfigurationListMulticast1_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_ConfigurationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtPucchConfigurationListMulticast1R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pucch_ConfigurationListMulticast1_r17).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast1_r17_Release:
				(*ie.Pucch_ConfigurationListMulticast1_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast1_r17_Setup:
				(*ie.Pucch_ConfigurationListMulticast1_r17).Setup = new(PUCCH_ConfigurationList_r16)
				if err := (*ie.Pucch_ConfigurationListMulticast1_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Pucch_ConfigurationListMulticast2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_ConfigurationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtPucchConfigurationListMulticast2R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pucch_ConfigurationListMulticast2_r17).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast2_r17_Release:
				(*ie.Pucch_ConfigurationListMulticast2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Ext_Pucch_ConfigurationListMulticast2_r17_Setup:
				(*ie.Pucch_ConfigurationListMulticast2_r17).Setup = new(PUCCH_ConfigurationList_r16)
				if err := (*ie.Pucch_ConfigurationListMulticast2_r17).Setup.Decode(dx); err != nil {
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
				{Name: "pucch-ConfigMulticast1-r17", Optional: true},
				{Name: "pucch-ConfigMulticast2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pucch_ConfigMulticast1_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_Config
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtPucchConfigMulticast1R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pucch_ConfigMulticast1_r17).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast1_r17_Release:
				(*ie.Pucch_ConfigMulticast1_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast1_r17_Setup:
				(*ie.Pucch_ConfigMulticast1_r17).Setup = new(PUCCH_Config)
				if err := (*ie.Pucch_ConfigMulticast1_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Pucch_ConfigMulticast2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_Config
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkDedicatedExtPucchConfigMulticast2R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pucch_ConfigMulticast2_r17).Choice = int(index)
			switch index {
			case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast2_r17_Release:
				(*ie.Pucch_ConfigMulticast2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicated_Ext_Pucch_ConfigMulticast2_r17_Setup:
				(*ie.Pucch_ConfigMulticast2_r17).Setup = new(PUCCH_Config)
				if err := (*ie.Pucch_ConfigMulticast2_r17).Setup.Decode(dx); err != nil {
					return err
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
				{Name: "pathlossReferenceRSToAddModList-r17", Optional: true},
				{Name: "pathlossReferenceRSToReleaseList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(bWPUplinkDedicatedExtPathlossReferenceRSToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRSToAddModList_r17 = make([]PathlossReferenceRS_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRSToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(bWPUplinkDedicatedExtPathlossReferenceRSToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRSToReleaseList_r17 = make([]PathlossReferenceRS_Id_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRSToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
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
				{Name: "sbfd-Config2-Transmission-r19", Optional: true},
				{Name: "sbfd-Config2-PUSCH-RB-Offset-r19", Optional: true},
				{Name: "ul-Muting-NonSBFD-Symbol-r19", Optional: true},
				{Name: "twoTA-Without-MultiDCI-MultiTRP-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bWPUplinkDedicatedExtSbfdConfig2TransmissionR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_Config2_Transmission_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(bWPUplinkDedicatedSbfdConfig2PUSCHRBOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_Config2_PUSCH_RB_Offset_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bWPUplinkDedicatedExtUlMutingNonSBFDSymbolR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_Muting_NonSBFD_Symbol_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bWPUplinkDedicatedExtTwoTAWithoutMultiDCIMultiTRPR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoTA_Without_MultiDCI_MultiTRP_r19 = &v
		}
	}

	return nil
}
