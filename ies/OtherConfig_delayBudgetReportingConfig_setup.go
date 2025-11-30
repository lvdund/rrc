package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_delayBudgetReportingConfig_setup struct {
	DelayBudgetReportingProhibitTimer OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer `madatory`
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.DelayBudgetReportingProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode DelayBudgetReportingProhibitTimer", err)
	}
	return nil
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.DelayBudgetReportingProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode DelayBudgetReportingProhibitTimer", err)
	}
	return nil
}
