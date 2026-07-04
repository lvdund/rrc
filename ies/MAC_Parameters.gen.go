// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-Parameters (line 20944).

var mACParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mac-ParametersCommon", Optional: true},
		{Name: "mac-ParametersXDD-Diff", Optional: true},
	},
}

type MAC_Parameters struct {
	Mac_ParametersCommon   *MAC_ParametersCommon
	Mac_ParametersXDD_Diff *MAC_ParametersXDD_Diff
}

func (ie *MAC_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_ParametersCommon != nil, ie.Mac_ParametersXDD_Diff != nil}); err != nil {
		return err
	}

	// 2. mac-ParametersCommon: ref
	{
		if ie.Mac_ParametersCommon != nil {
			if err := ie.Mac_ParametersCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mac-ParametersXDD-Diff: ref
	{
		if ie.Mac_ParametersXDD_Diff != nil {
			if err := ie.Mac_ParametersXDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MAC_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mac-ParametersCommon: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mac_ParametersCommon = new(MAC_ParametersCommon)
			if err := ie.Mac_ParametersCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mac-ParametersXDD-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mac_ParametersXDD_Diff = new(MAC_ParametersXDD_Diff)
			if err := ie.Mac_ParametersXDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
