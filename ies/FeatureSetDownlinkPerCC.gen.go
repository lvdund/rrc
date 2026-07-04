// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC (line 19880).

var featureSetDownlinkPerCCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedSubcarrierSpacingDL"},
		{Name: "supportedBandwidthDL"},
		{Name: "channelBW-90mhz", Optional: true},
		{Name: "maxNumberMIMO-LayersPDSCH", Optional: true},
		{Name: "supportedModulationOrderDL", Optional: true},
	},
}

const (
	FeatureSetDownlinkPerCC_ChannelBW_90mhz_Supported = 0
)

var featureSetDownlinkPerCCChannelBW90mhzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_ChannelBW_90mhz_Supported},
}

type FeatureSetDownlinkPerCC struct {
	SupportedSubcarrierSpacingDL SubcarrierSpacing
	SupportedBandwidthDL         SupportedBandwidth
	ChannelBW_90mhz              *int64
	MaxNumberMIMO_LayersPDSCH    *MIMO_LayersDL
	SupportedModulationOrderDL   *ModulationOrder
}

func (ie *FeatureSetDownlinkPerCC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ChannelBW_90mhz != nil, ie.MaxNumberMIMO_LayersPDSCH != nil, ie.SupportedModulationOrderDL != nil}); err != nil {
		return err
	}

	// 2. supportedSubcarrierSpacingDL: ref
	{
		if err := ie.SupportedSubcarrierSpacingDL.Encode(e); err != nil {
			return err
		}
	}

	// 3. supportedBandwidthDL: ref
	{
		if err := ie.SupportedBandwidthDL.Encode(e); err != nil {
			return err
		}
	}

	// 4. channelBW-90mhz: enumerated
	{
		if ie.ChannelBW_90mhz != nil {
			if err := e.EncodeEnumerated(*ie.ChannelBW_90mhz, featureSetDownlinkPerCCChannelBW90mhzConstraints); err != nil {
				return err
			}
		}
	}

	// 5. maxNumberMIMO-LayersPDSCH: ref
	{
		if ie.MaxNumberMIMO_LayersPDSCH != nil {
			if err := ie.MaxNumberMIMO_LayersPDSCH.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. supportedModulationOrderDL: ref
	{
		if ie.SupportedModulationOrderDL != nil {
			if err := ie.SupportedModulationOrderDL.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedSubcarrierSpacingDL: ref
	{
		if err := ie.SupportedSubcarrierSpacingDL.Decode(d); err != nil {
			return err
		}
	}

	// 3. supportedBandwidthDL: ref
	{
		if err := ie.SupportedBandwidthDL.Decode(d); err != nil {
			return err
		}
	}

	// 4. channelBW-90mhz: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCChannelBW90mhzConstraints)
			if err != nil {
				return err
			}
			ie.ChannelBW_90mhz = &idx
		}
	}

	// 5. maxNumberMIMO-LayersPDSCH: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MaxNumberMIMO_LayersPDSCH = new(MIMO_LayersDL)
			if err := ie.MaxNumberMIMO_LayersPDSCH.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. supportedModulationOrderDL: ref
	{
		if seq.IsComponentPresent(4) {
			ie.SupportedModulationOrderDL = new(ModulationOrder)
			if err := ie.SupportedModulationOrderDL.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
