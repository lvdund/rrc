package ies

import "rrc/utils"

// Phy-ParametersFR1-pdcch-MonitoringSingleSpanFirst4Sym-r16 ::= ENUMERATED
type PhyParametersfr1PdcchMonitoringsinglespanfirst4symR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr1PdcchMonitoringsinglespanfirst4symR16EnumeratedNothing = iota
	PhyParametersfr1PdcchMonitoringsinglespanfirst4symR16EnumeratedSupported
)
