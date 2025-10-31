package ies

import "rrc/utils"

// SL-SyncConfig-r16-sl-SyncRefDiffHyst-r16 ::= ENUMERATED
type SlSyncconfigR16SlSyncrefdiffhystR16 struct {
	Value utils.ENUMERATED
}

const (
	SlSyncconfigR16SlSyncrefdiffhystR16EnumeratedNothing = iota
	SlSyncconfigR16SlSyncrefdiffhystR16EnumeratedDb0
	SlSyncconfigR16SlSyncrefdiffhystR16EnumeratedDb3
	SlSyncconfigR16SlSyncrefdiffhystR16EnumeratedDb6
	SlSyncconfigR16SlSyncrefdiffhystR16EnumeratedDb9
	SlSyncconfigR16SlSyncrefdiffhystR16EnumeratedDb12
	SlSyncconfigR16SlSyncrefdiffhystR16EnumeratedDbinf
)
