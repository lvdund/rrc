package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-shortDRX-Cycle ::= ENUMERATED
type MacParametersxddDiffShortdrxCycle struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffShortdrxCycleEnumeratedNothing = iota
	MacParametersxddDiffShortdrxCycleEnumeratedSupported
)
