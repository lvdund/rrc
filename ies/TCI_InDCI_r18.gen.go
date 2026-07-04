// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TCI-InDCI-r18 (line 5305).

var tCIInDCIR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tci-SelectionPresentInDCI-r18", Optional: true},
		{Name: "applyIndicatedTCI-StateDCI-1-0-r18", Optional: true},
	},
}

const (
	TCI_InDCI_r18_Tci_SelectionPresentInDCI_r18_Enabled = 0
)

var tCIInDCIR18TciSelectionPresentInDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TCI_InDCI_r18_Tci_SelectionPresentInDCI_r18_Enabled},
}

const (
	TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_First  = 0
	TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_Second = 1
	TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_Both   = 2
	TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_Spare1 = 3
)

var tCIInDCIR18ApplyIndicatedTCIStateDCI10R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_First, TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_Second, TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_Both, TCI_InDCI_r18_ApplyIndicatedTCI_StateDCI_1_0_r18_Spare1},
}

type TCI_InDCI_r18 struct {
	Tci_SelectionPresentInDCI_r18      *int64
	ApplyIndicatedTCI_StateDCI_1_0_r18 *int64
}

func (ie *TCI_InDCI_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tCIInDCIR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tci_SelectionPresentInDCI_r18 != nil, ie.ApplyIndicatedTCI_StateDCI_1_0_r18 != nil}); err != nil {
		return err
	}

	// 2. tci-SelectionPresentInDCI-r18: enumerated
	{
		if ie.Tci_SelectionPresentInDCI_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Tci_SelectionPresentInDCI_r18, tCIInDCIR18TciSelectionPresentInDCIR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. applyIndicatedTCI-StateDCI-1-0-r18: enumerated
	{
		if ie.ApplyIndicatedTCI_StateDCI_1_0_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ApplyIndicatedTCI_StateDCI_1_0_r18, tCIInDCIR18ApplyIndicatedTCIStateDCI10R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TCI_InDCI_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tCIInDCIR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. tci-SelectionPresentInDCI-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(tCIInDCIR18TciSelectionPresentInDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.Tci_SelectionPresentInDCI_r18 = &idx
		}
	}

	// 3. applyIndicatedTCI-StateDCI-1-0-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(tCIInDCIR18ApplyIndicatedTCIStateDCI10R18Constraints)
			if err != nil {
				return err
			}
			ie.ApplyIndicatedTCI_StateDCI_1_0_r18 = &idx
		}
	}

	return nil
}
