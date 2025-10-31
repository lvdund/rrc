package ies

// BandSidelink-r16-sync-Sidelink-r16 ::= SEQUENCE
type BandsidelinkR16SyncSidelinkR16 struct {
	GnbSyncR16                           *BandsidelinkR16SyncSidelinkR16GnbSyncR16
	GnbGnssUeSyncwithpriorityongnbEnbR16 *BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnbEnbR16
	GnbGnssUeSyncwithpriorityongnssR16   *BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnssR16
}
