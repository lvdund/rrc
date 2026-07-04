// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-Config-r18 (line 8751).

var lTMConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-ReferenceConfiguration-r18", Optional: true},
		{Name: "ltm-CandidateToReleaseList-r18", Optional: true},
		{Name: "ltm-CandidateToAddModList-r18", Optional: true},
		{Name: "ltm-ServingCellNoResetID-r18", Optional: true},
		{Name: "ltm-CSI-ResourceConfigToAddModList-r18", Optional: true},
		{Name: "ltm-CSI-ResourceConfigToReleaseList-r18", Optional: true},
		{Name: "attemptLTM-Switch-r18", Optional: true},
		{Name: "ltm-ServingCellUE-MeasuredTA-ID-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var lTMConfigR18LtmCandidateToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofLTM_Configs_r18)

var lTMConfigR18LtmCandidateToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofLTM_Configs_r18)

var lTMConfigR18LtmServingCellNoResetIDR18Constraints = per.Constrained(1, common.MaxNrofLTM_Configs_Plus1_r18)

const (
	LTM_Config_r18_AttemptLTM_Switch_r18_True = 0
)

var lTMConfigR18AttemptLTMSwitchR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_Config_r18_AttemptLTM_Switch_r18_True},
}

var lTMConfigR18LtmServingCellUEMeasuredTAIDR18Constraints = per.Constrained(1, common.MaxNrofLTM_Configs_Plus1_r18)

var lTMConfigR18ExtLtmServingCellExecutionConditionR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	LTM_Config_r18_Ext_Ltm_ServingCellExecutionCondition_r19_Release = 0
	LTM_Config_r18_Ext_Ltm_ServingCellExecutionCondition_r19_Setup   = 1
)

type LTM_Config_r18 struct {
	Ltm_ReferenceConfiguration_r18          *LTM_ReferenceConfiguration_r18
	Ltm_CandidateToReleaseList_r18          []LTM_CandidateId_r18
	Ltm_CandidateToAddModList_r18           []LTM_Candidate_r18
	Ltm_ServingCellNoResetID_r18            *int64
	Ltm_CSI_ResourceConfigToAddModList_r18  *LTM_CSI_ResourceConfigToAddModList_r18
	Ltm_CSI_ResourceConfigToReleaseList_r18 *LTM_CSI_ResourceConfigToReleaseList_r18
	AttemptLTM_Switch_r18                   *int64
	Ltm_ServingCellUE_MeasuredTA_ID_r18     *int64
	Ltm_ServingCellNoSecurityChangeID_r19   *LTM_NoSecurityChangeId_r19
	Ltm_ServingCellExecutionCondition_r19   *struct {
		Choice  int
		Release *struct{}
		Setup   *LTM_ExecutionConditionList_r19
	}
}

func (ie *LTM_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMConfigR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ltm_ServingCellNoSecurityChangeID_r19 != nil || ie.Ltm_ServingCellExecutionCondition_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ltm_ReferenceConfiguration_r18 != nil, ie.Ltm_CandidateToReleaseList_r18 != nil, ie.Ltm_CandidateToAddModList_r18 != nil, ie.Ltm_ServingCellNoResetID_r18 != nil, ie.Ltm_CSI_ResourceConfigToAddModList_r18 != nil, ie.Ltm_CSI_ResourceConfigToReleaseList_r18 != nil, ie.AttemptLTM_Switch_r18 != nil, ie.Ltm_ServingCellUE_MeasuredTA_ID_r18 != nil}); err != nil {
		return err
	}

	// 3. ltm-ReferenceConfiguration-r18: ref
	{
		if ie.Ltm_ReferenceConfiguration_r18 != nil {
			if err := ie.Ltm_ReferenceConfiguration_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ltm-CandidateToReleaseList-r18: sequence-of
	{
		if ie.Ltm_CandidateToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMConfigR18LtmCandidateToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_CandidateToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_CandidateToReleaseList_r18 {
				if err := ie.Ltm_CandidateToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. ltm-CandidateToAddModList-r18: sequence-of
	{
		if ie.Ltm_CandidateToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMConfigR18LtmCandidateToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_CandidateToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_CandidateToAddModList_r18 {
				if err := ie.Ltm_CandidateToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. ltm-ServingCellNoResetID-r18: integer
	{
		if ie.Ltm_ServingCellNoResetID_r18 != nil {
			if err := e.EncodeInteger(*ie.Ltm_ServingCellNoResetID_r18, lTMConfigR18LtmServingCellNoResetIDR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ltm-CSI-ResourceConfigToAddModList-r18: ref
	{
		if ie.Ltm_CSI_ResourceConfigToAddModList_r18 != nil {
			if err := ie.Ltm_CSI_ResourceConfigToAddModList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. ltm-CSI-ResourceConfigToReleaseList-r18: ref
	{
		if ie.Ltm_CSI_ResourceConfigToReleaseList_r18 != nil {
			if err := ie.Ltm_CSI_ResourceConfigToReleaseList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. attemptLTM-Switch-r18: enumerated
	{
		if ie.AttemptLTM_Switch_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AttemptLTM_Switch_r18, lTMConfigR18AttemptLTMSwitchR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. ltm-ServingCellUE-MeasuredTA-ID-r18: integer
	{
		if ie.Ltm_ServingCellUE_MeasuredTA_ID_r18 != nil {
			if err := e.EncodeInteger(*ie.Ltm_ServingCellUE_MeasuredTA_ID_r18, lTMConfigR18LtmServingCellUEMeasuredTAIDR18Constraints); err != nil {
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
					{Name: "ltm-ServingCellNoSecurityChangeID-r19", Optional: true},
					{Name: "ltm-ServingCellExecutionCondition-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_ServingCellNoSecurityChangeID_r19 != nil, ie.Ltm_ServingCellExecutionCondition_r19 != nil}); err != nil {
				return err
			}

			if ie.Ltm_ServingCellNoSecurityChangeID_r19 != nil {
				if err := ie.Ltm_ServingCellNoSecurityChangeID_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_ServingCellExecutionCondition_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(lTMConfigR18ExtLtmServingCellExecutionConditionR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ltm_ServingCellExecutionCondition_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ltm_ServingCellExecutionCondition_r19).Choice {
				case LTM_Config_r18_Ext_Ltm_ServingCellExecutionCondition_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case LTM_Config_r18_Ext_Ltm_ServingCellExecutionCondition_r19_Setup:
					if err := (*ie.Ltm_ServingCellExecutionCondition_r19).Setup.Encode(ex); err != nil {
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

func (ie *LTM_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-ReferenceConfiguration-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ltm_ReferenceConfiguration_r18 = new(LTM_ReferenceConfiguration_r18)
			if err := ie.Ltm_ReferenceConfiguration_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ltm-CandidateToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(lTMConfigR18LtmCandidateToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_CandidateToReleaseList_r18 = make([]LTM_CandidateId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_CandidateToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. ltm-CandidateToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(lTMConfigR18LtmCandidateToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_CandidateToAddModList_r18 = make([]LTM_Candidate_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_CandidateToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. ltm-ServingCellNoResetID-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(lTMConfigR18LtmServingCellNoResetIDR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_ServingCellNoResetID_r18 = &v
		}
	}

	// 7. ltm-CSI-ResourceConfigToAddModList-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ltm_CSI_ResourceConfigToAddModList_r18 = new(LTM_CSI_ResourceConfigToAddModList_r18)
			if err := ie.Ltm_CSI_ResourceConfigToAddModList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. ltm-CSI-ResourceConfigToReleaseList-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Ltm_CSI_ResourceConfigToReleaseList_r18 = new(LTM_CSI_ResourceConfigToReleaseList_r18)
			if err := ie.Ltm_CSI_ResourceConfigToReleaseList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. attemptLTM-Switch-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(lTMConfigR18AttemptLTMSwitchR18Constraints)
			if err != nil {
				return err
			}
			ie.AttemptLTM_Switch_r18 = &idx
		}
	}

	// 10. ltm-ServingCellUE-MeasuredTA-ID-r18: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(lTMConfigR18LtmServingCellUEMeasuredTAIDR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_ServingCellUE_MeasuredTA_ID_r18 = &v
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
				{Name: "ltm-ServingCellNoSecurityChangeID-r19", Optional: true},
				{Name: "ltm-ServingCellExecutionCondition-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_ServingCellNoSecurityChangeID_r19 = new(LTM_NoSecurityChangeId_r19)
			if err := ie.Ltm_ServingCellNoSecurityChangeID_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ltm_ServingCellExecutionCondition_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTM_ExecutionConditionList_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(lTMConfigR18ExtLtmServingCellExecutionConditionR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ltm_ServingCellExecutionCondition_r19).Choice = int(index)
			switch index {
			case LTM_Config_r18_Ext_Ltm_ServingCellExecutionCondition_r19_Release:
				(*ie.Ltm_ServingCellExecutionCondition_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case LTM_Config_r18_Ext_Ltm_ServingCellExecutionCondition_r19_Setup:
				(*ie.Ltm_ServingCellExecutionCondition_r19).Setup = new(LTM_ExecutionConditionList_r19)
				if err := (*ie.Ltm_ServingCellExecutionCondition_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
