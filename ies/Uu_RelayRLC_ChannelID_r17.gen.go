// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: Uu-RelayRLC-ChannelID-r17 (line 16365).
// Uu-RelayRLC-ChannelID-r17 ::= INTEGER (1..maxLC-ID)

var uuRelayRLCChannelIDR17Constraints = per.Constrained(1, common.MaxLC_ID)

type Uu_RelayRLC_ChannelID_r17 struct {
	Value int64
}

func (v *Uu_RelayRLC_ChannelID_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, uuRelayRLCChannelIDR17Constraints)
}

func (v *Uu_RelayRLC_ChannelID_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(uuRelayRLCChannelIDR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
