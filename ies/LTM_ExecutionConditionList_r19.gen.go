// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-ExecutionConditionList-r19 (line 8951).
// LTM-ExecutionConditionList-r19 ::= SEQUENCE (SIZE (1..maxNrofLTM-Configs-r18)) OF LTM-ExecutionCondition-r19

var lTMExecutionConditionListR19SizeConstraints = per.SizeRange(1, common.MaxNrofLTM_Configs_r18)

type LTM_ExecutionConditionList_r19 []LTM_ExecutionCondition_r19

func (ie *LTM_ExecutionConditionList_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lTMExecutionConditionListR19SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *LTM_ExecutionConditionList_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lTMExecutionConditionListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(LTM_ExecutionConditionList_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
