// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TwoPUCCH-Grp-Configurations-r17 (line 18320).

var twoPUCCHGrpConfigurationsR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "primaryPUCCH-GroupConfig-r17"},
		{Name: "secondaryPUCCH-GroupConfig-r17"},
	},
}

type TwoPUCCH_Grp_Configurations_r17 struct {
	PrimaryPUCCH_GroupConfig_r17   PUCCH_Group_Config_r17
	SecondaryPUCCH_GroupConfig_r17 PUCCH_Group_Config_r17
}

func (ie *TwoPUCCH_Grp_Configurations_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(twoPUCCHGrpConfigurationsR17Constraints)
	_ = seq

	// 1. primaryPUCCH-GroupConfig-r17: ref
	{
		if err := ie.PrimaryPUCCH_GroupConfig_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. secondaryPUCCH-GroupConfig-r17: ref
	{
		if err := ie.SecondaryPUCCH_GroupConfig_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TwoPUCCH_Grp_Configurations_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(twoPUCCHGrpConfigurationsR17Constraints)
	_ = seq

	// 1. primaryPUCCH-GroupConfig-r17: ref
	{
		if err := ie.PrimaryPUCCH_GroupConfig_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. secondaryPUCCH-GroupConfig-r17: ref
	{
		if err := ie.SecondaryPUCCH_GroupConfig_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
