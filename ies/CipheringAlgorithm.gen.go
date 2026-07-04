// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CipheringAlgorithm (line 14564).

const (
	CipheringAlgorithm_Nea0   = 0
	CipheringAlgorithm_Nea1   = 1
	CipheringAlgorithm_Nea2   = 2
	CipheringAlgorithm_Nea3   = 3
	CipheringAlgorithm_Spare4 = 4
	CipheringAlgorithm_Spare3 = 5
	CipheringAlgorithm_Spare2 = 6
	CipheringAlgorithm_Spare1 = 7
)

var cipheringAlgorithmConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{CipheringAlgorithm_Nea0, CipheringAlgorithm_Nea1, CipheringAlgorithm_Nea2, CipheringAlgorithm_Nea3, CipheringAlgorithm_Spare4, CipheringAlgorithm_Spare3, CipheringAlgorithm_Spare2, CipheringAlgorithm_Spare1},
}

type CipheringAlgorithm struct {
	Value int64
}

func (v *CipheringAlgorithm) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, cipheringAlgorithmConstraints)
}

func (v *CipheringAlgorithm) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(cipheringAlgorithmConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
