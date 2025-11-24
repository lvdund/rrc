package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	OtherConfig_delayBudgetReportingConfig_Choice_nothing uint64 = iota
	OtherConfig_delayBudgetReportingConfig_Choice_Release
	OtherConfig_delayBudgetReportingConfig_Choice_Setup
)

type OtherConfig_delayBudgetReportingConfig struct {
	Choice  uint64
	Release uper.NULL `madatory`
	Setup   *OtherConfig_delayBudgetReportingConfig_setup
}

func (ie *OtherConfig_delayBudgetReportingConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case OtherConfig_delayBudgetReportingConfig_Choice_Release:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Release", err)
		}
	case OtherConfig_delayBudgetReportingConfig_Choice_Setup:
		if err = ie.Setup.Encode(w); err != nil {
			err = utils.WrapError("Encode Setup", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *OtherConfig_delayBudgetReportingConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case OtherConfig_delayBudgetReportingConfig_Choice_Release:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Release", err)
		}
	case OtherConfig_delayBudgetReportingConfig_Choice_Setup:
		ie.Setup = new(OtherConfig_delayBudgetReportingConfig_setup)
		if err = ie.Setup.Decode(r); err != nil {
			return utils.WrapError("Decode Setup", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
