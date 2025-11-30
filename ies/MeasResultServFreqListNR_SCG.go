package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServFreqListNR_SCG struct {
	Value []MeasResult2NR `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *MeasResultServFreqListNR_SCG) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2NR]([]*MeasResult2NR{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultServFreqListNR_SCG", err)
	}
	return nil
}

func (ie *MeasResultServFreqListNR_SCG) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2NR]([]*MeasResult2NR{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *MeasResult2NR {
		return new(MeasResult2NR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultServFreqListNR_SCG", err)
	}
	ie.Value = []MeasResult2NR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
