// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-CG-MaxTransNum-r16 (line 27081).

var sLCGMaxTransNumR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-Priority-r16"},
		{Name: "sl-MaxTransNum-r16"},
	},
}

var sLCGMaxTransNumR16SlPriorityR16Constraints = per.Constrained(1, 8)

var sLCGMaxTransNumR16SlMaxTransNumR16Constraints = per.Constrained(1, 32)

type SL_CG_MaxTransNum_r16 struct {
	Sl_Priority_r16    int64
	Sl_MaxTransNum_r16 int64
}

func (ie *SL_CG_MaxTransNum_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLCGMaxTransNumR16Constraints)
	_ = seq

	// 1. sl-Priority-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_Priority_r16, sLCGMaxTransNumR16SlPriorityR16Constraints); err != nil {
			return err
		}
	}

	// 2. sl-MaxTransNum-r16: integer
	{
		if err := e.EncodeInteger(ie.Sl_MaxTransNum_r16, sLCGMaxTransNumR16SlMaxTransNumR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_CG_MaxTransNum_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLCGMaxTransNumR16Constraints)
	_ = seq

	// 1. sl-Priority-r16: integer
	{
		v0, err := d.DecodeInteger(sLCGMaxTransNumR16SlPriorityR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Priority_r16 = v0
	}

	// 2. sl-MaxTransNum-r16: integer
	{
		v1, err := d.DecodeInteger(sLCGMaxTransNumR16SlMaxTransNumR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_MaxTransNum_r16 = v1
	}

	return nil
}
