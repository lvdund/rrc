// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PRS-ResourcePoolConfig-r18 (line 26868).

var sLPRSResourcePoolConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-ResourcePoolID-r18"},
		{Name: "sl-PRS-ResourcePool-r18", Optional: true},
	},
}

type SL_PRS_ResourcePoolConfig_r18 struct {
	Sl_PRS_ResourcePoolID_r18 SL_PRS_ResourcePoolID_r18
	Sl_PRS_ResourcePool_r18   *SL_PRS_ResourcePool_r18
}

func (ie *SL_PRS_ResourcePoolConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPRSResourcePoolConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_ResourcePool_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-ResourcePoolID-r18: ref
	{
		if err := ie.Sl_PRS_ResourcePoolID_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-PRS-ResourcePool-r18: ref
	{
		if ie.Sl_PRS_ResourcePool_r18 != nil {
			if err := ie.Sl_PRS_ResourcePool_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PRS_ResourcePoolConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPRSResourcePoolConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-ResourcePoolID-r18: ref
	{
		if err := ie.Sl_PRS_ResourcePoolID_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-PRS-ResourcePool-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_PRS_ResourcePool_r18 = new(SL_PRS_ResourcePool_r18)
			if err := ie.Sl_PRS_ResourcePool_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
