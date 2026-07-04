// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: N3C-IndirectPathAddChange-r18 (line 10261).

var n3CIndirectPathAddChangeR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "n3c-RelayIdentification-r18"},
	},
}

type N3C_IndirectPathAddChange_r18 struct {
	N3c_RelayIdentification_r18 N3C_RelayUE_Info_r18
}

func (ie *N3C_IndirectPathAddChange_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(n3CIndirectPathAddChangeR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. n3c-RelayIdentification-r18: ref
	{
		if err := ie.N3c_RelayIdentification_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *N3C_IndirectPathAddChange_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(n3CIndirectPathAddChangeR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. n3c-RelayIdentification-r18: ref
	{
		if err := ie.N3c_RelayIdentification_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
