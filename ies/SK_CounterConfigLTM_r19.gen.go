// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SK-CounterConfigLTM-r19 (line 15129).

var sKCounterConfigLTMR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-NoSecurityChangeID-r19"},
		{Name: "sk-CounterList-r19"},
	},
}

var sKCounterConfigLTMR19SkCounterListR19Constraints = per.SizeRange(1, common.MaxSK_Counter_r18)

type SK_CounterConfigLTM_r19 struct {
	Ltm_NoSecurityChangeID_r19 LTM_NoSecurityChangeId_r19
	Sk_CounterList_r19         []SK_Counter
}

func (ie *SK_CounterConfigLTM_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sKCounterConfigLTMR19Constraints)
	_ = seq

	// 1. ltm-NoSecurityChangeID-r19: ref
	{
		if err := ie.Ltm_NoSecurityChangeID_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. sk-CounterList-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sKCounterConfigLTMR19SkCounterListR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sk_CounterList_r19))); err != nil {
			return err
		}
		for i := range ie.Sk_CounterList_r19 {
			if err := ie.Sk_CounterList_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SK_CounterConfigLTM_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sKCounterConfigLTMR19Constraints)
	_ = seq

	// 1. ltm-NoSecurityChangeID-r19: ref
	{
		if err := ie.Ltm_NoSecurityChangeID_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. sk-CounterList-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sKCounterConfigLTMR19SkCounterListR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sk_CounterList_r19 = make([]SK_Counter, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sk_CounterList_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
