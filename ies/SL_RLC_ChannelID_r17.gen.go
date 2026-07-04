// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RLC-ChannelID-r17 (line 28167).
// SL-RLC-ChannelID-r17 ::=    INTEGER (1..maxSL-LCID-r16)

var sLRLCChannelIDR17Constraints = per.Constrained(1, common.MaxSL_LCID_r16)

type SL_RLC_ChannelID_r17 struct {
	Value int64
}

func (v *SL_RLC_ChannelID_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLRLCChannelIDR17Constraints)
}

func (v *SL_RLC_ChannelID_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLRLCChannelIDR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
