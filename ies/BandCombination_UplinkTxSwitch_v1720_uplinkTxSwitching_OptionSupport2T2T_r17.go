package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17_Enum_switchedUL aper.Enumerated = 0
	BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17_Enum_dualUL     aper.Enumerated = 1
	BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17_Enum_both       aper.Enumerated = 2
)

type BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17", err)
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
