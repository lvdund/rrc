package ies

import "rrc/utils"

// PUSCH-Config-priorityIndicatorDCI-0-1-r16 ::= ENUMERATED
type PuschConfigPriorityindicatordci01R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigPriorityindicatordci01R16EnumeratedNothing = iota
	PuschConfigPriorityindicatordci01R16EnumeratedEnabled
)
