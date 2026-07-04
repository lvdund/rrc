// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1900 (line 19952).

var featureSetDownlinkPerCCV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthDL-v1900", Optional: true},
		{Name: "supportedMinBandwidthDL-v1900", Optional: true},
		{Name: "support32-DL-HARQ-ProcessTN-r19", Optional: true},
		{Name: "maxNumberMIMO-LayersPDSCH-v1900", Optional: true},
	},
}

const (
	FeatureSetDownlinkPerCC_v1900_Support32_DL_HARQ_ProcessTN_r19_Supported = 0
)

var featureSetDownlinkPerCCV1900Support32DLHARQProcessTNR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1900_Support32_DL_HARQ_ProcessTN_r19_Supported},
}

const (
	FeatureSetDownlinkPerCC_v1900_MaxNumberMIMO_LayersPDSCH_v1900_Sixlayers = 0
)

var featureSetDownlinkPerCCV1900MaxNumberMIMOLayersPDSCHV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1900_MaxNumberMIMO_LayersPDSCH_v1900_Sixlayers},
}

type FeatureSetDownlinkPerCC_v1900 struct {
	SupportedBandwidthDL_v1900      *SupportedBandwidth_v1900
	SupportedMinBandwidthDL_v1900   *SupportedBandwidth_v1900
	Support32_DL_HARQ_ProcessTN_r19 *int64
	MaxNumberMIMO_LayersPDSCH_v1900 *int64
}

func (ie *FeatureSetDownlinkPerCC_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthDL_v1900 != nil, ie.SupportedMinBandwidthDL_v1900 != nil, ie.Support32_DL_HARQ_ProcessTN_r19 != nil, ie.MaxNumberMIMO_LayersPDSCH_v1900 != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthDL-v1900: ref
	{
		if ie.SupportedBandwidthDL_v1900 != nil {
			if err := ie.SupportedBandwidthDL_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthDL-v1900: ref
	{
		if ie.SupportedMinBandwidthDL_v1900 != nil {
			if err := ie.SupportedMinBandwidthDL_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. support32-DL-HARQ-ProcessTN-r19: enumerated
	{
		if ie.Support32_DL_HARQ_ProcessTN_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Support32_DL_HARQ_ProcessTN_r19, featureSetDownlinkPerCCV1900Support32DLHARQProcessTNR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maxNumberMIMO-LayersPDSCH-v1900: enumerated
	{
		if ie.MaxNumberMIMO_LayersPDSCH_v1900 != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberMIMO_LayersPDSCH_v1900, featureSetDownlinkPerCCV1900MaxNumberMIMOLayersPDSCHV1900Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthDL-v1900: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandwidthDL_v1900 = new(SupportedBandwidth_v1900)
			if err := ie.SupportedBandwidthDL_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthDL-v1900: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedMinBandwidthDL_v1900 = new(SupportedBandwidth_v1900)
			if err := ie.SupportedMinBandwidthDL_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. support32-DL-HARQ-ProcessTN-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1900Support32DLHARQProcessTNR19Constraints)
			if err != nil {
				return err
			}
			ie.Support32_DL_HARQ_ProcessTN_r19 = &idx
		}
	}

	// 5. maxNumberMIMO-LayersPDSCH-v1900: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1900MaxNumberMIMOLayersPDSCHV1900Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberMIMO_LayersPDSCH_v1900 = &idx
		}
	}

	return nil
}
