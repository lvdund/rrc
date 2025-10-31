package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-sfn-SchemeA-DynamicSwitching-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700SfnSchemeaDynamicswitchingR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700SfnSchemeaDynamicswitchingR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700SfnSchemeaDynamicswitchingR17EnumeratedSupported
)
