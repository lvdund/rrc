package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportConfigNR_reportType_Choice_nothing uint64 = iota
	ReportConfigNR_reportType_Choice_Periodical
	ReportConfigNR_reportType_Choice_EventTriggered
	ReportConfigNR_reportType_Choice_ReportCGI
	ReportConfigNR_reportType_Choice_ReportSFTD
	ReportConfigNR_reportType_Choice_CondTriggerConfig_r16
	ReportConfigNR_reportType_Choice_Cli_Periodical_r16
	ReportConfigNR_reportType_Choice_Cli_EventTriggered_r16
	ReportConfigNR_reportType_Choice_RxTxPeriodical_r17
)

type ReportConfigNR_reportType struct {
	Choice                 uint64
	Periodical             *PeriodicalReportConfig
	EventTriggered         *EventTriggerConfig
	ReportCGI              *ReportCGI
	ReportSFTD             *ReportSFTD_NR
	CondTriggerConfig_r16  *CondTriggerConfig_r16
	Cli_Periodical_r16     *CLI_PeriodicalReportConfig_r16
	Cli_EventTriggered_r16 *CLI_EventTriggerConfig_r16
	RxTxPeriodical_r17     *RxTxPeriodical_r17
}

func (ie *ReportConfigNR_reportType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigNR_reportType_Choice_Periodical:
		if err = ie.Periodical.Encode(w); err != nil {
			err = utils.WrapError("Encode Periodical", err)
		}
	case ReportConfigNR_reportType_Choice_EventTriggered:
		if err = ie.EventTriggered.Encode(w); err != nil {
			err = utils.WrapError("Encode EventTriggered", err)
		}
	case ReportConfigNR_reportType_Choice_ReportCGI:
		if err = ie.ReportCGI.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportCGI", err)
		}
	case ReportConfigNR_reportType_Choice_ReportSFTD:
		if err = ie.ReportSFTD.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportSFTD", err)
		}
	case ReportConfigNR_reportType_Choice_CondTriggerConfig_r16:
		if err = ie.CondTriggerConfig_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode CondTriggerConfig_r16", err)
		}
	case ReportConfigNR_reportType_Choice_Cli_Periodical_r16:
		if err = ie.Cli_Periodical_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Cli_Periodical_r16", err)
		}
	case ReportConfigNR_reportType_Choice_Cli_EventTriggered_r16:
		if err = ie.Cli_EventTriggered_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Cli_EventTriggered_r16", err)
		}
	case ReportConfigNR_reportType_Choice_RxTxPeriodical_r17:
		if err = ie.RxTxPeriodical_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode RxTxPeriodical_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReportConfigNR_reportType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigNR_reportType_Choice_Periodical:
		ie.Periodical = new(PeriodicalReportConfig)
		if err = ie.Periodical.Decode(r); err != nil {
			return utils.WrapError("Decode Periodical", err)
		}
	case ReportConfigNR_reportType_Choice_EventTriggered:
		ie.EventTriggered = new(EventTriggerConfig)
		if err = ie.EventTriggered.Decode(r); err != nil {
			return utils.WrapError("Decode EventTriggered", err)
		}
	case ReportConfigNR_reportType_Choice_ReportCGI:
		ie.ReportCGI = new(ReportCGI)
		if err = ie.ReportCGI.Decode(r); err != nil {
			return utils.WrapError("Decode ReportCGI", err)
		}
	case ReportConfigNR_reportType_Choice_ReportSFTD:
		ie.ReportSFTD = new(ReportSFTD_NR)
		if err = ie.ReportSFTD.Decode(r); err != nil {
			return utils.WrapError("Decode ReportSFTD", err)
		}
	case ReportConfigNR_reportType_Choice_CondTriggerConfig_r16:
		ie.CondTriggerConfig_r16 = new(CondTriggerConfig_r16)
		if err = ie.CondTriggerConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondTriggerConfig_r16", err)
		}
	case ReportConfigNR_reportType_Choice_Cli_Periodical_r16:
		ie.Cli_Periodical_r16 = new(CLI_PeriodicalReportConfig_r16)
		if err = ie.Cli_Periodical_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Cli_Periodical_r16", err)
		}
	case ReportConfigNR_reportType_Choice_Cli_EventTriggered_r16:
		ie.Cli_EventTriggered_r16 = new(CLI_EventTriggerConfig_r16)
		if err = ie.Cli_EventTriggered_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Cli_EventTriggered_r16", err)
		}
	case ReportConfigNR_reportType_Choice_RxTxPeriodical_r17:
		ie.RxTxPeriodical_r17 = new(RxTxPeriodical_r17)
		if err = ie.RxTxPeriodical_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RxTxPeriodical_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
