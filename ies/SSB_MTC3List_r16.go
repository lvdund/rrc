package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC3List_r16 struct {
	Value []SSB_MTC3_r16 `lb:1,ub:4,madatory`
}

func (ie *SSB_MTC3List_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*SSB_MTC3_r16]([]*SSB_MTC3_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SSB_MTC3List_r16", err)
	}
	return nil
}

func (ie *SSB_MTC3List_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*SSB_MTC3_r16]([]*SSB_MTC3_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
	fn := func() *SSB_MTC3_r16 {
		return new(SSB_MTC3_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SSB_MTC3List_r16", err)
	}
	ie.Value = []SSB_MTC3_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
