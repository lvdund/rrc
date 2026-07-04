// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AvailableRAN-VisibleMetrics-r18 (line 26001).

var availableRANVisibleMetricsR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "appLayerBufferLevelList-r18", Optional: true},
		{Name: "playoutDelayForMediaStartup-r18", Optional: true},
	},
}

const (
	AvailableRAN_VisibleMetrics_r18_AppLayerBufferLevelList_r18_True = 0
)

var availableRANVisibleMetricsR18AppLayerBufferLevelListR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AvailableRAN_VisibleMetrics_r18_AppLayerBufferLevelList_r18_True},
}

const (
	AvailableRAN_VisibleMetrics_r18_PlayoutDelayForMediaStartup_r18_True = 0
)

var availableRANVisibleMetricsR18PlayoutDelayForMediaStartupR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AvailableRAN_VisibleMetrics_r18_PlayoutDelayForMediaStartup_r18_True},
}

type AvailableRAN_VisibleMetrics_r18 struct {
	AppLayerBufferLevelList_r18     *int64
	PlayoutDelayForMediaStartup_r18 *int64
}

func (ie *AvailableRAN_VisibleMetrics_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(availableRANVisibleMetricsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AppLayerBufferLevelList_r18 != nil, ie.PlayoutDelayForMediaStartup_r18 != nil}); err != nil {
		return err
	}

	// 3. appLayerBufferLevelList-r18: enumerated
	{
		if ie.AppLayerBufferLevelList_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AppLayerBufferLevelList_r18, availableRANVisibleMetricsR18AppLayerBufferLevelListR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. playoutDelayForMediaStartup-r18: enumerated
	{
		if ie.PlayoutDelayForMediaStartup_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PlayoutDelayForMediaStartup_r18, availableRANVisibleMetricsR18PlayoutDelayForMediaStartupR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *AvailableRAN_VisibleMetrics_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(availableRANVisibleMetricsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. appLayerBufferLevelList-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(availableRANVisibleMetricsR18AppLayerBufferLevelListR18Constraints)
			if err != nil {
				return err
			}
			ie.AppLayerBufferLevelList_r18 = &idx
		}
	}

	// 4. playoutDelayForMediaStartup-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(availableRANVisibleMetricsR18PlayoutDelayForMediaStartupR18Constraints)
			if err != nil {
				return err
			}
			ie.PlayoutDelayForMediaStartup_r18 = &idx
		}
	}

	return nil
}
