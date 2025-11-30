package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportFrameworkExt_r16 struct {
	MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 int64 `lb:5,ub:8,madatory`
}

func (ie *CSI_ReportFrameworkExt_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16, &aper.Constraint{Lb: 5, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16", err)
	}
	return nil
}

func (ie *CSI_ReportFrameworkExt_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 int64
	if tmp_int_MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16, err = r.ReadInteger(&aper.Constraint{Lb: 5, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16", err)
	}
	ie.MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 = tmp_int_MaxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16
	return nil
}
