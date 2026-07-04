// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDU-SessionID (line 11549).
// PDU-SessionID ::=   INTEGER (0..255)

var pDUSessionIDConstraints = per.Constrained(0, 255)

type PDU_SessionID struct {
	Value int64
}

func (v *PDU_SessionID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, pDUSessionIDConstraints)
}

func (v *PDU_SessionID) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(pDUSessionIDConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
