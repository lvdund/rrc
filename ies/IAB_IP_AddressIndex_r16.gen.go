// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IAB-IP-AddressIndex-r16 (line 26227).
// IAB-IP-AddressIndex-r16 ::= INTEGER (1..maxIAB-IP-Address-r16)

var iABIPAddressIndexR16Constraints = per.Constrained(1, common.MaxIAB_IP_Address_r16)

type IAB_IP_AddressIndex_r16 struct {
	Value int64
}

func (v *IAB_IP_AddressIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, iABIPAddressIndexR16Constraints)
}

func (v *IAB_IP_AddressIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(iABIPAddressIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
