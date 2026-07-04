// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationRequest-r16-IEs (line 2905).

var uEInformationRequestR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idleModeMeasurementReq-r16", Optional: true},
		{Name: "logMeasReportReq-r16", Optional: true},
		{Name: "connEstFailReportReq-r16", Optional: true},
		{Name: "ra-ReportReq-r16", Optional: true},
		{Name: "rlf-ReportReq-r16", Optional: true},
		{Name: "mobilityHistoryReportReq-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UEInformationRequest_r16_IEs_IdleModeMeasurementReq_r16_True = 0
)

var uEInformationRequestR16IEsIdleModeMeasurementReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_r16_IEs_IdleModeMeasurementReq_r16_True},
}

const (
	UEInformationRequest_r16_IEs_LogMeasReportReq_r16_True = 0
)

var uEInformationRequestR16IEsLogMeasReportReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_r16_IEs_LogMeasReportReq_r16_True},
}

const (
	UEInformationRequest_r16_IEs_ConnEstFailReportReq_r16_True = 0
)

var uEInformationRequestR16IEsConnEstFailReportReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_r16_IEs_ConnEstFailReportReq_r16_True},
}

const (
	UEInformationRequest_r16_IEs_Ra_ReportReq_r16_True = 0
)

var uEInformationRequestR16IEsRaReportReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_r16_IEs_Ra_ReportReq_r16_True},
}

const (
	UEInformationRequest_r16_IEs_Rlf_ReportReq_r16_True = 0
)

var uEInformationRequestR16IEsRlfReportReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_r16_IEs_Rlf_ReportReq_r16_True},
}

const (
	UEInformationRequest_r16_IEs_MobilityHistoryReportReq_r16_True = 0
)

var uEInformationRequestR16IEsMobilityHistoryReportReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_r16_IEs_MobilityHistoryReportReq_r16_True},
}

var uEInformationRequestR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UEInformationRequest_r16_IEs struct {
	IdleModeMeasurementReq_r16   *int64
	LogMeasReportReq_r16         *int64
	ConnEstFailReportReq_r16     *int64
	Ra_ReportReq_r16             *int64
	Rlf_ReportReq_r16            *int64
	MobilityHistoryReportReq_r16 *int64
	LateNonCriticalExtension     []byte
	NonCriticalExtension         *UEInformationRequest_v1700_IEs
}

func (ie *UEInformationRequest_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationRequestR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IdleModeMeasurementReq_r16 != nil, ie.LogMeasReportReq_r16 != nil, ie.ConnEstFailReportReq_r16 != nil, ie.Ra_ReportReq_r16 != nil, ie.Rlf_ReportReq_r16 != nil, ie.MobilityHistoryReportReq_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. idleModeMeasurementReq-r16: enumerated
	{
		if ie.IdleModeMeasurementReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IdleModeMeasurementReq_r16, uEInformationRequestR16IEsIdleModeMeasurementReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. logMeasReportReq-r16: enumerated
	{
		if ie.LogMeasReportReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LogMeasReportReq_r16, uEInformationRequestR16IEsLogMeasReportReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. connEstFailReportReq-r16: enumerated
	{
		if ie.ConnEstFailReportReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ConnEstFailReportReq_r16, uEInformationRequestR16IEsConnEstFailReportReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ra-ReportReq-r16: enumerated
	{
		if ie.Ra_ReportReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ra_ReportReq_r16, uEInformationRequestR16IEsRaReportReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. rlf-ReportReq-r16: enumerated
	{
		if ie.Rlf_ReportReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Rlf_ReportReq_r16, uEInformationRequestR16IEsRlfReportReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. mobilityHistoryReportReq-r16: enumerated
	{
		if ie.MobilityHistoryReportReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MobilityHistoryReportReq_r16, uEInformationRequestR16IEsMobilityHistoryReportReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uEInformationRequestR16IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 9. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEInformationRequest_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationRequestR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idleModeMeasurementReq-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uEInformationRequestR16IEsIdleModeMeasurementReqR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleModeMeasurementReq_r16 = &idx
		}
	}

	// 3. logMeasReportReq-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uEInformationRequestR16IEsLogMeasReportReqR16Constraints)
			if err != nil {
				return err
			}
			ie.LogMeasReportReq_r16 = &idx
		}
	}

	// 4. connEstFailReportReq-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uEInformationRequestR16IEsConnEstFailReportReqR16Constraints)
			if err != nil {
				return err
			}
			ie.ConnEstFailReportReq_r16 = &idx
		}
	}

	// 5. ra-ReportReq-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(uEInformationRequestR16IEsRaReportReqR16Constraints)
			if err != nil {
				return err
			}
			ie.Ra_ReportReq_r16 = &idx
		}
	}

	// 6. rlf-ReportReq-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uEInformationRequestR16IEsRlfReportReqR16Constraints)
			if err != nil {
				return err
			}
			ie.Rlf_ReportReq_r16 = &idx
		}
	}

	// 7. mobilityHistoryReportReq-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(uEInformationRequestR16IEsMobilityHistoryReportReqR16Constraints)
			if err != nil {
				return err
			}
			ie.MobilityHistoryReportReq_r16 = &idx
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(uEInformationRequestR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 9. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(7) {
			ie.NonCriticalExtension = new(UEInformationRequest_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
