// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1760 (line 16769).

var bandCombinationV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1760"},
		{Name: "ca-ParametersNRDC-v1760"},
	},
}

type BandCombination_v1760 struct {
	Ca_ParametersNR_v1760   CA_ParametersNR_v1760
	Ca_ParametersNRDC_v1760 CA_ParametersNRDC_v1760
}

func (ie *BandCombination_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1760Constraints)
	_ = seq

	// 1. ca-ParametersNR-v1760: ref
	{
		if err := ie.Ca_ParametersNR_v1760.Encode(e); err != nil {
			return err
		}
	}

	// 2. ca-ParametersNRDC-v1760: ref
	{
		if err := ie.Ca_ParametersNRDC_v1760.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandCombination_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1760Constraints)
	_ = seq

	// 1. ca-ParametersNR-v1760: ref
	{
		if err := ie.Ca_ParametersNR_v1760.Decode(d); err != nil {
			return err
		}
	}

	// 2. ca-ParametersNRDC-v1760: ref
	{
		if err := ie.Ca_ParametersNRDC_v1760.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
