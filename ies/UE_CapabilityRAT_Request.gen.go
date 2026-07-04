// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-CapabilityRAT-Request (line 25473).

var uECapabilityRATRequestConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rat-Type"},
		{Name: "capabilityRequestFilter", Optional: true},
	},
}

var uECapabilityRATRequestCapabilityRequestFilterConstraints = per.SizeConstraints{}

type UE_CapabilityRAT_Request struct {
	Rat_Type                RAT_Type
	CapabilityRequestFilter []byte
}

func (ie *UE_CapabilityRAT_Request) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityRATRequestConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CapabilityRequestFilter != nil}); err != nil {
		return err
	}

	// 3. rat-Type: ref
	{
		if err := ie.Rat_Type.Encode(e); err != nil {
			return err
		}
	}

	// 4. capabilityRequestFilter: octet-string
	{
		if ie.CapabilityRequestFilter != nil {
			if err := e.EncodeOctetString(ie.CapabilityRequestFilter, uECapabilityRATRequestCapabilityRequestFilterConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_CapabilityRAT_Request) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityRATRequestConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rat-Type: ref
	{
		if err := ie.Rat_Type.Decode(d); err != nil {
			return err
		}
	}

	// 4. capabilityRequestFilter: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uECapabilityRATRequestCapabilityRequestFilterConstraints)
			if err != nil {
				return err
			}
			ie.CapabilityRequestFilter = v
		}
	}

	return nil
}
