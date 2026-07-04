// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeComplete-v1800-IEs (line 1636).

var rRCResumeCompleteV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "needForInterruptionInfoNR-r18", Optional: true},
		{Name: "musim-CapRestrictionInd-r18", Optional: true},
		{Name: "flightPathInfoAvailable-r18", Optional: true},
		{Name: "measConfigReportAppLayerAvailable-r18", Optional: true},
		{Name: "measResultReselectionNR-r18", Optional: true},
		{Name: "reselectionMeasAvailable-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCResumeComplete_v1800_IEs_Musim_CapRestrictionInd_r18_True = 0
)

var rRCResumeCompleteV1800IEsMusimCapRestrictionIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResumeComplete_v1800_IEs_Musim_CapRestrictionInd_r18_True},
}

const (
	RRCResumeComplete_v1800_IEs_FlightPathInfoAvailable_r18_True = 0
)

var rRCResumeCompleteV1800IEsFlightPathInfoAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResumeComplete_v1800_IEs_FlightPathInfoAvailable_r18_True},
}

const (
	RRCResumeComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True = 0
)

var rRCResumeCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResumeComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True},
}

const (
	RRCResumeComplete_v1800_IEs_ReselectionMeasAvailable_r18_True = 0
)

var rRCResumeCompleteV1800IEsReselectionMeasAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResumeComplete_v1800_IEs_ReselectionMeasAvailable_r18_True},
}

type RRCResumeComplete_v1800_IEs struct {
	NeedForInterruptionInfoNR_r18         *NeedForInterruptionInfoNR_r18
	Musim_CapRestrictionInd_r18           *int64
	FlightPathInfoAvailable_r18           *int64
	MeasConfigReportAppLayerAvailable_r18 *int64
	MeasResultReselectionNR_r18           *MeasResultIdleNR_r16
	ReselectionMeasAvailable_r18          *int64
	NonCriticalExtension                  *RRCResumeComplete_v1900_IEs
}

func (ie *RRCResumeComplete_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeCompleteV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NeedForInterruptionInfoNR_r18 != nil, ie.Musim_CapRestrictionInd_r18 != nil, ie.FlightPathInfoAvailable_r18 != nil, ie.MeasConfigReportAppLayerAvailable_r18 != nil, ie.MeasResultReselectionNR_r18 != nil, ie.ReselectionMeasAvailable_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. needForInterruptionInfoNR-r18: ref
	{
		if ie.NeedForInterruptionInfoNR_r18 != nil {
			if err := ie.NeedForInterruptionInfoNR_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. musim-CapRestrictionInd-r18: enumerated
	{
		if ie.Musim_CapRestrictionInd_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_CapRestrictionInd_r18, rRCResumeCompleteV1800IEsMusimCapRestrictionIndR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. flightPathInfoAvailable-r18: enumerated
	{
		if ie.FlightPathInfoAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FlightPathInfoAvailable_r18, rRCResumeCompleteV1800IEsFlightPathInfoAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if ie.MeasConfigReportAppLayerAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MeasConfigReportAppLayerAvailable_r18, rRCResumeCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. measResultReselectionNR-r18: ref
	{
		if ie.MeasResultReselectionNR_r18 != nil {
			if err := ie.MeasResultReselectionNR_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. reselectionMeasAvailable-r18: enumerated
	{
		if ie.ReselectionMeasAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ReselectionMeasAvailable_r18, rRCResumeCompleteV1800IEsReselectionMeasAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCResumeComplete_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeCompleteV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. needForInterruptionInfoNR-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.NeedForInterruptionInfoNR_r18 = new(NeedForInterruptionInfoNR_r18)
			if err := ie.NeedForInterruptionInfoNR_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. musim-CapRestrictionInd-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCResumeCompleteV1800IEsMusimCapRestrictionIndR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_CapRestrictionInd_r18 = &idx
		}
	}

	// 4. flightPathInfoAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rRCResumeCompleteV1800IEsFlightPathInfoAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.FlightPathInfoAvailable_r18 = &idx
		}
	}

	// 5. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(rRCResumeCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.MeasConfigReportAppLayerAvailable_r18 = &idx
		}
	}

	// 6. measResultReselectionNR-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.MeasResultReselectionNR_r18 = new(MeasResultIdleNR_r16)
			if err := ie.MeasResultReselectionNR_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. reselectionMeasAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(rRCResumeCompleteV1800IEsReselectionMeasAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.ReselectionMeasAvailable_r18 = &idx
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(6) {
			ie.NonCriticalExtension = new(RRCResumeComplete_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
