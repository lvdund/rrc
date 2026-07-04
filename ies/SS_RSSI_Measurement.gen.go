// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SS-RSSI-Measurement (line 15932).

var sSRSSIMeasurementConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measurementSlots"},
		{Name: "endSymbol"},
	},
}

var sSRSSIMeasurementMeasurementSlotsConstraints = per.SizeRange(1, 80)

var sSRSSIMeasurementEndSymbolConstraints = per.Constrained(0, 3)

type SS_RSSI_Measurement struct {
	MeasurementSlots per.BitString
	EndSymbol        int64
}

func (ie *SS_RSSI_Measurement) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSRSSIMeasurementConstraints)
	_ = seq

	// 1. measurementSlots: bit-string
	{
		if err := e.EncodeBitString(ie.MeasurementSlots, sSRSSIMeasurementMeasurementSlotsConstraints); err != nil {
			return err
		}
	}

	// 2. endSymbol: integer
	{
		if err := e.EncodeInteger(ie.EndSymbol, sSRSSIMeasurementEndSymbolConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SS_RSSI_Measurement) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSRSSIMeasurementConstraints)
	_ = seq

	// 1. measurementSlots: bit-string
	{
		v0, err := d.DecodeBitString(sSRSSIMeasurementMeasurementSlotsConstraints)
		if err != nil {
			return err
		}
		ie.MeasurementSlots = v0
	}

	// 2. endSymbol: integer
	{
		v1, err := d.DecodeInteger(sSRSSIMeasurementEndSymbolConstraints)
		if err != nil {
			return err
		}
		ie.EndSymbol = v1
	}

	return nil
}
