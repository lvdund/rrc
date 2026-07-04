// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParametersSidelink-r16 (line 17199).

var bandParametersSidelinkR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "freqBandSidelink-r16"},
	},
}

type BandParametersSidelink_r16 struct {
	FreqBandSidelink_r16 FreqBandIndicatorNR
}

func (ie *BandParametersSidelink_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersSidelinkR16Constraints)
	_ = seq

	// 1. freqBandSidelink-r16: ref
	{
		if err := ie.FreqBandSidelink_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandParametersSidelink_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersSidelinkR16Constraints)
	_ = seq

	// 1. freqBandSidelink-r16: ref
	{
		if err := ie.FreqBandSidelink_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
