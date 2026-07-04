// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-MonitoringOccasions-r16 (line 19797).

var pDCCHMonitoringOccasionsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "period7span3-r16", Optional: true},
		{Name: "period4span3-r16", Optional: true},
		{Name: "period2span2-r16", Optional: true},
	},
}

const (
	PDCCH_MonitoringOccasions_r16_Period7span3_r16_Supported = 0
)

var pDCCHMonitoringOccasionsR16Period7span3R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_MonitoringOccasions_r16_Period7span3_r16_Supported},
}

const (
	PDCCH_MonitoringOccasions_r16_Period4span3_r16_Supported = 0
)

var pDCCHMonitoringOccasionsR16Period4span3R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_MonitoringOccasions_r16_Period4span3_r16_Supported},
}

const (
	PDCCH_MonitoringOccasions_r16_Period2span2_r16_Supported = 0
)

var pDCCHMonitoringOccasionsR16Period2span2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_MonitoringOccasions_r16_Period2span2_r16_Supported},
}

type PDCCH_MonitoringOccasions_r16 struct {
	Period7span3_r16 *int64
	Period4span3_r16 *int64
	Period2span2_r16 *int64
}

func (ie *PDCCH_MonitoringOccasions_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHMonitoringOccasionsR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Period7span3_r16 != nil, ie.Period4span3_r16 != nil, ie.Period2span2_r16 != nil}); err != nil {
		return err
	}

	// 2. period7span3-r16: enumerated
	{
		if ie.Period7span3_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Period7span3_r16, pDCCHMonitoringOccasionsR16Period7span3R16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. period4span3-r16: enumerated
	{
		if ie.Period4span3_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Period4span3_r16, pDCCHMonitoringOccasionsR16Period4span3R16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. period2span2-r16: enumerated
	{
		if ie.Period2span2_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Period2span2_r16, pDCCHMonitoringOccasionsR16Period2span2R16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCCH_MonitoringOccasions_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHMonitoringOccasionsR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. period7span3-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pDCCHMonitoringOccasionsR16Period7span3R16Constraints)
			if err != nil {
				return err
			}
			ie.Period7span3_r16 = &idx
		}
	}

	// 3. period4span3-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pDCCHMonitoringOccasionsR16Period4span3R16Constraints)
			if err != nil {
				return err
			}
			ie.Period4span3_r16 = &idx
		}
	}

	// 4. period2span2-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDCCHMonitoringOccasionsR16Period2span2R16Constraints)
			if err != nil {
				return err
			}
			ie.Period2span2_r16 = &idx
		}
	}

	return nil
}
