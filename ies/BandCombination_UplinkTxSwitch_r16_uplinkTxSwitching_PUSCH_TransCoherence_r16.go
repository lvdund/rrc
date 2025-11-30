package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16_Enum_nonCoherent  aper.Enumerated = 0
	BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16_Enum_fullCoherent aper.Enumerated = 1
)

type BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16", err)
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
