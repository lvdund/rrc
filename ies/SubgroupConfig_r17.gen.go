// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SubgroupConfig-r17 (line 7935).

var subgroupConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "subgroupsNumPerPO-r17"},
		{Name: "subgroupsNumForUEID-r17", Optional: true},
	},
}

var subgroupConfigR17SubgroupsNumPerPOR17Constraints = per.Constrained(1, common.MaxNrofPagingSubgroups_r17)

var subgroupConfigR17SubgroupsNumForUEIDR17Constraints = per.Constrained(1, common.MaxNrofPagingSubgroups_r17)

type SubgroupConfig_r17 struct {
	SubgroupsNumPerPO_r17   int64
	SubgroupsNumForUEID_r17 *int64
}

func (ie *SubgroupConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(subgroupConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SubgroupsNumForUEID_r17 != nil}); err != nil {
		return err
	}

	// 3. subgroupsNumPerPO-r17: integer
	{
		if err := e.EncodeInteger(ie.SubgroupsNumPerPO_r17, subgroupConfigR17SubgroupsNumPerPOR17Constraints); err != nil {
			return err
		}
	}

	// 4. subgroupsNumForUEID-r17: integer
	{
		if ie.SubgroupsNumForUEID_r17 != nil {
			if err := e.EncodeInteger(*ie.SubgroupsNumForUEID_r17, subgroupConfigR17SubgroupsNumForUEIDR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SubgroupConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(subgroupConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. subgroupsNumPerPO-r17: integer
	{
		v0, err := d.DecodeInteger(subgroupConfigR17SubgroupsNumPerPOR17Constraints)
		if err != nil {
			return err
		}
		ie.SubgroupsNumPerPO_r17 = v0
	}

	// 4. subgroupsNumForUEID-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(subgroupConfigR17SubgroupsNumForUEIDR17Constraints)
			if err != nil {
				return err
			}
			ie.SubgroupsNumForUEID_r17 = &v
		}
	}

	return nil
}
