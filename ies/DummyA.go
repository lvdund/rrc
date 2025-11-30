package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DummyA struct {
	MaxNumberNZP_CSI_RS_PerCC                       int64                                                  `lb:1,ub:32,madatory`
	MaxNumberPortsAcrossNZP_CSI_RS_PerCC            DummyA_maxNumberPortsAcrossNZP_CSI_RS_PerCC            `madatory`
	MaxNumberCS_IM_PerCC                            DummyA_maxNumberCS_IM_PerCC                            `madatory`
	MaxNumberSimultaneousCSI_RS_ActBWP_AllCC        DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC        `madatory`
	TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC `madatory`
}

func (ie *DummyA) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumberNZP_CSI_RS_PerCC, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberNZP_CSI_RS_PerCC", err)
	}
	if err = ie.MaxNumberPortsAcrossNZP_CSI_RS_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	if err = ie.MaxNumberCS_IM_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberCS_IM_PerCC", err)
	}
	if err = ie.MaxNumberSimultaneousCSI_RS_ActBWP_AllCC.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	if err = ie.TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC.Encode(w); err != nil {
		return utils.WrapError("Encode TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}

func (ie *DummyA) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_MaxNumberNZP_CSI_RS_PerCC int64
	if tmp_int_MaxNumberNZP_CSI_RS_PerCC, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberNZP_CSI_RS_PerCC", err)
	}
	ie.MaxNumberNZP_CSI_RS_PerCC = tmp_int_MaxNumberNZP_CSI_RS_PerCC
	if err = ie.MaxNumberPortsAcrossNZP_CSI_RS_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	if err = ie.MaxNumberCS_IM_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberCS_IM_PerCC", err)
	}
	if err = ie.MaxNumberSimultaneousCSI_RS_ActBWP_AllCC.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	if err = ie.TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC.Decode(r); err != nil {
		return utils.WrapError("Decode TotalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}
