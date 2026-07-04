// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellIdentity (line 5882).
// CellIdentity ::=                         BIT STRING (SIZE (36))

var cellIdentityConstraints = per.FixedSize(36)

type CellIdentity struct {
	Value per.BitString
}

func (v *CellIdentity) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, cellIdentityConstraints)
}

func (v *CellIdentity) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(cellIdentityConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
