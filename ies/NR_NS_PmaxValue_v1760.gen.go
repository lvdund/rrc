// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-NS-PmaxValue-v1760 (line 10753).

var nRNSPmaxValueV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalSpectrumEmission-v1760", Optional: true},
	},
}

type NR_NS_PmaxValue_v1760 struct {
	AdditionalSpectrumEmission_v1760 *AdditionalSpectrumEmission_v1760
}

func (ie *NR_NS_PmaxValue_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRNSPmaxValueV1760Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalSpectrumEmission_v1760 != nil}); err != nil {
		return err
	}

	// 2. additionalSpectrumEmission-v1760: ref
	{
		if ie.AdditionalSpectrumEmission_v1760 != nil {
			if err := ie.AdditionalSpectrumEmission_v1760.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NR_NS_PmaxValue_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRNSPmaxValueV1760Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. additionalSpectrumEmission-v1760: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AdditionalSpectrumEmission_v1760 = new(AdditionalSpectrumEmission_v1760)
			if err := ie.AdditionalSpectrumEmission_v1760.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
