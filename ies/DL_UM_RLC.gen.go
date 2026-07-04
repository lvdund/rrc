// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-UM-RLC (line 14084).

var dLUMRLCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sn-FieldLength", Optional: true},
		{Name: "t-Reassembly"},
	},
}

type DL_UM_RLC struct {
	Sn_FieldLength *SN_FieldLengthUM
	T_Reassembly   T_Reassembly
}

func (ie *DL_UM_RLC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLUMRLCConstraints)

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

	// 3. t-Reassembly: ref
	{
		if err := ie.T_Reassembly.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DL_UM_RLC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLUMRLCConstraints)

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

	// 3. t-Reassembly: ref
	{
		if err := ie.T_Reassembly.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
