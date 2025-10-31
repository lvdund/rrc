package ies

// SL-EventTriggerConfig-r16 ::= SEQUENCE
// Extensible
type SlEventtriggerconfigR16 struct {
	SlEventidR16        SlEventtriggerconfigR16SlEventidR16
	SlReportintervalR16 Reportinterval
	SlReportamountR16   SlEventtriggerconfigR16SlReportamountR16
	SlReportquantityR16 SlMeasreportquantityR16
	SlRsTypeR16         SlRsTypeR16
}
