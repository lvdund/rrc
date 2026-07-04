// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LogicalChannelIdentityExt-r17 (line 14043).
// LogicalChannelIdentityExt-r17 ::=           INTEGER (320..65855)

var logicalChannelIdentityExtR17Constraints = per.Constrained(320, 65855)

type LogicalChannelIdentityExt_r17 struct {
	Value int64
}

func (v *LogicalChannelIdentityExt_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, logicalChannelIdentityExtR17Constraints)
}

func (v *LogicalChannelIdentityExt_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(logicalChannelIdentityExtR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
