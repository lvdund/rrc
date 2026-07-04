// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ResourcePoolConfig-r16 (line 26837).

var sLResourcePoolConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ResourcePoolID-r16"},
		{Name: "sl-ResourcePool-r16", Optional: true},
	},
}

type SL_ResourcePoolConfig_r16 struct {
	Sl_ResourcePoolID_r16 SL_ResourcePoolID_r16
	Sl_ResourcePool_r16   *SL_ResourcePool_r16
}

func (ie *SL_ResourcePoolConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLResourcePoolConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ResourcePool_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-ResourcePoolID-r16: ref
	{
		if err := ie.Sl_ResourcePoolID_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-ResourcePool-r16: ref
	{
		if ie.Sl_ResourcePool_r16 != nil {
			if err := ie.Sl_ResourcePool_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_ResourcePoolConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLResourcePoolConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ResourcePoolID-r16: ref
	{
		if err := ie.Sl_ResourcePoolID_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-ResourcePool-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_ResourcePool_r16 = new(SL_ResourcePool_r16)
			if err := ie.Sl_ResourcePool_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
