// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BWP (line 5230).

var bWPConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "locationAndBandwidth"},
		{Name: "subcarrierSpacing"},
		{Name: "cyclicPrefix", Optional: true},
	},
}

var bWPLocationAndBandwidthConstraints = per.Constrained(0, 37949)

const (
	BWP_CyclicPrefix_Extended = 0
)

var bWPCyclicPrefixConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_CyclicPrefix_Extended},
}

type BWP struct {
	LocationAndBandwidth int64
	SubcarrierSpacing    SubcarrierSpacing
	CyclicPrefix         *int64
}

func (ie *BWP) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CyclicPrefix != nil}); err != nil {
		return err
	}

	// 2. locationAndBandwidth: integer
	{
		if err := e.EncodeInteger(ie.LocationAndBandwidth, bWPLocationAndBandwidthConstraints); err != nil {
			return err
		}
	}

	// 3. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Encode(e); err != nil {
			return err
		}
	}

	// 4. cyclicPrefix: enumerated
	{
		if ie.CyclicPrefix != nil {
			if err := e.EncodeEnumerated(*ie.CyclicPrefix, bWPCyclicPrefixConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BWP) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. locationAndBandwidth: integer
	{
		v0, err := d.DecodeInteger(bWPLocationAndBandwidthConstraints)
		if err != nil {
			return err
		}
		ie.LocationAndBandwidth = v0
	}

	// 3. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Decode(d); err != nil {
			return err
		}
	}

	// 4. cyclicPrefix: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(bWPCyclicPrefixConstraints)
			if err != nil {
				return err
			}
			ie.CyclicPrefix = &idx
		}
	}

	return nil
}
