// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AdditionalSpectrumEmission (line 4943).
// AdditionalSpectrumEmission ::=              INTEGER (0..7)

var additionalSpectrumEmissionConstraints = per.Constrained(0, 7)

type AdditionalSpectrumEmission struct {
	Value int64
}

func (v *AdditionalSpectrumEmission) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, additionalSpectrumEmissionConstraints)
}

func (v *AdditionalSpectrumEmission) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(additionalSpectrumEmissionConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
