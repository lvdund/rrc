package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportConfigInterRAT_reportType_Choice_nothing uint64 = iota
	ReportConfigInterRAT_reportType_Choice_Periodical
	ReportConfigInterRAT_reportType_Choice_EventTriggered
	ReportConfigInterRAT_reportType_Choice_ReportCGI
	ReportConfigInterRAT_reportType_Choice_ReportSFTD
)

type ReportConfigInterRAT_reportType struct {
	Choice         uint64
	Periodical     *PeriodicalReportConfigInterRAT
	EventTriggered *EventTriggerConfigInterRAT
	ReportCGI      *ReportCGI_EUTRA
	ReportSFTD     *ReportSFTD_EUTRA
}

func (ie *ReportConfigInterRAT_reportType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigInterRAT_reportType_Choice_Periodical:
		if err = ie.Periodical.Encode(w); err != nil {
			err = utils.WrapError("Encode Periodical", err)
		}
	case ReportConfigInterRAT_reportType_Choice_EventTriggered:
		if err = ie.EventTriggered.Encode(w); err != nil {
			err = utils.WrapError("Encode EventTriggered", err)
		}
	case ReportConfigInterRAT_reportType_Choice_ReportCGI:
		if err = ie.ReportCGI.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportCGI", err)
		}
	case ReportConfigInterRAT_reportType_Choice_ReportSFTD:
		if err = ie.ReportSFTD.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportSFTD", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReportConfigInterRAT_reportType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigInterRAT_reportType_Choice_Periodical:
		ie.Periodical = new(PeriodicalReportConfigInterRAT)
		if err = ie.Periodical.Decode(r); err != nil {
			return utils.WrapError("Decode Periodical", err)
		}
	case ReportConfigInterRAT_reportType_Choice_EventTriggered:
		ie.EventTriggered = new(EventTriggerConfigInterRAT)
		if err = ie.EventTriggered.Decode(r); err != nil {
			return utils.WrapError("Decode EventTriggered", err)
		}
	case ReportConfigInterRAT_reportType_Choice_ReportCGI:
		ie.ReportCGI = new(ReportCGI_EUTRA)
		if err = ie.ReportCGI.Decode(r); err != nil {
			return utils.WrapError("Decode ReportCGI", err)
		}
	case ReportConfigInterRAT_reportType_Choice_ReportSFTD:
		ie.ReportSFTD = new(ReportSFTD_EUTRA)
		if err = ie.ReportSFTD.Decode(r); err != nil {
			return utils.WrapError("Decode ReportSFTD", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
