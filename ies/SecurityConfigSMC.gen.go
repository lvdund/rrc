// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SecurityConfigSMC (line 1940).

var securityConfigSMCConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "securityAlgorithmConfig"},
	},
}

type SecurityConfigSMC struct {
	SecurityAlgorithmConfig SecurityAlgorithmConfig
}

func (ie *SecurityConfigSMC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityConfigSMCConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. securityAlgorithmConfig: ref
	{
		if err := ie.SecurityAlgorithmConfig.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SecurityConfigSMC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityConfigSMCConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. securityAlgorithmConfig: ref
	{
		if err := ie.SecurityAlgorithmConfig.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
