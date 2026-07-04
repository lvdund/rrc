// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1720 (line 19907).

var featureSetDownlinkPerCCV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxModulationOrderForMulticastDataRateCalculation-r17", Optional: true},
		{Name: "fdm-BroadcastUnicast-r17", Optional: true},
		{Name: "fdm-MulticastUnicast-r17", Optional: true},
	},
}

const (
	FeatureSetDownlinkPerCC_v1720_MaxModulationOrderForMulticastDataRateCalculation_r17_Qam64   = 0
	FeatureSetDownlinkPerCC_v1720_MaxModulationOrderForMulticastDataRateCalculation_r17_Qam256  = 1
	FeatureSetDownlinkPerCC_v1720_MaxModulationOrderForMulticastDataRateCalculation_r17_Qam1024 = 2
)

var featureSetDownlinkPerCCV1720MaxModulationOrderForMulticastDataRateCalculationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1720_MaxModulationOrderForMulticastDataRateCalculation_r17_Qam64, FeatureSetDownlinkPerCC_v1720_MaxModulationOrderForMulticastDataRateCalculation_r17_Qam256, FeatureSetDownlinkPerCC_v1720_MaxModulationOrderForMulticastDataRateCalculation_r17_Qam1024},
}

const (
	FeatureSetDownlinkPerCC_v1720_Fdm_BroadcastUnicast_r17_Supported = 0
)

var featureSetDownlinkPerCCV1720FdmBroadcastUnicastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1720_Fdm_BroadcastUnicast_r17_Supported},
}

const (
	FeatureSetDownlinkPerCC_v1720_Fdm_MulticastUnicast_r17_Supported = 0
)

var featureSetDownlinkPerCCV1720FdmMulticastUnicastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1720_Fdm_MulticastUnicast_r17_Supported},
}

type FeatureSetDownlinkPerCC_v1720 struct {
	MaxModulationOrderForMulticastDataRateCalculation_r17 *int64
	Fdm_BroadcastUnicast_r17                              *int64
	Fdm_MulticastUnicast_r17                              *int64
}

func (ie *FeatureSetDownlinkPerCC_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxModulationOrderForMulticastDataRateCalculation_r17 != nil, ie.Fdm_BroadcastUnicast_r17 != nil, ie.Fdm_MulticastUnicast_r17 != nil}); err != nil {
		return err
	}

	// 2. maxModulationOrderForMulticastDataRateCalculation-r17: enumerated
	{
		if ie.MaxModulationOrderForMulticastDataRateCalculation_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MaxModulationOrderForMulticastDataRateCalculation_r17, featureSetDownlinkPerCCV1720MaxModulationOrderForMulticastDataRateCalculationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. fdm-BroadcastUnicast-r17: enumerated
	{
		if ie.Fdm_BroadcastUnicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Fdm_BroadcastUnicast_r17, featureSetDownlinkPerCCV1720FdmBroadcastUnicastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. fdm-MulticastUnicast-r17: enumerated
	{
		if ie.Fdm_MulticastUnicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Fdm_MulticastUnicast_r17, featureSetDownlinkPerCCV1720FdmMulticastUnicastR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxModulationOrderForMulticastDataRateCalculation-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1720MaxModulationOrderForMulticastDataRateCalculationR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxModulationOrderForMulticastDataRateCalculation_r17 = &idx
		}
	}

	// 3. fdm-BroadcastUnicast-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1720FdmBroadcastUnicastR17Constraints)
			if err != nil {
				return err
			}
			ie.Fdm_BroadcastUnicast_r17 = &idx
		}
	}

	// 4. fdm-MulticastUnicast-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1720FdmMulticastUnicastR17Constraints)
			if err != nil {
				return err
			}
			ie.Fdm_MulticastUnicast_r17 = &idx
		}
	}

	return nil
}
