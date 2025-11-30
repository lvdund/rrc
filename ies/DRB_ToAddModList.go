package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRB_ToAddModList struct {
	Value []DRB_ToAddMod `lb:1,ub:maxDRB,madatory`
}

func (ie *DRB_ToAddModList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*DRB_ToAddMod]([]*DRB_ToAddMod{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode DRB_ToAddModList", err)
	}
	return nil
}

func (ie *DRB_ToAddModList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*DRB_ToAddMod]([]*DRB_ToAddMod{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn := func() *DRB_ToAddMod {
		return new(DRB_ToAddMod)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode DRB_ToAddModList", err)
	}
	ie.Value = []DRB_ToAddMod{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
