// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: QFI (line 12749).
// QFI ::=                             INTEGER (0..maxQFI)

var qFIConstraints = per.Constrained(0, common.MaxQFI)

type QFI struct {
	Value int64
}

func (v *QFI) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, qFIConstraints)
}

func (v *QFI) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(qFIConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
