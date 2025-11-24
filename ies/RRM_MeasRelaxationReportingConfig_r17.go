package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRM_MeasRelaxationReportingConfig_r17 struct {
	S_SearchDeltaP_Stationary_r17 RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17 `madatory`
	T_SearchDeltaP_Stationary_r17 RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17 `madatory`
}

func (ie *RRM_MeasRelaxationReportingConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.S_SearchDeltaP_Stationary_r17.Encode(w); err != nil {
		return utils.WrapError("Encode S_SearchDeltaP_Stationary_r17", err)
	}
	if err = ie.T_SearchDeltaP_Stationary_r17.Encode(w); err != nil {
		return utils.WrapError("Encode T_SearchDeltaP_Stationary_r17", err)
	}
	return nil
}

func (ie *RRM_MeasRelaxationReportingConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.S_SearchDeltaP_Stationary_r17.Decode(r); err != nil {
		return utils.WrapError("Decode S_SearchDeltaP_Stationary_r17", err)
	}
	if err = ie.T_SearchDeltaP_Stationary_r17.Decode(r); err != nil {
		return utils.WrapError("Decode T_SearchDeltaP_Stationary_r17", err)
	}
	return nil
}
