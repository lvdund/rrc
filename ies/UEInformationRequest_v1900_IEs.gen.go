// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationRequest-v1900-IEs (line 2930).

var uEInformationRequestV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-LogMeasReportReq-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UEInformationRequest_v1900_IEs_Csi_LogMeasReportReq_r19_True = 0
)

var uEInformationRequestV1900IEsCsiLogMeasReportReqR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_v1900_IEs_Csi_LogMeasReportReq_r19_True},
}

type UEInformationRequest_v1900_IEs struct {
	Csi_LogMeasReportReq_r19 *int64
	NonCriticalExtension     *struct{}
}

func (ie *UEInformationRequest_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationRequestV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_LogMeasReportReq_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. csi-LogMeasReportReq-r19: enumerated
	{
		if ie.Csi_LogMeasReportReq_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_LogMeasReportReq_r19, uEInformationRequestV1900IEsCsiLogMeasReportReqR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UEInformationRequest_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationRequestV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. csi-LogMeasReportReq-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uEInformationRequestV1900IEsCsiLogMeasReportReqR19Constraints)
			if err != nil {
				return err
			}
			ie.Csi_LogMeasReportReq_r19 = &idx
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
