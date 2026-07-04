// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC (line 20544).

var featureSetUplinkPerCCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedSubcarrierSpacingUL"},
		{Name: "supportedBandwidthUL"},
		{Name: "channelBW-90mhz", Optional: true},
		{Name: "mimo-CB-PUSCH", Optional: true},
		{Name: "maxNumberMIMO-LayersNonCB-PUSCH", Optional: true},
		{Name: "supportedModulationOrderUL", Optional: true},
	},
}

const (
	FeatureSetUplinkPerCC_ChannelBW_90mhz_Supported = 0
)

var featureSetUplinkPerCCChannelBW90mhzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_ChannelBW_90mhz_Supported},
}

var featureSetUplinkPerCCMimoCBPUSCHConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberMIMO-LayersCB-PUSCH", Optional: true},
		{Name: "maxNumberSRS-ResourcePerSet"},
	},
}

type FeatureSetUplinkPerCC struct {
	SupportedSubcarrierSpacingUL SubcarrierSpacing
	SupportedBandwidthUL         SupportedBandwidth
	ChannelBW_90mhz              *int64
	Mimo_CB_PUSCH                *struct {
		MaxNumberMIMO_LayersCB_PUSCH *MIMO_LayersUL
		MaxNumberSRS_ResourcePerSet  int64
	}
	MaxNumberMIMO_LayersNonCB_PUSCH *MIMO_LayersUL
	SupportedModulationOrderUL      *ModulationOrder
}

func (ie *FeatureSetUplinkPerCC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ChannelBW_90mhz != nil, ie.Mimo_CB_PUSCH != nil, ie.MaxNumberMIMO_LayersNonCB_PUSCH != nil, ie.SupportedModulationOrderUL != nil}); err != nil {
		return err
	}

	// 2. supportedSubcarrierSpacingUL: ref
	{
		if err := ie.SupportedSubcarrierSpacingUL.Encode(e); err != nil {
			return err
		}
	}

	// 3. supportedBandwidthUL: ref
	{
		if err := ie.SupportedBandwidthUL.Encode(e); err != nil {
			return err
		}
	}

	// 4. channelBW-90mhz: enumerated
	{
		if ie.ChannelBW_90mhz != nil {
			if err := e.EncodeEnumerated(*ie.ChannelBW_90mhz, featureSetUplinkPerCCChannelBW90mhzConstraints); err != nil {
				return err
			}
		}
	}

	// 5. mimo-CB-PUSCH: sequence
	{
		if ie.Mimo_CB_PUSCH != nil {
			c := ie.Mimo_CB_PUSCH
			featureSetUplinkPerCCMimoCBPUSCHSeq := e.NewSequenceEncoder(featureSetUplinkPerCCMimoCBPUSCHConstraints)
			if err := featureSetUplinkPerCCMimoCBPUSCHSeq.EncodePreamble([]bool{c.MaxNumberMIMO_LayersCB_PUSCH != nil}); err != nil {
				return err
			}
			if c.MaxNumberMIMO_LayersCB_PUSCH != nil {
				if err := c.MaxNumberMIMO_LayersCB_PUSCH.Encode(e); err != nil {
					return err
				}
			}
			if err := e.EncodeInteger(c.MaxNumberSRS_ResourcePerSet, per.Constrained(1, 2)); err != nil {
				return err
			}
		}
	}

	// 6. maxNumberMIMO-LayersNonCB-PUSCH: ref
	{
		if ie.MaxNumberMIMO_LayersNonCB_PUSCH != nil {
			if err := ie.MaxNumberMIMO_LayersNonCB_PUSCH.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. supportedModulationOrderUL: ref
	{
		if ie.SupportedModulationOrderUL != nil {
			if err := ie.SupportedModulationOrderUL.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedSubcarrierSpacingUL: ref
	{
		if err := ie.SupportedSubcarrierSpacingUL.Decode(d); err != nil {
			return err
		}
	}

	// 3. supportedBandwidthUL: ref
	{
		if err := ie.SupportedBandwidthUL.Decode(d); err != nil {
			return err
		}
	}

	// 4. channelBW-90mhz: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCChannelBW90mhzConstraints)
			if err != nil {
				return err
			}
			ie.ChannelBW_90mhz = &idx
		}
	}

	// 5. mimo-CB-PUSCH: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.Mimo_CB_PUSCH = &struct {
				MaxNumberMIMO_LayersCB_PUSCH *MIMO_LayersUL
				MaxNumberSRS_ResourcePerSet  int64
			}{}
			c := ie.Mimo_CB_PUSCH
			featureSetUplinkPerCCMimoCBPUSCHSeq := d.NewSequenceDecoder(featureSetUplinkPerCCMimoCBPUSCHConstraints)
			if err := featureSetUplinkPerCCMimoCBPUSCHSeq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkPerCCMimoCBPUSCHSeq.IsComponentPresent(0) {
				c.MaxNumberMIMO_LayersCB_PUSCH = new(MIMO_LayersUL)
				if err := (*c.MaxNumberMIMO_LayersCB_PUSCH).Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet = v
			}
		}
	}

	// 6. maxNumberMIMO-LayersNonCB-PUSCH: ref
	{
		if seq.IsComponentPresent(4) {
			ie.MaxNumberMIMO_LayersNonCB_PUSCH = new(MIMO_LayersUL)
			if err := ie.MaxNumberMIMO_LayersNonCB_PUSCH.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. supportedModulationOrderUL: ref
	{
		if seq.IsComponentPresent(5) {
			ie.SupportedModulationOrderUL = new(ModulationOrder)
			if err := ie.SupportedModulationOrderUL.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
