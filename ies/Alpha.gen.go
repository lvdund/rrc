// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Alpha (line 4978).
// Alpha ::=                       ENUMERATED {alpha0, alpha04, alpha05, alpha06, alpha07, alpha08, alpha09, alpha1}

const (
	Alpha_Alpha0  = 0
	Alpha_Alpha04 = 1
	Alpha_Alpha05 = 2
	Alpha_Alpha06 = 3
	Alpha_Alpha07 = 4
	Alpha_Alpha08 = 5
	Alpha_Alpha09 = 6
	Alpha_Alpha1  = 7
)

var alphaConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Alpha_Alpha0, Alpha_Alpha04, Alpha_Alpha05, Alpha_Alpha06, Alpha_Alpha07, Alpha_Alpha08, Alpha_Alpha09, Alpha_Alpha1},
}

type Alpha struct {
	Value int64
}

func (v *Alpha) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, alphaConstraints)
}

func (v *Alpha) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(alphaConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
