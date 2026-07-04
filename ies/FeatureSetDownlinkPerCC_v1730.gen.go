// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlinkPerCC-v1730 (line 19916).

var featureSetDownlinkPerCCV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraSlotTDM-UnicastGroupCommonPDSCH-r17", Optional: true},
		{Name: "sps-MulticastSCell-r17", Optional: true},
		{Name: "sps-MulticastSCellMultiConfig-r17", Optional: true},
		{Name: "dci-BroadcastWith16Repetitions-r17", Optional: true},
	},
}

const (
	FeatureSetDownlinkPerCC_v1730_IntraSlotTDM_UnicastGroupCommonPDSCH_r17_Yes = 0
	FeatureSetDownlinkPerCC_v1730_IntraSlotTDM_UnicastGroupCommonPDSCH_r17_No  = 1
)

var featureSetDownlinkPerCCV1730IntraSlotTDMUnicastGroupCommonPDSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1730_IntraSlotTDM_UnicastGroupCommonPDSCH_r17_Yes, FeatureSetDownlinkPerCC_v1730_IntraSlotTDM_UnicastGroupCommonPDSCH_r17_No},
}

const (
	FeatureSetDownlinkPerCC_v1730_Sps_MulticastSCell_r17_Supported = 0
)

var featureSetDownlinkPerCCV1730SpsMulticastSCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1730_Sps_MulticastSCell_r17_Supported},
}

var featureSetDownlinkPerCCV1730SpsMulticastSCellMultiConfigR17Constraints = per.Constrained(1, 8)

const (
	FeatureSetDownlinkPerCC_v1730_Dci_BroadcastWith16Repetitions_r17_Supported = 0
)

var featureSetDownlinkPerCCV1730DciBroadcastWith16RepetitionsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlinkPerCC_v1730_Dci_BroadcastWith16Repetitions_r17_Supported},
}

type FeatureSetDownlinkPerCC_v1730 struct {
	IntraSlotTDM_UnicastGroupCommonPDSCH_r17 *int64
	Sps_MulticastSCell_r17                   *int64
	Sps_MulticastSCellMultiConfig_r17        *int64
	Dci_BroadcastWith16Repetitions_r17       *int64
}

func (ie *FeatureSetDownlinkPerCC_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkPerCCV1730Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17 != nil, ie.Sps_MulticastSCell_r17 != nil, ie.Sps_MulticastSCellMultiConfig_r17 != nil, ie.Dci_BroadcastWith16Repetitions_r17 != nil}); err != nil {
		return err
	}

	// 2. intraSlotTDM-UnicastGroupCommonPDSCH-r17: enumerated
	{
		if ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17, featureSetDownlinkPerCCV1730IntraSlotTDMUnicastGroupCommonPDSCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sps-MulticastSCell-r17: enumerated
	{
		if ie.Sps_MulticastSCell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sps_MulticastSCell_r17, featureSetDownlinkPerCCV1730SpsMulticastSCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sps-MulticastSCellMultiConfig-r17: integer
	{
		if ie.Sps_MulticastSCellMultiConfig_r17 != nil {
			if err := e.EncodeInteger(*ie.Sps_MulticastSCellMultiConfig_r17, featureSetDownlinkPerCCV1730SpsMulticastSCellMultiConfigR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. dci-BroadcastWith16Repetitions-r17: enumerated
	{
		if ie.Dci_BroadcastWith16Repetitions_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dci_BroadcastWith16Repetitions_r17, featureSetDownlinkPerCCV1730DciBroadcastWith16RepetitionsR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkPerCCV1730Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intraSlotTDM-UnicastGroupCommonPDSCH-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1730IntraSlotTDMUnicastGroupCommonPDSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.IntraSlotTDM_UnicastGroupCommonPDSCH_r17 = &idx
		}
	}

	// 3. sps-MulticastSCell-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1730SpsMulticastSCellR17Constraints)
			if err != nil {
				return err
			}
			ie.Sps_MulticastSCell_r17 = &idx
		}
	}

	// 4. sps-MulticastSCellMultiConfig-r17: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(featureSetDownlinkPerCCV1730SpsMulticastSCellMultiConfigR17Constraints)
			if err != nil {
				return err
			}
			ie.Sps_MulticastSCellMultiConfig_r17 = &v
		}
	}

	// 5. dci-BroadcastWith16Repetitions-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkPerCCV1730DciBroadcastWith16RepetitionsR17Constraints)
			if err != nil {
				return err
			}
			ie.Dci_BroadcastWith16Repetitions_r17 = &idx
		}
	}

	return nil
}
