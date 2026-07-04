// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AggregatedBandwidth (line 20751).

const (
	AggregatedBandwidth_Mhz50  = 0
	AggregatedBandwidth_Mhz100 = 1
	AggregatedBandwidth_Mhz150 = 2
	AggregatedBandwidth_Mhz200 = 3
	AggregatedBandwidth_Mhz250 = 4
	AggregatedBandwidth_Mhz300 = 5
	AggregatedBandwidth_Mhz350 = 6
	AggregatedBandwidth_Mhz400 = 7
	AggregatedBandwidth_Mhz450 = 8
	AggregatedBandwidth_Mhz500 = 9
	AggregatedBandwidth_Mhz550 = 10
	AggregatedBandwidth_Mhz600 = 11
	AggregatedBandwidth_Mhz650 = 12
	AggregatedBandwidth_Mhz700 = 13
	AggregatedBandwidth_Mhz750 = 14
	AggregatedBandwidth_Mhz800 = 15
)

var aggregatedBandwidthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AggregatedBandwidth_Mhz50, AggregatedBandwidth_Mhz100, AggregatedBandwidth_Mhz150, AggregatedBandwidth_Mhz200, AggregatedBandwidth_Mhz250, AggregatedBandwidth_Mhz300, AggregatedBandwidth_Mhz350, AggregatedBandwidth_Mhz400, AggregatedBandwidth_Mhz450, AggregatedBandwidth_Mhz500, AggregatedBandwidth_Mhz550, AggregatedBandwidth_Mhz600, AggregatedBandwidth_Mhz650, AggregatedBandwidth_Mhz700, AggregatedBandwidth_Mhz750, AggregatedBandwidth_Mhz800},
}

type AggregatedBandwidth struct {
	Value int64
}

func (v *AggregatedBandwidth) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, aggregatedBandwidthConstraints)
}

func (v *AggregatedBandwidth) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(aggregatedBandwidthConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
