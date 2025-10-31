package ies

import "rrc/utils"

// ProcessingTimelineSet-r15 ::= ENUMERATED
type ProcessingtimelinesetR15 struct {
	Value utils.ENUMERATED
}

const (
	ProcessingtimelinesetR15EnumeratedNothing = iota
	ProcessingtimelinesetR15EnumeratedSet1
	ProcessingtimelinesetR15EnumeratedSet2
)
