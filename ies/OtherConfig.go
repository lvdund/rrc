package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig struct {
	DelayBudgetReportingConfig *OtherConfig_delayBudgetReportingConfig `optional`
}

func (ie *OtherConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DelayBudgetReportingConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DelayBudgetReportingConfig != nil {
		if err = ie.DelayBudgetReportingConfig.Encode(w); err != nil {
			return utils.WrapError("Encode DelayBudgetReportingConfig", err)
		}
	}
	return nil
}

func (ie *OtherConfig) Decode(r *aper.AperReader) error {
	var err error
	var DelayBudgetReportingConfigPresent bool
	if DelayBudgetReportingConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DelayBudgetReportingConfigPresent {
		ie.DelayBudgetReportingConfig = new(OtherConfig_delayBudgetReportingConfig)
		if err = ie.DelayBudgetReportingConfig.Decode(r); err != nil {
			return utils.WrapError("Decode DelayBudgetReportingConfig", err)
		}
	}
	return nil
}
