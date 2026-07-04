// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SearchSpaceSwitchConfig-r19 (line 11033).

var searchSpaceSwitchConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sr-TriggeredSearchSpaceSwitchGroupId-r19", Optional: true},
	},
}

var searchSpaceSwitchConfigR19SrTriggeredSearchSpaceSwitchGroupIdR19Constraints = per.Constrained(0, common.MaxNrofSearchSpaceGroups_1_r17)

type SearchSpaceSwitchConfig_r19 struct {
	Sr_TriggeredSearchSpaceSwitchGroupId_r19 *int64
}

func (ie *SearchSpaceSwitchConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceSwitchConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sr_TriggeredSearchSpaceSwitchGroupId_r19 != nil}); err != nil {
		return err
	}

	// 2. sr-TriggeredSearchSpaceSwitchGroupId-r19: integer
	{
		if ie.Sr_TriggeredSearchSpaceSwitchGroupId_r19 != nil {
			if err := e.EncodeInteger(*ie.Sr_TriggeredSearchSpaceSwitchGroupId_r19, searchSpaceSwitchConfigR19SrTriggeredSearchSpaceSwitchGroupIdR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SearchSpaceSwitchConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceSwitchConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sr-TriggeredSearchSpaceSwitchGroupId-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(searchSpaceSwitchConfigR19SrTriggeredSearchSpaceSwitchGroupIdR19Constraints)
			if err != nil {
				return err
			}
			ie.Sr_TriggeredSearchSpaceSwitchGroupId_r19 = &v
		}
	}

	return nil
}
