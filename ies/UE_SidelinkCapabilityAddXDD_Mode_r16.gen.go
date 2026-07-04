// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-SidelinkCapabilityAddXDD-Mode-r16 (line 25070).

var uESidelinkCapabilityAddXDDModeR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mac-ParametersSidelinkXDD-Diff-r16", Optional: true},
	},
}

type UE_SidelinkCapabilityAddXDD_Mode_r16 struct {
	Mac_ParametersSidelinkXDD_Diff_r16 *MAC_ParametersSidelinkXDD_Diff_r16
}

func (ie *UE_SidelinkCapabilityAddXDD_Mode_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uESidelinkCapabilityAddXDDModeR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_ParametersSidelinkXDD_Diff_r16 != nil}); err != nil {
		return err
	}

	// 2. mac-ParametersSidelinkXDD-Diff-r16: ref
	{
		if ie.Mac_ParametersSidelinkXDD_Diff_r16 != nil {
			if err := ie.Mac_ParametersSidelinkXDD_Diff_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_SidelinkCapabilityAddXDD_Mode_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uESidelinkCapabilityAddXDDModeR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mac-ParametersSidelinkXDD-Diff-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mac_ParametersSidelinkXDD_Diff_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16)
			if err := ie.Mac_ParametersSidelinkXDD_Diff_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
