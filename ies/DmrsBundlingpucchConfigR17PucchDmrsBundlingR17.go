package ies

import "rrc/utils"

// DMRS-BundlingPUCCH-Config-r17-pucch-DMRS-Bundling-r17 ::= ENUMERATED
type DmrsBundlingpucchConfigR17PucchDmrsBundlingR17 struct {
	Value utils.ENUMERATED
}

const (
	DmrsBundlingpucchConfigR17PucchDmrsBundlingR17EnumeratedNothing = iota
	DmrsBundlingpucchConfigR17PucchDmrsBundlingR17EnumeratedEnabled
)
