// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SDT-CG-Config-r17 (line 1376).
// SDT-CG-Config-r17 ::= OCTET STRING (CONTAINING SDT-MAC-PHY-CG-Config-r17)

var sDTCGConfigR17Constraints = per.SizeConstraints{}

type SDT_CG_Config_r17 struct {
	Value []byte
}

func (v *SDT_CG_Config_r17) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, sDTCGConfigR17Constraints)
}

func (v *SDT_CG_Config_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(sDTCGConfigR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
