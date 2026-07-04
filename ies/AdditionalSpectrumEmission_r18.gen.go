// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AdditionalSpectrumEmission-r18 (line 4947).
// AdditionalSpectrumEmission-r18 ::=          INTEGER (0..39)

var additionalSpectrumEmissionR18Constraints = per.Constrained(0, 39)

type AdditionalSpectrumEmission_r18 struct {
	Value int64
}

func (v *AdditionalSpectrumEmission_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, additionalSpectrumEmissionR18Constraints)
}

func (v *AdditionalSpectrumEmission_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(additionalSpectrumEmissionR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
