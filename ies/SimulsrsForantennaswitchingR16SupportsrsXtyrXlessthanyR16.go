package ies

import "rrc/utils"

// SimulSRS-ForAntennaSwitching-r16-supportSRS-xTyR-xLessThanY-r16 ::= ENUMERATED
type SimulsrsForantennaswitchingR16SupportsrsXtyrXlessthanyR16 struct {
	Value utils.ENUMERATED
}

const (
	SimulsrsForantennaswitchingR16SupportsrsXtyrXlessthanyR16EnumeratedNothing = iota
	SimulsrsForantennaswitchingR16SupportsrsXtyrXlessthanyR16EnumeratedSupported
)
