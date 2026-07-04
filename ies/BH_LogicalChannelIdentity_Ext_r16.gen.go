// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BH-LogicalChannelIdentity-Ext-r16 (line 5196).
// BH-LogicalChannelIdentity-Ext-r16 ::=   INTEGER (320.. maxLC-ID-Iab-r16)

var bHLogicalChannelIdentityExtR16Constraints = per.Constrained(320, common.MaxLC_ID_Iab_r16)

type BH_LogicalChannelIdentity_Ext_r16 struct {
	Value int64
}

func (v *BH_LogicalChannelIdentity_Ext_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, bHLogicalChannelIdentityExtR16Constraints)
}

func (v *BH_LogicalChannelIdentity_Ext_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(bHLogicalChannelIdentityExtR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
