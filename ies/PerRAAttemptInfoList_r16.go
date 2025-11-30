package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PerRAAttemptInfoList_r16 struct {
	Value []PerRAAttemptInfo_r16 `lb:1,ub:200,madatory`
}

func (ie *PerRAAttemptInfoList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PerRAAttemptInfo_r16]([]*PerRAAttemptInfo_r16{}, aper.Constraint{Lb: 1, Ub: 200}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PerRAAttemptInfoList_r16", err)
	}
	return nil
}

func (ie *PerRAAttemptInfoList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PerRAAttemptInfo_r16]([]*PerRAAttemptInfo_r16{}, aper.Constraint{Lb: 1, Ub: 200}, false)
	fn := func() *PerRAAttemptInfo_r16 {
		return new(PerRAAttemptInfo_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PerRAAttemptInfoList_r16", err)
	}
	ie.Value = []PerRAAttemptInfo_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
