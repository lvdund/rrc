// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SelectionWindowConfig-r16 (line 28084).

var sLSelectionWindowConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Priority-r16"},
		{Name: "sl-SelectionWindow-r16"},
	},
}

var sLSelectionWindowConfigR16SlPriorityR16Constraints = per.Constrained(1, 8)

const (
	SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N1  = 0
	SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N5  = 1
	SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N10 = 2
	SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N20 = 3
)

var sLSelectionWindowConfigR16SlSelectionWindowR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N1, SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N5, SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N10, SL_SelectionWindowConfig_r16_Sl_SelectionWindow_r16_N20},
}

type SL_SelectionWindowConfig_r16 struct {
	Sl_Priority_r16        int64
	Sl_SelectionWindow_r16 int64
}

func (ie *SL_SelectionWindowConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSelectionWindowConfigR16Constraints)
	_ = seq

	// 1. sl-Priority-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_Priority_r16, sLSelectionWindowConfigR16SlPriorityR16Constraints); err != nil {
			return err
		}
	}

	// 2. sl-SelectionWindow-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_SelectionWindow_r16, sLSelectionWindowConfigR16SlSelectionWindowR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_SelectionWindowConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSelectionWindowConfigR16Constraints)
	_ = seq

	// 1. sl-Priority-r16: integer
	{
		v0, err := d.DecodeInteger(sLSelectionWindowConfigR16SlPriorityR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Priority_r16 = v0
	}

	// 2. sl-SelectionWindow-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sLSelectionWindowConfigR16SlSelectionWindowR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_SelectionWindow_r16 = v1
	}

	return nil
}
