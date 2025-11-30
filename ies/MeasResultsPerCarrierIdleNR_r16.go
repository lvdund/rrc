package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCarrierIdleNR_r16 struct {
	CarrierFreq_r16                  ARFCN_ValueNR                  `madatory`
	MeasResultsPerCellListIdleNR_r16 []MeasResultsPerCellIdleNR_r16 `lb:1,ub:maxCellMeasIdle_r16,madatory`
}

func (ie *MeasResultsPerCarrierIdleNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	tmp_MeasResultsPerCellListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCellIdleNR_r16]([]*MeasResultsPerCellIdleNR_r16{}, aper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	for _, i := range ie.MeasResultsPerCellListIdleNR_r16 {
		tmp_MeasResultsPerCellListIdleNR_r16.Value = append(tmp_MeasResultsPerCellListIdleNR_r16.Value, &i)
	}
	if err = tmp_MeasResultsPerCellListIdleNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultsPerCellListIdleNR_r16", err)
	}
	return nil
}

func (ie *MeasResultsPerCarrierIdleNR_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	tmp_MeasResultsPerCellListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCellIdleNR_r16]([]*MeasResultsPerCellIdleNR_r16{}, aper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	fn_MeasResultsPerCellListIdleNR_r16 := func() *MeasResultsPerCellIdleNR_r16 {
		return new(MeasResultsPerCellIdleNR_r16)
	}
	if err = tmp_MeasResultsPerCellListIdleNR_r16.Decode(r, fn_MeasResultsPerCellListIdleNR_r16); err != nil {
		return utils.WrapError("Decode MeasResultsPerCellListIdleNR_r16", err)
	}
	ie.MeasResultsPerCellListIdleNR_r16 = []MeasResultsPerCellIdleNR_r16{}
	for _, i := range tmp_MeasResultsPerCellListIdleNR_r16.Value {
		ie.MeasResultsPerCellListIdleNR_r16 = append(ie.MeasResultsPerCellListIdleNR_r16, *i)
	}
	return nil
}
