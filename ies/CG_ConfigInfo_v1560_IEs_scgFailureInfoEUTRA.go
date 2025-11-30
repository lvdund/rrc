package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA struct {
	FailureTypeEUTRA    CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA `madatory`
	MeasResultSCG_EUTRA []byte                                                       `madatory`
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.FailureTypeEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode FailureTypeEUTRA", err)
	}
	if err = w.WriteOctetString(ie.MeasResultSCG_EUTRA, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString MeasResultSCG_EUTRA", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.FailureTypeEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode FailureTypeEUTRA", err)
	}
	var tmp_os_MeasResultSCG_EUTRA []byte
	if tmp_os_MeasResultSCG_EUTRA, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString MeasResultSCG_EUTRA", err)
	}
	ie.MeasResultSCG_EUTRA = tmp_os_MeasResultSCG_EUTRA
	return nil
}
