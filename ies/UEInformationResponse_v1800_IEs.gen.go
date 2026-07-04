// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationResponse-v1800-IEs (line 2970).

var uEInformationResponseV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "flightPathInfoReport-r18", Optional: true},
		{Name: "successPSCell-Report-r18", Optional: true},
		{Name: "measResultReselectionNR-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UEInformationResponse_v1800_IEs struct {
	FlightPathInfoReport_r18    *FlightPathInfoReport_r18
	SuccessPSCell_Report_r18    *SuccessPSCell_Report_r18
	MeasResultReselectionNR_r18 *MeasResultIdleNR_r16
	NonCriticalExtension        *UEInformationResponse_v1900_IEs
}

func (ie *UEInformationResponse_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationResponseV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FlightPathInfoReport_r18 != nil, ie.SuccessPSCell_Report_r18 != nil, ie.MeasResultReselectionNR_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. flightPathInfoReport-r18: ref
	{
		if ie.FlightPathInfoReport_r18 != nil {
			if err := ie.FlightPathInfoReport_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. successPSCell-Report-r18: ref
	{
		if ie.SuccessPSCell_Report_r18 != nil {
			if err := ie.SuccessPSCell_Report_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measResultReselectionNR-r18: ref
	{
		if ie.MeasResultReselectionNR_r18 != nil {
			if err := ie.MeasResultReselectionNR_r18.Encode(e); err != nil {
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

func (ie *UEInformationResponse_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationResponseV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. flightPathInfoReport-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FlightPathInfoReport_r18 = new(FlightPathInfoReport_r18)
			if err := ie.FlightPathInfoReport_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. successPSCell-Report-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SuccessPSCell_Report_r18 = new(SuccessPSCell_Report_r18)
			if err := ie.SuccessPSCell_Report_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measResultReselectionNR-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultReselectionNR_r18 = new(MeasResultIdleNR_r16)
			if err := ie.MeasResultReselectionNR_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(UEInformationResponse_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
