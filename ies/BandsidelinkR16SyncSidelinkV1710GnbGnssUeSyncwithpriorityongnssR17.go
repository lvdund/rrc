package ies

import "rrc/utils"

// BandSidelink-r16-sync-Sidelink-v1710-gNB-GNSS-UE-SyncWithPriorityOnGNSS-r17 ::= ENUMERATED
type BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnssR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnssR17EnumeratedNothing = iota
	BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnssR17EnumeratedSupported
)
