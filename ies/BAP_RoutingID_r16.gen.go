// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BAP-RoutingID-r16 (line 5100).

var bAPRoutingIDR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bap-Address-r16"},
		{Name: "bap-PathId-r16"},
	},
}

var bAPRoutingIDR16BapAddressR16Constraints = per.FixedSize(10)

var bAPRoutingIDR16BapPathIdR16Constraints = per.FixedSize(10)

type BAP_RoutingID_r16 struct {
	Bap_Address_r16 per.BitString
	Bap_PathId_r16  per.BitString
}

func (ie *BAP_RoutingID_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bAPRoutingIDR16Constraints)
	_ = seq

	// 1. bap-Address-r16: bit-string
	{
		if err := e.EncodeBitString(ie.Bap_Address_r16, bAPRoutingIDR16BapAddressR16Constraints); err != nil {
			return err
		}
	}

	// 2. bap-PathId-r16: bit-string
	{
		if err := e.EncodeBitString(ie.Bap_PathId_r16, bAPRoutingIDR16BapPathIdR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BAP_RoutingID_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bAPRoutingIDR16Constraints)
	_ = seq

	// 1. bap-Address-r16: bit-string
	{
		v0, err := d.DecodeBitString(bAPRoutingIDR16BapAddressR16Constraints)
		if err != nil {
			return err
		}
		ie.Bap_Address_r16 = v0
	}

	// 2. bap-PathId-r16: bit-string
	{
		v1, err := d.DecodeBitString(bAPRoutingIDR16BapPathIdR16Constraints)
		if err != nil {
			return err
		}
		ie.Bap_PathId_r16 = v1
	}

	return nil
}
