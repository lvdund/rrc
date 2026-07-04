// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-NS-PmaxValue (line 26180).

var eUTRANSPmaxValueConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalPmax", Optional: true},
		{Name: "additionalSpectrumEmission", Optional: true},
	},
}

var eUTRANSPmaxValueAdditionalPmaxConstraints = per.Constrained(-30, 33)

var eUTRANSPmaxValueAdditionalSpectrumEmissionConstraints = per.Constrained(1, 288)

type EUTRA_NS_PmaxValue struct {
	AdditionalPmax             *int64
	AdditionalSpectrumEmission *int64
}

func (ie *EUTRA_NS_PmaxValue) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRANSPmaxValueConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalPmax != nil, ie.AdditionalSpectrumEmission != nil}); err != nil {
		return err
	}

	// 2. additionalPmax: integer
	{
		if ie.AdditionalPmax != nil {
			if err := e.EncodeInteger(*ie.AdditionalPmax, eUTRANSPmaxValueAdditionalPmaxConstraints); err != nil {
				return err
			}
		}
	}

	// 3. additionalSpectrumEmission: integer
	{
		if ie.AdditionalSpectrumEmission != nil {
			if err := e.EncodeInteger(*ie.AdditionalSpectrumEmission, eUTRANSPmaxValueAdditionalSpectrumEmissionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EUTRA_NS_PmaxValue) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRANSPmaxValueConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. additionalPmax: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(eUTRANSPmaxValueAdditionalPmaxConstraints)
			if err != nil {
				return err
			}
			ie.AdditionalPmax = &v
		}
	}

	// 3. additionalSpectrumEmission: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(eUTRANSPmaxValueAdditionalSpectrumEmissionConstraints)
			if err != nil {
				return err
			}
			ie.AdditionalSpectrumEmission = &v
		}
	}

	return nil
}
