// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReestablishmentCause (line 958).
// ReestablishmentCause ::=            ENUMERATED {reconfigurationFailure, handoverFailure, otherFailure, spare1}

const (
	ReestablishmentCause_ReconfigurationFailure = 0
	ReestablishmentCause_HandoverFailure        = 1
	ReestablishmentCause_OtherFailure           = 2
	ReestablishmentCause_Spare1                 = 3
)

var reestablishmentCauseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReestablishmentCause_ReconfigurationFailure, ReestablishmentCause_HandoverFailure, ReestablishmentCause_OtherFailure, ReestablishmentCause_Spare1},
}

type ReestablishmentCause struct {
	Value int64
}

func (v *ReestablishmentCause) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, reestablishmentCauseConstraints)
}

func (v *ReestablishmentCause) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(reestablishmentCauseConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
