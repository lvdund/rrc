// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasRSSI-ReportConfig-r16 (line 13871).

var measRSSIReportConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "channelOccupancyThreshold-r16", Optional: true},
	},
}

type MeasRSSI_ReportConfig_r16 struct {
	ChannelOccupancyThreshold_r16 *RSSI_Range_r16
}

func (ie *MeasRSSI_ReportConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measRSSIReportConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ChannelOccupancyThreshold_r16 != nil}); err != nil {
		return err
	}

	// 2. channelOccupancyThreshold-r16: ref
	{
		if ie.ChannelOccupancyThreshold_r16 != nil {
			if err := ie.ChannelOccupancyThreshold_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasRSSI_ReportConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measRSSIReportConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. channelOccupancyThreshold-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ChannelOccupancyThreshold_r16 = new(RSSI_Range_r16)
			if err := ie.ChannelOccupancyThreshold_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
