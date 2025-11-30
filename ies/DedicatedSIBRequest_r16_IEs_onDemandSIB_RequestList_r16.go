package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16 struct {
	RequestedSIB_List_r16    []SIB_ReqInfo_r16    `lb:1,ub:maxOnDemandSIB_r16,optional`
	RequestedPosSIB_List_r16 []PosSIB_ReqInfo_r16 `lb:1,ub:maxOnDemandPosSIB_r16,optional`
}

func (ie *DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.RequestedSIB_List_r16) > 0, len(ie.RequestedPosSIB_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.RequestedSIB_List_r16) > 0 {
		tmp_RequestedSIB_List_r16 := utils.NewSequence[*SIB_ReqInfo_r16]([]*SIB_ReqInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxOnDemandSIB_r16}, false)
		for _, i := range ie.RequestedSIB_List_r16 {
			tmp_RequestedSIB_List_r16.Value = append(tmp_RequestedSIB_List_r16.Value, &i)
		}
		if err = tmp_RequestedSIB_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RequestedSIB_List_r16", err)
		}
	}
	if len(ie.RequestedPosSIB_List_r16) > 0 {
		tmp_RequestedPosSIB_List_r16 := utils.NewSequence[*PosSIB_ReqInfo_r16]([]*PosSIB_ReqInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxOnDemandPosSIB_r16}, false)
		for _, i := range ie.RequestedPosSIB_List_r16 {
			tmp_RequestedPosSIB_List_r16.Value = append(tmp_RequestedPosSIB_List_r16.Value, &i)
		}
		if err = tmp_RequestedPosSIB_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RequestedPosSIB_List_r16", err)
		}
	}
	return nil
}

func (ie *DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16) Decode(r *aper.AperReader) error {
	var err error
	var RequestedSIB_List_r16Present bool
	if RequestedSIB_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RequestedPosSIB_List_r16Present bool
	if RequestedPosSIB_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RequestedSIB_List_r16Present {
		tmp_RequestedSIB_List_r16 := utils.NewSequence[*SIB_ReqInfo_r16]([]*SIB_ReqInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxOnDemandSIB_r16}, false)
		fn_RequestedSIB_List_r16 := func() *SIB_ReqInfo_r16 {
			return new(SIB_ReqInfo_r16)
		}
		if err = tmp_RequestedSIB_List_r16.Decode(r, fn_RequestedSIB_List_r16); err != nil {
			return utils.WrapError("Decode RequestedSIB_List_r16", err)
		}
		ie.RequestedSIB_List_r16 = []SIB_ReqInfo_r16{}
		for _, i := range tmp_RequestedSIB_List_r16.Value {
			ie.RequestedSIB_List_r16 = append(ie.RequestedSIB_List_r16, *i)
		}
	}
	if RequestedPosSIB_List_r16Present {
		tmp_RequestedPosSIB_List_r16 := utils.NewSequence[*PosSIB_ReqInfo_r16]([]*PosSIB_ReqInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxOnDemandPosSIB_r16}, false)
		fn_RequestedPosSIB_List_r16 := func() *PosSIB_ReqInfo_r16 {
			return new(PosSIB_ReqInfo_r16)
		}
		if err = tmp_RequestedPosSIB_List_r16.Decode(r, fn_RequestedPosSIB_List_r16); err != nil {
			return utils.WrapError("Decode RequestedPosSIB_List_r16", err)
		}
		ie.RequestedPosSIB_List_r16 = []PosSIB_ReqInfo_r16{}
		for _, i := range tmp_RequestedPosSIB_List_r16.Value {
			ie.RequestedPosSIB_List_r16 = append(ie.RequestedPosSIB_List_r16, *i)
		}
	}
	return nil
}
