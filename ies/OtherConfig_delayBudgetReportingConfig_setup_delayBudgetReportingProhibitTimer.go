package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s0     aper.Enumerated = 0
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s0dot4 aper.Enumerated = 1
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s0dot8 aper.Enumerated = 2
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s1dot6 aper.Enumerated = 3
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s3     aper.Enumerated = 4
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s6     aper.Enumerated = 5
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s12    aper.Enumerated = 6
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s30    aper.Enumerated = 7
)

type OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer", err)
	}
	return nil
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
