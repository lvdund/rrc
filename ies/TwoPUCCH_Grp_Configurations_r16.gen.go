// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TwoPUCCH-Grp-Configurations-r16 (line 18315).

var twoPUCCHGrpConfigurationsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-PrimaryGroupMapping-r16"},
		{Name: "pucch-SecondaryGroupMapping-r16"},
	},
}

type TwoPUCCH_Grp_Configurations_r16 struct {
	Pucch_PrimaryGroupMapping_r16   TwoPUCCH_Grp_ConfigParams_r16
	Pucch_SecondaryGroupMapping_r16 TwoPUCCH_Grp_ConfigParams_r16
}

func (ie *TwoPUCCH_Grp_Configurations_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(twoPUCCHGrpConfigurationsR16Constraints)
	_ = seq

	// 1. pucch-PrimaryGroupMapping-r16: ref
	{
		if err := ie.Pucch_PrimaryGroupMapping_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. pucch-SecondaryGroupMapping-r16: ref
	{
		if err := ie.Pucch_SecondaryGroupMapping_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TwoPUCCH_Grp_Configurations_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(twoPUCCHGrpConfigurationsR16Constraints)
	_ = seq

	// 1. pucch-PrimaryGroupMapping-r16: ref
	{
		if err := ie.Pucch_PrimaryGroupMapping_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. pucch-SecondaryGroupMapping-r16: ref
	{
		if err := ie.Pucch_SecondaryGroupMapping_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
