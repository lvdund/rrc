package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationResponse_v1700_IEs struct {
	SuccessHO_Report_r17      *SuccessHO_Report_r17      `optional`
	ConnEstFailReportList_r17 *ConnEstFailReportList_r17 `optional`
	CoarseLocationInfo_r17    *[]byte                    `optional`
	NonCriticalExtension      interface{}                `optional`
}

func (ie *UEInformationResponse_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SuccessHO_Report_r17 != nil, ie.ConnEstFailReportList_r17 != nil, ie.CoarseLocationInfo_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SuccessHO_Report_r17 != nil {
		if err = ie.SuccessHO_Report_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SuccessHO_Report_r17", err)
		}
	}
	if ie.ConnEstFailReportList_r17 != nil {
		if err = ie.ConnEstFailReportList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ConnEstFailReportList_r17", err)
		}
	}
	if ie.CoarseLocationInfo_r17 != nil {
		if err = w.WriteOctetString(*ie.CoarseLocationInfo_r17, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode CoarseLocationInfo_r17", err)
		}
	}
	return nil
}

func (ie *UEInformationResponse_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var SuccessHO_Report_r17Present bool
	if SuccessHO_Report_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConnEstFailReportList_r17Present bool
	if ConnEstFailReportList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CoarseLocationInfo_r17Present bool
	if CoarseLocationInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SuccessHO_Report_r17Present {
		ie.SuccessHO_Report_r17 = new(SuccessHO_Report_r17)
		if err = ie.SuccessHO_Report_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SuccessHO_Report_r17", err)
		}
	}
	if ConnEstFailReportList_r17Present {
		ie.ConnEstFailReportList_r17 = new(ConnEstFailReportList_r17)
		if err = ie.ConnEstFailReportList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ConnEstFailReportList_r17", err)
		}
	}
	if CoarseLocationInfo_r17Present {
		var tmp_os_CoarseLocationInfo_r17 []byte
		if tmp_os_CoarseLocationInfo_r17, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode CoarseLocationInfo_r17", err)
		}
		ie.CoarseLocationInfo_r17 = &tmp_os_CoarseLocationInfo_r17
	}
	return nil
}
