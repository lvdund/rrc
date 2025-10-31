package ies

import "rrc/utils"

// PDCCH-MonitoringOccasions-r16-period4span3-r16 ::= ENUMERATED
type PdcchMonitoringoccasionsR16Period4span3R16 struct {
	Value utils.ENUMERATED
}

const (
	PdcchMonitoringoccasionsR16Period4span3R16EnumeratedNothing = iota
	PdcchMonitoringoccasionsR16Period4span3R16EnumeratedSupported
)
