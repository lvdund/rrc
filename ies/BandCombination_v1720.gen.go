// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1720 (line 16754).

var bandCombinationV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1720", Optional: true},
		{Name: "ca-ParametersNRDC-v1720", Optional: true},
	},
}

type BandCombination_v1720 struct {
	Ca_ParametersNR_v1720   *CA_ParametersNR_v1720
	Ca_ParametersNRDC_v1720 *CA_ParametersNRDC_v1720
}

func (ie *BandCombination_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1720 != nil, ie.Ca_ParametersNRDC_v1720 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1720: ref
	{
		if ie.Ca_ParametersNR_v1720 != nil {
			if err := ie.Ca_ParametersNR_v1720.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1720: ref
	{
		if ie.Ca_ParametersNRDC_v1720 != nil {
			if err := ie.Ca_ParametersNRDC_v1720.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1720: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1720 = new(CA_ParametersNR_v1720)
			if err := ie.Ca_ParametersNR_v1720.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1720: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_v1720 = new(CA_ParametersNRDC_v1720)
			if err := ie.Ca_ParametersNRDC_v1720.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
