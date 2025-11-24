package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TraceReference_r16 struct {
	Plmn_Identity_r16 PLMN_Identity `madatory`
	TraceId_r16       []byte        `lb:3,ub:3,madatory`
}

func (ie *TraceReference_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_Identity_r16", err)
	}
	if err = w.WriteOctetString(ie.TraceId_r16, &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteOctetString TraceId_r16", err)
	}
	return nil
}

func (ie *TraceReference_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_Identity_r16", err)
	}
	var tmp_os_TraceId_r16 []byte
	if tmp_os_TraceId_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadOctetString TraceId_r16", err)
	}
	ie.TraceId_r16 = tmp_os_TraceId_r16
	return nil
}
