package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_dB3    aper.Enumerated = 0
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_dB6    aper.Enumerated = 1
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_dB9    aper.Enumerated = 2
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_dB12   aper.Enumerated = 3
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_dB15   aper.Enumerated = 4
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_spare3 aper.Enumerated = 5
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_spare2 aper.Enumerated = 6
	SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17_Enum_spare1 aper.Enumerated = 7
)

type SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17", err)
	}
	return nil
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
