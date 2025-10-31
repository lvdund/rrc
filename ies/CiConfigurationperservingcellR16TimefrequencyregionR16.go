package ies

import "rrc/utils"

// CI-ConfigurationPerServingCell-r16-timeFrequencyRegion-r16 ::= SEQUENCE
// Extensible
type CiConfigurationperservingcellR16TimefrequencyregionR16 struct {
	TimedurationforciR16    *CiConfigurationperservingcellR16TimefrequencyregionR16TimedurationforciR16
	TimegranularityforciR16 CiConfigurationperservingcellR16TimefrequencyregionR16TimegranularityforciR16
	FrequencyregionforciR16 utils.INTEGER `lb:0,ub:37949`
	DeltaoffsetR16          utils.INTEGER `lb:0,ub:2`
}
