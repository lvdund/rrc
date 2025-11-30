package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarConnEstFailReportList_r17 struct {
	ConnEstFailReportList_r17 []VarConnEstFailReport_r16 `lb:1,ub:maxCEFReport_r17,madatory`
}

func (ie *VarConnEstFailReportList_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp_ConnEstFailReportList_r17 := utils.NewSequence[*VarConnEstFailReport_r16]([]*VarConnEstFailReport_r16{}, aper.Constraint{Lb: 1, Ub: maxCEFReport_r17}, false)
	for _, i := range ie.ConnEstFailReportList_r17 {
		tmp_ConnEstFailReportList_r17.Value = append(tmp_ConnEstFailReportList_r17.Value, &i)
	}
	if err = tmp_ConnEstFailReportList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ConnEstFailReportList_r17", err)
	}
	return nil
}

func (ie *VarConnEstFailReportList_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp_ConnEstFailReportList_r17 := utils.NewSequence[*VarConnEstFailReport_r16]([]*VarConnEstFailReport_r16{}, aper.Constraint{Lb: 1, Ub: maxCEFReport_r17}, false)
	fn_ConnEstFailReportList_r17 := func() *VarConnEstFailReport_r16 {
		return new(VarConnEstFailReport_r16)
	}
	if err = tmp_ConnEstFailReportList_r17.Decode(r, fn_ConnEstFailReportList_r17); err != nil {
		return utils.WrapError("Decode ConnEstFailReportList_r17", err)
	}
	ie.ConnEstFailReportList_r17 = []VarConnEstFailReport_r16{}
	for _, i := range tmp_ConnEstFailReportList_r17.Value {
		ie.ConnEstFailReportList_r17 = append(ie.ConnEstFailReportList_r17, *i)
	}
	return nil
}
