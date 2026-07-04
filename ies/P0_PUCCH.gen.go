// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: P0-PUCCH (line 12250).

var p0PUCCHConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "p0-PUCCH-Id"},
		{Name: "p0-PUCCH-Value"},
	},
}

var p0PUCCHP0PUCCHValueConstraints = per.Constrained(-16, 15)

type P0_PUCCH struct {
	P0_PUCCH_Id    P0_PUCCH_Id
	P0_PUCCH_Value int64
}

func (ie *P0_PUCCH) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(p0PUCCHConstraints)
	_ = seq

	// 1. p0-PUCCH-Id: ref
	{
		if err := ie.P0_PUCCH_Id.Encode(e); err != nil {
			return err
		}
	}

	// 2. p0-PUCCH-Value: integer
	{
		if err := e.EncodeInteger(ie.P0_PUCCH_Value, p0PUCCHP0PUCCHValueConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *P0_PUCCH) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(p0PUCCHConstraints)
	_ = seq

	// 1. p0-PUCCH-Id: ref
	{
		if err := ie.P0_PUCCH_Id.Decode(d); err != nil {
			return err
		}
	}

	// 2. p0-PUCCH-Value: integer
	{
		v1, err := d.DecodeInteger(p0PUCCHP0PUCCHValueConstraints)
		if err != nil {
			return err
		}
		ie.P0_PUCCH_Value = v1
	}

	return nil
}
