package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportFramework struct {
	MaxNumberPeriodicCSI_PerBWP_ForCSI_Report       int64                                                          `lb:1,ub:4,madatory`
	MaxNumberAperiodicCSI_PerBWP_ForCSI_Report      int64                                                          `lb:1,ub:4,madatory`
	MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report int64                                                          `lb:0,ub:4,madatory`
	MaxNumberPeriodicCSI_PerBWP_ForBeamReport       int64                                                          `lb:1,ub:4,madatory`
	MaxNumberAperiodicCSI_PerBWP_ForBeamReport      int64                                                          `lb:1,ub:4,madatory`
	MaxNumberAperiodicCSI_triggeringStatePerCC      CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC `madatory`
	MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport int64                                                          `lb:0,ub:4,madatory`
	SimultaneousCSI_ReportsPerCC                    int64                                                          `lb:1,ub:8,madatory`
}

func (ie *CSI_ReportFramework) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumberPeriodicCSI_PerBWP_ForCSI_Report, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberPeriodicCSI_PerBWP_ForCSI_Report", err)
	}
	if err = w.WriteInteger(ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_Report, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberAperiodicCSI_PerBWP_ForCSI_Report", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report, &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report", err)
	}
	if err = w.WriteInteger(ie.MaxNumberPeriodicCSI_PerBWP_ForBeamReport, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberPeriodicCSI_PerBWP_ForBeamReport", err)
	}
	if err = w.WriteInteger(ie.MaxNumberAperiodicCSI_PerBWP_ForBeamReport, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberAperiodicCSI_PerBWP_ForBeamReport", err)
	}
	if err = ie.MaxNumberAperiodicCSI_triggeringStatePerCC.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	if err = w.WriteInteger(ie.MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport, &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport", err)
	}
	if err = w.WriteInteger(ie.SimultaneousCSI_ReportsPerCC, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger SimultaneousCSI_ReportsPerCC", err)
	}
	return nil
}

func (ie *CSI_ReportFramework) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_MaxNumberPeriodicCSI_PerBWP_ForCSI_Report int64
	if tmp_int_MaxNumberPeriodicCSI_PerBWP_ForCSI_Report, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberPeriodicCSI_PerBWP_ForCSI_Report", err)
	}
	ie.MaxNumberPeriodicCSI_PerBWP_ForCSI_Report = tmp_int_MaxNumberPeriodicCSI_PerBWP_ForCSI_Report
	var tmp_int_MaxNumberAperiodicCSI_PerBWP_ForCSI_Report int64
	if tmp_int_MaxNumberAperiodicCSI_PerBWP_ForCSI_Report, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberAperiodicCSI_PerBWP_ForCSI_Report", err)
	}
	ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_Report = tmp_int_MaxNumberAperiodicCSI_PerBWP_ForCSI_Report
	var tmp_int_MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report int64
	if tmp_int_MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report", err)
	}
	ie.MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report = tmp_int_MaxNumberSemiPersistentCSI_PerBWP_ForCSI_Report
	var tmp_int_MaxNumberPeriodicCSI_PerBWP_ForBeamReport int64
	if tmp_int_MaxNumberPeriodicCSI_PerBWP_ForBeamReport, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberPeriodicCSI_PerBWP_ForBeamReport", err)
	}
	ie.MaxNumberPeriodicCSI_PerBWP_ForBeamReport = tmp_int_MaxNumberPeriodicCSI_PerBWP_ForBeamReport
	var tmp_int_MaxNumberAperiodicCSI_PerBWP_ForBeamReport int64
	if tmp_int_MaxNumberAperiodicCSI_PerBWP_ForBeamReport, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberAperiodicCSI_PerBWP_ForBeamReport", err)
	}
	ie.MaxNumberAperiodicCSI_PerBWP_ForBeamReport = tmp_int_MaxNumberAperiodicCSI_PerBWP_ForBeamReport
	if err = ie.MaxNumberAperiodicCSI_triggeringStatePerCC.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	var tmp_int_MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport int64
	if tmp_int_MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport", err)
	}
	ie.MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport = tmp_int_MaxNumberSemiPersistentCSI_PerBWP_ForBeamReport
	var tmp_int_SimultaneousCSI_ReportsPerCC int64
	if tmp_int_SimultaneousCSI_ReportsPerCC, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger SimultaneousCSI_ReportsPerCC", err)
	}
	ie.SimultaneousCSI_ReportsPerCC = tmp_int_SimultaneousCSI_ReportsPerCC
	return nil
}
