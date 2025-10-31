package ies

import "rrc/utils"

// SCG-Configuration-r12-setup-scg-ConfigPartMCG-r12 ::= SEQUENCE
// Extensible
type ScgConfigurationR12SetupScgConfigpartmcgR12 struct {
	ScgCounterR12            *utils.INTEGER `lb:0,ub:65535`
	PowercoordinationinfoR12 *PowercoordinationinfoR12
}
