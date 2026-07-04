// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ConfiguredGrantConfigIndexMAC-r16 (line 6747).
// ConfiguredGrantConfigIndexMAC-r16 ::= INTEGER (0.. maxNrofConfiguredGrantConfigMAC-1-r16)

var configuredGrantConfigIndexMACR16Constraints = per.Constrained(0, common.MaxNrofConfiguredGrantConfigMAC_1_r16)

type ConfiguredGrantConfigIndexMAC_r16 struct {
	Value int64
}

func (v *ConfiguredGrantConfigIndexMAC_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, configuredGrantConfigIndexMACR16Constraints)
}

func (v *ConfiguredGrantConfigIndexMAC_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(configuredGrantConfigIndexMACR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
