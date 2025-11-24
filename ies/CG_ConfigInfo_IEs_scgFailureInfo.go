package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_IEs_scgFailureInfo struct {
	FailureType   CG_ConfigInfo_IEs_scgFailureInfo_failureType `madatory`
	MeasResultSCG []byte                                       `madatory`
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.FailureType.Encode(w); err != nil {
		return utils.WrapError("Encode FailureType", err)
	}
	if err = w.WriteOctetString(ie.MeasResultSCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString MeasResultSCG", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.FailureType.Decode(r); err != nil {
		return utils.WrapError("Decode FailureType", err)
	}
	var tmp_os_MeasResultSCG []byte
	if tmp_os_MeasResultSCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString MeasResultSCG", err)
	}
	ie.MeasResultSCG = tmp_os_MeasResultSCG
	return nil
}
