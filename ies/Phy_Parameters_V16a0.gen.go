// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-Parameters-v16a0 (line 22783).

var phyParametersV16a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "phy-ParametersCommon-v16a0", Optional: true},
	},
}

type Phy_Parameters_V16a0 struct {
	Phy_ParametersCommon_V16a0 *Phy_ParametersCommon_V16a0
}

func (ie *Phy_Parameters_V16a0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersV16a0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Phy_ParametersCommon_V16a0 != nil}); err != nil {
		return err
	}

	// 2. phy-ParametersCommon-v16a0: ref
	{
		if ie.Phy_ParametersCommon_V16a0 != nil {
			if err := ie.Phy_ParametersCommon_V16a0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Phy_Parameters_V16a0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersV16a0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. phy-ParametersCommon-v16a0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Phy_ParametersCommon_V16a0 = new(Phy_ParametersCommon_V16a0)
			if err := ie.Phy_ParametersCommon_V16a0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
