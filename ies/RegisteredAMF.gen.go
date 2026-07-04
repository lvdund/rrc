// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RegisteredAMF (line 1777).

var registeredAMFConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity", Optional: true},
		{Name: "amf-Identifier"},
	},
}

type RegisteredAMF struct {
	Plmn_Identity  *PLMN_Identity
	Amf_Identifier AMF_Identifier
}

func (ie *RegisteredAMF) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(registeredAMFConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Plmn_Identity != nil}); err != nil {
		return err
	}

	// 2. plmn-Identity: ref
	{
		if ie.Plmn_Identity != nil {
			if err := ie.Plmn_Identity.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. amf-Identifier: ref
	{
		if err := ie.Amf_Identifier.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RegisteredAMF) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(registeredAMFConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. plmn-Identity: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Plmn_Identity = new(PLMN_Identity)
			if err := ie.Plmn_Identity.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. amf-Identifier: ref
	{
		if err := ie.Amf_Identifier.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
