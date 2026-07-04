// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TxHoppingConfig-r18 (line 15721).

var txHoppingConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "overlapValue-r18"},
		{Name: "numberOfHops-r18"},
		{Name: "slotOffsetForRemainingHopsList-r18"},
	},
}

const (
	TxHoppingConfig_r18_OverlapValue_r18_ZeroRB = 0
	TxHoppingConfig_r18_OverlapValue_r18_OneRB  = 1
	TxHoppingConfig_r18_OverlapValue_r18_TwoRB  = 2
	TxHoppingConfig_r18_OverlapValue_r18_FourRB = 3
)

var txHoppingConfigR18OverlapValueR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TxHoppingConfig_r18_OverlapValue_r18_ZeroRB, TxHoppingConfig_r18_OverlapValue_r18_OneRB, TxHoppingConfig_r18_OverlapValue_r18_TwoRB, TxHoppingConfig_r18_OverlapValue_r18_FourRB},
}

var txHoppingConfigR18NumberOfHopsR18Constraints = per.Constrained(1, 6)

var txHoppingConfigR18SlotOffsetForRemainingHopsListR18Constraints = per.SizeRange(1, common.MaxNrofHops_1_r18)

type TxHoppingConfig_r18 struct {
	OverlapValue_r18                   int64
	NumberOfHops_r18                   int64
	SlotOffsetForRemainingHopsList_r18 []SlotOffsetForRemainingHops_r18
}

func (ie *TxHoppingConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(txHoppingConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. overlapValue-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.OverlapValue_r18, txHoppingConfigR18OverlapValueR18Constraints); err != nil {
			return err
		}
	}

	// 3. numberOfHops-r18: integer
	{
		if err := e.EncodeInteger(ie.NumberOfHops_r18, txHoppingConfigR18NumberOfHopsR18Constraints); err != nil {
			return err
		}
	}

	// 4. slotOffsetForRemainingHopsList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(txHoppingConfigR18SlotOffsetForRemainingHopsListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.SlotOffsetForRemainingHopsList_r18))); err != nil {
			return err
		}
		for i := range ie.SlotOffsetForRemainingHopsList_r18 {
			if err := ie.SlotOffsetForRemainingHopsList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TxHoppingConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(txHoppingConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. overlapValue-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(txHoppingConfigR18OverlapValueR18Constraints)
		if err != nil {
			return err
		}
		ie.OverlapValue_r18 = v0
	}

	// 3. numberOfHops-r18: integer
	{
		v1, err := d.DecodeInteger(txHoppingConfigR18NumberOfHopsR18Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfHops_r18 = v1
	}

	// 4. slotOffsetForRemainingHopsList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(txHoppingConfigR18SlotOffsetForRemainingHopsListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SlotOffsetForRemainingHopsList_r18 = make([]SlotOffsetForRemainingHops_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SlotOffsetForRemainingHopsList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
