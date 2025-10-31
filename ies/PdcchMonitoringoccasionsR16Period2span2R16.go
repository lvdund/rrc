package ies

import "rrc/utils"

// PDCCH-MonitoringOccasions-r16-period2span2-r16 ::= ENUMERATED
type PdcchMonitoringoccasionsR16Period2span2R16 struct {
	Value utils.ENUMERATED
}

const (
	PdcchMonitoringoccasionsR16Period2span2R16EnumeratedNothing = iota
	PdcchMonitoringoccasionsR16Period2span2R16EnumeratedSupported
)
