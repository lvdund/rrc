package ies

import "rrc/utils"

// ReportCGI-EUTRA-useAutonomousGaps-r16 ::= ENUMERATED
type ReportcgiEutraUseautonomousgapsR16 struct {
	Value utils.ENUMERATED
}

const (
	ReportcgiEutraUseautonomousgapsR16EnumeratedNothing = iota
	ReportcgiEutraUseautonomousgapsR16EnumeratedSetup
)
