// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RLC-ChannelToAddModList-r17 (line 26996).
// SL-RLC-ChannelToAddModList-r17 ::= SEQUENCE (SIZE (1..maxSL-LCID-r16)) OF SL-RLC-ChannelConfig-r17

var sLRLCChannelToAddModListR17SizeConstraints = per.SizeRange(1, common.MaxSL_LCID_r16)

type SL_RLC_ChannelToAddModList_r17 []SL_RLC_ChannelConfig_r17

func (ie *SL_RLC_ChannelToAddModList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sLRLCChannelToAddModListR17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SL_RLC_ChannelToAddModList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sLRLCChannelToAddModListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SL_RLC_ChannelToAddModList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
