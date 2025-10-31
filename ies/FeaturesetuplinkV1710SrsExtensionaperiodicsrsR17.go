package ies

import "rrc/utils"

// FeatureSetUplink-v1710-srs-ExtensionAperiodicSRS-r17 ::= ENUMERATED
type FeaturesetuplinkV1710SrsExtensionaperiodicsrsR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1710SrsExtensionaperiodicsrsR17EnumeratedNothing = iota
	FeaturesetuplinkV1710SrsExtensionaperiodicsrsR17EnumeratedSupported
)
