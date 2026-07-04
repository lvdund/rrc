// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-ResourceOfModeB-r19 (line 7402).

var pUSCHResourceOfModeBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "configuredGrantConfigIndex-r19"},
		{Name: "ul-BWP-Id-r19"},
	},
}

type PUSCH_ResourceOfModeB_r19 struct {
	ConfiguredGrantConfigIndex_r19 ConfiguredGrantConfigIndex_r16
	Ul_BWP_Id_r19                  BWP_Id
}

func (ie *PUSCH_ResourceOfModeB_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHResourceOfModeBR19Constraints)
	_ = seq

	// 1. configuredGrantConfigIndex-r19: ref
	{
		if err := ie.ConfiguredGrantConfigIndex_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. ul-BWP-Id-r19: ref
	{
		if err := ie.Ul_BWP_Id_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUSCH_ResourceOfModeB_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHResourceOfModeBR19Constraints)
	_ = seq

	// 1. configuredGrantConfigIndex-r19: ref
	{
		if err := ie.ConfiguredGrantConfigIndex_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. ul-BWP-Id-r19: ref
	{
		if err := ie.Ul_BWP_Id_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
