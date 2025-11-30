package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_SL_r16 struct {
	MeasResultListCBR_NR_r16 []MeasResultCBR_NR_r16 `lb:1,ub:maxNrofSL_PoolToMeasureNR_r16,madatory`
}

func (ie *MeasResultNR_SL_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp_MeasResultListCBR_NR_r16 := utils.NewSequence[*MeasResultCBR_NR_r16]([]*MeasResultCBR_NR_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_PoolToMeasureNR_r16}, false)
	for _, i := range ie.MeasResultListCBR_NR_r16 {
		tmp_MeasResultListCBR_NR_r16.Value = append(tmp_MeasResultListCBR_NR_r16.Value, &i)
	}
	if err = tmp_MeasResultListCBR_NR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListCBR_NR_r16", err)
	}
	return nil
}

func (ie *MeasResultNR_SL_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp_MeasResultListCBR_NR_r16 := utils.NewSequence[*MeasResultCBR_NR_r16]([]*MeasResultCBR_NR_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_PoolToMeasureNR_r16}, false)
	fn_MeasResultListCBR_NR_r16 := func() *MeasResultCBR_NR_r16 {
		return new(MeasResultCBR_NR_r16)
	}
	if err = tmp_MeasResultListCBR_NR_r16.Decode(r, fn_MeasResultListCBR_NR_r16); err != nil {
		return utils.WrapError("Decode MeasResultListCBR_NR_r16", err)
	}
	ie.MeasResultListCBR_NR_r16 = []MeasResultCBR_NR_r16{}
	for _, i := range tmp_MeasResultListCBR_NR_r16.Value {
		ie.MeasResultListCBR_NR_r16 = append(ie.MeasResultListCBR_NR_r16, *i)
	}
	return nil
}
