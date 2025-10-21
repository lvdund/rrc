package ies

import "rrc/utils"

// ProcessingTimelineSet-r15 ::= ENUMERATED
type ProcessingtimelinesetR15 struct {
	Value utils.ENUMERATED
}

const (
	ProcessingtimelinesetR15Set1 = 0
	ProcessingtimelinesetR15Set2 = 1
)
