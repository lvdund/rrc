package ies

import "rrc/utils"

// BandSidelink-r16-sync-Sidelink-r16-gNB-Sync-r16 ::= ENUMERATED
type BandsidelinkR16SyncSidelinkR16GnbSyncR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SyncSidelinkR16GnbSyncR16EnumeratedNothing = iota
	BandsidelinkR16SyncSidelinkR16GnbSyncR16EnumeratedSupported
)
