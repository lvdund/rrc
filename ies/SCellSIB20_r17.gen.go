// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCellSIB20-r17 (line 5757).
// SCellSIB20-r17 ::= OCTET STRING (CONTAINING SystemInformation)

var sCellSIB20R17Constraints = per.SizeConstraints{}

type SCellSIB20_r17 struct {
	Value []byte
}

func (v *SCellSIB20_r17) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, sCellSIB20R17Constraints)
}

func (v *SCellSIB20_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(sCellSIB20R17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
