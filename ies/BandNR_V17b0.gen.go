// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandNR-v17b0 (line 24667).

var bandNRV17b0Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mimo-ParametersPerBand-v17b0", Optional: true},
	},
}

type BandNR_V17b0 struct {
	Mimo_ParametersPerBand_V17b0 *MIMO_ParametersPerBand_V17b0
}

func (ie *BandNR_V17b0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandNRV17b0Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mimo_ParametersPerBand_V17b0 != nil}); err != nil {
		return err
	}

	// 3. mimo-ParametersPerBand-v17b0: ref
	{
		if ie.Mimo_ParametersPerBand_V17b0 != nil {
			if err := ie.Mimo_ParametersPerBand_V17b0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandNR_V17b0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandNRV17b0Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mimo-ParametersPerBand-v17b0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mimo_ParametersPerBand_V17b0 = new(MIMO_ParametersPerBand_V17b0)
			if err := ie.Mimo_ParametersPerBand_V17b0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
