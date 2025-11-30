package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_oneEighth aper.Enumerated = 0
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_oneFourth aper.Enumerated = 1
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_oneHalf   aper.Enumerated = 2
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_one       aper.Enumerated = 3
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_two       aper.Enumerated = 4
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_four      aper.Enumerated = 5
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_eight     aper.Enumerated = 6
	BeamFailureRecoveryConfig_ssb_perRACH_Occasion_Enum_sixteen   aper.Enumerated = 7
)

type BeamFailureRecoveryConfig_ssb_perRACH_Occasion struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BeamFailureRecoveryConfig_ssb_perRACH_Occasion) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BeamFailureRecoveryConfig_ssb_perRACH_Occasion", err)
	}
	return nil
}

func (ie *BeamFailureRecoveryConfig_ssb_perRACH_Occasion) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BeamFailureRecoveryConfig_ssb_perRACH_Occasion", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
