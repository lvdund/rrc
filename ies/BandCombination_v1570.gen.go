// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1570 (line 16675).

var bandCombinationV1570Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersEUTRA-v1570"},
	},
}

type BandCombination_v1570 struct {
	Ca_ParametersEUTRA_v1570 CA_ParametersEUTRA_v1570
}

func (ie *BandCombination_v1570) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1570Constraints)
	_ = seq

	// 1. ca-ParametersEUTRA-v1570: ref
	{
		if err := ie.Ca_ParametersEUTRA_v1570.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandCombination_v1570) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1570Constraints)
	_ = seq

	// 1. ca-ParametersEUTRA-v1570: ref
	{
		if err := ie.Ca_ParametersEUTRA_v1570.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
