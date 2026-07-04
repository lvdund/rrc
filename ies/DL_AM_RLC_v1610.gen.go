// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-AM-RLC-v1610 (line 14156).

var dLAMRLCV1610Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "t-StatusProhibit-v1610", Optional: true},
	},
}

type DL_AM_RLC_v1610 struct {
	T_StatusProhibit_v1610 *T_StatusProhibit_v1610
}

func (ie *DL_AM_RLC_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLAMRLCV1610Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T_StatusProhibit_v1610 != nil}); err != nil {
		return err
	}

	// 3. t-StatusProhibit-v1610: ref
	{
		if ie.T_StatusProhibit_v1610 != nil {
			if err := ie.T_StatusProhibit_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DL_AM_RLC_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLAMRLCV1610Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. t-StatusProhibit-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.T_StatusProhibit_v1610 = new(T_StatusProhibit_v1610)
			if err := ie.T_StatusProhibit_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
