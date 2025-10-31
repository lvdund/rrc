package ies

import "rrc/utils"

// PDSCH-Config-priorityIndicatorDCI-1-2-r16 ::= ENUMERATED
type PdschConfigPriorityindicatordci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPriorityindicatordci12R16EnumeratedNothing = iota
	PdschConfigPriorityindicatordci12R16EnumeratedEnabled
)
