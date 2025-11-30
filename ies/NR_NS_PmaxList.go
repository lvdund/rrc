package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_NS_PmaxList struct {
	Value []NR_NS_PmaxValue `lb:1,ub:maxNR_NS_Pmax,madatory`
}

func (ie *NR_NS_PmaxList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*NR_NS_PmaxValue]([]*NR_NS_PmaxValue{}, aper.Constraint{Lb: 1, Ub: maxNR_NS_Pmax}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NR_NS_PmaxList", err)
	}
	return nil
}

func (ie *NR_NS_PmaxList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*NR_NS_PmaxValue]([]*NR_NS_PmaxValue{}, aper.Constraint{Lb: 1, Ub: maxNR_NS_Pmax}, false)
	fn := func() *NR_NS_PmaxValue {
		return new(NR_NS_PmaxValue)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NR_NS_PmaxList", err)
	}
	ie.Value = []NR_NS_PmaxValue{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
