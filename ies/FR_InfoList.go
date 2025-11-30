package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FR_InfoList struct {
	Value []FR_Info `lb:1,ub:maxNrofServingCells_1,madatory`
}

func (ie *FR_InfoList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*FR_Info]([]*FR_Info{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FR_InfoList", err)
	}
	return nil
}

func (ie *FR_InfoList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*FR_Info]([]*FR_Info{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
	fn := func() *FR_Info {
		return new(FR_Info)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FR_InfoList", err)
	}
	ie.Value = []FR_Info{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
