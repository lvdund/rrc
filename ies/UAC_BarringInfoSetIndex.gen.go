// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UAC-BarringInfoSetIndex (line 16173).
// UAC-BarringInfoSetIndex ::=                INTEGER (1..maxBarringInfoSet)

var uACBarringInfoSetIndexConstraints = per.Constrained(1, common.MaxBarringInfoSet)

type UAC_BarringInfoSetIndex struct {
	Value int64
}

func (v *UAC_BarringInfoSetIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, uACBarringInfoSetIndexConstraints)
}

func (v *UAC_BarringInfoSetIndex) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(uACBarringInfoSetIndexConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
