package ies

import "rrc/utils"

// PDSCH-Config-priorityIndicatorDCI-4-2-r17 ::= ENUMERATED
type PdschConfigPriorityindicatordci42R17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPriorityindicatordci42R17EnumeratedNothing = iota
	PdschConfigPriorityindicatordci42R17EnumeratedEnabled
)
