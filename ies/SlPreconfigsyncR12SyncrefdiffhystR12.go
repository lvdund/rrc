package ies

import "rrc/utils"

// SL-PreconfigSync-r12-syncRefDiffHyst-r12 ::= ENUMERATED
type SlPreconfigsyncR12SyncrefdiffhystR12 struct {
	Value utils.ENUMERATED
}

const (
	SlPreconfigsyncR12SyncrefdiffhystR12EnumeratedNothing = iota
	SlPreconfigsyncR12SyncrefdiffhystR12EnumeratedDb0
	SlPreconfigsyncR12SyncrefdiffhystR12EnumeratedDb3
	SlPreconfigsyncR12SyncrefdiffhystR12EnumeratedDb6
	SlPreconfigsyncR12SyncrefdiffhystR12EnumeratedDb9
	SlPreconfigsyncR12SyncrefdiffhystR12EnumeratedDb12
	SlPreconfigsyncR12SyncrefdiffhystR12EnumeratedDbinf
)
