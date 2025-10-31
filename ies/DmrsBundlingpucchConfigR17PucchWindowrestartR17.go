package ies

import "rrc/utils"

// DMRS-BundlingPUCCH-Config-r17-pucch-WindowRestart-r17 ::= ENUMERATED
type DmrsBundlingpucchConfigR17PucchWindowrestartR17 struct {
	Value utils.ENUMERATED
}

const (
	DmrsBundlingpucchConfigR17PucchWindowrestartR17EnumeratedNothing = iota
	DmrsBundlingpucchConfigR17PucchWindowrestartR17EnumeratedEnabled
)
