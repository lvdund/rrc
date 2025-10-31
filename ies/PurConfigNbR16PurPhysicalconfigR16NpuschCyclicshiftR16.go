package ies

import "rrc/utils"

// PUR-Config-NB-r16-pur-PhysicalConfig-r16-npusch-CyclicShift-r16 ::= ENUMERATED
type PurConfigNbR16PurPhysicalconfigR16NpuschCyclicshiftR16 struct {
	Value utils.ENUMERATED
}

const (
	PurConfigNbR16PurPhysicalconfigR16NpuschCyclicshiftR16EnumeratedNothing = iota
	PurConfigNbR16PurPhysicalconfigR16NpuschCyclicshiftR16EnumeratedN0
	PurConfigNbR16PurPhysicalconfigR16NpuschCyclicshiftR16EnumeratedN6
)
