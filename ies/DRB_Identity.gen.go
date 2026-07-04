// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRB-Identity (line 8105).
// DRB-Identity ::=                    INTEGER (1..32)

var dRBIdentityConstraints = per.Constrained(1, 32)

type DRB_Identity struct {
	Value int64
}

func (v *DRB_Identity) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, dRBIdentityConstraints)
}

func (v *DRB_Identity) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(dRBIdentityConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
