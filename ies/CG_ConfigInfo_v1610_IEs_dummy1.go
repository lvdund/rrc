package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1610_IEs_dummy1 struct {
	FailureTypeEUTRA_r16    CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16 `madatory`
	MeasResultSCG_EUTRA_r16 []byte                                              `madatory`
}

func (ie *CG_ConfigInfo_v1610_IEs_dummy1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.FailureTypeEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FailureTypeEUTRA_r16", err)
	}
	if err = w.WriteOctetString(ie.MeasResultSCG_EUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString MeasResultSCG_EUTRA_r16", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1610_IEs_dummy1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.FailureTypeEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FailureTypeEUTRA_r16", err)
	}
	var tmp_os_MeasResultSCG_EUTRA_r16 []byte
	if tmp_os_MeasResultSCG_EUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString MeasResultSCG_EUTRA_r16", err)
	}
	ie.MeasResultSCG_EUTRA_r16 = tmp_os_MeasResultSCG_EUTRA_r16
	return nil
}
