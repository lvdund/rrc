// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SK-CounterConfig-r18 (line 6522).

var sKCounterConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "securityCellSetId-r18"},
		{Name: "sk-CounterList-r18"},
	},
}

var sKCounterConfigR18SkCounterListR18Constraints = per.SizeRange(1, common.MaxSK_Counter_r18)

type SK_CounterConfig_r18 struct {
	SecurityCellSetId_r18 SecurityCellSetId_r18
	Sk_CounterList_r18    []SK_Counter
}

func (ie *SK_CounterConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sKCounterConfigR18Constraints)
	_ = seq

	// 1. securityCellSetId-r18: ref
	{
		if err := ie.SecurityCellSetId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. sk-CounterList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sKCounterConfigR18SkCounterListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sk_CounterList_r18))); err != nil {
			return err
		}
		for i := range ie.Sk_CounterList_r18 {
			if err := ie.Sk_CounterList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SK_CounterConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sKCounterConfigR18Constraints)
	_ = seq

	// 1. securityCellSetId-r18: ref
	{
		if err := ie.SecurityCellSetId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. sk-CounterList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sKCounterConfigR18SkCounterListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sk_CounterList_r18 = make([]SK_Counter, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sk_CounterList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
