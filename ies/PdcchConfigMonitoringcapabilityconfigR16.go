package ies

import "rrc/utils"

// PDCCH-Config-monitoringCapabilityConfig-r16 ::= ENUMERATED
type PdcchConfigMonitoringcapabilityconfigR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcchConfigMonitoringcapabilityconfigR16EnumeratedNothing = iota
	PdcchConfigMonitoringcapabilityconfigR16EnumeratedR15monitoringcapability
	PdcchConfigMonitoringcapabilityconfigR16EnumeratedR16monitoringcapability
)
