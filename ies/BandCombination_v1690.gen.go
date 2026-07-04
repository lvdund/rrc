// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1690 (line 16731).

var bandCombinationV1690Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy", Optional: true},
	},
}

type BandCombination_v1690 struct {
	Dummy *CA_ParametersNR_v1690
}

func (ie *BandCombination_v1690) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1690Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dummy != nil}); err != nil {
		return err
	}

	// 2. dummy: ref
	{
		if ie.Dummy != nil {
			if err := ie.Dummy.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1690) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1690Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dummy: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Dummy = new(CA_ParametersNR_v1690)
			if err := ie.Dummy.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
