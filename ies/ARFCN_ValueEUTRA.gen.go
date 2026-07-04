// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ARFCN-ValueEUTRA (line 5019).
// ARFCN-ValueEUTRA ::=                INTEGER (0..maxEARFCN)

var aRFCNValueEUTRAConstraints = per.Constrained(0, common.MaxEARFCN)

type ARFCN_ValueEUTRA struct {
	Value int64
}

func (v *ARFCN_ValueEUTRA) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, aRFCNValueEUTRAConstraints)
}

func (v *ARFCN_ValueEUTRA) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(aRFCNValueEUTRAConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
