package ies

import "rrc/utils"

// MAC-ParametersCommon-simultaneousSR-PUSCH-DiffPUCCH-groups-r17 ::= ENUMERATED
type MacParameterscommonSimultaneoussrPuschDiffpucchGroupsR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonSimultaneoussrPuschDiffpucchGroupsR17EnumeratedNothing = iota
	MacParameterscommonSimultaneoussrPuschDiffpucchGroupsR17EnumeratedSupported
)
