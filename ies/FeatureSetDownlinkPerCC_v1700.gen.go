// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1700 (line 19895).

var featureSetDownlinkPerCCV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedMinBandwidthDL-r17", Optional: true},
		{Name: "broadcastSCell-r17", Optional: true},
		{Name: "maxNumberMIMO-LayersMulticastPDSCH-r17", Optional: true},
		{Name: "dynamicMulticastSCell-r17", Optional: true},
		{Name: "supportedBandwidthDL-v1710", Optional: true},
		{Name: "supportedCRS-InterfMitigation-r17", Optional: true},
	},
}

const (
	FeatureSetDownlinkPerCC_v1700_BroadcastSCell_r17_Supported = 0
)

var featureSetDownlinkPerCCV1700BroadcastSCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1700_BroadcastSCell_r17_Supported},
}

const (
	FeatureSetDownlinkPerCC_v1700_MaxNumberMIMO_LayersMulticastPDSCH_r17_N2 = 0
	FeatureSetDownlinkPerCC_v1700_MaxNumberMIMO_LayersMulticastPDSCH_r17_N4 = 1
	FeatureSetDownlinkPerCC_v1700_MaxNumberMIMO_LayersMulticastPDSCH_r17_N8 = 2
)

var featureSetDownlinkPerCCV1700MaxNumberMIMOLayersMulticastPDSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1700_MaxNumberMIMO_LayersMulticastPDSCH_r17_N2, FeatureSetDownlinkPerCC_v1700_MaxNumberMIMO_LayersMulticastPDSCH_r17_N4, FeatureSetDownlinkPerCC_v1700_MaxNumberMIMO_LayersMulticastPDSCH_r17_N8},
}

const (
	FeatureSetDownlinkPerCC_v1700_DynamicMulticastSCell_r17_Supported = 0
)

var featureSetDownlinkPerCCV1700DynamicMulticastSCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1700_DynamicMulticastSCell_r17_Supported},
}

type FeatureSetDownlinkPerCC_v1700 struct {
	SupportedMinBandwidthDL_r17            *SupportedBandwidth_v1700
	BroadcastSCell_r17                     *int64
	MaxNumberMIMO_LayersMulticastPDSCH_r17 *int64
	DynamicMulticastSCell_r17              *int64
	SupportedBandwidthDL_v1710             *SupportedBandwidth_v1700
	SupportedCRS_InterfMitigation_r17      *CRS_InterfMitigation_r17
}

func (ie *FeatureSetDownlinkPerCC_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedMinBandwidthDL_r17 != nil, ie.BroadcastSCell_r17 != nil, ie.MaxNumberMIMO_LayersMulticastPDSCH_r17 != nil, ie.DynamicMulticastSCell_r17 != nil, ie.SupportedBandwidthDL_v1710 != nil, ie.SupportedCRS_InterfMitigation_r17 != nil}); err != nil {
		return err
	}

	// 2. supportedMinBandwidthDL-r17: ref
	{
		if ie.SupportedMinBandwidthDL_r17 != nil {
			if err := ie.SupportedMinBandwidthDL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. broadcastSCell-r17: enumerated
	{
		if ie.BroadcastSCell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.BroadcastSCell_r17, featureSetDownlinkPerCCV1700BroadcastSCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maxNumberMIMO-LayersMulticastPDSCH-r17: enumerated
	{
		if ie.MaxNumberMIMO_LayersMulticastPDSCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberMIMO_LayersMulticastPDSCH_r17, featureSetDownlinkPerCCV1700MaxNumberMIMOLayersMulticastPDSCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. dynamicMulticastSCell-r17: enumerated
	{
		if ie.DynamicMulticastSCell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.DynamicMulticastSCell_r17, featureSetDownlinkPerCCV1700DynamicMulticastSCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. supportedBandwidthDL-v1710: ref
	{
		if ie.SupportedBandwidthDL_v1710 != nil {
			if err := ie.SupportedBandwidthDL_v1710.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. supportedCRS-InterfMitigation-r17: ref
	{
		if ie.SupportedCRS_InterfMitigation_r17 != nil {
			if err := ie.SupportedCRS_InterfMitigation_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedMinBandwidthDL-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedMinBandwidthDL_r17 = new(SupportedBandwidth_v1700)
			if err := ie.SupportedMinBandwidthDL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. broadcastSCell-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1700BroadcastSCellR17Constraints)
			if err != nil {
				return err
			}
			ie.BroadcastSCell_r17 = &idx
		}
	}

	// 4. maxNumberMIMO-LayersMulticastPDSCH-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1700MaxNumberMIMOLayersMulticastPDSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberMIMO_LayersMulticastPDSCH_r17 = &idx
		}
	}

	// 5. dynamicMulticastSCell-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1700DynamicMulticastSCellR17Constraints)
			if err != nil {
				return err
			}
			ie.DynamicMulticastSCell_r17 = &idx
		}
	}

	// 6. supportedBandwidthDL-v1710: ref
	{
		if seq.IsComponentPresent(4) {
			ie.SupportedBandwidthDL_v1710 = new(SupportedBandwidth_v1700)
			if err := ie.SupportedBandwidthDL_v1710.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. supportedCRS-InterfMitigation-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.SupportedCRS_InterfMitigation_r17 = new(CRS_InterfMitigation_r17)
			if err := ie.SupportedCRS_InterfMitigation_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
