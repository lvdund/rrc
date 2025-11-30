package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultIdleEUTRA_r16 struct {
	MeasResultsPerCarrierListIdleEUTRA_r16 []MeasResultsPerCarrierIdleEUTRA_r16 `lb:1,ub:maxFreqIdle_r16,madatory`
}

func (ie *MeasResultIdleEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp_MeasResultsPerCarrierListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleEUTRA_r16]([]*MeasResultsPerCarrierIdleEUTRA_r16{}, aper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
	for _, i := range ie.MeasResultsPerCarrierListIdleEUTRA_r16 {
		tmp_MeasResultsPerCarrierListIdleEUTRA_r16.Value = append(tmp_MeasResultsPerCarrierListIdleEUTRA_r16.Value, &i)
	}
	if err = tmp_MeasResultsPerCarrierListIdleEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultsPerCarrierListIdleEUTRA_r16", err)
	}
	return nil
}

func (ie *MeasResultIdleEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp_MeasResultsPerCarrierListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleEUTRA_r16]([]*MeasResultsPerCarrierIdleEUTRA_r16{}, aper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
	fn_MeasResultsPerCarrierListIdleEUTRA_r16 := func() *MeasResultsPerCarrierIdleEUTRA_r16 {
		return new(MeasResultsPerCarrierIdleEUTRA_r16)
	}
	if err = tmp_MeasResultsPerCarrierListIdleEUTRA_r16.Decode(r, fn_MeasResultsPerCarrierListIdleEUTRA_r16); err != nil {
		return utils.WrapError("Decode MeasResultsPerCarrierListIdleEUTRA_r16", err)
	}
	ie.MeasResultsPerCarrierListIdleEUTRA_r16 = []MeasResultsPerCarrierIdleEUTRA_r16{}
	for _, i := range tmp_MeasResultsPerCarrierListIdleEUTRA_r16.Value {
		ie.MeasResultsPerCarrierListIdleEUTRA_r16 = append(ie.MeasResultsPerCarrierListIdleEUTRA_r16, *i)
	}
	return nil
}
