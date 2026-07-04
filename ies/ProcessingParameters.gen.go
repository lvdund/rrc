// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ProcessingParameters (line 23476).

var processingParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fallback"},
		{Name: "differentTB-PerSlot", Optional: true},
	},
}

const (
	ProcessingParameters_Fallback_Sc        = 0
	ProcessingParameters_Fallback_Cap1_Only = 1
)

var processingParametersFallbackConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ProcessingParameters_Fallback_Sc, ProcessingParameters_Fallback_Cap1_Only},
}

var processingParametersDifferentTBPerSlotConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "upto1", Optional: true},
		{Name: "upto2", Optional: true},
		{Name: "upto4", Optional: true},
		{Name: "upto7", Optional: true},
	},
}

type ProcessingParameters struct {
	Fallback            int64
	DifferentTB_PerSlot *struct {
		Upto1 *NumberOfCarriers
		Upto2 *NumberOfCarriers
		Upto4 *NumberOfCarriers
		Upto7 *NumberOfCarriers
	}
}

func (ie *ProcessingParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(processingParametersConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DifferentTB_PerSlot != nil}); err != nil {
		return err
	}

	// 2. fallback: enumerated
	{
		if err := e.EncodeEnumerated(ie.Fallback, processingParametersFallbackConstraints); err != nil {
			return err
		}
	}

	// 3. differentTB-PerSlot: sequence
	{
		if ie.DifferentTB_PerSlot != nil {
			c := ie.DifferentTB_PerSlot
			processingParametersDifferentTBPerSlotSeq := e.NewSequenceEncoder(processingParametersDifferentTBPerSlotConstraints)
			if err := processingParametersDifferentTBPerSlotSeq.EncodePreamble([]bool{c.Upto1 != nil, c.Upto2 != nil, c.Upto4 != nil, c.Upto7 != nil}); err != nil {
				return err
			}
			if c.Upto1 != nil {
				if err := c.Upto1.Encode(e); err != nil {
					return err
				}
			}
			if c.Upto2 != nil {
				if err := c.Upto2.Encode(e); err != nil {
					return err
				}
			}
			if c.Upto4 != nil {
				if err := c.Upto4.Encode(e); err != nil {
					return err
				}
			}
			if c.Upto7 != nil {
				if err := c.Upto7.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *ProcessingParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(processingParametersConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fallback: enumerated
	{
		v0, err := d.DecodeEnumerated(processingParametersFallbackConstraints)
		if err != nil {
			return err
		}
		ie.Fallback = v0
	}

	// 3. differentTB-PerSlot: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.DifferentTB_PerSlot = &struct {
				Upto1 *NumberOfCarriers
				Upto2 *NumberOfCarriers
				Upto4 *NumberOfCarriers
				Upto7 *NumberOfCarriers
			}{}
			c := ie.DifferentTB_PerSlot
			processingParametersDifferentTBPerSlotSeq := d.NewSequenceDecoder(processingParametersDifferentTBPerSlotConstraints)
			if err := processingParametersDifferentTBPerSlotSeq.DecodePreamble(); err != nil {
				return err
			}
			if processingParametersDifferentTBPerSlotSeq.IsComponentPresent(0) {
				c.Upto1 = new(NumberOfCarriers)
				if err := (*c.Upto1).Decode(d); err != nil {
					return err
				}
			}
			if processingParametersDifferentTBPerSlotSeq.IsComponentPresent(1) {
				c.Upto2 = new(NumberOfCarriers)
				if err := (*c.Upto2).Decode(d); err != nil {
					return err
				}
			}
			if processingParametersDifferentTBPerSlotSeq.IsComponentPresent(2) {
				c.Upto4 = new(NumberOfCarriers)
				if err := (*c.Upto4).Decode(d); err != nil {
					return err
				}
			}
			if processingParametersDifferentTBPerSlotSeq.IsComponentPresent(3) {
				c.Upto7 = new(NumberOfCarriers)
				if err := (*c.Upto7).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
