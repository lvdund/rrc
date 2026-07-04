// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v15n0 (line 16694).

var bandCombinationV15n0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mrdc-Parameters-v15n0"},
	},
}

type BandCombination_V15n0 struct {
	Mrdc_Parameters_V15n0 MRDC_Parameters_V15n0
}

func (ie *BandCombination_V15n0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV15n0Constraints)
	_ = seq

	// 1. mrdc-Parameters-v15n0: ref
	{
		if err := ie.Mrdc_Parameters_V15n0.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandCombination_V15n0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV15n0Constraints)
	_ = seq

	// 1. mrdc-Parameters-v15n0: ref
	{
		if err := ie.Mrdc_Parameters_V15n0.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
