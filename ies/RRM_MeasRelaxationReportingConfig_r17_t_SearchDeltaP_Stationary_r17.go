package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s5     aper.Enumerated = 0
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s10    aper.Enumerated = 1
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s20    aper.Enumerated = 2
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s30    aper.Enumerated = 3
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s60    aper.Enumerated = 4
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s120   aper.Enumerated = 5
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s180   aper.Enumerated = 6
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s240   aper.Enumerated = 7
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s300   aper.Enumerated = 8
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare7 aper.Enumerated = 9
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare6 aper.Enumerated = 10
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare5 aper.Enumerated = 11
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare4 aper.Enumerated = 12
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare3 aper.Enumerated = 13
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare2 aper.Enumerated = 14
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare1 aper.Enumerated = 15
)

type RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17", err)
	}
	return nil
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
