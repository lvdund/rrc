package ies

import "rrc/utils"

// FeatureSetDownlink-v1610-pdcch-MonitoringMixed-r16 ::= ENUMERATED
type FeaturesetdownlinkV1610PdcchMonitoringmixedR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1610PdcchMonitoringmixedR16EnumeratedNothing = iota
	FeaturesetdownlinkV1610PdcchMonitoringmixedR16EnumeratedSupported
)
