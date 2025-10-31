package ies

import "rrc/utils"

// TDD-Config-NB-r15-subframeAssignment-r15 ::= ENUMERATED
type TddConfigNbR15SubframeassignmentR15 struct {
	Value utils.ENUMERATED
}

const (
	TddConfigNbR15SubframeassignmentR15EnumeratedNothing = iota
	TddConfigNbR15SubframeassignmentR15EnumeratedSa1
	TddConfigNbR15SubframeassignmentR15EnumeratedSa2
	TddConfigNbR15SubframeassignmentR15EnumeratedSa3
	TddConfigNbR15SubframeassignmentR15EnumeratedSa4
	TddConfigNbR15SubframeassignmentR15EnumeratedSa5
)
