package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UL_PDCP_DelayValueResultList_r16 struct {
	Value []UL_PDCP_DelayValueResult_r16 `lb:1,ub:maxDRB,madatory`
}

func (ie *UL_PDCP_DelayValueResultList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*UL_PDCP_DelayValueResult_r16]([]*UL_PDCP_DelayValueResult_r16{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UL_PDCP_DelayValueResultList_r16", err)
	}
	return nil
}

func (ie *UL_PDCP_DelayValueResultList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*UL_PDCP_DelayValueResult_r16]([]*UL_PDCP_DelayValueResult_r16{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn := func() *UL_PDCP_DelayValueResult_r16 {
		return new(UL_PDCP_DelayValueResult_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UL_PDCP_DelayValueResultList_r16", err)
	}
	ie.Value = []UL_PDCP_DelayValueResult_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
