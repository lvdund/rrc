package ies

import "rrc/utils"

// BandSidelinkPC5-r16-sl-openLoopPC-RSRP-ReportSidelink-r16 ::= ENUMERATED
type Bandsidelinkpc5R16SlOpenlooppcRsrpReportsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16SlOpenlooppcRsrpReportsidelinkR16EnumeratedNothing = iota
	Bandsidelinkpc5R16SlOpenlooppcRsrpReportsidelinkR16EnumeratedSupported
)
