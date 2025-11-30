package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultListNR struct {
	Value []MeasResultNR `lb:1,ub:maxCellReport,madatory`
}

func (ie *MeasResultListNR) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultNR]([]*MeasResultNR{}, aper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListNR", err)
	}
	return nil
}

func (ie *MeasResultListNR) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultNR]([]*MeasResultNR{}, aper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	fn := func() *MeasResultNR {
		return new(MeasResultNR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultListNR", err)
	}
	ie.Value = []MeasResultNR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
