// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReferenceConfiguration-r18 (line 13311).
// ReferenceConfiguration-r18 ::= OCTET STRING (CONTAINING RRCReconfiguration)

var referenceConfigurationR18Constraints = per.SizeConstraints{}

type ReferenceConfiguration_r18 struct {
	Value []byte
}

func (v *ReferenceConfiguration_r18) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, referenceConfigurationR18Constraints)
}

func (v *ReferenceConfiguration_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(referenceConfigurationR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
