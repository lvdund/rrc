package ies

// SL-SyncAllowed-r14 ::= SEQUENCE
type SlSyncallowedR14 struct {
	GnssSyncR14 *bool
	EnbSyncR14  *bool
	UeSyncR14   *bool
}
