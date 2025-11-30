package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym14  aper.Enumerated = 0
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym28  aper.Enumerated = 1
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym48  aper.Enumerated = 2
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym224 aper.Enumerated = 3
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym336 aper.Enumerated = 4
)

type MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
