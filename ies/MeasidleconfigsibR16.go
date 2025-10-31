package ies

// MeasIdleConfigSIB-r16 ::= SEQUENCE
// Extensible
type MeasidleconfigsibR16 struct {
	MeasidlecarrierlistnrR16    *[]MeasidlecarriernrR16    `lb:1,ub:maxFreqIdleR16`
	MeasidlecarrierlisteutraR16 *[]MeasidlecarriereutraR16 `lb:1,ub:maxFreqIdleR16`
}
