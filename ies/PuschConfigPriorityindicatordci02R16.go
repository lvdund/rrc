package ies

import "rrc/utils"

// PUSCH-Config-priorityIndicatorDCI-0-2-r16 ::= ENUMERATED
type PuschConfigPriorityindicatordci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigPriorityindicatordci02R16EnumeratedNothing = iota
	PuschConfigPriorityindicatordci02R16EnumeratedEnabled
)
