// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReducedMaxCCs-r16 (line 2695).

var reducedMaxCCsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedCCsDL-r16"},
		{Name: "reducedCCsUL-r16"},
	},
}

var reducedMaxCCsR16ReducedCCsDLR16Constraints = per.Constrained(0, 31)

var reducedMaxCCsR16ReducedCCsULR16Constraints = per.Constrained(0, 31)

type ReducedMaxCCs_r16 struct {
	ReducedCCsDL_r16 int64
	ReducedCCsUL_r16 int64
}

func (ie *ReducedMaxCCs_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reducedMaxCCsR16Constraints)
	_ = seq

	// 1. reducedCCsDL-r16: integer
	{
		if err := e.EncodeInteger(ie.ReducedCCsDL_r16, reducedMaxCCsR16ReducedCCsDLR16Constraints); err != nil {
			return err
		}
	}

	// 2. reducedCCsUL-r16: integer
	{
		if err := e.EncodeInteger(ie.ReducedCCsUL_r16, reducedMaxCCsR16ReducedCCsULR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReducedMaxCCs_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reducedMaxCCsR16Constraints)
	_ = seq

	// 1. reducedCCsDL-r16: integer
	{
		v0, err := d.DecodeInteger(reducedMaxCCsR16ReducedCCsDLR16Constraints)
		if err != nil {
			return err
		}
		ie.ReducedCCsDL_r16 = v0
	}

	// 2. reducedCCsUL-r16: integer
	{
		v1, err := d.DecodeInteger(reducedMaxCCsR16ReducedCCsULR16Constraints)
		if err != nil {
			return err
		}
		ie.ReducedCCsUL_r16 = v1
	}

	return nil
}
