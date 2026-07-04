// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqSeparationClassDL-v1620 (line 20759).
// FreqSeparationClassDL-v1620 ::= ENUMERATED {mhz1000, mhz1600, mhz1800, mhz2000, mhz2200, mhz2400}

const (
	FreqSeparationClassDL_v1620_Mhz1000 = 0
	FreqSeparationClassDL_v1620_Mhz1600 = 1
	FreqSeparationClassDL_v1620_Mhz1800 = 2
	FreqSeparationClassDL_v1620_Mhz2000 = 3
	FreqSeparationClassDL_v1620_Mhz2200 = 4
	FreqSeparationClassDL_v1620_Mhz2400 = 5
)

var freqSeparationClassDLV1620Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FreqSeparationClassDL_v1620_Mhz1000, FreqSeparationClassDL_v1620_Mhz1600, FreqSeparationClassDL_v1620_Mhz1800, FreqSeparationClassDL_v1620_Mhz2000, FreqSeparationClassDL_v1620_Mhz2200, FreqSeparationClassDL_v1620_Mhz2400},
}

type FreqSeparationClassDL_v1620 struct {
	Value int64
}

func (v *FreqSeparationClassDL_v1620) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, freqSeparationClassDLV1620Constraints)
}

func (v *FreqSeparationClassDL_v1620) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(freqSeparationClassDLV1620Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
