// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PagingCycle (line 1326).
// PagingCycle ::=                     ENUMERATED {rf32, rf64, rf128, rf256}

const (
	PagingCycle_Rf32  = 0
	PagingCycle_Rf64  = 1
	PagingCycle_Rf128 = 2
	PagingCycle_Rf256 = 3
)

var pagingCycleConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PagingCycle_Rf32, PagingCycle_Rf64, PagingCycle_Rf128, PagingCycle_Rf256},
}

type PagingCycle struct {
	Value int64
}

func (v *PagingCycle) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, pagingCycleConstraints)
}

func (v *PagingCycle) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(pagingCycleConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
