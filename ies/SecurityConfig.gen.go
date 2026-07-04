// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SecurityConfig (line 13172).

var securityConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "securityAlgorithmConfig", Optional: true},
		{Name: "keyToUse", Optional: true},
	},
}

const (
	SecurityConfig_KeyToUse_Master    = 0
	SecurityConfig_KeyToUse_Secondary = 1
)

var securityConfigKeyToUseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SecurityConfig_KeyToUse_Master, SecurityConfig_KeyToUse_Secondary},
}

type SecurityConfig struct {
	SecurityAlgorithmConfig *SecurityAlgorithmConfig
	KeyToUse                *int64
}

func (ie *SecurityConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SecurityAlgorithmConfig != nil, ie.KeyToUse != nil}); err != nil {
		return err
	}

	// 3. securityAlgorithmConfig: ref
	{
		if ie.SecurityAlgorithmConfig != nil {
			if err := ie.SecurityAlgorithmConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. keyToUse: enumerated
	{
		if ie.KeyToUse != nil {
			if err := e.EncodeEnumerated(*ie.KeyToUse, securityConfigKeyToUseConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SecurityConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. securityAlgorithmConfig: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SecurityAlgorithmConfig = new(SecurityAlgorithmConfig)
			if err := ie.SecurityAlgorithmConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. keyToUse: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(securityConfigKeyToUseConstraints)
			if err != nil {
				return err
			}
			ie.KeyToUse = &idx
		}
	}

	return nil
}
