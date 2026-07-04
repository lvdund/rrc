// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultForRSSI-r16 (line 9848).

var measResultForRSSIR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rssi-Result-r16"},
		{Name: "channelOccupancy-r16"},
	},
}

var measResultForRSSIR16ChannelOccupancyR16Constraints = per.Constrained(0, 100)

type MeasResultForRSSI_r16 struct {
	Rssi_Result_r16      RSSI_Range_r16
	ChannelOccupancy_r16 int64
}

func (ie *MeasResultForRSSI_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultForRSSIR16Constraints)
	_ = seq

	// 1. rssi-Result-r16: ref
	{
		if err := ie.Rssi_Result_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. channelOccupancy-r16: integer
	{
		if err := e.EncodeInteger(ie.ChannelOccupancy_r16, measResultForRSSIR16ChannelOccupancyR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultForRSSI_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultForRSSIR16Constraints)
	_ = seq

	// 1. rssi-Result-r16: ref
	{
		if err := ie.Rssi_Result_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. channelOccupancy-r16: integer
	{
		v1, err := d.DecodeInteger(measResultForRSSIR16ChannelOccupancyR16Constraints)
		if err != nil {
			return err
		}
		ie.ChannelOccupancy_r16 = v1
	}

	return nil
}
