// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PLMN-Identity (line 11885).

var pLMNIdentityConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mcc", Optional: true},
		{Name: "mnc"},
	},
}

type PLMN_Identity struct {
	Mcc *MCC
	Mnc MNC
}

func (ie *PLMN_Identity) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pLMNIdentityConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mcc != nil}); err != nil {
		return err
	}

	// 2. mcc: ref
	{
		if ie.Mcc != nil {
			if err := ie.Mcc.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mnc: ref
	{
		if err := ie.Mnc.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PLMN_Identity) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pLMNIdentityConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mcc: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mcc = new(MCC)
			if err := ie.Mcc.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mnc: ref
	{
		if err := ie.Mnc.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
