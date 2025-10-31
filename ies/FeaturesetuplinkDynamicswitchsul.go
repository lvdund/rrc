package ies

import "rrc/utils"

// FeatureSetUplink-dynamicSwitchSUL ::= ENUMERATED
type FeaturesetuplinkDynamicswitchsul struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkDynamicswitchsulEnumeratedNothing = iota
	FeaturesetuplinkDynamicswitchsulEnumeratedSupported
)
