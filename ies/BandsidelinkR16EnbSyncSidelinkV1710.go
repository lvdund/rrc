package ies

import "rrc/utils"

// BandSidelink-r16-enb-sync-Sidelink-v1710 ::= ENUMERATED
type BandsidelinkR16EnbSyncSidelinkV1710 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16EnbSyncSidelinkV1710EnumeratedNothing = iota
	BandsidelinkR16EnbSyncSidelinkV1710EnumeratedSupported
)
