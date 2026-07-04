// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReducedMaxBW-FRx-r16 (line 2690).

var reducedMaxBWFRxR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reducedBW-DL-r16"},
		{Name: "reducedBW-UL-r16"},
	},
}

type ReducedMaxBW_FRx_r16 struct {
	ReducedBW_DL_r16 ReducedAggregatedBandwidth
	ReducedBW_UL_r16 ReducedAggregatedBandwidth
}

func (ie *ReducedMaxBW_FRx_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reducedMaxBWFRxR16Constraints)
	_ = seq

	// 1. reducedBW-DL-r16: ref
	{
		if err := ie.ReducedBW_DL_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. reducedBW-UL-r16: ref
	{
		if err := ie.ReducedBW_UL_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReducedMaxBW_FRx_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reducedMaxBWFRxR16Constraints)
	_ = seq

	// 1. reducedBW-DL-r16: ref
	{
		if err := ie.ReducedBW_DL_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. reducedBW-UL-r16: ref
	{
		if err := ie.ReducedBW_UL_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
