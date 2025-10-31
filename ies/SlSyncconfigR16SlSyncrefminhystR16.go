package ies

import "rrc/utils"

// SL-SyncConfig-r16-sl-SyncRefMinHyst-r16 ::= ENUMERATED
type SlSyncconfigR16SlSyncrefminhystR16 struct {
	Value utils.ENUMERATED
}

const (
	SlSyncconfigR16SlSyncrefminhystR16EnumeratedNothing = iota
	SlSyncconfigR16SlSyncrefminhystR16EnumeratedDb0
	SlSyncconfigR16SlSyncrefminhystR16EnumeratedDb3
	SlSyncconfigR16SlSyncrefminhystR16EnumeratedDb6
	SlSyncconfigR16SlSyncrefminhystR16EnumeratedDb9
	SlSyncconfigR16SlSyncrefminhystR16EnumeratedDb12
)
