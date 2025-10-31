package ies

import "rrc/utils"

// PDCCH-Config-monitoringCapabilityConfig-v1710 ::= ENUMERATED
type PdcchConfigMonitoringcapabilityconfigV1710 struct {
	Value utils.ENUMERATED
}

const (
	PdcchConfigMonitoringcapabilityconfigV1710EnumeratedNothing = iota
	PdcchConfigMonitoringcapabilityconfigV1710EnumeratedR17monitoringcapability
)
