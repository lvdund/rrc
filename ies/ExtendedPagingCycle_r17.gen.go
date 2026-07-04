// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ExtendedPagingCycle-r17 (line 1510).
// ExtendedPagingCycle-r17 ::=         ENUMERATED {rf256, rf512, rf1024, spare1}

const (
	ExtendedPagingCycle_r17_Rf256  = 0
	ExtendedPagingCycle_r17_Rf512  = 1
	ExtendedPagingCycle_r17_Rf1024 = 2
	ExtendedPagingCycle_r17_Spare1 = 3
)

var extendedPagingCycleR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ExtendedPagingCycle_r17_Rf256, ExtendedPagingCycle_r17_Rf512, ExtendedPagingCycle_r17_Rf1024, ExtendedPagingCycle_r17_Spare1},
}

type ExtendedPagingCycle_r17 struct {
	Value int64
}

func (v *ExtendedPagingCycle_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, extendedPagingCycleR17Constraints)
}

func (v *ExtendedPagingCycle_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(extendedPagingCycleR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
