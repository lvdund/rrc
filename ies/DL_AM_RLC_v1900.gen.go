// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-AM-RLC-v1900 (line 14165).

var dLAMRLCV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "t-RxDiscard-r19", Optional: true},
	},
}

type DL_AM_RLC_v1900 struct {
	T_RxDiscard_r19 *T_RxDiscard_r19
}

func (ie *DL_AM_RLC_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLAMRLCV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T_RxDiscard_r19 != nil}); err != nil {
		return err
	}

	// 2. t-RxDiscard-r19: ref
	{
		if ie.T_RxDiscard_r19 != nil {
			if err := ie.T_RxDiscard_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DL_AM_RLC_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLAMRLCV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. t-RxDiscard-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.T_RxDiscard_r19 = new(T_RxDiscard_r19)
			if err := ie.T_RxDiscard_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
