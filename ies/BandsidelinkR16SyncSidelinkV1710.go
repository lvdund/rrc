package ies

// BandSidelink-r16-sync-Sidelink-v1710 ::= SEQUENCE
type BandsidelinkR16SyncSidelinkV1710 struct {
	SyncGnssR17                          *BandsidelinkR16SyncSidelinkV1710SyncGnssR17
	GnbSyncR17                           *BandsidelinkR16SyncSidelinkV1710GnbSyncR17
	GnbGnssUeSyncwithpriorityongnbEnbR17 *BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnbEnbR17
	GnbGnssUeSyncwithpriorityongnssR17   *BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnssR17
}
