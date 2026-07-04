// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReestablishmentComplete-v1800-IEs (line 931).

var rRCReestablishmentCompleteV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "flightPathInfoAvailable-r18", Optional: true},
		{Name: "measConfigReportAppLayerAvailable-r18", Optional: true},
		{Name: "musim-CapRestrictionInd-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCReestablishmentComplete_v1800_IEs_FlightPathInfoAvailable_r18_True = 0
)

var rRCReestablishmentCompleteV1800IEsFlightPathInfoAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReestablishmentComplete_v1800_IEs_FlightPathInfoAvailable_r18_True},
}

const (
	RRCReestablishmentComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True = 0
)

var rRCReestablishmentCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReestablishmentComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True},
}

const (
	RRCReestablishmentComplete_v1800_IEs_Musim_CapRestrictionInd_r18_True = 0
)

var rRCReestablishmentCompleteV1800IEsMusimCapRestrictionIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReestablishmentComplete_v1800_IEs_Musim_CapRestrictionInd_r18_True},
}

type RRCReestablishmentComplete_v1800_IEs struct {
	FlightPathInfoAvailable_r18           *int64
	MeasConfigReportAppLayerAvailable_r18 *int64
	Musim_CapRestrictionInd_r18           *int64
	NonCriticalExtension                  *struct{}
}

func (ie *RRCReestablishmentComplete_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestablishmentCompleteV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FlightPathInfoAvailable_r18 != nil, ie.MeasConfigReportAppLayerAvailable_r18 != nil, ie.Musim_CapRestrictionInd_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. flightPathInfoAvailable-r18: enumerated
	{
		if ie.FlightPathInfoAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FlightPathInfoAvailable_r18, rRCReestablishmentCompleteV1800IEsFlightPathInfoAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if ie.MeasConfigReportAppLayerAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MeasConfigReportAppLayerAvailable_r18, rRCReestablishmentCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. musim-CapRestrictionInd-r18: enumerated
	{
		if ie.Musim_CapRestrictionInd_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_CapRestrictionInd_r18, rRCReestablishmentCompleteV1800IEsMusimCapRestrictionIndR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *RRCReestablishmentComplete_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestablishmentCompleteV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. flightPathInfoAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCReestablishmentCompleteV1800IEsFlightPathInfoAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.FlightPathInfoAvailable_r18 = &idx
		}
	}

	// 3. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCReestablishmentCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.MeasConfigReportAppLayerAvailable_r18 = &idx
		}
	}

	// 4. musim-CapRestrictionInd-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rRCReestablishmentCompleteV1800IEsMusimCapRestrictionIndR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_CapRestrictionInd_r18 = &idx
		}
	}

	// 5. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
