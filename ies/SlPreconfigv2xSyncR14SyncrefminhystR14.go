package ies

import "rrc/utils"

// SL-PreconfigV2X-Sync-r14-syncRefMinHyst-r14 ::= ENUMERATED
type SlPreconfigv2xSyncR14SyncrefminhystR14 struct {
	Value utils.ENUMERATED
}

const (
	SlPreconfigv2xSyncR14SyncrefminhystR14EnumeratedNothing = iota
	SlPreconfigv2xSyncR14SyncrefminhystR14EnumeratedDb0
	SlPreconfigv2xSyncR14SyncrefminhystR14EnumeratedDb3
	SlPreconfigv2xSyncR14SyncrefminhystR14EnumeratedDb6
	SlPreconfigv2xSyncR14SyncrefminhystR14EnumeratedDb9
	SlPreconfigv2xSyncR14SyncrefminhystR14EnumeratedDb12
)
