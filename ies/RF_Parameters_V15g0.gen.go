// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RF-Parameters-v15g0 (line 23653).

var rFParametersV15g0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandCombinationList-v15g0", Optional: true},
	},
}

type RF_Parameters_V15g0 struct {
	SupportedBandCombinationList_V15g0 *BandCombinationList_V15g0
}

func (ie *RF_Parameters_V15g0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersV15g0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandCombinationList_V15g0 != nil}); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-v15g0: ref
	{
		if ie.SupportedBandCombinationList_V15g0 != nil {
			if err := ie.SupportedBandCombinationList_V15g0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RF_Parameters_V15g0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersV15g0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-v15g0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_V15g0 = new(BandCombinationList_V15g0)
			if err := ie.SupportedBandCombinationList_V15g0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
