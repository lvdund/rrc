package ies

import "rrc/utils"

// DMRS-BundlingPUSCH-Config-r17-pusch-DMRS-Bundling-r17 ::= ENUMERATED
type DmrsBundlingpuschConfigR17PuschDmrsBundlingR17 struct {
	Value utils.ENUMERATED
}

const (
	DmrsBundlingpuschConfigR17PuschDmrsBundlingR17EnumeratedNothing = iota
	DmrsBundlingpuschConfigR17PuschDmrsBundlingR17EnumeratedEnabled
)
