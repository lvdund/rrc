// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: WLAN-RSSI-Range-r16 (line 26277).
// WLAN-RSSI-Range-r16 ::= INTEGER(0..141)

var wLANRSSIRangeR16Constraints = per.Constrained(0, 141)

type WLAN_RSSI_Range_r16 struct {
	Value int64
}

func (v *WLAN_RSSI_Range_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, wLANRSSIRangeR16Constraints)
}

func (v *WLAN_RSSI_Range_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(wLANRSSIRangeR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
