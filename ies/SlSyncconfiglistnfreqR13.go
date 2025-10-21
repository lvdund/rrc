package ies

import "rrc/utils"

// SL-SyncConfigListNFreq-r13 ::= SEQUENCE OF SL-SyncConfigNFreq-r13
// SIZE (1..maxSL-SyncConfig-r12)
type SlSyncconfiglistnfreqR13 struct {
	Value utils.Sequence[SlSyncconfignfreqR13]
}
