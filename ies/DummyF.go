package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyF struct {
	MaxNumberPeriodicCSI_ReportPerBWP       int64 `lb:1,ub:4,madatory`
	MaxNumberAperiodicCSI_ReportPerBWP      int64 `lb:1,ub:4,madatory`
	MaxNumberSemiPersistentCSI_ReportPerBWP int64 `lb:0,ub:4,madatory`
	SimultaneousCSI_ReportsAllCC            int64 `lb:5,ub:32,madatory`
}

func (ie *DummyF) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumberPeriodicCSI_ReportPerBWP, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberPeriodicCSI_ReportPerBWP", err)
	}
	if err = w.WriteInteger(ie.MaxNumberAperiodicCSI_ReportPerBWP, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberAperiodicCSI_ReportPerBWP", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSemiPersistentCSI_ReportPerBWP, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSemiPersistentCSI_ReportPerBWP", err)
	}
	if err = w.WriteInteger(ie.SimultaneousCSI_ReportsAllCC, &uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger SimultaneousCSI_ReportsAllCC", err)
	}
	return nil
}

func (ie *DummyF) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_MaxNumberPeriodicCSI_ReportPerBWP int64
	if tmp_int_MaxNumberPeriodicCSI_ReportPerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberPeriodicCSI_ReportPerBWP", err)
	}
	ie.MaxNumberPeriodicCSI_ReportPerBWP = tmp_int_MaxNumberPeriodicCSI_ReportPerBWP
	var tmp_int_MaxNumberAperiodicCSI_ReportPerBWP int64
	if tmp_int_MaxNumberAperiodicCSI_ReportPerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberAperiodicCSI_ReportPerBWP", err)
	}
	ie.MaxNumberAperiodicCSI_ReportPerBWP = tmp_int_MaxNumberAperiodicCSI_ReportPerBWP
	var tmp_int_MaxNumberSemiPersistentCSI_ReportPerBWP int64
	if tmp_int_MaxNumberSemiPersistentCSI_ReportPerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSemiPersistentCSI_ReportPerBWP", err)
	}
	ie.MaxNumberSemiPersistentCSI_ReportPerBWP = tmp_int_MaxNumberSemiPersistentCSI_ReportPerBWP
	var tmp_int_SimultaneousCSI_ReportsAllCC int64
	if tmp_int_SimultaneousCSI_ReportsAllCC, err = r.ReadInteger(&uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger SimultaneousCSI_ReportsAllCC", err)
	}
	ie.SimultaneousCSI_ReportsAllCC = tmp_int_SimultaneousCSI_ReportsAllCC
	return nil
}
