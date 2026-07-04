// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AttemptedBWP-Info-r18 (line 3139).

var attemptedBWPInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "locationAndBandwidth-r18"},
		{Name: "subcarrierSpacing-r18"},
	},
}

var attemptedBWPInfoR18LocationAndBandwidthR18Constraints = per.Constrained(0, 37949)

type AttemptedBWP_Info_r18 struct {
	LocationAndBandwidth_r18 int64
	SubcarrierSpacing_r18    SubcarrierSpacing
}

func (ie *AttemptedBWP_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(attemptedBWPInfoR18Constraints)
	_ = seq

	// 1. locationAndBandwidth-r18: integer
	{
		if err := e.EncodeInteger(ie.LocationAndBandwidth_r18, attemptedBWPInfoR18LocationAndBandwidthR18Constraints); err != nil {
			return err
		}
	}

	// 2. subcarrierSpacing-r18: ref
	{
		if err := ie.SubcarrierSpacing_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AttemptedBWP_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(attemptedBWPInfoR18Constraints)
	_ = seq

	// 1. locationAndBandwidth-r18: integer
	{
		v0, err := d.DecodeInteger(attemptedBWPInfoR18LocationAndBandwidthR18Constraints)
		if err != nil {
			return err
		}
		ie.LocationAndBandwidth_r18 = v0
	}

	// 2. subcarrierSpacing-r18: ref
	{
		if err := ie.SubcarrierSpacing_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
