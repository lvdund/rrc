// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqInfoMBS-r18 (line 28475).

var freqInfoMBSR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreqMBS-r18"},
		{Name: "freqBandIndicatorMBS-r18"},
	},
}

type FreqInfoMBS_r18 struct {
	CarrierFreqMBS_r18       ARFCN_ValueNR
	FreqBandIndicatorMBS_r18 FreqBandIndicatorNR
}

func (ie *FreqInfoMBS_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(freqInfoMBSR18Constraints)
	_ = seq

	// 1. carrierFreqMBS-r18: ref
	{
		if err := ie.CarrierFreqMBS_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. freqBandIndicatorMBS-r18: ref
	{
		if err := ie.FreqBandIndicatorMBS_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FreqInfoMBS_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(freqInfoMBSR18Constraints)
	_ = seq

	// 1. carrierFreqMBS-r18: ref
	{
		if err := ie.CarrierFreqMBS_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. freqBandIndicatorMBS-r18: ref
	{
		if err := ie.FreqBandIndicatorMBS_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
