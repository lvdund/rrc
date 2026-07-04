// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqSeparationClass (line 20757).
// FreqSeparationClass ::= ENUMERATED { mhz800, mhz1200, mhz1400, ..., mhz400-v1650, mhz600-v1650}

const (
	FreqSeparationClass_Mhz800  = 0
	FreqSeparationClass_Mhz1200 = 1
	FreqSeparationClass_Mhz1400 = 2
)

var freqSeparationClassConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{FreqSeparationClass_Mhz800, FreqSeparationClass_Mhz1200, FreqSeparationClass_Mhz1400},
	ExtValues:  []int64{3, 4},
}

type FreqSeparationClass struct {
	Value int64
}

func (v *FreqSeparationClass) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, freqSeparationClassConstraints)
}

func (v *FreqSeparationClass) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(freqSeparationClassConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
