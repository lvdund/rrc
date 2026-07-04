// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationRequest-v1800-IEs (line 2922).

var uEInformationRequestV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "flightPathInfoReq-r18", Optional: true},
		{Name: "successPSCell-ReportReq-r18", Optional: true},
		{Name: "reselectionMeasurementReq-r18", Optional: true},
		{Name: "validatedMeasurementsReq-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UEInformationRequest_v1800_IEs_SuccessPSCell_ReportReq_r18_True = 0
)

var uEInformationRequestV1800IEsSuccessPSCellReportReqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_v1800_IEs_SuccessPSCell_ReportReq_r18_True},
}

const (
	UEInformationRequest_v1800_IEs_ReselectionMeasurementReq_r18_True = 0
)

var uEInformationRequestV1800IEsReselectionMeasurementReqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_v1800_IEs_ReselectionMeasurementReq_r18_True},
}

const (
	UEInformationRequest_v1800_IEs_ValidatedMeasurementsReq_r18_True = 0
)

var uEInformationRequestV1800IEsValidatedMeasurementsReqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_v1800_IEs_ValidatedMeasurementsReq_r18_True},
}

type UEInformationRequest_v1800_IEs struct {
	FlightPathInfoReq_r18         *FlightPathInfoReportConfig_r18
	SuccessPSCell_ReportReq_r18   *int64
	ReselectionMeasurementReq_r18 *int64
	ValidatedMeasurementsReq_r18  *int64
	NonCriticalExtension          *UEInformationRequest_v1900_IEs
}

func (ie *UEInformationRequest_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationRequestV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FlightPathInfoReq_r18 != nil, ie.SuccessPSCell_ReportReq_r18 != nil, ie.ReselectionMeasurementReq_r18 != nil, ie.ValidatedMeasurementsReq_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. flightPathInfoReq-r18: ref
	{
		if ie.FlightPathInfoReq_r18 != nil {
			if err := ie.FlightPathInfoReq_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. successPSCell-ReportReq-r18: enumerated
	{
		if ie.SuccessPSCell_ReportReq_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SuccessPSCell_ReportReq_r18, uEInformationRequestV1800IEsSuccessPSCellReportReqR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. reselectionMeasurementReq-r18: enumerated
	{
		if ie.ReselectionMeasurementReq_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ReselectionMeasurementReq_r18, uEInformationRequestV1800IEsReselectionMeasurementReqR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. validatedMeasurementsReq-r18: enumerated
	{
		if ie.ValidatedMeasurementsReq_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ValidatedMeasurementsReq_r18, uEInformationRequestV1800IEsValidatedMeasurementsReqR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEInformationRequest_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationRequestV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. flightPathInfoReq-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FlightPathInfoReq_r18 = new(FlightPathInfoReportConfig_r18)
			if err := ie.FlightPathInfoReq_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. successPSCell-ReportReq-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uEInformationRequestV1800IEsSuccessPSCellReportReqR18Constraints)
			if err != nil {
				return err
			}
			ie.SuccessPSCell_ReportReq_r18 = &idx
		}
	}

	// 4. reselectionMeasurementReq-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uEInformationRequestV1800IEsReselectionMeasurementReqR18Constraints)
			if err != nil {
				return err
			}
			ie.ReselectionMeasurementReq_r18 = &idx
		}
	}

	// 5. validatedMeasurementsReq-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(uEInformationRequestV1800IEsValidatedMeasurementsReqR18Constraints)
			if err != nil {
				return err
			}
			ie.ValidatedMeasurementsReq_r18 = &idx
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(UEInformationRequest_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
