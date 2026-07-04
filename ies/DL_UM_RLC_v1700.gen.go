// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-UM-RLC-v1700 (line 14173).

var dLUMRLCV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "t-ReassemblyExt-r17", Optional: true},
	},
}

type DL_UM_RLC_v1700 struct {
	T_ReassemblyExt_r17 *T_ReassemblyExt_r17
}

func (ie *DL_UM_RLC_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLUMRLCV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T_ReassemblyExt_r17 != nil}); err != nil {
		return err
	}

	// 2. t-ReassemblyExt-r17: ref
	{
		if ie.T_ReassemblyExt_r17 != nil {
			if err := ie.T_ReassemblyExt_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DL_UM_RLC_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLUMRLCV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. t-ReassemblyExt-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.T_ReassemblyExt_r17 = new(T_ReassemblyExt_r17)
			if err := ie.T_ReassemblyExt_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
