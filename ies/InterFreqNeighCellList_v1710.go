package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqNeighCellList_v1710 struct {
	Value []InterFreqNeighCellInfo_v1710 `lb:1,ub:maxCellInter,madatory`
}

func (ie *InterFreqNeighCellList_v1710) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqNeighCellInfo_v1710]([]*InterFreqNeighCellInfo_v1710{}, aper.Constraint{Lb: 1, Ub: maxCellInter}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqNeighCellList_v1710", err)
	}
	return nil
}

func (ie *InterFreqNeighCellList_v1710) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqNeighCellInfo_v1710]([]*InterFreqNeighCellInfo_v1710{}, aper.Constraint{Lb: 1, Ub: maxCellInter}, false)
	fn := func() *InterFreqNeighCellInfo_v1710 {
		return new(InterFreqNeighCellInfo_v1710)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqNeighCellList_v1710", err)
	}
	ie.Value = []InterFreqNeighCellInfo_v1710{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
