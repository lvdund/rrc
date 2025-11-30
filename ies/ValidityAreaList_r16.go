package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ValidityAreaList_r16 struct {
	Value []ValidityArea_r16 `lb:1,ub:maxFreqIdle_r16,madatory`
}

func (ie *ValidityAreaList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*ValidityArea_r16]([]*ValidityArea_r16{}, aper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ValidityAreaList_r16", err)
	}
	return nil
}

func (ie *ValidityAreaList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*ValidityArea_r16]([]*ValidityArea_r16{}, aper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
	fn := func() *ValidityArea_r16 {
		return new(ValidityArea_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ValidityAreaList_r16", err)
	}
	ie.Value = []ValidityArea_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
