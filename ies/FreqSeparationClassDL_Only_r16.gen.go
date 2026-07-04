// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqSeparationClassDL-Only-r16 (line 20766).
// FreqSeparationClassDL-Only-r16 ::= ENUMERATED {mhz200, mhz400, mhz600, mhz800, mhz1000, mhz1200}

const (
	FreqSeparationClassDL_Only_r16_Mhz200  = 0
	FreqSeparationClassDL_Only_r16_Mhz400  = 1
	FreqSeparationClassDL_Only_r16_Mhz600  = 2
	FreqSeparationClassDL_Only_r16_Mhz800  = 3
	FreqSeparationClassDL_Only_r16_Mhz1000 = 4
	FreqSeparationClassDL_Only_r16_Mhz1200 = 5
)

var freqSeparationClassDLOnlyR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FreqSeparationClassDL_Only_r16_Mhz200, FreqSeparationClassDL_Only_r16_Mhz400, FreqSeparationClassDL_Only_r16_Mhz600, FreqSeparationClassDL_Only_r16_Mhz800, FreqSeparationClassDL_Only_r16_Mhz1000, FreqSeparationClassDL_Only_r16_Mhz1200},
}

type FreqSeparationClassDL_Only_r16 struct {
	Value int64
}

func (v *FreqSeparationClassDL_Only_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, freqSeparationClassDLOnlyR16Constraints)
}

func (v *FreqSeparationClassDL_Only_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(freqSeparationClassDLOnlyR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
