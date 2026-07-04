// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultLogging2NR-r16 (line 3417).

var measResultLogging2NRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r16"},
		{Name: "measResultListLoggingNR-r16"},
	},
}

type MeasResultLogging2NR_r16 struct {
	CarrierFreq_r16             ARFCN_ValueNR
	MeasResultListLoggingNR_r16 MeasResultListLoggingNR_r16
}

func (ie *MeasResultLogging2NR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultLogging2NRR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. measResultListLoggingNR-r16: ref
	{
		if err := ie.MeasResultListLoggingNR_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultLogging2NR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultLogging2NRR16Constraints)
	_ = seq

	// 1. carrierFreq-r16: ref
	{
		if err := ie.CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. measResultListLoggingNR-r16: ref
	{
		if err := ie.MeasResultListLoggingNR_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
