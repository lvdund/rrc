// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-Parameters-v1700 (line 20953).

var mACParametersV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mac-ParametersFR2-2-r17", Optional: true},
	},
}

type MAC_Parameters_v1700 struct {
	Mac_ParametersFR2_2_r17 *MAC_ParametersFR2_2_r17
}

func (ie *MAC_Parameters_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mac_ParametersFR2_2_r17 != nil}); err != nil {
		return err
	}

	// 2. mac-ParametersFR2-2-r17: ref
	{
		if ie.Mac_ParametersFR2_2_r17 != nil {
			if err := ie.Mac_ParametersFR2_2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MAC_Parameters_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mac-ParametersFR2-2-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mac_ParametersFR2_2_r17 = new(MAC_ParametersFR2_2_r17)
			if err := ie.Mac_ParametersFR2_2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
