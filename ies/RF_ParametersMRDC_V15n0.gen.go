// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RF-ParametersMRDC-v15n0 (line 24804).

var rFParametersMRDCV15n0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandCombinationList-v15n0", Optional: true},
	},
}

type RF_ParametersMRDC_V15n0 struct {
	SupportedBandCombinationList_V15n0 *BandCombinationList_V15n0
}

func (ie *RF_ParametersMRDC_V15n0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersMRDCV15n0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandCombinationList_V15n0 != nil}); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-v15n0: ref
	{
		if ie.SupportedBandCombinationList_V15n0 != nil {
			if err := ie.SupportedBandCombinationList_V15n0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RF_ParametersMRDC_V15n0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersMRDCV15n0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-v15n0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_V15n0 = new(BandCombinationList_V15n0)
			if err := ie.SupportedBandCombinationList_V15n0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
