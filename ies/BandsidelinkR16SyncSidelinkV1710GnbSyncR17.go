package ies

import "rrc/utils"

// BandSidelink-r16-sync-Sidelink-v1710-gNB-Sync-r17 ::= ENUMERATED
type BandsidelinkR16SyncSidelinkV1710GnbSyncR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SyncSidelinkV1710GnbSyncR17EnumeratedNothing = iota
	BandsidelinkR16SyncSidelinkV1710GnbSyncR17EnumeratedSupported
)
