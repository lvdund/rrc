package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationRequest_v1700_IEs struct {
	SuccessHO_ReportReq_r17   *UEInformationRequest_v1700_IEs_successHO_ReportReq_r17   `optional`
	CoarseLocationRequest_r17 *UEInformationRequest_v1700_IEs_coarseLocationRequest_r17 `optional`
	NonCriticalExtension      interface{}                                               `optional`
}

func (ie *UEInformationRequest_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SuccessHO_ReportReq_r17 != nil, ie.CoarseLocationRequest_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SuccessHO_ReportReq_r17 != nil {
		if err = ie.SuccessHO_ReportReq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SuccessHO_ReportReq_r17", err)
		}
	}
	if ie.CoarseLocationRequest_r17 != nil {
		if err = ie.CoarseLocationRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CoarseLocationRequest_r17", err)
		}
	}
	return nil
}

func (ie *UEInformationRequest_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var SuccessHO_ReportReq_r17Present bool
	if SuccessHO_ReportReq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CoarseLocationRequest_r17Present bool
	if CoarseLocationRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SuccessHO_ReportReq_r17Present {
		ie.SuccessHO_ReportReq_r17 = new(UEInformationRequest_v1700_IEs_successHO_ReportReq_r17)
		if err = ie.SuccessHO_ReportReq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SuccessHO_ReportReq_r17", err)
		}
	}
	if CoarseLocationRequest_r17Present {
		ie.CoarseLocationRequest_r17 = new(UEInformationRequest_v1700_IEs_coarseLocationRequest_r17)
		if err = ie.CoarseLocationRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CoarseLocationRequest_r17", err)
		}
	}
	return nil
}
