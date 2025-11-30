package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16_Enum_n35us  aper.Enumerated = 0
	ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16_Enum_n140us aper.Enumerated = 1
	ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16_Enum_n210us aper.Enumerated = 2
)

type ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16", err)
	}
	return nil
}

func (ie *ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
