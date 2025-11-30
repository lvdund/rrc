package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqList_v1610 struct {
	Value []InterFreqCarrierFreqInfo_v1610 `lb:1,ub:maxFreq,madatory`
}

func (ie *InterFreqCarrierFreqList_v1610) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1610]([]*InterFreqCarrierFreqInfo_v1610{}, aper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqCarrierFreqList_v1610", err)
	}
	return nil
}

func (ie *InterFreqCarrierFreqList_v1610) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1610]([]*InterFreqCarrierFreqInfo_v1610{}, aper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *InterFreqCarrierFreqInfo_v1610 {
		return new(InterFreqCarrierFreqInfo_v1610)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqCarrierFreqList_v1610", err)
	}
	ie.Value = []InterFreqCarrierFreqInfo_v1610{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
