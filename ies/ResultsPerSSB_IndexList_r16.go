package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_IndexList_r16 struct {
	Value []ResultsPerSSB_IndexIdle_r16 `lb:1,ub:maxNrofIndexesToReport,madatory`
}

func (ie *ResultsPerSSB_IndexList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*ResultsPerSSB_IndexIdle_r16]([]*ResultsPerSSB_IndexIdle_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ResultsPerSSB_IndexList_r16", err)
	}
	return nil
}

func (ie *ResultsPerSSB_IndexList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*ResultsPerSSB_IndexIdle_r16]([]*ResultsPerSSB_IndexIdle_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false)
	fn := func() *ResultsPerSSB_IndexIdle_r16 {
		return new(ResultsPerSSB_IndexIdle_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ResultsPerSSB_IndexList_r16", err)
	}
	ie.Value = []ResultsPerSSB_IndexIdle_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
