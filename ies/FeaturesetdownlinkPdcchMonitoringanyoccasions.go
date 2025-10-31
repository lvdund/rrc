package ies

import "rrc/utils"

// FeatureSetDownlink-pdcch-MonitoringAnyOccasions ::= ENUMERATED
type FeaturesetdownlinkPdcchMonitoringanyoccasions struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkPdcchMonitoringanyoccasionsEnumeratedNothing = iota
	FeaturesetdownlinkPdcchMonitoringanyoccasionsEnumeratedWithoutdci_Gap
	FeaturesetdownlinkPdcchMonitoringanyoccasionsEnumeratedWithdci_Gap
)
