package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentMoreCarrierList_r17 struct {
	Value []CC_Group_r17 `lb:1,ub:maxNrofCC_Group_r17,madatory`
}

func (ie *UplinkTxDirectCurrentMoreCarrierList_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*CC_Group_r17]([]*CC_Group_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofCC_Group_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkTxDirectCurrentMoreCarrierList_r17", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentMoreCarrierList_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*CC_Group_r17]([]*CC_Group_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofCC_Group_r17}, false)
	fn := func() *CC_Group_r17 {
		return new(CC_Group_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UplinkTxDirectCurrentMoreCarrierList_r17", err)
	}
	ie.Value = []CC_Group_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
