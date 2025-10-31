package ies

import "rrc/utils"

// SL-PreconfigSync-r12-syncRefMinHyst-r12 ::= ENUMERATED
type SlPreconfigsyncR12SyncrefminhystR12 struct {
	Value utils.ENUMERATED
}

const (
	SlPreconfigsyncR12SyncrefminhystR12EnumeratedNothing = iota
	SlPreconfigsyncR12SyncrefminhystR12EnumeratedDb0
	SlPreconfigsyncR12SyncrefminhystR12EnumeratedDb3
	SlPreconfigsyncR12SyncrefminhystR12EnumeratedDb6
	SlPreconfigsyncR12SyncrefminhystR12EnumeratedDb9
	SlPreconfigsyncR12SyncrefminhystR12EnumeratedDb12
)
