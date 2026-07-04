// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SecurityAlgorithmConfig (line 14554).

var securityAlgorithmConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cipheringAlgorithm"},
		{Name: "integrityProtAlgorithm", Optional: true},
	},
}

type SecurityAlgorithmConfig struct {
	CipheringAlgorithm     CipheringAlgorithm
	IntegrityProtAlgorithm *IntegrityProtAlgorithm
}

func (ie *SecurityAlgorithmConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityAlgorithmConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntegrityProtAlgorithm != nil}); err != nil {
		return err
	}

	// 3. cipheringAlgorithm: ref
	{
		if err := ie.CipheringAlgorithm.Encode(e); err != nil {
			return err
		}
	}

	// 4. integrityProtAlgorithm: ref
	{
		if ie.IntegrityProtAlgorithm != nil {
			if err := ie.IntegrityProtAlgorithm.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SecurityAlgorithmConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityAlgorithmConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cipheringAlgorithm: ref
	{
		if err := ie.CipheringAlgorithm.Decode(d); err != nil {
			return err
		}
	}

	// 4. integrityProtAlgorithm: ref
	{
		if seq.IsComponentPresent(1) {
			ie.IntegrityProtAlgorithm = new(IntegrityProtAlgorithm)
			if err := ie.IntegrityProtAlgorithm.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
