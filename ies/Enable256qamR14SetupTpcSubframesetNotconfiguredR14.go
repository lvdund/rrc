package ies

import "rrc/utils"

// Enable256QAM-r14-setup-tpc-SubframeSet-NotConfigured-r14 ::= SEQUENCE
type Enable256qamR14SetupTpcSubframesetNotconfiguredR14 struct {
	DciFormat0R14 utils.BOOLEAN
	DciFormat4R14 utils.BOOLEAN
}
