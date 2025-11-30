package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms10  aper.Enumerated = 0
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms20  aper.Enumerated = 1
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms40  aper.Enumerated = 2
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms60  aper.Enumerated = 3
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms80  aper.Enumerated = 4
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms100 aper.Enumerated = 5
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms150 aper.Enumerated = 6
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms200 aper.Enumerated = 7
)

type BeamFailureRecoveryConfig_beamFailureRecoveryTimer struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BeamFailureRecoveryConfig_beamFailureRecoveryTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BeamFailureRecoveryConfig_beamFailureRecoveryTimer", err)
	}
	return nil
}

func (ie *BeamFailureRecoveryConfig_beamFailureRecoveryTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BeamFailureRecoveryConfig_beamFailureRecoveryTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
