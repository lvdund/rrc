// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosRRC-Inactive-r17 (line 1432).
// SRS-PosRRC-Inactive-r17 ::= OCTET STRING (CONTAINING SRS-PosRRC-InactiveConfig-r17)

var sRSPosRRCInactiveR17Constraints = per.SizeConstraints{}

type SRS_PosRRC_Inactive_r17 struct {
	Value []byte
}

func (v *SRS_PosRRC_Inactive_r17) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, sRSPosRRCInactiveR17Constraints)
}

func (v *SRS_PosRRC_Inactive_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(sRSPosRRCInactiveR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
