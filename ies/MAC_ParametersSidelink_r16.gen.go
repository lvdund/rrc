// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-ParametersSidelink-r16 (line 25064).

var mACParametersSidelinkR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mac-ParametersSidelinkCommon-r16", Optional: true},
		{Name: "mac-ParametersSidelinkXDD-Diff-r16", Optional: true},
	},
}

type MAC_ParametersSidelink_r16 struct {
	Mac_ParametersSidelinkCommon_r16   *MAC_ParametersSidelinkCommon_r16
	Mac_ParametersSidelinkXDD_Diff_r16 *MAC_ParametersSidelinkXDD_Diff_r16
}

func (ie *MAC_ParametersSidelink_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersSidelinkR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_ParametersSidelinkCommon_r16 != nil, ie.Mac_ParametersSidelinkXDD_Diff_r16 != nil}); err != nil {
		return err
	}

	// 3. mac-ParametersSidelinkCommon-r16: ref
	{
		if ie.Mac_ParametersSidelinkCommon_r16 != nil {
			if err := ie.Mac_ParametersSidelinkCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mac-ParametersSidelinkXDD-Diff-r16: ref
	{
		if ie.Mac_ParametersSidelinkXDD_Diff_r16 != nil {
			if err := ie.Mac_ParametersSidelinkXDD_Diff_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MAC_ParametersSidelink_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersSidelinkR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mac-ParametersSidelinkCommon-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mac_ParametersSidelinkCommon_r16 = new(MAC_ParametersSidelinkCommon_r16)
			if err := ie.Mac_ParametersSidelinkCommon_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mac-ParametersSidelinkXDD-Diff-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mac_ParametersSidelinkXDD_Diff_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16)
			if err := ie.Mac_ParametersSidelinkXDD_Diff_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
