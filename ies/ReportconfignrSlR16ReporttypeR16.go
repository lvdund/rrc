package ies

// ReportConfigNR-SL-r16-reportType-r16 ::= CHOICE
const (
	ReportconfignrSlR16ReporttypeR16ChoiceNothing = iota
	ReportconfignrSlR16ReporttypeR16ChoicePeriodicalR16
	ReportconfignrSlR16ReporttypeR16ChoiceEventtriggeredR16
)

type ReportconfignrSlR16ReporttypeR16 struct {
	Choice            uint64
	PeriodicalR16     *PeriodicalreportconfignrSlR16
	EventtriggeredR16 *EventtriggerconfignrSlR16
}
