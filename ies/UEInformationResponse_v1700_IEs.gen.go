// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationResponse-v1700-IEs (line 2963).

var uEInformationResponseV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "successHO-Report-r17", Optional: true},
		{Name: "connEstFailReportList-r17", Optional: true},
		{Name: "coarseLocationInfo-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEInformationResponseV1700IEsCoarseLocationInfoR17Constraints = per.SizeConstraints{}

type UEInformationResponse_v1700_IEs struct {
	SuccessHO_Report_r17      *SuccessHO_Report_r17
	ConnEstFailReportList_r17 *ConnEstFailReportList_r17
	CoarseLocationInfo_r17    []byte
	NonCriticalExtension      *UEInformationResponse_v1800_IEs
}

func (ie *UEInformationResponse_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationResponseV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SuccessHO_Report_r17 != nil, ie.ConnEstFailReportList_r17 != nil, ie.CoarseLocationInfo_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. successHO-Report-r17: ref
	{
		if ie.SuccessHO_Report_r17 != nil {
			if err := ie.SuccessHO_Report_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. connEstFailReportList-r17: ref
	{
		if ie.ConnEstFailReportList_r17 != nil {
			if err := ie.ConnEstFailReportList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. coarseLocationInfo-r17: octet-string
	{
		if ie.CoarseLocationInfo_r17 != nil {
			if err := e.EncodeOctetString(ie.CoarseLocationInfo_r17, uEInformationResponseV1700IEsCoarseLocationInfoR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEInformationResponse_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationResponseV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. successHO-Report-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SuccessHO_Report_r17 = new(SuccessHO_Report_r17)
			if err := ie.SuccessHO_Report_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. connEstFailReportList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ConnEstFailReportList_r17 = new(ConnEstFailReportList_r17)
			if err := ie.ConnEstFailReportList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. coarseLocationInfo-r17: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(uEInformationResponseV1700IEsCoarseLocationInfoR17Constraints)
			if err != nil {
				return err
			}
			ie.CoarseLocationInfo_r17 = v
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(UEInformationResponse_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
