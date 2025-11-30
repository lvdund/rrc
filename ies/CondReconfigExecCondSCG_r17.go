package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CondReconfigExecCondSCG_r17 struct {
	Value []MeasId `lb:1,ub:2,madatory`
}

func (ie *CondReconfigExecCondSCG_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasId]([]*MeasId{}, aper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CondReconfigExecCondSCG_r17", err)
	}
	return nil
}

func (ie *CondReconfigExecCondSCG_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasId]([]*MeasId{}, aper.Constraint{Lb: 1, Ub: 2}, false)
	fn := func() *MeasId {
		return new(MeasId)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CondReconfigExecCondSCG_r17", err)
	}
	ie.Value = []MeasId{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
