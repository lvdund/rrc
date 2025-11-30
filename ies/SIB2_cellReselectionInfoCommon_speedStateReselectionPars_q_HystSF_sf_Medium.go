package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium_Enum_dB_6 aper.Enumerated = 0
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium_Enum_dB_4 aper.Enumerated = 1
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium_Enum_dB_2 aper.Enumerated = 2
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium_Enum_dB0  aper.Enumerated = 3
)

type SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium", err)
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
