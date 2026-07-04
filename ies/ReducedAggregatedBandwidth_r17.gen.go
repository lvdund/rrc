// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReducedAggregatedBandwidth-r17 (line 2441).
// ReducedAggregatedBandwidth-r17 ::= ENUMERATED {mhz0, mhz100, mhz200, mhz400, mhz800, mhz1200, mhz1600, mhz2000}

const (
	ReducedAggregatedBandwidth_r17_Mhz0    = 0
	ReducedAggregatedBandwidth_r17_Mhz100  = 1
	ReducedAggregatedBandwidth_r17_Mhz200  = 2
	ReducedAggregatedBandwidth_r17_Mhz400  = 3
	ReducedAggregatedBandwidth_r17_Mhz800  = 4
	ReducedAggregatedBandwidth_r17_Mhz1200 = 5
	ReducedAggregatedBandwidth_r17_Mhz1600 = 6
	ReducedAggregatedBandwidth_r17_Mhz2000 = 7
)

var reducedAggregatedBandwidthR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReducedAggregatedBandwidth_r17_Mhz0, ReducedAggregatedBandwidth_r17_Mhz100, ReducedAggregatedBandwidth_r17_Mhz200, ReducedAggregatedBandwidth_r17_Mhz400, ReducedAggregatedBandwidth_r17_Mhz800, ReducedAggregatedBandwidth_r17_Mhz1200, ReducedAggregatedBandwidth_r17_Mhz1600, ReducedAggregatedBandwidth_r17_Mhz2000},
}

type ReducedAggregatedBandwidth_r17 struct {
	Value int64
}

func (v *ReducedAggregatedBandwidth_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, reducedAggregatedBandwidthR17Constraints)
}

func (v *ReducedAggregatedBandwidth_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(reducedAggregatedBandwidthR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
