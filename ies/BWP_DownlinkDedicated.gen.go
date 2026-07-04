// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BWP-DownlinkDedicated (line 5259).

var bWPDownlinkDedicatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-Config", Optional: true},
		{Name: "pdsch-Config", Optional: true},
		{Name: "sps-Config", Optional: true},
		{Name: "radioLinkMonitoringConfig", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var bWP_DownlinkDedicatedPdcchConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Pdcch_Config_Release = 0
	BWP_DownlinkDedicated_Pdcch_Config_Setup   = 1
)

var bWP_DownlinkDedicatedPdschConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Pdsch_Config_Release = 0
	BWP_DownlinkDedicated_Pdsch_Config_Setup   = 1
)

var bWP_DownlinkDedicatedSpsConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Sps_Config_Release = 0
	BWP_DownlinkDedicated_Sps_Config_Setup   = 1
)

var bWP_DownlinkDedicatedRadioLinkMonitoringConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_RadioLinkMonitoringConfig_Release = 0
	BWP_DownlinkDedicated_RadioLinkMonitoringConfig_Setup   = 1
)

var bWPDownlinkDedicatedExtBeamFailureRecoverySCellConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Ext_BeamFailureRecoverySCellConfig_r16_Release = 0
	BWP_DownlinkDedicated_Ext_BeamFailureRecoverySCellConfig_r16_Setup   = 1
)

var bWPDownlinkDedicatedExtSlPDCCHConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Ext_Sl_PDCCH_Config_r16_Release = 0
	BWP_DownlinkDedicated_Ext_Sl_PDCCH_Config_r16_Setup   = 1
)

var bWPDownlinkDedicatedExtSlV2XPDCCHConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Ext_Sl_V2X_PDCCH_Config_r16_Release = 0
	BWP_DownlinkDedicated_Ext_Sl_V2X_PDCCH_Config_r16_Setup   = 1
)

var bWPDownlinkDedicatedPreConfGapStatusR17Constraints = per.FixedSize(common.MaxNrofGapId_r17)

var bWPDownlinkDedicatedExtBeamFailureRecoverySpCellConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Ext_BeamFailureRecoverySpCellConfig_r17_Release = 0
	BWP_DownlinkDedicated_Ext_BeamFailureRecoverySpCellConfig_r17_Setup   = 1
)

var bWPDownlinkDedicatedExtCfrConfigMulticastR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Ext_Cfr_ConfigMulticast_r17_Release = 0
	BWP_DownlinkDedicated_Ext_Cfr_ConfigMulticast_r17_Setup   = 1
)

var bWPDownlinkDedicatedExtTciInDCIR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicated_Ext_Tci_InDCI_r18_Release = 0
	BWP_DownlinkDedicated_Ext_Tci_InDCI_r18_Setup   = 1
)

const (
	BWP_DownlinkDedicated_Ext_Sbfd_Config2_Reception_r19_Enabled = 0
)

var bWPDownlinkDedicatedExtSbfdConfig2ReceptionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_DownlinkDedicated_Ext_Sbfd_Config2_Reception_r19_Enabled},
}

const (
	BWP_DownlinkDedicated_Ext_PathlossOffsetPRACH_DCI_1_0_r19_Enabled = 0
)

var bWPDownlinkDedicatedExtPathlossOffsetPRACHDCI10R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_DownlinkDedicated_Ext_PathlossOffsetPRACH_DCI_1_0_r19_Enabled},
}

const (
	BWP_DownlinkDedicated_Ext_PrachAssociationDCI_1_0_r19_Enabled = 0
)

var bWPDownlinkDedicatedExtPrachAssociationDCI10R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_DownlinkDedicated_Ext_PrachAssociationDCI_1_0_r19_Enabled},
}

type BWP_DownlinkDedicated struct {
	Pdcch_Config *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_Config
	}
	Pdsch_Config *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_Config
	}
	Sps_Config *struct {
		Choice  int
		Release *struct{}
		Setup   *SPS_Config
	}
	RadioLinkMonitoringConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *RadioLinkMonitoringConfig
	}
	Sps_ConfigToAddModList_r16          *SPS_ConfigToAddModList_r16
	Sps_ConfigToReleaseList_r16         *SPS_ConfigToReleaseList_r16
	Sps_ConfigDeactivationStateList_r16 *SPS_ConfigDeactivationStateList_r16
	BeamFailureRecoverySCellConfig_r16  *struct {
		Choice  int
		Release *struct{}
		Setup   *BeamFailureRecoveryRSConfig_r16
	}
	Sl_PDCCH_Config_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_Config
	}
	Sl_V2X_PDCCH_Config_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_Config
	}
	PreConfGapStatus_r17                *per.BitString
	BeamFailureRecoverySpCellConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BeamFailureRecoveryRSConfig_r16
	}
	Harq_FeedbackEnablingforSPSactive_r17 *bool
	Cfr_ConfigMulticast_r17               *struct {
		Choice  int
		Release *struct{}
		Setup   *CFR_ConfigMulticast_r17
	}
	Dl_PPW_PreConfigToAddModList_r17  *DL_PPW_PreConfigToAddModList_r17
	Dl_PPW_PreConfigToReleaseList_r17 *DL_PPW_PreConfigToReleaseList_r17
	NonCellDefiningSSB_r17            *NonCellDefiningSSB_r17
	ServingCellMO_r17                 *MeasObjectId
	Tci_InDCI_r18                     *struct {
		Choice  int
		Release *struct{}
		Setup   *TCI_InDCI_r18
	}
	Sbfd_Config2_Reception_r19      *int64
	PathlossOffsetPRACH_DCI_1_0_r19 *int64
	PrachAssociationDCI_1_0_r19     *int64
}

func (ie *BWP_DownlinkDedicated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPDownlinkDedicatedConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sps_ConfigToAddModList_r16 != nil || ie.Sps_ConfigToReleaseList_r16 != nil || ie.Sps_ConfigDeactivationStateList_r16 != nil || ie.BeamFailureRecoverySCellConfig_r16 != nil || ie.Sl_PDCCH_Config_r16 != nil || ie.Sl_V2X_PDCCH_Config_r16 != nil
	hasExtGroup1 := ie.PreConfGapStatus_r17 != nil || ie.BeamFailureRecoverySpCellConfig_r17 != nil || ie.Harq_FeedbackEnablingforSPSactive_r17 != nil || ie.Cfr_ConfigMulticast_r17 != nil || ie.Dl_PPW_PreConfigToAddModList_r17 != nil || ie.Dl_PPW_PreConfigToReleaseList_r17 != nil || ie.NonCellDefiningSSB_r17 != nil || ie.ServingCellMO_r17 != nil
	hasExtGroup2 := ie.Tci_InDCI_r18 != nil
	hasExtGroup3 := ie.Sbfd_Config2_Reception_r19 != nil || ie.PathlossOffsetPRACH_DCI_1_0_r19 != nil || ie.PrachAssociationDCI_1_0_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_Config != nil, ie.Pdsch_Config != nil, ie.Sps_Config != nil, ie.RadioLinkMonitoringConfig != nil}); err != nil {
		return err
	}

	// 3. pdcch-Config: choice
	{
		if ie.Pdcch_Config != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkDedicatedPdcchConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdcch_Config).Choice {
			case BWP_DownlinkDedicated_Pdcch_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Pdcch_Config_Setup:
				if err := (*ie.Pdcch_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdcch_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. pdsch-Config: choice
	{
		if ie.Pdsch_Config != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkDedicatedPdschConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdsch_Config).Choice {
			case BWP_DownlinkDedicated_Pdsch_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Pdsch_Config_Setup:
				if err := (*ie.Pdsch_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdsch_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. sps-Config: choice
	{
		if ie.Sps_Config != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkDedicatedSpsConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sps_Config).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sps_Config).Choice {
			case BWP_DownlinkDedicated_Sps_Config_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Sps_Config_Setup:
				if err := (*ie.Sps_Config).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sps_Config).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. radioLinkMonitoringConfig: choice
	{
		if ie.RadioLinkMonitoringConfig != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkDedicatedRadioLinkMonitoringConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.RadioLinkMonitoringConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.RadioLinkMonitoringConfig).Choice {
			case BWP_DownlinkDedicated_RadioLinkMonitoringConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_RadioLinkMonitoringConfig_Setup:
				if err := (*ie.RadioLinkMonitoringConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.RadioLinkMonitoringConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sps-ConfigToAddModList-r16", Optional: true},
					{Name: "sps-ConfigToReleaseList-r16", Optional: true},
					{Name: "sps-ConfigDeactivationStateList-r16", Optional: true},
					{Name: "beamFailureRecoverySCellConfig-r16", Optional: true},
					{Name: "sl-PDCCH-Config-r16", Optional: true},
					{Name: "sl-V2X-PDCCH-Config-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sps_ConfigToAddModList_r16 != nil, ie.Sps_ConfigToReleaseList_r16 != nil, ie.Sps_ConfigDeactivationStateList_r16 != nil, ie.BeamFailureRecoverySCellConfig_r16 != nil, ie.Sl_PDCCH_Config_r16 != nil, ie.Sl_V2X_PDCCH_Config_r16 != nil}); err != nil {
				return err
			}

			if ie.Sps_ConfigToAddModList_r16 != nil {
				if err := ie.Sps_ConfigToAddModList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sps_ConfigToReleaseList_r16 != nil {
				if err := ie.Sps_ConfigToReleaseList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sps_ConfigDeactivationStateList_r16 != nil {
				if err := ie.Sps_ConfigDeactivationStateList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.BeamFailureRecoverySCellConfig_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPDownlinkDedicatedExtBeamFailureRecoverySCellConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.BeamFailureRecoverySCellConfig_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.BeamFailureRecoverySCellConfig_r16).Choice {
				case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySCellConfig_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySCellConfig_r16_Setup:
					if err := (*ie.BeamFailureRecoverySCellConfig_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_PDCCH_Config_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPDownlinkDedicatedExtSlPDCCHConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PDCCH_Config_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_PDCCH_Config_r16).Choice {
				case BWP_DownlinkDedicated_Ext_Sl_PDCCH_Config_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_DownlinkDedicated_Ext_Sl_PDCCH_Config_r16_Setup:
					if err := (*ie.Sl_PDCCH_Config_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_V2X_PDCCH_Config_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPDownlinkDedicatedExtSlV2XPDCCHConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_V2X_PDCCH_Config_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_V2X_PDCCH_Config_r16).Choice {
				case BWP_DownlinkDedicated_Ext_Sl_V2X_PDCCH_Config_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_DownlinkDedicated_Ext_Sl_V2X_PDCCH_Config_r16_Setup:
					if err := (*ie.Sl_V2X_PDCCH_Config_r16).Setup.Encode(ex); err != nil {
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
					{Name: "preConfGapStatus-r17", Optional: true},
					{Name: "beamFailureRecoverySpCellConfig-r17", Optional: true},
					{Name: "harq-FeedbackEnablingforSPSactive-r17", Optional: true},
					{Name: "cfr-ConfigMulticast-r17", Optional: true},
					{Name: "dl-PPW-PreConfigToAddModList-r17", Optional: true},
					{Name: "dl-PPW-PreConfigToReleaseList-r17", Optional: true},
					{Name: "nonCellDefiningSSB-r17", Optional: true},
					{Name: "servingCellMO-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PreConfGapStatus_r17 != nil, ie.BeamFailureRecoverySpCellConfig_r17 != nil, ie.Harq_FeedbackEnablingforSPSactive_r17 != nil, ie.Cfr_ConfigMulticast_r17 != nil, ie.Dl_PPW_PreConfigToAddModList_r17 != nil, ie.Dl_PPW_PreConfigToReleaseList_r17 != nil, ie.NonCellDefiningSSB_r17 != nil, ie.ServingCellMO_r17 != nil}); err != nil {
				return err
			}

			if ie.PreConfGapStatus_r17 != nil {
				if err := ex.EncodeBitString(*ie.PreConfGapStatus_r17, bWPDownlinkDedicatedPreConfGapStatusR17Constraints); err != nil {
					return err
				}
			}

			if ie.BeamFailureRecoverySpCellConfig_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPDownlinkDedicatedExtBeamFailureRecoverySpCellConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.BeamFailureRecoverySpCellConfig_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.BeamFailureRecoverySpCellConfig_r17).Choice {
				case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySpCellConfig_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySpCellConfig_r17_Setup:
					if err := (*ie.BeamFailureRecoverySpCellConfig_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_FeedbackEnablingforSPSactive_r17 != nil {
				if err := ex.EncodeBoolean(*ie.Harq_FeedbackEnablingforSPSactive_r17); err != nil {
					return err
				}
			}

			if ie.Cfr_ConfigMulticast_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPDownlinkDedicatedExtCfrConfigMulticastR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Cfr_ConfigMulticast_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Cfr_ConfigMulticast_r17).Choice {
				case BWP_DownlinkDedicated_Ext_Cfr_ConfigMulticast_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_DownlinkDedicated_Ext_Cfr_ConfigMulticast_r17_Setup:
					if err := (*ie.Cfr_ConfigMulticast_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dl_PPW_PreConfigToAddModList_r17 != nil {
				if err := ie.Dl_PPW_PreConfigToAddModList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Dl_PPW_PreConfigToReleaseList_r17 != nil {
				if err := ie.Dl_PPW_PreConfigToReleaseList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.NonCellDefiningSSB_r17 != nil {
				if err := ie.NonCellDefiningSSB_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ServingCellMO_r17 != nil {
				if err := ie.ServingCellMO_r17.Encode(ex); err != nil {
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
					{Name: "tci-InDCI-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Tci_InDCI_r18 != nil}); err != nil {
				return err
			}

			if ie.Tci_InDCI_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPDownlinkDedicatedExtTciInDCIR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Tci_InDCI_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Tci_InDCI_r18).Choice {
				case BWP_DownlinkDedicated_Ext_Tci_InDCI_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_DownlinkDedicated_Ext_Tci_InDCI_r18_Setup:
					if err := (*ie.Tci_InDCI_r18).Setup.Encode(ex); err != nil {
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
					{Name: "sbfd-Config2-Reception-r19", Optional: true},
					{Name: "pathlossOffsetPRACH-DCI-1-0-r19", Optional: true},
					{Name: "prachAssociationDCI-1-0-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sbfd_Config2_Reception_r19 != nil, ie.PathlossOffsetPRACH_DCI_1_0_r19 != nil, ie.PrachAssociationDCI_1_0_r19 != nil}); err != nil {
				return err
			}

			if ie.Sbfd_Config2_Reception_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Sbfd_Config2_Reception_r19, bWPDownlinkDedicatedExtSbfdConfig2ReceptionR19Constraints); err != nil {
					return err
				}
			}

			if ie.PathlossOffsetPRACH_DCI_1_0_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffsetPRACH_DCI_1_0_r19, bWPDownlinkDedicatedExtPathlossOffsetPRACHDCI10R19Constraints); err != nil {
					return err
				}
			}

			if ie.PrachAssociationDCI_1_0_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PrachAssociationDCI_1_0_r19, bWPDownlinkDedicatedExtPrachAssociationDCI10R19Constraints); err != nil {
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

func (ie *BWP_DownlinkDedicated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPDownlinkDedicatedConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pdcch-Config: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Pdcch_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkDedicatedPdcchConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkDedicated_Pdcch_Config_Release:
				(*ie.Pdcch_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Pdcch_Config_Setup:
				(*ie.Pdcch_Config).Setup = new(PDCCH_Config)
				if err := (*ie.Pdcch_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. pdsch-Config: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Pdsch_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkDedicatedPdschConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkDedicated_Pdsch_Config_Release:
				(*ie.Pdsch_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Pdsch_Config_Setup:
				(*ie.Pdsch_Config).Setup = new(PDSCH_Config)
				if err := (*ie.Pdsch_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sps-Config: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Sps_Config = &struct {
				Choice  int
				Release *struct{}
				Setup   *SPS_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkDedicatedSpsConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sps_Config).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkDedicated_Sps_Config_Release:
				(*ie.Sps_Config).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Sps_Config_Setup:
				(*ie.Sps_Config).Setup = new(SPS_Config)
				if err := (*ie.Sps_Config).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. radioLinkMonitoringConfig: choice
	{
		if seq.IsComponentPresent(3) {
			ie.RadioLinkMonitoringConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *RadioLinkMonitoringConfig
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkDedicatedRadioLinkMonitoringConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.RadioLinkMonitoringConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkDedicated_RadioLinkMonitoringConfig_Release:
				(*ie.RadioLinkMonitoringConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_RadioLinkMonitoringConfig_Setup:
				(*ie.RadioLinkMonitoringConfig).Setup = new(RadioLinkMonitoringConfig)
				if err := (*ie.RadioLinkMonitoringConfig).Setup.Decode(d); err != nil {
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
				{Name: "sps-ConfigToAddModList-r16", Optional: true},
				{Name: "sps-ConfigToReleaseList-r16", Optional: true},
				{Name: "sps-ConfigDeactivationStateList-r16", Optional: true},
				{Name: "beamFailureRecoverySCellConfig-r16", Optional: true},
				{Name: "sl-PDCCH-Config-r16", Optional: true},
				{Name: "sl-V2X-PDCCH-Config-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sps_ConfigToAddModList_r16 = new(SPS_ConfigToAddModList_r16)
			if err := ie.Sps_ConfigToAddModList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sps_ConfigToReleaseList_r16 = new(SPS_ConfigToReleaseList_r16)
			if err := ie.Sps_ConfigToReleaseList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Sps_ConfigDeactivationStateList_r16 = new(SPS_ConfigDeactivationStateList_r16)
			if err := ie.Sps_ConfigDeactivationStateList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.BeamFailureRecoverySCellConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BeamFailureRecoveryRSConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPDownlinkDedicatedExtBeamFailureRecoverySCellConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BeamFailureRecoverySCellConfig_r16).Choice = int(index)
			switch index {
			case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySCellConfig_r16_Release:
				(*ie.BeamFailureRecoverySCellConfig_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySCellConfig_r16_Setup:
				(*ie.BeamFailureRecoverySCellConfig_r16).Setup = new(BeamFailureRecoveryRSConfig_r16)
				if err := (*ie.BeamFailureRecoverySCellConfig_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Sl_PDCCH_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_Config
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPDownlinkDedicatedExtSlPDCCHConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PDCCH_Config_r16).Choice = int(index)
			switch index {
			case BWP_DownlinkDedicated_Ext_Sl_PDCCH_Config_r16_Release:
				(*ie.Sl_PDCCH_Config_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Ext_Sl_PDCCH_Config_r16_Setup:
				(*ie.Sl_PDCCH_Config_r16).Setup = new(PDCCH_Config)
				if err := (*ie.Sl_PDCCH_Config_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Sl_V2X_PDCCH_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_Config
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPDownlinkDedicatedExtSlV2XPDCCHConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_V2X_PDCCH_Config_r16).Choice = int(index)
			switch index {
			case BWP_DownlinkDedicated_Ext_Sl_V2X_PDCCH_Config_r16_Release:
				(*ie.Sl_V2X_PDCCH_Config_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Ext_Sl_V2X_PDCCH_Config_r16_Setup:
				(*ie.Sl_V2X_PDCCH_Config_r16).Setup = new(PDCCH_Config)
				if err := (*ie.Sl_V2X_PDCCH_Config_r16).Setup.Decode(dx); err != nil {
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
				{Name: "preConfGapStatus-r17", Optional: true},
				{Name: "beamFailureRecoverySpCellConfig-r17", Optional: true},
				{Name: "harq-FeedbackEnablingforSPSactive-r17", Optional: true},
				{Name: "cfr-ConfigMulticast-r17", Optional: true},
				{Name: "dl-PPW-PreConfigToAddModList-r17", Optional: true},
				{Name: "dl-PPW-PreConfigToReleaseList-r17", Optional: true},
				{Name: "nonCellDefiningSSB-r17", Optional: true},
				{Name: "servingCellMO-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBitString(bWPDownlinkDedicatedPreConfGapStatusR17Constraints)
			if err != nil {
				return err
			}
			ie.PreConfGapStatus_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.BeamFailureRecoverySpCellConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BeamFailureRecoveryRSConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPDownlinkDedicatedExtBeamFailureRecoverySpCellConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BeamFailureRecoverySpCellConfig_r17).Choice = int(index)
			switch index {
			case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySpCellConfig_r17_Release:
				(*ie.BeamFailureRecoverySpCellConfig_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Ext_BeamFailureRecoverySpCellConfig_r17_Setup:
				(*ie.BeamFailureRecoverySpCellConfig_r17).Setup = new(BeamFailureRecoveryRSConfig_r16)
				if err := (*ie.BeamFailureRecoverySpCellConfig_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Harq_FeedbackEnablingforSPSactive_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Cfr_ConfigMulticast_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *CFR_ConfigMulticast_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPDownlinkDedicatedExtCfrConfigMulticastR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cfr_ConfigMulticast_r17).Choice = int(index)
			switch index {
			case BWP_DownlinkDedicated_Ext_Cfr_ConfigMulticast_r17_Release:
				(*ie.Cfr_ConfigMulticast_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Ext_Cfr_ConfigMulticast_r17_Setup:
				(*ie.Cfr_ConfigMulticast_r17).Setup = new(CFR_ConfigMulticast_r17)
				if err := (*ie.Cfr_ConfigMulticast_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Dl_PPW_PreConfigToAddModList_r17 = new(DL_PPW_PreConfigToAddModList_r17)
			if err := ie.Dl_PPW_PreConfigToAddModList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Dl_PPW_PreConfigToReleaseList_r17 = new(DL_PPW_PreConfigToReleaseList_r17)
			if err := ie.Dl_PPW_PreConfigToReleaseList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.NonCellDefiningSSB_r17 = new(NonCellDefiningSSB_r17)
			if err := ie.NonCellDefiningSSB_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.ServingCellMO_r17 = new(MeasObjectId)
			if err := ie.ServingCellMO_r17.Decode(dx); err != nil {
				return err
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
				{Name: "tci-InDCI-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Tci_InDCI_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *TCI_InDCI_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPDownlinkDedicatedExtTciInDCIR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Tci_InDCI_r18).Choice = int(index)
			switch index {
			case BWP_DownlinkDedicated_Ext_Tci_InDCI_r18_Release:
				(*ie.Tci_InDCI_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicated_Ext_Tci_InDCI_r18_Setup:
				(*ie.Tci_InDCI_r18).Setup = new(TCI_InDCI_r18)
				if err := (*ie.Tci_InDCI_r18).Setup.Decode(dx); err != nil {
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
				{Name: "sbfd-Config2-Reception-r19", Optional: true},
				{Name: "pathlossOffsetPRACH-DCI-1-0-r19", Optional: true},
				{Name: "prachAssociationDCI-1-0-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bWPDownlinkDedicatedExtSbfdConfig2ReceptionR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_Config2_Reception_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bWPDownlinkDedicatedExtPathlossOffsetPRACHDCI10R19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffsetPRACH_DCI_1_0_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bWPDownlinkDedicatedExtPrachAssociationDCI10R19Constraints)
			if err != nil {
				return err
			}
			ie.PrachAssociationDCI_1_0_r19 = &v
		}
	}

	return nil
}
