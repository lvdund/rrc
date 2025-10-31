package ies

import "rrc/utils"

// MAC-ParametersCommon-survivalTime-r17 ::= ENUMERATED
type MacParameterscommonSurvivaltimeR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonSurvivaltimeR17EnumeratedNothing = iota
	MacParameterscommonSurvivaltimeR17EnumeratedSupported
)
