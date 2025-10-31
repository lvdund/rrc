package ies

import "rrc/utils"

// TDD-Config-subframeAssignment ::= ENUMERATED
type TddConfigSubframeassignment struct {
	Value utils.ENUMERATED
}

const (
	TddConfigSubframeassignmentEnumeratedNothing = iota
	TddConfigSubframeassignmentEnumeratedSa0
	TddConfigSubframeassignmentEnumeratedSa1
	TddConfigSubframeassignmentEnumeratedSa2
	TddConfigSubframeassignmentEnumeratedSa3
	TddConfigSubframeassignmentEnumeratedSa4
	TddConfigSubframeassignmentEnumeratedSa5
	TddConfigSubframeassignmentEnumeratedSa6
)
