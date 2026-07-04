// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1860 (line 16812).

var bandCombinationV1860Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1860", Optional: true},
	},
}

type BandCombination_v1860 struct {
	Ca_ParametersNR_v1860 *CA_ParametersNR_v1860
}

func (ie *BandCombination_v1860) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1860Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1860 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1860: ref
	{
		if ie.Ca_ParametersNR_v1860 != nil {
			if err := ie.Ca_ParametersNR_v1860.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1860) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1860Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1860: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1860 = new(CA_ParametersNR_v1860)
			if err := ie.Ca_ParametersNR_v1860.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
