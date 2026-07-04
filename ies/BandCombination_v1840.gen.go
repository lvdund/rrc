// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1840 (line 16808).

var bandCombinationV1840Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mrdc-Parameters-v1840", Optional: true},
	},
}

type BandCombination_v1840 struct {
	Mrdc_Parameters_v1840 *MRDC_Parameters_v1840
}

func (ie *BandCombination_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1840Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mrdc_Parameters_v1840 != nil}); err != nil {
		return err
	}

	// 2. mrdc-Parameters-v1840: ref
	{
		if ie.Mrdc_Parameters_v1840 != nil {
			if err := ie.Mrdc_Parameters_v1840.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1840Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mrdc-Parameters-v1840: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mrdc_Parameters_v1840 = new(MRDC_Parameters_v1840)
			if err := ie.Mrdc_Parameters_v1840.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
