// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqSeparationClassUL-v1620 (line 20761).
// FreqSeparationClassUL-v1620 ::= ENUMERATED {mhz1000}

const (
	FreqSeparationClassUL_v1620_Mhz1000 = 0
)

var freqSeparationClassULV1620Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FreqSeparationClassUL_v1620_Mhz1000},
}

type FreqSeparationClassUL_v1620 struct {
	Value int64
}

func (v *FreqSeparationClassUL_v1620) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, freqSeparationClassULV1620Constraints)
}

func (v *FreqSeparationClassUL_v1620) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(freqSeparationClassULV1620Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
