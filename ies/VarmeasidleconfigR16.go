package ies

// VarMeasIdleConfig-r16 ::= SEQUENCE
type VarmeasidleconfigR16 struct {
	MeasidlecarrierlistnrR16    *[]MeasidlecarriernrR16    `lb:1,ub:maxFreqIdleR16`
	MeasidlecarrierlisteutraR16 *[]MeasidlecarriereutraR16 `lb:1,ub:maxFreqIdleR16`
	MeasidledurationR16         VarmeasidleconfigR16MeasidledurationR16
	ValidityarealistR16         *ValidityarealistR16
}
