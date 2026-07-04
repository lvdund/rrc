// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-MRDC-CapabilityAddFRX-Mode (line 25621).

var uEMRDCCapabilityAddFRXModeConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersMRDC-FRX-Diff"},
	},
}

type UE_MRDC_CapabilityAddFRX_Mode struct {
	MeasAndMobParametersMRDC_FRX_Diff MeasAndMobParametersMRDC_FRX_Diff
}

func (ie *UE_MRDC_CapabilityAddFRX_Mode) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEMRDCCapabilityAddFRXModeConstraints)
	_ = seq

	// 1. measAndMobParametersMRDC-FRX-Diff: ref
	{
		if err := ie.MeasAndMobParametersMRDC_FRX_Diff.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UE_MRDC_CapabilityAddFRX_Mode) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEMRDCCapabilityAddFRXModeConstraints)
	_ = seq

	// 1. measAndMobParametersMRDC-FRX-Diff: ref
	{
		if err := ie.MeasAndMobParametersMRDC_FRX_Diff.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
