package ies

import "rrc/utils"

// ReportCGI-useAutonomousGaps-r16 ::= ENUMERATED
type ReportcgiUseautonomousgapsR16 struct {
	Value utils.ENUMERATED
}

const (
	ReportcgiUseautonomousgapsR16EnumeratedNothing = iota
	ReportcgiUseautonomousgapsR16EnumeratedSetup
)
