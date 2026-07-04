// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TwoPUCCH-Grp-ConfigParams-r16 (line 18325).

var twoPUCCHGrpConfigParamsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-GroupMapping-r16"},
		{Name: "pucch-TX-r16"},
	},
}

type TwoPUCCH_Grp_ConfigParams_r16 struct {
	Pucch_GroupMapping_r16 PUCCH_Grp_CarrierTypes_r16
	Pucch_TX_r16           PUCCH_Grp_CarrierTypes_r16
}

func (ie *TwoPUCCH_Grp_ConfigParams_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(twoPUCCHGrpConfigParamsR16Constraints)
	_ = seq

	// 1. pucch-GroupMapping-r16: ref
	{
		if err := ie.Pucch_GroupMapping_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. pucch-TX-r16: ref
	{
		if err := ie.Pucch_TX_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TwoPUCCH_Grp_ConfigParams_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(twoPUCCHGrpConfigParamsR16Constraints)
	_ = seq

	// 1. pucch-GroupMapping-r16: ref
	{
		if err := ie.Pucch_GroupMapping_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. pucch-TX-r16: ref
	{
		if err := ie.Pucch_TX_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
