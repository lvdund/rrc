// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntegrityProtAlgorithm (line 14560).

const (
	IntegrityProtAlgorithm_Nia0   = 0
	IntegrityProtAlgorithm_Nia1   = 1
	IntegrityProtAlgorithm_Nia2   = 2
	IntegrityProtAlgorithm_Nia3   = 3
	IntegrityProtAlgorithm_Spare4 = 4
	IntegrityProtAlgorithm_Spare3 = 5
	IntegrityProtAlgorithm_Spare2 = 6
	IntegrityProtAlgorithm_Spare1 = 7
)

var integrityProtAlgorithmConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{IntegrityProtAlgorithm_Nia0, IntegrityProtAlgorithm_Nia1, IntegrityProtAlgorithm_Nia2, IntegrityProtAlgorithm_Nia3, IntegrityProtAlgorithm_Spare4, IntegrityProtAlgorithm_Spare3, IntegrityProtAlgorithm_Spare2, IntegrityProtAlgorithm_Spare1},
}

type IntegrityProtAlgorithm struct {
	Value int64
}

func (v *IntegrityProtAlgorithm) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, integrityProtAlgorithmConstraints)
}

func (v *IntegrityProtAlgorithm) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(integrityProtAlgorithmConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
