// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReducedAggregatedBandwidth (line 2439).
// ReducedAggregatedBandwidth ::= ENUMERATED {mhz0, mhz10, mhz20, mhz30, mhz40, mhz50, mhz60, mhz80, mhz100, mhz200, mhz300, mhz400}

const (
	ReducedAggregatedBandwidth_Mhz0   = 0
	ReducedAggregatedBandwidth_Mhz10  = 1
	ReducedAggregatedBandwidth_Mhz20  = 2
	ReducedAggregatedBandwidth_Mhz30  = 3
	ReducedAggregatedBandwidth_Mhz40  = 4
	ReducedAggregatedBandwidth_Mhz50  = 5
	ReducedAggregatedBandwidth_Mhz60  = 6
	ReducedAggregatedBandwidth_Mhz80  = 7
	ReducedAggregatedBandwidth_Mhz100 = 8
	ReducedAggregatedBandwidth_Mhz200 = 9
	ReducedAggregatedBandwidth_Mhz300 = 10
	ReducedAggregatedBandwidth_Mhz400 = 11
)

var reducedAggregatedBandwidthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReducedAggregatedBandwidth_Mhz0, ReducedAggregatedBandwidth_Mhz10, ReducedAggregatedBandwidth_Mhz20, ReducedAggregatedBandwidth_Mhz30, ReducedAggregatedBandwidth_Mhz40, ReducedAggregatedBandwidth_Mhz50, ReducedAggregatedBandwidth_Mhz60, ReducedAggregatedBandwidth_Mhz80, ReducedAggregatedBandwidth_Mhz100, ReducedAggregatedBandwidth_Mhz200, ReducedAggregatedBandwidth_Mhz300, ReducedAggregatedBandwidth_Mhz400},
}

type ReducedAggregatedBandwidth struct {
	Value int64
}

func (v *ReducedAggregatedBandwidth) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, reducedAggregatedBandwidthConstraints)
}

func (v *ReducedAggregatedBandwidth) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(reducedAggregatedBandwidthConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
