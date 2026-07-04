// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationRequest-v1700-IEs (line 2916).

var uEInformationRequestV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "successHO-ReportReq-r17", Optional: true},
		{Name: "coarseLocationRequest-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UEInformationRequest_v1700_IEs_SuccessHO_ReportReq_r17_True = 0
)

var uEInformationRequestV1700IEsSuccessHOReportReqR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_v1700_IEs_SuccessHO_ReportReq_r17_True},
}

const (
	UEInformationRequest_v1700_IEs_CoarseLocationRequest_r17_True = 0
)

var uEInformationRequestV1700IEsCoarseLocationRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEInformationRequest_v1700_IEs_CoarseLocationRequest_r17_True},
}

type UEInformationRequest_v1700_IEs struct {
	SuccessHO_ReportReq_r17   *int64
	CoarseLocationRequest_r17 *int64
	NonCriticalExtension      *UEInformationRequest_v1800_IEs
}

func (ie *UEInformationRequest_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationRequestV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SuccessHO_ReportReq_r17 != nil, ie.CoarseLocationRequest_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. successHO-ReportReq-r17: enumerated
	{
		if ie.SuccessHO_ReportReq_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SuccessHO_ReportReq_r17, uEInformationRequestV1700IEsSuccessHOReportReqR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. coarseLocationRequest-r17: enumerated
	{
		if ie.CoarseLocationRequest_r17 != nil {
			if err := e.EncodeEnumerated(*ie.CoarseLocationRequest_r17, uEInformationRequestV1700IEsCoarseLocationRequestR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEInformationRequest_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationRequestV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. successHO-ReportReq-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uEInformationRequestV1700IEsSuccessHOReportReqR17Constraints)
			if err != nil {
				return err
			}
			ie.SuccessHO_ReportReq_r17 = &idx
		}
	}

	// 3. coarseLocationRequest-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uEInformationRequestV1700IEsCoarseLocationRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.CoarseLocationRequest_r17 = &idx
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UEInformationRequest_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
