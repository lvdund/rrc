// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RejectWaitTime (line 13338).
// RejectWaitTime ::=                  INTEGER (1..16)

var rejectWaitTimeConstraints = per.Constrained(1, 16)

type RejectWaitTime struct {
	Value int64
}

func (v *RejectWaitTime) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rejectWaitTimeConstraints)
}

func (v *RejectWaitTime) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rejectWaitTimeConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
