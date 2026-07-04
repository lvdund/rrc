// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-CapabilityRAT-Container (line 25463).

var uECapabilityRATContainerConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rat-Type"},
		{Name: "ue-CapabilityRAT-Container"},
	},
}

var uECapabilityRATContainerUeCapabilityRATContainerConstraints = per.SizeConstraints{}

type UE_CapabilityRAT_Container struct {
	Rat_Type                   RAT_Type
	Ue_CapabilityRAT_Container []byte
}

func (ie *UE_CapabilityRAT_Container) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityRATContainerConstraints)
	_ = seq

	// 1. rat-Type: ref
	{
		if err := ie.Rat_Type.Encode(e); err != nil {
			return err
		}
	}

	// 2. ue-CapabilityRAT-Container: octet-string
	{
		if err := e.EncodeOctetString(ie.Ue_CapabilityRAT_Container, uECapabilityRATContainerUeCapabilityRATContainerConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UE_CapabilityRAT_Container) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityRATContainerConstraints)
	_ = seq

	// 1. rat-Type: ref
	{
		if err := ie.Rat_Type.Decode(d); err != nil {
			return err
		}
	}

	// 2. ue-CapabilityRAT-Container: octet-string
	{
		v1, err := d.DecodeOctetString(uECapabilityRATContainerUeCapabilityRATContainerConstraints)
		if err != nil {
			return err
		}
		ie.Ue_CapabilityRAT_Container = v1
	}

	return nil
}
