package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s5     aper.Enumerated = 0
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s10    aper.Enumerated = 1
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s20    aper.Enumerated = 2
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s30    aper.Enumerated = 3
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s60    aper.Enumerated = 4
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s120   aper.Enumerated = 5
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s180   aper.Enumerated = 6
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s240   aper.Enumerated = 7
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_s300   aper.Enumerated = 8
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_spare7 aper.Enumerated = 9
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_spare6 aper.Enumerated = 10
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_spare5 aper.Enumerated = 11
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_spare4 aper.Enumerated = 12
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_spare3 aper.Enumerated = 13
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_spare2 aper.Enumerated = 14
	SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17_Enum_spare1 aper.Enumerated = 15
)

type SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17", err)
	}
	return nil
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
