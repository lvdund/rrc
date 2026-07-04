// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-CapabilityAddXDD-Mode (line 25908).

var uENRCapabilityAddXDDModeConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "phy-ParametersXDD-Diff", Optional: true},
		{Name: "mac-ParametersXDD-Diff", Optional: true},
		{Name: "measAndMobParametersXDD-Diff", Optional: true},
	},
}

type UE_NR_CapabilityAddXDD_Mode struct {
	Phy_ParametersXDD_Diff       *Phy_ParametersXDD_Diff
	Mac_ParametersXDD_Diff       *MAC_ParametersXDD_Diff
	MeasAndMobParametersXDD_Diff *MeasAndMobParametersXDD_Diff
}

func (ie *UE_NR_CapabilityAddXDD_Mode) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityAddXDDModeConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Phy_ParametersXDD_Diff != nil, ie.Mac_ParametersXDD_Diff != nil, ie.MeasAndMobParametersXDD_Diff != nil}); err != nil {
		return err
	}

	// 2. phy-ParametersXDD-Diff: ref
	{
		if ie.Phy_ParametersXDD_Diff != nil {
			if err := ie.Phy_ParametersXDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mac-ParametersXDD-Diff: ref
	{
		if ie.Mac_ParametersXDD_Diff != nil {
			if err := ie.Mac_ParametersXDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measAndMobParametersXDD-Diff: ref
	{
		if ie.MeasAndMobParametersXDD_Diff != nil {
			if err := ie.MeasAndMobParametersXDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_CapabilityAddXDD_Mode) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityAddXDDModeConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. phy-ParametersXDD-Diff: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Phy_ParametersXDD_Diff = new(Phy_ParametersXDD_Diff)
			if err := ie.Phy_ParametersXDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mac-ParametersXDD-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mac_ParametersXDD_Diff = new(MAC_ParametersXDD_Diff)
			if err := ie.Mac_ParametersXDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measAndMobParametersXDD-Diff: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasAndMobParametersXDD_Diff = new(MeasAndMobParametersXDD_Diff)
			if err := ie.MeasAndMobParametersXDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
