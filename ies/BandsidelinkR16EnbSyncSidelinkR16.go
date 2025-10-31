package ies

import "rrc/utils"

// BandSidelink-r16-enb-sync-Sidelink-r16 ::= ENUMERATED
type BandsidelinkR16EnbSyncSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16EnbSyncSidelinkR16EnumeratedNothing = iota
	BandsidelinkR16EnbSyncSidelinkR16EnumeratedSupported
)
