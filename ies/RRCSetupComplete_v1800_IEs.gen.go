// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupComplete-v1800-IEs (line 1749).

var rRCSetupCompleteV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ncr-NodeIndication-r18", Optional: true},
		{Name: "musim-CapRestrictionInd-r18", Optional: true},
		{Name: "flightPathInfoAvailable-r18", Optional: true},
		{Name: "measConfigReportAppLayerAvailable-r18", Optional: true},
		{Name: "mobileIAB-NodeIndication-r18", Optional: true},
		{Name: "reselectionMeasAvailable-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCSetupComplete_v1800_IEs_Ncr_NodeIndication_r18_True = 0
)

var rRCSetupCompleteV1800IEsNcrNodeIndicationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1800_IEs_Ncr_NodeIndication_r18_True},
}

const (
	RRCSetupComplete_v1800_IEs_Musim_CapRestrictionInd_r18_True = 0
)

var rRCSetupCompleteV1800IEsMusimCapRestrictionIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1800_IEs_Musim_CapRestrictionInd_r18_True},
}

const (
	RRCSetupComplete_v1800_IEs_FlightPathInfoAvailable_r18_True = 0
)

var rRCSetupCompleteV1800IEsFlightPathInfoAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1800_IEs_FlightPathInfoAvailable_r18_True},
}

const (
	RRCSetupComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True = 0
)

var rRCSetupCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True},
}

const (
	RRCSetupComplete_v1800_IEs_MobileIAB_NodeIndication_r18_True = 0
)

var rRCSetupCompleteV1800IEsMobileIABNodeIndicationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1800_IEs_MobileIAB_NodeIndication_r18_True},
}

const (
	RRCSetupComplete_v1800_IEs_ReselectionMeasAvailable_r18_True = 0
)

var rRCSetupCompleteV1800IEsReselectionMeasAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1800_IEs_ReselectionMeasAvailable_r18_True},
}

type RRCSetupComplete_v1800_IEs struct {
	Ncr_NodeIndication_r18                *int64
	Musim_CapRestrictionInd_r18           *int64
	FlightPathInfoAvailable_r18           *int64
	MeasConfigReportAppLayerAvailable_r18 *int64
	MobileIAB_NodeIndication_r18          *int64
	ReselectionMeasAvailable_r18          *int64
	NonCriticalExtension                  *struct{}
}

func (ie *RRCSetupComplete_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ncr_NodeIndication_r18 != nil, ie.Musim_CapRestrictionInd_r18 != nil, ie.FlightPathInfoAvailable_r18 != nil, ie.MeasConfigReportAppLayerAvailable_r18 != nil, ie.MobileIAB_NodeIndication_r18 != nil, ie.ReselectionMeasAvailable_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ncr-NodeIndication-r18: enumerated
	{
		if ie.Ncr_NodeIndication_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ncr_NodeIndication_r18, rRCSetupCompleteV1800IEsNcrNodeIndicationR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. musim-CapRestrictionInd-r18: enumerated
	{
		if ie.Musim_CapRestrictionInd_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_CapRestrictionInd_r18, rRCSetupCompleteV1800IEsMusimCapRestrictionIndR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. flightPathInfoAvailable-r18: enumerated
	{
		if ie.FlightPathInfoAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FlightPathInfoAvailable_r18, rRCSetupCompleteV1800IEsFlightPathInfoAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if ie.MeasConfigReportAppLayerAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MeasConfigReportAppLayerAvailable_r18, rRCSetupCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. mobileIAB-NodeIndication-r18: enumerated
	{
		if ie.MobileIAB_NodeIndication_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MobileIAB_NodeIndication_r18, rRCSetupCompleteV1800IEsMobileIABNodeIndicationR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. reselectionMeasAvailable-r18: enumerated
	{
		if ie.ReselectionMeasAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ReselectionMeasAvailable_r18, rRCSetupCompleteV1800IEsReselectionMeasAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *RRCSetupComplete_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ncr-NodeIndication-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1800IEsNcrNodeIndicationR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_NodeIndication_r18 = &idx
		}
	}

	// 3. musim-CapRestrictionInd-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1800IEsMusimCapRestrictionIndR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_CapRestrictionInd_r18 = &idx
		}
	}

	// 4. flightPathInfoAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1800IEsFlightPathInfoAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.FlightPathInfoAvailable_r18 = &idx
		}
	}

	// 5. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.MeasConfigReportAppLayerAvailable_r18 = &idx
		}
	}

	// 6. mobileIAB-NodeIndication-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1800IEsMobileIABNodeIndicationR18Constraints)
			if err != nil {
				return err
			}
			ie.MobileIAB_NodeIndication_r18 = &idx
		}
	}

	// 7. reselectionMeasAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1800IEsReselectionMeasAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.ReselectionMeasAvailable_r18 = &idx
		}
	}

	// 8. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(6) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
