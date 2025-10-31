package ies

import "rrc/utils"

// BandSidelink-r16-sync-Sidelink-v1710-gNB-GNSS-UE-SyncWithPriorityOnGNB-ENB-r17 ::= ENUMERATED
type BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnbEnbR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnbEnbR17EnumeratedNothing = iota
	BandsidelinkR16SyncSidelinkV1710GnbGnssUeSyncwithpriorityongnbEnbR17EnumeratedSupported
)
