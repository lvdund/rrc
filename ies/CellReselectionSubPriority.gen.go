// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellReselectionSubPriority (line 5892).
// CellReselectionSubPriority ::=          ENUMERATED {oDot2, oDot4, oDot6, oDot8}

const (
	CellReselectionSubPriority_ODot2 = 0
	CellReselectionSubPriority_ODot4 = 1
	CellReselectionSubPriority_ODot6 = 2
	CellReselectionSubPriority_ODot8 = 3
)

var cellReselectionSubPriorityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellReselectionSubPriority_ODot2, CellReselectionSubPriority_ODot4, CellReselectionSubPriority_ODot6, CellReselectionSubPriority_ODot8},
}

type CellReselectionSubPriority struct {
	Value int64
}

func (v *CellReselectionSubPriority) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, cellReselectionSubPriorityConstraints)
}

func (v *CellReselectionSubPriority) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(cellReselectionSubPriorityConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
