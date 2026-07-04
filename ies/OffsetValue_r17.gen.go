// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OffsetValue-r17 (line 16403).

var offsetValueR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetValue-r17"},
		{Name: "shift7dot5kHz-r17"},
	},
}

var offsetValueR17OffsetValueR17Constraints = per.Constrained(-20000, 20000)

type OffsetValue_r17 struct {
	OffsetValue_r17   int64
	Shift7dot5kHz_r17 bool
}

func (ie *OffsetValue_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(offsetValueR17Constraints)
	_ = seq

	// 1. offsetValue-r17: integer
	{
		if err := e.EncodeInteger(ie.OffsetValue_r17, offsetValueR17OffsetValueR17Constraints); err != nil {
			return err
		}
	}

	// 2. shift7dot5kHz-r17: boolean
	{
		if err := e.EncodeBoolean(ie.Shift7dot5kHz_r17); err != nil {
			return err
		}
	}

	return nil
}

func (ie *OffsetValue_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(offsetValueR17Constraints)
	_ = seq

	// 1. offsetValue-r17: integer
	{
		v0, err := d.DecodeInteger(offsetValueR17OffsetValueR17Constraints)
		if err != nil {
			return err
		}
		ie.OffsetValue_r17 = v0
	}

	// 2. shift7dot5kHz-r17: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Shift7dot5kHz_r17 = v1
	}

	return nil
}
