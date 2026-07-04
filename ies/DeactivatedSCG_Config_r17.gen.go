// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DeactivatedSCG-Config-r17 (line 5759).

var deactivatedSCGConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bfd-and-RLM-r17"},
	},
}

type DeactivatedSCG_Config_r17 struct {
	Bfd_And_RLM_r17 bool
}

func (ie *DeactivatedSCG_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(deactivatedSCGConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. bfd-and-RLM-r17: boolean
	{
		if err := e.EncodeBoolean(ie.Bfd_And_RLM_r17); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DeactivatedSCG_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(deactivatedSCGConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. bfd-and-RLM-r17: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Bfd_And_RLM_r17 = v0
	}

	return nil
}
