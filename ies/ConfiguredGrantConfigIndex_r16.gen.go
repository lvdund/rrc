// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConfiguredGrantConfigIndex-r16 (line 6742).
// ConfiguredGrantConfigIndex-r16 ::= INTEGER (0.. maxNrofConfiguredGrantConfig-1-r16)

var configuredGrantConfigIndexR16Constraints = per.Constrained(0, common.MaxNrofConfiguredGrantConfig_1_r16)

type ConfiguredGrantConfigIndex_r16 struct {
	Value int64
}

func (v *ConfiguredGrantConfigIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, configuredGrantConfigIndexR16Constraints)
}

func (v *ConfiguredGrantConfigIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(configuredGrantConfigIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
