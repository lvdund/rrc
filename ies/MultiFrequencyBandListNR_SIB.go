package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MultiFrequencyBandListNR_SIB struct {
	Value []NR_MultiBandInfo `lb:1,ub:maxNrofMultiBands,madatory`
}

func (ie *MultiFrequencyBandListNR_SIB) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*NR_MultiBandInfo]([]*NR_MultiBandInfo{}, aper.Constraint{Lb: 1, Ub: maxNrofMultiBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MultiFrequencyBandListNR_SIB", err)
	}
	return nil
}

func (ie *MultiFrequencyBandListNR_SIB) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*NR_MultiBandInfo]([]*NR_MultiBandInfo{}, aper.Constraint{Lb: 1, Ub: maxNrofMultiBands}, false)
	fn := func() *NR_MultiBandInfo {
		return new(NR_MultiBandInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MultiFrequencyBandListNR_SIB", err)
	}
	ie.Value = []NR_MultiBandInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
