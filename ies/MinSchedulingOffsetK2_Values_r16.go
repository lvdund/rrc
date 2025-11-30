package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetK2_Values_r16 struct {
	Value []int64 `lb:1,ub:maxNrOfMinSchedulingOffsetValues_r16,madatory`
}

func (ie *MinSchedulingOffsetK2_Values_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrOfMinSchedulingOffsetValues_r16}, false)
	for _, i := range ie.Value {
		tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 0}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MinSchedulingOffsetK2_Values_r16", err)
	}
	return nil
}

func (ie *MinSchedulingOffsetK2_Values_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrOfMinSchedulingOffsetValues_r16}, false)
	fn := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 0}, false)
		return &ie
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MinSchedulingOffsetK2_Values_r16", err)
	}
	ie.Value = []int64{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, int64(i.Value))
	}
	return err
}
