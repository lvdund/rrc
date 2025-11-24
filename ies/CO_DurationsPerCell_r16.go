package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CO_DurationsPerCell_r16 struct {
	ServingCellId_r16     ServCellIndex     `madatory`
	PositionInDCI_r16     int64             `lb:0,ub:maxSFI_DCI_PayloadSize_1,madatory`
	SubcarrierSpacing_r16 SubcarrierSpacing `madatory`
	Co_DurationList_r16   []CO_Duration_r16 `lb:1,ub:64,madatory`
}

func (ie *CO_DurationsPerCell_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ServingCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ServingCellId_r16", err)
	}
	if err = w.WriteInteger(ie.PositionInDCI_r16, &uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("WriteInteger PositionInDCI_r16", err)
	}
	if err = ie.SubcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing_r16", err)
	}
	tmp_Co_DurationList_r16 := utils.NewSequence[*CO_Duration_r16]([]*CO_Duration_r16{}, uper.Constraint{Lb: 1, Ub: 64}, false)
	for _, i := range ie.Co_DurationList_r16 {
		tmp_Co_DurationList_r16.Value = append(tmp_Co_DurationList_r16.Value, &i)
	}
	if err = tmp_Co_DurationList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Co_DurationList_r16", err)
	}
	return nil
}

func (ie *CO_DurationsPerCell_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ServingCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ServingCellId_r16", err)
	}
	var tmp_int_PositionInDCI_r16 int64
	if tmp_int_PositionInDCI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("ReadInteger PositionInDCI_r16", err)
	}
	ie.PositionInDCI_r16 = tmp_int_PositionInDCI_r16
	if err = ie.SubcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing_r16", err)
	}
	tmp_Co_DurationList_r16 := utils.NewSequence[*CO_Duration_r16]([]*CO_Duration_r16{}, uper.Constraint{Lb: 1, Ub: 64}, false)
	fn_Co_DurationList_r16 := func() *CO_Duration_r16 {
		return new(CO_Duration_r16)
	}
	if err = tmp_Co_DurationList_r16.Decode(r, fn_Co_DurationList_r16); err != nil {
		return utils.WrapError("Decode Co_DurationList_r16", err)
	}
	ie.Co_DurationList_r16 = []CO_Duration_r16{}
	for _, i := range tmp_Co_DurationList_r16.Value {
		ie.Co_DurationList_r16 = append(ie.Co_DurationList_r16, *i)
	}
	return nil
}
