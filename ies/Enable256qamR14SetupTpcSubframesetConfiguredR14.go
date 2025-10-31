package ies

import "rrc/utils"

// Enable256QAM-r14-setup-tpc-SubframeSet-Configured-r14 ::= SEQUENCE
type Enable256qamR14SetupTpcSubframesetConfiguredR14 struct {
	Subframeset1DciFormat0R14 utils.BOOLEAN
	Subframeset1DciFormat4R14 utils.BOOLEAN
	Subframeset2DciFormat0R14 utils.BOOLEAN
	Subframeset2DciFormat4R14 utils.BOOLEAN
}
