package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-sfn-SchemeB-DynamicSwitching-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700SfnSchemebDynamicswitchingR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700SfnSchemebDynamicswitchingR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700SfnSchemebDynamicswitchingR17EnumeratedSupported
)
