package ies

import "rrc/utils"

// SimulSRS-ForAntennaSwitching-r16-supportSRS-AntennaSwitching-r16 ::= ENUMERATED
type SimulsrsForantennaswitchingR16SupportsrsAntennaswitchingR16 struct {
	Value utils.ENUMERATED
}

const (
	SimulsrsForantennaswitchingR16SupportsrsAntennaswitchingR16EnumeratedNothing = iota
	SimulsrsForantennaswitchingR16SupportsrsAntennaswitchingR16EnumeratedSupported
)
