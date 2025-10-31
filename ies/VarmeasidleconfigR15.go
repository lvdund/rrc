package ies

// VarMeasIdleConfig-r15 ::= SEQUENCE
type VarmeasidleconfigR15 struct {
	MeasidlecarrierlisteutraR15 *EutraCarrierlistR15
	MeasidledurationR15         VarmeasidleconfigR15MeasidledurationR15
}
