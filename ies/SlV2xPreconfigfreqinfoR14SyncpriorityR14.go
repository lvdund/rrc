package ies

import "rrc/utils"

// SL-V2X-PreconfigFreqInfo-r14-syncPriority-r14 ::= ENUMERATED
type SlV2xPreconfigfreqinfoR14SyncpriorityR14 struct {
	Value utils.ENUMERATED
}

const (
	SlV2xPreconfigfreqinfoR14SyncpriorityR14EnumeratedNothing = iota
	SlV2xPreconfigfreqinfoR14SyncpriorityR14EnumeratedGnss
	SlV2xPreconfigfreqinfoR14SyncpriorityR14EnumeratedEnb
)
