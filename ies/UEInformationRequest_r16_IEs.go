package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationRequest_r16_IEs struct {
	IdleModeMeasurementReq_r16   *UEInformationRequest_r16_IEs_idleModeMeasurementReq_r16   `optional`
	LogMeasReportReq_r16         *UEInformationRequest_r16_IEs_logMeasReportReq_r16         `optional`
	ConnEstFailReportReq_r16     *UEInformationRequest_r16_IEs_connEstFailReportReq_r16     `optional`
	Ra_ReportReq_r16             *UEInformationRequest_r16_IEs_ra_ReportReq_r16             `optional`
	Rlf_ReportReq_r16            *UEInformationRequest_r16_IEs_rlf_ReportReq_r16            `optional`
	MobilityHistoryReportReq_r16 *UEInformationRequest_r16_IEs_mobilityHistoryReportReq_r16 `optional`
	LateNonCriticalExtension     *[]byte                                                    `optional`
	NonCriticalExtension         *UEInformationRequest_v1700_IEs                            `optional`
}

func (ie *UEInformationRequest_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IdleModeMeasurementReq_r16 != nil, ie.LogMeasReportReq_r16 != nil, ie.ConnEstFailReportReq_r16 != nil, ie.Ra_ReportReq_r16 != nil, ie.Rlf_ReportReq_r16 != nil, ie.MobilityHistoryReportReq_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IdleModeMeasurementReq_r16 != nil {
		if err = ie.IdleModeMeasurementReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleModeMeasurementReq_r16", err)
		}
	}
	if ie.LogMeasReportReq_r16 != nil {
		if err = ie.LogMeasReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasReportReq_r16", err)
		}
	}
	if ie.ConnEstFailReportReq_r16 != nil {
		if err = ie.ConnEstFailReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConnEstFailReportReq_r16", err)
		}
	}
	if ie.Ra_ReportReq_r16 != nil {
		if err = ie.Ra_ReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_ReportReq_r16", err)
		}
	}
	if ie.Rlf_ReportReq_r16 != nil {
		if err = ie.Rlf_ReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rlf_ReportReq_r16", err)
		}
	}
	if ie.MobilityHistoryReportReq_r16 != nil {
		if err = ie.MobilityHistoryReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MobilityHistoryReportReq_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEInformationRequest_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var IdleModeMeasurementReq_r16Present bool
	if IdleModeMeasurementReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogMeasReportReq_r16Present bool
	if LogMeasReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConnEstFailReportReq_r16Present bool
	if ConnEstFailReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_ReportReq_r16Present bool
	if Ra_ReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlf_ReportReq_r16Present bool
	if Rlf_ReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MobilityHistoryReportReq_r16Present bool
	if MobilityHistoryReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if IdleModeMeasurementReq_r16Present {
		ie.IdleModeMeasurementReq_r16 = new(UEInformationRequest_r16_IEs_idleModeMeasurementReq_r16)
		if err = ie.IdleModeMeasurementReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleModeMeasurementReq_r16", err)
		}
	}
	if LogMeasReportReq_r16Present {
		ie.LogMeasReportReq_r16 = new(UEInformationRequest_r16_IEs_logMeasReportReq_r16)
		if err = ie.LogMeasReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasReportReq_r16", err)
		}
	}
	if ConnEstFailReportReq_r16Present {
		ie.ConnEstFailReportReq_r16 = new(UEInformationRequest_r16_IEs_connEstFailReportReq_r16)
		if err = ie.ConnEstFailReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConnEstFailReportReq_r16", err)
		}
	}
	if Ra_ReportReq_r16Present {
		ie.Ra_ReportReq_r16 = new(UEInformationRequest_r16_IEs_ra_ReportReq_r16)
		if err = ie.Ra_ReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_ReportReq_r16", err)
		}
	}
	if Rlf_ReportReq_r16Present {
		ie.Rlf_ReportReq_r16 = new(UEInformationRequest_r16_IEs_rlf_ReportReq_r16)
		if err = ie.Rlf_ReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rlf_ReportReq_r16", err)
		}
	}
	if MobilityHistoryReportReq_r16Present {
		ie.MobilityHistoryReportReq_r16 = new(UEInformationRequest_r16_IEs_mobilityHistoryReportReq_r16)
		if err = ie.MobilityHistoryReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityHistoryReportReq_r16", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UEInformationRequest_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
