// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TAG-Id (line 15961).
// TAG-Id ::=                          INTEGER (0..maxNrofTAGs-1)

var tAGIdConstraints = per.Constrained(0, common.MaxNrofTAGs_1)

type TAG_Id struct {
	Value int64
}

func (v *TAG_Id) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, tAGIdConstraints)
}

func (v *TAG_Id) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(tAGIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
