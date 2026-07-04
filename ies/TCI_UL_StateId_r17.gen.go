// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TCI-UL-StateId-r17 (line 16062).
// TCI-UL-StateId-r17 ::=              INTEGER (0..maxUL-TCI-1-r17)

var tCIULStateIdR17Constraints = per.Constrained(0, common.MaxUL_TCI_1_r17)

type TCI_UL_StateId_r17 struct {
	Value int64
}

func (v *TCI_UL_StateId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, tCIULStateIdR17Constraints)
}

func (v *TCI_UL_StateId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(tCIULStateIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
