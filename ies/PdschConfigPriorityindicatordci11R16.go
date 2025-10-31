package ies

import "rrc/utils"

// PDSCH-Config-priorityIndicatorDCI-1-1-r16 ::= ENUMERATED
type PdschConfigPriorityindicatordci11R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPriorityindicatordci11R16EnumeratedNothing = iota
	PdschConfigPriorityindicatordci11R16EnumeratedEnabled
)
