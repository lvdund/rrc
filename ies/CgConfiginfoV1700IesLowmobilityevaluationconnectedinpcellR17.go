package ies

import "rrc/utils"

// CG-ConfigInfo-v1700-IEs-lowMobilityEvaluationConnectedInPCell-r17 ::= ENUMERATED
type CgConfiginfoV1700IesLowmobilityevaluationconnectedinpcellR17 struct {
	Value utils.ENUMERATED
}

const (
	CgConfiginfoV1700IesLowmobilityevaluationconnectedinpcellR17EnumeratedNothing = iota
	CgConfiginfoV1700IesLowmobilityevaluationconnectedinpcellR17EnumeratedEnabled
)
