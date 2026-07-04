// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SchedulingRequestResourceId (line 14324).
// SchedulingRequestResourceId ::=     INTEGER (1..maxNrofSR-Resources)

var schedulingRequestResourceIdConstraints = per.Constrained(1, common.MaxNrofSR_Resources)

type SchedulingRequestResourceId struct {
	Value int64
}

func (v *SchedulingRequestResourceId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, schedulingRequestResourceIdConstraints)
}

func (v *SchedulingRequestResourceId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(schedulingRequestResourceIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
