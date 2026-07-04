// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRC-TransactionIdentifier (line 26645).
// RRC-TransactionIdentifier ::=       INTEGER (0..3)

var rRCTransactionIdentifierConstraints = per.Constrained(0, 3)

type RRC_TransactionIdentifier struct {
	Value int64
}

func (v *RRC_TransactionIdentifier) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, rRCTransactionIdentifierConstraints)
}

func (v *RRC_TransactionIdentifier) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(rRCTransactionIdentifierConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
