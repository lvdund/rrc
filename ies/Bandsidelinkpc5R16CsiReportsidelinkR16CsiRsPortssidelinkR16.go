package ies

import "rrc/utils"

// BandSidelinkPC5-r16-csi-ReportSidelink-r16-csi-RS-PortsSidelink-r16 ::= ENUMERATED
type Bandsidelinkpc5R16CsiReportsidelinkR16CsiRsPortssidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16CsiReportsidelinkR16CsiRsPortssidelinkR16EnumeratedNothing = iota
	Bandsidelinkpc5R16CsiReportsidelinkR16CsiRsPortssidelinkR16EnumeratedP1
	Bandsidelinkpc5R16CsiReportsidelinkR16CsiRsPortssidelinkR16EnumeratedP2
)
