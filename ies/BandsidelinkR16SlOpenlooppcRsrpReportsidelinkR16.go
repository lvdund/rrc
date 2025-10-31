package ies

import "rrc/utils"

// BandSidelink-r16-sl-openLoopPC-RSRP-ReportSidelink-r16 ::= ENUMERATED
type BandsidelinkR16SlOpenlooppcRsrpReportsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlOpenlooppcRsrpReportsidelinkR16EnumeratedNothing = iota
	BandsidelinkR16SlOpenlooppcRsrpReportsidelinkR16EnumeratedSupported
)
