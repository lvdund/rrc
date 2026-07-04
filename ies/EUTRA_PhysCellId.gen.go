// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-PhysCellId (line 26188).
// EUTRA-PhysCellId ::=                        INTEGER (0..503)

var eUTRAPhysCellIdConstraints = per.Constrained(0, 503)

type EUTRA_PhysCellId struct {
	Value int64
}

func (v *EUTRA_PhysCellId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, eUTRAPhysCellIdConstraints)
}

func (v *EUTRA_PhysCellId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(eUTRAPhysCellIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
