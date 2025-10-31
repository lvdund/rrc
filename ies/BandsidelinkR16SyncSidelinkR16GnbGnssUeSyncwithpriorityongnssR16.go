package ies

import "rrc/utils"

// BandSidelink-r16-sync-Sidelink-r16-gNB-GNSS-UE-SyncWithPriorityOnGNSS-r16 ::= ENUMERATED
type BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnssR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnssR16EnumeratedNothing = iota
	BandsidelinkR16SyncSidelinkR16GnbGnssUeSyncwithpriorityongnssR16EnumeratedSupported
)
