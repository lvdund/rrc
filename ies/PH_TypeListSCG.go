package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PH_TypeListSCG struct {
	Value []PH_InfoSCG `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *PH_TypeListSCG) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PH_InfoSCG]([]*PH_InfoSCG{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PH_TypeListSCG", err)
	}
	return nil
}

func (ie *PH_TypeListSCG) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PH_InfoSCG]([]*PH_InfoSCG{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *PH_InfoSCG {
		return new(PH_InfoSCG)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PH_TypeListSCG", err)
	}
	ie.Value = []PH_InfoSCG{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
