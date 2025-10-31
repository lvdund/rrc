package ies

import "rrc/utils"

// DMRS-BundlingPUCCH-Config-r17 ::= SEQUENCE
// Extensible
type DmrsBundlingpucchConfigR17 struct {
	PucchDmrsBundlingR17             *DmrsBundlingpucchConfigR17PucchDmrsBundlingR17
	PucchTimedomainwindowlengthR17   *utils.INTEGER `lb:0,ub:8`
	PucchWindowrestartR17            *DmrsBundlingpucchConfigR17PucchWindowrestartR17
	PucchFrequencyhoppingintervalR17 *DmrsBundlingpucchConfigR17PucchFrequencyhoppingintervalR17
}
