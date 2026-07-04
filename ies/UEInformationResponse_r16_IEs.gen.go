// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationResponse-r16-IEs (line 2951).

var uEInformationResponseR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultIdleEUTRA-r16", Optional: true},
		{Name: "measResultIdleNR-r16", Optional: true},
		{Name: "logMeasReport-r16", Optional: true},
		{Name: "connEstFailReport-r16", Optional: true},
		{Name: "ra-ReportList-r16", Optional: true},
		{Name: "rlf-Report-r16", Optional: true},
		{Name: "mobilityHistoryReport-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEInformationResponseR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UEInformationResponse_r16_IEs struct {
	MeasResultIdleEUTRA_r16   *MeasResultIdleEUTRA_r16
	MeasResultIdleNR_r16      *MeasResultIdleNR_r16
	LogMeasReport_r16         *LogMeasReport_r16
	ConnEstFailReport_r16     *ConnEstFailReport_r16
	Ra_ReportList_r16         *RA_ReportList_r16
	Rlf_Report_r16            *RLF_Report_r16
	MobilityHistoryReport_r16 *MobilityHistoryReport_r16
	LateNonCriticalExtension  []byte
	NonCriticalExtension      *UEInformationResponse_v1700_IEs
}

func (ie *UEInformationResponse_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationResponseR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultIdleEUTRA_r16 != nil, ie.MeasResultIdleNR_r16 != nil, ie.LogMeasReport_r16 != nil, ie.ConnEstFailReport_r16 != nil, ie.Ra_ReportList_r16 != nil, ie.Rlf_Report_r16 != nil, ie.MobilityHistoryReport_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measResultIdleEUTRA-r16: ref
	{
		if ie.MeasResultIdleEUTRA_r16 != nil {
			if err := ie.MeasResultIdleEUTRA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. measResultIdleNR-r16: ref
	{
		if ie.MeasResultIdleNR_r16 != nil {
			if err := ie.MeasResultIdleNR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. logMeasReport-r16: ref
	{
		if ie.LogMeasReport_r16 != nil {
			if err := ie.LogMeasReport_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. connEstFailReport-r16: ref
	{
		if ie.ConnEstFailReport_r16 != nil {
			if err := ie.ConnEstFailReport_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. ra-ReportList-r16: ref
	{
		if ie.Ra_ReportList_r16 != nil {
			if err := ie.Ra_ReportList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. rlf-Report-r16: ref
	{
		if ie.Rlf_Report_r16 != nil {
			if err := ie.Rlf_Report_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. mobilityHistoryReport-r16: ref
	{
		if ie.MobilityHistoryReport_r16 != nil {
			if err := ie.MobilityHistoryReport_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uEInformationResponseR16IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEInformationResponse_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationResponseR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measResultIdleEUTRA-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasResultIdleEUTRA_r16 = new(MeasResultIdleEUTRA_r16)
			if err := ie.MeasResultIdleEUTRA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. measResultIdleNR-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasResultIdleNR_r16 = new(MeasResultIdleNR_r16)
			if err := ie.MeasResultIdleNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. logMeasReport-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.LogMeasReport_r16 = new(LogMeasReport_r16)
			if err := ie.LogMeasReport_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. connEstFailReport-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.ConnEstFailReport_r16 = new(ConnEstFailReport_r16)
			if err := ie.ConnEstFailReport_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. ra-ReportList-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ra_ReportList_r16 = new(RA_ReportList_r16)
			if err := ie.Ra_ReportList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. rlf-Report-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Rlf_Report_r16 = new(RLF_Report_r16)
			if err := ie.Rlf_Report_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. mobilityHistoryReport-r16: ref
	{
		if seq.IsComponentPresent(6) {
			ie.MobilityHistoryReport_r16 = new(MobilityHistoryReport_r16)
			if err := ie.MobilityHistoryReport_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeOctetString(uEInformationResponseR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(8) {
			ie.NonCriticalExtension = new(UEInformationResponse_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
