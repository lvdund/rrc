package ies

import "rrc/utils"

// SL-FreqConfigCommon-r16-sl-SyncPriority-r16 ::= ENUMERATED
type SlFreqconfigcommonR16SlSyncpriorityR16 struct {
	Value utils.ENUMERATED
}

const (
	SlFreqconfigcommonR16SlSyncpriorityR16EnumeratedNothing = iota
	SlFreqconfigcommonR16SlSyncpriorityR16EnumeratedGnss
	SlFreqconfigcommonR16SlSyncpriorityR16EnumeratedGnbenb
)
