// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-NS-PmaxValue (line 10746).

var nRNSPmaxValueConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalPmax", Optional: true},
		{Name: "additionalSpectrumEmission"},
	},
}

type NR_NS_PmaxValue struct {
	AdditionalPmax             *P_Max
	AdditionalSpectrumEmission AdditionalSpectrumEmission
}

func (ie *NR_NS_PmaxValue) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRNSPmaxValueConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalPmax != nil}); err != nil {
		return err
	}

	// 2. additionalPmax: ref
	{
		if ie.AdditionalPmax != nil {
			if err := ie.AdditionalPmax.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. additionalSpectrumEmission: ref
	{
		if err := ie.AdditionalSpectrumEmission.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NR_NS_PmaxValue) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRNSPmaxValueConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. additionalPmax: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AdditionalPmax = new(P_Max)
			if err := ie.AdditionalPmax.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. additionalSpectrumEmission: ref
	{
		if err := ie.AdditionalSpectrumEmission.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
