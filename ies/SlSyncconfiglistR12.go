package ies

import "rrc/utils"

// SL-SyncConfigList-r12 ::= SEQUENCE OF SL-SyncConfig-r12
// SIZE (1..maxSL-SyncConfig-r12)
type SlSyncconfiglistR12 struct {
	Value utils.Sequence[SlSyncconfigR12]
}
