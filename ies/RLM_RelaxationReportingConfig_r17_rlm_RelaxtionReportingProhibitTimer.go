package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s0       aper.Enumerated = 0
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s0dot5   aper.Enumerated = 1
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s1       aper.Enumerated = 2
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s2       aper.Enumerated = 3
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s5       aper.Enumerated = 4
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s10      aper.Enumerated = 5
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s20      aper.Enumerated = 6
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s30      aper.Enumerated = 7
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s60      aper.Enumerated = 8
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s90      aper.Enumerated = 9
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s120     aper.Enumerated = 10
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s300     aper.Enumerated = 11
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s600     aper.Enumerated = 12
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_infinity aper.Enumerated = 13
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_spare2   aper.Enumerated = 14
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_spare1   aper.Enumerated = 15
)

type RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}

func (ie *RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
