// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NumberOfMsg3-Repetitions-r17 (line 5382).
// NumberOfMsg3-Repetitions-r17::=         ENUMERATED {n1, n2, n3, n4, n7, n8, n12, n16}

const (
	NumberOfMsg3_Repetitions_r17_N1  = 0
	NumberOfMsg3_Repetitions_r17_N2  = 1
	NumberOfMsg3_Repetitions_r17_N3  = 2
	NumberOfMsg3_Repetitions_r17_N4  = 3
	NumberOfMsg3_Repetitions_r17_N7  = 4
	NumberOfMsg3_Repetitions_r17_N8  = 5
	NumberOfMsg3_Repetitions_r17_N12 = 6
	NumberOfMsg3_Repetitions_r17_N16 = 7
)

var numberOfMsg3RepetitionsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NumberOfMsg3_Repetitions_r17_N1, NumberOfMsg3_Repetitions_r17_N2, NumberOfMsg3_Repetitions_r17_N3, NumberOfMsg3_Repetitions_r17_N4, NumberOfMsg3_Repetitions_r17_N7, NumberOfMsg3_Repetitions_r17_N8, NumberOfMsg3_Repetitions_r17_N12, NumberOfMsg3_Repetitions_r17_N16},
}

type NumberOfMsg3_Repetitions_r17 struct {
	Value int64
}

func (v *NumberOfMsg3_Repetitions_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, numberOfMsg3RepetitionsR17Constraints)
}

func (v *NumberOfMsg3_Repetitions_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(numberOfMsg3RepetitionsR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
