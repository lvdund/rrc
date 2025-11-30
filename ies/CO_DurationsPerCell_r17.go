package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CO_DurationsPerCell_r17 struct {
	ServingCellId_r17     ServCellIndex     `madatory`
	PositionInDCI_r17     int64             `lb:0,ub:maxSFI_DCI_PayloadSize_1,madatory`
	SubcarrierSpacing_r17 SubcarrierSpacing `madatory`
	Co_DurationList_r17   []CO_Duration_r17 `lb:1,ub:64,madatory`
}

func (ie *CO_DurationsPerCell_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ServingCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ServingCellId_r17", err)
	}
	if err = w.WriteInteger(ie.PositionInDCI_r17, &aper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("WriteInteger PositionInDCI_r17", err)
	}
	if err = ie.SubcarrierSpacing_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing_r17", err)
	}
	tmp_Co_DurationList_r17 := utils.NewSequence[*CO_Duration_r17]([]*CO_Duration_r17{}, aper.Constraint{Lb: 1, Ub: 64}, false)
	for _, i := range ie.Co_DurationList_r17 {
		tmp_Co_DurationList_r17.Value = append(tmp_Co_DurationList_r17.Value, &i)
	}
	if err = tmp_Co_DurationList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Co_DurationList_r17", err)
	}
	return nil
}

func (ie *CO_DurationsPerCell_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ServingCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ServingCellId_r17", err)
	}
	var tmp_int_PositionInDCI_r17 int64
	if tmp_int_PositionInDCI_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("ReadInteger PositionInDCI_r17", err)
	}
	ie.PositionInDCI_r17 = tmp_int_PositionInDCI_r17
	if err = ie.SubcarrierSpacing_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing_r17", err)
	}
	tmp_Co_DurationList_r17 := utils.NewSequence[*CO_Duration_r17]([]*CO_Duration_r17{}, aper.Constraint{Lb: 1, Ub: 64}, false)
	fn_Co_DurationList_r17 := func() *CO_Duration_r17 {
		return new(CO_Duration_r17)
	}
	if err = tmp_Co_DurationList_r17.Decode(r, fn_Co_DurationList_r17); err != nil {
		return utils.WrapError("Decode Co_DurationList_r17", err)
	}
	ie.Co_DurationList_r17 = []CO_Duration_r17{}
	for _, i := range tmp_Co_DurationList_r17.Value {
		ie.Co_DurationList_r17 = append(ie.Co_DurationList_r17, *i)
	}
	return nil
}
