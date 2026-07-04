// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SizeDCI-2-9-r19 (line 11880).
// SizeDCI-2-9-r19 ::=                       INTEGER (1..maxDCI-2-9-Size-r18)

var sizeDCI29R19Constraints = per.Constrained(1, common.MaxDCI_2_9_Size_r18)

type SizeDCI_2_9_r19 struct {
	Value int64
}

func (v *SizeDCI_2_9_r19) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sizeDCI29R19Constraints)
}

func (v *SizeDCI_2_9_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sizeDCI29R19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
