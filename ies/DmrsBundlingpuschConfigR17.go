package ies

import "rrc/utils"

// DMRS-BundlingPUSCH-Config-r17 ::= SEQUENCE
// Extensible
type DmrsBundlingpuschConfigR17 struct {
	PuschDmrsBundlingR17             *DmrsBundlingpuschConfigR17PuschDmrsBundlingR17
	PuschTimedomainwindowlengthR17   *utils.INTEGER `lb:0,ub:32`
	PuschWindowrestartR17            *DmrsBundlingpuschConfigR17PuschWindowrestartR17
	PuschFrequencyhoppingintervalR17 *DmrsBundlingpuschConfigR17PuschFrequencyhoppingintervalR17
}
