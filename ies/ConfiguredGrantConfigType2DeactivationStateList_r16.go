package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConfiguredGrantConfigType2DeactivationStateList_r16 struct {
	Value []ConfiguredGrantConfigType2DeactivationState_r16 `lb:1,ub:maxNrofCG_Type2DeactivationState,madatory`
}

func (ie *ConfiguredGrantConfigType2DeactivationStateList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*ConfiguredGrantConfigType2DeactivationState_r16]([]*ConfiguredGrantConfigType2DeactivationState_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofCG_Type2DeactivationState}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfigType2DeactivationStateList_r16", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfigType2DeactivationStateList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*ConfiguredGrantConfigType2DeactivationState_r16]([]*ConfiguredGrantConfigType2DeactivationState_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofCG_Type2DeactivationState}, false)
	fn := func() *ConfiguredGrantConfigType2DeactivationState_r16 {
		return new(ConfiguredGrantConfigType2DeactivationState_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfigType2DeactivationStateList_r16", err)
	}
	ie.Value = []ConfiguredGrantConfigType2DeactivationState_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
