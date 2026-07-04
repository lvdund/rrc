// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RAT-Type (line 23523).
// RAT-Type ::= ENUMERATED {nr, eutra-nr, eutra, utra-fdd-v1610, ...}

const (
	RAT_Type_Nr             = 0
	RAT_Type_Eutra_Nr       = 1
	RAT_Type_Eutra          = 2
	RAT_Type_Utra_Fdd_v1610 = 3
)

var rATTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{RAT_Type_Nr, RAT_Type_Eutra_Nr, RAT_Type_Eutra, RAT_Type_Utra_Fdd_v1610},
}

type RAT_Type struct {
	Value int64
}

func (v *RAT_Type) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, rATTypeConstraints)
}

func (v *RAT_Type) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(rATTypeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
