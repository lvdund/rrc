package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ReportConfig_r16_sl_ReportType_r16_Choice_nothing uint64 = iota
	SL_ReportConfig_r16_sl_ReportType_r16_Choice_Sl_Periodical_r16
	SL_ReportConfig_r16_sl_ReportType_r16_Choice_Sl_EventTriggered_r16
)

type SL_ReportConfig_r16_sl_ReportType_r16 struct {
	Choice                uint64
	Sl_Periodical_r16     *SL_PeriodicalReportConfig_r16
	Sl_EventTriggered_r16 *SL_EventTriggerConfig_r16
}

func (ie *SL_ReportConfig_r16_sl_ReportType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_Sl_Periodical_r16:
		if err = ie.Sl_Periodical_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Sl_Periodical_r16", err)
		}
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_Sl_EventTriggered_r16:
		if err = ie.Sl_EventTriggered_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Sl_EventTriggered_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_ReportConfig_r16_sl_ReportType_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_Sl_Periodical_r16:
		ie.Sl_Periodical_r16 = new(SL_PeriodicalReportConfig_r16)
		if err = ie.Sl_Periodical_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Periodical_r16", err)
		}
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_Sl_EventTriggered_r16:
		ie.Sl_EventTriggered_r16 = new(SL_EventTriggerConfig_r16)
		if err = ie.Sl_EventTriggered_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_EventTriggered_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
