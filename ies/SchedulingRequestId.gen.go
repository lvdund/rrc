// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingRequestId (line 14276).
// SchedulingRequestId ::=             INTEGER (0..7)

var schedulingRequestIdConstraints = per.Constrained(0, 7)

type SchedulingRequestId struct {
	Value int64
}

func (v *SchedulingRequestId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, schedulingRequestIdConstraints)
}

func (v *SchedulingRequestId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(schedulingRequestIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
