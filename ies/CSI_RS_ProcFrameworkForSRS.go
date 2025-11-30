package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ProcFrameworkForSRS struct {
	MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP  int64 `lb:1,ub:4,madatory`
	MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP int64 `lb:1,ub:4,madatory`
	MaxNumberSP_SRS_AssocCSI_RS_PerBWP       int64 `lb:0,ub:4,madatory`
	SimultaneousSRS_AssocCSI_RS_PerCC        int64 `lb:1,ub:8,madatory`
}

func (ie *CSI_RS_ProcFrameworkForSRS) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSP_SRS_AssocCSI_RS_PerBWP, &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSP_SRS_AssocCSI_RS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.SimultaneousSRS_AssocCSI_RS_PerCC, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger SimultaneousSRS_AssocCSI_RS_PerCC", err)
	}
	return nil
}

func (ie *CSI_RS_ProcFrameworkForSRS) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP int64
	if tmp_int_MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP", err)
	}
	ie.MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP = tmp_int_MaxNumberPeriodicSRS_AssocCSI_RS_PerBWP
	var tmp_int_MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP int64
	if tmp_int_MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP", err)
	}
	ie.MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP = tmp_int_MaxNumberAperiodicSRS_AssocCSI_RS_PerBWP
	var tmp_int_MaxNumberSP_SRS_AssocCSI_RS_PerBWP int64
	if tmp_int_MaxNumberSP_SRS_AssocCSI_RS_PerBWP, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSP_SRS_AssocCSI_RS_PerBWP", err)
	}
	ie.MaxNumberSP_SRS_AssocCSI_RS_PerBWP = tmp_int_MaxNumberSP_SRS_AssocCSI_RS_PerBWP
	var tmp_int_SimultaneousSRS_AssocCSI_RS_PerCC int64
	if tmp_int_SimultaneousSRS_AssocCSI_RS_PerCC, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger SimultaneousSRS_AssocCSI_RS_PerCC", err)
	}
	ie.SimultaneousSRS_AssocCSI_RS_PerCC = tmp_int_SimultaneousSRS_AssocCSI_RS_PerCC
	return nil
}
