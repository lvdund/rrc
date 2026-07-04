// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-Parameters-v1610 (line 20949).

var mACParametersV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mac-ParametersFRX-Diff-r16", Optional: true},
	},
}

type MAC_Parameters_v1610 struct {
	Mac_ParametersFRX_Diff_r16 *MAC_ParametersFRX_Diff_r16
}

func (ie *MAC_Parameters_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_ParametersFRX_Diff_r16 != nil}); err != nil {
		return err
	}

	// 2. mac-ParametersFRX-Diff-r16: ref
	{
		if ie.Mac_ParametersFRX_Diff_r16 != nil {
			if err := ie.Mac_ParametersFRX_Diff_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MAC_Parameters_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mac-ParametersFRX-Diff-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mac_ParametersFRX_Diff_r16 = new(MAC_ParametersFRX_Diff_r16)
			if err := ie.Mac_ParametersFRX_Diff_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
