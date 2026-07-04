// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LogicalChannelIdentity (line 8639).
// LogicalChannelIdentity ::=          INTEGER (1..maxLC-ID)

var logicalChannelIdentityConstraints = per.Constrained(1, common.MaxLC_ID)

type LogicalChannelIdentity struct {
	Value int64
}

func (v *LogicalChannelIdentity) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, logicalChannelIdentityConstraints)
}

func (v *LogicalChannelIdentity) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(logicalChannelIdentityConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
