// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-AM-RLC (line 14074).

var dLAMRLCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sn-FieldLength", Optional: true},
		{Name: "t-Reassembly"},
		{Name: "t-StatusProhibit"},
	},
}

type DL_AM_RLC struct {
	Sn_FieldLength   *SN_FieldLengthAM
	T_Reassembly     T_Reassembly
	T_StatusProhibit T_StatusProhibit
}

func (ie *DL_AM_RLC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLAMRLCConstraints)

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

	// 4. t-StatusProhibit: ref
	{
		if err := ie.T_StatusProhibit.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DL_AM_RLC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLAMRLCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sn-FieldLength: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sn_FieldLength = new(SN_FieldLengthAM)
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

	// 4. t-StatusProhibit: ref
	{
		if err := ie.T_StatusProhibit.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
