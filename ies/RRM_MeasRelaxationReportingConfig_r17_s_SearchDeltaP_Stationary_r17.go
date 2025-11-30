package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB2    aper.Enumerated = 0
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB3    aper.Enumerated = 1
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB6    aper.Enumerated = 2
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB9    aper.Enumerated = 3
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB12   aper.Enumerated = 4
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB15   aper.Enumerated = 5
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_spare2 aper.Enumerated = 6
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_spare1 aper.Enumerated = 7
)

type RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17", err)
	}
	return nil
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
