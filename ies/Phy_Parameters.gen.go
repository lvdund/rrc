// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-Parameters (line 22775).

var phyParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "phy-ParametersCommon", Optional: true},
		{Name: "phy-ParametersXDD-Diff", Optional: true},
		{Name: "phy-ParametersFRX-Diff", Optional: true},
		{Name: "phy-ParametersFR1", Optional: true},
		{Name: "phy-ParametersFR2", Optional: true},
	},
}

type Phy_Parameters struct {
	Phy_ParametersCommon   *Phy_ParametersCommon
	Phy_ParametersXDD_Diff *Phy_ParametersXDD_Diff
	Phy_ParametersFRX_Diff *Phy_ParametersFRX_Diff
	Phy_ParametersFR1      *Phy_ParametersFR1
	Phy_ParametersFR2      *Phy_ParametersFR2
}

func (ie *Phy_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Phy_ParametersCommon != nil, ie.Phy_ParametersXDD_Diff != nil, ie.Phy_ParametersFRX_Diff != nil, ie.Phy_ParametersFR1 != nil, ie.Phy_ParametersFR2 != nil}); err != nil {
		return err
	}

	// 2. phy-ParametersCommon: ref
	{
		if ie.Phy_ParametersCommon != nil {
			if err := ie.Phy_ParametersCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. phy-ParametersXDD-Diff: ref
	{
		if ie.Phy_ParametersXDD_Diff != nil {
			if err := ie.Phy_ParametersXDD_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. phy-ParametersFRX-Diff: ref
	{
		if ie.Phy_ParametersFRX_Diff != nil {
			if err := ie.Phy_ParametersFRX_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. phy-ParametersFR1: ref
	{
		if ie.Phy_ParametersFR1 != nil {
			if err := ie.Phy_ParametersFR1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. phy-ParametersFR2: ref
	{
		if ie.Phy_ParametersFR2 != nil {
			if err := ie.Phy_ParametersFR2.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Phy_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. phy-ParametersCommon: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Phy_ParametersCommon = new(Phy_ParametersCommon)
			if err := ie.Phy_ParametersCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. phy-ParametersXDD-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Phy_ParametersXDD_Diff = new(Phy_ParametersXDD_Diff)
			if err := ie.Phy_ParametersXDD_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. phy-ParametersFRX-Diff: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Phy_ParametersFRX_Diff = new(Phy_ParametersFRX_Diff)
			if err := ie.Phy_ParametersFRX_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. phy-ParametersFR1: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Phy_ParametersFR1 = new(Phy_ParametersFR1)
			if err := ie.Phy_ParametersFR1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. phy-ParametersFR2: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Phy_ParametersFR2 = new(Phy_ParametersFR2)
			if err := ie.Phy_ParametersFR2.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
