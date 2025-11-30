package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SpCellConfig_lowMobilityEvaluationConnected_r17 struct {
	S_SearchDeltaP_Connected_r17 SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17 `madatory`
	T_SearchDeltaP_Connected_r17 SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17 `madatory`
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.S_SearchDeltaP_Connected_r17.Encode(w); err != nil {
		return utils.WrapError("Encode S_SearchDeltaP_Connected_r17", err)
	}
	if err = ie.T_SearchDeltaP_Connected_r17.Encode(w); err != nil {
		return utils.WrapError("Encode T_SearchDeltaP_Connected_r17", err)
	}
	return nil
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.S_SearchDeltaP_Connected_r17.Decode(r); err != nil {
		return utils.WrapError("Decode S_SearchDeltaP_Connected_r17", err)
	}
	if err = ie.T_SearchDeltaP_Connected_r17.Decode(r); err != nil {
		return utils.WrapError("Decode T_SearchDeltaP_Connected_r17", err)
	}
	return nil
}
