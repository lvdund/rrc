package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationInfoList struct {
	Value []BandCombinationInfo `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationInfoList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombinationInfo]([]*BandCombinationInfo{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationInfoList", err)
	}
	return nil
}

func (ie *BandCombinationInfoList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombinationInfo]([]*BandCombinationInfo{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombinationInfo {
		return new(BandCombinationInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationInfoList", err)
	}
	ie.Value = []BandCombinationInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
