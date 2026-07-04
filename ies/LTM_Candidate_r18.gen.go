// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-Candidate-r18 (line 8675).

var lTMCandidateR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-CandidateId-r18"},
		{Name: "ltm-CandidatePCI-r18", Optional: true},
		{Name: "ltm-SSB-Config-r18", Optional: true},
		{Name: "ltm-CandidateConfig-r18", Optional: true},
		{Name: "ltm-ConfigComplete-r18", Optional: true},
		{Name: "ltm-EarlyUL-SyncConfig-r18", Optional: true},
		{Name: "ltm-EarlyUL-SyncConfigSUL-r18", Optional: true},
		{Name: "ltm-TCI-Info-r18", Optional: true},
		{Name: "ltm-NoResetID-r18", Optional: true},
		{Name: "ltm-UE-MeasuredTA-ID-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var lTMCandidateR18LtmCandidateConfigR18Constraints = per.SizeConstraints{}

const (
	LTM_Candidate_r18_Ltm_ConfigComplete_r18_True = 0
)

var lTMCandidateR18LtmConfigCompleteR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_Candidate_r18_Ltm_ConfigComplete_r18_True},
}

var lTMCandidateR18LtmEarlyULSyncConfigR18Constraints = per.SizeConstraints{}

var lTMCandidateR18LtmEarlyULSyncConfigSULR18Constraints = per.SizeConstraints{}

var lTMCandidateR18LtmNoResetIDR18Constraints = per.Constrained(1, common.MaxNrofLTM_Configs_Plus1_r18)

var lTMCandidateR18ExtLtmExecutionConditionR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LTM_Candidate_r18_Ext_Ltm_ExecutionCondition_r19_Release = 0
	LTM_Candidate_r18_Ext_Ltm_ExecutionCondition_r19_Setup   = 1
)

var lTMCandidateR18ExtLtmCSIReportConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LTM_Candidate_r18_Ext_Ltm_CSI_ReportConfig_r19_Release = 0
	LTM_Candidate_r18_Ext_Ltm_CSI_ReportConfig_r19_Setup   = 1
)

type LTM_Candidate_r18 struct {
	Ltm_CandidateId_r18           LTM_CandidateId_r18
	Ltm_CandidatePCI_r18          *PhysCellId
	Ltm_SSB_Config_r18            *LTM_SSB_Config_r18
	Ltm_CandidateConfig_r18       []byte
	Ltm_ConfigComplete_r18        *int64
	Ltm_EarlyUL_SyncConfig_r18    []byte
	Ltm_EarlyUL_SyncConfigSUL_r18 []byte
	Ltm_TCI_Info_r18              *LTM_TCI_Info_r18
	Ltm_NoResetID_r18             *int64
	Ltm_UE_MeasuredTA_ID_r18      *LTM_UE_MeasuredTA_ID_r18
	Ltm_NoSecurityChangeID_r19    *LTM_NoSecurityChangeId_r19
	Ltm_ExecutionCondition_r19    *struct {
		Choice  int
		Release *struct{}
		Setup   *LTM_ExecutionConditionList_r19
	}
	Ltm_NZP_CSI_RS_ResourceToAddModList_r19     *LTM_NZP_CSI_RS_ResourceToAddModList_r19
	Ltm_NZP_CSI_RS_ResourceToReleaseList_r19    *LTM_NZP_CSI_RS_ResourceToReleaseList_r19
	Ltm_NZP_CSI_RS_ResourceSetToAddModList_r19  *LTM_NZP_CSI_RS_ResourceSetToAddModList_r19
	Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r19 *LTM_NZP_CSI_RS_ResourceSetToReleaseList_r19
	Ltm_CSI_ReportConfig_r19                    *struct {
		Choice  int
		Release *struct{}
		Setup   *LTM_CSI_ReportConfig_r18
	}
	Ltm_CSI_IM_ResourceToAddModList_r19     *LTM_CSI_IM_ResourceToAddModList_r19
	Ltm_CSI_IM_ResourceToReleaseList_r19    *LTM_CSI_IM_ResourceToReleaseList_r19
	Ltm_CSI_IM_ResourceSetToAddModList_r19  *LTM_CSI_IM_ResourceSetToAddModList_r19
	Ltm_CSI_IM_ResourceSetToReleaseList_r19 *LTM_CSI_IM_ResourceSetToReleaseList_r19
}

func (ie *LTM_Candidate_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMCandidateR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ltm_NoSecurityChangeID_r19 != nil || ie.Ltm_ExecutionCondition_r19 != nil || ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r19 != nil || ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r19 != nil || ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r19 != nil || ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r19 != nil || ie.Ltm_CSI_ReportConfig_r19 != nil || ie.Ltm_CSI_IM_ResourceToAddModList_r19 != nil || ie.Ltm_CSI_IM_ResourceToReleaseList_r19 != nil || ie.Ltm_CSI_IM_ResourceSetToAddModList_r19 != nil || ie.Ltm_CSI_IM_ResourceSetToReleaseList_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ltm_CandidatePCI_r18 != nil, ie.Ltm_SSB_Config_r18 != nil, ie.Ltm_CandidateConfig_r18 != nil, ie.Ltm_ConfigComplete_r18 != nil, ie.Ltm_EarlyUL_SyncConfig_r18 != nil, ie.Ltm_EarlyUL_SyncConfigSUL_r18 != nil, ie.Ltm_TCI_Info_r18 != nil, ie.Ltm_NoResetID_r18 != nil, ie.Ltm_UE_MeasuredTA_ID_r18 != nil}); err != nil {
		return err
	}

	// 3. ltm-CandidateId-r18: ref
	{
		if err := ie.Ltm_CandidateId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. ltm-CandidatePCI-r18: ref
	{
		if ie.Ltm_CandidatePCI_r18 != nil {
			if err := ie.Ltm_CandidatePCI_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. ltm-SSB-Config-r18: ref
	{
		if ie.Ltm_SSB_Config_r18 != nil {
			if err := ie.Ltm_SSB_Config_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. ltm-CandidateConfig-r18: octet-string
	{
		if ie.Ltm_CandidateConfig_r18 != nil {
			if err := e.EncodeOctetString(ie.Ltm_CandidateConfig_r18, lTMCandidateR18LtmCandidateConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ltm-ConfigComplete-r18: enumerated
	{
		if ie.Ltm_ConfigComplete_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ltm_ConfigComplete_r18, lTMCandidateR18LtmConfigCompleteR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. ltm-EarlyUL-SyncConfig-r18: octet-string
	{
		if ie.Ltm_EarlyUL_SyncConfig_r18 != nil {
			if err := e.EncodeOctetString(ie.Ltm_EarlyUL_SyncConfig_r18, lTMCandidateR18LtmEarlyULSyncConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. ltm-EarlyUL-SyncConfigSUL-r18: octet-string
	{
		if ie.Ltm_EarlyUL_SyncConfigSUL_r18 != nil {
			if err := e.EncodeOctetString(ie.Ltm_EarlyUL_SyncConfigSUL_r18, lTMCandidateR18LtmEarlyULSyncConfigSULR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. ltm-TCI-Info-r18: ref
	{
		if ie.Ltm_TCI_Info_r18 != nil {
			if err := ie.Ltm_TCI_Info_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. ltm-NoResetID-r18: integer
	{
		if ie.Ltm_NoResetID_r18 != nil {
			if err := e.EncodeInteger(*ie.Ltm_NoResetID_r18, lTMCandidateR18LtmNoResetIDR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. ltm-UE-MeasuredTA-ID-r18: ref
	{
		if ie.Ltm_UE_MeasuredTA_ID_r18 != nil {
			if err := ie.Ltm_UE_MeasuredTA_ID_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ltm-NoSecurityChangeID-r19", Optional: true},
					{Name: "ltm-ExecutionCondition-r19", Optional: true},
					{Name: "ltm-NZP-CSI-RS-ResourceToAddModList-r19", Optional: true},
					{Name: "ltm-NZP-CSI-RS-ResourceToReleaseList-r19", Optional: true},
					{Name: "ltm-NZP-CSI-RS-ResourceSetToAddModList-r19", Optional: true},
					{Name: "ltm-NZP-CSI-RS-ResourceSetToReleaseList-r19", Optional: true},
					{Name: "ltm-CSI-ReportConfig-r19", Optional: true},
					{Name: "ltm-CSI-IM-ResourceToAddModList-r19", Optional: true},
					{Name: "ltm-CSI-IM-ResourceToReleaseList-r19", Optional: true},
					{Name: "ltm-CSI-IM-ResourceSetToAddModList-r19", Optional: true},
					{Name: "ltm-CSI-IM-ResourceSetToReleaseList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_NoSecurityChangeID_r19 != nil, ie.Ltm_ExecutionCondition_r19 != nil, ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r19 != nil, ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r19 != nil, ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r19 != nil, ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r19 != nil, ie.Ltm_CSI_ReportConfig_r19 != nil, ie.Ltm_CSI_IM_ResourceToAddModList_r19 != nil, ie.Ltm_CSI_IM_ResourceToReleaseList_r19 != nil, ie.Ltm_CSI_IM_ResourceSetToAddModList_r19 != nil, ie.Ltm_CSI_IM_ResourceSetToReleaseList_r19 != nil}); err != nil {
				return err
			}

			if ie.Ltm_NoSecurityChangeID_r19 != nil {
				if err := ie.Ltm_NoSecurityChangeID_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_ExecutionCondition_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(lTMCandidateR18ExtLtmExecutionConditionR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ltm_ExecutionCondition_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ltm_ExecutionCondition_r19).Choice {
				case LTM_Candidate_r18_Ext_Ltm_ExecutionCondition_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case LTM_Candidate_r18_Ext_Ltm_ExecutionCondition_r19_Setup:
					if err := (*ie.Ltm_ExecutionCondition_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r19 != nil {
				if err := ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r19 != nil {
				if err := ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r19 != nil {
				if err := ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r19 != nil {
				if err := ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_CSI_ReportConfig_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(lTMCandidateR18ExtLtmCSIReportConfigR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ltm_CSI_ReportConfig_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ltm_CSI_ReportConfig_r19).Choice {
				case LTM_Candidate_r18_Ext_Ltm_CSI_ReportConfig_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case LTM_Candidate_r18_Ext_Ltm_CSI_ReportConfig_r19_Setup:
					if err := (*ie.Ltm_CSI_ReportConfig_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ltm_CSI_IM_ResourceToAddModList_r19 != nil {
				if err := ie.Ltm_CSI_IM_ResourceToAddModList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_CSI_IM_ResourceToReleaseList_r19 != nil {
				if err := ie.Ltm_CSI_IM_ResourceToReleaseList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_CSI_IM_ResourceSetToAddModList_r19 != nil {
				if err := ie.Ltm_CSI_IM_ResourceSetToAddModList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_CSI_IM_ResourceSetToReleaseList_r19 != nil {
				if err := ie.Ltm_CSI_IM_ResourceSetToReleaseList_r19.Encode(ex); err != nil {
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

func (ie *LTM_Candidate_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMCandidateR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-CandidateId-r18: ref
	{
		if err := ie.Ltm_CandidateId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. ltm-CandidatePCI-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ltm_CandidatePCI_r18 = new(PhysCellId)
			if err := ie.Ltm_CandidatePCI_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. ltm-SSB-Config-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ltm_SSB_Config_r18 = new(LTM_SSB_Config_r18)
			if err := ie.Ltm_SSB_Config_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. ltm-CandidateConfig-r18: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(lTMCandidateR18LtmCandidateConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_CandidateConfig_r18 = v
		}
	}

	// 7. ltm-ConfigComplete-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(lTMCandidateR18LtmConfigCompleteR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_ConfigComplete_r18 = &idx
		}
	}

	// 8. ltm-EarlyUL-SyncConfig-r18: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(lTMCandidateR18LtmEarlyULSyncConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_EarlyUL_SyncConfig_r18 = v
		}
	}

	// 9. ltm-EarlyUL-SyncConfigSUL-r18: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(lTMCandidateR18LtmEarlyULSyncConfigSULR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_EarlyUL_SyncConfigSUL_r18 = v
		}
	}

	// 10. ltm-TCI-Info-r18: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Ltm_TCI_Info_r18 = new(LTM_TCI_Info_r18)
			if err := ie.Ltm_TCI_Info_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. ltm-NoResetID-r18: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(lTMCandidateR18LtmNoResetIDR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_NoResetID_r18 = &v
		}
	}

	// 12. ltm-UE-MeasuredTA-ID-r18: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Ltm_UE_MeasuredTA_ID_r18 = new(LTM_UE_MeasuredTA_ID_r18)
			if err := ie.Ltm_UE_MeasuredTA_ID_r18.Decode(d); err != nil {
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
				{Name: "ltm-NoSecurityChangeID-r19", Optional: true},
				{Name: "ltm-ExecutionCondition-r19", Optional: true},
				{Name: "ltm-NZP-CSI-RS-ResourceToAddModList-r19", Optional: true},
				{Name: "ltm-NZP-CSI-RS-ResourceToReleaseList-r19", Optional: true},
				{Name: "ltm-NZP-CSI-RS-ResourceSetToAddModList-r19", Optional: true},
				{Name: "ltm-NZP-CSI-RS-ResourceSetToReleaseList-r19", Optional: true},
				{Name: "ltm-CSI-ReportConfig-r19", Optional: true},
				{Name: "ltm-CSI-IM-ResourceToAddModList-r19", Optional: true},
				{Name: "ltm-CSI-IM-ResourceToReleaseList-r19", Optional: true},
				{Name: "ltm-CSI-IM-ResourceSetToAddModList-r19", Optional: true},
				{Name: "ltm-CSI-IM-ResourceSetToReleaseList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_NoSecurityChangeID_r19 = new(LTM_NoSecurityChangeId_r19)
			if err := ie.Ltm_NoSecurityChangeID_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ltm_ExecutionCondition_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTM_ExecutionConditionList_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(lTMCandidateR18ExtLtmExecutionConditionR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ltm_ExecutionCondition_r19).Choice = int(index)
			switch index {
			case LTM_Candidate_r18_Ext_Ltm_ExecutionCondition_r19_Release:
				(*ie.Ltm_ExecutionCondition_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case LTM_Candidate_r18_Ext_Ltm_ExecutionCondition_r19_Setup:
				(*ie.Ltm_ExecutionCondition_r19).Setup = new(LTM_ExecutionConditionList_r19)
				if err := (*ie.Ltm_ExecutionCondition_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r19 = new(LTM_NZP_CSI_RS_ResourceToAddModList_r19)
			if err := ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r19 = new(LTM_NZP_CSI_RS_ResourceToReleaseList_r19)
			if err := ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r19 = new(LTM_NZP_CSI_RS_ResourceSetToAddModList_r19)
			if err := ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r19 = new(LTM_NZP_CSI_RS_ResourceSetToReleaseList_r19)
			if err := ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Ltm_CSI_ReportConfig_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTM_CSI_ReportConfig_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(lTMCandidateR18ExtLtmCSIReportConfigR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ltm_CSI_ReportConfig_r19).Choice = int(index)
			switch index {
			case LTM_Candidate_r18_Ext_Ltm_CSI_ReportConfig_r19_Release:
				(*ie.Ltm_CSI_ReportConfig_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case LTM_Candidate_r18_Ext_Ltm_CSI_ReportConfig_r19_Setup:
				(*ie.Ltm_CSI_ReportConfig_r19).Setup = new(LTM_CSI_ReportConfig_r18)
				if err := (*ie.Ltm_CSI_ReportConfig_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Ltm_CSI_IM_ResourceToAddModList_r19 = new(LTM_CSI_IM_ResourceToAddModList_r19)
			if err := ie.Ltm_CSI_IM_ResourceToAddModList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(8) {
			ie.Ltm_CSI_IM_ResourceToReleaseList_r19 = new(LTM_CSI_IM_ResourceToReleaseList_r19)
			if err := ie.Ltm_CSI_IM_ResourceToReleaseList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Ltm_CSI_IM_ResourceSetToAddModList_r19 = new(LTM_CSI_IM_ResourceSetToAddModList_r19)
			if err := ie.Ltm_CSI_IM_ResourceSetToAddModList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(10) {
			ie.Ltm_CSI_IM_ResourceSetToReleaseList_r19 = new(LTM_CSI_IM_ResourceSetToReleaseList_r19)
			if err := ie.Ltm_CSI_IM_ResourceSetToReleaseList_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
