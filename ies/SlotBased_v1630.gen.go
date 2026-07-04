// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SlotBased-v1630 (line 13362).

var slotBasedV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tciMapping-r16"},
		{Name: "sequenceOffsetForRV-r16"},
	},
}

const (
	SlotBased_v1630_TciMapping_r16_CyclicMapping     = 0
	SlotBased_v1630_TciMapping_r16_SequentialMapping = 1
)

var slotBasedV1630TciMappingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SlotBased_v1630_TciMapping_r16_CyclicMapping, SlotBased_v1630_TciMapping_r16_SequentialMapping},
}

var slotBasedV1630SequenceOffsetForRVR16Constraints = per.Constrained(0, 0)

type SlotBased_v1630 struct {
	TciMapping_r16          int64
	SequenceOffsetForRV_r16 int64
}

func (ie *SlotBased_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(slotBasedV1630Constraints)
	_ = seq

	// 1. tciMapping-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.TciMapping_r16, slotBasedV1630TciMappingR16Constraints); err != nil {
			return err
		}
	}

	// 2. sequenceOffsetForRV-r16: integer
	{
		if err := e.EncodeInteger(ie.SequenceOffsetForRV_r16, slotBasedV1630SequenceOffsetForRVR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SlotBased_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(slotBasedV1630Constraints)
	_ = seq

	// 1. tciMapping-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(slotBasedV1630TciMappingR16Constraints)
		if err != nil {
			return err
		}
		ie.TciMapping_r16 = v0
	}

	// 2. sequenceOffsetForRV-r16: integer
	{
		v1, err := d.DecodeInteger(slotBasedV1630SequenceOffsetForRVR16Constraints)
		if err != nil {
			return err
		}
		ie.SequenceOffsetForRV_r16 = v1
	}

	return nil
}
