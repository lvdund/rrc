package ies

import "rrc/utils"

// CrossCarrierSchedulingSCell-SpCell-r17-pdcch-MonitoringOccasion-r17 ::= ENUMERATED
type CrosscarrierschedulingscellSpcellR17PdcchMonitoringoccasionR17 struct {
	Value utils.ENUMERATED
}

const (
	CrosscarrierschedulingscellSpcellR17PdcchMonitoringoccasionR17EnumeratedNothing = iota
	CrosscarrierschedulingscellSpcellR17PdcchMonitoringoccasionR17EnumeratedVal1
	CrosscarrierschedulingscellSpcellR17PdcchMonitoringoccasionR17EnumeratedVal2
)
