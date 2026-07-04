// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AdditionalSpectrumEmission-v1760 (line 4945).
// AdditionalSpectrumEmission-v1760 ::=        INTEGER (8..39)

var additionalSpectrumEmissionV1760Constraints = per.Constrained(8, 39)

type AdditionalSpectrumEmission_v1760 struct {
	Value int64
}

func (v *AdditionalSpectrumEmission_v1760) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, additionalSpectrumEmissionV1760Constraints)
}

func (v *AdditionalSpectrumEmission_v1760) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(additionalSpectrumEmissionV1760Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
