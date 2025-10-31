package ies

// MeasIdleConfigDedicated-r15 ::= SEQUENCE
// Extensible
type MeasidleconfigdedicatedR15 struct {
	MeasidlecarrierlisteutraR15 *EutraCarrierlistR15
	MeasidledurationR15         MeasidleconfigdedicatedR15MeasidledurationR15
}
