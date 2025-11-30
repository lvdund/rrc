package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IntraBandCC_CombinationReqList_r17 struct {
	ServCellIndexList_r17  []ServCellIndex               `lb:1,ub:maxNrofServingCells,madatory`
	Cc_CombinationList_r17 []IntraBandCC_Combination_r17 `lb:1,ub:maxNrofReqComDC_Location_r17,madatory`
}

func (ie *IntraBandCC_CombinationReqList_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp_ServCellIndexList_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.ServCellIndexList_r17 {
		tmp_ServCellIndexList_r17.Value = append(tmp_ServCellIndexList_r17.Value, &i)
	}
	if err = tmp_ServCellIndexList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellIndexList_r17", err)
	}
	tmp_Cc_CombinationList_r17 := utils.NewSequence[*IntraBandCC_Combination_r17]([]*IntraBandCC_Combination_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofReqComDC_Location_r17}, false)
	for _, i := range ie.Cc_CombinationList_r17 {
		tmp_Cc_CombinationList_r17.Value = append(tmp_Cc_CombinationList_r17.Value, &i)
	}
	if err = tmp_Cc_CombinationList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Cc_CombinationList_r17", err)
	}
	return nil
}

func (ie *IntraBandCC_CombinationReqList_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp_ServCellIndexList_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_ServCellIndexList_r17 := func() *ServCellIndex {
		return new(ServCellIndex)
	}
	if err = tmp_ServCellIndexList_r17.Decode(r, fn_ServCellIndexList_r17); err != nil {
		return utils.WrapError("Decode ServCellIndexList_r17", err)
	}
	ie.ServCellIndexList_r17 = []ServCellIndex{}
	for _, i := range tmp_ServCellIndexList_r17.Value {
		ie.ServCellIndexList_r17 = append(ie.ServCellIndexList_r17, *i)
	}
	tmp_Cc_CombinationList_r17 := utils.NewSequence[*IntraBandCC_Combination_r17]([]*IntraBandCC_Combination_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofReqComDC_Location_r17}, false)
	fn_Cc_CombinationList_r17 := func() *IntraBandCC_Combination_r17 {
		return new(IntraBandCC_Combination_r17)
	}
	if err = tmp_Cc_CombinationList_r17.Decode(r, fn_Cc_CombinationList_r17); err != nil {
		return utils.WrapError("Decode Cc_CombinationList_r17", err)
	}
	ie.Cc_CombinationList_r17 = []IntraBandCC_Combination_r17{}
	for _, i := range tmp_Cc_CombinationList_r17.Value {
		ie.Cc_CombinationList_r17 = append(ie.Cc_CombinationList_r17, *i)
	}
	return nil
}
