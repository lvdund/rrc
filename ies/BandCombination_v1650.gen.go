// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1650 (line 16723).

var bandCombinationV1650Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNRDC-v1650", Optional: true},
	},
}

type BandCombination_v1650 struct {
	Ca_ParametersNRDC_v1650 *CA_ParametersNRDC_v1650
}

func (ie *BandCombination_v1650) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1650Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNRDC_v1650 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNRDC-v1650: ref
	{
		if ie.Ca_ParametersNRDC_v1650 != nil {
			if err := ie.Ca_ParametersNRDC_v1650.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1650) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1650Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNRDC-v1650: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNRDC_v1650 = new(CA_ParametersNRDC_v1650)
			if err := ie.Ca_ParametersNRDC_v1650.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
