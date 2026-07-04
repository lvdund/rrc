// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SlotFormatCombination (line 15151).

var slotFormatCombinationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "slotFormatCombinationId"},
		{Name: "slotFormats"},
	},
}

var slotFormatCombinationSlotFormatsConstraints = per.SizeRange(1, common.MaxNrofSlotFormatsPerCombination)

type SlotFormatCombination struct {
	SlotFormatCombinationId SlotFormatCombinationId
	SlotFormats             []int64
}

func (ie *SlotFormatCombination) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(slotFormatCombinationConstraints)
	_ = seq

	// 1. slotFormatCombinationId: ref
	{
		if err := ie.SlotFormatCombinationId.Encode(e); err != nil {
			return err
		}
	}

	// 2. slotFormats: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(slotFormatCombinationSlotFormatsConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.SlotFormats))); err != nil {
			return err
		}
		for i := range ie.SlotFormats {
			if err := e.EncodeInteger(ie.SlotFormats[i], per.Constrained(0, 255)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SlotFormatCombination) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(slotFormatCombinationConstraints)
	_ = seq

	// 1. slotFormatCombinationId: ref
	{
		if err := ie.SlotFormatCombinationId.Decode(d); err != nil {
			return err
		}
	}

	// 2. slotFormats: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(slotFormatCombinationSlotFormatsConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SlotFormats = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			ie.SlotFormats[i] = v
		}
	}

	return nil
}
