// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-MeasConfig (line 7010).

var cSIMeasConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nzp-CSI-RS-ResourceToAddModList", Optional: true},
		{Name: "nzp-CSI-RS-ResourceToReleaseList", Optional: true},
		{Name: "nzp-CSI-RS-ResourceSetToAddModList", Optional: true},
		{Name: "nzp-CSI-RS-ResourceSetToReleaseList", Optional: true},
		{Name: "csi-IM-ResourceToAddModList", Optional: true},
		{Name: "csi-IM-ResourceToReleaseList", Optional: true},
		{Name: "csi-IM-ResourceSetToAddModList", Optional: true},
		{Name: "csi-IM-ResourceSetToReleaseList", Optional: true},
		{Name: "csi-SSB-ResourceSetToAddModList", Optional: true},
		{Name: "csi-SSB-ResourceSetToReleaseList", Optional: true},
		{Name: "csi-ResourceConfigToAddModList", Optional: true},
		{Name: "csi-ResourceConfigToReleaseList", Optional: true},
		{Name: "csi-ReportConfigToAddModList", Optional: true},
		{Name: "csi-ReportConfigToReleaseList", Optional: true},
		{Name: "reportTriggerSize", Optional: true},
		{Name: "aperiodicTriggerStateList", Optional: true},
		{Name: "semiPersistentOnPUSCH-TriggerStateList", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var cSIMeasConfigNzpCSIRSResourceToAddModListConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_Resources)

var cSIMeasConfigNzpCSIRSResourceToReleaseListConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_Resources)

var cSIMeasConfigNzpCSIRSResourceSetToAddModListConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourceSets)

var cSIMeasConfigNzpCSIRSResourceSetToReleaseListConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourceSets)

var cSIMeasConfigCsiIMResourceToAddModListConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_Resources)

var cSIMeasConfigCsiIMResourceToReleaseListConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_Resources)

var cSIMeasConfigCsiIMResourceSetToAddModListConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_ResourceSets)

var cSIMeasConfigCsiIMResourceSetToReleaseListConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_ResourceSets)

var cSIMeasConfigCsiSSBResourceSetToAddModListConstraints = per.SizeRange(1, common.MaxNrofCSI_SSB_ResourceSets)

var cSIMeasConfigCsiSSBResourceSetToReleaseListConstraints = per.SizeRange(1, common.MaxNrofCSI_SSB_ResourceSets)

var cSIMeasConfigCsiResourceConfigToAddModListConstraints = per.SizeRange(1, common.MaxNrofCSI_ResourceConfigurations)

var cSIMeasConfigCsiResourceConfigToReleaseListConstraints = per.SizeRange(1, common.MaxNrofCSI_ResourceConfigurations)

var cSIMeasConfigCsiReportConfigToAddModListConstraints = per.SizeRange(1, common.MaxNrofCSI_ReportConfigurations)

var cSIMeasConfigCsiReportConfigToReleaseListConstraints = per.SizeRange(1, common.MaxNrofCSI_ReportConfigurations)

var cSIMeasConfigReportTriggerSizeConstraints = per.Constrained(0, 6)

var cSI_MeasConfigAperiodicTriggerStateListConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CSI_MeasConfig_AperiodicTriggerStateList_Release = 0
	CSI_MeasConfig_AperiodicTriggerStateList_Setup   = 1
)

var cSI_MeasConfigSemiPersistentOnPUSCHTriggerStateListConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CSI_MeasConfig_SemiPersistentOnPUSCH_TriggerStateList_Release = 0
	CSI_MeasConfig_SemiPersistentOnPUSCH_TriggerStateList_Setup   = 1
)

var cSIMeasConfigReportTriggerSizeDCI02R16Constraints = per.Constrained(0, 6)

var cSIMeasConfigExtSCellActivationRSConfigToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofSCellActRS_r17)

var cSIMeasConfigExtSCellActivationRSConfigToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofSCellActRS_r17)

var cSIMeasConfigExtLtmCSIReportConfigToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofLTM_CSI_ReportConfigurations_r18)

var cSIMeasConfigExtLtmCSIReportConfigToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofLTM_CSI_ReportConfigurations_r18)

var cSIMeasConfigExtCliRSSIMeasResourceToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofCLI_RSSI_MeasResources_r19)

var cSIMeasConfigExtCliRSSIMeasResourceToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofCLI_RSSI_MeasResources_r19)

var cSIMeasConfigExtCliRSSIMeasResourceSetToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofCLI_RSSI_MeasResourceSets_r19)

var cSIMeasConfigExtCliRSSIMeasResourceSetToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofCLI_RSSI_MeasResourceSets_r19)

var cSIMeasConfigExtSrsRSRPMeasResourceToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofSRS_RSRP_MeasResources_r19)

var cSIMeasConfigExtSrsRSRPMeasResourceToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofSRS_RSRP_MeasResources_r19)

var cSIMeasConfigExtSrsRSRPMeasResourceSetToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofSRS_RSRP_MeasResourceSets_r19)

var cSIMeasConfigExtSrsRSRPMeasResourceSetToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofSRS_RSRP_MeasResourceSets_r19)

var cSIMeasConfigExtCsiLoggedMeasurementConfigToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofLoggedMeasurementConfigurations_r19)

var cSIMeasConfigExtCsiLoggedMeasurementConfigToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofLoggedMeasurementConfigurations_r19)

type CSI_MeasConfig struct {
	Nzp_CSI_RS_ResourceToAddModList     []NZP_CSI_RS_Resource
	Nzp_CSI_RS_ResourceToReleaseList    []NZP_CSI_RS_ResourceId
	Nzp_CSI_RS_ResourceSetToAddModList  []NZP_CSI_RS_ResourceSet
	Nzp_CSI_RS_ResourceSetToReleaseList []NZP_CSI_RS_ResourceSetId
	Csi_IM_ResourceToAddModList         []CSI_IM_Resource
	Csi_IM_ResourceToReleaseList        []CSI_IM_ResourceId
	Csi_IM_ResourceSetToAddModList      []CSI_IM_ResourceSet
	Csi_IM_ResourceSetToReleaseList     []CSI_IM_ResourceSetId
	Csi_SSB_ResourceSetToAddModList     []CSI_SSB_ResourceSet
	Csi_SSB_ResourceSetToReleaseList    []CSI_SSB_ResourceSetId
	Csi_ResourceConfigToAddModList      []CSI_ResourceConfig
	Csi_ResourceConfigToReleaseList     []CSI_ResourceConfigId
	Csi_ReportConfigToAddModList        []CSI_ReportConfig
	Csi_ReportConfigToReleaseList       []CSI_ReportConfigId
	ReportTriggerSize                   *int64
	AperiodicTriggerStateList           *struct {
		Choice  int
		Release *struct{}
		Setup   *CSI_AperiodicTriggerStateList
	}
	SemiPersistentOnPUSCH_TriggerStateList *struct {
		Choice  int
		Release *struct{}
		Setup   *CSI_SemiPersistentOnPUSCH_TriggerStateList
	}
	ReportTriggerSizeDCI_0_2_r16                 *int64
	SCellActivationRS_ConfigToAddModList_r17     []SCellActivationRS_Config_r17
	SCellActivationRS_ConfigToReleaseList_r17    []SCellActivationRS_ConfigId_r17
	Ltm_CSI_ReportConfigToAddModList_r18         []LTM_CSI_ReportConfig_r18
	Ltm_CSI_ReportConfigToReleaseList_r18        []LTM_CSI_ReportConfigId_r18
	Cli_RSSI_MeasResourceToAddModList_r19        []CLI_RSSI_MeasResource_r19
	Cli_RSSI_MeasResourceToReleaseList_r19       []CLI_RSSI_MeasResourceId_r19
	Cli_RSSI_MeasResourceSetToAddModList_r19     []CLI_RSSI_MeasResourceSet_r19
	Cli_RSSI_MeasResourceSetToReleaseList_r19    []CLI_RSSI_MeasResourceSetId_r19
	Srs_RSRP_MeasResourceToAddModList_r19        []SRS_RSRP_MeasResource_r19
	Srs_RSRP_MeasResourceToReleaseList_r19       []SRS_RSRP_MeasResourceId_r19
	Srs_RSRP_MeasResourceSetToAddModList_r19     []SRS_RSRP_MeasResourceSet_r19
	Srs_RSRP_MeasResourceSetToReleaseList_r19    []SRS_RSRP_MeasResourceSetId_r19
	Csi_LoggedMeasurementConfigToAddModList_r19  []CSI_LoggedMeasurementConfig_r19
	Csi_LoggedMeasurementConfigToReleaseList_r19 []CSI_LoggedMeasurementConfigId_r19
}

func (ie *CSI_MeasConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIMeasConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReportTriggerSizeDCI_0_2_r16 != nil
	hasExtGroup1 := ie.SCellActivationRS_ConfigToAddModList_r17 != nil || ie.SCellActivationRS_ConfigToReleaseList_r17 != nil
	hasExtGroup2 := ie.Ltm_CSI_ReportConfigToAddModList_r18 != nil || ie.Ltm_CSI_ReportConfigToReleaseList_r18 != nil
	hasExtGroup3 := ie.Cli_RSSI_MeasResourceToAddModList_r19 != nil || ie.Cli_RSSI_MeasResourceToReleaseList_r19 != nil || ie.Cli_RSSI_MeasResourceSetToAddModList_r19 != nil || ie.Cli_RSSI_MeasResourceSetToReleaseList_r19 != nil || ie.Srs_RSRP_MeasResourceToAddModList_r19 != nil || ie.Srs_RSRP_MeasResourceToReleaseList_r19 != nil || ie.Srs_RSRP_MeasResourceSetToAddModList_r19 != nil || ie.Srs_RSRP_MeasResourceSetToReleaseList_r19 != nil || ie.Csi_LoggedMeasurementConfigToAddModList_r19 != nil || ie.Csi_LoggedMeasurementConfigToReleaseList_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nzp_CSI_RS_ResourceToAddModList != nil, ie.Nzp_CSI_RS_ResourceToReleaseList != nil, ie.Nzp_CSI_RS_ResourceSetToAddModList != nil, ie.Nzp_CSI_RS_ResourceSetToReleaseList != nil, ie.Csi_IM_ResourceToAddModList != nil, ie.Csi_IM_ResourceToReleaseList != nil, ie.Csi_IM_ResourceSetToAddModList != nil, ie.Csi_IM_ResourceSetToReleaseList != nil, ie.Csi_SSB_ResourceSetToAddModList != nil, ie.Csi_SSB_ResourceSetToReleaseList != nil, ie.Csi_ResourceConfigToAddModList != nil, ie.Csi_ResourceConfigToReleaseList != nil, ie.Csi_ReportConfigToAddModList != nil, ie.Csi_ReportConfigToReleaseList != nil, ie.ReportTriggerSize != nil, ie.AperiodicTriggerStateList != nil, ie.SemiPersistentOnPUSCH_TriggerStateList != nil}); err != nil {
		return err
	}

	// 3. nzp-CSI-RS-ResourceToAddModList: sequence-of
	{
		if ie.Nzp_CSI_RS_ResourceToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigNzpCSIRSResourceToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Nzp_CSI_RS_ResourceToAddModList))); err != nil {
				return err
			}
			for i := range ie.Nzp_CSI_RS_ResourceToAddModList {
				if err := ie.Nzp_CSI_RS_ResourceToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. nzp-CSI-RS-ResourceToReleaseList: sequence-of
	{
		if ie.Nzp_CSI_RS_ResourceToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigNzpCSIRSResourceToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Nzp_CSI_RS_ResourceToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Nzp_CSI_RS_ResourceToReleaseList {
				if err := ie.Nzp_CSI_RS_ResourceToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. nzp-CSI-RS-ResourceSetToAddModList: sequence-of
	{
		if ie.Nzp_CSI_RS_ResourceSetToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigNzpCSIRSResourceSetToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Nzp_CSI_RS_ResourceSetToAddModList))); err != nil {
				return err
			}
			for i := range ie.Nzp_CSI_RS_ResourceSetToAddModList {
				if err := ie.Nzp_CSI_RS_ResourceSetToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. nzp-CSI-RS-ResourceSetToReleaseList: sequence-of
	{
		if ie.Nzp_CSI_RS_ResourceSetToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigNzpCSIRSResourceSetToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Nzp_CSI_RS_ResourceSetToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Nzp_CSI_RS_ResourceSetToReleaseList {
				if err := ie.Nzp_CSI_RS_ResourceSetToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. csi-IM-ResourceToAddModList: sequence-of
	{
		if ie.Csi_IM_ResourceToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiIMResourceToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_IM_ResourceToAddModList))); err != nil {
				return err
			}
			for i := range ie.Csi_IM_ResourceToAddModList {
				if err := ie.Csi_IM_ResourceToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. csi-IM-ResourceToReleaseList: sequence-of
	{
		if ie.Csi_IM_ResourceToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiIMResourceToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_IM_ResourceToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Csi_IM_ResourceToReleaseList {
				if err := ie.Csi_IM_ResourceToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. csi-IM-ResourceSetToAddModList: sequence-of
	{
		if ie.Csi_IM_ResourceSetToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiIMResourceSetToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_IM_ResourceSetToAddModList))); err != nil {
				return err
			}
			for i := range ie.Csi_IM_ResourceSetToAddModList {
				if err := ie.Csi_IM_ResourceSetToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. csi-IM-ResourceSetToReleaseList: sequence-of
	{
		if ie.Csi_IM_ResourceSetToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiIMResourceSetToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_IM_ResourceSetToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Csi_IM_ResourceSetToReleaseList {
				if err := ie.Csi_IM_ResourceSetToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 11. csi-SSB-ResourceSetToAddModList: sequence-of
	{
		if ie.Csi_SSB_ResourceSetToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiSSBResourceSetToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_SSB_ResourceSetToAddModList))); err != nil {
				return err
			}
			for i := range ie.Csi_SSB_ResourceSetToAddModList {
				if err := ie.Csi_SSB_ResourceSetToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 12. csi-SSB-ResourceSetToReleaseList: sequence-of
	{
		if ie.Csi_SSB_ResourceSetToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiSSBResourceSetToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_SSB_ResourceSetToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Csi_SSB_ResourceSetToReleaseList {
				if err := ie.Csi_SSB_ResourceSetToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. csi-ResourceConfigToAddModList: sequence-of
	{
		if ie.Csi_ResourceConfigToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiResourceConfigToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_ResourceConfigToAddModList))); err != nil {
				return err
			}
			for i := range ie.Csi_ResourceConfigToAddModList {
				if err := ie.Csi_ResourceConfigToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 14. csi-ResourceConfigToReleaseList: sequence-of
	{
		if ie.Csi_ResourceConfigToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiResourceConfigToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_ResourceConfigToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Csi_ResourceConfigToReleaseList {
				if err := ie.Csi_ResourceConfigToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 15. csi-ReportConfigToAddModList: sequence-of
	{
		if ie.Csi_ReportConfigToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiReportConfigToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_ReportConfigToAddModList))); err != nil {
				return err
			}
			for i := range ie.Csi_ReportConfigToAddModList {
				if err := ie.Csi_ReportConfigToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 16. csi-ReportConfigToReleaseList: sequence-of
	{
		if ie.Csi_ReportConfigToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cSIMeasConfigCsiReportConfigToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_ReportConfigToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Csi_ReportConfigToReleaseList {
				if err := ie.Csi_ReportConfigToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 17. reportTriggerSize: integer
	{
		if ie.ReportTriggerSize != nil {
			if err := e.EncodeInteger(*ie.ReportTriggerSize, cSIMeasConfigReportTriggerSizeConstraints); err != nil {
				return err
			}
		}
	}

	// 18. aperiodicTriggerStateList: choice
	{
		if ie.AperiodicTriggerStateList != nil {
			choiceEnc := e.NewChoiceEncoder(cSI_MeasConfigAperiodicTriggerStateListConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.AperiodicTriggerStateList).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.AperiodicTriggerStateList).Choice {
			case CSI_MeasConfig_AperiodicTriggerStateList_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case CSI_MeasConfig_AperiodicTriggerStateList_Setup:
				if err := (*ie.AperiodicTriggerStateList).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.AperiodicTriggerStateList).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 19. semiPersistentOnPUSCH-TriggerStateList: choice
	{
		if ie.SemiPersistentOnPUSCH_TriggerStateList != nil {
			choiceEnc := e.NewChoiceEncoder(cSI_MeasConfigSemiPersistentOnPUSCHTriggerStateListConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SemiPersistentOnPUSCH_TriggerStateList).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SemiPersistentOnPUSCH_TriggerStateList).Choice {
			case CSI_MeasConfig_SemiPersistentOnPUSCH_TriggerStateList_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case CSI_MeasConfig_SemiPersistentOnPUSCH_TriggerStateList_Setup:
				if err := (*ie.SemiPersistentOnPUSCH_TriggerStateList).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SemiPersistentOnPUSCH_TriggerStateList).Choice), Constraint: "index out of range"}
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
					{Name: "reportTriggerSizeDCI-0-2-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportTriggerSizeDCI_0_2_r16 != nil}); err != nil {
				return err
			}

			if ie.ReportTriggerSizeDCI_0_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.ReportTriggerSizeDCI_0_2_r16, cSIMeasConfigReportTriggerSizeDCI02R16Constraints); err != nil {
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
					{Name: "sCellActivationRS-ConfigToAddModList-r17", Optional: true},
					{Name: "sCellActivationRS-ConfigToReleaseList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SCellActivationRS_ConfigToAddModList_r17 != nil, ie.SCellActivationRS_ConfigToReleaseList_r17 != nil}); err != nil {
				return err
			}

			if ie.SCellActivationRS_ConfigToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtSCellActivationRSConfigToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SCellActivationRS_ConfigToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.SCellActivationRS_ConfigToAddModList_r17 {
					if err := ie.SCellActivationRS_ConfigToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SCellActivationRS_ConfigToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtSCellActivationRSConfigToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SCellActivationRS_ConfigToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.SCellActivationRS_ConfigToReleaseList_r17 {
					if err := ie.SCellActivationRS_ConfigToReleaseList_r17[i].Encode(ex); err != nil {
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
					{Name: "ltm-CSI-ReportConfigToAddModList-r18", Optional: true},
					{Name: "ltm-CSI-ReportConfigToReleaseList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_CSI_ReportConfigToAddModList_r18 != nil, ie.Ltm_CSI_ReportConfigToReleaseList_r18 != nil}); err != nil {
				return err
			}

			if ie.Ltm_CSI_ReportConfigToAddModList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtLtmCSIReportConfigToAddModListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Ltm_CSI_ReportConfigToAddModList_r18))); err != nil {
					return err
				}
				for i := range ie.Ltm_CSI_ReportConfigToAddModList_r18 {
					if err := ie.Ltm_CSI_ReportConfigToAddModList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ltm_CSI_ReportConfigToReleaseList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtLtmCSIReportConfigToReleaseListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Ltm_CSI_ReportConfigToReleaseList_r18))); err != nil {
					return err
				}
				for i := range ie.Ltm_CSI_ReportConfigToReleaseList_r18 {
					if err := ie.Ltm_CSI_ReportConfigToReleaseList_r18[i].Encode(ex); err != nil {
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
					{Name: "cli-RSSI-MeasResourceToAddModList-r19", Optional: true},
					{Name: "cli-RSSI-MeasResourceToReleaseList-r19", Optional: true},
					{Name: "cli-RSSI-MeasResourceSetToAddModList-r19", Optional: true},
					{Name: "cli-RSSI-MeasResourceSetToReleaseList-r19", Optional: true},
					{Name: "srs-RSRP-MeasResourceToAddModList-r19", Optional: true},
					{Name: "srs-RSRP-MeasResourceToReleaseList-r19", Optional: true},
					{Name: "srs-RSRP-MeasResourceSetToAddModList-r19", Optional: true},
					{Name: "srs-RSRP-MeasResourceSetToReleaseList-r19", Optional: true},
					{Name: "csi-LoggedMeasurementConfigToAddModList-r19", Optional: true},
					{Name: "csi-LoggedMeasurementConfigToReleaseList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cli_RSSI_MeasResourceToAddModList_r19 != nil, ie.Cli_RSSI_MeasResourceToReleaseList_r19 != nil, ie.Cli_RSSI_MeasResourceSetToAddModList_r19 != nil, ie.Cli_RSSI_MeasResourceSetToReleaseList_r19 != nil, ie.Srs_RSRP_MeasResourceToAddModList_r19 != nil, ie.Srs_RSRP_MeasResourceToReleaseList_r19 != nil, ie.Srs_RSRP_MeasResourceSetToAddModList_r19 != nil, ie.Srs_RSRP_MeasResourceSetToReleaseList_r19 != nil, ie.Csi_LoggedMeasurementConfigToAddModList_r19 != nil, ie.Csi_LoggedMeasurementConfigToReleaseList_r19 != nil}); err != nil {
				return err
			}

			if ie.Cli_RSSI_MeasResourceToAddModList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtCliRSSIMeasResourceToAddModListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Cli_RSSI_MeasResourceToAddModList_r19))); err != nil {
					return err
				}
				for i := range ie.Cli_RSSI_MeasResourceToAddModList_r19 {
					if err := ie.Cli_RSSI_MeasResourceToAddModList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Cli_RSSI_MeasResourceToReleaseList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtCliRSSIMeasResourceToReleaseListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Cli_RSSI_MeasResourceToReleaseList_r19))); err != nil {
					return err
				}
				for i := range ie.Cli_RSSI_MeasResourceToReleaseList_r19 {
					if err := ie.Cli_RSSI_MeasResourceToReleaseList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Cli_RSSI_MeasResourceSetToAddModList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtCliRSSIMeasResourceSetToAddModListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Cli_RSSI_MeasResourceSetToAddModList_r19))); err != nil {
					return err
				}
				for i := range ie.Cli_RSSI_MeasResourceSetToAddModList_r19 {
					if err := ie.Cli_RSSI_MeasResourceSetToAddModList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Cli_RSSI_MeasResourceSetToReleaseList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtCliRSSIMeasResourceSetToReleaseListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Cli_RSSI_MeasResourceSetToReleaseList_r19))); err != nil {
					return err
				}
				for i := range ie.Cli_RSSI_MeasResourceSetToReleaseList_r19 {
					if err := ie.Cli_RSSI_MeasResourceSetToReleaseList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_RSRP_MeasResourceToAddModList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtSrsRSRPMeasResourceToAddModListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_RSRP_MeasResourceToAddModList_r19))); err != nil {
					return err
				}
				for i := range ie.Srs_RSRP_MeasResourceToAddModList_r19 {
					if err := ie.Srs_RSRP_MeasResourceToAddModList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_RSRP_MeasResourceToReleaseList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtSrsRSRPMeasResourceToReleaseListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_RSRP_MeasResourceToReleaseList_r19))); err != nil {
					return err
				}
				for i := range ie.Srs_RSRP_MeasResourceToReleaseList_r19 {
					if err := ie.Srs_RSRP_MeasResourceToReleaseList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_RSRP_MeasResourceSetToAddModList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtSrsRSRPMeasResourceSetToAddModListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_RSRP_MeasResourceSetToAddModList_r19))); err != nil {
					return err
				}
				for i := range ie.Srs_RSRP_MeasResourceSetToAddModList_r19 {
					if err := ie.Srs_RSRP_MeasResourceSetToAddModList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_RSRP_MeasResourceSetToReleaseList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtSrsRSRPMeasResourceSetToReleaseListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_RSRP_MeasResourceSetToReleaseList_r19))); err != nil {
					return err
				}
				for i := range ie.Srs_RSRP_MeasResourceSetToReleaseList_r19 {
					if err := ie.Srs_RSRP_MeasResourceSetToReleaseList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Csi_LoggedMeasurementConfigToAddModList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtCsiLoggedMeasurementConfigToAddModListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Csi_LoggedMeasurementConfigToAddModList_r19))); err != nil {
					return err
				}
				for i := range ie.Csi_LoggedMeasurementConfigToAddModList_r19 {
					if err := ie.Csi_LoggedMeasurementConfigToAddModList_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Csi_LoggedMeasurementConfigToReleaseList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIMeasConfigExtCsiLoggedMeasurementConfigToReleaseListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Csi_LoggedMeasurementConfigToReleaseList_r19))); err != nil {
					return err
				}
				for i := range ie.Csi_LoggedMeasurementConfigToReleaseList_r19 {
					if err := ie.Csi_LoggedMeasurementConfigToReleaseList_r19[i].Encode(ex); err != nil {
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

func (ie *CSI_MeasConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIMeasConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. nzp-CSI-RS-ResourceToAddModList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigNzpCSIRSResourceToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Nzp_CSI_RS_ResourceToAddModList = make([]NZP_CSI_RS_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Nzp_CSI_RS_ResourceToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. nzp-CSI-RS-ResourceToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigNzpCSIRSResourceToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Nzp_CSI_RS_ResourceToReleaseList = make([]NZP_CSI_RS_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Nzp_CSI_RS_ResourceToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. nzp-CSI-RS-ResourceSetToAddModList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigNzpCSIRSResourceSetToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Nzp_CSI_RS_ResourceSetToAddModList = make([]NZP_CSI_RS_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Nzp_CSI_RS_ResourceSetToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. nzp-CSI-RS-ResourceSetToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigNzpCSIRSResourceSetToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Nzp_CSI_RS_ResourceSetToReleaseList = make([]NZP_CSI_RS_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Nzp_CSI_RS_ResourceSetToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. csi-IM-ResourceToAddModList: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiIMResourceToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_IM_ResourceToAddModList = make([]CSI_IM_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_IM_ResourceToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. csi-IM-ResourceToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiIMResourceToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_IM_ResourceToReleaseList = make([]CSI_IM_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_IM_ResourceToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. csi-IM-ResourceSetToAddModList: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiIMResourceSetToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_IM_ResourceSetToAddModList = make([]CSI_IM_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_IM_ResourceSetToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. csi-IM-ResourceSetToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiIMResourceSetToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_IM_ResourceSetToReleaseList = make([]CSI_IM_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_IM_ResourceSetToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. csi-SSB-ResourceSetToAddModList: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiSSBResourceSetToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_SSB_ResourceSetToAddModList = make([]CSI_SSB_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_SSB_ResourceSetToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. csi-SSB-ResourceSetToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiSSBResourceSetToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_SSB_ResourceSetToReleaseList = make([]CSI_SSB_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_SSB_ResourceSetToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. csi-ResourceConfigToAddModList: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiResourceConfigToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_ResourceConfigToAddModList = make([]CSI_ResourceConfig, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_ResourceConfigToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 14. csi-ResourceConfigToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(11) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiResourceConfigToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_ResourceConfigToReleaseList = make([]CSI_ResourceConfigId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_ResourceConfigToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 15. csi-ReportConfigToAddModList: sequence-of
	{
		if seq.IsComponentPresent(12) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiReportConfigToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_ReportConfigToAddModList = make([]CSI_ReportConfig, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_ReportConfigToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 16. csi-ReportConfigToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(13) {
			seqOf := d.NewSequenceOfDecoder(cSIMeasConfigCsiReportConfigToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_ReportConfigToReleaseList = make([]CSI_ReportConfigId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_ReportConfigToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 17. reportTriggerSize: integer
	{
		if seq.IsComponentPresent(14) {
			v, err := d.DecodeInteger(cSIMeasConfigReportTriggerSizeConstraints)
			if err != nil {
				return err
			}
			ie.ReportTriggerSize = &v
		}
	}

	// 18. aperiodicTriggerStateList: choice
	{
		if seq.IsComponentPresent(15) {
			ie.AperiodicTriggerStateList = &struct {
				Choice  int
				Release *struct{}
				Setup   *CSI_AperiodicTriggerStateList
			}{}
			choiceDec := d.NewChoiceDecoder(cSI_MeasConfigAperiodicTriggerStateListConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AperiodicTriggerStateList).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CSI_MeasConfig_AperiodicTriggerStateList_Release:
				(*ie.AperiodicTriggerStateList).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case CSI_MeasConfig_AperiodicTriggerStateList_Setup:
				(*ie.AperiodicTriggerStateList).Setup = new(CSI_AperiodicTriggerStateList)
				if err := (*ie.AperiodicTriggerStateList).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 19. semiPersistentOnPUSCH-TriggerStateList: choice
	{
		if seq.IsComponentPresent(16) {
			ie.SemiPersistentOnPUSCH_TriggerStateList = &struct {
				Choice  int
				Release *struct{}
				Setup   *CSI_SemiPersistentOnPUSCH_TriggerStateList
			}{}
			choiceDec := d.NewChoiceDecoder(cSI_MeasConfigSemiPersistentOnPUSCHTriggerStateListConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SemiPersistentOnPUSCH_TriggerStateList).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CSI_MeasConfig_SemiPersistentOnPUSCH_TriggerStateList_Release:
				(*ie.SemiPersistentOnPUSCH_TriggerStateList).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case CSI_MeasConfig_SemiPersistentOnPUSCH_TriggerStateList_Setup:
				(*ie.SemiPersistentOnPUSCH_TriggerStateList).Setup = new(CSI_SemiPersistentOnPUSCH_TriggerStateList)
				if err := (*ie.SemiPersistentOnPUSCH_TriggerStateList).Setup.Decode(d); err != nil {
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
				{Name: "reportTriggerSizeDCI-0-2-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(cSIMeasConfigReportTriggerSizeDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.ReportTriggerSizeDCI_0_2_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sCellActivationRS-ConfigToAddModList-r17", Optional: true},
				{Name: "sCellActivationRS-ConfigToReleaseList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtSCellActivationRSConfigToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SCellActivationRS_ConfigToAddModList_r17 = make([]SCellActivationRS_Config_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SCellActivationRS_ConfigToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtSCellActivationRSConfigToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SCellActivationRS_ConfigToReleaseList_r17 = make([]SCellActivationRS_ConfigId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SCellActivationRS_ConfigToReleaseList_r17[i].Decode(dx); err != nil {
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
				{Name: "ltm-CSI-ReportConfigToAddModList-r18", Optional: true},
				{Name: "ltm-CSI-ReportConfigToReleaseList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtLtmCSIReportConfigToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_CSI_ReportConfigToAddModList_r18 = make([]LTM_CSI_ReportConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_CSI_ReportConfigToAddModList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtLtmCSIReportConfigToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_CSI_ReportConfigToReleaseList_r18 = make([]LTM_CSI_ReportConfigId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_CSI_ReportConfigToReleaseList_r18[i].Decode(dx); err != nil {
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
				{Name: "cli-RSSI-MeasResourceToAddModList-r19", Optional: true},
				{Name: "cli-RSSI-MeasResourceToReleaseList-r19", Optional: true},
				{Name: "cli-RSSI-MeasResourceSetToAddModList-r19", Optional: true},
				{Name: "cli-RSSI-MeasResourceSetToReleaseList-r19", Optional: true},
				{Name: "srs-RSRP-MeasResourceToAddModList-r19", Optional: true},
				{Name: "srs-RSRP-MeasResourceToReleaseList-r19", Optional: true},
				{Name: "srs-RSRP-MeasResourceSetToAddModList-r19", Optional: true},
				{Name: "srs-RSRP-MeasResourceSetToReleaseList-r19", Optional: true},
				{Name: "csi-LoggedMeasurementConfigToAddModList-r19", Optional: true},
				{Name: "csi-LoggedMeasurementConfigToReleaseList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtCliRSSIMeasResourceToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cli_RSSI_MeasResourceToAddModList_r19 = make([]CLI_RSSI_MeasResource_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cli_RSSI_MeasResourceToAddModList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtCliRSSIMeasResourceToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cli_RSSI_MeasResourceToReleaseList_r19 = make([]CLI_RSSI_MeasResourceId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cli_RSSI_MeasResourceToReleaseList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtCliRSSIMeasResourceSetToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cli_RSSI_MeasResourceSetToAddModList_r19 = make([]CLI_RSSI_MeasResourceSet_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cli_RSSI_MeasResourceSetToAddModList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtCliRSSIMeasResourceSetToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cli_RSSI_MeasResourceSetToReleaseList_r19 = make([]CLI_RSSI_MeasResourceSetId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cli_RSSI_MeasResourceSetToReleaseList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtSrsRSRPMeasResourceToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_RSRP_MeasResourceToAddModList_r19 = make([]SRS_RSRP_MeasResource_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_RSRP_MeasResourceToAddModList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtSrsRSRPMeasResourceToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_RSRP_MeasResourceToReleaseList_r19 = make([]SRS_RSRP_MeasResourceId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_RSRP_MeasResourceToReleaseList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtSrsRSRPMeasResourceSetToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_RSRP_MeasResourceSetToAddModList_r19 = make([]SRS_RSRP_MeasResourceSet_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_RSRP_MeasResourceSetToAddModList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtSrsRSRPMeasResourceSetToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_RSRP_MeasResourceSetToReleaseList_r19 = make([]SRS_RSRP_MeasResourceSetId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_RSRP_MeasResourceSetToReleaseList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtCsiLoggedMeasurementConfigToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_LoggedMeasurementConfigToAddModList_r19 = make([]CSI_LoggedMeasurementConfig_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_LoggedMeasurementConfigToAddModList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			seqOf := dx.NewSequenceOfDecoder(cSIMeasConfigExtCsiLoggedMeasurementConfigToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_LoggedMeasurementConfigToReleaseList_r19 = make([]CSI_LoggedMeasurementConfigId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_LoggedMeasurementConfigToReleaseList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
