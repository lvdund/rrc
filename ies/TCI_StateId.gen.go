// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TCI-StateId (line 16030).
// TCI-StateId ::=                     INTEGER (0..maxNrofTCI-States-1)

var tCIStateIdConstraints = per.Constrained(0, common.MaxNrofTCI_States_1)

type TCI_StateId struct {
	Value int64
}

func (v *TCI_StateId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, tCIStateIdConstraints)
}

func (v *TCI_StateId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(tCIStateIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
