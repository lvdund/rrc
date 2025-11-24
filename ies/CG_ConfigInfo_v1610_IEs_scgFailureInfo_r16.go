package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16 struct {
	FailureType_r16   CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16_failureType_r16 `madatory`
	MeasResultSCG_r16 []byte                                                     `madatory`
}

func (ie *CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.FailureType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FailureType_r16", err)
	}
	if err = w.WriteOctetString(ie.MeasResultSCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString MeasResultSCG_r16", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.FailureType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FailureType_r16", err)
	}
	var tmp_os_MeasResultSCG_r16 []byte
	if tmp_os_MeasResultSCG_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString MeasResultSCG_r16", err)
	}
	ie.MeasResultSCG_r16 = tmp_os_MeasResultSCG_r16
	return nil
}
