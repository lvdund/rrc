package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_IM_ReceptionForFeedback struct {
	MaxConfigNumberNZP_CSI_RS_PerCC              int64                                                      `lb:1,ub:64,madatory`
	MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC   int64                                                      `lb:2,ub:256,madatory`
	MaxConfigNumberCSI_IM_PerCC                  CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC `madatory`
	MaxNumberSimultaneousNZP_CSI_RS_PerCC        int64                                                      `lb:1,ub:64,madatory`
	TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC int64                                                      `lb:2,ub:256,madatory`
}

func (ie *CSI_RS_IM_ReceptionForFeedback) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxConfigNumberNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger MaxConfigNumberNZP_CSI_RS_PerCC", err)
	}
	if err = w.WriteInteger(ie.MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	if err = ie.MaxConfigNumberCSI_IM_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode MaxConfigNumberCSI_IM_PerCC", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSimultaneousNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSimultaneousNZP_CSI_RS_PerCC", err)
	}
	if err = w.WriteInteger(ie.TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC", err)
	}
	return nil
}

func (ie *CSI_RS_IM_ReceptionForFeedback) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_MaxConfigNumberNZP_CSI_RS_PerCC int64
	if tmp_int_MaxConfigNumberNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger MaxConfigNumberNZP_CSI_RS_PerCC", err)
	}
	ie.MaxConfigNumberNZP_CSI_RS_PerCC = tmp_int_MaxConfigNumberNZP_CSI_RS_PerCC
	var tmp_int_MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC int64
	if tmp_int_MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	ie.MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC = tmp_int_MaxConfigNumberPortsAcrossNZP_CSI_RS_PerCC
	if err = ie.MaxConfigNumberCSI_IM_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode MaxConfigNumberCSI_IM_PerCC", err)
	}
	var tmp_int_MaxNumberSimultaneousNZP_CSI_RS_PerCC int64
	if tmp_int_MaxNumberSimultaneousNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSimultaneousNZP_CSI_RS_PerCC", err)
	}
	ie.MaxNumberSimultaneousNZP_CSI_RS_PerCC = tmp_int_MaxNumberSimultaneousNZP_CSI_RS_PerCC
	var tmp_int_TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC int64
	if tmp_int_TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC", err)
	}
	ie.TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC = tmp_int_TotalNumberPortsSimultaneousNZP_CSI_RS_PerCC
	return nil
}
