package ies

// SL-ReportConfig-r16-sl-ReportType-r16 ::= CHOICE
// Extensible
const (
	SlReportconfigR16SlReporttypeR16ChoiceNothing = iota
	SlReportconfigR16SlReporttypeR16ChoiceSlPeriodicalR16
	SlReportconfigR16SlReporttypeR16ChoiceSlEventtriggeredR16
)

type SlReportconfigR16SlReporttypeR16 struct {
	Choice              uint64
	SlPeriodicalR16     *SlPeriodicalreportconfigR16
	SlEventtriggeredR16 *SlEventtriggerconfigR16
}
