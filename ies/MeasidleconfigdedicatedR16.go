package ies

// MeasIdleConfigDedicated-r16 ::= SEQUENCE
// Extensible
type MeasidleconfigdedicatedR16 struct {
	MeasidlecarrierlistnrR16    *[]MeasidlecarriernrR16    `lb:1,ub:maxFreqIdleR16`
	MeasidlecarrierlisteutraR16 *[]MeasidlecarriereutraR16 `lb:1,ub:maxFreqIdleR16`
	MeasidledurationR16         MeasidleconfigdedicatedR16MeasidledurationR16
	ValidityarealistR16         *ValidityarealistR16
}
