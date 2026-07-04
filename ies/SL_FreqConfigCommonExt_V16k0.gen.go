// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-FreqConfigCommonExt-v16k0 (line 27321).

var sLFreqConfigCommonExtV16k0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalSpectrumEmission-r16", Optional: true},
	},
}

type SL_FreqConfigCommonExt_V16k0 struct {
	AdditionalSpectrumEmission_r16 *AdditionalSpectrumEmission
}

func (ie *SL_FreqConfigCommonExt_V16k0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLFreqConfigCommonExtV16k0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalSpectrumEmission_r16 != nil}); err != nil {
		return err
	}

	// 2. additionalSpectrumEmission-r16: ref
	{
		if ie.AdditionalSpectrumEmission_r16 != nil {
			if err := ie.AdditionalSpectrumEmission_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_FreqConfigCommonExt_V16k0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLFreqConfigCommonExtV16k0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. additionalSpectrumEmission-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AdditionalSpectrumEmission_r16 = new(AdditionalSpectrumEmission)
			if err := ie.AdditionalSpectrumEmission_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
