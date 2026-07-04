// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-CapabilityAddFRX-Mode (line 25918).

var uENRCapabilityAddFRXModeConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "phy-ParametersFRX-Diff", Optional: true},
		{Name: "measAndMobParametersFRX-Diff", Optional: true},
	},
}

type UE_NR_CapabilityAddFRX_Mode struct {
	Phy_ParametersFRX_Diff       *Phy_ParametersFRX_Diff
	MeasAndMobParametersFRX_Diff *MeasAndMobParametersFRX_Diff
}

func (ie *UE_NR_CapabilityAddFRX_Mode) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityAddFRXModeConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Phy_ParametersFRX_Diff != nil, ie.MeasAndMobParametersFRX_Diff != nil}); err != nil {
		return err
	}

	// 2. phy-ParametersFRX-Diff: ref
	{
		if ie.Phy_ParametersFRX_Diff != nil {
			if err := ie.Phy_ParametersFRX_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersFRX-Diff: ref
	{
		if ie.MeasAndMobParametersFRX_Diff != nil {
			if err := ie.MeasAndMobParametersFRX_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityAddFRXModeConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. phy-ParametersFRX-Diff: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Phy_ParametersFRX_Diff = new(Phy_ParametersFRX_Diff)
			if err := ie.Phy_ParametersFRX_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. measAndMobParametersFRX-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasAndMobParametersFRX_Diff = new(MeasAndMobParametersFRX_Diff)
			if err := ie.MeasAndMobParametersFRX_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
