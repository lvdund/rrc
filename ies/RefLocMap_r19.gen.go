// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RefLocMap-r19 (line 3864).
// RefLocMap-r19      ::= INTEGER (1..7)

var refLocMapR19Constraints = per.Constrained(1, 7)

type RefLocMap_r19 struct {
	Value int64
}

func (v *RefLocMap_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, refLocMapR19Constraints)
}

func (v *RefLocMap_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(refLocMapR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
