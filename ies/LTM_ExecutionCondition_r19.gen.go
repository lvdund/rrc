// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-ExecutionCondition-r19 (line 8953).

var lTMExecutionConditionR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-CandidateId-r19"},
		{Name: "executionCondition-r19", Optional: true},
	},
}

var lTM_ExecutionCondition_r19ExecutionConditionR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "l1-Conditions-r19"},
		{Name: "l3-Conditions-r19"},
	},
}

const (
	LTM_ExecutionCondition_r19_ExecutionCondition_r19_L1_Conditions_r19 = 0
	LTM_ExecutionCondition_r19_ExecutionCondition_r19_L3_Conditions_r19 = 1
)

var lTMExecutionConditionR19ExecutionConditionR19L3ConditionsR19Constraints = per.SizeRange(1, 2)

type LTM_ExecutionCondition_r19 struct {
	Ltm_CandidateId_r19    LTM_CandidateId_r18
	ExecutionCondition_r19 *struct {
		Choice            int
		L1_Conditions_r19 *LTM_CSI_ReportConfigId_r18
		L3_Conditions_r19 *[]MeasId
	}
}

func (ie *LTM_ExecutionCondition_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMExecutionConditionR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ExecutionCondition_r19 != nil}); err != nil {
		return err
	}

	// 3. ltm-CandidateId-r19: ref
	{
		if err := ie.Ltm_CandidateId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. executionCondition-r19: choice
	{
		if ie.ExecutionCondition_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lTM_ExecutionCondition_r19ExecutionConditionR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ExecutionCondition_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ExecutionCondition_r19).Choice {
			case LTM_ExecutionCondition_r19_ExecutionCondition_r19_L1_Conditions_r19:
				if err := (*ie.ExecutionCondition_r19).L1_Conditions_r19.Encode(e); err != nil {
					return err
				}
			case LTM_ExecutionCondition_r19_ExecutionCondition_r19_L3_Conditions_r19:
				seqOf := e.NewSequenceOfEncoder(lTMExecutionConditionR19ExecutionConditionR19L3ConditionsR19Constraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.ExecutionCondition_r19).L3_Conditions_r19)))); err != nil {
					return err
				}
				for i := range *(*ie.ExecutionCondition_r19).L3_Conditions_r19 {
					if err := (*(*ie.ExecutionCondition_r19).L3_Conditions_r19)[i].Encode(e); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ExecutionCondition_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *LTM_ExecutionCondition_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMExecutionConditionR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-CandidateId-r19: ref
	{
		if err := ie.Ltm_CandidateId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. executionCondition-r19: choice
	{
		if seq.IsComponentPresent(1) {
			ie.ExecutionCondition_r19 = &struct {
				Choice            int
				L1_Conditions_r19 *LTM_CSI_ReportConfigId_r18
				L3_Conditions_r19 *[]MeasId
			}{}
			choiceDec := d.NewChoiceDecoder(lTM_ExecutionCondition_r19ExecutionConditionR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ExecutionCondition_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LTM_ExecutionCondition_r19_ExecutionCondition_r19_L1_Conditions_r19:
				(*ie.ExecutionCondition_r19).L1_Conditions_r19 = new(LTM_CSI_ReportConfigId_r18)
				if err := (*ie.ExecutionCondition_r19).L1_Conditions_r19.Decode(d); err != nil {
					return err
				}
			case LTM_ExecutionCondition_r19_ExecutionCondition_r19_L3_Conditions_r19:
				(*ie.ExecutionCondition_r19).L3_Conditions_r19 = new([]MeasId)
				seqOf := d.NewSequenceOfDecoder(lTMExecutionConditionR19ExecutionConditionR19L3ConditionsR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.ExecutionCondition_r19).L3_Conditions_r19) = make([]MeasId, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.ExecutionCondition_r19).L3_Conditions_r19)[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
