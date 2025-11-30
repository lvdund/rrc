package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IntraFreqNeighCellList struct {
	Value []IntraFreqNeighCellInfo `lb:1,ub:maxCellIntra,madatory`
}

func (ie *IntraFreqNeighCellList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo]([]*IntraFreqNeighCellInfo{}, aper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreqNeighCellList", err)
	}
	return nil
}

func (ie *IntraFreqNeighCellList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo]([]*IntraFreqNeighCellInfo{}, aper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	fn := func() *IntraFreqNeighCellInfo {
		return new(IntraFreqNeighCellInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode IntraFreqNeighCellList", err)
	}
	ie.Value = []IntraFreqNeighCellInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
