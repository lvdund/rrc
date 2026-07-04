// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosRRC-InactiveEnhanced-r18 (line 1452).
// SRS-PosRRC-InactiveEnhanced-r18 ::= OCTET STRING (CONTAINING SRS-PosRRC-InactiveEnhancedConfig-r18)

var sRSPosRRCInactiveEnhancedR18Constraints = per.SizeConstraints{}

type SRS_PosRRC_InactiveEnhanced_r18 struct {
	Value []byte
}

func (v *SRS_PosRRC_InactiveEnhanced_r18) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(v.Value, sRSPosRRCInactiveEnhancedR18Constraints)
}

func (v *SRS_PosRRC_InactiveEnhanced_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeOctetString(sRSPosRRCInactiveEnhancedR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
