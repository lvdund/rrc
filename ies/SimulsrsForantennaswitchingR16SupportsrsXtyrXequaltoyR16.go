package ies

import "rrc/utils"

// SimulSRS-ForAntennaSwitching-r16-supportSRS-xTyR-xEqualToY-r16 ::= ENUMERATED
type SimulsrsForantennaswitchingR16SupportsrsXtyrXequaltoyR16 struct {
	Value utils.ENUMERATED
}

const (
	SimulsrsForantennaswitchingR16SupportsrsXtyrXequaltoyR16EnumeratedNothing = iota
	SimulsrsForantennaswitchingR16SupportsrsXtyrXequaltoyR16EnumeratedSupported
)
