package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type HRNN_List_r16 struct {
	Value []HRNN_r16 `lb:1,ub:maxNPN_r16,madatory`
}

func (ie *HRNN_List_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*HRNN_r16]([]*HRNN_r16{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode HRNN_List_r16", err)
	}
	return nil
}

func (ie *HRNN_List_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*HRNN_r16]([]*HRNN_r16{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn := func() *HRNN_r16 {
		return new(HRNN_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode HRNN_List_r16", err)
	}
	ie.Value = []HRNN_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
