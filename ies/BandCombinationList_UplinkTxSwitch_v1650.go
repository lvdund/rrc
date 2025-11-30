package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationList_UplinkTxSwitch_v1650 struct {
	Value []BandCombination_UplinkTxSwitch_v1650 `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationList_UplinkTxSwitch_v1650) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_UplinkTxSwitch_v1650]([]*BandCombination_UplinkTxSwitch_v1650{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationList_UplinkTxSwitch_v1650", err)
	}
	return nil
}

func (ie *BandCombinationList_UplinkTxSwitch_v1650) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_UplinkTxSwitch_v1650]([]*BandCombination_UplinkTxSwitch_v1650{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombination_UplinkTxSwitch_v1650 {
		return new(BandCombination_UplinkTxSwitch_v1650)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationList_UplinkTxSwitch_v1650", err)
	}
	ie.Value = []BandCombination_UplinkTxSwitch_v1650{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
