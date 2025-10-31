package ies

import "rrc/utils"

// BandSidelink-r16-sync-Sidelink-r16-gNB-GNSS-UE-SyncWithPriorityOnGNB-ENB-r16 ::= ENUMERATED
type BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnbEnbR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnbEnbR16EnumeratedNothing = iota
	BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnbEnbR16EnumeratedSupported
)
