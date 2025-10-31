package ies

import "rrc/utils"

// SL-FreqConfig-r16-sl-SyncPriority-r16 ::= ENUMERATED
type SlFreqconfigR16SlSyncpriorityR16 struct {
	Value utils.ENUMERATED
}

const (
	SlFreqconfigR16SlSyncpriorityR16EnumeratedNothing = iota
	SlFreqconfigR16SlSyncpriorityR16EnumeratedGnss
	SlFreqconfigR16SlSyncpriorityR16EnumeratedGnbenb
)
