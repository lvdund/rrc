package ies

import "rrc/utils"

// VarMeasIdleConfig-r16 ::= SEQUENCE
type VarmeasidleconfigR16 struct {
	MeasidlecarrierlistnrR16 *NrCarrierlistR16
	ValidityarealistR16      *ValidityarealistR16
}
