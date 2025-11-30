package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultList2NR_r16 struct {
	Value []MeasResult2NR_r16 `lb:1,ub:maxFreq,madatory`
}

func (ie *MeasResultList2NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2NR_r16]([]*MeasResult2NR_r16{}, aper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultList2NR_r16", err)
	}
	return nil
}

func (ie *MeasResultList2NR_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2NR_r16]([]*MeasResult2NR_r16{}, aper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *MeasResult2NR_r16 {
		return new(MeasResult2NR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultList2NR_r16", err)
	}
	ie.Value = []MeasResult2NR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
