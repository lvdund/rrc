package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym4  aper.Enumerated = 0
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym8  aper.Enumerated = 1
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym14 aper.Enumerated = 2
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym28 aper.Enumerated = 3
)

type MIMO_ParametersPerBand_beamReportTiming_scs_30kHz struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MIMO_ParametersPerBand_beamReportTiming_scs_30kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_beamReportTiming_scs_30kHz", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamReportTiming_scs_30kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_beamReportTiming_scs_30kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
