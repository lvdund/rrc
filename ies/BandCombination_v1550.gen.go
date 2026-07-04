// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1550 (line 16665).

var bandCombinationV1550Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1550"},
	},
}

type BandCombination_v1550 struct {
	Ca_ParametersNR_v1550 CA_ParametersNR_v1550
}

func (ie *BandCombination_v1550) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1550Constraints)
	_ = seq

	// 1. ca-ParametersNR-v1550: ref
	{
		if err := ie.Ca_ParametersNR_v1550.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandCombination_v1550) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1550Constraints)
	_ = seq

	// 1. ca-ParametersNR-v1550: ref
	{
		if err := ie.Ca_ParametersNR_v1550.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
