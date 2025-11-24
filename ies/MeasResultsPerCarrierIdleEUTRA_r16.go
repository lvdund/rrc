package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCarrierIdleEUTRA_r16 struct {
	CarrierFreqEUTRA_r16                ARFCN_ValueEUTRA                  `madatory`
	MeasResultsPerCellListIdleEUTRA_r16 []MeasResultsPerCellIdleEUTRA_r16 `lb:1,ub:maxCellMeasIdle_r16,madatory`
}

func (ie *MeasResultsPerCarrierIdleEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CarrierFreqEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreqEUTRA_r16", err)
	}
	tmp_MeasResultsPerCellListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCellIdleEUTRA_r16]([]*MeasResultsPerCellIdleEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	for _, i := range ie.MeasResultsPerCellListIdleEUTRA_r16 {
		tmp_MeasResultsPerCellListIdleEUTRA_r16.Value = append(tmp_MeasResultsPerCellListIdleEUTRA_r16.Value, &i)
	}
	if err = tmp_MeasResultsPerCellListIdleEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultsPerCellListIdleEUTRA_r16", err)
	}
	return nil
}

func (ie *MeasResultsPerCarrierIdleEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CarrierFreqEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreqEUTRA_r16", err)
	}
	tmp_MeasResultsPerCellListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCellIdleEUTRA_r16]([]*MeasResultsPerCellIdleEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	fn_MeasResultsPerCellListIdleEUTRA_r16 := func() *MeasResultsPerCellIdleEUTRA_r16 {
		return new(MeasResultsPerCellIdleEUTRA_r16)
	}
	if err = tmp_MeasResultsPerCellListIdleEUTRA_r16.Decode(r, fn_MeasResultsPerCellListIdleEUTRA_r16); err != nil {
		return utils.WrapError("Decode MeasResultsPerCellListIdleEUTRA_r16", err)
	}
	ie.MeasResultsPerCellListIdleEUTRA_r16 = []MeasResultsPerCellIdleEUTRA_r16{}
	for _, i := range tmp_MeasResultsPerCellListIdleEUTRA_r16.Value {
		ie.MeasResultsPerCellListIdleEUTRA_r16 = append(ie.MeasResultsPerCellListIdleEUTRA_r16, *i)
	}
	return nil
}
