// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasSequence-r18 (line 10040).
// MeasSequence-r18 ::=             INTEGER (1..maxMeasSequence-r18)

var measSequenceR18Constraints = per.Constrained(1, common.MaxMeasSequence_r18)

type MeasSequence_r18 struct {
	Value int64
}

func (v *MeasSequence_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, measSequenceR18Constraints)
}

func (v *MeasSequence_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(measSequenceR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
