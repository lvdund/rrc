// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SlotBased-r16 (line 13357).

var slotBasedR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tciMapping-r16"},
		{Name: "sequenceOffsetForRV-r16"},
	},
}

const (
	SlotBased_r16_TciMapping_r16_CyclicMapping     = 0
	SlotBased_r16_TciMapping_r16_SequentialMapping = 1
)

var slotBasedR16TciMappingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SlotBased_r16_TciMapping_r16_CyclicMapping, SlotBased_r16_TciMapping_r16_SequentialMapping},
}

var slotBasedR16SequenceOffsetForRVR16Constraints = per.Constrained(1, 3)

type SlotBased_r16 struct {
	TciMapping_r16          int64
	SequenceOffsetForRV_r16 int64
}

func (ie *SlotBased_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(slotBasedR16Constraints)
	_ = seq

	// 1. tciMapping-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.TciMapping_r16, slotBasedR16TciMappingR16Constraints); err != nil {
			return err
		}
	}

	// 2. sequenceOffsetForRV-r16: integer
	{
		if err := e.EncodeInteger(ie.SequenceOffsetForRV_r16, slotBasedR16SequenceOffsetForRVR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SlotBased_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(slotBasedR16Constraints)
	_ = seq

	// 1. tciMapping-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(slotBasedR16TciMappingR16Constraints)
		if err != nil {
			return err
		}
		ie.TciMapping_r16 = v0
	}

	// 2. sequenceOffsetForRV-r16: integer
	{
		v1, err := d.DecodeInteger(slotBasedR16SequenceOffsetForRVR16Constraints)
		if err != nil {
			return err
		}
		ie.SequenceOffsetForRV_r16 = v1
	}

	return nil
}
