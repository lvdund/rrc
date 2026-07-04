// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasObjectCLI-r16 (line 9337).

var measObjectCLIR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cli-ResourceConfig-r16"},
	},
}

type MeasObjectCLI_r16 struct {
	Cli_ResourceConfig_r16 CLI_ResourceConfig_r16
}

func (ie *MeasObjectCLI_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectCLIR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. cli-ResourceConfig-r16: ref
	{
		if err := ie.Cli_ResourceConfig_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasObjectCLI_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectCLIR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. cli-ResourceConfig-r16: ref
	{
		if err := ie.Cli_ResourceConfig_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
