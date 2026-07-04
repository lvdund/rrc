// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRB-Identity (line 15276).
// SRB-Identity ::=                    INTEGER (1..3)

var sRBIdentityConstraints = per.Constrained(1, 3)

type SRB_Identity struct {
	Value int64
}

func (v *SRB_Identity) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRBIdentityConstraints)
}

func (v *SRB_Identity) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRBIdentityConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
