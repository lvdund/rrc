package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RLM_RelaxationReportingConfig_r17 struct {
	Rlm_RelaxtionReportingProhibitTimer RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer `madatory`
}

func (ie *RLM_RelaxationReportingConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Rlm_RelaxtionReportingProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Rlm_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}

func (ie *RLM_RelaxationReportingConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Rlm_RelaxtionReportingProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Rlm_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}
