package ies

import "rrc/utils"

// SubframeAssignment-r15 ::= ENUMERATED
type SubframeassignmentR15 struct {
	Value utils.ENUMERATED
}

const (
	SubframeassignmentR15EnumeratedNothing = iota
	SubframeassignmentR15EnumeratedSa0
	SubframeassignmentR15EnumeratedSa1
	SubframeassignmentR15EnumeratedSa2
	SubframeassignmentR15EnumeratedSa3
	SubframeassignmentR15EnumeratedSa4
	SubframeassignmentR15EnumeratedSa5
	SubframeassignmentR15EnumeratedSa6
)
