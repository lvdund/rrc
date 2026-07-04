// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BWP-Id (line 5313).
// BWP-Id ::=                          INTEGER (0..maxNrofBWPs)

var bWPIdConstraints = per.Constrained(0, common.MaxNrofBWPs)

type BWP_Id struct {
	Value int64
}

func (v *BWP_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, bWPIdConstraints)
}

func (v *BWP_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(bWPIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
