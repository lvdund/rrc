package ies

import "rrc/utils"

// BandSidelink-r16-sync-Sidelink-v1710-sync-GNSS-r17 ::= ENUMERATED
type BandsidelinkR16SyncSidelinkV1710SyncGnssR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SyncSidelinkV1710SyncGnssR17EnumeratedNothing = iota
	BandsidelinkR16SyncSidelinkV1710SyncGnssR17EnumeratedSupported
)
