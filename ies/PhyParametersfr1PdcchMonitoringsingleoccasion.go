package ies

import "rrc/utils"

// Phy-ParametersFR1-pdcch-MonitoringSingleOccasion ::= ENUMERATED
type PhyParametersfr1PdcchMonitoringsingleoccasion struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr1PdcchMonitoringsingleoccasionEnumeratedNothing = iota
	PhyParametersfr1PdcchMonitoringsingleoccasionEnumeratedSupported
)
