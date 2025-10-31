package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-longDRX-Cycle ::= ENUMERATED
type MacParametersxddDiffLongdrxCycle struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffLongdrxCycleEnumeratedNothing = iota
	MacParametersxddDiffLongdrxCycleEnumeratedSupported
)
