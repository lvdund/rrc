package ies

import "rrc/utils"

// PhysicalConfigDedicated-NB-r13 ::= SEQUENCE
// Extensible
type PhysicalconfigdedicatedNbR13 struct {
	CarrierconfigdedicatedR13      *CarrierconfigdedicatedNbR13
	NpdcchConfigdedicatedR13       *NpdcchConfigdedicatedNbR13
	NpuschConfigdedicatedR13       *NpuschConfigdedicatedNbR13
	UplinkpowercontroldedicatedR13 *UplinkpowercontroldedicatedNbR13
}
