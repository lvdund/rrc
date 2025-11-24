package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	VarLogMeasConfig_r16_IEs_reportType_Choice_nothing uint64 = iota
	VarLogMeasConfig_r16_IEs_reportType_Choice_Periodical
	VarLogMeasConfig_r16_IEs_reportType_Choice_EventTriggered
)

type VarLogMeasConfig_r16_IEs_reportType struct {
	Choice         uint64
	Periodical     *LoggedPeriodicalReportConfig_r16
	EventTriggered *LoggedEventTriggerConfig_r16
}

func (ie *VarLogMeasConfig_r16_IEs_reportType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VarLogMeasConfig_r16_IEs_reportType_Choice_Periodical:
		if err = ie.Periodical.Encode(w); err != nil {
			err = utils.WrapError("Encode Periodical", err)
		}
	case VarLogMeasConfig_r16_IEs_reportType_Choice_EventTriggered:
		if err = ie.EventTriggered.Encode(w); err != nil {
			err = utils.WrapError("Encode EventTriggered", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *VarLogMeasConfig_r16_IEs_reportType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VarLogMeasConfig_r16_IEs_reportType_Choice_Periodical:
		ie.Periodical = new(LoggedPeriodicalReportConfig_r16)
		if err = ie.Periodical.Decode(r); err != nil {
			return utils.WrapError("Decode Periodical", err)
		}
	case VarLogMeasConfig_r16_IEs_reportType_Choice_EventTriggered:
		ie.EventTriggered = new(LoggedEventTriggerConfig_r16)
		if err = ie.EventTriggered.Decode(r); err != nil {
			return utils.WrapError("Decode EventTriggered", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
