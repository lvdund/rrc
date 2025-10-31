package ies

import "rrc/utils"

// DMRS-BundlingPUSCH-Config-r17-pusch-WindowRestart-r17 ::= ENUMERATED
type DmrsBundlingpuschConfigR17PuschWindowrestartR17 struct {
	Value utils.ENUMERATED
}

const (
	DmrsBundlingpuschConfigR17PuschWindowrestartR17EnumeratedNothing = iota
	DmrsBundlingpuschConfigR17PuschWindowrestartR17EnumeratedEnabled
)
