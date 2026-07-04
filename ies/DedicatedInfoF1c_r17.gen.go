// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DedicatedInfoF1c-r17 (line 26129).
// DedicatedInfoF1c-r17 ::=        OCTET STRING

var dedicatedInfoF1cR17Constraints = per.SizeConstraints{}

type DedicatedInfoF1c_r17 struct {
	Value []byte
}

func (v *DedicatedInfoF1c_r17) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, dedicatedInfoF1cR17Constraints)
}

func (v *DedicatedInfoF1c_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(dedicatedInfoF1cR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
