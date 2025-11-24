package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationResponse_r16_IEs struct {
	MeasResultIdleEUTRA_r16   *MeasResultIdleEUTRA_r16         `optional`
	MeasResultIdleNR_r16      *MeasResultIdleNR_r16            `optional`
	LogMeasReport_r16         *LogMeasReport_r16               `optional`
	ConnEstFailReport_r16     *ConnEstFailReport_r16           `optional`
	Ra_ReportList_r16         *RA_ReportList_r16               `optional`
	Rlf_Report_r16            *RLF_Report_r16                  `optional`
	MobilityHistoryReport_r16 *MobilityHistoryReport_r16       `optional`
	LateNonCriticalExtension  *[]byte                          `optional`
	NonCriticalExtension      *UEInformationResponse_v1700_IEs `optional`
}

func (ie *UEInformationResponse_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultIdleEUTRA_r16 != nil, ie.MeasResultIdleNR_r16 != nil, ie.LogMeasReport_r16 != nil, ie.ConnEstFailReport_r16 != nil, ie.Ra_ReportList_r16 != nil, ie.Rlf_Report_r16 != nil, ie.MobilityHistoryReport_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasResultIdleEUTRA_r16 != nil {
		if err = ie.MeasResultIdleEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultIdleEUTRA_r16", err)
		}
	}
	if ie.MeasResultIdleNR_r16 != nil {
		if err = ie.MeasResultIdleNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultIdleNR_r16", err)
		}
	}
	if ie.LogMeasReport_r16 != nil {
		if err = ie.LogMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasReport_r16", err)
		}
	}
	if ie.ConnEstFailReport_r16 != nil {
		if err = ie.ConnEstFailReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConnEstFailReport_r16", err)
		}
	}
	if ie.Ra_ReportList_r16 != nil {
		if err = ie.Ra_ReportList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_ReportList_r16", err)
		}
	}
	if ie.Rlf_Report_r16 != nil {
		if err = ie.Rlf_Report_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rlf_Report_r16", err)
		}
	}
	if ie.MobilityHistoryReport_r16 != nil {
		if err = ie.MobilityHistoryReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MobilityHistoryReport_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
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

func (ie *UEInformationResponse_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var MeasResultIdleEUTRA_r16Present bool
	if MeasResultIdleEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultIdleNR_r16Present bool
	if MeasResultIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogMeasReport_r16Present bool
	if LogMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConnEstFailReport_r16Present bool
	if ConnEstFailReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_ReportList_r16Present bool
	if Ra_ReportList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlf_Report_r16Present bool
	if Rlf_Report_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MobilityHistoryReport_r16Present bool
	if MobilityHistoryReport_r16Present, err = r.ReadBool(); err != nil {
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
	if MeasResultIdleEUTRA_r16Present {
		ie.MeasResultIdleEUTRA_r16 = new(MeasResultIdleEUTRA_r16)
		if err = ie.MeasResultIdleEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultIdleEUTRA_r16", err)
		}
	}
	if MeasResultIdleNR_r16Present {
		ie.MeasResultIdleNR_r16 = new(MeasResultIdleNR_r16)
		if err = ie.MeasResultIdleNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultIdleNR_r16", err)
		}
	}
	if LogMeasReport_r16Present {
		ie.LogMeasReport_r16 = new(LogMeasReport_r16)
		if err = ie.LogMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasReport_r16", err)
		}
	}
	if ConnEstFailReport_r16Present {
		ie.ConnEstFailReport_r16 = new(ConnEstFailReport_r16)
		if err = ie.ConnEstFailReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConnEstFailReport_r16", err)
		}
	}
	if Ra_ReportList_r16Present {
		ie.Ra_ReportList_r16 = new(RA_ReportList_r16)
		if err = ie.Ra_ReportList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_ReportList_r16", err)
		}
	}
	if Rlf_Report_r16Present {
		ie.Rlf_Report_r16 = new(RLF_Report_r16)
		if err = ie.Rlf_Report_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rlf_Report_r16", err)
		}
	}
	if MobilityHistoryReport_r16Present {
		ie.MobilityHistoryReport_r16 = new(MobilityHistoryReport_r16)
		if err = ie.MobilityHistoryReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityHistoryReport_r16", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UEInformationResponse_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
