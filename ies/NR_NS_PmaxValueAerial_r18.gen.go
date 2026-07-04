// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-NS-PmaxValueAerial-r18 (line 10759).

var nRNSPmaxValueAerialR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalPmax-r18", Optional: true},
		{Name: "additionalSpectrumEmission-r18"},
	},
}

type NR_NS_PmaxValueAerial_r18 struct {
	AdditionalPmax_r18             *P_Max
	AdditionalSpectrumEmission_r18 AdditionalSpectrumEmission_r18
}

func (ie *NR_NS_PmaxValueAerial_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRNSPmaxValueAerialR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalPmax_r18 != nil}); err != nil {
		return err
	}

	// 2. additionalPmax-r18: ref
	{
		if ie.AdditionalPmax_r18 != nil {
			if err := ie.AdditionalPmax_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. additionalSpectrumEmission-r18: ref
	{
		if err := ie.AdditionalSpectrumEmission_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NR_NS_PmaxValueAerial_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRNSPmaxValueAerialR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. additionalPmax-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AdditionalPmax_r18 = new(P_Max)
			if err := ie.AdditionalPmax_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. additionalSpectrumEmission-r18: ref
	{
		if err := ie.AdditionalSpectrumEmission_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
