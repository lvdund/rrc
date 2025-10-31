package ies

import "rrc/utils"

// EIMTA-MainConfig-r12-setup ::= SEQUENCE
type EimtaMainconfigR12Setup struct {
	EimtaRntiR12               CRnti
	EimtaCommandperiodicityR12 EimtaMainconfigR12SetupEimtaCommandperiodicityR12
	EimtaCommandsubframesetR12 utils.BITSTRING `lb:10,ub:10`
}
