// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: WLAN-Name-r16 (line 26742).
// WLAN-Name-r16 ::= OCTET STRING (SIZE (1..32))

var wLANNameR16Constraints = per.SizeRange(1, 32)

type WLAN_Name_r16 struct {
	Value []byte
}

func (v *WLAN_Name_r16) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, wLANNameR16Constraints)
}

func (v *WLAN_Name_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(wLANNameR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
