// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SearchSpaceSwitchConfig-r16 (line 11023).

var searchSpaceSwitchConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellGroupsForSwitchList-r16", Optional: true},
		{Name: "searchSpaceSwitchDelay-r16", Optional: true},
	},
}

var searchSpaceSwitchConfigR16CellGroupsForSwitchListR16Constraints = per.SizeRange(1, 4)

var searchSpaceSwitchConfigR16SearchSpaceSwitchDelayR16Constraints = per.Constrained(10, 52)

type SearchSpaceSwitchConfig_r16 struct {
	CellGroupsForSwitchList_r16 []CellGroupForSwitch_r16
	SearchSpaceSwitchDelay_r16  *int64
}

func (ie *SearchSpaceSwitchConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceSwitchConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CellGroupsForSwitchList_r16 != nil, ie.SearchSpaceSwitchDelay_r16 != nil}); err != nil {
		return err
	}

	// 2. cellGroupsForSwitchList-r16: sequence-of
	{
		if ie.CellGroupsForSwitchList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(searchSpaceSwitchConfigR16CellGroupsForSwitchListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.CellGroupsForSwitchList_r16))); err != nil {
				return err
			}
			for i := range ie.CellGroupsForSwitchList_r16 {
				if err := ie.CellGroupsForSwitchList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. searchSpaceSwitchDelay-r16: integer
	{
		if ie.SearchSpaceSwitchDelay_r16 != nil {
			if err := e.EncodeInteger(*ie.SearchSpaceSwitchDelay_r16, searchSpaceSwitchConfigR16SearchSpaceSwitchDelayR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SearchSpaceSwitchConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceSwitchConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cellGroupsForSwitchList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(searchSpaceSwitchConfigR16CellGroupsForSwitchListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CellGroupsForSwitchList_r16 = make([]CellGroupForSwitch_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CellGroupsForSwitchList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. searchSpaceSwitchDelay-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(searchSpaceSwitchConfigR16SearchSpaceSwitchDelayR16Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSwitchDelay_r16 = &v
		}
	}

	return nil
}
