package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellGrouping_r16 struct {
	Mcg_r16  []FreqBandIndicatorNR     `lb:1,ub:maxBands,madatory`
	Scg_r16  []FreqBandIndicatorNR     `lb:1,ub:maxBands,madatory`
	Mode_r16 CellGrouping_r16_mode_r16 `madatory`
}

func (ie *CellGrouping_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Mcg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.Mcg_r16 {
		tmp_Mcg_r16.Value = append(tmp_Mcg_r16.Value, &i)
	}
	if err = tmp_Mcg_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Mcg_r16", err)
	}
	tmp_Scg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.Scg_r16 {
		tmp_Scg_r16.Value = append(tmp_Scg_r16.Value, &i)
	}
	if err = tmp_Scg_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Scg_r16", err)
	}
	if err = ie.Mode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Mode_r16", err)
	}
	return nil
}

func (ie *CellGrouping_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp_Mcg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn_Mcg_r16 := func() *FreqBandIndicatorNR {
		return new(FreqBandIndicatorNR)
	}
	if err = tmp_Mcg_r16.Decode(r, fn_Mcg_r16); err != nil {
		return utils.WrapError("Decode Mcg_r16", err)
	}
	ie.Mcg_r16 = []FreqBandIndicatorNR{}
	for _, i := range tmp_Mcg_r16.Value {
		ie.Mcg_r16 = append(ie.Mcg_r16, *i)
	}
	tmp_Scg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn_Scg_r16 := func() *FreqBandIndicatorNR {
		return new(FreqBandIndicatorNR)
	}
	if err = tmp_Scg_r16.Decode(r, fn_Scg_r16); err != nil {
		return utils.WrapError("Decode Scg_r16", err)
	}
	ie.Scg_r16 = []FreqBandIndicatorNR{}
	for _, i := range tmp_Scg_r16.Value {
		ie.Scg_r16 = append(ie.Scg_r16, *i)
	}
	if err = ie.Mode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Mode_r16", err)
	}
	return nil
}
