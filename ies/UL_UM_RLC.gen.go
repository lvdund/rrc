// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-UM-RLC (line 14080).

var uLUMRLCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sn-FieldLength", Optional: true},
	},
}

type UL_UM_RLC struct {
	Sn_FieldLength *SN_FieldLengthUM
}

func (ie *UL_UM_RLC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLUMRLCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sn_FieldLength != nil}); err != nil {
		return err
	}

	// 2. sn-FieldLength: ref
	{
		if ie.Sn_FieldLength != nil {
			if err := ie.Sn_FieldLength.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UL_UM_RLC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLUMRLCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sn-FieldLength: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sn_FieldLength = new(SN_FieldLengthUM)
			if err := ie.Sn_FieldLength.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
