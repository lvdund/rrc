// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AdditionalTDDConfig-perPCI-ToAddMod-r18 (line 14827).

var additionalTDDConfigPerPCIToAddModR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalTDDConfig-Index-r18"},
		{Name: "tdd-UL-DL-ConfigurationCommon-r18"},
	},
}

type AdditionalTDDConfig_PerPCI_ToAddMod_r18 struct {
	AdditionalTDDConfig_Index_r18     AdditionalPCIIndex_r17
	Tdd_UL_DL_ConfigurationCommon_r18 TDD_UL_DL_ConfigCommon
}

func (ie *AdditionalTDDConfig_PerPCI_ToAddMod_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(additionalTDDConfigPerPCIToAddModR18Constraints)
	_ = seq

	// 1. additionalTDDConfig-Index-r18: ref
	{
		if err := ie.AdditionalTDDConfig_Index_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. tdd-UL-DL-ConfigurationCommon-r18: ref
	{
		if err := ie.Tdd_UL_DL_ConfigurationCommon_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AdditionalTDDConfig_PerPCI_ToAddMod_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(additionalTDDConfigPerPCIToAddModR18Constraints)
	_ = seq

	// 1. additionalTDDConfig-Index-r18: ref
	{
		if err := ie.AdditionalTDDConfig_Index_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. tdd-UL-DL-ConfigurationCommon-r18: ref
	{
		if err := ie.Tdd_UL_DL_ConfigurationCommon_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
