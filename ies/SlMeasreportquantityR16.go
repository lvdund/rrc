package ies

import "rrc/utils"

// SL-MeasReportQuantity-r16 ::= CHOICE
// Extensible
const (
	SlMeasreportquantityR16ChoiceNothing = iota
	SlMeasreportquantityR16ChoiceSlRsrpR16
)

type SlMeasreportquantityR16 struct {
	Choice    uint64
	SlRsrpR16 *utils.BOOLEAN
}
