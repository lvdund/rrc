package ies

import "rrc/utils"

// SL-SyncAllowed-r16 ::= SEQUENCE
type SlSyncallowedR16 struct {
	GnssSyncR16   *utils.BOOLEAN
	GnbenbSyncR16 *utils.BOOLEAN
	UeSyncR16     *utils.BOOLEAN
}
