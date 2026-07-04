// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ValidityArea-r16 (line 9265).

var validityAreaR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "validityCellList-r16", Optional: true},
	},
}

type ValidityArea_r16 struct {
	CarrierFreq_r16      ARFCN_ValueNR
	ValidityCellList_r16 *ValidityCellList
}

func (ie *ValidityArea_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(validityAreaR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ValidityCellList_r16 != nil}); err != nil {
		return err
	}

	// 2. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. validityCellList-r16: ref
	{
		if ie.ValidityCellList_r16 != nil {
			if err := ie.ValidityCellList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ValidityArea_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(validityAreaR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. validityCellList-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ValidityCellList_r16 = new(ValidityCellList)
			if err := ie.ValidityCellList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
