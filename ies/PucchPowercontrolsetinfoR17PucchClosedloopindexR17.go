package ies

import "rrc/utils"

// PUCCH-PowerControlSetInfo-r17-pucch-ClosedLoopIndex-r17 ::= ENUMERATED
type PucchPowercontrolsetinfoR17PucchClosedloopindexR17 struct {
	Value utils.ENUMERATED
}

const (
	PucchPowercontrolsetinfoR17PucchClosedloopindexR17EnumeratedNothing = iota
	PucchPowercontrolsetinfoR17PucchClosedloopindexR17EnumeratedI0
	PucchPowercontrolsetinfoR17PucchClosedloopindexR17EnumeratedI1
)
