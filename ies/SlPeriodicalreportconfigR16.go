package ies

// SL-PeriodicalReportConfig-r16 ::= SEQUENCE
// Extensible
type SlPeriodicalreportconfigR16 struct {
	SlReportintervalR16 Reportinterval
	SlReportamountR16   SlPeriodicalreportconfigR16SlReportamountR16
	SlReportquantityR16 SlMeasreportquantityR16
	SlRsTypeR16         SlRsTypeR16
}
