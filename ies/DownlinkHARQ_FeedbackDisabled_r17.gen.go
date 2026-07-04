// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DownlinkHARQ-FeedbackDisabled-r17 (line 11507).
// DownlinkHARQ-FeedbackDisabled-r17 ::= BIT STRING (SIZE (32))

var downlinkHARQFeedbackDisabledR17Constraints = per.FixedSize(32)

type DownlinkHARQ_FeedbackDisabled_r17 struct {
	Value per.BitString
}

func (v *DownlinkHARQ_FeedbackDisabled_r17) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, downlinkHARQFeedbackDisabledR17Constraints)
}

func (v *DownlinkHARQ_FeedbackDisabled_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(downlinkHARQFeedbackDisabledR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
